package codedeploy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployApp",
		reflect.TypeOf((*CodedeployApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "computePlatform", GoGetter: "ComputePlatform"},
			_jsii_.MemberProperty{JsiiProperty: "computePlatformInput", GoGetter: "ComputePlatformInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "githubAccountName", GoGetter: "GithubAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "linkedToGithub", GoGetter: "LinkedToGithub"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputePlatform", GoMethod: "ResetComputePlatform"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployApp{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployAppConfig",
		reflect.TypeOf((*CodedeployAppConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfig",
		reflect.TypeOf((*CodedeployDeploymentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "computePlatform", GoGetter: "ComputePlatform"},
			_jsii_.MemberProperty{JsiiProperty: "computePlatformInput", GoGetter: "ComputePlatformInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigId", GoGetter: "DeploymentConfigId"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigNameInput", GoGetter: "DeploymentConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "minimumHealthyHosts", GoGetter: "MinimumHealthyHosts"},
			_jsii_.MemberProperty{JsiiProperty: "minimumHealthyHostsInput", GoGetter: "MinimumHealthyHostsInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putMinimumHealthyHosts", GoMethod: "PutMinimumHealthyHosts"},
			_jsii_.MemberMethod{JsiiMethod: "putTrafficRoutingConfig", GoMethod: "PutTrafficRoutingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputePlatform", GoMethod: "ResetComputePlatform"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumHealthyHosts", GoMethod: "ResetMinimumHealthyHosts"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficRoutingConfig", GoMethod: "ResetTrafficRoutingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "trafficRoutingConfig", GoGetter: "TrafficRoutingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "trafficRoutingConfigInput", GoGetter: "TrafficRoutingConfigInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigConfig",
		reflect.TypeOf((*CodedeployDeploymentConfigConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigMinimumHealthyHosts",
		reflect.TypeOf((*CodedeployDeploymentConfigMinimumHealthyHosts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigMinimumHealthyHostsOutputReference",
		reflect.TypeOf((*CodedeployDeploymentConfigMinimumHealthyHostsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentConfigMinimumHealthyHostsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfig",
		reflect.TypeOf((*CodedeployDeploymentConfigTrafficRoutingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigOutputReference",
		reflect.TypeOf((*CodedeployDeploymentConfigTrafficRoutingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeBasedCanary", GoMethod: "PutTimeBasedCanary"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeBasedLinear", GoMethod: "PutTimeBasedLinear"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeBasedCanary", GoMethod: "ResetTimeBasedCanary"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeBasedLinear", GoMethod: "ResetTimeBasedLinear"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeBasedCanary", GoGetter: "TimeBasedCanary"},
			_jsii_.MemberProperty{JsiiProperty: "timeBasedCanaryInput", GoGetter: "TimeBasedCanaryInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeBasedLinear", GoGetter: "TimeBasedLinear"},
			_jsii_.MemberProperty{JsiiProperty: "timeBasedLinearInput", GoGetter: "TimeBasedLinearInput"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary",
		reflect.TypeOf((*CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference",
		reflect.TypeOf((*CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInput", GoGetter: "IntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "percentage", GoGetter: "Percentage"},
			_jsii_.MemberProperty{JsiiProperty: "percentageInput", GoGetter: "PercentageInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInterval", GoMethod: "ResetInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetPercentage", GoMethod: "ResetPercentage"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear",
		reflect.TypeOf((*CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference",
		reflect.TypeOf((*CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInput", GoGetter: "IntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "percentage", GoGetter: "Percentage"},
			_jsii_.MemberProperty{JsiiProperty: "percentageInput", GoGetter: "PercentageInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInterval", GoMethod: "ResetInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetPercentage", GoMethod: "ResetPercentage"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroup",
		reflect.TypeOf((*CodedeployDeploymentGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alarmConfiguration", GoGetter: "AlarmConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "alarmConfigurationInput", GoGetter: "AlarmConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "appName", GoGetter: "AppName"},
			_jsii_.MemberProperty{JsiiProperty: "appNameInput", GoGetter: "AppNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "autoRollbackConfiguration", GoGetter: "AutoRollbackConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "autoRollbackConfigurationInput", GoGetter: "AutoRollbackConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroups", GoGetter: "AutoscalingGroups"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingGroupsInput", GoGetter: "AutoscalingGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "blueGreenDeploymentConfig", GoGetter: "BlueGreenDeploymentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "blueGreenDeploymentConfigInput", GoGetter: "BlueGreenDeploymentConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "computePlatform", GoGetter: "ComputePlatform"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigName", GoGetter: "DeploymentConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigNameInput", GoGetter: "DeploymentConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupId", GoGetter: "DeploymentGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupName", GoGetter: "DeploymentGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentGroupNameInput", GoGetter: "DeploymentGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStyle", GoGetter: "DeploymentStyle"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentStyleInput", GoGetter: "DeploymentStyleInput"},
			_jsii_.MemberProperty{JsiiProperty: "ec2TagFilter", GoGetter: "Ec2TagFilter"},
			_jsii_.MemberProperty{JsiiProperty: "ec2TagFilterInput", GoGetter: "Ec2TagFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "ec2TagSet", GoGetter: "Ec2TagSet"},
			_jsii_.MemberProperty{JsiiProperty: "ec2TagSetInput", GoGetter: "Ec2TagSetInput"},
			_jsii_.MemberProperty{JsiiProperty: "ecsService", GoGetter: "EcsService"},
			_jsii_.MemberProperty{JsiiProperty: "ecsServiceInput", GoGetter: "EcsServiceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerInfo", GoGetter: "LoadBalancerInfo"},
			_jsii_.MemberProperty{JsiiProperty: "loadBalancerInfoInput", GoGetter: "LoadBalancerInfoInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onPremisesInstanceTagFilter", GoGetter: "OnPremisesInstanceTagFilter"},
			_jsii_.MemberProperty{JsiiProperty: "onPremisesInstanceTagFilterInput", GoGetter: "OnPremisesInstanceTagFilterInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putAlarmConfiguration", GoMethod: "PutAlarmConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putAutoRollbackConfiguration", GoMethod: "PutAutoRollbackConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putBlueGreenDeploymentConfig", GoMethod: "PutBlueGreenDeploymentConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putDeploymentStyle", GoMethod: "PutDeploymentStyle"},
			_jsii_.MemberMethod{JsiiMethod: "putEcsService", GoMethod: "PutEcsService"},
			_jsii_.MemberMethod{JsiiMethod: "putLoadBalancerInfo", GoMethod: "PutLoadBalancerInfo"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlarmConfiguration", GoMethod: "ResetAlarmConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRollbackConfiguration", GoMethod: "ResetAutoRollbackConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoscalingGroups", GoMethod: "ResetAutoscalingGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlueGreenDeploymentConfig", GoMethod: "ResetBlueGreenDeploymentConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentConfigName", GoMethod: "ResetDeploymentConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentStyle", GoMethod: "ResetDeploymentStyle"},
			_jsii_.MemberMethod{JsiiMethod: "resetEc2TagFilter", GoMethod: "ResetEc2TagFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetEc2TagSet", GoMethod: "ResetEc2TagSet"},
			_jsii_.MemberMethod{JsiiMethod: "resetEcsService", GoMethod: "ResetEcsService"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoadBalancerInfo", GoMethod: "ResetLoadBalancerInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnPremisesInstanceTagFilter", GoMethod: "ResetOnPremisesInstanceTagFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggerConfiguration", GoMethod: "ResetTriggerConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArn", GoGetter: "ServiceRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArnInput", GoGetter: "ServiceRoleArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfiguration", GoGetter: "TriggerConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "triggerConfigurationInput", GoGetter: "TriggerConfigurationInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupAlarmConfiguration",
		reflect.TypeOf((*CodedeployDeploymentGroupAlarmConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupAlarmConfigurationOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupAlarmConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarms", GoGetter: "Alarms"},
			_jsii_.MemberProperty{JsiiProperty: "alarmsInput", GoGetter: "AlarmsInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ignorePollAlarmFailure", GoGetter: "IgnorePollAlarmFailure"},
			_jsii_.MemberProperty{JsiiProperty: "ignorePollAlarmFailureInput", GoGetter: "IgnorePollAlarmFailureInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlarms", GoMethod: "ResetAlarms"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnorePollAlarmFailure", GoMethod: "ResetIgnorePollAlarmFailure"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupAlarmConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupAutoRollbackConfiguration",
		reflect.TypeOf((*CodedeployDeploymentGroupAutoRollbackConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "events", GoGetter: "Events"},
			_jsii_.MemberProperty{JsiiProperty: "eventsInput", GoGetter: "EventsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEvents", GoMethod: "ResetEvents"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupAutoRollbackConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfig",
		reflect.TypeOf((*CodedeployDeploymentGroupBlueGreenDeploymentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption",
		reflect.TypeOf((*CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOption)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionOnTimeout", GoGetter: "ActionOnTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "actionOnTimeoutInput", GoGetter: "ActionOnTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetActionOnTimeout", GoMethod: "ResetActionOnTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetWaitTimeInMinutes", GoMethod: "ResetWaitTimeInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "waitTimeInMinutes", GoGetter: "WaitTimeInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "waitTimeInMinutesInput", GoGetter: "WaitTimeInMinutesInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption",
		reflect.TypeOf((*CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOption)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetAction", GoMethod: "ResetAction"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentReadyOption", GoGetter: "DeploymentReadyOption"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentReadyOptionInput", GoGetter: "DeploymentReadyOptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "greenFleetProvisioningOption", GoGetter: "GreenFleetProvisioningOption"},
			_jsii_.MemberProperty{JsiiProperty: "greenFleetProvisioningOptionInput", GoGetter: "GreenFleetProvisioningOptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putDeploymentReadyOption", GoMethod: "PutDeploymentReadyOption"},
			_jsii_.MemberMethod{JsiiMethod: "putGreenFleetProvisioningOption", GoMethod: "PutGreenFleetProvisioningOption"},
			_jsii_.MemberMethod{JsiiMethod: "putTerminateBlueInstancesOnDeploymentSuccess", GoMethod: "PutTerminateBlueInstancesOnDeploymentSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentReadyOption", GoMethod: "ResetDeploymentReadyOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetGreenFleetProvisioningOption", GoMethod: "ResetGreenFleetProvisioningOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetTerminateBlueInstancesOnDeploymentSuccess", GoMethod: "ResetTerminateBlueInstancesOnDeploymentSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "terminateBlueInstancesOnDeploymentSuccess", GoGetter: "TerminateBlueInstancesOnDeploymentSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "terminateBlueInstancesOnDeploymentSuccessInput", GoGetter: "TerminateBlueInstancesOnDeploymentSuccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess",
		reflect.TypeOf((*CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetAction", GoMethod: "ResetAction"},
			_jsii_.MemberMethod{JsiiMethod: "resetTerminationWaitTimeInMinutes", GoMethod: "ResetTerminationWaitTimeInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "terminationWaitTimeInMinutes", GoGetter: "TerminationWaitTimeInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "terminationWaitTimeInMinutesInput", GoGetter: "TerminationWaitTimeInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupConfig",
		reflect.TypeOf((*CodedeployDeploymentGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupDeploymentStyle",
		reflect.TypeOf((*CodedeployDeploymentGroupDeploymentStyle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupDeploymentStyleOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupDeploymentStyleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deploymentOption", GoGetter: "DeploymentOption"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentOptionInput", GoGetter: "DeploymentOptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentType", GoGetter: "DeploymentType"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentTypeInput", GoGetter: "DeploymentTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentOption", GoMethod: "ResetDeploymentOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentType", GoMethod: "ResetDeploymentType"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupDeploymentStyleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupEc2TagFilter",
		reflect.TypeOf((*CodedeployDeploymentGroupEc2TagFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupEc2TagSet",
		reflect.TypeOf((*CodedeployDeploymentGroupEc2TagSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupEc2TagSetEc2TagFilter",
		reflect.TypeOf((*CodedeployDeploymentGroupEc2TagSetEc2TagFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupEcsService",
		reflect.TypeOf((*CodedeployDeploymentGroupEcsService)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupEcsServiceOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupEcsServiceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterNameInput", GoGetter: "ClusterNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNameInput", GoGetter: "ServiceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupEcsServiceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfo",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoElbInfo",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoElbInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "elbInfo", GoGetter: "ElbInfo"},
			_jsii_.MemberProperty{JsiiProperty: "elbInfoInput", GoGetter: "ElbInfoInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putTargetGroupPairInfo", GoMethod: "PutTargetGroupPairInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetElbInfo", GoMethod: "ResetElbInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetGroupInfo", GoMethod: "ResetTargetGroupInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetGroupPairInfo", GoMethod: "ResetTargetGroupPairInfo"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupInfo", GoGetter: "TargetGroupInfo"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupInfoInput", GoGetter: "TargetGroupInfoInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupPairInfo", GoGetter: "TargetGroupPairInfo"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupPairInfoInput", GoGetter: "TargetGroupPairInfoInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "prodTrafficRoute", GoGetter: "ProdTrafficRoute"},
			_jsii_.MemberProperty{JsiiProperty: "prodTrafficRouteInput", GoGetter: "ProdTrafficRouteInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProdTrafficRoute", GoMethod: "PutProdTrafficRoute"},
			_jsii_.MemberMethod{JsiiMethod: "putTestTrafficRoute", GoMethod: "PutTestTrafficRoute"},
			_jsii_.MemberMethod{JsiiMethod: "resetTestTrafficRoute", GoMethod: "ResetTestTrafficRoute"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroup", GoGetter: "TargetGroup"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroupInput", GoGetter: "TargetGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "testTrafficRoute", GoGetter: "TestTrafficRoute"},
			_jsii_.MemberProperty{JsiiProperty: "testTrafficRouteInput", GoGetter: "TestTrafficRouteInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArns", GoGetter: "ListenerArns"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArnsInput", GoGetter: "ListenerArnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference",
		reflect.TypeOf((*CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArns", GoGetter: "ListenerArns"},
			_jsii_.MemberProperty{JsiiProperty: "listenerArnsInput", GoGetter: "ListenerArnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoTestTrafficRouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupOnPremisesInstanceTagFilter",
		reflect.TypeOf((*CodedeployDeploymentGroupOnPremisesInstanceTagFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeDeploy.CodedeployDeploymentGroupTriggerConfiguration",
		reflect.TypeOf((*CodedeployDeploymentGroupTriggerConfiguration)(nil)).Elem(),
	)
}
