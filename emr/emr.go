package emr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/emr/internal"
)

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
	PutCoreInstanceFleet(value *EmrClusterCoreInstanceFleet)
	PutCoreInstanceGroup(value *EmrClusterCoreInstanceGroup)
	PutEc2Attributes(value *EmrClusterEc2Attributes)
	PutKerberosAttributes(value *EmrClusterKerberosAttributes)
	PutMasterInstanceFleet(value *EmrClusterMasterInstanceFleet)
	PutMasterInstanceGroup(value *EmrClusterMasterInstanceGroup)
	ResetAdditionalInfo()
	ResetApplications()
	ResetAutoscalingRole()
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
		"hashicorp_aws.EMR.EmrCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html aws_emr_cluster} Resource.
func NewEmrCluster_Override(e EmrCluster, scope constructs.Construct, id *string, config *EmrClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrCluster",
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
		"hashicorp_aws.EMR.EmrCluster",
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
		"hashicorp_aws.EMR.EmrCluster",
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

type EmrClusterBootstrapAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#name EmrCluster#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#path EmrCluster#path}.
	Path *string `json:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster.html#args EmrCluster#args}.
	Args *[]*string `json:"args"`
}

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
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference_Override(e EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
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
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterCoreInstanceFleetOutputReference_Override(e EmrClusterCoreInstanceFleetOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetOutputReference",
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
		"hashicorp_aws.EMR.EmrClusterCoreInstanceGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterCoreInstanceGroupOutputReference_Override(e EmrClusterCoreInstanceGroupOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceGroupOutputReference",
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
		"hashicorp_aws.EMR.EmrClusterEc2AttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterEc2AttributesOutputReference_Override(e EmrClusterEc2AttributesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrClusterEc2AttributesOutputReference",
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
		"hashicorp_aws.EMR.EmrClusterKerberosAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterKerberosAttributesOutputReference_Override(e EmrClusterKerberosAttributesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrClusterKerberosAttributesOutputReference",
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
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference_Override(e EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
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
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterMasterInstanceFleetOutputReference_Override(e EmrClusterMasterInstanceFleetOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetOutputReference",
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
		"hashicorp_aws.EMR.EmrClusterMasterInstanceGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrClusterMasterInstanceGroupOutputReference_Override(e EmrClusterMasterInstanceGroupOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceGroupOutputReference",
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
		"hashicorp_aws.EMR.EmrInstanceFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_fleet.html aws_emr_instance_fleet} Resource.
func NewEmrInstanceFleet_Override(e EmrInstanceFleet, scope constructs.Construct, id *string, config *EmrInstanceFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrInstanceFleet",
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
		"hashicorp_aws.EMR.EmrInstanceFleet",
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
		"hashicorp_aws.EMR.EmrInstanceFleet",
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
		"hashicorp_aws.EMR.EmrInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewEmrInstanceFleetLaunchSpecificationsOutputReference_Override(e EmrInstanceFleetLaunchSpecificationsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrInstanceFleetLaunchSpecificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
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
		"hashicorp_aws.EMR.EmrInstanceGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_instance_group.html aws_emr_instance_group} Resource.
func NewEmrInstanceGroup_Override(e EmrInstanceGroup, scope constructs.Construct, id *string, config *EmrInstanceGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrInstanceGroup",
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
		"hashicorp_aws.EMR.EmrInstanceGroup",
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
		"hashicorp_aws.EMR.EmrInstanceGroup",
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
		"hashicorp_aws.EMR.EmrManagedScalingPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_managed_scaling_policy.html aws_emr_managed_scaling_policy} Resource.
func NewEmrManagedScalingPolicy_Override(e EmrManagedScalingPolicy, scope constructs.Construct, id *string, config *EmrManagedScalingPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrManagedScalingPolicy",
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
		"hashicorp_aws.EMR.EmrManagedScalingPolicy",
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
		"hashicorp_aws.EMR.EmrManagedScalingPolicy",
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
		"hashicorp_aws.EMR.EmrSecurityConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/emr_security_configuration.html aws_emr_security_configuration} Resource.
func NewEmrSecurityConfiguration_Override(e EmrSecurityConfiguration, scope constructs.Construct, id *string, config *EmrSecurityConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.EMR.EmrSecurityConfiguration",
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
		"hashicorp_aws.EMR.EmrSecurityConfiguration",
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
		"hashicorp_aws.EMR.EmrSecurityConfiguration",
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
