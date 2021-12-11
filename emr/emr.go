package emr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/emr/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/emr_release_labels.html aws_emr_release_labels}.
type DataAwsEmrReleaseLabels interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Filters() DataAwsEmrReleaseLabelsFiltersOutputReference
	FiltersInput() *DataAwsEmrReleaseLabelsFilters
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ReleaseLabels() *[]*string
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
	PutFilters(value *DataAwsEmrReleaseLabelsFilters)
	ResetFilters()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsEmrReleaseLabels
type jsiiProxy_DataAwsEmrReleaseLabels struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) Filters() DataAwsEmrReleaseLabelsFiltersOutputReference {
	var returns DataAwsEmrReleaseLabelsFiltersOutputReference
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) FiltersInput() *DataAwsEmrReleaseLabelsFilters {
	var returns *DataAwsEmrReleaseLabelsFilters
	_jsii_.Get(
		j,
		"filtersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) ReleaseLabels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"releaseLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/emr_release_labels.html aws_emr_release_labels} Data Source.
func NewDataAwsEmrReleaseLabels(scope constructs.Construct, id *string, config *DataAwsEmrReleaseLabelsConfig) DataAwsEmrReleaseLabels {
	_init_.Initialize()

	j := jsiiProxy_DataAwsEmrReleaseLabels{}

	_jsii_.Create(
		"hashicorp_aws.emr.DataAwsEmrReleaseLabels",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/emr_release_labels.html aws_emr_release_labels} Data Source.
func NewDataAwsEmrReleaseLabels_Override(d DataAwsEmrReleaseLabels, scope constructs.Construct, id *string, config *DataAwsEmrReleaseLabelsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.DataAwsEmrReleaseLabels",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabels) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsEmrReleaseLabels_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.emr.DataAwsEmrReleaseLabels",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsEmrReleaseLabels_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.emr.DataAwsEmrReleaseLabels",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsEmrReleaseLabels) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsEmrReleaseLabels) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabels) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabels) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabels) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabels) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabels) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsEmrReleaseLabels) PutFilters(value *DataAwsEmrReleaseLabelsFilters) {
	_jsii_.InvokeVoid(
		d,
		"putFilters",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsEmrReleaseLabels) ResetFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetFilters",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsEmrReleaseLabels) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEmrReleaseLabels) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabels) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabels) ToString() *string {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabels) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Elastic MapReduce.
type DataAwsEmrReleaseLabelsConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/emr_release_labels.html#filters DataAwsEmrReleaseLabels#filters}
	Filters *DataAwsEmrReleaseLabelsFilters `json:"filters"`
}

type DataAwsEmrReleaseLabelsFilters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/emr_release_labels.html#application DataAwsEmrReleaseLabels#application}.
	Application *string `json:"application"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/emr_release_labels.html#prefix DataAwsEmrReleaseLabels#prefix}.
	Prefix *string `json:"prefix"`
}

type DataAwsEmrReleaseLabelsFiltersOutputReference interface {
	cdktf.ComplexObject
	Application() *string
	SetApplication(val *string)
	ApplicationInput() *string
	InternalValue() *DataAwsEmrReleaseLabelsFilters
	SetInternalValue(val *DataAwsEmrReleaseLabelsFilters)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
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
	ResetApplication()
	ResetPrefix()
}

// The jsii proxy struct for DataAwsEmrReleaseLabelsFiltersOutputReference
type jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) Application() *string {
	var returns *string
	_jsii_.Get(
		j,
		"application",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) ApplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) InternalValue() *DataAwsEmrReleaseLabelsFilters {
	var returns *DataAwsEmrReleaseLabelsFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDataAwsEmrReleaseLabelsFiltersOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) DataAwsEmrReleaseLabelsFiltersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.DataAwsEmrReleaseLabelsFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDataAwsEmrReleaseLabelsFiltersOutputReference_Override(d DataAwsEmrReleaseLabelsFiltersOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.DataAwsEmrReleaseLabelsFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) SetApplication(val *string) {
	_jsii_.Set(
		j,
		"application",
		val,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) SetInternalValue(val *DataAwsEmrReleaseLabelsFilters) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) SetPrefix(val *string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) ResetApplication() {
	_jsii_.InvokeVoid(
		d,
		"resetApplication",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsEmrReleaseLabelsFiltersOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetPrefix",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html aws_emr_cluster}.
type EmrCluster interface {
	cdktf.TerraformResource
	AdditionalInfo() *string
	SetAdditionalInfo(val *string)
	AdditionalInfoInput() *string
	Applications() *[]*string
	SetApplications(val *[]*string)
	ApplicationsInput() *[]*string
	Arn() *string
	AutoscalingRole() *string
	SetAutoscalingRole(val *string)
	AutoscalingRoleInput() *string
	AutoTerminationPolicy() EmrClusterAutoTerminationPolicyOutputReference
	AutoTerminationPolicyInput() *EmrClusterAutoTerminationPolicy
	BootstrapAction() *[]*EmrClusterBootstrapAction
	SetBootstrapAction(val *[]*EmrClusterBootstrapAction)
	BootstrapActionInput() *[]*EmrClusterBootstrapAction
	CdktfStack() cdktf.TerraformStack
	ClusterState() *string
	Configurations() *string
	SetConfigurations(val *string)
	ConfigurationsInput() *string
	ConfigurationsJson() *string
	SetConfigurationsJson(val *string)
	ConfigurationsJsonInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	CoreInstanceFleet() EmrClusterCoreInstanceFleetOutputReference
	CoreInstanceFleetInput() *EmrClusterCoreInstanceFleet
	CoreInstanceGroup() EmrClusterCoreInstanceGroupOutputReference
	CoreInstanceGroupInput() *EmrClusterCoreInstanceGroup
	Count() interface{}
	SetCount(val interface{})
	CustomAmiId() *string
	SetCustomAmiId(val *string)
	CustomAmiIdInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EbsRootVolumeSize() *float64
	SetEbsRootVolumeSize(val *float64)
	EbsRootVolumeSizeInput() *float64
	Ec2Attributes() EmrClusterEc2AttributesOutputReference
	Ec2AttributesInput() *EmrClusterEc2Attributes
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KeepJobFlowAliveWhenNoSteps() interface{}
	SetKeepJobFlowAliveWhenNoSteps(val interface{})
	KeepJobFlowAliveWhenNoStepsInput() interface{}
	KerberosAttributes() EmrClusterKerberosAttributesOutputReference
	KerberosAttributesInput() *EmrClusterKerberosAttributes
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogEncryptionKmsKeyId() *string
	SetLogEncryptionKmsKeyId(val *string)
	LogEncryptionKmsKeyIdInput() *string
	LogUri() *string
	SetLogUri(val *string)
	LogUriInput() *string
	MasterInstanceFleet() EmrClusterMasterInstanceFleetOutputReference
	MasterInstanceFleetInput() *EmrClusterMasterInstanceFleet
	MasterInstanceGroup() EmrClusterMasterInstanceGroupOutputReference
	MasterInstanceGroupInput() *EmrClusterMasterInstanceGroup
	MasterPublicDns() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	ReleaseLabelInput() *string
	ScaleDownBehavior() *string
	SetScaleDownBehavior(val *string)
	ScaleDownBehaviorInput() *string
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	SecurityConfigurationInput() *string
	ServiceRole() *string
	SetServiceRole(val *string)
	ServiceRoleInput() *string
	Step() *[]*EmrClusterStep
	SetStep(val *[]*EmrClusterStep)
	StepConcurrencyLevel() *float64
	SetStepConcurrencyLevel(val *float64)
	StepConcurrencyLevelInput() *float64
	StepInput() *[]*EmrClusterStep
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerminationProtection() interface{}
	SetTerminationProtection(val interface{})
	TerminationProtectionInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VisibleToAllUsers() interface{}
	SetVisibleToAllUsers(val interface{})
	VisibleToAllUsersInput() interface{}
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutAutoTerminationPolicy(value *EmrClusterAutoTerminationPolicy)
	PutCoreInstanceFleet(value *EmrClusterCoreInstanceFleet)
	PutCoreInstanceGroup(value *EmrClusterCoreInstanceGroup)
	PutEc2Attributes(value *EmrClusterEc2Attributes)
	PutKerberosAttributes(value *EmrClusterKerberosAttributes)
	PutMasterInstanceFleet(value *EmrClusterMasterInstanceFleet)
	PutMasterInstanceGroup(value *EmrClusterMasterInstanceGroup)
	ResetAdditionalInfo()
	ResetApplications()
	ResetAutoscalingRole()
	ResetAutoTerminationPolicy()
	ResetBootstrapAction()
	ResetConfigurations()
	ResetConfigurationsJson()
	ResetCoreInstanceFleet()
	ResetCoreInstanceGroup()
	ResetCustomAmiId()
	ResetEbsRootVolumeSize()
	ResetEc2Attributes()
	ResetKeepJobFlowAliveWhenNoSteps()
	ResetKerberosAttributes()
	ResetLogEncryptionKmsKeyId()
	ResetLogUri()
	ResetMasterInstanceFleet()
	ResetMasterInstanceGroup()
	ResetOverrideLogicalId()
	ResetScaleDownBehavior()
	ResetSecurityConfiguration()
	ResetStep()
	ResetStepConcurrencyLevel()
	ResetTags()
	ResetTagsAll()
	ResetTerminationProtection()
	ResetVisibleToAllUsers()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for EmrCluster
type jsiiProxy_EmrCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrCluster) AdditionalInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AdditionalInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Applications() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ApplicationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AutoscalingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AutoscalingRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AutoTerminationPolicy() EmrClusterAutoTerminationPolicyOutputReference {
	var returns EmrClusterAutoTerminationPolicyOutputReference
	_jsii_.Get(
		j,
		"autoTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) AutoTerminationPolicyInput() *EmrClusterAutoTerminationPolicy {
	var returns *EmrClusterAutoTerminationPolicy
	_jsii_.Get(
		j,
		"autoTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) BootstrapAction() *[]*EmrClusterBootstrapAction {
	var returns *[]*EmrClusterBootstrapAction
	_jsii_.Get(
		j,
		"bootstrapAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) BootstrapActionInput() *[]*EmrClusterBootstrapAction {
	var returns *[]*EmrClusterBootstrapAction
	_jsii_.Get(
		j,
		"bootstrapActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ClusterState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Configurations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ConfigurationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ConfigurationsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ConfigurationsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CoreInstanceFleet() EmrClusterCoreInstanceFleetOutputReference {
	var returns EmrClusterCoreInstanceFleetOutputReference
	_jsii_.Get(
		j,
		"coreInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CoreInstanceFleetInput() *EmrClusterCoreInstanceFleet {
	var returns *EmrClusterCoreInstanceFleet
	_jsii_.Get(
		j,
		"coreInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CoreInstanceGroup() EmrClusterCoreInstanceGroupOutputReference {
	var returns EmrClusterCoreInstanceGroupOutputReference
	_jsii_.Get(
		j,
		"coreInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CoreInstanceGroupInput() *EmrClusterCoreInstanceGroup {
	var returns *EmrClusterCoreInstanceGroup
	_jsii_.Get(
		j,
		"coreInstanceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) CustomAmiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) EbsRootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) EbsRootVolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Ec2Attributes() EmrClusterEc2AttributesOutputReference {
	var returns EmrClusterEc2AttributesOutputReference
	_jsii_.Get(
		j,
		"ec2Attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Ec2AttributesInput() *EmrClusterEc2Attributes {
	var returns *EmrClusterEc2Attributes
	_jsii_.Get(
		j,
		"ec2AttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) KeepJobFlowAliveWhenNoSteps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) KeepJobFlowAliveWhenNoStepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keepJobFlowAliveWhenNoStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) KerberosAttributes() EmrClusterKerberosAttributesOutputReference {
	var returns EmrClusterKerberosAttributesOutputReference
	_jsii_.Get(
		j,
		"kerberosAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) KerberosAttributesInput() *EmrClusterKerberosAttributes {
	var returns *EmrClusterKerberosAttributes
	_jsii_.Get(
		j,
		"kerberosAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) LogEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) LogEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) LogUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterInstanceFleet() EmrClusterMasterInstanceFleetOutputReference {
	var returns EmrClusterMasterInstanceFleetOutputReference
	_jsii_.Get(
		j,
		"masterInstanceFleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterInstanceFleetInput() *EmrClusterMasterInstanceFleet {
	var returns *EmrClusterMasterInstanceFleet
	_jsii_.Get(
		j,
		"masterInstanceFleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterInstanceGroup() EmrClusterMasterInstanceGroupOutputReference {
	var returns EmrClusterMasterInstanceGroupOutputReference
	_jsii_.Get(
		j,
		"masterInstanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterInstanceGroupInput() *EmrClusterMasterInstanceGroup {
	var returns *EmrClusterMasterInstanceGroup
	_jsii_.Get(
		j,
		"masterInstanceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) MasterPublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPublicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ScaleDownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ScaleDownBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Step() *[]*EmrClusterStep {
	var returns *[]*EmrClusterStep
	_jsii_.Get(
		j,
		"step",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) StepConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) StepConcurrencyLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) StepInput() *[]*EmrClusterStep {
	var returns *[]*EmrClusterStep
	_jsii_.Get(
		j,
		"stepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerminationProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerminationProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) VisibleToAllUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCluster) VisibleToAllUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsersInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html aws_emr_cluster} Resource.
func NewEmrCluster(scope constructs.Construct, id *string, config *EmrClusterConfig) EmrCluster {
	_init_.Initialize()

	j := jsiiProxy_EmrCluster{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html aws_emr_cluster} Resource.
func NewEmrCluster_Override(e EmrCluster, scope constructs.Construct, id *string, config *EmrClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrCluster",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrCluster) SetAdditionalInfo(val *string) {
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetApplications(val *[]*string) {
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetAutoscalingRole(val *string) {
	_jsii_.Set(
		j,
		"autoscalingRole",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetBootstrapAction(val *[]*EmrClusterBootstrapAction) {
	_jsii_.Set(
		j,
		"bootstrapAction",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetConfigurations(val *string) {
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetConfigurationsJson(val *string) {
	_jsii_.Set(
		j,
		"configurationsJson",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetCustomAmiId(val *string) {
	_jsii_.Set(
		j,
		"customAmiId",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetEbsRootVolumeSize(val *float64) {
	_jsii_.Set(
		j,
		"ebsRootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetKeepJobFlowAliveWhenNoSteps(val interface{}) {
	_jsii_.Set(
		j,
		"keepJobFlowAliveWhenNoSteps",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetLogEncryptionKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"logEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetLogUri(val *string) {
	_jsii_.Set(
		j,
		"logUri",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetReleaseLabel(val *string) {
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetScaleDownBehavior(val *string) {
	_jsii_.Set(
		j,
		"scaleDownBehavior",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetStep(val *[]*EmrClusterStep) {
	_jsii_.Set(
		j,
		"step",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetStepConcurrencyLevel(val *float64) {
	_jsii_.Set(
		j,
		"stepConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetTerminationProtection(val interface{}) {
	_jsii_.Set(
		j,
		"terminationProtection",
		val,
	)
}

func (j *jsiiProxy_EmrCluster) SetVisibleToAllUsers(val interface{}) {
	_jsii_.Set(
		j,
		"visibleToAllUsers",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EmrCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.emr.EmrCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.emr.EmrCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_EmrCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_EmrCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EmrCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmrCluster) PutAutoTerminationPolicy(value *EmrClusterAutoTerminationPolicy) {
	_jsii_.InvokeVoid(
		e,
		"putAutoTerminationPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutCoreInstanceFleet(value *EmrClusterCoreInstanceFleet) {
	_jsii_.InvokeVoid(
		e,
		"putCoreInstanceFleet",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutCoreInstanceGroup(value *EmrClusterCoreInstanceGroup) {
	_jsii_.InvokeVoid(
		e,
		"putCoreInstanceGroup",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutEc2Attributes(value *EmrClusterEc2Attributes) {
	_jsii_.InvokeVoid(
		e,
		"putEc2Attributes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutKerberosAttributes(value *EmrClusterKerberosAttributes) {
	_jsii_.InvokeVoid(
		e,
		"putKerberosAttributes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutMasterInstanceFleet(value *EmrClusterMasterInstanceFleet) {
	_jsii_.InvokeVoid(
		e,
		"putMasterInstanceFleet",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) PutMasterInstanceGroup(value *EmrClusterMasterInstanceGroup) {
	_jsii_.InvokeVoid(
		e,
		"putMasterInstanceGroup",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrCluster) ResetAdditionalInfo() {
	_jsii_.InvokeVoid(
		e,
		"resetAdditionalInfo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetApplications() {
	_jsii_.InvokeVoid(
		e,
		"resetApplications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetAutoscalingRole() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscalingRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetAutoTerminationPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoTerminationPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetBootstrapAction() {
	_jsii_.InvokeVoid(
		e,
		"resetBootstrapAction",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetConfigurations() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigurations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetConfigurationsJson() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigurationsJson",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetCoreInstanceFleet() {
	_jsii_.InvokeVoid(
		e,
		"resetCoreInstanceFleet",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetCoreInstanceGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetCoreInstanceGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetCustomAmiId() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomAmiId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetEbsRootVolumeSize() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsRootVolumeSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetEc2Attributes() {
	_jsii_.InvokeVoid(
		e,
		"resetEc2Attributes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetKeepJobFlowAliveWhenNoSteps() {
	_jsii_.InvokeVoid(
		e,
		"resetKeepJobFlowAliveWhenNoSteps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetKerberosAttributes() {
	_jsii_.InvokeVoid(
		e,
		"resetKerberosAttributes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetLogEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		e,
		"resetLogEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetLogUri() {
	_jsii_.InvokeVoid(
		e,
		"resetLogUri",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetMasterInstanceFleet() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterInstanceFleet",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetMasterInstanceGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterInstanceGroup",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_EmrCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetScaleDownBehavior() {
	_jsii_.InvokeVoid(
		e,
		"resetScaleDownBehavior",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetStep() {
	_jsii_.InvokeVoid(
		e,
		"resetStep",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetStepConcurrencyLevel() {
	_jsii_.InvokeVoid(
		e,
		"resetStepConcurrencyLevel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetTerminationProtection() {
	_jsii_.InvokeVoid(
		e,
		"resetTerminationProtection",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) ResetVisibleToAllUsers() {
	_jsii_.InvokeVoid(
		e,
		"resetVisibleToAllUsers",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EmrCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_EmrCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EmrClusterAutoTerminationPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#idle_timeout EmrCluster#idle_timeout}.
	IdleTimeout *float64 `json:"idleTimeout"`
}

type EmrClusterAutoTerminationPolicyOutputReference interface {
	cdktf.ComplexObject
	IdleTimeout() *float64
	SetIdleTimeout(val *float64)
	IdleTimeoutInput() *float64
	InternalValue() *EmrClusterAutoTerminationPolicy
	SetInternalValue(val *EmrClusterAutoTerminationPolicy)
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
	ResetIdleTimeout()
}

// The jsii proxy struct for EmrClusterAutoTerminationPolicyOutputReference
type jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) IdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) IdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) InternalValue() *EmrClusterAutoTerminationPolicy {
	var returns *EmrClusterAutoTerminationPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrClusterAutoTerminationPolicyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrClusterAutoTerminationPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterAutoTerminationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterAutoTerminationPolicyOutputReference_Override(e EmrClusterAutoTerminationPolicyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterAutoTerminationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) SetIdleTimeout(val *float64) {
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) SetInternalValue(val *EmrClusterAutoTerminationPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterAutoTerminationPolicyOutputReference) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

type EmrClusterBootstrapAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#name EmrCluster#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#path EmrCluster#path}.
	Path *string `json:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#args EmrCluster#args}.
	Args *[]*string `json:"args"`
}

// AWS Elastic MapReduce.
type EmrClusterConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#name EmrCluster#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#release_label EmrCluster#release_label}.
	ReleaseLabel *string `json:"releaseLabel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#service_role EmrCluster#service_role}.
	ServiceRole *string `json:"serviceRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#additional_info EmrCluster#additional_info}.
	AdditionalInfo *string `json:"additionalInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#applications EmrCluster#applications}.
	Applications *[]*string `json:"applications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#autoscaling_role EmrCluster#autoscaling_role}.
	AutoscalingRole *string `json:"autoscalingRole"`
	// auto_termination_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#auto_termination_policy EmrCluster#auto_termination_policy}
	AutoTerminationPolicy *EmrClusterAutoTerminationPolicy `json:"autoTerminationPolicy"`
	// bootstrap_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#bootstrap_action EmrCluster#bootstrap_action}
	BootstrapAction *[]*EmrClusterBootstrapAction `json:"bootstrapAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#configurations EmrCluster#configurations}.
	Configurations *string `json:"configurations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#configurations_json EmrCluster#configurations_json}.
	ConfigurationsJson *string `json:"configurationsJson"`
	// core_instance_fleet block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#core_instance_fleet EmrCluster#core_instance_fleet}
	CoreInstanceFleet *EmrClusterCoreInstanceFleet `json:"coreInstanceFleet"`
	// core_instance_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#core_instance_group EmrCluster#core_instance_group}
	CoreInstanceGroup *EmrClusterCoreInstanceGroup `json:"coreInstanceGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#custom_ami_id EmrCluster#custom_ami_id}.
	CustomAmiId *string `json:"customAmiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#ebs_root_volume_size EmrCluster#ebs_root_volume_size}.
	EbsRootVolumeSize *float64 `json:"ebsRootVolumeSize"`
	// ec2_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#ec2_attributes EmrCluster#ec2_attributes}
	Ec2Attributes *EmrClusterEc2Attributes `json:"ec2Attributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#keep_job_flow_alive_when_no_steps EmrCluster#keep_job_flow_alive_when_no_steps}.
	KeepJobFlowAliveWhenNoSteps interface{} `json:"keepJobFlowAliveWhenNoSteps"`
	// kerberos_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#kerberos_attributes EmrCluster#kerberos_attributes}
	KerberosAttributes *EmrClusterKerberosAttributes `json:"kerberosAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#log_encryption_kms_key_id EmrCluster#log_encryption_kms_key_id}.
	LogEncryptionKmsKeyId *string `json:"logEncryptionKmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#log_uri EmrCluster#log_uri}.
	LogUri *string `json:"logUri"`
	// master_instance_fleet block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#master_instance_fleet EmrCluster#master_instance_fleet}
	MasterInstanceFleet *EmrClusterMasterInstanceFleet `json:"masterInstanceFleet"`
	// master_instance_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#master_instance_group EmrCluster#master_instance_group}
	MasterInstanceGroup *EmrClusterMasterInstanceGroup `json:"masterInstanceGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#scale_down_behavior EmrCluster#scale_down_behavior}.
	ScaleDownBehavior *string `json:"scaleDownBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#security_configuration EmrCluster#security_configuration}.
	SecurityConfiguration *string `json:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#step EmrCluster#step}.
	Step *[]*EmrClusterStep `json:"step"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#step_concurrency_level EmrCluster#step_concurrency_level}.
	StepConcurrencyLevel *float64 `json:"stepConcurrencyLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#tags EmrCluster#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#tags_all EmrCluster#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#termination_protection EmrCluster#termination_protection}.
	TerminationProtection interface{} `json:"terminationProtection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#visible_to_all_users EmrCluster#visible_to_all_users}.
	VisibleToAllUsers interface{} `json:"visibleToAllUsers"`
}

type EmrClusterCoreInstanceFleet struct {
	// instance_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#instance_type_configs EmrCluster#instance_type_configs}
	InstanceTypeConfigs *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigs `json:"instanceTypeConfigs"`
	// launch_specifications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#launch_specifications EmrCluster#launch_specifications}
	LaunchSpecifications *EmrClusterCoreInstanceFleetLaunchSpecifications `json:"launchSpecifications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#name EmrCluster#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#target_on_demand_capacity EmrCluster#target_on_demand_capacity}.
	TargetOnDemandCapacity *float64 `json:"targetOnDemandCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#target_spot_capacity EmrCluster#target_spot_capacity}.
	TargetSpotCapacity *float64 `json:"targetSpotCapacity"`
}

type EmrClusterCoreInstanceFleetInstanceTypeConfigs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#instance_type EmrCluster#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#bid_price EmrCluster#bid_price}.
	BidPrice *string `json:"bidPrice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#bid_price_as_percentage_of_on_demand_price EmrCluster#bid_price_as_percentage_of_on_demand_price}.
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice"`
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#configurations EmrCluster#configurations}
	Configurations *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigsConfigurations `json:"configurations"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#ebs_config EmrCluster#ebs_config}
	EbsConfig *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigsEbsConfig `json:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#weighted_capacity EmrCluster#weighted_capacity}.
	WeightedCapacity *float64 `json:"weightedCapacity"`
}

