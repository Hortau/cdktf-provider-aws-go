package licensemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/licensemanager/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_association.html aws_licensemanager_association}.
type LicensemanagerAssociation interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LicenseConfigurationArn() *string
	SetLicenseConfigurationArn(val *string)
	LicenseConfigurationArnInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceArn() *string
	SetResourceArn(val *string)
	ResourceArnInput() *string
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
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LicensemanagerAssociation
type jsiiProxy_LicensemanagerAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LicensemanagerAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) LicenseConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) LicenseConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseConfigurationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) ResourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_association.html aws_licensemanager_association} Resource.
func NewLicensemanagerAssociation(scope constructs.Construct, id *string, config *LicensemanagerAssociationConfig) LicensemanagerAssociation {
	_init_.Initialize()

	j := jsiiProxy_LicensemanagerAssociation{}

	_jsii_.Create(
		"hashicorp_aws.LicenseManager.LicensemanagerAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_association.html aws_licensemanager_association} Resource.
func NewLicensemanagerAssociation_Override(l LicensemanagerAssociation, scope constructs.Construct, id *string, config *LicensemanagerAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.LicenseManager.LicensemanagerAssociation",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LicensemanagerAssociation) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerAssociation) SetLicenseConfigurationArn(val *string) {
	_jsii_.Set(
		j,
		"licenseConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerAssociation) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LicensemanagerAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.LicenseManager.LicensemanagerAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LicensemanagerAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.LicenseManager.LicensemanagerAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LicensemanagerAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LicensemanagerAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LicensemanagerAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (l *jsiiProxy_LicensemanagerAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (l *jsiiProxy_LicensemanagerAssociation) GetStringAttribute(terraformAttribute *string) *string {
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
func (l *jsiiProxy_LicensemanagerAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LicensemanagerAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LicensemanagerAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerAssociation) SynthesizeAttributes() *map[string]interface{} {
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
func (l *jsiiProxy_LicensemanagerAssociation) ToMetadata() interface{} {
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
func (l *jsiiProxy_LicensemanagerAssociation) ToString() *string {
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
func (l *jsiiProxy_LicensemanagerAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LicensemanagerAssociationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_association.html#license_configuration_arn LicensemanagerAssociation#license_configuration_arn}.
	LicenseConfigurationArn *string `json:"licenseConfigurationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_association.html#resource_arn LicensemanagerAssociation#resource_arn}.
	ResourceArn *string `json:"resourceArn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html aws_licensemanager_license_configuration}.
type LicensemanagerLicenseConfiguration interface {
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
	LicenseCount() *float64
	SetLicenseCount(val *float64)
	LicenseCountHardLimit() interface{}
	SetLicenseCountHardLimit(val interface{})
	LicenseCountHardLimitInput() interface{}
	LicenseCountingType() *string
	SetLicenseCountingType(val *string)
	LicenseCountingTypeInput() *string
	LicenseCountInput() *float64
	LicenseRules() *[]*string
	SetLicenseRules(val *[]*string)
	LicenseRulesInput() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	OwnerAccountId() *string
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
	ResetDescription()
	ResetLicenseCount()
	ResetLicenseCountHardLimit()
	ResetLicenseRules()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for LicensemanagerLicenseConfiguration
type jsiiProxy_LicensemanagerLicenseConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) LicenseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"licenseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) LicenseCountHardLimit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseCountHardLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) LicenseCountHardLimitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseCountHardLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) LicenseCountingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseCountingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) LicenseCountingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseCountingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) LicenseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"licenseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) LicenseRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenseRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) LicenseRulesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenseRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) OwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html aws_licensemanager_license_configuration} Resource.
func NewLicensemanagerLicenseConfiguration(scope constructs.Construct, id *string, config *LicensemanagerLicenseConfigurationConfig) LicensemanagerLicenseConfiguration {
	_init_.Initialize()

	j := jsiiProxy_LicensemanagerLicenseConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.LicenseManager.LicensemanagerLicenseConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html aws_licensemanager_license_configuration} Resource.
func NewLicensemanagerLicenseConfiguration_Override(l LicensemanagerLicenseConfiguration, scope constructs.Construct, id *string, config *LicensemanagerLicenseConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.LicenseManager.LicensemanagerLicenseConfiguration",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetLicenseCount(val *float64) {
	_jsii_.Set(
		j,
		"licenseCount",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetLicenseCountHardLimit(val interface{}) {
	_jsii_.Set(
		j,
		"licenseCountHardLimit",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetLicenseCountingType(val *string) {
	_jsii_.Set(
		j,
		"licenseCountingType",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetLicenseRules(val *[]*string) {
	_jsii_.Set(
		j,
		"licenseRules",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfiguration) SetTagsAll(val interface{}) {
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
func LicensemanagerLicenseConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.LicenseManager.LicensemanagerLicenseConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LicensemanagerLicenseConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.LicenseManager.LicensemanagerLicenseConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) GetStringAttribute(terraformAttribute *string) *string {
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
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ResetLicenseCount() {
	_jsii_.InvokeVoid(
		l,
		"resetLicenseCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ResetLicenseCountHardLimit() {
	_jsii_.InvokeVoid(
		l,
		"resetLicenseCountHardLimit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ResetLicenseRules() {
	_jsii_.InvokeVoid(
		l,
		"resetLicenseRules",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ResetTagsAll() {
	_jsii_.InvokeVoid(
		l,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfiguration) SynthesizeAttributes() *map[string]interface{} {
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
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ToMetadata() interface{} {
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
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ToString() *string {
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
func (l *jsiiProxy_LicensemanagerLicenseConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type LicensemanagerLicenseConfigurationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html#license_counting_type LicensemanagerLicenseConfiguration#license_counting_type}.
	LicenseCountingType *string `json:"licenseCountingType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html#name LicensemanagerLicenseConfiguration#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html#description LicensemanagerLicenseConfiguration#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html#license_count LicensemanagerLicenseConfiguration#license_count}.
	LicenseCount *float64 `json:"licenseCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html#license_count_hard_limit LicensemanagerLicenseConfiguration#license_count_hard_limit}.
	LicenseCountHardLimit interface{} `json:"licenseCountHardLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html#license_rules LicensemanagerLicenseConfiguration#license_rules}.
	LicenseRules *[]*string `json:"licenseRules"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html#tags LicensemanagerLicenseConfiguration#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/licensemanager_license_configuration.html#tags_all LicensemanagerLicenseConfiguration#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}
