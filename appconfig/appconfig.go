package appconfig

import (
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hortau/cdktf-provider-aws-go/appconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appconfig_application aws_appconfig_application}.
type AppconfigApplication interface {
	cdktf.TerraformResource
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
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppconfigApplication
type jsiiProxy_AppconfigApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppconfigApplication) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_application aws_appconfig_application} Resource.
func NewAppconfigApplication(scope constructs.Construct, id *string, config *AppconfigApplicationConfig) AppconfigApplication {
	_init_.Initialize()

	j := jsiiProxy_AppconfigApplication{}

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_application aws_appconfig_application} Resource.
func NewAppconfigApplication_Override(a AppconfigApplication, scope constructs.Construct, id *string, config *AppconfigApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigApplication",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppconfigApplication) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppconfigApplication) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppconfigApplication) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppconfigApplication) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppconfigApplication) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppconfigApplication) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppconfigApplication) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppconfigApplication) SetTagsAll(val *map[string]*string) {
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
func AppconfigApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appconfig.AppconfigApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppconfigApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appconfig.AppconfigApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppconfigApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigApplication) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppconfigApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppconfigApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigApplication) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppconfigApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppconfigApplication) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppconfigApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigApplication) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigApplication) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigApplication) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppconfigApplication) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppconfigApplication) ToString() *string {
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
func (a *jsiiProxy_AppconfigApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppConfig.
type AppconfigApplicationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_application#name AppconfigApplication#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_application#description AppconfigApplication#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_application#tags AppconfigApplication#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_application#tags_all AppconfigApplication#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile aws_appconfig_configuration_profile}.
type AppconfigConfigurationProfile interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConfigurationProfileId() *string
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
	LocationUri() *string
	SetLocationUri(val *string)
	LocationUriInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RetrievalRoleArn() *string
	SetRetrievalRoleArn(val *string)
	RetrievalRoleArnInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Validator() interface{}
	SetValidator(val interface{})
	ValidatorInput() interface{}
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
	ResetOverrideLogicalId()
	ResetRetrievalRoleArn()
	ResetTags()
	ResetTagsAll()
	ResetValidator()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppconfigConfigurationProfile
type jsiiProxy_AppconfigConfigurationProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppconfigConfigurationProfile) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) ConfigurationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) LocationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) RetrievalRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) RetrievalRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) Validator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigConfigurationProfile) ValidatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validatorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile aws_appconfig_configuration_profile} Resource.
func NewAppconfigConfigurationProfile(scope constructs.Construct, id *string, config *AppconfigConfigurationProfileConfig) AppconfigConfigurationProfile {
	_init_.Initialize()

	j := jsiiProxy_AppconfigConfigurationProfile{}

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigConfigurationProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile aws_appconfig_configuration_profile} Resource.
func NewAppconfigConfigurationProfile_Override(a AppconfigConfigurationProfile, scope constructs.Construct, id *string, config *AppconfigConfigurationProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigConfigurationProfile",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetLocationUri(val *string) {
	_jsii_.Set(
		j,
		"locationUri",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetRetrievalRoleArn(val *string) {
	_jsii_.Set(
		j,
		"retrievalRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AppconfigConfigurationProfile) SetValidator(val interface{}) {
	_jsii_.Set(
		j,
		"validator",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AppconfigConfigurationProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appconfig.AppconfigConfigurationProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppconfigConfigurationProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appconfig.AppconfigConfigurationProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigConfigurationProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppconfigConfigurationProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigConfigurationProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigConfigurationProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigConfigurationProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppconfigConfigurationProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppconfigConfigurationProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigConfigurationProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigConfigurationProfile) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppconfigConfigurationProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigConfigurationProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigConfigurationProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppconfigConfigurationProfile) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppconfigConfigurationProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigConfigurationProfile) ResetRetrievalRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRetrievalRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigConfigurationProfile) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigConfigurationProfile) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigConfigurationProfile) ResetValidator() {
	_jsii_.InvokeVoid(
		a,
		"resetValidator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigConfigurationProfile) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppconfigConfigurationProfile) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppconfigConfigurationProfile) ToString() *string {
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
func (a *jsiiProxy_AppconfigConfigurationProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppConfig.
type AppconfigConfigurationProfileConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#application_id AppconfigConfigurationProfile#application_id}.
	ApplicationId *string `json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#location_uri AppconfigConfigurationProfile#location_uri}.
	LocationUri *string `json:"locationUri" yaml:"locationUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#name AppconfigConfigurationProfile#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#description AppconfigConfigurationProfile#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#retrieval_role_arn AppconfigConfigurationProfile#retrieval_role_arn}.
	RetrievalRoleArn *string `json:"retrievalRoleArn" yaml:"retrievalRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#tags AppconfigConfigurationProfile#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#tags_all AppconfigConfigurationProfile#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// validator block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#validator AppconfigConfigurationProfile#validator}
	Validator interface{} `json:"validator" yaml:"validator"`
}

type AppconfigConfigurationProfileValidator struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#type AppconfigConfigurationProfile#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_configuration_profile#content AppconfigConfigurationProfile#content}.
	Content *string `json:"content" yaml:"content"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment aws_appconfig_deployment}.
type AppconfigDeployment interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConfigurationProfileId() *string
	SetConfigurationProfileId(val *string)
	ConfigurationProfileIdInput() *string
	ConfigurationVersion() *string
	SetConfigurationVersion(val *string)
	ConfigurationVersionInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentNumber() *float64
	DeploymentStrategyId() *string
	SetDeploymentStrategyId(val *string)
	DeploymentStrategyIdInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EnvironmentId() *string
	SetEnvironmentId(val *string)
	EnvironmentIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	State() *string
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
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppconfigDeployment
type jsiiProxy_AppconfigDeployment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppconfigDeployment) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) ConfigurationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) ConfigurationProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) ConfigurationVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) ConfigurationVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) DeploymentNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) DeploymentStrategyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentStrategyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) DeploymentStrategyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentStrategyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) EnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeployment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment aws_appconfig_deployment} Resource.
func NewAppconfigDeployment(scope constructs.Construct, id *string, config *AppconfigDeploymentConfig) AppconfigDeployment {
	_init_.Initialize()

	j := jsiiProxy_AppconfigDeployment{}

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigDeployment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment aws_appconfig_deployment} Resource.
func NewAppconfigDeployment_Override(a AppconfigDeployment, scope constructs.Construct, id *string, config *AppconfigDeploymentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigDeployment",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetConfigurationProfileId(val *string) {
	_jsii_.Set(
		j,
		"configurationProfileId",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetConfigurationVersion(val *string) {
	_jsii_.Set(
		j,
		"configurationVersion",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetDeploymentStrategyId(val *string) {
	_jsii_.Set(
		j,
		"deploymentStrategyId",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetEnvironmentId(val *string) {
	_jsii_.Set(
		j,
		"environmentId",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeployment) SetTagsAll(val *map[string]*string) {
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
func AppconfigDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appconfig.AppconfigDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppconfigDeployment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appconfig.AppconfigDeployment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeployment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppconfigDeployment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeployment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigDeployment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeployment) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppconfigDeployment) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppconfigDeployment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeployment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeployment) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppconfigDeployment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeployment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigDeployment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppconfigDeployment) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppconfigDeployment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigDeployment) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigDeployment) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigDeployment) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppconfigDeployment) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppconfigDeployment) ToString() *string {
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
func (a *jsiiProxy_AppconfigDeployment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppConfig.
type AppconfigDeploymentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment#application_id AppconfigDeployment#application_id}.
	ApplicationId *string `json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment#configuration_profile_id AppconfigDeployment#configuration_profile_id}.
	ConfigurationProfileId *string `json:"configurationProfileId" yaml:"configurationProfileId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment#configuration_version AppconfigDeployment#configuration_version}.
	ConfigurationVersion *string `json:"configurationVersion" yaml:"configurationVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment#deployment_strategy_id AppconfigDeployment#deployment_strategy_id}.
	DeploymentStrategyId *string `json:"deploymentStrategyId" yaml:"deploymentStrategyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment#environment_id AppconfigDeployment#environment_id}.
	EnvironmentId *string `json:"environmentId" yaml:"environmentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment#description AppconfigDeployment#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment#tags AppconfigDeployment#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment#tags_all AppconfigDeployment#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy aws_appconfig_deployment_strategy}.
type AppconfigDeploymentStrategy interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentDurationInMinutes() *float64
	SetDeploymentDurationInMinutes(val *float64)
	DeploymentDurationInMinutesInput() *float64
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	FinalBakeTimeInMinutes() *float64
	SetFinalBakeTimeInMinutes(val *float64)
	FinalBakeTimeInMinutesInput() *float64
	Fqn() *string
	FriendlyUniqueId() *string
	GrowthFactor() *float64
	SetGrowthFactor(val *float64)
	GrowthFactorInput() *float64
	GrowthType() *string
	SetGrowthType(val *string)
	GrowthTypeInput() *string
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
	ReplicateTo() *string
	SetReplicateTo(val *string)
	ReplicateToInput() *string
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
	ResetFinalBakeTimeInMinutes()
	ResetGrowthType()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppconfigDeploymentStrategy
type jsiiProxy_AppconfigDeploymentStrategy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) DeploymentDurationInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentDurationInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) DeploymentDurationInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentDurationInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) FinalBakeTimeInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"finalBakeTimeInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) FinalBakeTimeInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"finalBakeTimeInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) GrowthFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"growthFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) GrowthFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"growthFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) GrowthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"growthType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) GrowthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"growthTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) ReplicateTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicateTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) ReplicateToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicateToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy aws_appconfig_deployment_strategy} Resource.
func NewAppconfigDeploymentStrategy(scope constructs.Construct, id *string, config *AppconfigDeploymentStrategyConfig) AppconfigDeploymentStrategy {
	_init_.Initialize()

	j := jsiiProxy_AppconfigDeploymentStrategy{}

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigDeploymentStrategy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy aws_appconfig_deployment_strategy} Resource.
func NewAppconfigDeploymentStrategy_Override(a AppconfigDeploymentStrategy, scope constructs.Construct, id *string, config *AppconfigDeploymentStrategyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigDeploymentStrategy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetDeploymentDurationInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"deploymentDurationInMinutes",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetFinalBakeTimeInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"finalBakeTimeInMinutes",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetGrowthFactor(val *float64) {
	_jsii_.Set(
		j,
		"growthFactor",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetGrowthType(val *string) {
	_jsii_.Set(
		j,
		"growthType",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetReplicateTo(val *string) {
	_jsii_.Set(
		j,
		"replicateTo",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppconfigDeploymentStrategy) SetTagsAll(val *map[string]*string) {
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
func AppconfigDeploymentStrategy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appconfig.AppconfigDeploymentStrategy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppconfigDeploymentStrategy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appconfig.AppconfigDeploymentStrategy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeploymentStrategy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppconfigDeploymentStrategy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeploymentStrategy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigDeploymentStrategy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeploymentStrategy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppconfigDeploymentStrategy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppconfigDeploymentStrategy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeploymentStrategy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeploymentStrategy) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppconfigDeploymentStrategy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigDeploymentStrategy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigDeploymentStrategy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppconfigDeploymentStrategy) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigDeploymentStrategy) ResetFinalBakeTimeInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetFinalBakeTimeInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigDeploymentStrategy) ResetGrowthType() {
	_jsii_.InvokeVoid(
		a,
		"resetGrowthType",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppconfigDeploymentStrategy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigDeploymentStrategy) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigDeploymentStrategy) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigDeploymentStrategy) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppconfigDeploymentStrategy) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppconfigDeploymentStrategy) ToString() *string {
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
func (a *jsiiProxy_AppconfigDeploymentStrategy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppConfig.
type AppconfigDeploymentStrategyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy#deployment_duration_in_minutes AppconfigDeploymentStrategy#deployment_duration_in_minutes}.
	DeploymentDurationInMinutes *float64 `json:"deploymentDurationInMinutes" yaml:"deploymentDurationInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy#growth_factor AppconfigDeploymentStrategy#growth_factor}.
	GrowthFactor *float64 `json:"growthFactor" yaml:"growthFactor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy#name AppconfigDeploymentStrategy#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy#replicate_to AppconfigDeploymentStrategy#replicate_to}.
	ReplicateTo *string `json:"replicateTo" yaml:"replicateTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy#description AppconfigDeploymentStrategy#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy#final_bake_time_in_minutes AppconfigDeploymentStrategy#final_bake_time_in_minutes}.
	FinalBakeTimeInMinutes *float64 `json:"finalBakeTimeInMinutes" yaml:"finalBakeTimeInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy#growth_type AppconfigDeploymentStrategy#growth_type}.
	GrowthType *string `json:"growthType" yaml:"growthType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy#tags AppconfigDeploymentStrategy#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_deployment_strategy#tags_all AppconfigDeploymentStrategy#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment aws_appconfig_environment}.
type AppconfigEnvironment interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
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
	EnvironmentId() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Monitor() interface{}
	SetMonitor(val interface{})
	MonitorInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	State() *string
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
	ResetMonitor()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppconfigEnvironment
type jsiiProxy_AppconfigEnvironment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppconfigEnvironment) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Monitor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) MonitorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment aws_appconfig_environment} Resource.
func NewAppconfigEnvironment(scope constructs.Construct, id *string, config *AppconfigEnvironmentConfig) AppconfigEnvironment {
	_init_.Initialize()

	j := jsiiProxy_AppconfigEnvironment{}

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment aws_appconfig_environment} Resource.
func NewAppconfigEnvironment_Override(a AppconfigEnvironment, scope constructs.Construct, id *string, config *AppconfigEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigEnvironment",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetMonitor(val interface{}) {
	_jsii_.Set(
		j,
		"monitor",
		val,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppconfigEnvironment) SetTagsAll(val *map[string]*string) {
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
func AppconfigEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appconfig.AppconfigEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppconfigEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appconfig.AppconfigEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigEnvironment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppconfigEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppconfigEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppconfigEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppconfigEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigEnvironment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppconfigEnvironment) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigEnvironment) ResetMonitor() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitor",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppconfigEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigEnvironment) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigEnvironment) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppconfigEnvironment) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppconfigEnvironment) ToString() *string {
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
func (a *jsiiProxy_AppconfigEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppConfig.
type AppconfigEnvironmentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment#application_id AppconfigEnvironment#application_id}.
	ApplicationId *string `json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment#name AppconfigEnvironment#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment#description AppconfigEnvironment#description}.
	Description *string `json:"description" yaml:"description"`
	// monitor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment#monitor AppconfigEnvironment#monitor}
	Monitor interface{} `json:"monitor" yaml:"monitor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment#tags AppconfigEnvironment#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment#tags_all AppconfigEnvironment#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

type AppconfigEnvironmentMonitor struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment#alarm_arn AppconfigEnvironment#alarm_arn}.
	AlarmArn *string `json:"alarmArn" yaml:"alarmArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_environment#alarm_role_arn AppconfigEnvironment#alarm_role_arn}.
	AlarmRoleArn *string `json:"alarmRoleArn" yaml:"alarmRoleArn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appconfig_hosted_configuration_version aws_appconfig_hosted_configuration_version}.
type AppconfigHostedConfigurationVersion interface {
	cdktf.TerraformResource
	ApplicationId() *string
	SetApplicationId(val *string)
	ApplicationIdInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConfigurationProfileId() *string
	SetConfigurationProfileId(val *string)
	ConfigurationProfileIdInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
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
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VersionNumber() *float64
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
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppconfigHostedConfigurationVersion
type jsiiProxy_AppconfigHostedConfigurationVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) ApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) ConfigurationProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) ConfigurationProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) VersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"versionNumber",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_hosted_configuration_version aws_appconfig_hosted_configuration_version} Resource.
func NewAppconfigHostedConfigurationVersion(scope constructs.Construct, id *string, config *AppconfigHostedConfigurationVersionConfig) AppconfigHostedConfigurationVersion {
	_init_.Initialize()

	j := jsiiProxy_AppconfigHostedConfigurationVersion{}

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigHostedConfigurationVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appconfig_hosted_configuration_version aws_appconfig_hosted_configuration_version} Resource.
func NewAppconfigHostedConfigurationVersion_Override(a AppconfigHostedConfigurationVersion, scope constructs.Construct, id *string, config *AppconfigHostedConfigurationVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appconfig.AppconfigHostedConfigurationVersion",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) SetApplicationId(val *string) {
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) SetConfigurationProfileId(val *string) {
	_jsii_.Set(
		j,
		"configurationProfileId",
		val,
	)
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) SetContent(val *string) {
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) SetContentType(val *string) {
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppconfigHostedConfigurationVersion) SetProvider(val cdktf.TerraformProvider) {
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
func AppconfigHostedConfigurationVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appconfig.AppconfigHostedConfigurationVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppconfigHostedConfigurationVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appconfig.AppconfigHostedConfigurationVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppconfigHostedConfigurationVersion) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppconfigHostedConfigurationVersion) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) ToString() *string {
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
func (a *jsiiProxy_AppconfigHostedConfigurationVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppConfig.
type AppconfigHostedConfigurationVersionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_hosted_configuration_version#application_id AppconfigHostedConfigurationVersion#application_id}.
	ApplicationId *string `json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_hosted_configuration_version#configuration_profile_id AppconfigHostedConfigurationVersion#configuration_profile_id}.
	ConfigurationProfileId *string `json:"configurationProfileId" yaml:"configurationProfileId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_hosted_configuration_version#content AppconfigHostedConfigurationVersion#content}.
	Content *string `json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_hosted_configuration_version#content_type AppconfigHostedConfigurationVersion#content_type}.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appconfig_hosted_configuration_version#description AppconfigHostedConfigurationVersion#description}.
	Description *string `json:"description" yaml:"description"`
}

