package codebuild

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProject",
		reflect.TypeOf((*CodebuildProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "artifacts", GoGetter: "Artifacts"},
			_jsii_.MemberProperty{JsiiProperty: "artifactsInput", GoGetter: "ArtifactsInput"},
			_jsii_.MemberProperty{JsiiProperty: "badgeEnabled", GoGetter: "BadgeEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "badgeEnabledInput", GoGetter: "BadgeEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "badgeUrl", GoGetter: "BadgeUrl"},
			_jsii_.MemberProperty{JsiiProperty: "buildBatchConfig", GoGetter: "BuildBatchConfig"},
			_jsii_.MemberProperty{JsiiProperty: "buildBatchConfigInput", GoGetter: "BuildBatchConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "buildTimeout", GoGetter: "BuildTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "buildTimeoutInput", GoGetter: "BuildTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "cache", GoGetter: "Cache"},
			_jsii_.MemberProperty{JsiiProperty: "cacheInput", GoGetter: "CacheInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "concurrentBuildLimit", GoGetter: "ConcurrentBuildLimit"},
			_jsii_.MemberProperty{JsiiProperty: "concurrentBuildLimitInput", GoGetter: "ConcurrentBuildLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKeyInput", GoGetter: "EncryptionKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "environmentInput", GoGetter: "EnvironmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemLocations", GoGetter: "FileSystemLocations"},
			_jsii_.MemberProperty{JsiiProperty: "fileSystemLocationsInput", GoGetter: "FileSystemLocationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logsConfig", GoGetter: "LogsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logsConfigInput", GoGetter: "LogsConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putArtifacts", GoMethod: "PutArtifacts"},
			_jsii_.MemberMethod{JsiiMethod: "putBuildBatchConfig", GoMethod: "PutBuildBatchConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putCache", GoMethod: "PutCache"},
			_jsii_.MemberMethod{JsiiMethod: "putEnvironment", GoMethod: "PutEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "putLogsConfig", GoMethod: "PutLogsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSource", GoMethod: "PutSource"},
			_jsii_.MemberMethod{JsiiMethod: "putVpcConfig", GoMethod: "PutVpcConfig"},
			_jsii_.MemberProperty{JsiiProperty: "queuedTimeout", GoGetter: "QueuedTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "queuedTimeoutInput", GoGetter: "QueuedTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBadgeEnabled", GoMethod: "ResetBadgeEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuildBatchConfig", GoMethod: "ResetBuildBatchConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuildTimeout", GoMethod: "ResetBuildTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetCache", GoMethod: "ResetCache"},
			_jsii_.MemberMethod{JsiiMethod: "resetConcurrentBuildLimit", GoMethod: "ResetConcurrentBuildLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionKey", GoMethod: "ResetEncryptionKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileSystemLocations", GoMethod: "ResetFileSystemLocations"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogsConfig", GoMethod: "ResetLogsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueuedTimeout", GoMethod: "ResetQueuedTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecondaryArtifacts", GoMethod: "ResetSecondaryArtifacts"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecondarySources", GoMethod: "ResetSecondarySources"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceVersion", GoMethod: "ResetSourceVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcConfig", GoMethod: "ResetVpcConfig"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryArtifacts", GoGetter: "SecondaryArtifacts"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryArtifactsInput", GoGetter: "SecondaryArtifactsInput"},
			_jsii_.MemberProperty{JsiiProperty: "secondarySources", GoGetter: "SecondarySources"},
			_jsii_.MemberProperty{JsiiProperty: "secondarySourcesInput", GoGetter: "SecondarySourcesInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleInput", GoGetter: "ServiceRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceInput", GoGetter: "SourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVersion", GoGetter: "SourceVersion"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVersionInput", GoGetter: "SourceVersionInput"},
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
			j := jsiiProxy_CodebuildProject{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectArtifacts",
		reflect.TypeOf((*CodebuildProjectArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectArtifactsOutputReference",
		reflect.TypeOf((*CodebuildProjectArtifactsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "artifactIdentifier", GoGetter: "ArtifactIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "artifactIdentifierInput", GoGetter: "ArtifactIdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionDisabled", GoGetter: "EncryptionDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionDisabledInput", GoGetter: "EncryptionDisabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceType", GoGetter: "NamespaceType"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceTypeInput", GoGetter: "NamespaceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "overrideArtifactName", GoGetter: "OverrideArtifactName"},
			_jsii_.MemberProperty{JsiiProperty: "overrideArtifactNameInput", GoGetter: "OverrideArtifactNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "packaging", GoGetter: "Packaging"},
			_jsii_.MemberProperty{JsiiProperty: "packagingInput", GoGetter: "PackagingInput"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetArtifactIdentifier", GoMethod: "ResetArtifactIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionDisabled", GoMethod: "ResetEncryptionDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocation", GoMethod: "ResetLocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespaceType", GoMethod: "ResetNamespaceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideArtifactName", GoMethod: "ResetOverrideArtifactName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPackaging", GoMethod: "ResetPackaging"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectArtifactsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectBuildBatchConfig",
		reflect.TypeOf((*CodebuildProjectBuildBatchConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectBuildBatchConfigOutputReference",
		reflect.TypeOf((*CodebuildProjectBuildBatchConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "combineArtifacts", GoGetter: "CombineArtifacts"},
			_jsii_.MemberProperty{JsiiProperty: "combineArtifactsInput", GoGetter: "CombineArtifactsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putRestrictions", GoMethod: "PutRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCombineArtifacts", GoMethod: "ResetCombineArtifacts"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestrictions", GoMethod: "ResetRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeoutInMins", GoMethod: "ResetTimeoutInMins"},
			_jsii_.MemberProperty{JsiiProperty: "restrictions", GoGetter: "Restrictions"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionsInput", GoGetter: "RestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleInput", GoGetter: "ServiceRoleInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInMins", GoGetter: "TimeoutInMins"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInMinsInput", GoGetter: "TimeoutInMinsInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectBuildBatchConfigRestrictions",
		reflect.TypeOf((*CodebuildProjectBuildBatchConfigRestrictions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectBuildBatchConfigRestrictionsOutputReference",
		reflect.TypeOf((*CodebuildProjectBuildBatchConfigRestrictionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "computeTypesAllowed", GoGetter: "ComputeTypesAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "computeTypesAllowedInput", GoGetter: "ComputeTypesAllowedInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maximumBuildsAllowed", GoGetter: "MaximumBuildsAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "maximumBuildsAllowedInput", GoGetter: "MaximumBuildsAllowedInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetComputeTypesAllowed", GoMethod: "ResetComputeTypesAllowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumBuildsAllowed", GoMethod: "ResetMaximumBuildsAllowed"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectCache",
		reflect.TypeOf((*CodebuildProjectCache)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectCacheOutputReference",
		reflect.TypeOf((*CodebuildProjectCacheOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "modes", GoGetter: "Modes"},
			_jsii_.MemberProperty{JsiiProperty: "modesInput", GoGetter: "ModesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocation", GoMethod: "ResetLocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetModes", GoMethod: "ResetModes"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectCacheOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectConfig",
		reflect.TypeOf((*CodebuildProjectConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectEnvironment",
		reflect.TypeOf((*CodebuildProjectEnvironment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectEnvironmentEnvironmentVariable",
		reflect.TypeOf((*CodebuildProjectEnvironmentEnvironmentVariable)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectEnvironmentOutputReference",
		reflect.TypeOf((*CodebuildProjectEnvironmentOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "computeType", GoGetter: "ComputeType"},
			_jsii_.MemberProperty{JsiiProperty: "computeTypeInput", GoGetter: "ComputeTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVariable", GoGetter: "EnvironmentVariable"},
			_jsii_.MemberProperty{JsiiProperty: "environmentVariableInput", GoGetter: "EnvironmentVariableInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imageInput", GoGetter: "ImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "imagePullCredentialsType", GoGetter: "ImagePullCredentialsType"},
			_jsii_.MemberProperty{JsiiProperty: "imagePullCredentialsTypeInput", GoGetter: "ImagePullCredentialsTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "privilegedMode", GoGetter: "PrivilegedMode"},
			_jsii_.MemberProperty{JsiiProperty: "privilegedModeInput", GoGetter: "PrivilegedModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putRegistryCredential", GoMethod: "PutRegistryCredential"},
			_jsii_.MemberProperty{JsiiProperty: "registryCredential", GoGetter: "RegistryCredential"},
			_jsii_.MemberProperty{JsiiProperty: "registryCredentialInput", GoGetter: "RegistryCredentialInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironmentVariable", GoMethod: "ResetEnvironmentVariable"},
			_jsii_.MemberMethod{JsiiMethod: "resetImagePullCredentialsType", GoMethod: "ResetImagePullCredentialsType"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivilegedMode", GoMethod: "ResetPrivilegedMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegistryCredential", GoMethod: "ResetRegistryCredential"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectEnvironmentOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectEnvironmentRegistryCredential",
		reflect.TypeOf((*CodebuildProjectEnvironmentRegistryCredential)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectEnvironmentRegistryCredentialOutputReference",
		reflect.TypeOf((*CodebuildProjectEnvironmentRegistryCredentialOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "credential", GoGetter: "Credential"},
			_jsii_.MemberProperty{JsiiProperty: "credentialInput", GoGetter: "CredentialInput"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProvider", GoGetter: "CredentialProvider"},
			_jsii_.MemberProperty{JsiiProperty: "credentialProviderInput", GoGetter: "CredentialProviderInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectFileSystemLocations",
		reflect.TypeOf((*CodebuildProjectFileSystemLocations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectLogsConfig",
		reflect.TypeOf((*CodebuildProjectLogsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectLogsConfigCloudwatchLogs",
		reflect.TypeOf((*CodebuildProjectLogsConfigCloudwatchLogs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectLogsConfigCloudwatchLogsOutputReference",
		reflect.TypeOf((*CodebuildProjectLogsConfigCloudwatchLogsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
			_jsii_.MemberProperty{JsiiProperty: "groupNameInput", GoGetter: "GroupNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupName", GoMethod: "ResetGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetStreamName", GoMethod: "ResetStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "streamName", GoGetter: "StreamName"},
			_jsii_.MemberProperty{JsiiProperty: "streamNameInput", GoGetter: "StreamNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectLogsConfigOutputReference",
		reflect.TypeOf((*CodebuildProjectLogsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogs", GoGetter: "CloudwatchLogs"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogsInput", GoGetter: "CloudwatchLogsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putCloudwatchLogs", GoMethod: "PutCloudwatchLogs"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Logs", GoMethod: "PutS3Logs"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLogs", GoMethod: "ResetCloudwatchLogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3Logs", GoMethod: "ResetS3Logs"},
			_jsii_.MemberProperty{JsiiProperty: "s3Logs", GoGetter: "S3Logs"},
			_jsii_.MemberProperty{JsiiProperty: "s3LogsInput", GoGetter: "S3LogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectLogsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectLogsConfigS3Logs",
		reflect.TypeOf((*CodebuildProjectLogsConfigS3Logs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectLogsConfigS3LogsOutputReference",
		reflect.TypeOf((*CodebuildProjectLogsConfigS3LogsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "encryptionDisabled", GoGetter: "EncryptionDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionDisabledInput", GoGetter: "EncryptionDisabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionDisabled", GoMethod: "ResetEncryptionDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocation", GoMethod: "ResetLocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectSecondaryArtifacts",
		reflect.TypeOf((*CodebuildProjectSecondaryArtifacts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectSecondarySources",
		reflect.TypeOf((*CodebuildProjectSecondarySources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectSecondarySourcesAuth",
		reflect.TypeOf((*CodebuildProjectSecondarySourcesAuth)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectSecondarySourcesAuthOutputReference",
		reflect.TypeOf((*CodebuildProjectSecondarySourcesAuthOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetResource", GoMethod: "ResetResource"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceInput", GoGetter: "ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectSecondarySourcesBuildStatusConfig",
		reflect.TypeOf((*CodebuildProjectSecondarySourcesBuildStatusConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference",
		reflect.TypeOf((*CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "contextInput", GoGetter: "ContextInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetContext", GoMethod: "ResetContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetUrl", GoMethod: "ResetTargetUrl"},
			_jsii_.MemberProperty{JsiiProperty: "targetUrl", GoGetter: "TargetUrl"},
			_jsii_.MemberProperty{JsiiProperty: "targetUrlInput", GoGetter: "TargetUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectSecondarySourcesGitSubmodulesConfig",
		reflect.TypeOf((*CodebuildProjectSecondarySourcesGitSubmodulesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference",
		reflect.TypeOf((*CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fetchSubmodules", GoGetter: "FetchSubmodules"},
			_jsii_.MemberProperty{JsiiProperty: "fetchSubmodulesInput", GoGetter: "FetchSubmodulesInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectSource",
		reflect.TypeOf((*CodebuildProjectSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectSourceAuth",
		reflect.TypeOf((*CodebuildProjectSourceAuth)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectSourceAuthOutputReference",
		reflect.TypeOf((*CodebuildProjectSourceAuthOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetResource", GoMethod: "ResetResource"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberProperty{JsiiProperty: "resourceInput", GoGetter: "ResourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectSourceAuthOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectSourceBuildStatusConfig",
		reflect.TypeOf((*CodebuildProjectSourceBuildStatusConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectSourceBuildStatusConfigOutputReference",
		reflect.TypeOf((*CodebuildProjectSourceBuildStatusConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "contextInput", GoGetter: "ContextInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetContext", GoMethod: "ResetContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetUrl", GoMethod: "ResetTargetUrl"},
			_jsii_.MemberProperty{JsiiProperty: "targetUrl", GoGetter: "TargetUrl"},
			_jsii_.MemberProperty{JsiiProperty: "targetUrlInput", GoGetter: "TargetUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectSourceGitSubmodulesConfig",
		reflect.TypeOf((*CodebuildProjectSourceGitSubmodulesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectSourceGitSubmodulesConfigOutputReference",
		reflect.TypeOf((*CodebuildProjectSourceGitSubmodulesConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fetchSubmodules", GoGetter: "FetchSubmodules"},
			_jsii_.MemberProperty{JsiiProperty: "fetchSubmodulesInput", GoGetter: "FetchSubmodulesInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectSourceOutputReference",
		reflect.TypeOf((*CodebuildProjectSourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "auth", GoGetter: "Auth"},
			_jsii_.MemberProperty{JsiiProperty: "authInput", GoGetter: "AuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "buildspec", GoGetter: "Buildspec"},
			_jsii_.MemberProperty{JsiiProperty: "buildspecInput", GoGetter: "BuildspecInput"},
			_jsii_.MemberProperty{JsiiProperty: "buildStatusConfig", GoGetter: "BuildStatusConfig"},
			_jsii_.MemberProperty{JsiiProperty: "buildStatusConfigInput", GoGetter: "BuildStatusConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gitCloneDepth", GoGetter: "GitCloneDepth"},
			_jsii_.MemberProperty{JsiiProperty: "gitCloneDepthInput", GoGetter: "GitCloneDepthInput"},
			_jsii_.MemberProperty{JsiiProperty: "gitSubmodulesConfig", GoGetter: "GitSubmodulesConfig"},
			_jsii_.MemberProperty{JsiiProperty: "gitSubmodulesConfigInput", GoGetter: "GitSubmodulesConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecureSsl", GoGetter: "InsecureSsl"},
			_jsii_.MemberProperty{JsiiProperty: "insecureSslInput", GoGetter: "InsecureSslInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuth", GoMethod: "PutAuth"},
			_jsii_.MemberMethod{JsiiMethod: "putBuildStatusConfig", GoMethod: "PutBuildStatusConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putGitSubmodulesConfig", GoMethod: "PutGitSubmodulesConfig"},
			_jsii_.MemberProperty{JsiiProperty: "reportBuildStatus", GoGetter: "ReportBuildStatus"},
			_jsii_.MemberProperty{JsiiProperty: "reportBuildStatusInput", GoGetter: "ReportBuildStatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuth", GoMethod: "ResetAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuildspec", GoMethod: "ResetBuildspec"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuildStatusConfig", GoMethod: "ResetBuildStatusConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetGitCloneDepth", GoMethod: "ResetGitCloneDepth"},
			_jsii_.MemberMethod{JsiiMethod: "resetGitSubmodulesConfig", GoMethod: "ResetGitSubmodulesConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureSsl", GoMethod: "ResetInsecureSsl"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocation", GoMethod: "ResetLocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetReportBuildStatus", GoMethod: "ResetReportBuildStatus"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectSourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildProjectVpcConfig",
		reflect.TypeOf((*CodebuildProjectVpcConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildProjectVpcConfigOutputReference",
		reflect.TypeOf((*CodebuildProjectVpcConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberProperty{JsiiProperty: "subnetsInput", GoGetter: "SubnetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIdInput", GoGetter: "VpcIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildProjectVpcConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildReportGroup",
		reflect.TypeOf((*CodebuildReportGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "created", GoGetter: "Created"},
			_jsii_.MemberProperty{JsiiProperty: "deleteReports", GoGetter: "DeleteReports"},
			_jsii_.MemberProperty{JsiiProperty: "deleteReportsInput", GoGetter: "DeleteReportsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "exportConfig", GoGetter: "ExportConfig"},
			_jsii_.MemberProperty{JsiiProperty: "exportConfigInput", GoGetter: "ExportConfigInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putExportConfig", GoMethod: "PutExportConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteReports", GoMethod: "ResetDeleteReports"},
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
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildReportGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildReportGroupConfig",
		reflect.TypeOf((*CodebuildReportGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildReportGroupExportConfig",
		reflect.TypeOf((*CodebuildReportGroupExportConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildReportGroupExportConfigOutputReference",
		reflect.TypeOf((*CodebuildReportGroupExportConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putS3Destination", GoMethod: "PutS3Destination"},
			_jsii_.MemberMethod{JsiiMethod: "resetS3Destination", GoMethod: "ResetS3Destination"},
			_jsii_.MemberProperty{JsiiProperty: "s3Destination", GoGetter: "S3Destination"},
			_jsii_.MemberProperty{JsiiProperty: "s3DestinationInput", GoGetter: "S3DestinationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildReportGroupExportConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildReportGroupExportConfigS3Destination",
		reflect.TypeOf((*CodebuildReportGroupExportConfigS3Destination)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildReportGroupExportConfigS3DestinationOutputReference",
		reflect.TypeOf((*CodebuildReportGroupExportConfigS3DestinationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionDisabled", GoGetter: "EncryptionDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionDisabledInput", GoGetter: "EncryptionDisabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKeyInput", GoGetter: "EncryptionKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "packaging", GoGetter: "Packaging"},
			_jsii_.MemberProperty{JsiiProperty: "packagingInput", GoGetter: "PackagingInput"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionDisabled", GoMethod: "ResetEncryptionDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetPackaging", GoMethod: "ResetPackaging"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildSourceCredential",
		reflect.TypeOf((*CodebuildSourceCredential)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "authType", GoGetter: "AuthType"},
			_jsii_.MemberProperty{JsiiProperty: "authTypeInput", GoGetter: "AuthTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetUserName", GoMethod: "ResetUserName"},
			_jsii_.MemberProperty{JsiiProperty: "serverType", GoGetter: "ServerType"},
			_jsii_.MemberProperty{JsiiProperty: "serverTypeInput", GoGetter: "ServerTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
			_jsii_.MemberProperty{JsiiProperty: "userNameInput", GoGetter: "UserNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildSourceCredential{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildSourceCredentialConfig",
		reflect.TypeOf((*CodebuildSourceCredentialConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CodeBuild.CodebuildWebhook",
		reflect.TypeOf((*CodebuildWebhook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "branchFilter", GoGetter: "BranchFilter"},
			_jsii_.MemberProperty{JsiiProperty: "branchFilterInput", GoGetter: "BranchFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "buildType", GoGetter: "BuildType"},
			_jsii_.MemberProperty{JsiiProperty: "buildTypeInput", GoGetter: "BuildTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "filterGroup", GoGetter: "FilterGroup"},
			_jsii_.MemberProperty{JsiiProperty: "filterGroupInput", GoGetter: "FilterGroupInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "payloadUrl", GoGetter: "PayloadUrl"},
			_jsii_.MemberProperty{JsiiProperty: "projectName", GoGetter: "ProjectName"},
			_jsii_.MemberProperty{JsiiProperty: "projectNameInput", GoGetter: "ProjectNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBranchFilter", GoMethod: "ResetBranchFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuildType", GoMethod: "ResetBuildType"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilterGroup", GoMethod: "ResetFilterGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_CodebuildWebhook{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildWebhookConfig",
		reflect.TypeOf((*CodebuildWebhookConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildWebhookFilterGroup",
		reflect.TypeOf((*CodebuildWebhookFilterGroup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CodeBuild.CodebuildWebhookFilterGroupFilter",
		reflect.TypeOf((*CodebuildWebhookFilterGroupFilter)(nil)).Elem(),
	)
}
