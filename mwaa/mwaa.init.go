package mwaa

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironment",
		reflect.TypeOf((*MwaaEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "airflowConfigurationOptions", GoGetter: "AirflowConfigurationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "airflowConfigurationOptionsInput", GoGetter: "AirflowConfigurationOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "airflowVersion", GoGetter: "AirflowVersion"},
			_jsii_.MemberProperty{JsiiProperty: "airflowVersionInput", GoGetter: "AirflowVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "dagS3Path", GoGetter: "DagS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "dagS3PathInput", GoGetter: "DagS3PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "environmentClass", GoGetter: "EnvironmentClass"},
			_jsii_.MemberProperty{JsiiProperty: "environmentClassInput", GoGetter: "EnvironmentClassInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyInput", GoGetter: "KmsKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "lastUpdated", GoMethod: "LastUpdated"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfiguration", GoGetter: "LoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigurationInput", GoGetter: "LoggingConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxWorkers", GoGetter: "MaxWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "maxWorkersInput", GoGetter: "MaxWorkersInput"},
			_jsii_.MemberProperty{JsiiProperty: "minWorkers", GoGetter: "MinWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "minWorkersInput", GoGetter: "MinWorkersInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfiguration", GoGetter: "NetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "networkConfigurationInput", GoGetter: "NetworkConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3ObjectVersion", GoGetter: "PluginsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3ObjectVersionInput", GoGetter: "PluginsS3ObjectVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3Path", GoGetter: "PluginsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "pluginsS3PathInput", GoGetter: "PluginsS3PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfiguration", GoMethod: "PutLoggingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkConfiguration", GoMethod: "PutNetworkConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3ObjectVersion", GoGetter: "RequirementsS3ObjectVersion"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3ObjectVersionInput", GoGetter: "RequirementsS3ObjectVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3Path", GoGetter: "RequirementsS3Path"},
			_jsii_.MemberProperty{JsiiProperty: "requirementsS3PathInput", GoGetter: "RequirementsS3PathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAirflowConfigurationOptions", GoMethod: "ResetAirflowConfigurationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAirflowVersion", GoMethod: "ResetAirflowVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironmentClass", GoMethod: "ResetEnvironmentClass"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKey", GoMethod: "ResetKmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfiguration", GoMethod: "ResetLoggingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxWorkers", GoMethod: "ResetMaxWorkers"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinWorkers", GoMethod: "ResetMinWorkers"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPluginsS3ObjectVersion", GoMethod: "ResetPluginsS3ObjectVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetPluginsS3Path", GoMethod: "ResetPluginsS3Path"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequirementsS3ObjectVersion", GoMethod: "ResetRequirementsS3ObjectVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequirementsS3Path", GoMethod: "ResetRequirementsS3Path"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebserverAccessMode", GoMethod: "ResetWebserverAccessMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetWeeklyMaintenanceWindowStart", GoMethod: "ResetWeeklyMaintenanceWindowStart"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRoleArn", GoGetter: "ServiceRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceBucketArn", GoGetter: "SourceBucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "sourceBucketArnInput", GoGetter: "SourceBucketArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
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
			_jsii_.MemberProperty{JsiiProperty: "webserverAccessMode", GoGetter: "WebserverAccessMode"},
			_jsii_.MemberProperty{JsiiProperty: "webserverAccessModeInput", GoGetter: "WebserverAccessModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "webserverUrl", GoGetter: "WebserverUrl"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyMaintenanceWindowStart", GoGetter: "WeeklyMaintenanceWindowStart"},
			_jsii_.MemberProperty{JsiiProperty: "weeklyMaintenanceWindowStartInput", GoGetter: "WeeklyMaintenanceWindowStartInput"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.mwaa.MwaaEnvironmentConfig",
		reflect.TypeOf((*MwaaEnvironmentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironmentLastUpdated",
		reflect.TypeOf((*MwaaEnvironmentLastUpdated)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "error", GoGetter: "Error"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironmentLastUpdated{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironmentLastUpdatedError",
		reflect.TypeOf((*MwaaEnvironmentLastUpdatedError)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "errorCode", GoGetter: "ErrorCode"},
			_jsii_.MemberProperty{JsiiProperty: "errorMessage", GoGetter: "ErrorMessage"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironmentLastUpdatedError{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfiguration",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationDagProcessingLogs",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationDagProcessingLogs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "logLevelInput", GoGetter: "LogLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogLevel", GoMethod: "ResetLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironmentLoggingConfigurationDagProcessingLogsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationOutputReference",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dagProcessingLogs", GoGetter: "DagProcessingLogs"},
			_jsii_.MemberProperty{JsiiProperty: "dagProcessingLogsInput", GoGetter: "DagProcessingLogsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putDagProcessingLogs", GoMethod: "PutDagProcessingLogs"},
			_jsii_.MemberMethod{JsiiMethod: "putSchedulerLogs", GoMethod: "PutSchedulerLogs"},
			_jsii_.MemberMethod{JsiiMethod: "putTaskLogs", GoMethod: "PutTaskLogs"},
			_jsii_.MemberMethod{JsiiMethod: "putWebserverLogs", GoMethod: "PutWebserverLogs"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkerLogs", GoMethod: "PutWorkerLogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDagProcessingLogs", GoMethod: "ResetDagProcessingLogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchedulerLogs", GoMethod: "ResetSchedulerLogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetTaskLogs", GoMethod: "ResetTaskLogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebserverLogs", GoMethod: "ResetWebserverLogs"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkerLogs", GoMethod: "ResetWorkerLogs"},
			_jsii_.MemberProperty{JsiiProperty: "schedulerLogs", GoGetter: "SchedulerLogs"},
			_jsii_.MemberProperty{JsiiProperty: "schedulerLogsInput", GoGetter: "SchedulerLogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskLogs", GoGetter: "TaskLogs"},
			_jsii_.MemberProperty{JsiiProperty: "taskLogsInput", GoGetter: "TaskLogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "webserverLogs", GoGetter: "WebserverLogs"},
			_jsii_.MemberProperty{JsiiProperty: "webserverLogsInput", GoGetter: "WebserverLogsInput"},
			_jsii_.MemberProperty{JsiiProperty: "workerLogs", GoGetter: "WorkerLogs"},
			_jsii_.MemberProperty{JsiiProperty: "workerLogsInput", GoGetter: "WorkerLogsInput"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironmentLoggingConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationSchedulerLogs",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationSchedulerLogs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "logLevelInput", GoGetter: "LogLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogLevel", GoMethod: "ResetLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironmentLoggingConfigurationSchedulerLogsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationTaskLogs",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationTaskLogs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "logLevelInput", GoGetter: "LogLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogLevel", GoMethod: "ResetLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironmentLoggingConfigurationTaskLogsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationWebserverLogs",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationWebserverLogs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "logLevelInput", GoGetter: "LogLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogLevel", GoMethod: "ResetLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironmentLoggingConfigurationWebserverLogsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationWorkerLogs",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationWorkerLogs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference",
		reflect.TypeOf((*MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "logLevelInput", GoGetter: "LogLevelInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogLevel", GoMethod: "ResetLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironmentLoggingConfigurationWorkerLogsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.mwaa.MwaaEnvironmentNetworkConfiguration",
		reflect.TypeOf((*MwaaEnvironmentNetworkConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.mwaa.MwaaEnvironmentNetworkConfigurationOutputReference",
		reflect.TypeOf((*MwaaEnvironmentNetworkConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIds", GoGetter: "SubnetIds"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdsInput", GoGetter: "SubnetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_MwaaEnvironmentNetworkConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
