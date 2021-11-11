package mq

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.DataAwsMqBroker",
		reflect.TypeOf((*DataAwsMqBroker)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationStrategy", GoGetter: "AuthenticationStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "autoMinorVersionUpgrade", GoGetter: "AutoMinorVersionUpgrade"},
			_jsii_.MemberProperty{JsiiProperty: "brokerId", GoGetter: "BrokerId"},
			_jsii_.MemberProperty{JsiiProperty: "brokerIdInput", GoGetter: "BrokerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "brokerName", GoGetter: "BrokerName"},
			_jsii_.MemberProperty{JsiiProperty: "brokerNameInput", GoGetter: "BrokerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberMethod{JsiiMethod: "configuration", GoMethod: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentMode", GoGetter: "DeploymentMode"},
			_jsii_.MemberMethod{JsiiMethod: "encryptionOptions", GoMethod: "EncryptionOptions"},
			_jsii_.MemberProperty{JsiiProperty: "engineType", GoGetter: "EngineType"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hostInstanceType", GoGetter: "HostInstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "instances", GoMethod: "Instances"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "ldapServerMetadata", GoMethod: "LdapServerMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "logs", GoMethod: "Logs"},
			_jsii_.MemberMethod{JsiiMethod: "maintenanceWindowStartTime", GoMethod: "MaintenanceWindowStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAccessible", GoGetter: "PubliclyAccessible"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBrokerId", GoMethod: "ResetBrokerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetBrokerName", GoMethod: "ResetBrokerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "storageType", GoGetter: "StorageType"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "user", GoMethod: "User"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsMqBroker{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.MQ.DataAwsMqBrokerConfig",
		reflect.TypeOf((*DataAwsMqBrokerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.DataAwsMqBrokerConfiguration",
		reflect.TypeOf((*DataAwsMqBrokerConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "revision", GoGetter: "Revision"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsMqBrokerConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.DataAwsMqBrokerEncryptionOptions",
		reflect.TypeOf((*DataAwsMqBrokerEncryptionOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "useAwsOwnedKey", GoGetter: "UseAwsOwnedKey"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsMqBrokerEncryptionOptions{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.DataAwsMqBrokerInstances",
		reflect.TypeOf((*DataAwsMqBrokerInstances)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "consoleUrl", GoGetter: "ConsoleUrl"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsMqBrokerInstances{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.DataAwsMqBrokerLdapServerMetadata",
		reflect.TypeOf((*DataAwsMqBrokerLdapServerMetadata)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hosts", GoGetter: "Hosts"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "roleBase", GoGetter: "RoleBase"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberProperty{JsiiProperty: "roleSearchMatching", GoGetter: "RoleSearchMatching"},
			_jsii_.MemberProperty{JsiiProperty: "roleSearchSubtree", GoGetter: "RoleSearchSubtree"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountPassword", GoGetter: "ServiceAccountPassword"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountUsername", GoGetter: "ServiceAccountUsername"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userBase", GoGetter: "UserBase"},
			_jsii_.MemberProperty{JsiiProperty: "userRoleName", GoGetter: "UserRoleName"},
			_jsii_.MemberProperty{JsiiProperty: "userSearchMatching", GoGetter: "UserSearchMatching"},
			_jsii_.MemberProperty{JsiiProperty: "userSearchSubtree", GoGetter: "UserSearchSubtree"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsMqBrokerLdapServerMetadata{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.DataAwsMqBrokerLogs",
		reflect.TypeOf((*DataAwsMqBrokerLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "audit", GoGetter: "Audit"},
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "general", GoGetter: "General"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsMqBrokerLogs{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.DataAwsMqBrokerMaintenanceWindowStartTime",
		reflect.TypeOf((*DataAwsMqBrokerMaintenanceWindowStartTime)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "dayOfWeek", GoGetter: "DayOfWeek"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfDay", GoGetter: "TimeOfDay"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsMqBrokerMaintenanceWindowStartTime{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.DataAwsMqBrokerUser",
		reflect.TypeOf((*DataAwsMqBrokerUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "consoleAccess", GoGetter: "ConsoleAccess"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groups", GoGetter: "Groups"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsMqBrokerUser{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.MqBroker",
		reflect.TypeOf((*MqBroker)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applyImmediately", GoGetter: "ApplyImmediately"},
			_jsii_.MemberProperty{JsiiProperty: "applyImmediatelyInput", GoGetter: "ApplyImmediatelyInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationStrategy", GoGetter: "AuthenticationStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationStrategyInput", GoGetter: "AuthenticationStrategyInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoMinorVersionUpgrade", GoGetter: "AutoMinorVersionUpgrade"},
			_jsii_.MemberProperty{JsiiProperty: "autoMinorVersionUpgradeInput", GoGetter: "AutoMinorVersionUpgradeInput"},
			_jsii_.MemberProperty{JsiiProperty: "brokerName", GoGetter: "BrokerName"},
			_jsii_.MemberProperty{JsiiProperty: "brokerNameInput", GoGetter: "BrokerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "configurationInput", GoGetter: "ConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentMode", GoGetter: "DeploymentMode"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentModeInput", GoGetter: "DeploymentModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionOptions", GoGetter: "EncryptionOptions"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionOptionsInput", GoGetter: "EncryptionOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "engineType", GoGetter: "EngineType"},
			_jsii_.MemberProperty{JsiiProperty: "engineTypeInput", GoGetter: "EngineTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersionInput", GoGetter: "EngineVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hostInstanceType", GoGetter: "HostInstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "hostInstanceTypeInput", GoGetter: "HostInstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "instances", GoMethod: "Instances"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ldapServerMetadata", GoGetter: "LdapServerMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "ldapServerMetadataInput", GoGetter: "LdapServerMetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logs", GoGetter: "Logs"},
			_jsii_.MemberProperty{JsiiProperty: "logsInput", GoGetter: "LogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowStartTime", GoGetter: "MaintenanceWindowStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceWindowStartTimeInput", GoGetter: "MaintenanceWindowStartTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAccessible", GoGetter: "PubliclyAccessible"},
			_jsii_.MemberProperty{JsiiProperty: "publiclyAccessibleInput", GoGetter: "PubliclyAccessibleInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConfiguration", GoMethod: "PutConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionOptions", GoMethod: "PutEncryptionOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putLdapServerMetadata", GoMethod: "PutLdapServerMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "putLogs", GoMethod: "PutLogs"},
			_jsii_.MemberMethod{JsiiMethod: "putMaintenanceWindowStartTime", GoMethod: "PutMaintenanceWindowStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplyImmediately", GoMethod: "ResetApplyImmediately"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthenticationStrategy", GoMethod: "ResetAuthenticationStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoMinorVersionUpgrade", GoMethod: "ResetAutoMinorVersionUpgrade"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfiguration", GoMethod: "ResetConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentMode", GoMethod: "ResetDeploymentMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionOptions", GoMethod: "ResetEncryptionOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetLdapServerMetadata", GoMethod: "ResetLdapServerMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogs", GoMethod: "ResetLogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceWindowStartTime", GoMethod: "ResetMaintenanceWindowStartTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPubliclyAccessible", GoMethod: "ResetPubliclyAccessible"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroups", GoMethod: "ResetSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageType", GoMethod: "ResetStorageType"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetIds", GoMethod: "ResetSubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupsInput", GoGetter: "SecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "storageType", GoGetter: "StorageType"},
			_jsii_.MemberProperty{JsiiProperty: "storageTypeInput", GoGetter: "StorageTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userInput", GoGetter: "UserInput"},
		},
		func() interface{} {
			j := jsiiProxy_MqBroker{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.MQ.MqBrokerConfig",
		reflect.TypeOf((*MqBrokerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.MQ.MqBrokerConfiguration",
		reflect.TypeOf((*MqBrokerConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.MqBrokerConfigurationOutputReference",
		reflect.TypeOf((*MqBrokerConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetRevision", GoMethod: "ResetRevision"},
			_jsii_.MemberProperty{JsiiProperty: "revision", GoGetter: "Revision"},
			_jsii_.MemberProperty{JsiiProperty: "revisionInput", GoGetter: "RevisionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MqBrokerConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.MQ.MqBrokerEncryptionOptions",
		reflect.TypeOf((*MqBrokerEncryptionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.MqBrokerEncryptionOptionsOutputReference",
		reflect.TypeOf((*MqBrokerEncryptionOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseAwsOwnedKey", GoMethod: "ResetUseAwsOwnedKey"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "useAwsOwnedKey", GoGetter: "UseAwsOwnedKey"},
			_jsii_.MemberProperty{JsiiProperty: "useAwsOwnedKeyInput", GoGetter: "UseAwsOwnedKeyInput"},
		},
		func() interface{} {
			j := jsiiProxy_MqBrokerEncryptionOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.MqBrokerInstances",
		reflect.TypeOf((*MqBrokerInstances)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "consoleUrl", GoGetter: "ConsoleUrl"},
			_jsii_.MemberProperty{JsiiProperty: "endpoints", GoGetter: "Endpoints"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MqBrokerInstances{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.MQ.MqBrokerLdapServerMetadata",
		reflect.TypeOf((*MqBrokerLdapServerMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.MqBrokerLdapServerMetadataOutputReference",
		reflect.TypeOf((*MqBrokerLdapServerMetadataOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hosts", GoGetter: "Hosts"},
			_jsii_.MemberProperty{JsiiProperty: "hostsInput", GoGetter: "HostsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetHosts", GoMethod: "ResetHosts"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleBase", GoMethod: "ResetRoleBase"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleName", GoMethod: "ResetRoleName"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleSearchMatching", GoMethod: "ResetRoleSearchMatching"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleSearchSubtree", GoMethod: "ResetRoleSearchSubtree"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccountPassword", GoMethod: "ResetServiceAccountPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccountUsername", GoMethod: "ResetServiceAccountUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserBase", GoMethod: "ResetUserBase"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserRoleName", GoMethod: "ResetUserRoleName"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserSearchMatching", GoMethod: "ResetUserSearchMatching"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserSearchSubtree", GoMethod: "ResetUserSearchSubtree"},
			_jsii_.MemberProperty{JsiiProperty: "roleBase", GoGetter: "RoleBase"},
			_jsii_.MemberProperty{JsiiProperty: "roleBaseInput", GoGetter: "RoleBaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleName", GoGetter: "RoleName"},
			_jsii_.MemberProperty{JsiiProperty: "roleNameInput", GoGetter: "RoleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleSearchMatching", GoGetter: "RoleSearchMatching"},
			_jsii_.MemberProperty{JsiiProperty: "roleSearchMatchingInput", GoGetter: "RoleSearchMatchingInput"},
			_jsii_.MemberProperty{JsiiProperty: "roleSearchSubtree", GoGetter: "RoleSearchSubtree"},
			_jsii_.MemberProperty{JsiiProperty: "roleSearchSubtreeInput", GoGetter: "RoleSearchSubtreeInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountPassword", GoGetter: "ServiceAccountPassword"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountPasswordInput", GoGetter: "ServiceAccountPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountUsername", GoGetter: "ServiceAccountUsername"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountUsernameInput", GoGetter: "ServiceAccountUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userBase", GoGetter: "UserBase"},
			_jsii_.MemberProperty{JsiiProperty: "userBaseInput", GoGetter: "UserBaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "userRoleName", GoGetter: "UserRoleName"},
			_jsii_.MemberProperty{JsiiProperty: "userRoleNameInput", GoGetter: "UserRoleNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "userSearchMatching", GoGetter: "UserSearchMatching"},
			_jsii_.MemberProperty{JsiiProperty: "userSearchMatchingInput", GoGetter: "UserSearchMatchingInput"},
			_jsii_.MemberProperty{JsiiProperty: "userSearchSubtree", GoGetter: "UserSearchSubtree"},
			_jsii_.MemberProperty{JsiiProperty: "userSearchSubtreeInput", GoGetter: "UserSearchSubtreeInput"},
		},
		func() interface{} {
			j := jsiiProxy_MqBrokerLdapServerMetadataOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.MQ.MqBrokerLogs",
		reflect.TypeOf((*MqBrokerLogs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.MqBrokerLogsOutputReference",
		reflect.TypeOf((*MqBrokerLogsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "audit", GoGetter: "Audit"},
			_jsii_.MemberProperty{JsiiProperty: "auditInput", GoGetter: "AuditInput"},
			_jsii_.MemberProperty{JsiiProperty: "general", GoGetter: "General"},
			_jsii_.MemberProperty{JsiiProperty: "generalInput", GoGetter: "GeneralInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudit", GoMethod: "ResetAudit"},
			_jsii_.MemberMethod{JsiiMethod: "resetGeneral", GoMethod: "ResetGeneral"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MqBrokerLogsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.MQ.MqBrokerMaintenanceWindowStartTime",
		reflect.TypeOf((*MqBrokerMaintenanceWindowStartTime)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.MqBrokerMaintenanceWindowStartTimeOutputReference",
		reflect.TypeOf((*MqBrokerMaintenanceWindowStartTimeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dayOfWeek", GoGetter: "DayOfWeek"},
			_jsii_.MemberProperty{JsiiProperty: "dayOfWeekInput", GoGetter: "DayOfWeekInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfDay", GoGetter: "TimeOfDay"},
			_jsii_.MemberProperty{JsiiProperty: "timeOfDayInput", GoGetter: "TimeOfDayInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeZone", GoGetter: "TimeZone"},
			_jsii_.MemberProperty{JsiiProperty: "timeZoneInput", GoGetter: "TimeZoneInput"},
		},
		func() interface{} {
			j := jsiiProxy_MqBrokerMaintenanceWindowStartTimeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.MQ.MqBrokerUser",
		reflect.TypeOf((*MqBrokerUser)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.MQ.MqConfiguration",
		reflect.TypeOf((*MqConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationStrategy", GoGetter: "AuthenticationStrategy"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationStrategyInput", GoGetter: "AuthenticationStrategyInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "dataInput", GoGetter: "DataInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "engineType", GoGetter: "EngineType"},
			_jsii_.MemberProperty{JsiiProperty: "engineTypeInput", GoGetter: "EngineTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersion", GoGetter: "EngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "engineVersionInput", GoGetter: "EngineVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "latestRevision", GoGetter: "LatestRevision"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthenticationStrategy", GoMethod: "ResetAuthenticationStrategy"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
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
			j := jsiiProxy_MqConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.MQ.MqConfigurationConfig",
		reflect.TypeOf((*MqConfigurationConfig)(nil)).Elem(),
	)
}
