package workspaces

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesBundle",
		reflect.TypeOf((*DataAwsWorkspacesBundle)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "bundleId", GoGetter: "BundleId"},
			_jsii_.MemberProperty{JsiiProperty: "bundleIdInput", GoGetter: "BundleIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberMethod{JsiiMethod: "computeType", GoMethod: "ComputeType"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "ownerInput", GoGetter: "OwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBundleId", GoMethod: "ResetBundleId"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOwner", GoMethod: "ResetOwner"},
			_jsii_.MemberMethod{JsiiMethod: "rootStorage", GoMethod: "RootStorage"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "userStorage", GoMethod: "UserStorage"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesBundle{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesBundleComputeType",
		reflect.TypeOf((*DataAwsWorkspacesBundleComputeType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesBundleComputeType{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesBundleConfig",
		reflect.TypeOf((*DataAwsWorkspacesBundleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesBundleRootStorage",
		reflect.TypeOf((*DataAwsWorkspacesBundleRootStorage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "capacity", GoGetter: "Capacity"},
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesBundleRootStorage{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesBundleUserStorage",
		reflect.TypeOf((*DataAwsWorkspacesBundleUserStorage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "capacity", GoGetter: "Capacity"},
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesBundleUserStorage{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesDirectory",
		reflect.TypeOf((*DataAwsWorkspacesDirectory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customerUserName", GoGetter: "CustomerUserName"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "directoryId", GoGetter: "DirectoryId"},
			_jsii_.MemberProperty{JsiiProperty: "directoryIdInput", GoGetter: "DirectoryIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "directoryName", GoGetter: "DirectoryName"},
			_jsii_.MemberProperty{JsiiProperty: "directoryType", GoGetter: "DirectoryType"},
			_jsii_.MemberProperty{JsiiProperty: "dnsIpAddresses", GoGetter: "DnsIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamRoleId", GoGetter: "IamRoleId"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipGroupIds", GoGetter: "IpGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "registrationCode", GoGetter: "RegistrationCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "selfServicePermissions", GoMethod: "SelfServicePermissions"},
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
			_jsii_.MemberMethod{JsiiMethod: "workspaceAccessProperties", GoMethod: "WorkspaceAccessProperties"},
			_jsii_.MemberMethod{JsiiMethod: "workspaceCreationProperties", GoMethod: "WorkspaceCreationProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceSecurityGroupId", GoGetter: "WorkspaceSecurityGroupId"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesDirectory{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesDirectoryConfig",
		reflect.TypeOf((*DataAwsWorkspacesDirectoryConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesDirectorySelfServicePermissions",
		reflect.TypeOf((*DataAwsWorkspacesDirectorySelfServicePermissions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "changeComputeType", GoGetter: "ChangeComputeType"},
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "increaseVolumeSize", GoGetter: "IncreaseVolumeSize"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "rebuildWorkspace", GoGetter: "RebuildWorkspace"},
			_jsii_.MemberProperty{JsiiProperty: "restartWorkspace", GoGetter: "RestartWorkspace"},
			_jsii_.MemberProperty{JsiiProperty: "switchRunningMode", GoGetter: "SwitchRunningMode"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesDirectoryWorkspaceAccessProperties",
		reflect.TypeOf((*DataAwsWorkspacesDirectoryWorkspaceAccessProperties)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeAndroid", GoGetter: "DeviceTypeAndroid"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeChromeos", GoGetter: "DeviceTypeChromeos"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeIos", GoGetter: "DeviceTypeIos"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeLinux", GoGetter: "DeviceTypeLinux"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeOsx", GoGetter: "DeviceTypeOsx"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWeb", GoGetter: "DeviceTypeWeb"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWindows", GoGetter: "DeviceTypeWindows"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeZeroclient", GoGetter: "DeviceTypeZeroclient"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesDirectoryWorkspaceCreationProperties",
		reflect.TypeOf((*DataAwsWorkspacesDirectoryWorkspaceCreationProperties)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "customSecurityGroupId", GoGetter: "CustomSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "defaultOu", GoGetter: "DefaultOu"},
			_jsii_.MemberProperty{JsiiProperty: "enableInternetAccess", GoGetter: "EnableInternetAccess"},
			_jsii_.MemberProperty{JsiiProperty: "enableMaintenanceMode", GoGetter: "EnableMaintenanceMode"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userEnabledAsLocalAdministrator", GoGetter: "UserEnabledAsLocalAdministrator"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesImage",
		reflect.TypeOf((*DataAwsWorkspacesImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "imageId", GoGetter: "ImageId"},
			_jsii_.MemberProperty{JsiiProperty: "imageIdInput", GoGetter: "ImageIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "operatingSystemType", GoGetter: "OperatingSystemType"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requiredTenancy", GoGetter: "RequiredTenancy"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesImage{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesImageConfig",
		reflect.TypeOf((*DataAwsWorkspacesImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesWorkspace",
		reflect.TypeOf((*DataAwsWorkspacesWorkspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "bundleId", GoGetter: "BundleId"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "computerName", GoGetter: "ComputerName"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "directoryId", GoGetter: "DirectoryId"},
			_jsii_.MemberProperty{JsiiProperty: "directoryIdInput", GoGetter: "DirectoryIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirectoryId", GoMethod: "ResetDirectoryId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserName", GoMethod: "ResetUserName"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceId", GoMethod: "ResetWorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "rootVolumeEncryptionEnabled", GoGetter: "RootVolumeEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
			_jsii_.MemberProperty{JsiiProperty: "userNameInput", GoGetter: "UserNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "userVolumeEncryptionEnabled", GoGetter: "UserVolumeEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "volumeEncryptionKey", GoGetter: "VolumeEncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceIdInput", GoGetter: "WorkspaceIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "workspaceProperties", GoMethod: "WorkspaceProperties"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesWorkspace{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesWorkspaceConfig",
		reflect.TypeOf((*DataAwsWorkspacesWorkspaceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.DataAwsWorkspacesWorkspaceWorkspaceProperties",
		reflect.TypeOf((*DataAwsWorkspacesWorkspaceWorkspaceProperties)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "computeTypeName", GoGetter: "ComputeTypeName"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "rootVolumeSizeGib", GoGetter: "RootVolumeSizeGib"},
			_jsii_.MemberProperty{JsiiProperty: "runningMode", GoGetter: "RunningMode"},
			_jsii_.MemberProperty{JsiiProperty: "runningModeAutoStopTimeoutInMinutes", GoGetter: "RunningModeAutoStopTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userVolumeSizeGib", GoGetter: "UserVolumeSizeGib"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.WorkspacesDirectory",
		reflect.TypeOf((*WorkspacesDirectory)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customerUserName", GoGetter: "CustomerUserName"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "directoryId", GoGetter: "DirectoryId"},
			_jsii_.MemberProperty{JsiiProperty: "directoryIdInput", GoGetter: "DirectoryIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "directoryName", GoGetter: "DirectoryName"},
			_jsii_.MemberProperty{JsiiProperty: "directoryType", GoGetter: "DirectoryType"},
			_jsii_.MemberProperty{JsiiProperty: "dnsIpAddresses", GoGetter: "DnsIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamRoleId", GoGetter: "IamRoleId"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipGroupIds", GoGetter: "IpGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "ipGroupIdsInput", GoGetter: "IpGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putSelfServicePermissions", GoMethod: "PutSelfServicePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceAccessProperties", GoMethod: "PutWorkspaceAccessProperties"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceCreationProperties", GoMethod: "PutWorkspaceCreationProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "registrationCode", GoGetter: "RegistrationCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpGroupIds", GoMethod: "ResetIpGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelfServicePermissions", GoMethod: "ResetSelfServicePermissions"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetIds", GoMethod: "ResetSubnetIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceAccessProperties", GoMethod: "ResetWorkspaceAccessProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceCreationProperties", GoMethod: "ResetWorkspaceCreationProperties"},
			_jsii_.MemberProperty{JsiiProperty: "selfServicePermissions", GoGetter: "SelfServicePermissions"},
			_jsii_.MemberProperty{JsiiProperty: "selfServicePermissionsInput", GoGetter: "SelfServicePermissionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "workspaceAccessProperties", GoGetter: "WorkspaceAccessProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceAccessPropertiesInput", GoGetter: "WorkspaceAccessPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceCreationProperties", GoGetter: "WorkspaceCreationProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceCreationPropertiesInput", GoGetter: "WorkspaceCreationPropertiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceSecurityGroupId", GoGetter: "WorkspaceSecurityGroupId"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesDirectory{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.WorkspacesDirectoryConfig",
		reflect.TypeOf((*WorkspacesDirectoryConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.WorkspacesDirectorySelfServicePermissions",
		reflect.TypeOf((*WorkspacesDirectorySelfServicePermissions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.WorkspacesDirectorySelfServicePermissionsOutputReference",
		reflect.TypeOf((*WorkspacesDirectorySelfServicePermissionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "changeComputeType", GoGetter: "ChangeComputeType"},
			_jsii_.MemberProperty{JsiiProperty: "changeComputeTypeInput", GoGetter: "ChangeComputeTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "increaseVolumeSize", GoGetter: "IncreaseVolumeSize"},
			_jsii_.MemberProperty{JsiiProperty: "increaseVolumeSizeInput", GoGetter: "IncreaseVolumeSizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "rebuildWorkspace", GoGetter: "RebuildWorkspace"},
			_jsii_.MemberProperty{JsiiProperty: "rebuildWorkspaceInput", GoGetter: "RebuildWorkspaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetChangeComputeType", GoMethod: "ResetChangeComputeType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncreaseVolumeSize", GoMethod: "ResetIncreaseVolumeSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetRebuildWorkspace", GoMethod: "ResetRebuildWorkspace"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestartWorkspace", GoMethod: "ResetRestartWorkspace"},
			_jsii_.MemberMethod{JsiiMethod: "resetSwitchRunningMode", GoMethod: "ResetSwitchRunningMode"},
			_jsii_.MemberProperty{JsiiProperty: "restartWorkspace", GoGetter: "RestartWorkspace"},
			_jsii_.MemberProperty{JsiiProperty: "restartWorkspaceInput", GoGetter: "RestartWorkspaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "switchRunningMode", GoGetter: "SwitchRunningMode"},
			_jsii_.MemberProperty{JsiiProperty: "switchRunningModeInput", GoGetter: "SwitchRunningModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.WorkspacesDirectoryWorkspaceAccessProperties",
		reflect.TypeOf((*WorkspacesDirectoryWorkspaceAccessProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference",
		reflect.TypeOf((*WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeAndroid", GoGetter: "DeviceTypeAndroid"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeAndroidInput", GoGetter: "DeviceTypeAndroidInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeChromeos", GoGetter: "DeviceTypeChromeos"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeChromeosInput", GoGetter: "DeviceTypeChromeosInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeIos", GoGetter: "DeviceTypeIos"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeIosInput", GoGetter: "DeviceTypeIosInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeLinux", GoGetter: "DeviceTypeLinux"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeLinuxInput", GoGetter: "DeviceTypeLinuxInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeOsx", GoGetter: "DeviceTypeOsx"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeOsxInput", GoGetter: "DeviceTypeOsxInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWeb", GoGetter: "DeviceTypeWeb"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWebInput", GoGetter: "DeviceTypeWebInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWindows", GoGetter: "DeviceTypeWindows"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeWindowsInput", GoGetter: "DeviceTypeWindowsInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeZeroclient", GoGetter: "DeviceTypeZeroclient"},
			_jsii_.MemberProperty{JsiiProperty: "deviceTypeZeroclientInput", GoGetter: "DeviceTypeZeroclientInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeAndroid", GoMethod: "ResetDeviceTypeAndroid"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeChromeos", GoMethod: "ResetDeviceTypeChromeos"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeIos", GoMethod: "ResetDeviceTypeIos"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeLinux", GoMethod: "ResetDeviceTypeLinux"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeOsx", GoMethod: "ResetDeviceTypeOsx"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeWeb", GoMethod: "ResetDeviceTypeWeb"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeWindows", GoMethod: "ResetDeviceTypeWindows"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceTypeZeroclient", GoMethod: "ResetDeviceTypeZeroclient"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.WorkspacesDirectoryWorkspaceCreationProperties",
		reflect.TypeOf((*WorkspacesDirectoryWorkspaceCreationProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference",
		reflect.TypeOf((*WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "customSecurityGroupId", GoGetter: "CustomSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "customSecurityGroupIdInput", GoGetter: "CustomSecurityGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultOu", GoGetter: "DefaultOu"},
			_jsii_.MemberProperty{JsiiProperty: "defaultOuInput", GoGetter: "DefaultOuInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableInternetAccess", GoGetter: "EnableInternetAccess"},
			_jsii_.MemberProperty{JsiiProperty: "enableInternetAccessInput", GoGetter: "EnableInternetAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableMaintenanceMode", GoGetter: "EnableMaintenanceMode"},
			_jsii_.MemberProperty{JsiiProperty: "enableMaintenanceModeInput", GoGetter: "EnableMaintenanceModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomSecurityGroupId", GoMethod: "ResetCustomSecurityGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultOu", GoMethod: "ResetDefaultOu"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableInternetAccess", GoMethod: "ResetEnableInternetAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableMaintenanceMode", GoMethod: "ResetEnableMaintenanceMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserEnabledAsLocalAdministrator", GoMethod: "ResetUserEnabledAsLocalAdministrator"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userEnabledAsLocalAdministrator", GoGetter: "UserEnabledAsLocalAdministrator"},
			_jsii_.MemberProperty{JsiiProperty: "userEnabledAsLocalAdministratorInput", GoGetter: "UserEnabledAsLocalAdministratorInput"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.WorkspacesIpGroup",
		reflect.TypeOf((*WorkspacesIpGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRules", GoMethod: "ResetRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberProperty{JsiiProperty: "rulesInput", GoGetter: "RulesInput"},
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
			j := jsiiProxy_WorkspacesIpGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.WorkspacesIpGroupConfig",
		reflect.TypeOf((*WorkspacesIpGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.WorkspacesIpGroupRules",
		reflect.TypeOf((*WorkspacesIpGroupRules)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.WorkspacesWorkspace",
		reflect.TypeOf((*WorkspacesWorkspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "bundleId", GoGetter: "BundleId"},
			_jsii_.MemberProperty{JsiiProperty: "bundleIdInput", GoGetter: "BundleIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "computerName", GoGetter: "ComputerName"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "directoryId", GoGetter: "DirectoryId"},
			_jsii_.MemberProperty{JsiiProperty: "directoryIdInput", GoGetter: "DirectoryIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddress", GoGetter: "IpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkspaceProperties", GoMethod: "PutWorkspaceProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRootVolumeEncryptionEnabled", GoMethod: "ResetRootVolumeEncryptionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserVolumeEncryptionEnabled", GoMethod: "ResetUserVolumeEncryptionEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeEncryptionKey", GoMethod: "ResetVolumeEncryptionKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkspaceProperties", GoMethod: "ResetWorkspaceProperties"},
			_jsii_.MemberProperty{JsiiProperty: "rootVolumeEncryptionEnabled", GoGetter: "RootVolumeEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "rootVolumeEncryptionEnabledInput", GoGetter: "RootVolumeEncryptionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
			_jsii_.MemberProperty{JsiiProperty: "userNameInput", GoGetter: "UserNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "userVolumeEncryptionEnabled", GoGetter: "UserVolumeEncryptionEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "userVolumeEncryptionEnabledInput", GoGetter: "UserVolumeEncryptionEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "volumeEncryptionKey", GoGetter: "VolumeEncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "volumeEncryptionKeyInput", GoGetter: "VolumeEncryptionKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceProperties", GoGetter: "WorkspaceProperties"},
			_jsii_.MemberProperty{JsiiProperty: "workspacePropertiesInput", GoGetter: "WorkspacePropertiesInput"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesWorkspace{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.WorkspacesWorkspaceConfig",
		reflect.TypeOf((*WorkspacesWorkspaceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.WorkspacesWorkspaceTimeouts",
		reflect.TypeOf((*WorkspacesWorkspaceTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.WorkspacesWorkspaceTimeoutsOutputReference",
		reflect.TypeOf((*WorkspacesWorkspaceTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Workspaces.WorkspacesWorkspaceWorkspaceProperties",
		reflect.TypeOf((*WorkspacesWorkspaceWorkspaceProperties)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Workspaces.WorkspacesWorkspaceWorkspacePropertiesOutputReference",
		reflect.TypeOf((*WorkspacesWorkspaceWorkspacePropertiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "computeTypeName", GoGetter: "ComputeTypeName"},
			_jsii_.MemberProperty{JsiiProperty: "computeTypeNameInput", GoGetter: "ComputeTypeNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeTypeName", GoMethod: "ResetComputeTypeName"},
			_jsii_.MemberMethod{JsiiMethod: "resetRootVolumeSizeGib", GoMethod: "ResetRootVolumeSizeGib"},
			_jsii_.MemberMethod{JsiiMethod: "resetRunningMode", GoMethod: "ResetRunningMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetRunningModeAutoStopTimeoutInMinutes", GoMethod: "ResetRunningModeAutoStopTimeoutInMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserVolumeSizeGib", GoMethod: "ResetUserVolumeSizeGib"},
			_jsii_.MemberProperty{JsiiProperty: "rootVolumeSizeGib", GoGetter: "RootVolumeSizeGib"},
			_jsii_.MemberProperty{JsiiProperty: "rootVolumeSizeGibInput", GoGetter: "RootVolumeSizeGibInput"},
			_jsii_.MemberProperty{JsiiProperty: "runningMode", GoGetter: "RunningMode"},
			_jsii_.MemberProperty{JsiiProperty: "runningModeAutoStopTimeoutInMinutes", GoGetter: "RunningModeAutoStopTimeoutInMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "runningModeAutoStopTimeoutInMinutesInput", GoGetter: "RunningModeAutoStopTimeoutInMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "runningModeInput", GoGetter: "RunningModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userVolumeSizeGib", GoGetter: "UserVolumeSizeGib"},
			_jsii_.MemberProperty{JsiiProperty: "userVolumeSizeGibInput", GoGetter: "UserVolumeSizeGibInput"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
