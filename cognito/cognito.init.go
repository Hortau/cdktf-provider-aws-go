package cognito

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoIdentityPool",
		reflect.TypeOf((*CognitoIdentityPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowClassicFlow", GoGetter: "AllowClassicFlow"},
			_jsii_.MemberProperty{JsiiProperty: "allowClassicFlowInput", GoGetter: "AllowClassicFlowInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowUnauthenticatedIdentities", GoGetter: "AllowUnauthenticatedIdentities"},
			_jsii_.MemberProperty{JsiiProperty: "allowUnauthenticatedIdentitiesInput", GoGetter: "AllowUnauthenticatedIdentitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cognitoIdentityProviders", GoGetter: "CognitoIdentityProviders"},
			_jsii_.MemberProperty{JsiiProperty: "cognitoIdentityProvidersInput", GoGetter: "CognitoIdentityProvidersInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "developerProviderName", GoGetter: "DeveloperProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "developerProviderNameInput", GoGetter: "DeveloperProviderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolName", GoGetter: "IdentityPoolName"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolNameInput", GoGetter: "IdentityPoolNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openidConnectProviderArns", GoGetter: "OpenidConnectProviderArns"},
			_jsii_.MemberProperty{JsiiProperty: "openidConnectProviderArnsInput", GoGetter: "OpenidConnectProviderArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowClassicFlow", GoMethod: "ResetAllowClassicFlow"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowUnauthenticatedIdentities", GoMethod: "ResetAllowUnauthenticatedIdentities"},
			_jsii_.MemberMethod{JsiiMethod: "resetCognitoIdentityProviders", GoMethod: "ResetCognitoIdentityProviders"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeveloperProviderName", GoMethod: "ResetDeveloperProviderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenidConnectProviderArns", GoMethod: "ResetOpenidConnectProviderArns"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSamlProviderArns", GoMethod: "ResetSamlProviderArns"},
			_jsii_.MemberMethod{JsiiMethod: "resetSupportedLoginProviders", GoMethod: "ResetSupportedLoginProviders"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "samlProviderArns", GoGetter: "SamlProviderArns"},
			_jsii_.MemberProperty{JsiiProperty: "samlProviderArnsInput", GoGetter: "SamlProviderArnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "supportedLoginProviders", GoGetter: "SupportedLoginProviders"},
			_jsii_.MemberProperty{JsiiProperty: "supportedLoginProvidersInput", GoGetter: "SupportedLoginProvidersInput"},
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
			j := jsiiProxy_CognitoIdentityPool{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoIdentityPoolCognitoIdentityProviders",
		reflect.TypeOf((*CognitoIdentityPoolCognitoIdentityProviders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoIdentityPoolConfig",
		reflect.TypeOf((*CognitoIdentityPoolConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoIdentityPoolRolesAttachment",
		reflect.TypeOf((*CognitoIdentityPoolRolesAttachment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityPoolId", GoGetter: "IdentityPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "identityPoolIdInput", GoGetter: "IdentityPoolIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleMapping", GoMethod: "ResetRoleMapping"},
			_jsii_.MemberProperty{JsiiProperty: "roleMapping", GoGetter: "RoleMapping"},
			_jsii_.MemberProperty{JsiiProperty: "roleMappingInput", GoGetter: "RoleMappingInput"},
			_jsii_.MemberProperty{JsiiProperty: "roles", GoGetter: "Roles"},
			_jsii_.MemberProperty{JsiiProperty: "rolesInput", GoGetter: "RolesInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoIdentityPoolRolesAttachment{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoIdentityPoolRolesAttachmentConfig",
		reflect.TypeOf((*CognitoIdentityPoolRolesAttachmentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoIdentityPoolRolesAttachmentRoleMapping",
		reflect.TypeOf((*CognitoIdentityPoolRolesAttachmentRoleMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoIdentityPoolRolesAttachmentRoleMappingMappingRule",
		reflect.TypeOf((*CognitoIdentityPoolRolesAttachmentRoleMappingMappingRule)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoIdentityProvider",
		reflect.TypeOf((*CognitoIdentityProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "attributeMapping", GoGetter: "AttributeMapping"},
			_jsii_.MemberProperty{JsiiProperty: "attributeMappingInput", GoGetter: "AttributeMappingInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idpIdentifiers", GoGetter: "IdpIdentifiers"},
			_jsii_.MemberProperty{JsiiProperty: "idpIdentifiersInput", GoGetter: "IdpIdentifiersInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "providerDetails", GoGetter: "ProviderDetails"},
			_jsii_.MemberProperty{JsiiProperty: "providerDetailsInput", GoGetter: "ProviderDetailsInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerName", GoGetter: "ProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "providerNameInput", GoGetter: "ProviderNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerType", GoGetter: "ProviderType"},
			_jsii_.MemberProperty{JsiiProperty: "providerTypeInput", GoGetter: "ProviderTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttributeMapping", GoMethod: "ResetAttributeMapping"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdpIdentifiers", GoMethod: "ResetIdpIdentifiers"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolId", GoGetter: "UserPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdInput", GoGetter: "UserPoolIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoIdentityProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoIdentityProviderConfig",
		reflect.TypeOf((*CognitoIdentityProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoResourceServer",
		reflect.TypeOf((*CognitoResourceServer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
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
			_jsii_.MemberProperty{JsiiProperty: "identifier", GoGetter: "Identifier"},
			_jsii_.MemberProperty{JsiiProperty: "identifierInput", GoGetter: "IdentifierInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetScope", GoMethod: "ResetScope"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "scopeIdentifiers", GoGetter: "ScopeIdentifiers"},
			_jsii_.MemberProperty{JsiiProperty: "scopeInput", GoGetter: "ScopeInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolId", GoGetter: "UserPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdInput", GoGetter: "UserPoolIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoResourceServer{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoResourceServerConfig",
		reflect.TypeOf((*CognitoResourceServerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoResourceServerScope",
		reflect.TypeOf((*CognitoResourceServerScope)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserGroup",
		reflect.TypeOf((*CognitoUserGroup)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "precedence", GoGetter: "Precedence"},
			_jsii_.MemberProperty{JsiiProperty: "precedenceInput", GoGetter: "PrecedenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrecedence", GoMethod: "ResetPrecedence"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleArn", GoMethod: "ResetRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolId", GoGetter: "UserPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdInput", GoGetter: "UserPoolIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserGroupConfig",
		reflect.TypeOf((*CognitoUserGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPool",
		reflect.TypeOf((*CognitoUserPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountRecoverySetting", GoGetter: "AccountRecoverySetting"},
			_jsii_.MemberProperty{JsiiProperty: "accountRecoverySettingInput", GoGetter: "AccountRecoverySettingInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "adminCreateUserConfig", GoGetter: "AdminCreateUserConfig"},
			_jsii_.MemberProperty{JsiiProperty: "adminCreateUserConfigInput", GoGetter: "AdminCreateUserConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "aliasAttributes", GoGetter: "AliasAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "aliasAttributesInput", GoGetter: "AliasAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "autoVerifiedAttributes", GoGetter: "AutoVerifiedAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "autoVerifiedAttributesInput", GoGetter: "AutoVerifiedAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creationDate", GoGetter: "CreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "customDomain", GoGetter: "CustomDomain"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "deviceConfiguration", GoGetter: "DeviceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "deviceConfigurationInput", GoGetter: "DeviceConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "emailConfiguration", GoGetter: "EmailConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "emailConfigurationInput", GoGetter: "EmailConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailVerificationMessage", GoGetter: "EmailVerificationMessage"},
			_jsii_.MemberProperty{JsiiProperty: "emailVerificationMessageInput", GoGetter: "EmailVerificationMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailVerificationSubject", GoGetter: "EmailVerificationSubject"},
			_jsii_.MemberProperty{JsiiProperty: "emailVerificationSubjectInput", GoGetter: "EmailVerificationSubjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedNumberOfUsers", GoGetter: "EstimatedNumberOfUsers"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaConfig", GoGetter: "LambdaConfig"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaConfigInput", GoGetter: "LambdaConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "lastModifiedDate", GoGetter: "LastModifiedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mfaConfiguration", GoGetter: "MfaConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "mfaConfigurationInput", GoGetter: "MfaConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordPolicy", GoGetter: "PasswordPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "passwordPolicyInput", GoGetter: "PasswordPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putAccountRecoverySetting", GoMethod: "PutAccountRecoverySetting"},
			_jsii_.MemberMethod{JsiiMethod: "putAdminCreateUserConfig", GoMethod: "PutAdminCreateUserConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putDeviceConfiguration", GoMethod: "PutDeviceConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putEmailConfiguration", GoMethod: "PutEmailConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putLambdaConfig", GoMethod: "PutLambdaConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putPasswordPolicy", GoMethod: "PutPasswordPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putSmsConfiguration", GoMethod: "PutSmsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putSoftwareTokenMfaConfiguration", GoMethod: "PutSoftwareTokenMfaConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putUsernameConfiguration", GoMethod: "PutUsernameConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putUserPoolAddOns", GoMethod: "PutUserPoolAddOns"},
			_jsii_.MemberMethod{JsiiMethod: "putVerificationMessageTemplate", GoMethod: "PutVerificationMessageTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountRecoverySetting", GoMethod: "ResetAccountRecoverySetting"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminCreateUserConfig", GoMethod: "ResetAdminCreateUserConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAliasAttributes", GoMethod: "ResetAliasAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoVerifiedAttributes", GoMethod: "ResetAutoVerifiedAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceConfiguration", GoMethod: "ResetDeviceConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailConfiguration", GoMethod: "ResetEmailConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailVerificationMessage", GoMethod: "ResetEmailVerificationMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailVerificationSubject", GoMethod: "ResetEmailVerificationSubject"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaConfig", GoMethod: "ResetLambdaConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMfaConfiguration", GoMethod: "ResetMfaConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordPolicy", GoMethod: "ResetPasswordPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchema", GoMethod: "ResetSchema"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmsAuthenticationMessage", GoMethod: "ResetSmsAuthenticationMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmsConfiguration", GoMethod: "ResetSmsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmsVerificationMessage", GoMethod: "ResetSmsVerificationMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetSoftwareTokenMfaConfiguration", GoMethod: "ResetSoftwareTokenMfaConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsernameAttributes", GoMethod: "ResetUsernameAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsernameConfiguration", GoMethod: "ResetUsernameConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserPoolAddOns", GoMethod: "ResetUserPoolAddOns"},
			_jsii_.MemberMethod{JsiiMethod: "resetVerificationMessageTemplate", GoMethod: "ResetVerificationMessageTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "smsAuthenticationMessage", GoGetter: "SmsAuthenticationMessage"},
			_jsii_.MemberProperty{JsiiProperty: "smsAuthenticationMessageInput", GoGetter: "SmsAuthenticationMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "smsConfiguration", GoGetter: "SmsConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "smsConfigurationInput", GoGetter: "SmsConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "smsVerificationMessage", GoGetter: "SmsVerificationMessage"},
			_jsii_.MemberProperty{JsiiProperty: "smsVerificationMessageInput", GoGetter: "SmsVerificationMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "softwareTokenMfaConfiguration", GoGetter: "SoftwareTokenMfaConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "softwareTokenMfaConfigurationInput", GoGetter: "SoftwareTokenMfaConfigurationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "usernameAttributes", GoGetter: "UsernameAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "usernameAttributesInput", GoGetter: "UsernameAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "usernameConfiguration", GoGetter: "UsernameConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "usernameConfigurationInput", GoGetter: "UsernameConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolAddOns", GoGetter: "UserPoolAddOns"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolAddOnsInput", GoGetter: "UserPoolAddOnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "verificationMessageTemplate", GoGetter: "VerificationMessageTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "verificationMessageTemplateInput", GoGetter: "VerificationMessageTemplateInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPool{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolAccountRecoverySetting",
		reflect.TypeOf((*CognitoUserPoolAccountRecoverySetting)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolAccountRecoverySettingOutputReference",
		reflect.TypeOf((*CognitoUserPoolAccountRecoverySettingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryMechanism", GoGetter: "RecoveryMechanism"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryMechanismInput", GoGetter: "RecoveryMechanismInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolAccountRecoverySettingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolAccountRecoverySettingRecoveryMechanism",
		reflect.TypeOf((*CognitoUserPoolAccountRecoverySettingRecoveryMechanism)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolAdminCreateUserConfig",
		reflect.TypeOf((*CognitoUserPoolAdminCreateUserConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate",
		reflect.TypeOf((*CognitoUserPoolAdminCreateUserConfigInviteMessageTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference",
		reflect.TypeOf((*CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "emailMessage", GoGetter: "EmailMessage"},
			_jsii_.MemberProperty{JsiiProperty: "emailMessageInput", GoGetter: "EmailMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailSubject", GoGetter: "EmailSubject"},
			_jsii_.MemberProperty{JsiiProperty: "emailSubjectInput", GoGetter: "EmailSubjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailMessage", GoMethod: "ResetEmailMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailSubject", GoMethod: "ResetEmailSubject"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmsMessage", GoMethod: "ResetSmsMessage"},
			_jsii_.MemberProperty{JsiiProperty: "smsMessage", GoGetter: "SmsMessage"},
			_jsii_.MemberProperty{JsiiProperty: "smsMessageInput", GoGetter: "SmsMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolAdminCreateUserConfigInviteMessageTemplateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolAdminCreateUserConfigOutputReference",
		reflect.TypeOf((*CognitoUserPoolAdminCreateUserConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowAdminCreateUserOnly", GoGetter: "AllowAdminCreateUserOnly"},
			_jsii_.MemberProperty{JsiiProperty: "allowAdminCreateUserOnlyInput", GoGetter: "AllowAdminCreateUserOnlyInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "inviteMessageTemplate", GoGetter: "InviteMessageTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "inviteMessageTemplateInput", GoGetter: "InviteMessageTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putInviteMessageTemplate", GoMethod: "PutInviteMessageTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowAdminCreateUserOnly", GoMethod: "ResetAllowAdminCreateUserOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetInviteMessageTemplate", GoMethod: "ResetInviteMessageTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolAdminCreateUserConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolClient",
		reflect.TypeOf((*CognitoUserPoolClient)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessTokenValidity", GoGetter: "AccessTokenValidity"},
			_jsii_.MemberProperty{JsiiProperty: "accessTokenValidityInput", GoGetter: "AccessTokenValidityInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthFlows", GoGetter: "AllowedOauthFlows"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthFlowsInput", GoGetter: "AllowedOauthFlowsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthFlowsUserPoolClient", GoGetter: "AllowedOauthFlowsUserPoolClient"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthFlowsUserPoolClientInput", GoGetter: "AllowedOauthFlowsUserPoolClientInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthScopes", GoGetter: "AllowedOauthScopes"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOauthScopesInput", GoGetter: "AllowedOauthScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "analyticsConfiguration", GoGetter: "AnalyticsConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "analyticsConfigurationInput", GoGetter: "AnalyticsConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrls", GoGetter: "CallbackUrls"},
			_jsii_.MemberProperty{JsiiProperty: "callbackUrlsInput", GoGetter: "CallbackUrlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRedirectUri", GoGetter: "DefaultRedirectUri"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRedirectUriInput", GoGetter: "DefaultRedirectUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableTokenRevocation", GoGetter: "EnableTokenRevocation"},
			_jsii_.MemberProperty{JsiiProperty: "enableTokenRevocationInput", GoGetter: "EnableTokenRevocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "explicitAuthFlows", GoGetter: "ExplicitAuthFlows"},
			_jsii_.MemberProperty{JsiiProperty: "explicitAuthFlowsInput", GoGetter: "ExplicitAuthFlowsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "generateSecret", GoGetter: "GenerateSecret"},
			_jsii_.MemberProperty{JsiiProperty: "generateSecretInput", GoGetter: "GenerateSecretInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idTokenValidity", GoGetter: "IdTokenValidity"},
			_jsii_.MemberProperty{JsiiProperty: "idTokenValidityInput", GoGetter: "IdTokenValidityInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logoutUrls", GoGetter: "LogoutUrls"},
			_jsii_.MemberProperty{JsiiProperty: "logoutUrlsInput", GoGetter: "LogoutUrlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preventUserExistenceErrors", GoGetter: "PreventUserExistenceErrors"},
			_jsii_.MemberProperty{JsiiProperty: "preventUserExistenceErrorsInput", GoGetter: "PreventUserExistenceErrorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putAnalyticsConfiguration", GoMethod: "PutAnalyticsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putTokenValidityUnits", GoMethod: "PutTokenValidityUnits"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "readAttributes", GoGetter: "ReadAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "readAttributesInput", GoGetter: "ReadAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenValidity", GoGetter: "RefreshTokenValidity"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenValidityInput", GoGetter: "RefreshTokenValidityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessTokenValidity", GoMethod: "ResetAccessTokenValidity"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOauthFlows", GoMethod: "ResetAllowedOauthFlows"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOauthFlowsUserPoolClient", GoMethod: "ResetAllowedOauthFlowsUserPoolClient"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOauthScopes", GoMethod: "ResetAllowedOauthScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnalyticsConfiguration", GoMethod: "ResetAnalyticsConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetCallbackUrls", GoMethod: "ResetCallbackUrls"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRedirectUri", GoMethod: "ResetDefaultRedirectUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableTokenRevocation", GoMethod: "ResetEnableTokenRevocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetExplicitAuthFlows", GoMethod: "ResetExplicitAuthFlows"},
			_jsii_.MemberMethod{JsiiMethod: "resetGenerateSecret", GoMethod: "ResetGenerateSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdTokenValidity", GoMethod: "ResetIdTokenValidity"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogoutUrls", GoMethod: "ResetLogoutUrls"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreventUserExistenceErrors", GoMethod: "ResetPreventUserExistenceErrors"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadAttributes", GoMethod: "ResetReadAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshTokenValidity", GoMethod: "ResetRefreshTokenValidity"},
			_jsii_.MemberMethod{JsiiMethod: "resetSupportedIdentityProviders", GoMethod: "ResetSupportedIdentityProviders"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenValidityUnits", GoMethod: "ResetTokenValidityUnits"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteAttributes", GoMethod: "ResetWriteAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "supportedIdentityProviders", GoGetter: "SupportedIdentityProviders"},
			_jsii_.MemberProperty{JsiiProperty: "supportedIdentityProvidersInput", GoGetter: "SupportedIdentityProvidersInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tokenValidityUnits", GoGetter: "TokenValidityUnits"},
			_jsii_.MemberProperty{JsiiProperty: "tokenValidityUnitsInput", GoGetter: "TokenValidityUnitsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolId", GoGetter: "UserPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdInput", GoGetter: "UserPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "writeAttributes", GoGetter: "WriteAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "writeAttributesInput", GoGetter: "WriteAttributesInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClient{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolClientAnalyticsConfiguration",
		reflect.TypeOf((*CognitoUserPoolClientAnalyticsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolClientAnalyticsConfigurationOutputReference",
		reflect.TypeOf((*CognitoUserPoolClientAnalyticsConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "applicationArn", GoGetter: "ApplicationArn"},
			_jsii_.MemberProperty{JsiiProperty: "applicationArnInput", GoGetter: "ApplicationArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "applicationId", GoGetter: "ApplicationId"},
			_jsii_.MemberProperty{JsiiProperty: "applicationIdInput", GoGetter: "ApplicationIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalId", GoGetter: "ExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "externalIdInput", GoGetter: "ExternalIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationArn", GoMethod: "ResetApplicationArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetApplicationId", GoMethod: "ResetApplicationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalId", GoMethod: "ResetExternalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoleArn", GoMethod: "ResetRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserDataShared", GoMethod: "ResetUserDataShared"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userDataShared", GoGetter: "UserDataShared"},
			_jsii_.MemberProperty{JsiiProperty: "userDataSharedInput", GoGetter: "UserDataSharedInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClientAnalyticsConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolClientConfig",
		reflect.TypeOf((*CognitoUserPoolClientConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolClientTokenValidityUnits",
		reflect.TypeOf((*CognitoUserPoolClientTokenValidityUnits)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolClientTokenValidityUnitsOutputReference",
		reflect.TypeOf((*CognitoUserPoolClientTokenValidityUnitsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessToken", GoGetter: "AccessToken"},
			_jsii_.MemberProperty{JsiiProperty: "accessTokenInput", GoGetter: "AccessTokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idToken", GoGetter: "IdToken"},
			_jsii_.MemberProperty{JsiiProperty: "idTokenInput", GoGetter: "IdTokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "refreshToken", GoGetter: "RefreshToken"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenInput", GoGetter: "RefreshTokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessToken", GoMethod: "ResetAccessToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdToken", GoMethod: "ResetIdToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshToken", GoMethod: "ResetRefreshToken"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolClientTokenValidityUnitsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolConfig",
		reflect.TypeOf((*CognitoUserPoolConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolDeviceConfiguration",
		reflect.TypeOf((*CognitoUserPoolDeviceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolDeviceConfigurationOutputReference",
		reflect.TypeOf((*CognitoUserPoolDeviceConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "challengeRequiredOnNewDevice", GoGetter: "ChallengeRequiredOnNewDevice"},
			_jsii_.MemberProperty{JsiiProperty: "challengeRequiredOnNewDeviceInput", GoGetter: "ChallengeRequiredOnNewDeviceInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceOnlyRememberedOnUserPrompt", GoGetter: "DeviceOnlyRememberedOnUserPrompt"},
			_jsii_.MemberProperty{JsiiProperty: "deviceOnlyRememberedOnUserPromptInput", GoGetter: "DeviceOnlyRememberedOnUserPromptInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetChallengeRequiredOnNewDevice", GoMethod: "ResetChallengeRequiredOnNewDevice"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceOnlyRememberedOnUserPrompt", GoMethod: "ResetDeviceOnlyRememberedOnUserPrompt"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolDeviceConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolDomain",
		reflect.TypeOf((*CognitoUserPoolDomain)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certificateArn", GoGetter: "CertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "certificateArnInput", GoGetter: "CertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontDistributionArn", GoGetter: "CloudfrontDistributionArn"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainInput", GoGetter: "DomainInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateArn", GoMethod: "ResetCertificateArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "s3Bucket", GoGetter: "S3Bucket"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolId", GoGetter: "UserPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdInput", GoGetter: "UserPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolDomain{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolDomainConfig",
		reflect.TypeOf((*CognitoUserPoolDomainConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolEmailConfiguration",
		reflect.TypeOf((*CognitoUserPoolEmailConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolEmailConfigurationOutputReference",
		reflect.TypeOf((*CognitoUserPoolEmailConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configurationSet", GoGetter: "ConfigurationSet"},
			_jsii_.MemberProperty{JsiiProperty: "configurationSetInput", GoGetter: "ConfigurationSetInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailSendingAccount", GoGetter: "EmailSendingAccount"},
			_jsii_.MemberProperty{JsiiProperty: "emailSendingAccountInput", GoGetter: "EmailSendingAccountInput"},
			_jsii_.MemberProperty{JsiiProperty: "fromEmailAddress", GoGetter: "FromEmailAddress"},
			_jsii_.MemberProperty{JsiiProperty: "fromEmailAddressInput", GoGetter: "FromEmailAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "replyToEmailAddress", GoGetter: "ReplyToEmailAddress"},
			_jsii_.MemberProperty{JsiiProperty: "replyToEmailAddressInput", GoGetter: "ReplyToEmailAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigurationSet", GoMethod: "ResetConfigurationSet"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailSendingAccount", GoMethod: "ResetEmailSendingAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetFromEmailAddress", GoMethod: "ResetFromEmailAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetReplyToEmailAddress", GoMethod: "ResetReplyToEmailAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceArn", GoMethod: "ResetSourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArn", GoGetter: "SourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceArnInput", GoGetter: "SourceArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolEmailConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfig",
		reflect.TypeOf((*CognitoUserPoolLambdaConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigCustomEmailSender",
		reflect.TypeOf((*CognitoUserPoolLambdaConfigCustomEmailSender)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference",
		reflect.TypeOf((*CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaArn", GoGetter: "LambdaArn"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaArnInput", GoGetter: "LambdaArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaVersion", GoGetter: "LambdaVersion"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaVersionInput", GoGetter: "LambdaVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolLambdaConfigCustomEmailSenderOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigCustomSmsSender",
		reflect.TypeOf((*CognitoUserPoolLambdaConfigCustomSmsSender)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference",
		reflect.TypeOf((*CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaArn", GoGetter: "LambdaArn"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaArnInput", GoGetter: "LambdaArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaVersion", GoGetter: "LambdaVersion"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaVersionInput", GoGetter: "LambdaVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolLambdaConfigCustomSmsSenderOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolLambdaConfigOutputReference",
		reflect.TypeOf((*CognitoUserPoolLambdaConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "createAuthChallenge", GoGetter: "CreateAuthChallenge"},
			_jsii_.MemberProperty{JsiiProperty: "createAuthChallengeInput", GoGetter: "CreateAuthChallengeInput"},
			_jsii_.MemberProperty{JsiiProperty: "customEmailSender", GoGetter: "CustomEmailSender"},
			_jsii_.MemberProperty{JsiiProperty: "customEmailSenderInput", GoGetter: "CustomEmailSenderInput"},
			_jsii_.MemberProperty{JsiiProperty: "customMessage", GoGetter: "CustomMessage"},
			_jsii_.MemberProperty{JsiiProperty: "customMessageInput", GoGetter: "CustomMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "customSmsSender", GoGetter: "CustomSmsSender"},
			_jsii_.MemberProperty{JsiiProperty: "customSmsSenderInput", GoGetter: "CustomSmsSenderInput"},
			_jsii_.MemberProperty{JsiiProperty: "defineAuthChallenge", GoGetter: "DefineAuthChallenge"},
			_jsii_.MemberProperty{JsiiProperty: "defineAuthChallengeInput", GoGetter: "DefineAuthChallengeInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "postAuthentication", GoGetter: "PostAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "postAuthenticationInput", GoGetter: "PostAuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "postConfirmation", GoGetter: "PostConfirmation"},
			_jsii_.MemberProperty{JsiiProperty: "postConfirmationInput", GoGetter: "PostConfirmationInput"},
			_jsii_.MemberProperty{JsiiProperty: "preAuthentication", GoGetter: "PreAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "preAuthenticationInput", GoGetter: "PreAuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "preSignUp", GoGetter: "PreSignUp"},
			_jsii_.MemberProperty{JsiiProperty: "preSignUpInput", GoGetter: "PreSignUpInput"},
			_jsii_.MemberProperty{JsiiProperty: "preTokenGeneration", GoGetter: "PreTokenGeneration"},
			_jsii_.MemberProperty{JsiiProperty: "preTokenGenerationInput", GoGetter: "PreTokenGenerationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomEmailSender", GoMethod: "PutCustomEmailSender"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomSmsSender", GoMethod: "PutCustomSmsSender"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreateAuthChallenge", GoMethod: "ResetCreateAuthChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomEmailSender", GoMethod: "ResetCustomEmailSender"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomMessage", GoMethod: "ResetCustomMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomSmsSender", GoMethod: "ResetCustomSmsSender"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefineAuthChallenge", GoMethod: "ResetDefineAuthChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostAuthentication", GoMethod: "ResetPostAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostConfirmation", GoMethod: "ResetPostConfirmation"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreAuthentication", GoMethod: "ResetPreAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreSignUp", GoMethod: "ResetPreSignUp"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreTokenGeneration", GoMethod: "ResetPreTokenGeneration"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserMigration", GoMethod: "ResetUserMigration"},
			_jsii_.MemberMethod{JsiiMethod: "resetVerifyAuthChallengeResponse", GoMethod: "ResetVerifyAuthChallengeResponse"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "userMigration", GoGetter: "UserMigration"},
			_jsii_.MemberProperty{JsiiProperty: "userMigrationInput", GoGetter: "UserMigrationInput"},
			_jsii_.MemberProperty{JsiiProperty: "verifyAuthChallengeResponse", GoGetter: "VerifyAuthChallengeResponse"},
			_jsii_.MemberProperty{JsiiProperty: "verifyAuthChallengeResponseInput", GoGetter: "VerifyAuthChallengeResponseInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolLambdaConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolPasswordPolicy",
		reflect.TypeOf((*CognitoUserPoolPasswordPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolPasswordPolicyOutputReference",
		reflect.TypeOf((*CognitoUserPoolPasswordPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "minimumLength", GoGetter: "MinimumLength"},
			_jsii_.MemberProperty{JsiiProperty: "minimumLengthInput", GoGetter: "MinimumLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireLowercase", GoGetter: "RequireLowercase"},
			_jsii_.MemberProperty{JsiiProperty: "requireLowercaseInput", GoGetter: "RequireLowercaseInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireNumbers", GoGetter: "RequireNumbers"},
			_jsii_.MemberProperty{JsiiProperty: "requireNumbersInput", GoGetter: "RequireNumbersInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireSymbols", GoGetter: "RequireSymbols"},
			_jsii_.MemberProperty{JsiiProperty: "requireSymbolsInput", GoGetter: "RequireSymbolsInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireUppercase", GoGetter: "RequireUppercase"},
			_jsii_.MemberProperty{JsiiProperty: "requireUppercaseInput", GoGetter: "RequireUppercaseInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumLength", GoMethod: "ResetMinimumLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireLowercase", GoMethod: "ResetRequireLowercase"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireNumbers", GoMethod: "ResetRequireNumbers"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireSymbols", GoMethod: "ResetRequireSymbols"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireUppercase", GoMethod: "ResetRequireUppercase"},
			_jsii_.MemberMethod{JsiiMethod: "resetTemporaryPasswordValidityDays", GoMethod: "ResetTemporaryPasswordValidityDays"},
			_jsii_.MemberProperty{JsiiProperty: "temporaryPasswordValidityDays", GoGetter: "TemporaryPasswordValidityDays"},
			_jsii_.MemberProperty{JsiiProperty: "temporaryPasswordValidityDaysInput", GoGetter: "TemporaryPasswordValidityDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolPasswordPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolSchema",
		reflect.TypeOf((*CognitoUserPoolSchema)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolSchemaNumberAttributeConstraints",
		reflect.TypeOf((*CognitoUserPoolSchemaNumberAttributeConstraints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference",
		reflect.TypeOf((*CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxValue", GoGetter: "MaxValue"},
			_jsii_.MemberProperty{JsiiProperty: "maxValueInput", GoGetter: "MaxValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "minValue", GoGetter: "MinValue"},
			_jsii_.MemberProperty{JsiiProperty: "minValueInput", GoGetter: "MinValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxValue", GoMethod: "ResetMaxValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinValue", GoMethod: "ResetMinValue"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolSchemaNumberAttributeConstraintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolSchemaStringAttributeConstraints",
		reflect.TypeOf((*CognitoUserPoolSchemaStringAttributeConstraints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolSchemaStringAttributeConstraintsOutputReference",
		reflect.TypeOf((*CognitoUserPoolSchemaStringAttributeConstraintsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxLength", GoGetter: "MaxLength"},
			_jsii_.MemberProperty{JsiiProperty: "maxLengthInput", GoGetter: "MaxLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "minLength", GoGetter: "MinLength"},
			_jsii_.MemberProperty{JsiiProperty: "minLengthInput", GoGetter: "MinLengthInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxLength", GoMethod: "ResetMaxLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinLength", GoMethod: "ResetMinLength"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolSchemaStringAttributeConstraintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolSmsConfiguration",
		reflect.TypeOf((*CognitoUserPoolSmsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolSmsConfigurationOutputReference",
		reflect.TypeOf((*CognitoUserPoolSmsConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "externalId", GoGetter: "ExternalId"},
			_jsii_.MemberProperty{JsiiProperty: "externalIdInput", GoGetter: "ExternalIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "snsCallerArn", GoGetter: "SnsCallerArn"},
			_jsii_.MemberProperty{JsiiProperty: "snsCallerArnInput", GoGetter: "SnsCallerArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolSmsConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolSoftwareTokenMfaConfiguration",
		reflect.TypeOf((*CognitoUserPoolSoftwareTokenMfaConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference",
		reflect.TypeOf((*CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
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
			j := jsiiProxy_CognitoUserPoolSoftwareTokenMfaConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolUiCustomization",
		reflect.TypeOf((*CognitoUserPoolUiCustomization)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creationDate", GoGetter: "CreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "css", GoGetter: "Css"},
			_jsii_.MemberProperty{JsiiProperty: "cssInput", GoGetter: "CssInput"},
			_jsii_.MemberProperty{JsiiProperty: "cssVersion", GoGetter: "CssVersion"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "imageFile", GoGetter: "ImageFile"},
			_jsii_.MemberProperty{JsiiProperty: "imageFileInput", GoGetter: "ImageFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageUrl", GoGetter: "ImageUrl"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastModifiedDate", GoGetter: "LastModifiedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCss", GoMethod: "ResetCss"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageFile", GoMethod: "ResetImageFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolId", GoGetter: "UserPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "userPoolIdInput", GoGetter: "UserPoolIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolUiCustomization{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolUiCustomizationConfig",
		reflect.TypeOf((*CognitoUserPoolUiCustomizationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolUserPoolAddOns",
		reflect.TypeOf((*CognitoUserPoolUserPoolAddOns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolUserPoolAddOnsOutputReference",
		reflect.TypeOf((*CognitoUserPoolUserPoolAddOnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurityMode", GoGetter: "AdvancedSecurityMode"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurityModeInput", GoGetter: "AdvancedSecurityModeInput"},
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
			j := jsiiProxy_CognitoUserPoolUserPoolAddOnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolUsernameConfiguration",
		reflect.TypeOf((*CognitoUserPoolUsernameConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolUsernameConfigurationOutputReference",
		reflect.TypeOf((*CognitoUserPoolUsernameConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "caseSensitive", GoGetter: "CaseSensitive"},
			_jsii_.MemberProperty{JsiiProperty: "caseSensitiveInput", GoGetter: "CaseSensitiveInput"},
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
			j := jsiiProxy_CognitoUserPoolUsernameConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.CognitoUserPoolVerificationMessageTemplate",
		reflect.TypeOf((*CognitoUserPoolVerificationMessageTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.CognitoUserPoolVerificationMessageTemplateOutputReference",
		reflect.TypeOf((*CognitoUserPoolVerificationMessageTemplateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultEmailOption", GoGetter: "DefaultEmailOption"},
			_jsii_.MemberProperty{JsiiProperty: "defaultEmailOptionInput", GoGetter: "DefaultEmailOptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailMessage", GoGetter: "EmailMessage"},
			_jsii_.MemberProperty{JsiiProperty: "emailMessageByLink", GoGetter: "EmailMessageByLink"},
			_jsii_.MemberProperty{JsiiProperty: "emailMessageByLinkInput", GoGetter: "EmailMessageByLinkInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailMessageInput", GoGetter: "EmailMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailSubject", GoGetter: "EmailSubject"},
			_jsii_.MemberProperty{JsiiProperty: "emailSubjectByLink", GoGetter: "EmailSubjectByLink"},
			_jsii_.MemberProperty{JsiiProperty: "emailSubjectByLinkInput", GoGetter: "EmailSubjectByLinkInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailSubjectInput", GoGetter: "EmailSubjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultEmailOption", GoMethod: "ResetDefaultEmailOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailMessage", GoMethod: "ResetEmailMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailMessageByLink", GoMethod: "ResetEmailMessageByLink"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailSubject", GoMethod: "ResetEmailSubject"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailSubjectByLink", GoMethod: "ResetEmailSubjectByLink"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmsMessage", GoMethod: "ResetSmsMessage"},
			_jsii_.MemberProperty{JsiiProperty: "smsMessage", GoGetter: "SmsMessage"},
			_jsii_.MemberProperty{JsiiProperty: "smsMessageInput", GoGetter: "SmsMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CognitoUserPoolVerificationMessageTemplateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Cognito.DataAwsCognitoUserPools",
		reflect.TypeOf((*DataAwsCognitoUserPools)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arns", GoGetter: "Arns"},
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
			_jsii_.MemberProperty{JsiiProperty: "ids", GoGetter: "Ids"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
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
			j := jsiiProxy_DataAwsCognitoUserPools{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Cognito.DataAwsCognitoUserPoolsConfig",
		reflect.TypeOf((*DataAwsCognitoUserPoolsConfig)(nil)).Elem(),
	)
}
