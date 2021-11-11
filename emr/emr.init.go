package emr

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrCluster",
		reflect.TypeOf((*EmrCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalInfo", GoGetter: "AdditionalInfo"},
			_jsii_.MemberProperty{JsiiProperty: "additionalInfoInput", GoGetter: "AdditionalInfoInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applications", GoGetter: "Applications"},
			_jsii_.MemberProperty{JsiiProperty: "applicationsInput", GoGetter: "ApplicationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingRole", GoGetter: "AutoscalingRole"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingRoleInput", GoGetter: "AutoscalingRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapAction", GoGetter: "BootstrapAction"},
			_jsii_.MemberProperty{JsiiProperty: "bootstrapActionInput", GoGetter: "BootstrapActionInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterState", GoGetter: "ClusterState"},
			_jsii_.MemberProperty{JsiiProperty: "configurations", GoGetter: "Configurations"},
			_jsii_.MemberProperty{JsiiProperty: "configurationsInput", GoGetter: "ConfigurationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "configurationsJson", GoGetter: "ConfigurationsJson"},
			_jsii_.MemberProperty{JsiiProperty: "configurationsJsonInput", GoGetter: "ConfigurationsJsonInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "coreInstanceFleet", GoGetter: "CoreInstanceFleet"},
			_jsii_.MemberProperty{JsiiProperty: "coreInstanceFleetInput", GoGetter: "CoreInstanceFleetInput"},
			_jsii_.MemberProperty{JsiiProperty: "coreInstanceGroup", GoGetter: "CoreInstanceGroup"},
			_jsii_.MemberProperty{JsiiProperty: "coreInstanceGroupInput", GoGetter: "CoreInstanceGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customAmiId", GoGetter: "CustomAmiId"},
			_jsii_.MemberProperty{JsiiProperty: "customAmiIdInput", GoGetter: "CustomAmiIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "ebsRootVolumeSize", GoGetter: "EbsRootVolumeSize"},
			_jsii_.MemberProperty{JsiiProperty: "ebsRootVolumeSizeInput", GoGetter: "EbsRootVolumeSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "ec2Attributes", GoGetter: "Ec2Attributes"},
			_jsii_.MemberProperty{JsiiProperty: "ec2AttributesInput", GoGetter: "Ec2AttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keepJobFlowAliveWhenNoSteps", GoGetter: "KeepJobFlowAliveWhenNoSteps"},
			_jsii_.MemberProperty{JsiiProperty: "keepJobFlowAliveWhenNoStepsInput", GoGetter: "KeepJobFlowAliveWhenNoStepsInput"},
			_jsii_.MemberProperty{JsiiProperty: "kerberosAttributes", GoGetter: "KerberosAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "kerberosAttributesInput", GoGetter: "KerberosAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logEncryptionKmsKeyId", GoGetter: "LogEncryptionKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "logEncryptionKmsKeyIdInput", GoGetter: "LogEncryptionKmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "logUri", GoGetter: "LogUri"},
			_jsii_.MemberProperty{JsiiProperty: "logUriInput", GoGetter: "LogUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "masterInstanceFleet", GoGetter: "MasterInstanceFleet"},
			_jsii_.MemberProperty{JsiiProperty: "masterInstanceFleetInput", GoGetter: "MasterInstanceFleetInput"},
			_jsii_.MemberProperty{JsiiProperty: "masterInstanceGroup", GoGetter: "MasterInstanceGroup"},
			_jsii_.MemberProperty{JsiiProperty: "masterInstanceGroupInput", GoGetter: "MasterInstanceGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "masterPublicDns", GoGetter: "MasterPublicDns"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putCoreInstanceFleet", GoMethod: "PutCoreInstanceFleet"},
			_jsii_.MemberMethod{JsiiMethod: "putCoreInstanceGroup", GoMethod: "PutCoreInstanceGroup"},
			_jsii_.MemberMethod{JsiiMethod: "putEc2Attributes", GoMethod: "PutEc2Attributes"},
			_jsii_.MemberMethod{JsiiMethod: "putKerberosAttributes", GoMethod: "PutKerberosAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "putMasterInstanceFleet", GoMethod: "PutMasterInstanceFleet"},
			_jsii_.MemberMethod{JsiiMethod: "putMasterInstanceGroup", GoMethod: "PutMasterInstanceGroup"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "releaseLabel", GoGetter: "ReleaseLabel"},
			_jsii_.MemberProperty{JsiiProperty: "releaseLabelInput", GoGetter: "ReleaseLabelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalInfo", GoMethod: "ResetAdditionalInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplications", GoMethod: "ResetApplications"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoscalingRole", GoMethod: "ResetAutoscalingRole"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootstrapAction", GoMethod: "ResetBootstrapAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigurations", GoMethod: "ResetConfigurations"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigurationsJson", GoMethod: "ResetConfigurationsJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetCoreInstanceFleet", GoMethod: "ResetCoreInstanceFleet"},
			_jsii_.MemberMethod{JsiiMethod: "resetCoreInstanceGroup", GoMethod: "ResetCoreInstanceGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomAmiId", GoMethod: "ResetCustomAmiId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsRootVolumeSize", GoMethod: "ResetEbsRootVolumeSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetEc2Attributes", GoMethod: "ResetEc2Attributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeepJobFlowAliveWhenNoSteps", GoMethod: "ResetKeepJobFlowAliveWhenNoSteps"},
			_jsii_.MemberMethod{JsiiMethod: "resetKerberosAttributes", GoMethod: "ResetKerberosAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogEncryptionKmsKeyId", GoMethod: "ResetLogEncryptionKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogUri", GoMethod: "ResetLogUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetMasterInstanceFleet", GoMethod: "ResetMasterInstanceFleet"},
			_jsii_.MemberMethod{JsiiMethod: "resetMasterInstanceGroup", GoMethod: "ResetMasterInstanceGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetScaleDownBehavior", GoMethod: "ResetScaleDownBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityConfiguration", GoMethod: "ResetSecurityConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetStep", GoMethod: "ResetStep"},
			_jsii_.MemberMethod{JsiiMethod: "resetStepConcurrencyLevel", GoMethod: "ResetStepConcurrencyLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetTerminationProtection", GoMethod: "ResetTerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisibleToAllUsers", GoMethod: "ResetVisibleToAllUsers"},
			_jsii_.MemberProperty{JsiiProperty: "scaleDownBehavior", GoGetter: "ScaleDownBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "scaleDownBehaviorInput", GoGetter: "ScaleDownBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfiguration", GoGetter: "SecurityConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfigurationInput", GoGetter: "SecurityConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleInput", GoGetter: "ServiceRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "step", GoGetter: "Step"},
			_jsii_.MemberProperty{JsiiProperty: "stepConcurrencyLevel", GoGetter: "StepConcurrencyLevel"},
			_jsii_.MemberProperty{JsiiProperty: "stepConcurrencyLevelInput", GoGetter: "StepConcurrencyLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "stepInput", GoGetter: "StepInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtectionInput", GoGetter: "TerminationProtectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "visibleToAllUsers", GoGetter: "VisibleToAllUsers"},
			_jsii_.MemberProperty{JsiiProperty: "visibleToAllUsersInput", GoGetter: "VisibleToAllUsersInput"},
		},
		func() interface{} {
			j := jsiiProxy_EmrCluster{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterBootstrapAction",
		reflect.TypeOf((*EmrClusterBootstrapAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterConfig",
		reflect.TypeOf((*EmrClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleet",
		reflect.TypeOf((*EmrClusterCoreInstanceFleet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetInstanceTypeConfigs",
		reflect.TypeOf((*EmrClusterCoreInstanceFleetInstanceTypeConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetInstanceTypeConfigsConfigurations",
		reflect.TypeOf((*EmrClusterCoreInstanceFleetInstanceTypeConfigsConfigurations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetInstanceTypeConfigsEbsConfig",
		reflect.TypeOf((*EmrClusterCoreInstanceFleetInstanceTypeConfigsEbsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetLaunchSpecifications",
		reflect.TypeOf((*EmrClusterCoreInstanceFleetLaunchSpecifications)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification",
		reflect.TypeOf((*EmrClusterCoreInstanceFleetLaunchSpecificationsOnDemandSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference",
		reflect.TypeOf((*EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "onDemandSpecification", GoGetter: "OnDemandSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "onDemandSpecificationInput", GoGetter: "OnDemandSpecificationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnDemandSpecification", GoMethod: "ResetOnDemandSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpotSpecification", GoMethod: "ResetSpotSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "spotSpecification", GoGetter: "SpotSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "spotSpecificationInput", GoGetter: "SpotSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_EmrClusterCoreInstanceFleetLaunchSpecificationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification",
		reflect.TypeOf((*EmrClusterCoreInstanceFleetLaunchSpecificationsSpotSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceFleetOutputReference",
		reflect.TypeOf((*EmrClusterCoreInstanceFleetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeConfigs", GoGetter: "InstanceTypeConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeConfigsInput", GoGetter: "InstanceTypeConfigsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "launchSpecifications", GoGetter: "LaunchSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "launchSpecificationsInput", GoGetter: "LaunchSpecificationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLaunchSpecifications", GoMethod: "PutLaunchSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceTypeConfigs", GoMethod: "ResetInstanceTypeConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetLaunchSpecifications", GoMethod: "ResetLaunchSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetOnDemandCapacity", GoMethod: "ResetTargetOnDemandCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetSpotCapacity", GoMethod: "ResetTargetSpotCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "targetOnDemandCapacity", GoGetter: "TargetOnDemandCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "targetOnDemandCapacityInput", GoGetter: "TargetOnDemandCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetSpotCapacity", GoGetter: "TargetSpotCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "targetSpotCapacityInput", GoGetter: "TargetSpotCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_EmrClusterCoreInstanceFleetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceGroup",
		reflect.TypeOf((*EmrClusterCoreInstanceGroup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceGroupEbsConfig",
		reflect.TypeOf((*EmrClusterCoreInstanceGroupEbsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrClusterCoreInstanceGroupOutputReference",
		reflect.TypeOf((*EmrClusterCoreInstanceGroupOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoscalingPolicy", GoGetter: "AutoscalingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingPolicyInput", GoGetter: "AutoscalingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "bidPrice", GoGetter: "BidPrice"},
			_jsii_.MemberProperty{JsiiProperty: "bidPriceInput", GoGetter: "BidPriceInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebsConfig", GoGetter: "EbsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "ebsConfigInput", GoGetter: "EbsConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCount", GoGetter: "InstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCountInput", GoGetter: "InstanceCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoscalingPolicy", GoMethod: "ResetAutoscalingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetBidPrice", GoMethod: "ResetBidPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsConfig", GoMethod: "ResetEbsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceCount", GoMethod: "ResetInstanceCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_EmrClusterCoreInstanceGroupOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterEc2Attributes",
		reflect.TypeOf((*EmrClusterEc2Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrClusterEc2AttributesOutputReference",
		reflect.TypeOf((*EmrClusterEc2AttributesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalMasterSecurityGroups", GoGetter: "AdditionalMasterSecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "additionalMasterSecurityGroupsInput", GoGetter: "AdditionalMasterSecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "additionalSlaveSecurityGroups", GoGetter: "AdditionalSlaveSecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "additionalSlaveSecurityGroupsInput", GoGetter: "AdditionalSlaveSecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "emrManagedMasterSecurityGroup", GoGetter: "EmrManagedMasterSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "emrManagedMasterSecurityGroupInput", GoGetter: "EmrManagedMasterSecurityGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "emrManagedSlaveSecurityGroup", GoGetter: "EmrManagedSlaveSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "emrManagedSlaveSecurityGroupInput", GoGetter: "EmrManagedSlaveSecurityGroupInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfile", GoGetter: "InstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "instanceProfileInput", GoGetter: "InstanceProfileInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "keyNameInput", GoGetter: "KeyNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalMasterSecurityGroups", GoMethod: "ResetAdditionalMasterSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalSlaveSecurityGroups", GoMethod: "ResetAdditionalSlaveSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmrManagedMasterSecurityGroup", GoMethod: "ResetEmrManagedMasterSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmrManagedSlaveSecurityGroup", GoMethod: "ResetEmrManagedSlaveSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyName", GoMethod: "ResetKeyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccessSecurityGroup", GoMethod: "ResetServiceAccessSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetId", GoMethod: "ResetSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetIds", GoMethod: "ResetSubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccessSecurityGroup", GoGetter: "ServiceAccessSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccessSecurityGroupInput", GoGetter: "ServiceAccessSecurityGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdInput", GoGetter: "SubnetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_EmrClusterEc2AttributesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterKerberosAttributes",
		reflect.TypeOf((*EmrClusterKerberosAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrClusterKerberosAttributesOutputReference",
		reflect.TypeOf((*EmrClusterKerberosAttributesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "adDomainJoinPassword", GoGetter: "AdDomainJoinPassword"},
			_jsii_.MemberProperty{JsiiProperty: "adDomainJoinPasswordInput", GoGetter: "AdDomainJoinPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "adDomainJoinUser", GoGetter: "AdDomainJoinUser"},
			_jsii_.MemberProperty{JsiiProperty: "adDomainJoinUserInput", GoGetter: "AdDomainJoinUserInput"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustPrincipalPassword", GoGetter: "CrossRealmTrustPrincipalPassword"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustPrincipalPasswordInput", GoGetter: "CrossRealmTrustPrincipalPasswordInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kdcAdminPassword", GoGetter: "KdcAdminPassword"},
			_jsii_.MemberProperty{JsiiProperty: "kdcAdminPasswordInput", GoGetter: "KdcAdminPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "realm", GoGetter: "Realm"},
			_jsii_.MemberProperty{JsiiProperty: "realmInput", GoGetter: "RealmInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdDomainJoinPassword", GoMethod: "ResetAdDomainJoinPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdDomainJoinUser", GoMethod: "ResetAdDomainJoinUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossRealmTrustPrincipalPassword", GoMethod: "ResetCrossRealmTrustPrincipalPassword"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_EmrClusterKerberosAttributesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleet",
		reflect.TypeOf((*EmrClusterMasterInstanceFleet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetInstanceTypeConfigs",
		reflect.TypeOf((*EmrClusterMasterInstanceFleetInstanceTypeConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetInstanceTypeConfigsConfigurations",
		reflect.TypeOf((*EmrClusterMasterInstanceFleetInstanceTypeConfigsConfigurations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetInstanceTypeConfigsEbsConfig",
		reflect.TypeOf((*EmrClusterMasterInstanceFleetInstanceTypeConfigsEbsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetLaunchSpecifications",
		reflect.TypeOf((*EmrClusterMasterInstanceFleetLaunchSpecifications)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification",
		reflect.TypeOf((*EmrClusterMasterInstanceFleetLaunchSpecificationsOnDemandSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference",
		reflect.TypeOf((*EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "onDemandSpecification", GoGetter: "OnDemandSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "onDemandSpecificationInput", GoGetter: "OnDemandSpecificationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnDemandSpecification", GoMethod: "ResetOnDemandSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpotSpecification", GoMethod: "ResetSpotSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "spotSpecification", GoGetter: "SpotSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "spotSpecificationInput", GoGetter: "SpotSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_EmrClusterMasterInstanceFleetLaunchSpecificationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification",
		reflect.TypeOf((*EmrClusterMasterInstanceFleetLaunchSpecificationsSpotSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceFleetOutputReference",
		reflect.TypeOf((*EmrClusterMasterInstanceFleetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeConfigs", GoGetter: "InstanceTypeConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeConfigsInput", GoGetter: "InstanceTypeConfigsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "launchSpecifications", GoGetter: "LaunchSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "launchSpecificationsInput", GoGetter: "LaunchSpecificationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLaunchSpecifications", GoMethod: "PutLaunchSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceTypeConfigs", GoMethod: "ResetInstanceTypeConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetLaunchSpecifications", GoMethod: "ResetLaunchSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetOnDemandCapacity", GoMethod: "ResetTargetOnDemandCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetSpotCapacity", GoMethod: "ResetTargetSpotCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "targetOnDemandCapacity", GoGetter: "TargetOnDemandCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "targetOnDemandCapacityInput", GoGetter: "TargetOnDemandCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetSpotCapacity", GoGetter: "TargetSpotCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "targetSpotCapacityInput", GoGetter: "TargetSpotCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_EmrClusterMasterInstanceFleetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceGroup",
		reflect.TypeOf((*EmrClusterMasterInstanceGroup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceGroupEbsConfig",
		reflect.TypeOf((*EmrClusterMasterInstanceGroupEbsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrClusterMasterInstanceGroupOutputReference",
		reflect.TypeOf((*EmrClusterMasterInstanceGroupOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bidPrice", GoGetter: "BidPrice"},
			_jsii_.MemberProperty{JsiiProperty: "bidPriceInput", GoGetter: "BidPriceInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebsConfig", GoGetter: "EbsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "ebsConfigInput", GoGetter: "EbsConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCount", GoGetter: "InstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCountInput", GoGetter: "InstanceCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBidPrice", GoMethod: "ResetBidPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsConfig", GoMethod: "ResetEbsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceCount", GoMethod: "ResetInstanceCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_EmrClusterMasterInstanceGroupOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterStep",
		reflect.TypeOf((*EmrClusterStep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrClusterStepHadoopJarStep",
		reflect.TypeOf((*EmrClusterStepHadoopJarStep)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrInstanceFleet",
		reflect.TypeOf((*EmrInstanceFleet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterId", GoGetter: "ClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdInput", GoGetter: "ClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeConfigs", GoGetter: "InstanceTypeConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeConfigsInput", GoGetter: "InstanceTypeConfigsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "launchSpecifications", GoGetter: "LaunchSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "launchSpecificationsInput", GoGetter: "LaunchSpecificationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedOnDemandCapacity", GoGetter: "ProvisionedOnDemandCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedSpotCapacity", GoGetter: "ProvisionedSpotCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "putLaunchSpecifications", GoMethod: "PutLaunchSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceTypeConfigs", GoMethod: "ResetInstanceTypeConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetLaunchSpecifications", GoMethod: "ResetLaunchSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetOnDemandCapacity", GoMethod: "ResetTargetOnDemandCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetSpotCapacity", GoMethod: "ResetTargetSpotCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targetOnDemandCapacity", GoGetter: "TargetOnDemandCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "targetOnDemandCapacityInput", GoGetter: "TargetOnDemandCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetSpotCapacity", GoGetter: "TargetSpotCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "targetSpotCapacityInput", GoGetter: "TargetSpotCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_EmrInstanceFleet{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrInstanceFleetConfig",
		reflect.TypeOf((*EmrInstanceFleetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrInstanceFleetInstanceTypeConfigs",
		reflect.TypeOf((*EmrInstanceFleetInstanceTypeConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrInstanceFleetInstanceTypeConfigsConfigurations",
		reflect.TypeOf((*EmrInstanceFleetInstanceTypeConfigsConfigurations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrInstanceFleetInstanceTypeConfigsEbsConfig",
		reflect.TypeOf((*EmrInstanceFleetInstanceTypeConfigsEbsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrInstanceFleetLaunchSpecifications",
		reflect.TypeOf((*EmrInstanceFleetLaunchSpecifications)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrInstanceFleetLaunchSpecificationsOnDemandSpecification",
		reflect.TypeOf((*EmrInstanceFleetLaunchSpecificationsOnDemandSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrInstanceFleetLaunchSpecificationsOutputReference",
		reflect.TypeOf((*EmrInstanceFleetLaunchSpecificationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "onDemandSpecification", GoGetter: "OnDemandSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "onDemandSpecificationInput", GoGetter: "OnDemandSpecificationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnDemandSpecification", GoMethod: "ResetOnDemandSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpotSpecification", GoMethod: "ResetSpotSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "spotSpecification", GoGetter: "SpotSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "spotSpecificationInput", GoGetter: "SpotSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_EmrInstanceFleetLaunchSpecificationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrInstanceFleetLaunchSpecificationsSpotSpecification",
		reflect.TypeOf((*EmrInstanceFleetLaunchSpecificationsSpotSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrInstanceGroup",
		reflect.TypeOf((*EmrInstanceGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingPolicy", GoGetter: "AutoscalingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingPolicyInput", GoGetter: "AutoscalingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "bidPrice", GoGetter: "BidPrice"},
			_jsii_.MemberProperty{JsiiProperty: "bidPriceInput", GoGetter: "BidPriceInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterId", GoGetter: "ClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdInput", GoGetter: "ClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "configurationsJson", GoGetter: "ConfigurationsJson"},
			_jsii_.MemberProperty{JsiiProperty: "configurationsJsonInput", GoGetter: "ConfigurationsJsonInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "ebsConfig", GoGetter: "EbsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "ebsConfigInput", GoGetter: "EbsConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimized", GoGetter: "EbsOptimized"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimizedInput", GoGetter: "EbsOptimizedInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCount", GoGetter: "InstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCountInput", GoGetter: "InstanceCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoscalingPolicy", GoMethod: "ResetAutoscalingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetBidPrice", GoMethod: "ResetBidPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigurationsJson", GoMethod: "ResetConfigurationsJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsConfig", GoMethod: "ResetEbsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsOptimized", GoMethod: "ResetEbsOptimized"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceCount", GoMethod: "ResetInstanceCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "runningInstanceCount", GoGetter: "RunningInstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_EmrInstanceGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrInstanceGroupConfig",
		reflect.TypeOf((*EmrInstanceGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrInstanceGroupEbsConfig",
		reflect.TypeOf((*EmrInstanceGroupEbsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrManagedScalingPolicy",
		reflect.TypeOf((*EmrManagedScalingPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterId", GoGetter: "ClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdInput", GoGetter: "ClusterIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeLimits", GoGetter: "ComputeLimits"},
			_jsii_.MemberProperty{JsiiProperty: "computeLimitsInput", GoGetter: "ComputeLimitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_EmrManagedScalingPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrManagedScalingPolicyComputeLimits",
		reflect.TypeOf((*EmrManagedScalingPolicyComputeLimits)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrManagedScalingPolicyConfig",
		reflect.TypeOf((*EmrManagedScalingPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.EMR.EmrSecurityConfiguration",
		reflect.TypeOf((*EmrSecurityConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "configurationInput", GoGetter: "ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creationDate", GoGetter: "CreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefix", GoGetter: "NamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefixInput", GoGetter: "NamePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamePrefix", GoMethod: "ResetNamePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_EmrSecurityConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.EMR.EmrSecurityConfigurationConfig",
		reflect.TypeOf((*EmrSecurityConfigurationConfig)(nil)).Elem(),
	)
}
