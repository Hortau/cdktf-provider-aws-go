package sagemaker

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.DataAwsSagemakerPrebuiltEcrImage",
		reflect.TypeOf((*DataAwsSagemakerPrebuiltEcrImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSuffix", GoGetter: "DnsSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSuffixInput", GoGetter: "DnsSuffixInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "imageTag", GoGetter: "ImageTag"},
			_jsii_.MemberProperty{JsiiProperty: "imageTagInput", GoGetter: "ImageTagInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "registryId", GoGetter: "RegistryId"},
			_jsii_.MemberProperty{JsiiProperty: "registryPath", GoGetter: "RegistryPath"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryName", GoGetter: "RepositoryName"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryNameInput", GoGetter: "RepositoryNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsSuffix", GoMethod: "ResetDnsSuffix"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageTag", GoMethod: "ResetImageTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsSagemakerPrebuiltEcrImage{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.DataAwsSagemakerPrebuiltEcrImageConfig",
		reflect.TypeOf((*DataAwsSagemakerPrebuiltEcrImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerApp",
		reflect.TypeOf((*SagemakerApp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appName", GoGetter: "AppName"},
			_jsii_.MemberProperty{JsiiProperty: "appNameInput", GoGetter: "AppNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "appType", GoGetter: "AppType"},
			_jsii_.MemberProperty{JsiiProperty: "appTypeInput", GoGetter: "AppTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "domainIdInput", GoGetter: "DomainIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putResourceSpec", GoMethod: "PutResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceSpec", GoMethod: "ResetResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "resourceSpec", GoGetter: "ResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "resourceSpecInput", GoGetter: "ResourceSpecInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "userProfileName", GoGetter: "UserProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "userProfileNameInput", GoGetter: "UserProfileNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerApp{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerAppConfig",
		reflect.TypeOf((*SagemakerAppConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerAppImageConfig",
		reflect.TypeOf((*SagemakerAppImageConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigName", GoGetter: "AppImageConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "appImageConfigNameInput", GoGetter: "AppImageConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayImageConfig", GoGetter: "KernelGatewayImageConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayImageConfigInput", GoGetter: "KernelGatewayImageConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putKernelGatewayImageConfig", GoMethod: "PutKernelGatewayImageConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetKernelGatewayImageConfig", GoMethod: "ResetKernelGatewayImageConfig"},
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
			j := jsiiProxy_SagemakerAppImageConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerAppImageConfigConfig",
		reflect.TypeOf((*SagemakerAppImageConfigConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerAppImageConfigKernelGatewayImageConfig",
		reflect.TypeOf((*SagemakerAppImageConfigKernelGatewayImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig",
		reflect.TypeOf((*SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference",
		reflect.TypeOf((*SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultGid", GoGetter: "DefaultGid"},
			_jsii_.MemberProperty{JsiiProperty: "defaultGidInput", GoGetter: "DefaultGidInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUid", GoGetter: "DefaultUid"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUidInput", GoGetter: "DefaultUidInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "mountPath", GoGetter: "MountPath"},
			_jsii_.MemberProperty{JsiiProperty: "mountPathInput", GoGetter: "MountPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultGid", GoMethod: "ResetDefaultGid"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultUid", GoMethod: "ResetDefaultUid"},
			_jsii_.MemberMethod{JsiiMethod: "resetMountPath", GoMethod: "ResetMountPath"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerAppImageConfigKernelGatewayImageConfigKernelSpec",
		reflect.TypeOf((*SagemakerAppImageConfigKernelGatewayImageConfigKernelSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerAppImageConfigKernelGatewayImageConfigKernelSpecOutputReference",
		reflect.TypeOf((*SagemakerAppImageConfigKernelGatewayImageConfigKernelSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigKernelSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerAppImageConfigKernelGatewayImageConfigOutputReference",
		reflect.TypeOf((*SagemakerAppImageConfigKernelGatewayImageConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fileSystemConfig", GoGetter: "FileSystemConfig"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemConfigInput", GoGetter: "FileSystemConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kernelSpec", GoGetter: "KernelSpec"},
			_jsii_.MemberProperty{JsiiProperty: "kernelSpecInput", GoGetter: "KernelSpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "putFileSystemConfig", GoMethod: "PutFileSystemConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putKernelSpec", GoMethod: "PutKernelSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemConfig", GoMethod: "ResetFileSystemConfig"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerAppImageConfigKernelGatewayImageConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerAppResourceSpec",
		reflect.TypeOf((*SagemakerAppResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerAppResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerAppResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageArn", GoMethod: "ResetSagemakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageVersionArn", GoMethod: "ResetSagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArn", GoGetter: "SagemakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArnInput", GoGetter: "SagemakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArn", GoGetter: "SagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArnInput", GoGetter: "SagemakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerAppResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerCodeRepository",
		reflect.TypeOf((*SagemakerCodeRepository)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "codeRepositoryName", GoGetter: "CodeRepositoryName"},
			_jsii_.MemberProperty{JsiiProperty: "codeRepositoryNameInput", GoGetter: "CodeRepositoryNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gitConfig", GoGetter: "GitConfig"},
			_jsii_.MemberProperty{JsiiProperty: "gitConfigInput", GoGetter: "GitConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putGitConfig", GoMethod: "PutGitConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
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
			j := jsiiProxy_SagemakerCodeRepository{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerCodeRepositoryConfig",
		reflect.TypeOf((*SagemakerCodeRepositoryConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerCodeRepositoryGitConfig",
		reflect.TypeOf((*SagemakerCodeRepositoryGitConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerCodeRepositoryGitConfigOutputReference",
		reflect.TypeOf((*SagemakerCodeRepositoryGitConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "branch", GoGetter: "Branch"},
			_jsii_.MemberProperty{JsiiProperty: "branchInput", GoGetter: "BranchInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryUrl", GoGetter: "RepositoryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryUrlInput", GoGetter: "RepositoryUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBranch", GoMethod: "ResetBranch"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretArn", GoMethod: "ResetSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretArn", GoGetter: "SecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "secretArnInput", GoGetter: "SecretArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerCodeRepositoryGitConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDevice",
		reflect.TypeOf((*SagemakerDevice)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "agentVersion", GoGetter: "AgentVersion"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "device", GoGetter: "Device"},
			_jsii_.MemberProperty{JsiiProperty: "deviceFleetName", GoGetter: "DeviceFleetName"},
			_jsii_.MemberProperty{JsiiProperty: "deviceFleetNameInput", GoGetter: "DeviceFleetNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceInput", GoGetter: "DeviceInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putDevice", GoMethod: "PutDevice"},
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
			j := jsiiProxy_SagemakerDevice{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDeviceConfig",
		reflect.TypeOf((*SagemakerDeviceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDeviceDevice",
		reflect.TypeOf((*SagemakerDeviceDevice)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDeviceDeviceOutputReference",
		reflect.TypeOf((*SagemakerDeviceDeviceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceName", GoGetter: "DeviceName"},
			_jsii_.MemberProperty{JsiiProperty: "deviceNameInput", GoGetter: "DeviceNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iotThingName", GoGetter: "IotThingName"},
			_jsii_.MemberProperty{JsiiProperty: "iotThingNameInput", GoGetter: "IotThingNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetIotThingName", GoMethod: "ResetIotThingName"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDeviceDeviceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDeviceFleet",
		reflect.TypeOf((*SagemakerDeviceFleet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceFleetName", GoGetter: "DeviceFleetName"},
			_jsii_.MemberProperty{JsiiProperty: "deviceFleetNameInput", GoGetter: "DeviceFleetNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableIotRoleAlias", GoGetter: "EnableIotRoleAlias"},
			_jsii_.MemberProperty{JsiiProperty: "enableIotRoleAliasInput", GoGetter: "EnableIotRoleAliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iotRoleAlias", GoGetter: "IotRoleAlias"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputConfig", GoGetter: "OutputConfig"},
			_jsii_.MemberProperty{JsiiProperty: "outputConfigInput", GoGetter: "OutputConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putOutputConfig", GoMethod: "PutOutputConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableIotRoleAlias", GoMethod: "ResetEnableIotRoleAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
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
			j := jsiiProxy_SagemakerDeviceFleet{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDeviceFleetConfig",
		reflect.TypeOf((*SagemakerDeviceFleetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDeviceFleetOutputConfig",
		reflect.TypeOf((*SagemakerDeviceFleetOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDeviceFleetOutputConfigOutputReference",
		reflect.TypeOf((*SagemakerDeviceFleetOutputConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputLocation", GoGetter: "S3OutputLocation"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputLocationInput", GoGetter: "S3OutputLocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDeviceFleetOutputConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomain",
		reflect.TypeOf((*SagemakerDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appNetworkAccessType", GoGetter: "AppNetworkAccessType"},
			_jsii_.MemberProperty{JsiiProperty: "appNetworkAccessTypeInput", GoGetter: "AppNetworkAccessTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "authMode", GoGetter: "AuthMode"},
			_jsii_.MemberProperty{JsiiProperty: "authModeInput", GoGetter: "AuthModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUserSettings", GoGetter: "DefaultUserSettings"},
			_jsii_.MemberProperty{JsiiProperty: "defaultUserSettingsInput", GoGetter: "DefaultUserSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameInput", GoGetter: "DomainNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "homeEfsFileSystemId", GoGetter: "HomeEfsFileSystemId"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultUserSettings", GoMethod: "PutDefaultUserSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putRetentionPolicy", GoMethod: "PutRetentionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppNetworkAccessType", GoMethod: "ResetAppNetworkAccessType"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetentionPolicy", GoMethod: "ResetRetentionPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "retentionPolicy", GoGetter: "RetentionPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "retentionPolicyInput", GoGetter: "RetentionPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "singleSignOnManagedApplicationInstanceId", GoGetter: "SingleSignOnManagedApplicationInstanceId"},
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
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIdInput", GoGetter: "VpcIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomain{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainConfig",
		reflect.TypeOf((*SagemakerDomainConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsJupyterServerAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterServerAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageArn", GoMethod: "ResetSagemakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageVersionArn", GoMethod: "ResetSagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArn", GoGetter: "SagemakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArnInput", GoGetter: "SagemakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArn", GoGetter: "SagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArnInput", GoGetter: "SagemakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsJupyterServerAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImage",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsCustomImage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageArn", GoMethod: "ResetSagemakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageVersionArn", GoMethod: "ResetSagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArn", GoGetter: "SagemakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArnInput", GoGetter: "SagemakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArn", GoGetter: "SagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArnInput", GoGetter: "SagemakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "customImage", GoGetter: "CustomImage"},
			_jsii_.MemberProperty{JsiiProperty: "customImageInput", GoGetter: "CustomImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomImage", GoMethod: "ResetCustomImage"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsKernelGatewayAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleInput", GoGetter: "ExecutionRoleInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterServerAppSettings", GoGetter: "JupyterServerAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterServerAppSettingsInput", GoGetter: "JupyterServerAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayAppSettings", GoGetter: "KernelGatewayAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayAppSettingsInput", GoGetter: "KernelGatewayAppSettingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putJupyterServerAppSettings", GoMethod: "PutJupyterServerAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putKernelGatewayAppSettings", GoMethod: "PutKernelGatewayAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putSharingSettings", GoMethod: "PutSharingSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putTensorBoardAppSettings", GoMethod: "PutTensorBoardAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetJupyterServerAppSettings", GoMethod: "ResetJupyterServerAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetKernelGatewayAppSettings", GoMethod: "ResetKernelGatewayAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroups", GoMethod: "ResetSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharingSettings", GoMethod: "ResetSharingSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetTensorBoardAppSettings", GoMethod: "ResetTensorBoardAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupsInput", GoGetter: "SecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharingSettings", GoGetter: "SharingSettings"},
			_jsii_.MemberProperty{JsiiProperty: "sharingSettingsInput", GoGetter: "SharingSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tensorBoardAppSettings", GoGetter: "TensorBoardAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "tensorBoardAppSettingsInput", GoGetter: "TensorBoardAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsSharingSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsSharingSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsSharingSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsSharingSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "notebookOutputOption", GoGetter: "NotebookOutputOption"},
			_jsii_.MemberProperty{JsiiProperty: "notebookOutputOptionInput", GoGetter: "NotebookOutputOptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotebookOutputOption", GoMethod: "ResetNotebookOutputOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3KmsKeyId", GoMethod: "ResetS3KmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3OutputPath", GoMethod: "ResetS3OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "s3KmsKeyId", GoGetter: "S3KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "s3KmsKeyIdInput", GoGetter: "S3KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPath", GoGetter: "S3OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPathInput", GoGetter: "S3OutputPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsSharingSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsTensorBoardAppSettings",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsTensorBoardAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageArn", GoMethod: "ResetSagemakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageVersionArn", GoMethod: "ResetSagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArn", GoGetter: "SagemakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArnInput", GoGetter: "SagemakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArn", GoGetter: "SagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArnInput", GoGetter: "SagemakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultResourceSpec", GoMethod: "ResetDefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainDefaultUserSettingsTensorBoardAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerDomainRetentionPolicy",
		reflect.TypeOf((*SagemakerDomainRetentionPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerDomainRetentionPolicyOutputReference",
		reflect.TypeOf((*SagemakerDomainRetentionPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "homeEfsFileSystem", GoGetter: "HomeEfsFileSystem"},
			_jsii_.MemberProperty{JsiiProperty: "homeEfsFileSystemInput", GoGetter: "HomeEfsFileSystemInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetHomeEfsFileSystem", GoMethod: "ResetHomeEfsFileSystem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerDomainRetentionPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpoint",
		reflect.TypeOf((*SagemakerEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfig", GoGetter: "DeploymentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentConfigInput", GoGetter: "DeploymentConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfigName", GoGetter: "EndpointConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfigNameInput", GoGetter: "EndpointConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putDeploymentConfig", GoMethod: "PutDeploymentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeploymentConfig", GoMethod: "ResetDeploymentConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
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
			j := jsiiProxy_SagemakerEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfig",
		reflect.TypeOf((*SagemakerEndpointConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfiguration",
		reflect.TypeOf((*SagemakerEndpointConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "asyncInferenceConfig", GoGetter: "AsyncInferenceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "asyncInferenceConfigInput", GoGetter: "AsyncInferenceConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dataCaptureConfig", GoGetter: "DataCaptureConfig"},
			_jsii_.MemberProperty{JsiiProperty: "dataCaptureConfigInput", GoGetter: "DataCaptureConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyArn", GoGetter: "KmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyArnInput", GoGetter: "KmsKeyArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "productionVariants", GoGetter: "ProductionVariants"},
			_jsii_.MemberProperty{JsiiProperty: "productionVariantsInput", GoGetter: "ProductionVariantsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putAsyncInferenceConfig", GoMethod: "PutAsyncInferenceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putDataCaptureConfig", GoMethod: "PutDataCaptureConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAsyncInferenceConfig", GoMethod: "ResetAsyncInferenceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataCaptureConfig", GoMethod: "ResetDataCaptureConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyArn", GoMethod: "ResetKmsKeyArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
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
			j := jsiiProxy_SagemakerEndpointConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfig",
		reflect.TypeOf((*SagemakerEndpointConfigurationAsyncInferenceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig",
		reflect.TypeOf((*SagemakerEndpointConfigurationAsyncInferenceConfigClientConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfigClientConfigOutputReference",
		reflect.TypeOf((*SagemakerEndpointConfigurationAsyncInferenceConfigClientConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentInvocationsPerInstance", GoGetter: "MaxConcurrentInvocationsPerInstance"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentInvocationsPerInstanceInput", GoGetter: "MaxConcurrentInvocationsPerInstanceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrentInvocationsPerInstance", GoMethod: "ResetMaxConcurrentInvocationsPerInstance"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigClientConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig",
		reflect.TypeOf((*SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig",
		reflect.TypeOf((*SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference",
		reflect.TypeOf((*SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "errorTopic", GoGetter: "ErrorTopic"},
			_jsii_.MemberProperty{JsiiProperty: "errorTopicInput", GoGetter: "ErrorTopicInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorTopic", GoMethod: "ResetErrorTopic"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuccessTopic", GoMethod: "ResetSuccessTopic"},
			_jsii_.MemberProperty{JsiiProperty: "successTopic", GoGetter: "SuccessTopic"},
			_jsii_.MemberProperty{JsiiProperty: "successTopicInput", GoGetter: "SuccessTopicInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference",
		reflect.TypeOf((*SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfig", GoGetter: "NotificationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfigInput", GoGetter: "NotificationConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putNotificationConfig", GoMethod: "PutNotificationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationConfig", GoMethod: "ResetNotificationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPath", GoGetter: "S3OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPathInput", GoGetter: "S3OutputPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference",
		reflect.TypeOf((*SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientConfig", GoGetter: "ClientConfig"},
			_jsii_.MemberProperty{JsiiProperty: "clientConfigInput", GoGetter: "ClientConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "outputConfig", GoGetter: "OutputConfig"},
			_jsii_.MemberProperty{JsiiProperty: "outputConfigInput", GoGetter: "OutputConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putClientConfig", GoMethod: "PutClientConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putOutputConfig", GoMethod: "PutOutputConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientConfig", GoMethod: "ResetClientConfig"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointConfigurationAsyncInferenceConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationConfig",
		reflect.TypeOf((*SagemakerEndpointConfigurationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationDataCaptureConfig",
		reflect.TypeOf((*SagemakerEndpointConfigurationDataCaptureConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationDataCaptureConfigCaptureContentTypeHeader",
		reflect.TypeOf((*SagemakerEndpointConfigurationDataCaptureConfigCaptureContentTypeHeader)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationDataCaptureConfigCaptureContentTypeHeaderOutputReference",
		reflect.TypeOf((*SagemakerEndpointConfigurationDataCaptureConfigCaptureContentTypeHeaderOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "csvContentTypes", GoGetter: "CsvContentTypes"},
			_jsii_.MemberProperty{JsiiProperty: "csvContentTypesInput", GoGetter: "CsvContentTypesInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "jsonContentTypes", GoGetter: "JsonContentTypes"},
			_jsii_.MemberProperty{JsiiProperty: "jsonContentTypesInput", GoGetter: "JsonContentTypesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCsvContentTypes", GoMethod: "ResetCsvContentTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetJsonContentTypes", GoMethod: "ResetJsonContentTypes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointConfigurationDataCaptureConfigCaptureContentTypeHeaderOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationDataCaptureConfigCaptureOptions",
		reflect.TypeOf((*SagemakerEndpointConfigurationDataCaptureConfigCaptureOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationDataCaptureConfigOutputReference",
		reflect.TypeOf((*SagemakerEndpointConfigurationDataCaptureConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "captureContentTypeHeader", GoGetter: "CaptureContentTypeHeader"},
			_jsii_.MemberProperty{JsiiProperty: "captureContentTypeHeaderInput", GoGetter: "CaptureContentTypeHeaderInput"},
			_jsii_.MemberProperty{JsiiProperty: "captureOptions", GoGetter: "CaptureOptions"},
			_jsii_.MemberProperty{JsiiProperty: "captureOptionsInput", GoGetter: "CaptureOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "destinationS3Uri", GoGetter: "DestinationS3Uri"},
			_jsii_.MemberProperty{JsiiProperty: "destinationS3UriInput", GoGetter: "DestinationS3UriInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableCapture", GoGetter: "EnableCapture"},
			_jsii_.MemberProperty{JsiiProperty: "enableCaptureInput", GoGetter: "EnableCaptureInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "initialSamplingPercentage", GoGetter: "InitialSamplingPercentage"},
			_jsii_.MemberProperty{JsiiProperty: "initialSamplingPercentageInput", GoGetter: "InitialSamplingPercentageInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCaptureContentTypeHeader", GoMethod: "PutCaptureContentTypeHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaptureContentTypeHeader", GoMethod: "ResetCaptureContentTypeHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableCapture", GoMethod: "ResetEnableCapture"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointConfigurationDataCaptureConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointConfigurationProductionVariants",
		reflect.TypeOf((*SagemakerEndpointConfigurationProductionVariants)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfig",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigAutoRollbackConfiguration",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigAutoRollbackConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigAutoRollbackConfigurationAlarms",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigAutoRollbackConfigurationAlarms)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigAutoRollbackConfigurationOutputReference",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigAutoRollbackConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alarms", GoGetter: "Alarms"},
			_jsii_.MemberProperty{JsiiProperty: "alarmsInput", GoGetter: "AlarmsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlarms", GoMethod: "ResetAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointDeploymentConfigAutoRollbackConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicy",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyOutputReference",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maximumExecutionTimeoutInSeconds", GoGetter: "MaximumExecutionTimeoutInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "maximumExecutionTimeoutInSecondsInput", GoGetter: "MaximumExecutionTimeoutInSecondsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putTrafficRoutingConfiguration", GoMethod: "PutTrafficRoutingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumExecutionTimeoutInSeconds", GoMethod: "ResetMaximumExecutionTimeoutInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTerminationWaitInSeconds", GoMethod: "ResetTerminationWaitInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "terminationWaitInSeconds", GoGetter: "TerminationWaitInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "terminationWaitInSecondsInput", GoGetter: "TerminationWaitInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "trafficRoutingConfiguration", GoGetter: "TrafficRoutingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "trafficRoutingConfigurationInput", GoGetter: "TrafficRoutingConfigurationInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfiguration",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationCanarySize",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationCanarySize)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationCanarySizeOutputReference",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationCanarySizeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationCanarySizeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationLinearStepSize",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationLinearStepSize)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationLinearStepSizeOutputReference",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationLinearStepSizeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationLinearStepSizeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationOutputReference",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "canarySize", GoGetter: "CanarySize"},
			_jsii_.MemberProperty{JsiiProperty: "canarySizeInput", GoGetter: "CanarySizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "linearStepSize", GoGetter: "LinearStepSize"},
			_jsii_.MemberProperty{JsiiProperty: "linearStepSizeInput", GoGetter: "LinearStepSizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCanarySize", GoMethod: "PutCanarySize"},
			_jsii_.MemberMethod{JsiiMethod: "putLinearStepSize", GoMethod: "PutLinearStepSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetCanarySize", GoMethod: "ResetCanarySize"},
			_jsii_.MemberMethod{JsiiMethod: "resetLinearStepSize", GoMethod: "ResetLinearStepSize"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "waitIntervalInSeconds", GoGetter: "WaitIntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "waitIntervalInSecondsInput", GoGetter: "WaitIntervalInSecondsInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerEndpointDeploymentConfigOutputReference",
		reflect.TypeOf((*SagemakerEndpointDeploymentConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoRollbackConfiguration", GoGetter: "AutoRollbackConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "autoRollbackConfigurationInput", GoGetter: "AutoRollbackConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "blueGreenUpdatePolicy", GoGetter: "BlueGreenUpdatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "blueGreenUpdatePolicyInput", GoGetter: "BlueGreenUpdatePolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAutoRollbackConfiguration", GoMethod: "PutAutoRollbackConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putBlueGreenUpdatePolicy", GoMethod: "PutBlueGreenUpdatePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRollbackConfiguration", GoMethod: "ResetAutoRollbackConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerEndpointDeploymentConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroup",
		reflect.TypeOf((*SagemakerFeatureGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventTimeFeatureName", GoGetter: "EventTimeFeatureName"},
			_jsii_.MemberProperty{JsiiProperty: "eventTimeFeatureNameInput", GoGetter: "EventTimeFeatureNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "featureDefinition", GoGetter: "FeatureDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "featureDefinitionInput", GoGetter: "FeatureDefinitionInput"},
			_jsii_.MemberProperty{JsiiProperty: "featureGroupName", GoGetter: "FeatureGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "featureGroupNameInput", GoGetter: "FeatureGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "offlineStoreConfig", GoGetter: "OfflineStoreConfig"},
			_jsii_.MemberProperty{JsiiProperty: "offlineStoreConfigInput", GoGetter: "OfflineStoreConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "onlineStoreConfig", GoGetter: "OnlineStoreConfig"},
			_jsii_.MemberProperty{JsiiProperty: "onlineStoreConfigInput", GoGetter: "OnlineStoreConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putOfflineStoreConfig", GoMethod: "PutOfflineStoreConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putOnlineStoreConfig", GoMethod: "PutOnlineStoreConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recordIdentifierFeatureName", GoGetter: "RecordIdentifierFeatureName"},
			_jsii_.MemberProperty{JsiiProperty: "recordIdentifierFeatureNameInput", GoGetter: "RecordIdentifierFeatureNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetOfflineStoreConfig", GoMethod: "ResetOfflineStoreConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnlineStoreConfig", GoMethod: "ResetOnlineStoreConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
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
			j := jsiiProxy_SagemakerFeatureGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupConfig",
		reflect.TypeOf((*SagemakerFeatureGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupFeatureDefinition",
		reflect.TypeOf((*SagemakerFeatureGroupFeatureDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOfflineStoreConfig",
		reflect.TypeOf((*SagemakerFeatureGroupOfflineStoreConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig",
		reflect.TypeOf((*SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfigOutputReference",
		reflect.TypeOf((*SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "catalog", GoGetter: "Catalog"},
			_jsii_.MemberProperty{JsiiProperty: "catalogInput", GoGetter: "CatalogInput"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetCatalog", GoMethod: "ResetCatalog"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatabase", GoMethod: "ResetDatabase"},
			_jsii_.MemberMethod{JsiiMethod: "resetTableName", GoMethod: "ResetTableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "tableNameInput", GoGetter: "TableNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOfflineStoreConfigOutputReference",
		reflect.TypeOf((*SagemakerFeatureGroupOfflineStoreConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dataCatalogConfig", GoGetter: "DataCatalogConfig"},
			_jsii_.MemberProperty{JsiiProperty: "dataCatalogConfigInput", GoGetter: "DataCatalogConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableGlueTableCreation", GoGetter: "DisableGlueTableCreation"},
			_jsii_.MemberProperty{JsiiProperty: "disableGlueTableCreationInput", GoGetter: "DisableGlueTableCreationInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putDataCatalogConfig", GoMethod: "PutDataCatalogConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putS3StorageConfig", GoMethod: "PutS3StorageConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataCatalogConfig", GoMethod: "ResetDataCatalogConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableGlueTableCreation", GoMethod: "ResetDisableGlueTableCreation"},
			_jsii_.MemberProperty{JsiiProperty: "s3StorageConfig", GoGetter: "S3StorageConfig"},
			_jsii_.MemberProperty{JsiiProperty: "s3StorageConfigInput", GoGetter: "S3StorageConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig",
		reflect.TypeOf((*SagemakerFeatureGroupOfflineStoreConfigS3StorageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOfflineStoreConfigS3StorageConfigOutputReference",
		reflect.TypeOf((*SagemakerFeatureGroupOfflineStoreConfigS3StorageConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "s3Uri", GoGetter: "S3Uri"},
			_jsii_.MemberProperty{JsiiProperty: "s3UriInput", GoGetter: "S3UriInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFeatureGroupOfflineStoreConfigS3StorageConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOnlineStoreConfig",
		reflect.TypeOf((*SagemakerFeatureGroupOnlineStoreConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOnlineStoreConfigOutputReference",
		reflect.TypeOf((*SagemakerFeatureGroupOnlineStoreConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enableOnlineStore", GoGetter: "EnableOnlineStore"},
			_jsii_.MemberProperty{JsiiProperty: "enableOnlineStoreInput", GoGetter: "EnableOnlineStoreInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putSecurityConfig", GoMethod: "PutSecurityConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableOnlineStore", GoMethod: "ResetEnableOnlineStore"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityConfig", GoMethod: "ResetSecurityConfig"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfig", GoGetter: "SecurityConfig"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfigInput", GoGetter: "SecurityConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOnlineStoreConfigSecurityConfig",
		reflect.TypeOf((*SagemakerFeatureGroupOnlineStoreConfigSecurityConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFeatureGroupOnlineStoreConfigSecurityConfigOutputReference",
		reflect.TypeOf((*SagemakerFeatureGroupOnlineStoreConfigSecurityConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFeatureGroupOnlineStoreConfigSecurityConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinition",
		reflect.TypeOf((*SagemakerFlowDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "flowDefinitionName", GoGetter: "FlowDefinitionName"},
			_jsii_.MemberProperty{JsiiProperty: "flowDefinitionNameInput", GoGetter: "FlowDefinitionNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopActivationConfig", GoGetter: "HumanLoopActivationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopActivationConfigInput", GoGetter: "HumanLoopActivationConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopConfig", GoGetter: "HumanLoopConfig"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopConfigInput", GoGetter: "HumanLoopConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopRequestSource", GoGetter: "HumanLoopRequestSource"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopRequestSourceInput", GoGetter: "HumanLoopRequestSourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputConfig", GoGetter: "OutputConfig"},
			_jsii_.MemberProperty{JsiiProperty: "outputConfigInput", GoGetter: "OutputConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putHumanLoopActivationConfig", GoMethod: "PutHumanLoopActivationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putHumanLoopConfig", GoMethod: "PutHumanLoopConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putHumanLoopRequestSource", GoMethod: "PutHumanLoopRequestSource"},
			_jsii_.MemberMethod{JsiiMethod: "putOutputConfig", GoMethod: "PutOutputConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetHumanLoopActivationConfig", GoMethod: "ResetHumanLoopActivationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetHumanLoopRequestSource", GoMethod: "ResetHumanLoopRequestSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
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
			j := jsiiProxy_SagemakerFlowDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionConfig",
		reflect.TypeOf((*SagemakerFlowDefinitionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopActivationConfig",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopActivationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfig",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfigOutputReference",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopActivationConditions", GoGetter: "HumanLoopActivationConditions"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopActivationConditionsInput", GoGetter: "HumanLoopActivationConditionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopActivationConfigOutputReference",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopActivationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopActivationConditionsConfig", GoGetter: "HumanLoopActivationConditionsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "humanLoopActivationConditionsConfigInput", GoGetter: "HumanLoopActivationConditionsConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putHumanLoopActivationConditionsConfig", GoMethod: "PutHumanLoopActivationConditionsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetHumanLoopActivationConditionsConfig", GoMethod: "ResetHumanLoopActivationConditionsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFlowDefinitionHumanLoopActivationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopConfig",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopConfigOutputReference",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "humanTaskUiArn", GoGetter: "HumanTaskUiArn"},
			_jsii_.MemberProperty{JsiiProperty: "humanTaskUiArnInput", GoGetter: "HumanTaskUiArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "publicWorkforceTaskPrice", GoGetter: "PublicWorkforceTaskPrice"},
			_jsii_.MemberProperty{JsiiProperty: "publicWorkforceTaskPriceInput", GoGetter: "PublicWorkforceTaskPriceInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPublicWorkforceTaskPrice", GoMethod: "PutPublicWorkforceTaskPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicWorkforceTaskPrice", GoMethod: "ResetPublicWorkforceTaskPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetTaskAvailabilityLifetimeInSeconds", GoMethod: "ResetTaskAvailabilityLifetimeInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTaskKeywords", GoMethod: "ResetTaskKeywords"},
			_jsii_.MemberMethod{JsiiMethod: "resetTaskTimeLimitInSeconds", GoMethod: "ResetTaskTimeLimitInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "taskAvailabilityLifetimeInSeconds", GoGetter: "TaskAvailabilityLifetimeInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "taskAvailabilityLifetimeInSecondsInput", GoGetter: "TaskAvailabilityLifetimeInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskCount", GoGetter: "TaskCount"},
			_jsii_.MemberProperty{JsiiProperty: "taskCountInput", GoGetter: "TaskCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskDescription", GoGetter: "TaskDescription"},
			_jsii_.MemberProperty{JsiiProperty: "taskDescriptionInput", GoGetter: "TaskDescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskKeywords", GoGetter: "TaskKeywords"},
			_jsii_.MemberProperty{JsiiProperty: "taskKeywordsInput", GoGetter: "TaskKeywordsInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskTimeLimitInSeconds", GoGetter: "TaskTimeLimitInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "taskTimeLimitInSecondsInput", GoGetter: "TaskTimeLimitInSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskTitle", GoGetter: "TaskTitle"},
			_jsii_.MemberProperty{JsiiProperty: "taskTitleInput", GoGetter: "TaskTitleInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "workteamArn", GoGetter: "WorkteamArn"},
			_jsii_.MemberProperty{JsiiProperty: "workteamArnInput", GoGetter: "WorkteamArnInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFlowDefinitionHumanLoopConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPrice",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPrice)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsd",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsd)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsdOutputReference",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsdOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cents", GoGetter: "Cents"},
			_jsii_.MemberProperty{JsiiProperty: "centsInput", GoGetter: "CentsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dollars", GoGetter: "Dollars"},
			_jsii_.MemberProperty{JsiiProperty: "dollarsInput", GoGetter: "DollarsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetCents", GoMethod: "ResetCents"},
			_jsii_.MemberMethod{JsiiMethod: "resetDollars", GoMethod: "ResetDollars"},
			_jsii_.MemberMethod{JsiiMethod: "resetTenthFractionsOfACent", GoMethod: "ResetTenthFractionsOfACent"},
			_jsii_.MemberProperty{JsiiProperty: "tenthFractionsOfACent", GoGetter: "TenthFractionsOfACent"},
			_jsii_.MemberProperty{JsiiProperty: "tenthFractionsOfACentInput", GoGetter: "TenthFractionsOfACentInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsdOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceOutputReference",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amountInUsd", GoGetter: "AmountInUsd"},
			_jsii_.MemberProperty{JsiiProperty: "amountInUsdInput", GoGetter: "AmountInUsdInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAmountInUsd", GoMethod: "PutAmountInUsd"},
			_jsii_.MemberMethod{JsiiMethod: "resetAmountInUsd", GoMethod: "ResetAmountInUsd"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopRequestSource",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopRequestSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionHumanLoopRequestSourceOutputReference",
		reflect.TypeOf((*SagemakerFlowDefinitionHumanLoopRequestSourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsManagedHumanLoopRequestSource", GoGetter: "AwsManagedHumanLoopRequestSource"},
			_jsii_.MemberProperty{JsiiProperty: "awsManagedHumanLoopRequestSourceInput", GoGetter: "AwsManagedHumanLoopRequestSourceInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFlowDefinitionHumanLoopRequestSourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionOutputConfig",
		reflect.TypeOf((*SagemakerFlowDefinitionOutputConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerFlowDefinitionOutputConfigOutputReference",
		reflect.TypeOf((*SagemakerFlowDefinitionOutputConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPath", GoGetter: "S3OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPathInput", GoGetter: "S3OutputPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerFlowDefinitionOutputConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerHumanTaskUi",
		reflect.TypeOf((*SagemakerHumanTaskUi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "humanTaskUiName", GoGetter: "HumanTaskUiName"},
			_jsii_.MemberProperty{JsiiProperty: "humanTaskUiNameInput", GoGetter: "HumanTaskUiNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putUiTemplate", GoMethod: "PutUiTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
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
			_jsii_.MemberProperty{JsiiProperty: "uiTemplate", GoGetter: "UiTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "uiTemplateInput", GoGetter: "UiTemplateInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerHumanTaskUi{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerHumanTaskUiConfig",
		reflect.TypeOf((*SagemakerHumanTaskUiConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerHumanTaskUiUiTemplate",
		reflect.TypeOf((*SagemakerHumanTaskUiUiTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerHumanTaskUiUiTemplateOutputReference",
		reflect.TypeOf((*SagemakerHumanTaskUiUiTemplateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "contentInput", GoGetter: "ContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentSha256", GoGetter: "ContentSha256"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetContent", GoMethod: "ResetContent"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerHumanTaskUiUiTemplateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerImage",
		reflect.TypeOf((*SagemakerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageNameInput", GoGetter: "ImageNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
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
			j := jsiiProxy_SagemakerImage{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerImageConfig",
		reflect.TypeOf((*SagemakerImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerImageVersion",
		reflect.TypeOf((*SagemakerImageVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "baseImage", GoGetter: "BaseImage"},
			_jsii_.MemberProperty{JsiiProperty: "baseImageInput", GoGetter: "BaseImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "containerImage", GoGetter: "ContainerImage"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "imageArn", GoGetter: "ImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "imageNameInput", GoGetter: "ImageNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerImageVersion{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerImageVersionConfig",
		reflect.TypeOf((*SagemakerImageVersionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerModel",
		reflect.TypeOf((*SagemakerModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerInput", GoGetter: "ContainerInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableNetworkIsolation", GoGetter: "EnableNetworkIsolation"},
			_jsii_.MemberProperty{JsiiProperty: "enableNetworkIsolationInput", GoGetter: "EnableNetworkIsolationInput"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArnInput", GoGetter: "ExecutionRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceExecutionConfig", GoGetter: "InferenceExecutionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceExecutionConfigInput", GoGetter: "InferenceExecutionConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "primaryContainer", GoGetter: "PrimaryContainer"},
			_jsii_.MemberProperty{JsiiProperty: "primaryContainerInput", GoGetter: "PrimaryContainerInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putInferenceExecutionConfig", GoMethod: "PutInferenceExecutionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putPrimaryContainer", GoMethod: "PutPrimaryContainer"},
			_jsii_.MemberMethod{JsiiMethod: "putVpcConfig", GoMethod: "PutVpcConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetContainer", GoMethod: "ResetContainer"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableNetworkIsolation", GoMethod: "ResetEnableNetworkIsolation"},
			_jsii_.MemberMethod{JsiiMethod: "resetInferenceExecutionConfig", GoMethod: "ResetInferenceExecutionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryContainer", GoMethod: "ResetPrimaryContainer"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcConfig", GoMethod: "ResetVpcConfig"},
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
			_jsii_.MemberProperty{JsiiProperty: "vpcConfig", GoGetter: "VpcConfig"},
			_jsii_.MemberProperty{JsiiProperty: "vpcConfigInput", GoGetter: "VpcConfigInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerModel{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerModelConfig",
		reflect.TypeOf((*SagemakerModelConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerModelContainer",
		reflect.TypeOf((*SagemakerModelContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerModelContainerImageConfig",
		reflect.TypeOf((*SagemakerModelContainerImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerModelContainerImageConfigOutputReference",
		reflect.TypeOf((*SagemakerModelContainerImageConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryAccessMode", GoGetter: "RepositoryAccessMode"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryAccessModeInput", GoGetter: "RepositoryAccessModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerModelContainerImageConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerModelInferenceExecutionConfig",
		reflect.TypeOf((*SagemakerModelInferenceExecutionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerModelInferenceExecutionConfigOutputReference",
		reflect.TypeOf((*SagemakerModelInferenceExecutionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "modeInput", GoGetter: "ModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerModelInferenceExecutionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerModelPackageGroup",
		reflect.TypeOf((*SagemakerModelPackageGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "modelPackageGroupDescription", GoGetter: "ModelPackageGroupDescription"},
			_jsii_.MemberProperty{JsiiProperty: "modelPackageGroupDescriptionInput", GoGetter: "ModelPackageGroupDescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "modelPackageGroupName", GoGetter: "ModelPackageGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "modelPackageGroupNameInput", GoGetter: "ModelPackageGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelPackageGroupDescription", GoMethod: "ResetModelPackageGroupDescription"},
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
			j := jsiiProxy_SagemakerModelPackageGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerModelPackageGroupConfig",
		reflect.TypeOf((*SagemakerModelPackageGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerModelPackageGroupPolicy",
		reflect.TypeOf((*SagemakerModelPackageGroupPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "modelPackageGroupName", GoGetter: "ModelPackageGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "modelPackageGroupNameInput", GoGetter: "ModelPackageGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePolicy", GoGetter: "ResourcePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePolicyInput", GoGetter: "ResourcePolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerModelPackageGroupPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerModelPackageGroupPolicyConfig",
		reflect.TypeOf((*SagemakerModelPackageGroupPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerModelPrimaryContainer",
		reflect.TypeOf((*SagemakerModelPrimaryContainer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerModelPrimaryContainerImageConfig",
		reflect.TypeOf((*SagemakerModelPrimaryContainerImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerModelPrimaryContainerImageConfigOutputReference",
		reflect.TypeOf((*SagemakerModelPrimaryContainerImageConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryAccessMode", GoGetter: "RepositoryAccessMode"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryAccessModeInput", GoGetter: "RepositoryAccessModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerModelPrimaryContainerImageConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerModelPrimaryContainerOutputReference",
		reflect.TypeOf((*SagemakerModelPrimaryContainerOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "containerHostname", GoGetter: "ContainerHostname"},
			_jsii_.MemberProperty{JsiiProperty: "containerHostnameInput", GoGetter: "ContainerHostnameInput"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "environmentInput", GoGetter: "EnvironmentInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imageConfig", GoGetter: "ImageConfig"},
			_jsii_.MemberProperty{JsiiProperty: "imageConfigInput", GoGetter: "ImageConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageInput", GoGetter: "ImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "modeInput", GoGetter: "ModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "modelDataUrl", GoGetter: "ModelDataUrl"},
			_jsii_.MemberProperty{JsiiProperty: "modelDataUrlInput", GoGetter: "ModelDataUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "putImageConfig", GoMethod: "PutImageConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetContainerHostname", GoMethod: "ResetContainerHostname"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironment", GoMethod: "ResetEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageConfig", GoMethod: "ResetImageConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMode", GoMethod: "ResetMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelDataUrl", GoMethod: "ResetModelDataUrl"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerModelPrimaryContainerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerModelVpcConfig",
		reflect.TypeOf((*SagemakerModelVpcConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerModelVpcConfigOutputReference",
		reflect.TypeOf((*SagemakerModelVpcConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberProperty{JsiiProperty: "subnetsInput", GoGetter: "SubnetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerModelVpcConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerNotebookInstance",
		reflect.TypeOf((*SagemakerNotebookInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalCodeRepositories", GoGetter: "AdditionalCodeRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "additionalCodeRepositoriesInput", GoGetter: "AdditionalCodeRepositoriesInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCodeRepository", GoGetter: "DefaultCodeRepository"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCodeRepositoryInput", GoGetter: "DefaultCodeRepositoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "directInternetAccess", GoGetter: "DirectInternetAccess"},
			_jsii_.MemberProperty{JsiiProperty: "directInternetAccessInput", GoGetter: "DirectInternetAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigName", GoGetter: "LifecycleConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigNameInput", GoGetter: "LifecycleConfigNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "platformIdentifier", GoGetter: "PlatformIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "platformIdentifierInput", GoGetter: "PlatformIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalCodeRepositories", GoMethod: "ResetAdditionalCodeRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultCodeRepository", GoMethod: "ResetDefaultCodeRepository"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirectInternetAccess", GoMethod: "ResetDirectInternetAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigName", GoMethod: "ResetLifecycleConfigName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlatformIdentifier", GoMethod: "ResetPlatformIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetRootAccess", GoMethod: "ResetRootAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroups", GoMethod: "ResetSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetId", GoMethod: "ResetSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeSize", GoMethod: "ResetVolumeSize"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "rootAccess", GoGetter: "RootAccess"},
			_jsii_.MemberProperty{JsiiProperty: "rootAccessInput", GoGetter: "RootAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupsInput", GoGetter: "SecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdInput", GoGetter: "SubnetIdInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSize", GoGetter: "VolumeSize"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSizeInput", GoGetter: "VolumeSizeInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerNotebookInstance{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerNotebookInstanceConfig",
		reflect.TypeOf((*SagemakerNotebookInstanceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerNotebookInstanceLifecycleConfiguration",
		reflect.TypeOf((*SagemakerNotebookInstanceLifecycleConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "onCreate", GoGetter: "OnCreate"},
			_jsii_.MemberProperty{JsiiProperty: "onCreateInput", GoGetter: "OnCreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "onStart", GoGetter: "OnStart"},
			_jsii_.MemberProperty{JsiiProperty: "onStartInput", GoGetter: "OnStartInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnCreate", GoMethod: "ResetOnCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnStart", GoMethod: "ResetOnStart"},
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
			j := jsiiProxy_SagemakerNotebookInstanceLifecycleConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerNotebookInstanceLifecycleConfigurationConfig",
		reflect.TypeOf((*SagemakerNotebookInstanceLifecycleConfigurationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerProject",
		reflect.TypeOf((*SagemakerProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectDescription", GoGetter: "ProjectDescription"},
			_jsii_.MemberProperty{JsiiProperty: "projectDescriptionInput", GoGetter: "ProjectDescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectName", GoGetter: "ProjectName"},
			_jsii_.MemberProperty{JsiiProperty: "projectNameInput", GoGetter: "ProjectNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putServiceCatalogProvisioningDetails", GoMethod: "PutServiceCatalogProvisioningDetails"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectDescription", GoMethod: "ResetProjectDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "serviceCatalogProvisioningDetails", GoGetter: "ServiceCatalogProvisioningDetails"},
			_jsii_.MemberProperty{JsiiProperty: "serviceCatalogProvisioningDetailsInput", GoGetter: "ServiceCatalogProvisioningDetailsInput"},
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
			j := jsiiProxy_SagemakerProject{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerProjectConfig",
		reflect.TypeOf((*SagemakerProjectConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerProjectServiceCatalogProvisioningDetails",
		reflect.TypeOf((*SagemakerProjectServiceCatalogProvisioningDetails)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerProjectServiceCatalogProvisioningDetailsOutputReference",
		reflect.TypeOf((*SagemakerProjectServiceCatalogProvisioningDetailsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "pathId", GoGetter: "PathId"},
			_jsii_.MemberProperty{JsiiProperty: "pathIdInput", GoGetter: "PathIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "productId", GoGetter: "ProductId"},
			_jsii_.MemberProperty{JsiiProperty: "productIdInput", GoGetter: "ProductIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningArtifactId", GoGetter: "ProvisioningArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningArtifactIdInput", GoGetter: "ProvisioningArtifactIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningParameter", GoGetter: "ProvisioningParameter"},
			_jsii_.MemberProperty{JsiiProperty: "provisioningParameterInput", GoGetter: "ProvisioningParameterInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathId", GoMethod: "ResetPathId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvisioningArtifactId", GoMethod: "ResetProvisioningArtifactId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvisioningParameter", GoMethod: "ResetProvisioningParameter"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParameter",
		reflect.TypeOf((*SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParameter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerStudioLifecycleConfig",
		reflect.TypeOf((*SagemakerStudioLifecycleConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "studioLifecycleConfigAppType", GoGetter: "StudioLifecycleConfigAppType"},
			_jsii_.MemberProperty{JsiiProperty: "studioLifecycleConfigAppTypeInput", GoGetter: "StudioLifecycleConfigAppTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "studioLifecycleConfigContent", GoGetter: "StudioLifecycleConfigContent"},
			_jsii_.MemberProperty{JsiiProperty: "studioLifecycleConfigContentInput", GoGetter: "StudioLifecycleConfigContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "studioLifecycleConfigName", GoGetter: "StudioLifecycleConfigName"},
			_jsii_.MemberProperty{JsiiProperty: "studioLifecycleConfigNameInput", GoGetter: "StudioLifecycleConfigNameInput"},
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
			j := jsiiProxy_SagemakerStudioLifecycleConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerStudioLifecycleConfigConfig",
		reflect.TypeOf((*SagemakerStudioLifecycleConfigConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerUserProfile",
		reflect.TypeOf((*SagemakerUserProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "domainIdInput", GoGetter: "DomainIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "homeEfsFileSystemUid", GoGetter: "HomeEfsFileSystemUid"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putUserSettings", GoMethod: "PutUserSettings"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSingleSignOnUserIdentifier", GoMethod: "ResetSingleSignOnUserIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetSingleSignOnUserValue", GoMethod: "ResetSingleSignOnUserValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserSettings", GoMethod: "ResetUserSettings"},
			_jsii_.MemberProperty{JsiiProperty: "singleSignOnUserIdentifier", GoGetter: "SingleSignOnUserIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "singleSignOnUserIdentifierInput", GoGetter: "SingleSignOnUserIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "singleSignOnUserValue", GoGetter: "SingleSignOnUserValue"},
			_jsii_.MemberProperty{JsiiProperty: "singleSignOnUserValueInput", GoGetter: "SingleSignOnUserValueInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "userProfileName", GoGetter: "UserProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "userProfileNameInput", GoGetter: "UserProfileNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "userSettings", GoGetter: "UserSettings"},
			_jsii_.MemberProperty{JsiiProperty: "userSettingsInput", GoGetter: "UserSettingsInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerUserProfile{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileConfig",
		reflect.TypeOf((*SagemakerUserProfileConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettings",
		reflect.TypeOf((*SagemakerUserProfileUserSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsJupyterServerAppSettings",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsJupyterServerAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsJupyterServerAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsJupyterServerAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageArn", GoMethod: "ResetSagemakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageVersionArn", GoMethod: "ResetSagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArn", GoGetter: "SagemakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArnInput", GoGetter: "SagemakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArn", GoGetter: "SagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArnInput", GoGetter: "SagemakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerUserProfileUserSettingsJupyterServerAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsJupyterServerAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsJupyterServerAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerUserProfileUserSettingsJupyterServerAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsKernelGatewayAppSettings",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsKernelGatewayAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsKernelGatewayAppSettingsCustomImage",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsKernelGatewayAppSettingsCustomImage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsKernelGatewayAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsKernelGatewayAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageArn", GoMethod: "ResetSagemakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageVersionArn", GoMethod: "ResetSagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArn", GoGetter: "SagemakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArnInput", GoGetter: "SagemakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArn", GoGetter: "SagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArnInput", GoGetter: "SagemakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerUserProfileUserSettingsKernelGatewayAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsKernelGatewayAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsKernelGatewayAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "customImage", GoGetter: "CustomImage"},
			_jsii_.MemberProperty{JsiiProperty: "customImageInput", GoGetter: "CustomImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArns", GoGetter: "LifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnsInput", GoGetter: "LifecycleConfigArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomImage", GoMethod: "ResetCustomImage"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArns", GoMethod: "ResetLifecycleConfigArns"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerUserProfileUserSettingsKernelGatewayAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsOutputReference",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "executionRole", GoGetter: "ExecutionRole"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleInput", GoGetter: "ExecutionRoleInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterServerAppSettings", GoGetter: "JupyterServerAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterServerAppSettingsInput", GoGetter: "JupyterServerAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayAppSettings", GoGetter: "KernelGatewayAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "kernelGatewayAppSettingsInput", GoGetter: "KernelGatewayAppSettingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putJupyterServerAppSettings", GoMethod: "PutJupyterServerAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putKernelGatewayAppSettings", GoMethod: "PutKernelGatewayAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putSharingSettings", GoMethod: "PutSharingSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putTensorBoardAppSettings", GoMethod: "PutTensorBoardAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetJupyterServerAppSettings", GoMethod: "ResetJupyterServerAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetKernelGatewayAppSettings", GoMethod: "ResetKernelGatewayAppSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroups", GoMethod: "ResetSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharingSettings", GoMethod: "ResetSharingSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetTensorBoardAppSettings", GoMethod: "ResetTensorBoardAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupsInput", GoGetter: "SecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharingSettings", GoGetter: "SharingSettings"},
			_jsii_.MemberProperty{JsiiProperty: "sharingSettingsInput", GoGetter: "SharingSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tensorBoardAppSettings", GoGetter: "TensorBoardAppSettings"},
			_jsii_.MemberProperty{JsiiProperty: "tensorBoardAppSettingsInput", GoGetter: "TensorBoardAppSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerUserProfileUserSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsSharingSettings",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsSharingSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsSharingSettingsOutputReference",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsSharingSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "notebookOutputOption", GoGetter: "NotebookOutputOption"},
			_jsii_.MemberProperty{JsiiProperty: "notebookOutputOptionInput", GoGetter: "NotebookOutputOptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotebookOutputOption", GoMethod: "ResetNotebookOutputOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3KmsKeyId", GoMethod: "ResetS3KmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3OutputPath", GoMethod: "ResetS3OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "s3KmsKeyId", GoGetter: "S3KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "s3KmsKeyIdInput", GoGetter: "S3KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPath", GoGetter: "S3OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "s3OutputPathInput", GoGetter: "S3OutputPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerUserProfileUserSettingsSharingSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsTensorBoardAppSettings",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsTensorBoardAppSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsTensorBoardAppSettingsDefaultResourceSpec",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsTensorBoardAppSettingsDefaultResourceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsTensorBoardAppSettingsDefaultResourceSpecOutputReference",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsTensorBoardAppSettingsDefaultResourceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArn", GoGetter: "LifecycleConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigArnInput", GoGetter: "LifecycleConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfigArn", GoMethod: "ResetLifecycleConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageArn", GoMethod: "ResetSagemakerImageArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSagemakerImageVersionArn", GoMethod: "ResetSagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArn", GoGetter: "SagemakerImageArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageArnInput", GoGetter: "SagemakerImageArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArn", GoGetter: "SagemakerImageVersionArn"},
			_jsii_.MemberProperty{JsiiProperty: "sagemakerImageVersionArnInput", GoGetter: "SagemakerImageVersionArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerUserProfileUserSettingsTensorBoardAppSettingsDefaultResourceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerUserProfileUserSettingsTensorBoardAppSettingsOutputReference",
		reflect.TypeOf((*SagemakerUserProfileUserSettingsTensorBoardAppSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpec", GoGetter: "DefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "defaultResourceSpecInput", GoGetter: "DefaultResourceSpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultResourceSpec", GoMethod: "PutDefaultResourceSpec"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerUserProfileUserSettingsTensorBoardAppSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerWorkforce",
		reflect.TypeOf((*SagemakerWorkforce)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cognitoConfig", GoGetter: "CognitoConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cognitoConfigInput", GoGetter: "CognitoConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "oidcConfig", GoGetter: "OidcConfig"},
			_jsii_.MemberProperty{JsiiProperty: "oidcConfigInput", GoGetter: "OidcConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putCognitoConfig", GoMethod: "PutCognitoConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putOidcConfig", GoMethod: "PutOidcConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSourceIpConfig", GoMethod: "PutSourceIpConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCognitoConfig", GoMethod: "ResetCognitoConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidcConfig", GoMethod: "ResetOidcConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceIpConfig", GoMethod: "ResetSourceIpConfig"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIpConfig", GoGetter: "SourceIpConfig"},
			_jsii_.MemberProperty{JsiiProperty: "sourceIpConfigInput", GoGetter: "SourceIpConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "subdomain", GoGetter: "Subdomain"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "workforceName", GoGetter: "WorkforceName"},
			_jsii_.MemberProperty{JsiiProperty: "workforceNameInput", GoGetter: "WorkforceNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerWorkforce{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerWorkforceCognitoConfig",
		reflect.TypeOf((*SagemakerWorkforceCognitoConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerWorkforceCognitoConfigOutputReference",
		reflect.TypeOf((*SagemakerWorkforceCognitoConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userPool", GoGetter: "UserPool"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolInput", GoGetter: "UserPoolInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerWorkforceCognitoConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerWorkforceConfig",
		reflect.TypeOf((*SagemakerWorkforceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerWorkforceOidcConfig",
		reflect.TypeOf((*SagemakerWorkforceOidcConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerWorkforceOidcConfigOutputReference",
		reflect.TypeOf((*SagemakerWorkforceOidcConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizationEndpoint", GoGetter: "AuthorizationEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationEndpointInput", GoGetter: "AuthorizationEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretInput", GoGetter: "ClientSecretInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "issuer", GoGetter: "Issuer"},
			_jsii_.MemberProperty{JsiiProperty: "issuerInput", GoGetter: "IssuerInput"},
			_jsii_.MemberProperty{JsiiProperty: "jwksUri", GoGetter: "JwksUri"},
			_jsii_.MemberProperty{JsiiProperty: "jwksUriInput", GoGetter: "JwksUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "logoutEndpoint", GoGetter: "LogoutEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "logoutEndpointInput", GoGetter: "LogoutEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpoint", GoGetter: "TokenEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "tokenEndpointInput", GoGetter: "TokenEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "userInfoEndpoint", GoGetter: "UserInfoEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "userInfoEndpointInput", GoGetter: "UserInfoEndpointInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerWorkforceOidcConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerWorkforceSourceIpConfig",
		reflect.TypeOf((*SagemakerWorkforceSourceIpConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerWorkforceSourceIpConfigOutputReference",
		reflect.TypeOf((*SagemakerWorkforceSourceIpConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cidrs", GoGetter: "Cidrs"},
			_jsii_.MemberProperty{JsiiProperty: "cidrsInput", GoGetter: "CidrsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerWorkforceSourceIpConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerWorkteam",
		reflect.TypeOf((*SagemakerWorkteam)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "memberDefinition", GoGetter: "MemberDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "memberDefinitionInput", GoGetter: "MemberDefinitionInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfiguration", GoGetter: "NotificationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "notificationConfigurationInput", GoGetter: "NotificationConfigurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putNotificationConfiguration", GoMethod: "PutNotificationConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationConfiguration", GoMethod: "ResetNotificationConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "subdomain", GoGetter: "Subdomain"},
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
			_jsii_.MemberProperty{JsiiProperty: "workforceName", GoGetter: "WorkforceName"},
			_jsii_.MemberProperty{JsiiProperty: "workforceNameInput", GoGetter: "WorkforceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "workteamName", GoGetter: "WorkteamName"},
			_jsii_.MemberProperty{JsiiProperty: "workteamNameInput", GoGetter: "WorkteamNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerWorkteam{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerWorkteamConfig",
		reflect.TypeOf((*SagemakerWorkteamConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerWorkteamMemberDefinition",
		reflect.TypeOf((*SagemakerWorkteamMemberDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerWorkteamMemberDefinitionCognitoMemberDefinition",
		reflect.TypeOf((*SagemakerWorkteamMemberDefinitionCognitoMemberDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerWorkteamMemberDefinitionCognitoMemberDefinitionOutputReference",
		reflect.TypeOf((*SagemakerWorkteamMemberDefinitionCognitoMemberDefinitionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userGroup", GoGetter: "UserGroup"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupInput", GoGetter: "UserGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "userPool", GoGetter: "UserPool"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolInput", GoGetter: "UserPoolInput"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerWorkteamMemberDefinitionCognitoMemberDefinitionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerWorkteamMemberDefinitionOidcMemberDefinition",
		reflect.TypeOf((*SagemakerWorkteamMemberDefinitionOidcMemberDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerWorkteamMemberDefinitionOidcMemberDefinitionOutputReference",
		reflect.TypeOf((*SagemakerWorkteamMemberDefinitionOidcMemberDefinitionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groups", GoGetter: "Groups"},
			_jsii_.MemberProperty{JsiiProperty: "groupsInput", GoGetter: "GroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerWorkteamMemberDefinitionOidcMemberDefinitionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.sagemaker.SagemakerWorkteamNotificationConfiguration",
		reflect.TypeOf((*SagemakerWorkteamNotificationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.sagemaker.SagemakerWorkteamNotificationConfigurationOutputReference",
		reflect.TypeOf((*SagemakerWorkteamNotificationConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "notificationTopicArn", GoGetter: "NotificationTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "notificationTopicArnInput", GoGetter: "NotificationTopicArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotificationTopicArn", GoMethod: "ResetNotificationTopicArn"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_SagemakerWorkteamNotificationConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
