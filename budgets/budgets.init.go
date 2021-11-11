package budgets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.Budgets.BudgetsBudget",
		reflect.TypeOf((*BudgetsBudget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "budgetType", GoGetter: "BudgetType"},
			_jsii_.MemberProperty{JsiiProperty: "budgetTypeInput", GoGetter: "BudgetTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "costFilter", GoGetter: "CostFilter"},
			_jsii_.MemberProperty{JsiiProperty: "costFilterInput", GoGetter: "CostFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "costFilters", GoGetter: "CostFilters"},
			_jsii_.MemberProperty{JsiiProperty: "costFiltersInput", GoGetter: "CostFiltersInput"},
			_jsii_.MemberProperty{JsiiProperty: "costTypes", GoGetter: "CostTypes"},
			_jsii_.MemberProperty{JsiiProperty: "costTypesInput", GoGetter: "CostTypesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "limitAmount", GoGetter: "LimitAmount"},
			_jsii_.MemberProperty{JsiiProperty: "limitAmountInput", GoGetter: "LimitAmountInput"},
			_jsii_.MemberProperty{JsiiProperty: "limitUnit", GoGetter: "LimitUnit"},
			_jsii_.MemberProperty{JsiiProperty: "limitUnitInput", GoGetter: "LimitUnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefix", GoGetter: "NamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefixInput", GoGetter: "NamePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notification", GoGetter: "Notification"},
			_jsii_.MemberProperty{JsiiProperty: "notificationInput", GoGetter: "NotificationInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putCostTypes", GoMethod: "PutCostTypes"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCostFilter", GoMethod: "ResetCostFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetCostFilters", GoMethod: "ResetCostFilters"},
			_jsii_.MemberMethod{JsiiMethod: "resetCostTypes", GoMethod: "ResetCostTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamePrefix", GoMethod: "ResetNamePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotification", GoMethod: "ResetNotification"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimePeriodEnd", GoMethod: "ResetTimePeriodEnd"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimePeriodStart", GoMethod: "ResetTimePeriodStart"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriodEnd", GoGetter: "TimePeriodEnd"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriodEndInput", GoGetter: "TimePeriodEndInput"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriodStart", GoGetter: "TimePeriodStart"},
			_jsii_.MemberProperty{JsiiProperty: "timePeriodStartInput", GoGetter: "TimePeriodStartInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeUnit", GoGetter: "TimeUnit"},
			_jsii_.MemberProperty{JsiiProperty: "timeUnitInput", GoGetter: "TimeUnitInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudget{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Budgets.BudgetsBudgetAction",
		reflect.TypeOf((*BudgetsBudgetAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "actionId", GoGetter: "ActionId"},
			_jsii_.MemberProperty{JsiiProperty: "actionThreshold", GoGetter: "ActionThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "actionThresholdInput", GoGetter: "ActionThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "actionType", GoGetter: "ActionType"},
			_jsii_.MemberProperty{JsiiProperty: "actionTypeInput", GoGetter: "ActionTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "approvalModel", GoGetter: "ApprovalModel"},
			_jsii_.MemberProperty{JsiiProperty: "approvalModelInput", GoGetter: "ApprovalModelInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "budgetName", GoGetter: "BudgetName"},
			_jsii_.MemberProperty{JsiiProperty: "budgetNameInput", GoGetter: "BudgetNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "definitionInput", GoGetter: "DefinitionInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArnInput", GoGetter: "ExecutionRoleArnInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "notificationType", GoGetter: "NotificationType"},
			_jsii_.MemberProperty{JsiiProperty: "notificationTypeInput", GoGetter: "NotificationTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putActionThreshold", GoMethod: "PutActionThreshold"},
			_jsii_.MemberMethod{JsiiMethod: "putDefinition", GoMethod: "PutDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "subscriber", GoGetter: "Subscriber"},
			_jsii_.MemberProperty{JsiiProperty: "subscriberInput", GoGetter: "SubscriberInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetAction{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetActionActionThreshold",
		reflect.TypeOf((*BudgetsBudgetActionActionThreshold)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Budgets.BudgetsBudgetActionActionThresholdOutputReference",
		reflect.TypeOf((*BudgetsBudgetActionActionThresholdOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionThresholdType", GoGetter: "ActionThresholdType"},
			_jsii_.MemberProperty{JsiiProperty: "actionThresholdTypeInput", GoGetter: "ActionThresholdTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "actionThresholdValue", GoGetter: "ActionThresholdValue"},
			_jsii_.MemberProperty{JsiiProperty: "actionThresholdValueInput", GoGetter: "ActionThresholdValueInput"},
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
			j := jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetActionConfig",
		reflect.TypeOf((*BudgetsBudgetActionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinition",
		reflect.TypeOf((*BudgetsBudgetActionDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionIamActionDefinition",
		reflect.TypeOf((*BudgetsBudgetActionDefinitionIamActionDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference",
		reflect.TypeOf((*BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groups", GoGetter: "Groups"},
			_jsii_.MemberProperty{JsiiProperty: "groupsInput", GoGetter: "GroupsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "policyArn", GoGetter: "PolicyArn"},
			_jsii_.MemberProperty{JsiiProperty: "policyArnInput", GoGetter: "PolicyArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroups", GoMethod: "ResetGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetRoles", GoMethod: "ResetRoles"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsers", GoMethod: "ResetUsers"},
			_jsii_.MemberProperty{JsiiProperty: "roles", GoGetter: "Roles"},
			_jsii_.MemberProperty{JsiiProperty: "rolesInput", GoGetter: "RolesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
			_jsii_.MemberProperty{JsiiProperty: "usersInput", GoGetter: "UsersInput"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionOutputReference",
		reflect.TypeOf((*BudgetsBudgetActionDefinitionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamActionDefinition", GoGetter: "IamActionDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "iamActionDefinitionInput", GoGetter: "IamActionDefinitionInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putIamActionDefinition", GoMethod: "PutIamActionDefinition"},
			_jsii_.MemberMethod{JsiiMethod: "putScpActionDefinition", GoMethod: "PutScpActionDefinition"},
			_jsii_.MemberMethod{JsiiMethod: "putSsmActionDefinition", GoMethod: "PutSsmActionDefinition"},
			_jsii_.MemberMethod{JsiiMethod: "resetIamActionDefinition", GoMethod: "ResetIamActionDefinition"},
			_jsii_.MemberMethod{JsiiMethod: "resetScpActionDefinition", GoMethod: "ResetScpActionDefinition"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsmActionDefinition", GoMethod: "ResetSsmActionDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "scpActionDefinition", GoGetter: "ScpActionDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "scpActionDefinitionInput", GoGetter: "ScpActionDefinitionInput"},
			_jsii_.MemberProperty{JsiiProperty: "ssmActionDefinition", GoGetter: "SsmActionDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "ssmActionDefinitionInput", GoGetter: "SsmActionDefinitionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetActionDefinitionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionScpActionDefinition",
		reflect.TypeOf((*BudgetsBudgetActionDefinitionScpActionDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference",
		reflect.TypeOf((*BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdInput", GoGetter: "PolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetIds", GoGetter: "TargetIds"},
			_jsii_.MemberProperty{JsiiProperty: "targetIdsInput", GoGetter: "TargetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionSsmActionDefinition",
		reflect.TypeOf((*BudgetsBudgetActionDefinitionSsmActionDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference",
		reflect.TypeOf((*BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionSubType", GoGetter: "ActionSubType"},
			_jsii_.MemberProperty{JsiiProperty: "actionSubTypeInput", GoGetter: "ActionSubTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIds", GoGetter: "InstanceIds"},
			_jsii_.MemberProperty{JsiiProperty: "instanceIdsInput", GoGetter: "InstanceIdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetActionSubscriber",
		reflect.TypeOf((*BudgetsBudgetActionSubscriber)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetConfig",
		reflect.TypeOf((*BudgetsBudgetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetCostFilter",
		reflect.TypeOf((*BudgetsBudgetCostFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetCostTypes",
		reflect.TypeOf((*BudgetsBudgetCostTypes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.Budgets.BudgetsBudgetCostTypesOutputReference",
		reflect.TypeOf((*BudgetsBudgetCostTypesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeCredit", GoGetter: "IncludeCredit"},
			_jsii_.MemberProperty{JsiiProperty: "includeCreditInput", GoGetter: "IncludeCreditInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeDiscount", GoGetter: "IncludeDiscount"},
			_jsii_.MemberProperty{JsiiProperty: "includeDiscountInput", GoGetter: "IncludeDiscountInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeOtherSubscription", GoGetter: "IncludeOtherSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "includeOtherSubscriptionInput", GoGetter: "IncludeOtherSubscriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeRecurring", GoGetter: "IncludeRecurring"},
			_jsii_.MemberProperty{JsiiProperty: "includeRecurringInput", GoGetter: "IncludeRecurringInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeRefund", GoGetter: "IncludeRefund"},
			_jsii_.MemberProperty{JsiiProperty: "includeRefundInput", GoGetter: "IncludeRefundInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeSubscription", GoGetter: "IncludeSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "includeSubscriptionInput", GoGetter: "IncludeSubscriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeSupport", GoGetter: "IncludeSupport"},
			_jsii_.MemberProperty{JsiiProperty: "includeSupportInput", GoGetter: "IncludeSupportInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeTax", GoGetter: "IncludeTax"},
			_jsii_.MemberProperty{JsiiProperty: "includeTaxInput", GoGetter: "IncludeTaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "includeUpfront", GoGetter: "IncludeUpfront"},
			_jsii_.MemberProperty{JsiiProperty: "includeUpfrontInput", GoGetter: "IncludeUpfrontInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeCredit", GoMethod: "ResetIncludeCredit"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeDiscount", GoMethod: "ResetIncludeDiscount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeOtherSubscription", GoMethod: "ResetIncludeOtherSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeRecurring", GoMethod: "ResetIncludeRecurring"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeRefund", GoMethod: "ResetIncludeRefund"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeSubscription", GoMethod: "ResetIncludeSubscription"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeSupport", GoMethod: "ResetIncludeSupport"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeTax", GoMethod: "ResetIncludeTax"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeUpfront", GoMethod: "ResetIncludeUpfront"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseAmortized", GoMethod: "ResetUseAmortized"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseBlended", GoMethod: "ResetUseBlended"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "useAmortized", GoGetter: "UseAmortized"},
			_jsii_.MemberProperty{JsiiProperty: "useAmortizedInput", GoGetter: "UseAmortizedInput"},
			_jsii_.MemberProperty{JsiiProperty: "useBlended", GoGetter: "UseBlended"},
			_jsii_.MemberProperty{JsiiProperty: "useBlendedInput", GoGetter: "UseBlendedInput"},
		},
		func() interface{} {
			j := jsiiProxy_BudgetsBudgetCostTypesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.Budgets.BudgetsBudgetNotification",
		reflect.TypeOf((*BudgetsBudgetNotification)(nil)).Elem(),
	)
}
