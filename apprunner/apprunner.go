package apprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/apprunner/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apprunner_auto_scaling_configuration_version.html aws_apprunner_auto_scaling_configuration_version}.
type ApprunnerAutoScalingConfigurationVersion interface {
	cdktf.TerraformResource
	Arn() *string
	AutoScalingConfigurationName() *string
	SetAutoScalingConfigurationName(val *string)
	AutoScalingConfigurationNameInput() *string
	AutoScalingConfigurationRevision() *float64
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Latest() interface{}
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxConcurrency() *float64
	SetMaxConcurrency(val *float64)
	MaxConcurrencyInput() *float64
	MaxSize() *float64
	SetMaxSize(val *float64)
	MaxSizeInput() *float64
	MinSize() *float64
	SetMinSize(val *float64)
	MinSizeInput() *float64
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
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
	ResetMaxConcurrency()
	ResetMaxSize()
	ResetMinSize()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApprunnerAutoScalingConfigurationVersion
type jsiiProxy_ApprunnerAutoScalingConfigurationVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) AutoScalingConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) AutoScalingConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) AutoScalingConfigurationRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScalingConfigurationRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Latest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) MaxConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) MaxConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apprunner_auto_scaling_configuration_version.html aws_apprunner_auto_scaling_configuration_version} Resource.
func NewApprunnerAutoScalingConfigurationVersion(scope constructs.Construct, id *string, config *ApprunnerAutoScalingConfigurationVersionConfig) ApprunnerAutoScalingConfigurationVersion {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerAutoScalingConfigurationVersion{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerAutoScalingConfigurationVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apprunner_auto_scaling_configuration_version.html aws_apprunner_auto_scaling_configuration_version} Resource.
func NewApprunnerAutoScalingConfigurationVersion_Override(a ApprunnerAutoScalingConfigurationVersion, scope constructs.Construct, id *string, config *ApprunnerAutoScalingConfigurationVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerAutoScalingConfigurationVersion",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetAutoScalingConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"autoScalingConfigurationName",
		val,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetMaxConcurrency(val *float64) {
	_jsii_.Set(
		j,
		"maxConcurrency",
		val,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetMaxSize(val *float64) {
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetMinSize(val *float64) {
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SetTagsAll(val interface{}) {
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
func ApprunnerAutoScalingConfigurationVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.AppRunner.ApprunnerAutoScalingConfigurationVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApprunnerAutoScalingConfigurationVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.AppRunner.ApprunnerAutoScalingConfigurationVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ResetMaxConcurrency() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ResetMaxSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ResetMinSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMinSize",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ToMetadata() interface{} {
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
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ToString() *string {
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
func (a *jsiiProxy_ApprunnerAutoScalingConfigurationVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApprunnerAutoScalingConfigurationVersionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_auto_scaling_configuration_version.html#auto_scaling_configuration_name ApprunnerAutoScalingConfigurationVersion#auto_scaling_configuration_name}.
	AutoScalingConfigurationName *string `json:"autoScalingConfigurationName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_auto_scaling_configuration_version.html#max_concurrency ApprunnerAutoScalingConfigurationVersion#max_concurrency}.
	MaxConcurrency *float64 `json:"maxConcurrency"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_auto_scaling_configuration_version.html#max_size ApprunnerAutoScalingConfigurationVersion#max_size}.
	MaxSize *float64 `json:"maxSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_auto_scaling_configuration_version.html#min_size ApprunnerAutoScalingConfigurationVersion#min_size}.
	MinSize *float64 `json:"minSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_auto_scaling_configuration_version.html#tags ApprunnerAutoScalingConfigurationVersion#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_auto_scaling_configuration_version.html#tags_all ApprunnerAutoScalingConfigurationVersion#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apprunner_connection.html aws_apprunner_connection}.
type ApprunnerConnection interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConnectionName() *string
	SetConnectionName(val *string)
	ConnectionNameInput() *string
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
	ProviderType() *string
	SetProviderType(val *string)
	ProviderTypeInput() *string
	RawOverrides() interface{}
	Status() *string
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
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApprunnerConnection
type jsiiProxy_ApprunnerConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApprunnerConnection) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) ConnectionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) ProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) ProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"providerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apprunner_connection.html aws_apprunner_connection} Resource.
func NewApprunnerConnection(scope constructs.Construct, id *string, config *ApprunnerConnectionConfig) ApprunnerConnection {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerConnection{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apprunner_connection.html aws_apprunner_connection} Resource.
func NewApprunnerConnection_Override(a ApprunnerConnection, scope constructs.Construct, id *string, config *ApprunnerConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerConnection",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApprunnerConnection) SetConnectionName(val *string) {
	_jsii_.Set(
		j,
		"connectionName",
		val,
	)
}

func (j *jsiiProxy_ApprunnerConnection) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApprunnerConnection) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApprunnerConnection) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApprunnerConnection) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApprunnerConnection) SetProviderType(val *string) {
	_jsii_.Set(
		j,
		"providerType",
		val,
	)
}

func (j *jsiiProxy_ApprunnerConnection) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApprunnerConnection) SetTagsAll(val interface{}) {
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
func ApprunnerConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.AppRunner.ApprunnerConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApprunnerConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.AppRunner.ApprunnerConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApprunnerConnection) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerConnection) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerConnection) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApprunnerConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerConnection) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerConnection) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerConnection) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_ApprunnerConnection) ToMetadata() interface{} {
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
func (a *jsiiProxy_ApprunnerConnection) ToString() *string {
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
func (a *jsiiProxy_ApprunnerConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApprunnerConnectionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_connection.html#connection_name ApprunnerConnection#connection_name}.
	ConnectionName *string `json:"connectionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_connection.html#provider_type ApprunnerConnection#provider_type}.
	ProviderType *string `json:"providerType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_connection.html#tags ApprunnerConnection#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_connection.html#tags_all ApprunnerConnection#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apprunner_custom_domain_association.html aws_apprunner_custom_domain_association}.
type ApprunnerCustomDomainAssociation interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DnsTarget() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EnableWwwSubdomain() interface{}
	SetEnableWwwSubdomain(val interface{})
	EnableWwwSubdomainInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServiceArn() *string
	SetServiceArn(val *string)
	ServiceArnInput() *string
	Status() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	CertificateValidationRecords(index *string) ApprunnerCustomDomainAssociationCertificateValidationRecords
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetEnableWwwSubdomain()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApprunnerCustomDomainAssociation
type jsiiProxy_ApprunnerCustomDomainAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) DnsTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) EnableWwwSubdomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWwwSubdomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) EnableWwwSubdomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWwwSubdomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) ServiceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apprunner_custom_domain_association.html aws_apprunner_custom_domain_association} Resource.
func NewApprunnerCustomDomainAssociation(scope constructs.Construct, id *string, config *ApprunnerCustomDomainAssociationConfig) ApprunnerCustomDomainAssociation {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerCustomDomainAssociation{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerCustomDomainAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apprunner_custom_domain_association.html aws_apprunner_custom_domain_association} Resource.
func NewApprunnerCustomDomainAssociation_Override(a ApprunnerCustomDomainAssociation, scope constructs.Construct, id *string, config *ApprunnerCustomDomainAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerCustomDomainAssociation",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) SetEnableWwwSubdomain(val interface{}) {
	_jsii_.Set(
		j,
		"enableWwwSubdomain",
		val,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociation) SetServiceArn(val *string) {
	_jsii_.Set(
		j,
		"serviceArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApprunnerCustomDomainAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.AppRunner.ApprunnerCustomDomainAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApprunnerCustomDomainAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.AppRunner.ApprunnerCustomDomainAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApprunnerCustomDomainAssociation) CertificateValidationRecords(index *string) ApprunnerCustomDomainAssociationCertificateValidationRecords {
	var returns ApprunnerCustomDomainAssociationCertificateValidationRecords

	_jsii_.Invoke(
		a,
		"certificateValidationRecords",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApprunnerCustomDomainAssociation) ResetEnableWwwSubdomain() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableWwwSubdomain",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerCustomDomainAssociation) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) ToMetadata() interface{} {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) ToString() *string {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApprunnerCustomDomainAssociationCertificateValidationRecords interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Name() *string
	Status() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Type() *string
	Value() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for ApprunnerCustomDomainAssociationCertificateValidationRecords
type jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

// Experimental.
func NewApprunnerCustomDomainAssociationCertificateValidationRecords(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) ApprunnerCustomDomainAssociationCertificateValidationRecords {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerCustomDomainAssociationCertificateValidationRecords",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewApprunnerCustomDomainAssociationCertificateValidationRecords_Override(a ApprunnerCustomDomainAssociationCertificateValidationRecords, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerCustomDomainAssociationCertificateValidationRecords",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		a,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerCustomDomainAssociationCertificateValidationRecords) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type ApprunnerCustomDomainAssociationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_custom_domain_association.html#domain_name ApprunnerCustomDomainAssociation#domain_name}.
	DomainName *string `json:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_custom_domain_association.html#service_arn ApprunnerCustomDomainAssociation#service_arn}.
	ServiceArn *string `json:"serviceArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_custom_domain_association.html#enable_www_subdomain ApprunnerCustomDomainAssociation#enable_www_subdomain}.
	EnableWwwSubdomain interface{} `json:"enableWwwSubdomain"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html aws_apprunner_service}.
type ApprunnerService interface {
	cdktf.TerraformResource
	Arn() *string
	AutoScalingConfigurationArn() *string
	SetAutoScalingConfigurationArn(val *string)
	AutoScalingConfigurationArnInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EncryptionConfiguration() ApprunnerServiceEncryptionConfigurationOutputReference
	EncryptionConfigurationInput() *ApprunnerServiceEncryptionConfiguration
	Fqn() *string
	FriendlyUniqueId() *string
	HealthCheckConfiguration() ApprunnerServiceHealthCheckConfigurationOutputReference
	HealthCheckConfigurationInput() *ApprunnerServiceHealthCheckConfiguration
	Id() *string
	InstanceConfiguration() ApprunnerServiceInstanceConfigurationOutputReference
	InstanceConfigurationInput() *ApprunnerServiceInstanceConfiguration
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServiceId() *string
	ServiceName() *string
	SetServiceName(val *string)
	ServiceNameInput() *string
	ServiceUrl() *string
	SourceConfiguration() ApprunnerServiceSourceConfigurationOutputReference
	SourceConfigurationInput() *ApprunnerServiceSourceConfiguration
	Status() *string
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
	PutEncryptionConfiguration(value *ApprunnerServiceEncryptionConfiguration)
	PutHealthCheckConfiguration(value *ApprunnerServiceHealthCheckConfiguration)
	PutInstanceConfiguration(value *ApprunnerServiceInstanceConfiguration)
	PutSourceConfiguration(value *ApprunnerServiceSourceConfiguration)
	ResetAutoScalingConfigurationArn()
	ResetEncryptionConfiguration()
	ResetHealthCheckConfiguration()
	ResetInstanceConfiguration()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ApprunnerService
type jsiiProxy_ApprunnerService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApprunnerService) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) AutoScalingConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) AutoScalingConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingConfigurationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) EncryptionConfiguration() ApprunnerServiceEncryptionConfigurationOutputReference {
	var returns ApprunnerServiceEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"encryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) EncryptionConfigurationInput() *ApprunnerServiceEncryptionConfiguration {
	var returns *ApprunnerServiceEncryptionConfiguration
	_jsii_.Get(
		j,
		"encryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) HealthCheckConfiguration() ApprunnerServiceHealthCheckConfigurationOutputReference {
	var returns ApprunnerServiceHealthCheckConfigurationOutputReference
	_jsii_.Get(
		j,
		"healthCheckConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) HealthCheckConfigurationInput() *ApprunnerServiceHealthCheckConfiguration {
	var returns *ApprunnerServiceHealthCheckConfiguration
	_jsii_.Get(
		j,
		"healthCheckConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) InstanceConfiguration() ApprunnerServiceInstanceConfigurationOutputReference {
	var returns ApprunnerServiceInstanceConfigurationOutputReference
	_jsii_.Get(
		j,
		"instanceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) InstanceConfigurationInput() *ApprunnerServiceInstanceConfiguration {
	var returns *ApprunnerServiceInstanceConfiguration
	_jsii_.Get(
		j,
		"instanceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) ServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) ServiceUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) SourceConfiguration() ApprunnerServiceSourceConfigurationOutputReference {
	var returns ApprunnerServiceSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"sourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) SourceConfigurationInput() *ApprunnerServiceSourceConfiguration {
	var returns *ApprunnerServiceSourceConfiguration
	_jsii_.Get(
		j,
		"sourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html aws_apprunner_service} Resource.
func NewApprunnerService(scope constructs.Construct, id *string, config *ApprunnerServiceConfig) ApprunnerService {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerService{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html aws_apprunner_service} Resource.
func NewApprunnerService_Override(a ApprunnerService, scope constructs.Construct, id *string, config *ApprunnerServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerService",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApprunnerService) SetAutoScalingConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"autoScalingConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_ApprunnerService) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApprunnerService) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApprunnerService) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApprunnerService) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApprunnerService) SetServiceName(val *string) {
	_jsii_.Set(
		j,
		"serviceName",
		val,
	)
}

func (j *jsiiProxy_ApprunnerService) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ApprunnerService) SetTagsAll(val interface{}) {
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
func ApprunnerService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.AppRunner.ApprunnerService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApprunnerService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.AppRunner.ApprunnerService",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_ApprunnerService) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerService) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerService) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerService) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerService) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApprunnerService) PutEncryptionConfiguration(value *ApprunnerServiceEncryptionConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerService) PutHealthCheckConfiguration(value *ApprunnerServiceHealthCheckConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putHealthCheckConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerService) PutInstanceConfiguration(value *ApprunnerServiceInstanceConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putInstanceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerService) PutSourceConfiguration(value *ApprunnerServiceSourceConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putSourceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerService) ResetAutoScalingConfigurationArn() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoScalingConfigurationArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerService) ResetEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerService) ResetHealthCheckConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerService) ResetInstanceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceConfiguration",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_ApprunnerService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerService) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerService) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerService) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_ApprunnerService) ToMetadata() interface{} {
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
func (a *jsiiProxy_ApprunnerService) ToString() *string {
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
func (a *jsiiProxy_ApprunnerService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApprunnerServiceConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#service_name ApprunnerService#service_name}.
	ServiceName *string `json:"serviceName"`
	// source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#source_configuration ApprunnerService#source_configuration}
	SourceConfiguration *ApprunnerServiceSourceConfiguration `json:"sourceConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#auto_scaling_configuration_arn ApprunnerService#auto_scaling_configuration_arn}.
	AutoScalingConfigurationArn *string `json:"autoScalingConfigurationArn"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#encryption_configuration ApprunnerService#encryption_configuration}
	EncryptionConfiguration *ApprunnerServiceEncryptionConfiguration `json:"encryptionConfiguration"`
	// health_check_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#health_check_configuration ApprunnerService#health_check_configuration}
	HealthCheckConfiguration *ApprunnerServiceHealthCheckConfiguration `json:"healthCheckConfiguration"`
	// instance_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#instance_configuration ApprunnerService#instance_configuration}
	InstanceConfiguration *ApprunnerServiceInstanceConfiguration `json:"instanceConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#tags ApprunnerService#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#tags_all ApprunnerService#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type ApprunnerServiceEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#kms_key ApprunnerService#kms_key}.
	KmsKey *string `json:"kmsKey"`
}

type ApprunnerServiceEncryptionConfigurationOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
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

// The jsii proxy struct for ApprunnerServiceEncryptionConfigurationOutputReference
type jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApprunnerServiceEncryptionConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceEncryptionConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceEncryptionConfigurationOutputReference_Override(a ApprunnerServiceEncryptionConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceEncryptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) SetKmsKey(val *string) {
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceEncryptionConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type ApprunnerServiceHealthCheckConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#healthy_threshold ApprunnerService#healthy_threshold}.
	HealthyThreshold *float64 `json:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#interval ApprunnerService#interval}.
	Interval *float64 `json:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#path ApprunnerService#path}.
	Path *string `json:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#protocol ApprunnerService#protocol}.
	Protocol *string `json:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#timeout ApprunnerService#timeout}.
	Timeout *float64 `json:"timeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#unhealthy_threshold ApprunnerService#unhealthy_threshold}.
	UnhealthyThreshold *float64 `json:"unhealthyThreshold"`
}

type ApprunnerServiceHealthCheckConfigurationOutputReference interface {
	cdktf.ComplexObject
	HealthyThreshold() *float64
	SetHealthyThreshold(val *float64)
	HealthyThresholdInput() *float64
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	UnhealthyThreshold() *float64
	SetUnhealthyThreshold(val *float64)
	UnhealthyThresholdInput() *float64
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetHealthyThreshold()
	ResetInterval()
	ResetPath()
	ResetProtocol()
	ResetTimeout()
	ResetUnhealthyThreshold()
}

// The jsii proxy struct for ApprunnerServiceHealthCheckConfigurationOutputReference
type jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}

func NewApprunnerServiceHealthCheckConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceHealthCheckConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceHealthCheckConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceHealthCheckConfigurationOutputReference_Override(a ApprunnerServiceHealthCheckConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceHealthCheckConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) SetHealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) SetInterval(val *float64) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) SetUnhealthyThreshold(val *float64) {
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceHealthCheckConfigurationOutputReference) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

type ApprunnerServiceInstanceConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#instance_role_arn ApprunnerService#instance_role_arn}.
	InstanceRoleArn *string `json:"instanceRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#cpu ApprunnerService#cpu}.
	Cpu *string `json:"cpu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#memory ApprunnerService#memory}.
	Memory *string `json:"memory"`
}

type ApprunnerServiceInstanceConfigurationOutputReference interface {
	cdktf.ComplexObject
	Cpu() *string
	SetCpu(val *string)
	CpuInput() *string
	InstanceRoleArn() *string
	SetInstanceRoleArn(val *string)
	InstanceRoleArnInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Memory() *string
	SetMemory(val *string)
	MemoryInput() *string
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
	ResetCpu()
	ResetMemory()
}

// The jsii proxy struct for ApprunnerServiceInstanceConfigurationOutputReference
type jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) CpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) InstanceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) InstanceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) MemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApprunnerServiceInstanceConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceInstanceConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceInstanceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceInstanceConfigurationOutputReference_Override(a ApprunnerServiceInstanceConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceInstanceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) SetCpu(val *string) {
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) SetInstanceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"instanceRoleArn",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) SetMemory(val *string) {
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) ResetCpu() {
	_jsii_.InvokeVoid(
		a,
		"resetCpu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceInstanceConfigurationOutputReference) ResetMemory() {
	_jsii_.InvokeVoid(
		a,
		"resetMemory",
		nil, // no parameters
	)
}

type ApprunnerServiceSourceConfiguration struct {
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#authentication_configuration ApprunnerService#authentication_configuration}
	AuthenticationConfiguration *ApprunnerServiceSourceConfigurationAuthenticationConfiguration `json:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#auto_deployments_enabled ApprunnerService#auto_deployments_enabled}.
	AutoDeploymentsEnabled interface{} `json:"autoDeploymentsEnabled"`
	// code_repository block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#code_repository ApprunnerService#code_repository}
	CodeRepository *ApprunnerServiceSourceConfigurationCodeRepository `json:"codeRepository"`
	// image_repository block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#image_repository ApprunnerService#image_repository}
	ImageRepository *ApprunnerServiceSourceConfigurationImageRepository `json:"imageRepository"`
}

type ApprunnerServiceSourceConfigurationAuthenticationConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#access_role_arn ApprunnerService#access_role_arn}.
	AccessRoleArn *string `json:"accessRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#connection_arn ApprunnerService#connection_arn}.
	ConnectionArn *string `json:"connectionArn"`
}

type ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference interface {
	cdktf.ComplexObject
	AccessRoleArn() *string
	SetAccessRoleArn(val *string)
	AccessRoleArnInput() *string
	ConnectionArn() *string
	SetConnectionArn(val *string)
	ConnectionArnInput() *string
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
	ResetAccessRoleArn()
	ResetConnectionArn()
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) AccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) AccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) ConnectionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference_Override(a ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) SetAccessRoleArn(val *string) {
	_jsii_.Set(
		j,
		"accessRoleArn",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) SetConnectionArn(val *string) {
	_jsii_.Set(
		j,
		"connectionArn",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) ResetAccessRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference) ResetConnectionArn() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionArn",
		nil, // no parameters
	)
}

type ApprunnerServiceSourceConfigurationCodeRepository struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#repository_url ApprunnerService#repository_url}.
	RepositoryUrl *string `json:"repositoryUrl"`
	// source_code_version block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#source_code_version ApprunnerService#source_code_version}
	SourceCodeVersion *ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion `json:"sourceCodeVersion"`
	// code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#code_configuration ApprunnerService#code_configuration}
	CodeConfiguration *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration `json:"codeConfiguration"`
}

type ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#configuration_source ApprunnerService#configuration_source}.
	ConfigurationSource *string `json:"configurationSource"`
	// code_configuration_values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#code_configuration_values ApprunnerService#code_configuration_values}
	CodeConfigurationValues *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues `json:"codeConfigurationValues"`
}

type ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#runtime ApprunnerService#runtime}.
	Runtime *string `json:"runtime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#build_command ApprunnerService#build_command}.
	BuildCommand *string `json:"buildCommand"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#port ApprunnerService#port}.
	Port *string `json:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#runtime_environment_variables ApprunnerService#runtime_environment_variables}.
	RuntimeEnvironmentVariables interface{} `json:"runtimeEnvironmentVariables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#start_command ApprunnerService#start_command}.
	StartCommand *string `json:"startCommand"`
}

type ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference interface {
	cdktf.ComplexObject
	BuildCommand() *string
	SetBuildCommand(val *string)
	BuildCommandInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *string
	SetPort(val *string)
	PortInput() *string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeEnvironmentVariables() interface{}
	SetRuntimeEnvironmentVariables(val interface{})
	RuntimeEnvironmentVariablesInput() interface{}
	RuntimeInput() *string
	StartCommand() *string
	SetStartCommand(val *string)
	StartCommandInput() *string
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
	ResetBuildCommand()
	ResetPort()
	ResetRuntimeEnvironmentVariables()
	ResetStartCommand()
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) BuildCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) BuildCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) RuntimeEnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) RuntimeEnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) StartCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) StartCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference_Override(a ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) SetBuildCommand(val *string) {
	_jsii_.Set(
		j,
		"buildCommand",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) SetPort(val *string) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) SetRuntime(val *string) {
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) SetRuntimeEnvironmentVariables(val interface{}) {
	_jsii_.Set(
		j,
		"runtimeEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) SetStartCommand(val *string) {
	_jsii_.Set(
		j,
		"startCommand",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) ResetBuildCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) ResetRuntimeEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference) ResetStartCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetStartCommand",
		nil, // no parameters
	)
}

type ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference interface {
	cdktf.ComplexObject
	CodeConfigurationValues() ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference
	CodeConfigurationValuesInput() *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues
	ConfigurationSource() *string
	SetConfigurationSource(val *string)
	ConfigurationSourceInput() *string
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
	PutCodeConfigurationValues(value *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues)
	ResetCodeConfigurationValues()
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) CodeConfigurationValues() ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference {
	var returns ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValuesOutputReference
	_jsii_.Get(
		j,
		"codeConfigurationValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) CodeConfigurationValuesInput() *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues {
	var returns *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues
	_jsii_.Get(
		j,
		"codeConfigurationValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) ConfigurationSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) ConfigurationSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference_Override(a ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) SetConfigurationSource(val *string) {
	_jsii_.Set(
		j,
		"configurationSource",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) PutCodeConfigurationValues(value *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationCodeConfigurationValues) {
	_jsii_.InvokeVoid(
		a,
		"putCodeConfigurationValues",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference) ResetCodeConfigurationValues() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeConfigurationValues",
		nil, // no parameters
	)
}

type ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference interface {
	cdktf.ComplexObject
	CodeConfiguration() ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference
	CodeConfigurationInput() *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RepositoryUrl() *string
	SetRepositoryUrl(val *string)
	RepositoryUrlInput() *string
	SourceCodeVersion() ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference
	SourceCodeVersionInput() *ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion
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
	PutCodeConfiguration(value *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration)
	PutSourceCodeVersion(value *ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion)
	ResetCodeConfiguration()
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) CodeConfiguration() ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference {
	var returns ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfigurationOutputReference
	_jsii_.Get(
		j,
		"codeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) CodeConfigurationInput() *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration {
	var returns *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration
	_jsii_.Get(
		j,
		"codeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) RepositoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) RepositoryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) SourceCodeVersion() ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference {
	var returns ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference
	_jsii_.Get(
		j,
		"sourceCodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) SourceCodeVersionInput() *ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion {
	var returns *ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion
	_jsii_.Get(
		j,
		"sourceCodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApprunnerServiceSourceConfigurationCodeRepositoryOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationCodeRepositoryOutputReference_Override(a ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) SetRepositoryUrl(val *string) {
	_jsii_.Set(
		j,
		"repositoryUrl",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) PutCodeConfiguration(value *ApprunnerServiceSourceConfigurationCodeRepositoryCodeConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putCodeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) PutSourceCodeVersion(value *ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion) {
	_jsii_.InvokeVoid(
		a,
		"putSourceCodeVersion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference) ResetCodeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeConfiguration",
		nil, // no parameters
	)
}

type ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#type ApprunnerService#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#value ApprunnerService#value}.
	Value *string `json:"value"`
}

type ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference_Override(a ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationCodeRepositorySourceCodeVersionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type ApprunnerServiceSourceConfigurationImageRepository struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#image_identifier ApprunnerService#image_identifier}.
	ImageIdentifier *string `json:"imageIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#image_repository_type ApprunnerService#image_repository_type}.
	ImageRepositoryType *string `json:"imageRepositoryType"`
	// image_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#image_configuration ApprunnerService#image_configuration}
	ImageConfiguration *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration `json:"imageConfiguration"`
}

type ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#port ApprunnerService#port}.
	Port *string `json:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#runtime_environment_variables ApprunnerService#runtime_environment_variables}.
	RuntimeEnvironmentVariables interface{} `json:"runtimeEnvironmentVariables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/apprunner_service.html#start_command ApprunnerService#start_command}.
	StartCommand *string `json:"startCommand"`
}

type ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *string
	SetPort(val *string)
	PortInput() *string
	RuntimeEnvironmentVariables() interface{}
	SetRuntimeEnvironmentVariables(val interface{})
	RuntimeEnvironmentVariablesInput() interface{}
	StartCommand() *string
	SetStartCommand(val *string)
	StartCommandInput() *string
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
	ResetPort()
	ResetRuntimeEnvironmentVariables()
	ResetStartCommand()
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) RuntimeEnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) RuntimeEnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) StartCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) StartCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference_Override(a ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) SetPort(val *string) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) SetRuntimeEnvironmentVariables(val interface{}) {
	_jsii_.Set(
		j,
		"runtimeEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) SetStartCommand(val *string) {
	_jsii_.Set(
		j,
		"startCommand",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) ResetRuntimeEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference) ResetStartCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetStartCommand",
		nil, // no parameters
	)
}

