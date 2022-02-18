package securityhub

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubAccount",
		reflect.TypeOf((*SecurityhubAccount)(nil)).Elem(),
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
			j := jsiiProxy_SecurityhubAccount{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubAccountConfig",
		reflect.TypeOf((*SecurityhubAccountConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubActionTarget",
		reflect.TypeOf((*SecurityhubActionTarget)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubActionTarget{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubActionTargetConfig",
		reflect.TypeOf((*SecurityhubActionTargetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubFindingAggregator",
		reflect.TypeOf((*SecurityhubFindingAggregator)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "linkingMode", GoGetter: "LinkingMode"},
			_jsii_.MemberProperty{JsiiProperty: "linkingModeInput", GoGetter: "LinkingModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpecifiedRegions", GoMethod: "ResetSpecifiedRegions"},
			_jsii_.MemberProperty{JsiiProperty: "specifiedRegions", GoGetter: "SpecifiedRegions"},
			_jsii_.MemberProperty{JsiiProperty: "specifiedRegionsInput", GoGetter: "SpecifiedRegionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubFindingAggregator{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubFindingAggregatorConfig",
		reflect.TypeOf((*SecurityhubFindingAggregatorConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsight",
		reflect.TypeOf((*SecurityhubInsight)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "filters", GoGetter: "Filters"},
			_jsii_.MemberProperty{JsiiProperty: "filtersInput", GoGetter: "FiltersInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "groupByAttribute", GoGetter: "GroupByAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupByAttributeInput", GoGetter: "GroupByAttributeInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putFilters", GoMethod: "PutFilters"},
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
			j := jsiiProxy_SecurityhubInsight{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightConfig",
		reflect.TypeOf((*SecurityhubInsightConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFilters",
		reflect.TypeOf((*SecurityhubInsightFilters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersAwsAccountId",
		reflect.TypeOf((*SecurityhubInsightFiltersAwsAccountId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersCompanyName",
		reflect.TypeOf((*SecurityhubInsightFiltersCompanyName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersComplianceStatus",
		reflect.TypeOf((*SecurityhubInsightFiltersComplianceStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersConfidence",
		reflect.TypeOf((*SecurityhubInsightFiltersConfidence)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersCreatedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersCreatedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersCreatedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersCreatedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersCreatedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersCreatedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersCreatedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersCriticality",
		reflect.TypeOf((*SecurityhubInsightFiltersCriticality)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersDescription",
		reflect.TypeOf((*SecurityhubInsightFiltersDescription)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFindingProviderFieldsConfidence",
		reflect.TypeOf((*SecurityhubInsightFiltersFindingProviderFieldsConfidence)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFindingProviderFieldsCriticality",
		reflect.TypeOf((*SecurityhubInsightFiltersFindingProviderFieldsCriticality)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId",
		reflect.TypeOf((*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn",
		reflect.TypeOf((*SecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel",
		reflect.TypeOf((*SecurityhubInsightFiltersFindingProviderFieldsSeverityLabel)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal",
		reflect.TypeOf((*SecurityhubInsightFiltersFindingProviderFieldsSeverityOriginal)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFindingProviderFieldsTypes",
		reflect.TypeOf((*SecurityhubInsightFiltersFindingProviderFieldsTypes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFirstObservedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersFirstObservedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFirstObservedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersFirstObservedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersFirstObservedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersGeneratorId",
		reflect.TypeOf((*SecurityhubInsightFiltersGeneratorId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersId",
		reflect.TypeOf((*SecurityhubInsightFiltersId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersKeyword",
		reflect.TypeOf((*SecurityhubInsightFiltersKeyword)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersLastObservedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersLastObservedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersLastObservedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersLastObservedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersLastObservedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersMalwareName",
		reflect.TypeOf((*SecurityhubInsightFiltersMalwareName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersMalwarePath",
		reflect.TypeOf((*SecurityhubInsightFiltersMalwarePath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersMalwareState",
		reflect.TypeOf((*SecurityhubInsightFiltersMalwareState)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersMalwareType",
		reflect.TypeOf((*SecurityhubInsightFiltersMalwareType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkDestinationDomain",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkDestinationDomain)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkDestinationIpv4",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkDestinationIpv4)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkDestinationIpv6",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkDestinationIpv6)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkDestinationPort",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkDestinationPort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkDirection",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkDirection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkProtocol",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkProtocol)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkSourceDomain",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkSourceDomain)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkSourceIpv4",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkSourceIpv4)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkSourceIpv6",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkSourceIpv6)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkSourceMac",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkSourceMac)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNetworkSourcePort",
		reflect.TypeOf((*SecurityhubInsightFiltersNetworkSourcePort)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNoteText",
		reflect.TypeOf((*SecurityhubInsightFiltersNoteText)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNoteUpdatedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersNoteUpdatedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNoteUpdatedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersNoteUpdatedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersNoteUpdatedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersNoteUpdatedBy",
		reflect.TypeOf((*SecurityhubInsightFiltersNoteUpdatedBy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsAccountId", GoGetter: "AwsAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "awsAccountIdInput", GoGetter: "AwsAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "companyName", GoGetter: "CompanyName"},
			_jsii_.MemberProperty{JsiiProperty: "companyNameInput", GoGetter: "CompanyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complianceStatus", GoGetter: "ComplianceStatus"},
			_jsii_.MemberProperty{JsiiProperty: "complianceStatusInput", GoGetter: "ComplianceStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "confidence", GoGetter: "Confidence"},
			_jsii_.MemberProperty{JsiiProperty: "confidenceInput", GoGetter: "ConfidenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "createdAtInput", GoGetter: "CreatedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "criticality", GoGetter: "Criticality"},
			_jsii_.MemberProperty{JsiiProperty: "criticalityInput", GoGetter: "CriticalityInput"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsConfidence", GoGetter: "FindingProviderFieldsConfidence"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsConfidenceInput", GoGetter: "FindingProviderFieldsConfidenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsCriticality", GoGetter: "FindingProviderFieldsCriticality"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsCriticalityInput", GoGetter: "FindingProviderFieldsCriticalityInput"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsRelatedFindingsId", GoGetter: "FindingProviderFieldsRelatedFindingsId"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsRelatedFindingsIdInput", GoGetter: "FindingProviderFieldsRelatedFindingsIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsRelatedFindingsProductArn", GoGetter: "FindingProviderFieldsRelatedFindingsProductArn"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsRelatedFindingsProductArnInput", GoGetter: "FindingProviderFieldsRelatedFindingsProductArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsSeverityLabel", GoGetter: "FindingProviderFieldsSeverityLabel"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsSeverityLabelInput", GoGetter: "FindingProviderFieldsSeverityLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsSeverityOriginal", GoGetter: "FindingProviderFieldsSeverityOriginal"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsSeverityOriginalInput", GoGetter: "FindingProviderFieldsSeverityOriginalInput"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsTypes", GoGetter: "FindingProviderFieldsTypes"},
			_jsii_.MemberProperty{JsiiProperty: "findingProviderFieldsTypesInput", GoGetter: "FindingProviderFieldsTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "firstObservedAt", GoGetter: "FirstObservedAt"},
			_jsii_.MemberProperty{JsiiProperty: "firstObservedAtInput", GoGetter: "FirstObservedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "generatorId", GoGetter: "GeneratorId"},
			_jsii_.MemberProperty{JsiiProperty: "generatorIdInput", GoGetter: "GeneratorIdInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "keyword", GoGetter: "Keyword"},
			_jsii_.MemberProperty{JsiiProperty: "keywordInput", GoGetter: "KeywordInput"},
			_jsii_.MemberProperty{JsiiProperty: "lastObservedAt", GoGetter: "LastObservedAt"},
			_jsii_.MemberProperty{JsiiProperty: "lastObservedAtInput", GoGetter: "LastObservedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "malwareName", GoGetter: "MalwareName"},
			_jsii_.MemberProperty{JsiiProperty: "malwareNameInput", GoGetter: "MalwareNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "malwarePath", GoGetter: "MalwarePath"},
			_jsii_.MemberProperty{JsiiProperty: "malwarePathInput", GoGetter: "MalwarePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "malwareState", GoGetter: "MalwareState"},
			_jsii_.MemberProperty{JsiiProperty: "malwareStateInput", GoGetter: "MalwareStateInput"},
			_jsii_.MemberProperty{JsiiProperty: "malwareType", GoGetter: "MalwareType"},
			_jsii_.MemberProperty{JsiiProperty: "malwareTypeInput", GoGetter: "MalwareTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkDestinationDomain", GoGetter: "NetworkDestinationDomain"},
			_jsii_.MemberProperty{JsiiProperty: "networkDestinationDomainInput", GoGetter: "NetworkDestinationDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkDestinationIpv4", GoGetter: "NetworkDestinationIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "networkDestinationIpv4Input", GoGetter: "NetworkDestinationIpv4Input"},
			_jsii_.MemberProperty{JsiiProperty: "networkDestinationIpv6", GoGetter: "NetworkDestinationIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "networkDestinationIpv6Input", GoGetter: "NetworkDestinationIpv6Input"},
			_jsii_.MemberProperty{JsiiProperty: "networkDestinationPort", GoGetter: "NetworkDestinationPort"},
			_jsii_.MemberProperty{JsiiProperty: "networkDestinationPortInput", GoGetter: "NetworkDestinationPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkDirection", GoGetter: "NetworkDirection"},
			_jsii_.MemberProperty{JsiiProperty: "networkDirectionInput", GoGetter: "NetworkDirectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkProtocol", GoGetter: "NetworkProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "networkProtocolInput", GoGetter: "NetworkProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourceDomain", GoGetter: "NetworkSourceDomain"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourceDomainInput", GoGetter: "NetworkSourceDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourceIpv4", GoGetter: "NetworkSourceIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourceIpv4Input", GoGetter: "NetworkSourceIpv4Input"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourceIpv6", GoGetter: "NetworkSourceIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourceIpv6Input", GoGetter: "NetworkSourceIpv6Input"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourceMac", GoGetter: "NetworkSourceMac"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourceMacInput", GoGetter: "NetworkSourceMacInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourcePort", GoGetter: "NetworkSourcePort"},
			_jsii_.MemberProperty{JsiiProperty: "networkSourcePortInput", GoGetter: "NetworkSourcePortInput"},
			_jsii_.MemberProperty{JsiiProperty: "noteText", GoGetter: "NoteText"},
			_jsii_.MemberProperty{JsiiProperty: "noteTextInput", GoGetter: "NoteTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "noteUpdatedAt", GoGetter: "NoteUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "noteUpdatedAtInput", GoGetter: "NoteUpdatedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "noteUpdatedBy", GoGetter: "NoteUpdatedBy"},
			_jsii_.MemberProperty{JsiiProperty: "noteUpdatedByInput", GoGetter: "NoteUpdatedByInput"},
			_jsii_.MemberProperty{JsiiProperty: "processLaunchedAt", GoGetter: "ProcessLaunchedAt"},
			_jsii_.MemberProperty{JsiiProperty: "processLaunchedAtInput", GoGetter: "ProcessLaunchedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "processName", GoGetter: "ProcessName"},
			_jsii_.MemberProperty{JsiiProperty: "processNameInput", GoGetter: "ProcessNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "processParentPid", GoGetter: "ProcessParentPid"},
			_jsii_.MemberProperty{JsiiProperty: "processParentPidInput", GoGetter: "ProcessParentPidInput"},
			_jsii_.MemberProperty{JsiiProperty: "processPath", GoGetter: "ProcessPath"},
			_jsii_.MemberProperty{JsiiProperty: "processPathInput", GoGetter: "ProcessPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "processPid", GoGetter: "ProcessPid"},
			_jsii_.MemberProperty{JsiiProperty: "processPidInput", GoGetter: "ProcessPidInput"},
			_jsii_.MemberProperty{JsiiProperty: "processTerminatedAt", GoGetter: "ProcessTerminatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "processTerminatedAtInput", GoGetter: "ProcessTerminatedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "productArn", GoGetter: "ProductArn"},
			_jsii_.MemberProperty{JsiiProperty: "productArnInput", GoGetter: "ProductArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "productFields", GoGetter: "ProductFields"},
			_jsii_.MemberProperty{JsiiProperty: "productFieldsInput", GoGetter: "ProductFieldsInput"},
			_jsii_.MemberProperty{JsiiProperty: "productName", GoGetter: "ProductName"},
			_jsii_.MemberProperty{JsiiProperty: "productNameInput", GoGetter: "ProductNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "recommendationText", GoGetter: "RecommendationText"},
			_jsii_.MemberProperty{JsiiProperty: "recommendationTextInput", GoGetter: "RecommendationTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "recordState", GoGetter: "RecordState"},
			_jsii_.MemberProperty{JsiiProperty: "recordStateInput", GoGetter: "RecordStateInput"},
			_jsii_.MemberProperty{JsiiProperty: "relatedFindingsId", GoGetter: "RelatedFindingsId"},
			_jsii_.MemberProperty{JsiiProperty: "relatedFindingsIdInput", GoGetter: "RelatedFindingsIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "relatedFindingsProductArn", GoGetter: "RelatedFindingsProductArn"},
			_jsii_.MemberProperty{JsiiProperty: "relatedFindingsProductArnInput", GoGetter: "RelatedFindingsProductArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsAccountId", GoMethod: "ResetAwsAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompanyName", GoMethod: "ResetCompanyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetComplianceStatus", GoMethod: "ResetComplianceStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfidence", GoMethod: "ResetConfidence"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreatedAt", GoMethod: "ResetCreatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetCriticality", GoMethod: "ResetCriticality"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetFindingProviderFieldsConfidence", GoMethod: "ResetFindingProviderFieldsConfidence"},
			_jsii_.MemberMethod{JsiiMethod: "resetFindingProviderFieldsCriticality", GoMethod: "ResetFindingProviderFieldsCriticality"},
			_jsii_.MemberMethod{JsiiMethod: "resetFindingProviderFieldsRelatedFindingsId", GoMethod: "ResetFindingProviderFieldsRelatedFindingsId"},
			_jsii_.MemberMethod{JsiiMethod: "resetFindingProviderFieldsRelatedFindingsProductArn", GoMethod: "ResetFindingProviderFieldsRelatedFindingsProductArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetFindingProviderFieldsSeverityLabel", GoMethod: "ResetFindingProviderFieldsSeverityLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetFindingProviderFieldsSeverityOriginal", GoMethod: "ResetFindingProviderFieldsSeverityOriginal"},
			_jsii_.MemberMethod{JsiiMethod: "resetFindingProviderFieldsTypes", GoMethod: "ResetFindingProviderFieldsTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetFirstObservedAt", GoMethod: "ResetFirstObservedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetGeneratorId", GoMethod: "ResetGeneratorId"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyword", GoMethod: "ResetKeyword"},
			_jsii_.MemberMethod{JsiiMethod: "resetLastObservedAt", GoMethod: "ResetLastObservedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetMalwareName", GoMethod: "ResetMalwareName"},
			_jsii_.MemberMethod{JsiiMethod: "resetMalwarePath", GoMethod: "ResetMalwarePath"},
			_jsii_.MemberMethod{JsiiMethod: "resetMalwareState", GoMethod: "ResetMalwareState"},
			_jsii_.MemberMethod{JsiiMethod: "resetMalwareType", GoMethod: "ResetMalwareType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkDestinationDomain", GoMethod: "ResetNetworkDestinationDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkDestinationIpv4", GoMethod: "ResetNetworkDestinationIpv4"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkDestinationIpv6", GoMethod: "ResetNetworkDestinationIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkDestinationPort", GoMethod: "ResetNetworkDestinationPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkDirection", GoMethod: "ResetNetworkDirection"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkProtocol", GoMethod: "ResetNetworkProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkSourceDomain", GoMethod: "ResetNetworkSourceDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkSourceIpv4", GoMethod: "ResetNetworkSourceIpv4"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkSourceIpv6", GoMethod: "ResetNetworkSourceIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkSourceMac", GoMethod: "ResetNetworkSourceMac"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkSourcePort", GoMethod: "ResetNetworkSourcePort"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoteText", GoMethod: "ResetNoteText"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoteUpdatedAt", GoMethod: "ResetNoteUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoteUpdatedBy", GoMethod: "ResetNoteUpdatedBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessLaunchedAt", GoMethod: "ResetProcessLaunchedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessName", GoMethod: "ResetProcessName"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessParentPid", GoMethod: "ResetProcessParentPid"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessPath", GoMethod: "ResetProcessPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessPid", GoMethod: "ResetProcessPid"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessTerminatedAt", GoMethod: "ResetProcessTerminatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetProductArn", GoMethod: "ResetProductArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetProductFields", GoMethod: "ResetProductFields"},
			_jsii_.MemberMethod{JsiiMethod: "resetProductName", GoMethod: "ResetProductName"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecommendationText", GoMethod: "ResetRecommendationText"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecordState", GoMethod: "ResetRecordState"},
			_jsii_.MemberMethod{JsiiMethod: "resetRelatedFindingsId", GoMethod: "ResetRelatedFindingsId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRelatedFindingsProductArn", GoMethod: "ResetRelatedFindingsProductArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsEc2InstanceIamInstanceProfileArn", GoMethod: "ResetResourceAwsEc2InstanceIamInstanceProfileArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsEc2InstanceImageId", GoMethod: "ResetResourceAwsEc2InstanceImageId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsEc2InstanceIpv4Addresses", GoMethod: "ResetResourceAwsEc2InstanceIpv4Addresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsEc2InstanceIpv6Addresses", GoMethod: "ResetResourceAwsEc2InstanceIpv6Addresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsEc2InstanceKeyName", GoMethod: "ResetResourceAwsEc2InstanceKeyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsEc2InstanceLaunchedAt", GoMethod: "ResetResourceAwsEc2InstanceLaunchedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsEc2InstanceSubnetId", GoMethod: "ResetResourceAwsEc2InstanceSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsEc2InstanceType", GoMethod: "ResetResourceAwsEc2InstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsEc2InstanceVpcId", GoMethod: "ResetResourceAwsEc2InstanceVpcId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsIamAccessKeyCreatedAt", GoMethod: "ResetResourceAwsIamAccessKeyCreatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsIamAccessKeyStatus", GoMethod: "ResetResourceAwsIamAccessKeyStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsIamAccessKeyUserName", GoMethod: "ResetResourceAwsIamAccessKeyUserName"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsS3BucketOwnerId", GoMethod: "ResetResourceAwsS3BucketOwnerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceAwsS3BucketOwnerName", GoMethod: "ResetResourceAwsS3BucketOwnerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceContainerImageId", GoMethod: "ResetResourceContainerImageId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceContainerImageName", GoMethod: "ResetResourceContainerImageName"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceContainerLaunchedAt", GoMethod: "ResetResourceContainerLaunchedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceContainerName", GoMethod: "ResetResourceContainerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceDetailsOther", GoMethod: "ResetResourceDetailsOther"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceId", GoMethod: "ResetResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourcePartition", GoMethod: "ResetResourcePartition"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceRegion", GoMethod: "ResetResourceRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceTags", GoMethod: "ResetResourceTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceType", GoMethod: "ResetResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetSeverityLabel", GoMethod: "ResetSeverityLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceUrl", GoMethod: "ResetSourceUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetThreatIntelIndicatorCategory", GoMethod: "ResetThreatIntelIndicatorCategory"},
			_jsii_.MemberMethod{JsiiMethod: "resetThreatIntelIndicatorLastObservedAt", GoMethod: "ResetThreatIntelIndicatorLastObservedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetThreatIntelIndicatorSource", GoMethod: "ResetThreatIntelIndicatorSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetThreatIntelIndicatorSourceUrl", GoMethod: "ResetThreatIntelIndicatorSourceUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetThreatIntelIndicatorType", GoMethod: "ResetThreatIntelIndicatorType"},
			_jsii_.MemberMethod{JsiiMethod: "resetThreatIntelIndicatorValue", GoMethod: "ResetThreatIntelIndicatorValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetTitle", GoMethod: "ResetTitle"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdatedAt", GoMethod: "ResetUpdatedAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserDefinedValues", GoMethod: "ResetUserDefinedValues"},
			_jsii_.MemberMethod{JsiiMethod: "resetVerificationState", GoMethod: "ResetVerificationState"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkflowStatus", GoMethod: "ResetWorkflowStatus"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceIamInstanceProfileArn", GoGetter: "ResourceAwsEc2InstanceIamInstanceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceIamInstanceProfileArnInput", GoGetter: "ResourceAwsEc2InstanceIamInstanceProfileArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceImageId", GoGetter: "ResourceAwsEc2InstanceImageId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceImageIdInput", GoGetter: "ResourceAwsEc2InstanceImageIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceIpv4Addresses", GoGetter: "ResourceAwsEc2InstanceIpv4Addresses"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceIpv4AddressesInput", GoGetter: "ResourceAwsEc2InstanceIpv4AddressesInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceIpv6Addresses", GoGetter: "ResourceAwsEc2InstanceIpv6Addresses"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceIpv6AddressesInput", GoGetter: "ResourceAwsEc2InstanceIpv6AddressesInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceKeyName", GoGetter: "ResourceAwsEc2InstanceKeyName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceKeyNameInput", GoGetter: "ResourceAwsEc2InstanceKeyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceLaunchedAt", GoGetter: "ResourceAwsEc2InstanceLaunchedAt"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceLaunchedAtInput", GoGetter: "ResourceAwsEc2InstanceLaunchedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceSubnetId", GoGetter: "ResourceAwsEc2InstanceSubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceSubnetIdInput", GoGetter: "ResourceAwsEc2InstanceSubnetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceType", GoGetter: "ResourceAwsEc2InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceTypeInput", GoGetter: "ResourceAwsEc2InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceVpcId", GoGetter: "ResourceAwsEc2InstanceVpcId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsEc2InstanceVpcIdInput", GoGetter: "ResourceAwsEc2InstanceVpcIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsIamAccessKeyCreatedAt", GoGetter: "ResourceAwsIamAccessKeyCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsIamAccessKeyCreatedAtInput", GoGetter: "ResourceAwsIamAccessKeyCreatedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsIamAccessKeyStatus", GoGetter: "ResourceAwsIamAccessKeyStatus"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsIamAccessKeyStatusInput", GoGetter: "ResourceAwsIamAccessKeyStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsIamAccessKeyUserName", GoGetter: "ResourceAwsIamAccessKeyUserName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsIamAccessKeyUserNameInput", GoGetter: "ResourceAwsIamAccessKeyUserNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsS3BucketOwnerId", GoGetter: "ResourceAwsS3BucketOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsS3BucketOwnerIdInput", GoGetter: "ResourceAwsS3BucketOwnerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsS3BucketOwnerName", GoGetter: "ResourceAwsS3BucketOwnerName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceAwsS3BucketOwnerNameInput", GoGetter: "ResourceAwsS3BucketOwnerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceContainerImageId", GoGetter: "ResourceContainerImageId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceContainerImageIdInput", GoGetter: "ResourceContainerImageIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceContainerImageName", GoGetter: "ResourceContainerImageName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceContainerImageNameInput", GoGetter: "ResourceContainerImageNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceContainerLaunchedAt", GoGetter: "ResourceContainerLaunchedAt"},
			_jsii_.MemberProperty{JsiiProperty: "resourceContainerLaunchedAtInput", GoGetter: "ResourceContainerLaunchedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceContainerName", GoGetter: "ResourceContainerName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceContainerNameInput", GoGetter: "ResourceContainerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceDetailsOther", GoGetter: "ResourceDetailsOther"},
			_jsii_.MemberProperty{JsiiProperty: "resourceDetailsOtherInput", GoGetter: "ResourceDetailsOtherInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceId", GoGetter: "ResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceIdInput", GoGetter: "ResourceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePartition", GoGetter: "ResourcePartition"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePartitionInput", GoGetter: "ResourcePartitionInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceRegion", GoGetter: "ResourceRegion"},
			_jsii_.MemberProperty{JsiiProperty: "resourceRegionInput", GoGetter: "ResourceRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceTags", GoGetter: "ResourceTags"},
			_jsii_.MemberProperty{JsiiProperty: "resourceTagsInput", GoGetter: "ResourceTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "resourceTypeInput", GoGetter: "ResourceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "severityLabel", GoGetter: "SeverityLabel"},
			_jsii_.MemberProperty{JsiiProperty: "severityLabelInput", GoGetter: "SeverityLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceUrl", GoGetter: "SourceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "sourceUrlInput", GoGetter: "SourceUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorCategory", GoGetter: "ThreatIntelIndicatorCategory"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorCategoryInput", GoGetter: "ThreatIntelIndicatorCategoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorLastObservedAt", GoGetter: "ThreatIntelIndicatorLastObservedAt"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorLastObservedAtInput", GoGetter: "ThreatIntelIndicatorLastObservedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorSource", GoGetter: "ThreatIntelIndicatorSource"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorSourceInput", GoGetter: "ThreatIntelIndicatorSourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorSourceUrl", GoGetter: "ThreatIntelIndicatorSourceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorSourceUrlInput", GoGetter: "ThreatIntelIndicatorSourceUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorType", GoGetter: "ThreatIntelIndicatorType"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorTypeInput", GoGetter: "ThreatIntelIndicatorTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorValue", GoGetter: "ThreatIntelIndicatorValue"},
			_jsii_.MemberProperty{JsiiProperty: "threatIntelIndicatorValueInput", GoGetter: "ThreatIntelIndicatorValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberProperty{JsiiProperty: "titleInput", GoGetter: "TitleInput"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAtInput", GoGetter: "UpdatedAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "userDefinedValues", GoGetter: "UserDefinedValues"},
			_jsii_.MemberProperty{JsiiProperty: "userDefinedValuesInput", GoGetter: "UserDefinedValuesInput"},
			_jsii_.MemberProperty{JsiiProperty: "verificationState", GoGetter: "VerificationState"},
			_jsii_.MemberProperty{JsiiProperty: "verificationStateInput", GoGetter: "VerificationStateInput"},
			_jsii_.MemberProperty{JsiiProperty: "workflowStatus", GoGetter: "WorkflowStatus"},
			_jsii_.MemberProperty{JsiiProperty: "workflowStatusInput", GoGetter: "WorkflowStatusInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessLaunchedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessLaunchedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessLaunchedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessLaunchedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersProcessLaunchedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessName",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessParentPid",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessParentPid)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessPath",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessPid",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessPid)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessTerminatedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessTerminatedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessTerminatedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessTerminatedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersProcessTerminatedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProductArn",
		reflect.TypeOf((*SecurityhubInsightFiltersProductArn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProductFields",
		reflect.TypeOf((*SecurityhubInsightFiltersProductFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersProductName",
		reflect.TypeOf((*SecurityhubInsightFiltersProductName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersRecommendationText",
		reflect.TypeOf((*SecurityhubInsightFiltersRecommendationText)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersRecordState",
		reflect.TypeOf((*SecurityhubInsightFiltersRecordState)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersRelatedFindingsId",
		reflect.TypeOf((*SecurityhubInsightFiltersRelatedFindingsId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersRelatedFindingsProductArn",
		reflect.TypeOf((*SecurityhubInsightFiltersRelatedFindingsProductArn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceImageId",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceImageId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv4Addresses)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceIpv6Addresses)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceKeyName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceSubnetId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceType",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsEc2InstanceVpcId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsIamAccessKeyStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsIamAccessKeyUserName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsS3BucketOwnerId",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsS3BucketOwnerId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceAwsS3BucketOwnerName",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceAwsS3BucketOwnerName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceContainerImageId",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceContainerImageId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceContainerImageName",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceContainerImageName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceContainerLaunchedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceContainerLaunchedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceContainerLaunchedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersResourceContainerLaunchedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceContainerName",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceContainerName)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceDetailsOther",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceDetailsOther)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceId",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourcePartition",
		reflect.TypeOf((*SecurityhubInsightFiltersResourcePartition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceRegion",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceRegion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceTags",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceTags)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersResourceType",
		reflect.TypeOf((*SecurityhubInsightFiltersResourceType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersSeverityLabel",
		reflect.TypeOf((*SecurityhubInsightFiltersSeverityLabel)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersSourceUrl",
		reflect.TypeOf((*SecurityhubInsightFiltersSourceUrl)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorCategory",
		reflect.TypeOf((*SecurityhubInsightFiltersThreatIntelIndicatorCategory)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorSource",
		reflect.TypeOf((*SecurityhubInsightFiltersThreatIntelIndicatorSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl",
		reflect.TypeOf((*SecurityhubInsightFiltersThreatIntelIndicatorSourceUrl)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorType",
		reflect.TypeOf((*SecurityhubInsightFiltersThreatIntelIndicatorType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersThreatIntelIndicatorValue",
		reflect.TypeOf((*SecurityhubInsightFiltersThreatIntelIndicatorValue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersTitle",
		reflect.TypeOf((*SecurityhubInsightFiltersTitle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersType",
		reflect.TypeOf((*SecurityhubInsightFiltersType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersUpdatedAt",
		reflect.TypeOf((*SecurityhubInsightFiltersUpdatedAt)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersUpdatedAtDateRange",
		reflect.TypeOf((*SecurityhubInsightFiltersUpdatedAtDateRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference",
		reflect.TypeOf((*SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubInsightFiltersUpdatedAtDateRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersUserDefinedValues",
		reflect.TypeOf((*SecurityhubInsightFiltersUserDefinedValues)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersVerificationState",
		reflect.TypeOf((*SecurityhubInsightFiltersVerificationState)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInsightFiltersWorkflowStatus",
		reflect.TypeOf((*SecurityhubInsightFiltersWorkflowStatus)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubInviteAccepter",
		reflect.TypeOf((*SecurityhubInviteAccepter)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "invitationId", GoGetter: "InvitationId"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "masterId", GoGetter: "MasterId"},
			_jsii_.MemberProperty{JsiiProperty: "masterIdInput", GoGetter: "MasterIdInput"},
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
			j := jsiiProxy_SecurityhubInviteAccepter{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubInviteAccepterConfig",
		reflect.TypeOf((*SecurityhubInviteAccepterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubMember",
		reflect.TypeOf((*SecurityhubMember)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "invite", GoGetter: "Invite"},
			_jsii_.MemberProperty{JsiiProperty: "inviteInput", GoGetter: "InviteInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "masterId", GoGetter: "MasterId"},
			_jsii_.MemberProperty{JsiiProperty: "memberStatus", GoGetter: "MemberStatus"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetInvite", GoMethod: "ResetInvite"},
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
			j := jsiiProxy_SecurityhubMember{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubMemberConfig",
		reflect.TypeOf((*SecurityhubMemberConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubOrganizationAdminAccount",
		reflect.TypeOf((*SecurityhubOrganizationAdminAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "adminAccountId", GoGetter: "AdminAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "adminAccountIdInput", GoGetter: "AdminAccountIdInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubOrganizationAdminAccount{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubOrganizationAdminAccountConfig",
		reflect.TypeOf((*SecurityhubOrganizationAdminAccountConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubOrganizationConfiguration",
		reflect.TypeOf((*SecurityhubOrganizationConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoEnable", GoGetter: "AutoEnable"},
			_jsii_.MemberProperty{JsiiProperty: "autoEnableInput", GoGetter: "AutoEnableInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubOrganizationConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubOrganizationConfigurationConfig",
		reflect.TypeOf((*SecurityhubOrganizationConfigurationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubProductSubscription",
		reflect.TypeOf((*SecurityhubProductSubscription)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "productArn", GoGetter: "ProductArn"},
			_jsii_.MemberProperty{JsiiProperty: "productArnInput", GoGetter: "ProductArnInput"},
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
			j := jsiiProxy_SecurityhubProductSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubProductSubscriptionConfig",
		reflect.TypeOf((*SecurityhubProductSubscriptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubStandardsControl",
		reflect.TypeOf((*SecurityhubStandardsControl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "controlId", GoGetter: "ControlId"},
			_jsii_.MemberProperty{JsiiProperty: "controlStatus", GoGetter: "ControlStatus"},
			_jsii_.MemberProperty{JsiiProperty: "controlStatusInput", GoGetter: "ControlStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "controlStatusUpdatedAt", GoGetter: "ControlStatusUpdatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "disabledReason", GoGetter: "DisabledReason"},
			_jsii_.MemberProperty{JsiiProperty: "disabledReasonInput", GoGetter: "DisabledReasonInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "relatedRequirements", GoGetter: "RelatedRequirements"},
			_jsii_.MemberProperty{JsiiProperty: "remediationUrl", GoGetter: "RemediationUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisabledReason", GoMethod: "ResetDisabledReason"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "severityRating", GoGetter: "SeverityRating"},
			_jsii_.MemberProperty{JsiiProperty: "standardsControlArn", GoGetter: "StandardsControlArn"},
			_jsii_.MemberProperty{JsiiProperty: "standardsControlArnInput", GoGetter: "StandardsControlArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubStandardsControl{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubStandardsControlConfig",
		reflect.TypeOf((*SecurityhubStandardsControlConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.securityhub.SecurityhubStandardsSubscription",
		reflect.TypeOf((*SecurityhubStandardsSubscription)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "standardsArn", GoGetter: "StandardsArn"},
			_jsii_.MemberProperty{JsiiProperty: "standardsArnInput", GoGetter: "StandardsArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityhubStandardsSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.securityhub.SecurityhubStandardsSubscriptionConfig",
		reflect.TypeOf((*SecurityhubStandardsSubscriptionConfig)(nil)).Elem(),
	)
}