type EmrClusterCoreInstanceFleetInstanceTypeConfigsConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#classification EmrCluster#classification}.
	Classification *string `json:"classification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#properties EmrCluster#properties}.
	Properties interface{} `json:"properties"`
}

type EmrClusterCoreInstanceFleetInstanceTypeConfigsEbsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#size EmrCluster#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#type EmrCluster#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#iops EmrCluster#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#volumes_per_instance EmrCluster#volumes_per_instance}.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

type EmrClusterCoreInstanceFleetLaunchSpecifications struct {
	// on_demand_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#on_demand_specification EmrCluster#on_demand_specification}
	OnDemandSpecification *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification `json:"onDemandSpecification"`
	// spot_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#spot_specification EmrCluster#spot_specification}
	SpotSpecification *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification `json:"spotSpecification"`
}

type EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#allocation_strategy EmrCluster#allocation_strategy}.
	AllocationStrategy *string `json:"allocationStrategy"`
}

type EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *EmrClusterCoreInstanceFleetLaunchSpecifications
	SetInternalValue(val *EmrClusterCoreInstanceFleetLaunchSpecifications)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OnDemandSpecification() *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification
	SetOnDemandSpecification(val *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification)
	OnDemandSpecificationInput() *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification
	SpotSpecification() *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification
	SetSpotSpecification(val *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification)
	SpotSpecificationInput() *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification
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
	ResetOnDemandSpecification()
	ResetSpotSpecification()
}

// The jsii proxy struct for EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference
type jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) InternalValue() *EmrClusterCoreInstanceFleetLaunchSpecifications {
	var returns *EmrClusterCoreInstanceFleetLaunchSpecifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) OnDemandSpecification() *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification {
	var returns *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification
	_jsii_.Get(
		j,
		"onDemandSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) OnDemandSpecificationInput() *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification {
	var returns *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification
	_jsii_.Get(
		j,
		"onDemandSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) SpotSpecification() *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification {
	var returns *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification
	_jsii_.Get(
		j,
		"spotSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) SpotSpecificationInput() *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification {
	var returns *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification
	_jsii_.Get(
		j,
		"spotSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference_Override(e EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) SetInternalValue(val *EmrClusterCoreInstanceFleetLaunchSpecifications) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) SetOnDemandSpecification(val *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification) {
	_jsii_.Set(
		j,
		"onDemandSpecification",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) SetSpotSpecification(val *[]*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification) {
	_jsii_.Set(
		j,
		"spotSpecification",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) ResetOnDemandSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference) ResetSpotSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotSpecification",
		nil, // no parameters
	)
}

type EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#allocation_strategy EmrCluster#allocation_strategy}.
	AllocationStrategy *string `json:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#timeout_action EmrCluster#timeout_action}.
	TimeoutAction *string `json:"timeoutAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#timeout_duration_minutes EmrCluster#timeout_duration_minutes}.
	TimeoutDurationMinutes *float64 `json:"timeoutDurationMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#block_duration_minutes EmrCluster#block_duration_minutes}.
	BlockDurationMinutes *float64 `json:"blockDurationMinutes"`
}