type ApprunnerServiceSourceConfigurationImageRepositoryOutputReference interface {
	cdktf.ComplexObject
	ImageConfiguration() ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference
	ImageConfigurationInput() *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration
	ImageIdentifier() *string
	SetImageIdentifier(val *string)
	ImageIdentifierInput() *string
	ImageRepositoryType() *string
	SetImageRepositoryType(val *string)
	ImageRepositoryTypeInput() *string
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
	PutImageConfiguration(value *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration)
	ResetImageConfiguration()
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationImageRepositoryOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageConfiguration() ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference {
	var returns ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference
	_jsii_.Get(
		j,
		"imageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageConfigurationInput() *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration {
	var returns *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration
	_jsii_.Get(
		j,
		"imageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageRepositoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepositoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageRepositoryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepositoryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApprunnerServiceSourceConfigurationImageRepositoryOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceSourceConfigurationImageRepositoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationImageRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationImageRepositoryOutputReference_Override(a ApprunnerServiceSourceConfigurationImageRepositoryOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationImageRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) SetImageIdentifier(val *string) {
	_jsii_.Set(
		j,
		"imageIdentifier",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) SetImageRepositoryType(val *string) {
	_jsii_.Set(
		j,
		"imageRepositoryType",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) PutImageConfiguration(value *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putImageConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ResetImageConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetImageConfiguration",
		nil, // no parameters
	)
}

type ApprunnerServiceSourceConfigurationOutputReference interface {
	cdktf.ComplexObject
	AuthenticationConfiguration() ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference
	AuthenticationConfigurationInput() *ApprunnerServiceSourceConfigurationAuthenticationConfiguration
	AutoDeploymentsEnabled() interface{}
	SetAutoDeploymentsEnabled(val interface{})
	AutoDeploymentsEnabledInput() interface{}
	CodeRepository() ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference
	CodeRepositoryInput() *ApprunnerServiceSourceConfigurationCodeRepository
	ImageRepository() ApprunnerServiceSourceConfigurationImageRepositoryOutputReference
	ImageRepositoryInput() *ApprunnerServiceSourceConfigurationImageRepository
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
	PutAuthenticationConfiguration(value *ApprunnerServiceSourceConfigurationAuthenticationConfiguration)
	PutCodeRepository(value *ApprunnerServiceSourceConfigurationCodeRepository)
	PutImageRepository(value *ApprunnerServiceSourceConfigurationImageRepository)
	ResetAuthenticationConfiguration()
	ResetAutoDeploymentsEnabled()
	ResetCodeRepository()
	ResetImageRepository()
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) AuthenticationConfiguration() ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference {
	var returns ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) AuthenticationConfigurationInput() *ApprunnerServiceSourceConfigurationAuthenticationConfiguration {
	var returns *ApprunnerServiceSourceConfigurationAuthenticationConfiguration
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) AutoDeploymentsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploymentsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) AutoDeploymentsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploymentsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) CodeRepository() ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference {
	var returns ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) CodeRepositoryInput() *ApprunnerServiceSourceConfigurationCodeRepository {
	var returns *ApprunnerServiceSourceConfigurationCodeRepository
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ImageRepository() ApprunnerServiceSourceConfigurationImageRepositoryOutputReference {
	var returns ApprunnerServiceSourceConfigurationImageRepositoryOutputReference
	_jsii_.Get(
		j,
		"imageRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ImageRepositoryInput() *ApprunnerServiceSourceConfigurationImageRepository {
	var returns *ApprunnerServiceSourceConfigurationImageRepository
	_jsii_.Get(
		j,
		"imageRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewApprunnerServiceSourceConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ApprunnerServiceSourceConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationOutputReference_Override(a ApprunnerServiceSourceConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AppRunner.ApprunnerServiceSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) SetAutoDeploymentsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"autoDeploymentsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) PutAuthenticationConfiguration(value *ApprunnerServiceSourceConfigurationAuthenticationConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) PutCodeRepository(value *ApprunnerServiceSourceConfigurationCodeRepository) {
	_jsii_.InvokeVoid(
		a,
		"putCodeRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) PutImageRepository(value *ApprunnerServiceSourceConfigurationImageRepository) {
	_jsii_.InvokeVoid(
		a,
		"putImageRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ResetAutoDeploymentsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoDeploymentsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ResetImageRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetImageRepository",
		nil, // no parameters
	)
}