type EmrClusterCoreInstanceFleetOutputReference interface {
	cdktf.ComplexObject
	InstanceTypeConfigs() *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigs
	SetInstanceTypeConfigs(val *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigs)
	InstanceTypeConfigsInput() *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigs
	InternalValue() *EmrClusterCoreInstanceFleet
	SetInternalValue(val *EmrClusterCoreInstanceFleet)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LaunchSpecifications() EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference
	LaunchSpecificationsInput() *EmrClusterCoreInstanceFleetLaunchSpecifications
	Name() *string
	SetName(val *string)
	NameInput() *string
	TargetOnDemandCapacity() *float64
	SetTargetOnDemandCapacity(val *float64)
	TargetOnDemandCapacityInput() *float64
	TargetSpotCapacity() *float64
	SetTargetSpotCapacity(val *float64)
	TargetSpotCapacityInput() *float64
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
	PutLaunchSpecifications(value *EmrClusterCoreInstanceFleetLaunchSpecifications)
	ResetInstanceTypeConfigs()
	ResetLaunchSpecifications()
	ResetName()
	ResetTargetOnDemandCapacity()
	ResetTargetSpotCapacity()
}

// The jsii proxy struct for EmrClusterCoreInstanceFleetOutputReference
type jsiiProxy_EmrClusterCoreInstanceFleetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) InstanceTypeConfigs() *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigs {
	var returns *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigs
	_jsii_.Get(
		j,
		"instanceTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) InstanceTypeConfigsInput() *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigs {
	var returns *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigs
	_jsii_.Get(
		j,
		"instanceTypeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) InternalValue() *EmrClusterCoreInstanceFleet {
	var returns *EmrClusterCoreInstanceFleet
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) LaunchSpecifications() EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference {
	var returns EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference
	_jsii_.Get(
		j,
		"launchSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) LaunchSpecificationsInput() *EmrClusterCoreInstanceFleetLaunchSpecifications {
	var returns *EmrClusterCoreInstanceFleetLaunchSpecifications
	_jsii_.Get(
		j,
		"launchSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) TargetOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) TargetOnDemandCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) TargetSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) TargetSpotCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrClusterCoreInstanceFleetOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrClusterCoreInstanceFleetOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrClusterCoreInstanceFleetOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterCoreInstanceFleetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterCoreInstanceFleetOutputReference_Override(e EmrClusterCoreInstanceFleetOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterCoreInstanceFleetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) SetInstanceTypeConfigs(val *[]*EmrClusterCoreInstanceFleetInstanceTypeConfigs) {
	_jsii_.Set(
		j,
		"instanceTypeConfigs",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) SetInternalValue(val *EmrClusterCoreInstanceFleet) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) SetTargetOnDemandCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) SetTargetSpotCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetSpotCapacity",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) PutLaunchSpecifications(value *EmrClusterCoreInstanceFleetLaunchSpecifications) {
	_jsii_.InvokeVoid(
		e,
		"putLaunchSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) ResetInstanceTypeConfigs() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypeConfigs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) ResetLaunchSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) ResetTargetOnDemandCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetOnDemandCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetOutputReference) ResetTargetSpotCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetSpotCapacity",
		nil, // no parameters
	)
}

type EmrClusterCoreInstanceGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#instance_type EmrCluster#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#autoscaling_policy EmrCluster#autoscaling_policy}.
	AutoscalingPolicy *string `json:"autoscalingPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#bid_price EmrCluster#bid_price}.
	BidPrice *string `json:"bidPrice"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#ebs_config EmrCluster#ebs_config}
	EbsConfig *[]*EmrClusterCoreInstanceGroupEbsConfig `json:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#instance_count EmrCluster#instance_count}.
	InstanceCount *float64 `json:"instanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#name EmrCluster#name}.
	Name *string `json:"name"`
}

type EmrClusterCoreInstanceGroupEbsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#size EmrCluster#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#type EmrCluster#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#iops EmrCluster#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#volumes_per_instance EmrCluster#volumes_per_instance}.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

type EmrClusterCoreInstanceGroupOutputReference interface {
	cdktf.ComplexObject
	AutoscalingPolicy() *string
	SetAutoscalingPolicy(val *string)
	AutoscalingPolicyInput() *string
	BidPrice() *string
	SetBidPrice(val *string)
	BidPriceInput() *string
	EbsConfig() *[]*EmrClusterCoreInstanceGroupEbsConfig
	SetEbsConfig(val *[]*EmrClusterCoreInstanceGroupEbsConfig)
	EbsConfigInput() *[]*EmrClusterCoreInstanceGroupEbsConfig
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *EmrClusterCoreInstanceGroup
	SetInternalValue(val *EmrClusterCoreInstanceGroup)
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
	ResetAutoscalingPolicy()
	ResetBidPrice()
	ResetEbsConfig()
	ResetInstanceCount()
	ResetName()
}

// The jsii proxy struct for EmrClusterCoreInstanceGroupOutputReference
type jsiiProxy_EmrClusterCoreInstanceGroupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) AutoscalingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) AutoscalingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) BidPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) BidPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) EbsConfig() *[]*EmrClusterCoreInstanceGroupEbsConfig {
	var returns *[]*EmrClusterCoreInstanceGroupEbsConfig
	_jsii_.Get(
		j,
		"ebsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) EbsConfigInput() *[]*EmrClusterCoreInstanceGroupEbsConfig {
	var returns *[]*EmrClusterCoreInstanceGroupEbsConfig
	_jsii_.Get(
		j,
		"ebsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) InternalValue() *EmrClusterCoreInstanceGroup {
	var returns *EmrClusterCoreInstanceGroup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrClusterCoreInstanceGroupOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrClusterCoreInstanceGroupOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrClusterCoreInstanceGroupOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterCoreInstanceGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterCoreInstanceGroupOutputReference_Override(e EmrClusterCoreInstanceGroupOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterCoreInstanceGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetAutoscalingPolicy(val *string) {
	_jsii_.Set(
		j,
		"autoscalingPolicy",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetBidPrice(val *string) {
	_jsii_.Set(
		j,
		"bidPrice",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetEbsConfig(val *[]*EmrClusterCoreInstanceGroupEbsConfig) {
	_jsii_.Set(
		j,
		"ebsConfig",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetInternalValue(val *EmrClusterCoreInstanceGroup) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) ResetAutoscalingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscalingPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) ResetBidPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetBidPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) ResetEbsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterCoreInstanceGroupOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

type EmrClusterEc2Attributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#instance_profile EmrCluster#instance_profile}.
	InstanceProfile *string `json:"instanceProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#additional_master_security_groups EmrCluster#additional_master_security_groups}.
	AdditionalMasterSecurityGroups *string `json:"additionalMasterSecurityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#additional_slave_security_groups EmrCluster#additional_slave_security_groups}.
	AdditionalSlaveSecurityGroups *string `json:"additionalSlaveSecurityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#emr_managed_master_security_group EmrCluster#emr_managed_master_security_group}.
	EmrManagedMasterSecurityGroup *string `json:"emrManagedMasterSecurityGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#emr_managed_slave_security_group EmrCluster#emr_managed_slave_security_group}.
	EmrManagedSlaveSecurityGroup *string `json:"emrManagedSlaveSecurityGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#key_name EmrCluster#key_name}.
	KeyName *string `json:"keyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#service_access_security_group EmrCluster#service_access_security_group}.
	ServiceAccessSecurityGroup *string `json:"serviceAccessSecurityGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#subnet_id EmrCluster#subnet_id}.
	SubnetId *string `json:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#subnet_ids EmrCluster#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
}

type EmrClusterEc2AttributesOutputReference interface {
	cdktf.ComplexObject
	AdditionalMasterSecurityGroups() *string
	SetAdditionalMasterSecurityGroups(val *string)
	AdditionalMasterSecurityGroupsInput() *string
	AdditionalSlaveSecurityGroups() *string
	SetAdditionalSlaveSecurityGroups(val *string)
	AdditionalSlaveSecurityGroupsInput() *string
	EmrManagedMasterSecurityGroup() *string
	SetEmrManagedMasterSecurityGroup(val *string)
	EmrManagedMasterSecurityGroupInput() *string
	EmrManagedSlaveSecurityGroup() *string
	SetEmrManagedSlaveSecurityGroup(val *string)
	EmrManagedSlaveSecurityGroupInput() *string
	InstanceProfile() *string
	SetInstanceProfile(val *string)
	InstanceProfileInput() *string
	InternalValue() *EmrClusterEc2Attributes
	SetInternalValue(val *EmrClusterEc2Attributes)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KeyName() *string
	SetKeyName(val *string)
	KeyNameInput() *string
	ServiceAccessSecurityGroup() *string
	SetServiceAccessSecurityGroup(val *string)
	ServiceAccessSecurityGroupInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	ResetAdditionalMasterSecurityGroups()
	ResetAdditionalSlaveSecurityGroups()
	ResetEmrManagedMasterSecurityGroup()
	ResetEmrManagedSlaveSecurityGroup()
	ResetKeyName()
	ResetServiceAccessSecurityGroup()
	ResetSubnetId()
	ResetSubnetIds()
}

// The jsii proxy struct for EmrClusterEc2AttributesOutputReference
type jsiiProxy_EmrClusterEc2AttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) AdditionalMasterSecurityGroups() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalMasterSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) AdditionalMasterSecurityGroupsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalMasterSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) AdditionalSlaveSecurityGroups() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalSlaveSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) AdditionalSlaveSecurityGroupsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalSlaveSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) EmrManagedMasterSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedMasterSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) EmrManagedMasterSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedMasterSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) EmrManagedSlaveSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedSlaveSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) EmrManagedSlaveSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedSlaveSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) InstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) InstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) InternalValue() *EmrClusterEc2Attributes {
	var returns *EmrClusterEc2Attributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) ServiceAccessSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) ServiceAccessSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrClusterEc2AttributesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrClusterEc2AttributesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrClusterEc2AttributesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterEc2AttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterEc2AttributesOutputReference_Override(e EmrClusterEc2AttributesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterEc2AttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetAdditionalMasterSecurityGroups(val *string) {
	_jsii_.Set(
		j,
		"additionalMasterSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetAdditionalSlaveSecurityGroups(val *string) {
	_jsii_.Set(
		j,
		"additionalSlaveSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetEmrManagedMasterSecurityGroup(val *string) {
	_jsii_.Set(
		j,
		"emrManagedMasterSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetEmrManagedSlaveSecurityGroup(val *string) {
	_jsii_.Set(
		j,
		"emrManagedSlaveSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetInstanceProfile(val *string) {
	_jsii_.Set(
		j,
		"instanceProfile",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetInternalValue(val *EmrClusterEc2Attributes) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetKeyName(val *string) {
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetServiceAccessSecurityGroup(val *string) {
	_jsii_.Set(
		j,
		"serviceAccessSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterEc2AttributesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetAdditionalMasterSecurityGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetAdditionalMasterSecurityGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetAdditionalSlaveSecurityGroups() {
	_jsii_.InvokeVoid(
		e,
		"resetAdditionalSlaveSecurityGroups",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetEmrManagedMasterSecurityGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetEmrManagedMasterSecurityGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetEmrManagedSlaveSecurityGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetEmrManagedSlaveSecurityGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetKeyName() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetServiceAccessSecurityGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceAccessSecurityGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterEc2AttributesOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetIds",
		nil, // no parameters
	)
}

type EmrClusterKerberosAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#kdc_admin_password EmrCluster#kdc_admin_password}.
	KdcAdminPassword *string `json:"kdcAdminPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#realm EmrCluster#realm}.
	Realm *string `json:"realm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#ad_domain_join_password EmrCluster#ad_domain_join_password}.
	AdDomainJoinPassword *string `json:"adDomainJoinPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#ad_domain_join_user EmrCluster#ad_domain_join_user}.
	AdDomainJoinUser *string `json:"adDomainJoinUser"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#cross_realm_trust_principal_password EmrCluster#cross_realm_trust_principal_password}.
	CrossRealmTrustPrincipalPassword *string `json:"crossRealmTrustPrincipalPassword"`
}

type EmrClusterKerberosAttributesOutputReference interface {
	cdktf.ComplexObject
	AdDomainJoinPassword() *string
	SetAdDomainJoinPassword(val *string)
	AdDomainJoinPasswordInput() *string
	AdDomainJoinUser() *string
	SetAdDomainJoinUser(val *string)
	AdDomainJoinUserInput() *string
	CrossRealmTrustPrincipalPassword() *string
	SetCrossRealmTrustPrincipalPassword(val *string)
	CrossRealmTrustPrincipalPasswordInput() *string
	InternalValue() *EmrClusterKerberosAttributes
	SetInternalValue(val *EmrClusterKerberosAttributes)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KdcAdminPassword() *string
	SetKdcAdminPassword(val *string)
	KdcAdminPasswordInput() *string
	Realm() *string
	SetRealm(val *string)
	RealmInput() *string
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
	ResetAdDomainJoinPassword()
	ResetAdDomainJoinUser()
	ResetCrossRealmTrustPrincipalPassword()
}

// The jsii proxy struct for EmrClusterKerberosAttributesOutputReference
type jsiiProxy_EmrClusterKerberosAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) AdDomainJoinPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) AdDomainJoinPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) AdDomainJoinUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) AdDomainJoinUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) CrossRealmTrustPrincipalPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustPrincipalPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) CrossRealmTrustPrincipalPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustPrincipalPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) InternalValue() *EmrClusterKerberosAttributes {
	var returns *EmrClusterKerberosAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) KdcAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) KdcAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrClusterKerberosAttributesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrClusterKerberosAttributesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrClusterKerberosAttributesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterKerberosAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterKerberosAttributesOutputReference_Override(e EmrClusterKerberosAttributesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterKerberosAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) SetAdDomainJoinPassword(val *string) {
	_jsii_.Set(
		j,
		"adDomainJoinPassword",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) SetAdDomainJoinUser(val *string) {
	_jsii_.Set(
		j,
		"adDomainJoinUser",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) SetCrossRealmTrustPrincipalPassword(val *string) {
	_jsii_.Set(
		j,
		"crossRealmTrustPrincipalPassword",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) SetInternalValue(val *EmrClusterKerberosAttributes) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) SetKdcAdminPassword(val *string) {
	_jsii_.Set(
		j,
		"kdcAdminPassword",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) SetRealm(val *string) {
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterKerberosAttributesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ResetAdDomainJoinPassword() {
	_jsii_.InvokeVoid(
		e,
		"resetAdDomainJoinPassword",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ResetAdDomainJoinUser() {
	_jsii_.InvokeVoid(
		e,
		"resetAdDomainJoinUser",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterKerberosAttributesOutputReference) ResetCrossRealmTrustPrincipalPassword() {
	_jsii_.InvokeVoid(
		e,
		"resetCrossRealmTrustPrincipalPassword",
		nil, // no parameters
	)
}

type EmrClusterMasterInstanceFleet struct {
	// instance_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#instance_type_configs EmrCluster#instance_type_configs}
	InstanceTypeConfigs *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigs `json:"instanceTypeConfigs"`
	// launch_specifications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#launch_specifications EmrCluster#launch_specifications}
	LaunchSpecifications *EmrClusterMasterInstanceFleetLaunchSpecifications `json:"launchSpecifications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#name EmrCluster#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#target_on_demand_capacity EmrCluster#target_on_demand_capacity}.
	TargetOnDemandCapacity *float64 `json:"targetOnDemandCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#target_spot_capacity EmrCluster#target_spot_capacity}.
	TargetSpotCapacity *float64 `json:"targetSpotCapacity"`
}

type EmrClusterMasterInstanceFleetInstanceTypeConfigs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#instance_type EmrCluster#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#bid_price EmrCluster#bid_price}.
	BidPrice *string `json:"bidPrice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#bid_price_as_percentage_of_on_demand_price EmrCluster#bid_price_as_percentage_of_on_demand_price}.
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice"`
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#configurations EmrCluster#configurations}
	Configurations *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigsConfigurations `json:"configurations"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#ebs_config EmrCluster#ebs_config}
	EbsConfig *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigsEbsConfig `json:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#weighted_capacity EmrCluster#weighted_capacity}.
	WeightedCapacity *float64 `json:"weightedCapacity"`
}

type EmrClusterMasterInstanceFleetInstanceTypeConfigsConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#classification EmrCluster#classification}.
	Classification *string `json:"classification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#properties EmrCluster#properties}.
	Properties interface{} `json:"properties"`
}

type EmrClusterMasterInstanceFleetInstanceTypeConfigsEbsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#size EmrCluster#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#type EmrCluster#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#iops EmrCluster#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#volumes_per_instance EmrCluster#volumes_per_instance}.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

type EmrClusterMasterInstanceFleetLaunchSpecifications struct {
	// on_demand_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#on_demand_specification EmrCluster#on_demand_specification}
	OnDemandSpecification *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification `json:"onDemandSpecification"`
	// spot_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#spot_specification EmrCluster#spot_specification}
	SpotSpecification *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification `json:"spotSpecification"`
}

type EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#allocation_strategy EmrCluster#allocation_strategy}.
	AllocationStrategy *string `json:"allocationStrategy"`
}

type EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *EmrClusterMasterInstanceFleetLaunchSpecifications
	SetInternalValue(val *EmrClusterMasterInstanceFleetLaunchSpecifications)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OnDemandSpecification() *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification
	SetOnDemandSpecification(val *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification)
	OnDemandSpecificationInput() *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification
	SpotSpecification() *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification
	SetSpotSpecification(val *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification)
	SpotSpecificationInput() *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification
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
	ResetOnDemandSpecification()
	ResetSpotSpecification()
}

// The jsii proxy struct for EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference
type jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) InternalValue() *EmrClusterMasterInstanceFleetLaunchSpecifications {
	var returns *EmrClusterMasterInstanceFleetLaunchSpecifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) OnDemandSpecification() *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification {
	var returns *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification
	_jsii_.Get(
		j,
		"onDemandSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) OnDemandSpecificationInput() *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification {
	var returns *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification
	_jsii_.Get(
		j,
		"onDemandSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) SpotSpecification() *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification {
	var returns *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification
	_jsii_.Get(
		j,
		"spotSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) SpotSpecificationInput() *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification {
	var returns *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification
	_jsii_.Get(
		j,
		"spotSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference_Override(e EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) SetInternalValue(val *EmrClusterMasterInstanceFleetLaunchSpecifications) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) SetOnDemandSpecification(val *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification) {
	_jsii_.Set(
		j,
		"onDemandSpecification",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) SetSpotSpecification(val *[]*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification) {
	_jsii_.Set(
		j,
		"spotSpecification",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) ResetOnDemandSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference) ResetSpotSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotSpecification",
		nil, // no parameters
	)
}

type EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#allocation_strategy EmrCluster#allocation_strategy}.
	AllocationStrategy *string `json:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#timeout_action EmrCluster#timeout_action}.
	TimeoutAction *string `json:"timeoutAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#timeout_duration_minutes EmrCluster#timeout_duration_minutes}.
	TimeoutDurationMinutes *float64 `json:"timeoutDurationMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#block_duration_minutes EmrCluster#block_duration_minutes}.
	BlockDurationMinutes *float64 `json:"blockDurationMinutes"`
}

type EmrClusterMasterInstanceFleetOutputReference interface {
	cdktf.ComplexObject
	InstanceTypeConfigs() *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigs
	SetInstanceTypeConfigs(val *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigs)
	InstanceTypeConfigsInput() *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigs
	InternalValue() *EmrClusterMasterInstanceFleet
	SetInternalValue(val *EmrClusterMasterInstanceFleet)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LaunchSpecifications() EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference
	LaunchSpecificationsInput() *EmrClusterMasterInstanceFleetLaunchSpecifications
	Name() *string
	SetName(val *string)
	NameInput() *string
	TargetOnDemandCapacity() *float64
	SetTargetOnDemandCapacity(val *float64)
	TargetOnDemandCapacityInput() *float64
	TargetSpotCapacity() *float64
	SetTargetSpotCapacity(val *float64)
	TargetSpotCapacityInput() *float64
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
	PutLaunchSpecifications(value *EmrClusterMasterInstanceFleetLaunchSpecifications)
	ResetInstanceTypeConfigs()
	ResetLaunchSpecifications()
	ResetName()
	ResetTargetOnDemandCapacity()
	ResetTargetSpotCapacity()
}

// The jsii proxy struct for EmrClusterMasterInstanceFleetOutputReference
type jsiiProxy_EmrClusterMasterInstanceFleetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) InstanceTypeConfigs() *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigs {
	var returns *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigs
	_jsii_.Get(
		j,
		"instanceTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) InstanceTypeConfigsInput() *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigs {
	var returns *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigs
	_jsii_.Get(
		j,
		"instanceTypeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) InternalValue() *EmrClusterMasterInstanceFleet {
	var returns *EmrClusterMasterInstanceFleet
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) LaunchSpecifications() EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference {
	var returns EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference
	_jsii_.Get(
		j,
		"launchSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) LaunchSpecificationsInput() *EmrClusterMasterInstanceFleetLaunchSpecifications {
	var returns *EmrClusterMasterInstanceFleetLaunchSpecifications
	_jsii_.Get(
		j,
		"launchSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) TargetOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) TargetOnDemandCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) TargetSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) TargetSpotCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrClusterMasterInstanceFleetOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrClusterMasterInstanceFleetOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrClusterMasterInstanceFleetOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterMasterInstanceFleetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterMasterInstanceFleetOutputReference_Override(e EmrClusterMasterInstanceFleetOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterMasterInstanceFleetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) SetInstanceTypeConfigs(val *[]*EmrClusterMasterInstanceFleetInstanceTypeConfigs) {
	_jsii_.Set(
		j,
		"instanceTypeConfigs",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) SetInternalValue(val *EmrClusterMasterInstanceFleet) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) SetTargetOnDemandCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) SetTargetSpotCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetSpotCapacity",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) PutLaunchSpecifications(value *EmrClusterMasterInstanceFleetLaunchSpecifications) {
	_jsii_.InvokeVoid(
		e,
		"putLaunchSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) ResetInstanceTypeConfigs() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypeConfigs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) ResetLaunchSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) ResetTargetOnDemandCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetOnDemandCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceFleetOutputReference) ResetTargetSpotCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetSpotCapacity",
		nil, // no parameters
	)
}

type EmrClusterMasterInstanceGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#instance_type EmrCluster#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#bid_price EmrCluster#bid_price}.
	BidPrice *string `json:"bidPrice"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#ebs_config EmrCluster#ebs_config}
	EbsConfig *[]*EmrClusterMasterInstanceGroupEbsConfig `json:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#instance_count EmrCluster#instance_count}.
	InstanceCount *float64 `json:"instanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#name EmrCluster#name}.
	Name *string `json:"name"`
}

type EmrClusterMasterInstanceGroupEbsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#size EmrCluster#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#type EmrCluster#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#iops EmrCluster#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#volumes_per_instance EmrCluster#volumes_per_instance}.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

type EmrClusterMasterInstanceGroupOutputReference interface {
	cdktf.ComplexObject
	BidPrice() *string
	SetBidPrice(val *string)
	BidPriceInput() *string
	EbsConfig() *[]*EmrClusterMasterInstanceGroupEbsConfig
	SetEbsConfig(val *[]*EmrClusterMasterInstanceGroupEbsConfig)
	EbsConfigInput() *[]*EmrClusterMasterInstanceGroupEbsConfig
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *EmrClusterMasterInstanceGroup
	SetInternalValue(val *EmrClusterMasterInstanceGroup)
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
	ResetBidPrice()
	ResetEbsConfig()
	ResetInstanceCount()
	ResetName()
}

// The jsii proxy struct for EmrClusterMasterInstanceGroupOutputReference
type jsiiProxy_EmrClusterMasterInstanceGroupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) BidPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) BidPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) EbsConfig() *[]*EmrClusterMasterInstanceGroupEbsConfig {
	var returns *[]*EmrClusterMasterInstanceGroupEbsConfig
	_jsii_.Get(
		j,
		"ebsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) EbsConfigInput() *[]*EmrClusterMasterInstanceGroupEbsConfig {
	var returns *[]*EmrClusterMasterInstanceGroupEbsConfig
	_jsii_.Get(
		j,
		"ebsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) InternalValue() *EmrClusterMasterInstanceGroup {
	var returns *EmrClusterMasterInstanceGroup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrClusterMasterInstanceGroupOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrClusterMasterInstanceGroupOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrClusterMasterInstanceGroupOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterMasterInstanceGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterMasterInstanceGroupOutputReference_Override(e EmrClusterMasterInstanceGroupOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrClusterMasterInstanceGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) SetBidPrice(val *string) {
	_jsii_.Set(
		j,
		"bidPrice",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) SetEbsConfig(val *[]*EmrClusterMasterInstanceGroupEbsConfig) {
	_jsii_.Set(
		j,
		"ebsConfig",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) SetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) SetInternalValue(val *EmrClusterMasterInstanceGroup) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) ResetBidPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetBidPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) ResetEbsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrClusterMasterInstanceGroupOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

type EmrClusterStep struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#action_on_failure EmrCluster#action_on_failure}.
	ActionOnFailure *string `json:"actionOnFailure"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#hadoop_jar_step EmrCluster#hadoop_jar_step}.
	HadoopJarStep *[]*EmrClusterStepHadoopJarStep `json:"hadoopJarStep"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#name EmrCluster#name}.
	Name *string `json:"name"`
}

type EmrClusterStepHadoopJarStep struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#args EmrCluster#args}.
	Args *[]*string `json:"args"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#jar EmrCluster#jar}.
	Jar *string `json:"jar"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#main_class EmrCluster#main_class}.
	MainClass *string `json:"mainClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#properties EmrCluster#properties}.
	Properties interface{} `json:"properties"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html aws_emr_instance_fleet}.
type EmrInstanceFleet interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstanceTypeConfigs() *[]*EmrInstanceFleetInstanceTypeConfigs
	SetInstanceTypeConfigs(val *[]*EmrInstanceFleetInstanceTypeConfigs)
	InstanceTypeConfigsInput() *[]*EmrInstanceFleetInstanceTypeConfigs
	LaunchSpecifications() EmrInstanceFleetLaunchSpecificationsOutputReference
	LaunchSpecificationsInput() *EmrInstanceFleetLaunchSpecifications
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	ProvisionedOnDemandCapacity() *float64
	ProvisionedSpotCapacity() *float64
	RawOverrides() interface{}
	TargetOnDemandCapacity() *float64
	SetTargetOnDemandCapacity(val *float64)
	TargetOnDemandCapacityInput() *float64
	TargetSpotCapacity() *float64
	SetTargetSpotCapacity(val *float64)
	TargetSpotCapacityInput() *float64
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
	PutLaunchSpecifications(value *EmrInstanceFleetLaunchSpecifications)
	ResetInstanceTypeConfigs()
	ResetLaunchSpecifications()
	ResetName()
	ResetOverrideLogicalId()
	ResetTargetOnDemandCapacity()
	ResetTargetSpotCapacity()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for EmrInstanceFleet
type jsiiProxy_EmrInstanceFleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrInstanceFleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) InstanceTypeConfigs() *[]*EmrInstanceFleetInstanceTypeConfigs {
	var returns *[]*EmrInstanceFleetInstanceTypeConfigs
	_jsii_.Get(
		j,
		"instanceTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) InstanceTypeConfigsInput() *[]*EmrInstanceFleetInstanceTypeConfigs {
	var returns *[]*EmrInstanceFleetInstanceTypeConfigs
	_jsii_.Get(
		j,
		"instanceTypeConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) LaunchSpecifications() EmrInstanceFleetLaunchSpecificationsOutputReference {
	var returns EmrInstanceFleetLaunchSpecificationsOutputReference
	_jsii_.Get(
		j,
		"launchSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) LaunchSpecificationsInput() *EmrInstanceFleetLaunchSpecifications {
	var returns *EmrInstanceFleetLaunchSpecifications
	_jsii_.Get(
		j,
		"launchSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) ProvisionedOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) ProvisionedSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) TargetOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) TargetOnDemandCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) TargetSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) TargetSpotCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html aws_emr_instance_fleet} Resource.
func NewEmrInstanceFleet(scope constructs.Construct, id *string, config *EmrInstanceFleetConfig) EmrInstanceFleet {
	_init_.Initialize()

	j := jsiiProxy_EmrInstanceFleet{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrInstanceFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html aws_emr_instance_fleet} Resource.
func NewEmrInstanceFleet_Override(e EmrInstanceFleet, scope constructs.Construct, id *string, config *EmrInstanceFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrInstanceFleet",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrInstanceFleet) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleet) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleet) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleet) SetInstanceTypeConfigs(val *[]*EmrInstanceFleetInstanceTypeConfigs) {
	_jsii_.Set(
		j,
		"instanceTypeConfigs",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleet) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleet) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleet) SetTargetOnDemandCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleet) SetTargetSpotCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetSpotCapacity",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EmrInstanceFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.emr.EmrInstanceFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrInstanceFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.emr.EmrInstanceFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmrInstanceFleet) PutLaunchSpecifications(value *EmrInstanceFleetLaunchSpecifications) {
	_jsii_.InvokeVoid(
		e,
		"putLaunchSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EmrInstanceFleet) ResetInstanceTypeConfigs() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypeConfigs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceFleet) ResetLaunchSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceFleet) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceFleet) ResetTargetOnDemandCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetOnDemandCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceFleet) ResetTargetSpotCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetSpotCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EmrInstanceFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_EmrInstanceFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Elastic MapReduce.
type EmrInstanceFleetConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#cluster_id EmrInstanceFleet#cluster_id}.
	ClusterId *string `json:"clusterId"`
	// instance_type_configs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#instance_type_configs EmrInstanceFleet#instance_type_configs}
	InstanceTypeConfigs *[]*EmrInstanceFleetInstanceTypeConfigs `json:"instanceTypeConfigs"`
	// launch_specifications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#launch_specifications EmrInstanceFleet#launch_specifications}
	LaunchSpecifications *EmrInstanceFleetLaunchSpecifications `json:"launchSpecifications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#name EmrInstanceFleet#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#target_on_demand_capacity EmrInstanceFleet#target_on_demand_capacity}.
	TargetOnDemandCapacity *float64 `json:"targetOnDemandCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#target_spot_capacity EmrInstanceFleet#target_spot_capacity}.
	TargetSpotCapacity *float64 `json:"targetSpotCapacity"`
}

type EmrInstanceFleetInstanceTypeConfigs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#instance_type EmrInstanceFleet#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#bid_price EmrInstanceFleet#bid_price}.
	BidPrice *string `json:"bidPrice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#bid_price_as_percentage_of_on_demand_price EmrInstanceFleet#bid_price_as_percentage_of_on_demand_price}.
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice"`
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#configurations EmrInstanceFleet#configurations}
	Configurations *[]*EmrInstanceFleetInstanceTypeConfigsConfigurations `json:"configurations"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#ebs_config EmrInstanceFleet#ebs_config}
	EbsConfig *[]*EmrInstanceFleetInstanceTypeConfigsEbsConfig `json:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#weighted_capacity EmrInstanceFleet#weighted_capacity}.
	WeightedCapacity *float64 `json:"weightedCapacity"`
}

type EmrInstanceFleetInstanceTypeConfigsConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#classification EmrInstanceFleet#classification}.
	Classification *string `json:"classification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#properties EmrInstanceFleet#properties}.
	Properties interface{} `json:"properties"`
}

type EmrInstanceFleetInstanceTypeConfigsEbsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#size EmrInstanceFleet#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#type EmrInstanceFleet#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#iops EmrInstanceFleet#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#volumes_per_instance EmrInstanceFleet#volumes_per_instance}.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

type EmrInstanceFleetLaunchSpecifications struct {
	// on_demand_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#on_demand_specification EmrInstanceFleet#on_demand_specification}
	OnDemandSpecification *[]*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification `json:"onDemandSpecification"`
	// spot_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#spot_specification EmrInstanceFleet#spot_specification}
	SpotSpecification *[]*EmrInstanceFleetLaunchSpecificationsSpotSpecification `json:"spotSpecification"`
}

type EmrInstanceFleetLaunchSpecificationsOnDemandSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#allocation_strategy EmrInstanceFleet#allocation_strategy}.
	AllocationStrategy *string `json:"allocationStrategy"`
}

type EmrInstanceFleetLaunchSpecificationsOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *EmrInstanceFleetLaunchSpecifications
	SetInternalValue(val *EmrInstanceFleetLaunchSpecifications)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OnDemandSpecification() *[]*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification
	SetOnDemandSpecification(val *[]*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification)
	OnDemandSpecificationInput() *[]*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification
	SpotSpecification() *[]*EmrInstanceFleetLaunchSpecificationsSpotSpecification
	SetSpotSpecification(val *[]*EmrInstanceFleetLaunchSpecificationsSpotSpecification)
	SpotSpecificationInput() *[]*EmrInstanceFleetLaunchSpecificationsSpotSpecification
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
	ResetOnDemandSpecification()
	ResetSpotSpecification()
}

// The jsii proxy struct for EmrInstanceFleetLaunchSpecificationsOutputReference
type jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) InternalValue() *EmrInstanceFleetLaunchSpecifications {
	var returns *EmrInstanceFleetLaunchSpecifications
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) OnDemandSpecification() *[]*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification {
	var returns *[]*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification
	_jsii_.Get(
		j,
		"onDemandSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) OnDemandSpecificationInput() *[]*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification {
	var returns *[]*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification
	_jsii_.Get(
		j,
		"onDemandSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) SpotSpecification() *[]*EmrInstanceFleetLaunchSpecificationsSpotSpecification {
	var returns *[]*EmrInstanceFleetLaunchSpecificationsSpotSpecification
	_jsii_.Get(
		j,
		"spotSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) SpotSpecificationInput() *[]*EmrInstanceFleetLaunchSpecificationsSpotSpecification {
	var returns *[]*EmrInstanceFleetLaunchSpecificationsSpotSpecification
	_jsii_.Get(
		j,
		"spotSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewEmrInstanceFleetLaunchSpecificationsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) EmrInstanceFleetLaunchSpecificationsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrInstanceFleetLaunchSpecificationsOutputReference_Override(e EmrInstanceFleetLaunchSpecificationsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) SetInternalValue(val *EmrInstanceFleetLaunchSpecifications) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) SetOnDemandSpecification(val *[]*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification) {
	_jsii_.Set(
		j,
		"onDemandSpecification",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) SetSpotSpecification(val *[]*EmrInstanceFleetLaunchSpecificationsSpotSpecification) {
	_jsii_.Set(
		j,
		"spotSpecification",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) ResetOnDemandSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandSpecification",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference) ResetSpotSpecification() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotSpecification",
		nil, // no parameters
	)
}

type EmrInstanceFleetLaunchSpecificationsSpotSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#allocation_strategy EmrInstanceFleet#allocation_strategy}.
	AllocationStrategy *string `json:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#timeout_action EmrInstanceFleet#timeout_action}.
	TimeoutAction *string `json:"timeoutAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#timeout_duration_minutes EmrInstanceFleet#timeout_duration_minutes}.
	TimeoutDurationMinutes *float64 `json:"timeoutDurationMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html#block_duration_minutes EmrInstanceFleet#block_duration_minutes}.
	BlockDurationMinutes *float64 `json:"blockDurationMinutes"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html aws_emr_instance_group}.
type EmrInstanceGroup interface {
	cdktf.TerraformResource
	AutoscalingPolicy() *string
	SetAutoscalingPolicy(val *string)
	AutoscalingPolicyInput() *string
	BidPrice() *string
	SetBidPrice(val *string)
	BidPriceInput() *string
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ConfigurationsJson() *string
	SetConfigurationsJson(val *string)
	ConfigurationsJsonInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EbsConfig() *[]*EmrInstanceGroupEbsConfig
	SetEbsConfig(val *[]*EmrInstanceGroupEbsConfig)
	EbsConfigInput() *[]*EmrInstanceGroupEbsConfig
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RunningInstanceCount() *float64
	Status() *string
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
	ResetAutoscalingPolicy()
	ResetBidPrice()
	ResetConfigurationsJson()
	ResetEbsConfig()
	ResetEbsOptimized()
	ResetInstanceCount()
	ResetName()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for EmrInstanceGroup
type jsiiProxy_EmrInstanceGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrInstanceGroup) AutoscalingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) AutoscalingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) BidPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) BidPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) ConfigurationsJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) ConfigurationsJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationsJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) EbsConfig() *[]*EmrInstanceGroupEbsConfig {
	var returns *[]*EmrInstanceGroupEbsConfig
	_jsii_.Get(
		j,
		"ebsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) EbsConfigInput() *[]*EmrInstanceGroupEbsConfig {
	var returns *[]*EmrInstanceGroupEbsConfig
	_jsii_.Get(
		j,
		"ebsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) RunningInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrInstanceGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html aws_emr_instance_group} Resource.
func NewEmrInstanceGroup(scope constructs.Construct, id *string, config *EmrInstanceGroupConfig) EmrInstanceGroup {
	_init_.Initialize()

	j := jsiiProxy_EmrInstanceGroup{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrInstanceGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html aws_emr_instance_group} Resource.
func NewEmrInstanceGroup_Override(e EmrInstanceGroup, scope constructs.Construct, id *string, config *EmrInstanceGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrInstanceGroup",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetAutoscalingPolicy(val *string) {
	_jsii_.Set(
		j,
		"autoscalingPolicy",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetBidPrice(val *string) {
	_jsii_.Set(
		j,
		"bidPrice",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetConfigurationsJson(val *string) {
	_jsii_.Set(
		j,
		"configurationsJson",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetEbsConfig(val *[]*EmrInstanceGroupEbsConfig) {
	_jsii_.Set(
		j,
		"ebsConfig",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetEbsOptimized(val interface{}) {
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrInstanceGroup) SetProvider(val cdktf.TerraformProvider) {
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
func EmrInstanceGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.emr.EmrInstanceGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrInstanceGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.emr.EmrInstanceGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmrInstanceGroup) ResetAutoscalingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoscalingPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceGroup) ResetBidPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetBidPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceGroup) ResetConfigurationsJson() {
	_jsii_.InvokeVoid(
		e,
		"resetConfigurationsJson",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceGroup) ResetEbsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceGroup) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceGroup) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceGroup) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrInstanceGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EmrInstanceGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_EmrInstanceGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Elastic MapReduce.
type EmrInstanceGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#cluster_id EmrInstanceGroup#cluster_id}.
	ClusterId *string `json:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#instance_type EmrInstanceGroup#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#autoscaling_policy EmrInstanceGroup#autoscaling_policy}.
	AutoscalingPolicy *string `json:"autoscalingPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#bid_price EmrInstanceGroup#bid_price}.
	BidPrice *string `json:"bidPrice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#configurations_json EmrInstanceGroup#configurations_json}.
	ConfigurationsJson *string `json:"configurationsJson"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#ebs_config EmrInstanceGroup#ebs_config}
	EbsConfig *[]*EmrInstanceGroupEbsConfig `json:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#ebs_optimized EmrInstanceGroup#ebs_optimized}.
	EbsOptimized interface{} `json:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#instance_count EmrInstanceGroup#instance_count}.
	InstanceCount *float64 `json:"instanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#name EmrInstanceGroup#name}.
	Name *string `json:"name"`
}

type EmrInstanceGroupEbsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#size EmrInstanceGroup#size}.
	Size *float64 `json:"size"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#type EmrInstanceGroup#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#iops EmrInstanceGroup#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html#volumes_per_instance EmrInstanceGroup#volumes_per_instance}.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html aws_emr_managed_scaling_policy}.
type EmrManagedScalingPolicy interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ComputeLimits() *[]*EmrManagedScalingPolicyComputeLimits
	SetComputeLimits(val *[]*EmrManagedScalingPolicyComputeLimits)
	ComputeLimitsInput() *[]*EmrManagedScalingPolicyComputeLimits
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

// The jsii proxy struct for EmrManagedScalingPolicy
type jsiiProxy_EmrManagedScalingPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrManagedScalingPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) ComputeLimits() *[]*EmrManagedScalingPolicyComputeLimits {
	var returns *[]*EmrManagedScalingPolicyComputeLimits
	_jsii_.Get(
		j,
		"computeLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) ComputeLimitsInput() *[]*EmrManagedScalingPolicyComputeLimits {
	var returns *[]*EmrManagedScalingPolicyComputeLimits
	_jsii_.Get(
		j,
		"computeLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrManagedScalingPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html aws_emr_managed_scaling_policy} Resource.
func NewEmrManagedScalingPolicy(scope constructs.Construct, id *string, config *EmrManagedScalingPolicyConfig) EmrManagedScalingPolicy {
	_init_.Initialize()

	j := jsiiProxy_EmrManagedScalingPolicy{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrManagedScalingPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html aws_emr_managed_scaling_policy} Resource.
func NewEmrManagedScalingPolicy_Override(e EmrManagedScalingPolicy, scope constructs.Construct, id *string, config *EmrManagedScalingPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrManagedScalingPolicy",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicy) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicy) SetComputeLimits(val *[]*EmrManagedScalingPolicyComputeLimits) {
	_jsii_.Set(
		j,
		"computeLimits",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrManagedScalingPolicy) SetProvider(val cdktf.TerraformProvider) {
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
func EmrManagedScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.emr.EmrManagedScalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrManagedScalingPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.emr.EmrManagedScalingPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrManagedScalingPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EmrManagedScalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_EmrManagedScalingPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type EmrManagedScalingPolicyComputeLimits struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html#maximum_capacity_units EmrManagedScalingPolicy#maximum_capacity_units}.
	MaximumCapacityUnits *float64 `json:"maximumCapacityUnits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html#minimum_capacity_units EmrManagedScalingPolicy#minimum_capacity_units}.
	MinimumCapacityUnits *float64 `json:"minimumCapacityUnits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html#unit_type EmrManagedScalingPolicy#unit_type}.
	UnitType *string `json:"unitType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html#maximum_core_capacity_units EmrManagedScalingPolicy#maximum_core_capacity_units}.
	MaximumCoreCapacityUnits *float64 `json:"maximumCoreCapacityUnits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html#maximum_ondemand_capacity_units EmrManagedScalingPolicy#maximum_ondemand_capacity_units}.
	MaximumOndemandCapacityUnits *float64 `json:"maximumOndemandCapacityUnits"`
}

// AWS Elastic MapReduce.
type EmrManagedScalingPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html#cluster_id EmrManagedScalingPolicy#cluster_id}.
	ClusterId *string `json:"clusterId"`
	// compute_limits block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html#compute_limits EmrManagedScalingPolicy#compute_limits}
	ComputeLimits *[]*EmrManagedScalingPolicyComputeLimits `json:"computeLimits"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/emr_security_configuration.html aws_emr_security_configuration}.
type EmrSecurityConfiguration interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	Configuration() *string
	SetConfiguration(val *string)
	ConfigurationInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CreationDate() *string
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
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
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
	ResetName()
	ResetNamePrefix()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for EmrSecurityConfiguration
type jsiiProxy_EmrSecurityConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrSecurityConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) ConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSecurityConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_security_configuration.html aws_emr_security_configuration} Resource.
func NewEmrSecurityConfiguration(scope constructs.Construct, id *string, config *EmrSecurityConfigurationConfig) EmrSecurityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_EmrSecurityConfiguration{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrSecurityConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_security_configuration.html aws_emr_security_configuration} Resource.
func NewEmrSecurityConfiguration_Override(e EmrSecurityConfiguration, scope constructs.Construct, id *string, config *EmrSecurityConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrSecurityConfiguration",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrSecurityConfiguration) SetConfiguration(val *string) {
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_EmrSecurityConfiguration) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrSecurityConfiguration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrSecurityConfiguration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrSecurityConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrSecurityConfiguration) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_EmrSecurityConfiguration) SetProvider(val cdktf.TerraformProvider) {
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
func EmrSecurityConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.emr.EmrSecurityConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrSecurityConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.emr.EmrSecurityConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmrSecurityConfiguration) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrSecurityConfiguration) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetNamePrefix",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrSecurityConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EmrSecurityConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_EmrSecurityConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Elastic MapReduce.
type EmrSecurityConfigurationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_security_configuration.html#configuration EmrSecurityConfiguration#configuration}.
	Configuration *string `json:"configuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_security_configuration.html#name EmrSecurityConfiguration#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_security_configuration.html#name_prefix EmrSecurityConfiguration#name_prefix}.
	NamePrefix *string `json:"namePrefix"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html aws_emr_studio}.
type EmrStudio interface {
	cdktf.TerraformResource
	Arn() *string
	AuthMode() *string
	SetAuthMode(val *string)
	AuthModeInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultS3Location() *string
	SetDefaultS3Location(val *string)
	DefaultS3LocationInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EngineSecurityGroupId() *string
	SetEngineSecurityGroupId(val *string)
	EngineSecurityGroupIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdpAuthUrl() *string
	SetIdpAuthUrl(val *string)
	IdpAuthUrlInput() *string
	IdpRelayStateParameterName() *string
	SetIdpRelayStateParameterName(val *string)
	IdpRelayStateParameterNameInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServiceRole() *string
	SetServiceRole(val *string)
	ServiceRoleInput() *string
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
	Url() *string
	UserRole() *string
	SetUserRole(val *string)
	UserRoleInput() *string
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
	WorkspaceSecurityGroupId() *string
	SetWorkspaceSecurityGroupId(val *string)
	WorkspaceSecurityGroupIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDescription()
	ResetIdpAuthUrl()
	ResetIdpRelayStateParameterName()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetUserRole()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for EmrStudio
type jsiiProxy_EmrStudio struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrStudio) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) AuthMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) AuthModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) DefaultS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) DefaultS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) EngineSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) EngineSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) IdpAuthUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpAuthUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) IdpAuthUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpAuthUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) IdpRelayStateParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpRelayStateParameterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) IdpRelayStateParameterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpRelayStateParameterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) UserRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) UserRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) WorkspaceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudio) WorkspaceSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html aws_emr_studio} Resource.
func NewEmrStudio(scope constructs.Construct, id *string, config *EmrStudioConfig) EmrStudio {
	_init_.Initialize()

	j := jsiiProxy_EmrStudio{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrStudio",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html aws_emr_studio} Resource.
func NewEmrStudio_Override(e EmrStudio, scope constructs.Construct, id *string, config *EmrStudioConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrStudio",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrStudio) SetAuthMode(val *string) {
	_jsii_.Set(
		j,
		"authMode",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetDefaultS3Location(val *string) {
	_jsii_.Set(
		j,
		"defaultS3Location",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetEngineSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"engineSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetIdpAuthUrl(val *string) {
	_jsii_.Set(
		j,
		"idpAuthUrl",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetIdpRelayStateParameterName(val *string) {
	_jsii_.Set(
		j,
		"idpRelayStateParameterName",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetUserRole(val *string) {
	_jsii_.Set(
		j,
		"userRole",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_EmrStudio) SetWorkspaceSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"workspaceSecurityGroupId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EmrStudio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.emr.EmrStudio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrStudio_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.emr.EmrStudio",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudio) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_EmrStudio) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudio) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudio) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudio) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudio) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EmrStudio) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmrStudio) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetIdpAuthUrl() {
	_jsii_.InvokeVoid(
		e,
		"resetIdpAuthUrl",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetIdpRelayStateParameterName() {
	_jsii_.InvokeVoid(
		e,
		"resetIdpRelayStateParameterName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_EmrStudio) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) ResetUserRole() {
	_jsii_.InvokeVoid(
		e,
		"resetUserRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudio) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudio) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EmrStudio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_EmrStudio) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Elastic MapReduce.
type EmrStudioConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#auth_mode EmrStudio#auth_mode}.
	AuthMode *string `json:"authMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#default_s3_location EmrStudio#default_s3_location}.
	DefaultS3Location *string `json:"defaultS3Location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#engine_security_group_id EmrStudio#engine_security_group_id}.
	EngineSecurityGroupId *string `json:"engineSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#name EmrStudio#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#service_role EmrStudio#service_role}.
	ServiceRole *string `json:"serviceRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#subnet_ids EmrStudio#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#vpc_id EmrStudio#vpc_id}.
	VpcId *string `json:"vpcId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#workspace_security_group_id EmrStudio#workspace_security_group_id}.
	WorkspaceSecurityGroupId *string `json:"workspaceSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#description EmrStudio#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#idp_auth_url EmrStudio#idp_auth_url}.
	IdpAuthUrl *string `json:"idpAuthUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#idp_relay_state_parameter_name EmrStudio#idp_relay_state_parameter_name}.
	IdpRelayStateParameterName *string `json:"idpRelayStateParameterName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#tags EmrStudio#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#tags_all EmrStudio#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio.html#user_role EmrStudio#user_role}.
	UserRole *string `json:"userRole"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/emr_studio_session_mapping.html aws_emr_studio_session_mapping}.
type EmrStudioSessionMapping interface {
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
	IdentityId() *string
	SetIdentityId(val *string)
	IdentityIdInput() *string
	IdentityName() *string
	SetIdentityName(val *string)
	IdentityNameInput() *string
	IdentityType() *string
	SetIdentityType(val *string)
	IdentityTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SessionPolicyArn() *string
	SetSessionPolicyArn(val *string)
	SessionPolicyArnInput() *string
	StudioId() *string
	SetStudioId(val *string)
	StudioIdInput() *string
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
	ResetIdentityId()
	ResetIdentityName()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for EmrStudioSessionMapping
type jsiiProxy_EmrStudioSessionMapping struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EmrStudioSessionMapping) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) IdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) IdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) IdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) IdentityNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) SessionPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) SessionPolicyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) StudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) StudioIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrStudioSessionMapping) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_studio_session_mapping.html aws_emr_studio_session_mapping} Resource.
func NewEmrStudioSessionMapping(scope constructs.Construct, id *string, config *EmrStudioSessionMappingConfig) EmrStudioSessionMapping {
	_init_.Initialize()

	j := jsiiProxy_EmrStudioSessionMapping{}

	_jsii_.Create(
		"hashicorp_aws.emr.EmrStudioSessionMapping",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_studio_session_mapping.html aws_emr_studio_session_mapping} Resource.
func NewEmrStudioSessionMapping_Override(e EmrStudioSessionMapping, scope constructs.Construct, id *string, config *EmrStudioSessionMappingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.emr.EmrStudioSessionMapping",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EmrStudioSessionMapping) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EmrStudioSessionMapping) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EmrStudioSessionMapping) SetIdentityId(val *string) {
	_jsii_.Set(
		j,
		"identityId",
		val,
	)
}

func (j *jsiiProxy_EmrStudioSessionMapping) SetIdentityName(val *string) {
	_jsii_.Set(
		j,
		"identityName",
		val,
	)
}

func (j *jsiiProxy_EmrStudioSessionMapping) SetIdentityType(val *string) {
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_EmrStudioSessionMapping) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EmrStudioSessionMapping) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EmrStudioSessionMapping) SetSessionPolicyArn(val *string) {
	_jsii_.Set(
		j,
		"sessionPolicyArn",
		val,
	)
}

func (j *jsiiProxy_EmrStudioSessionMapping) SetStudioId(val *string) {
	_jsii_.Set(
		j,
		"studioId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EmrStudioSessionMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.emr.EmrStudioSessionMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EmrStudioSessionMapping_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.emr.EmrStudioSessionMapping",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EmrStudioSessionMapping) ResetIdentityId() {
	_jsii_.InvokeVoid(
		e,
		"resetIdentityId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudioSessionMapping) ResetIdentityName() {
	_jsii_.InvokeVoid(
		e,
		"resetIdentityName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EmrStudioSessionMapping) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EmrStudioSessionMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_EmrStudioSessionMapping) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Elastic MapReduce.
type EmrStudioSessionMappingConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio_session_mapping.html#identity_type EmrStudioSessionMapping#identity_type}.
	IdentityType *string `json:"identityType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio_session_mapping.html#session_policy_arn EmrStudioSessionMapping#session_policy_arn}.
	SessionPolicyArn *string `json:"sessionPolicyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio_session_mapping.html#studio_id EmrStudioSessionMapping#studio_id}.
	StudioId *string `json:"studioId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio_session_mapping.html#identity_id EmrStudioSessionMapping#identity_id}.
	IdentityId *string `json:"identityId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_studio_session_mapping.html#identity_name EmrStudioSessionMapping#identity_name}.
	IdentityName *string `json:"identityName"`
}
