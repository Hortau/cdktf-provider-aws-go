package appautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/appautoscaling/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy aws_appautoscaling_policy}.
type AppautoscalingPolicy interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ScalableDimension() *string
	SetScalableDimension(val *string)
	ScalableDimensionInput() *string
	ServiceNamespace() *string
	SetServiceNamespace(val *string)
	ServiceNamespaceInput() *string
	StepScalingPolicyConfiguration() AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference
	StepScalingPolicyConfigurationInput() *AppautoscalingPolicyStepScalingPolicyConfiguration
	TargetTrackingScalingPolicyConfiguration() AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference
	TargetTrackingScalingPolicyConfigurationInput() *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutStepScalingPolicyConfiguration(value *AppautoscalingPolicyStepScalingPolicyConfiguration)
	PutTargetTrackingScalingPolicyConfiguration(value *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration)
	ResetOverrideLogicalId()
	ResetPolicyType()
	ResetStepScalingPolicyConfiguration()
	ResetTargetTrackingScalingPolicyConfiguration()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppautoscalingPolicy
type jsiiProxy_AppautoscalingPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppautoscalingPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) ScalableDimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) ServiceNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) StepScalingPolicyConfiguration() AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference {
	var returns AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference
	_jsii_.Get(
		j,
		"stepScalingPolicyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) StepScalingPolicyConfigurationInput() *AppautoscalingPolicyStepScalingPolicyConfiguration {
	var returns *AppautoscalingPolicyStepScalingPolicyConfiguration
	_jsii_.Get(
		j,
		"stepScalingPolicyConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) TargetTrackingScalingPolicyConfiguration() AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference {
	var returns AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference
	_jsii_.Get(
		j,
		"targetTrackingScalingPolicyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) TargetTrackingScalingPolicyConfigurationInput() *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration {
	var returns *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration
	_jsii_.Get(
		j,
		"targetTrackingScalingPolicyConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy aws_appautoscaling_policy} Resource.
func NewAppautoscalingPolicy(scope constructs.Construct, id *string, config *AppautoscalingPolicyConfig) AppautoscalingPolicy {
	_init_.Initialize()

	j := jsiiProxy_AppautoscalingPolicy{}

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy aws_appautoscaling_policy} Resource.
func NewAppautoscalingPolicy_Override(a AppautoscalingPolicy, scope constructs.Construct, id *string, config *AppautoscalingPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingPolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicy) SetPolicyType(val *string) {
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicy) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicy) SetScalableDimension(val *string) {
	_jsii_.Set(
		j,
		"scalableDimension",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicy) SetServiceNamespace(val *string) {
	_jsii_.Set(
		j,
		"serviceNamespace",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AppautoscalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppautoscalingPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppautoscalingPolicy) PutStepScalingPolicyConfiguration(value *AppautoscalingPolicyStepScalingPolicyConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putStepScalingPolicyConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicy) PutTargetTrackingScalingPolicyConfiguration(value *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putTargetTrackingScalingPolicyConfiguration",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicy) ResetPolicyType() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicy) ResetStepScalingPolicyConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetStepScalingPolicyConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicy) ResetTargetTrackingScalingPolicyConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetTrackingScalingPolicyConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppautoscalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppautoscalingPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS AppAutoScaling.
type AppautoscalingPolicyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#name AppautoscalingPolicy#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#resource_id AppautoscalingPolicy#resource_id}.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#scalable_dimension AppautoscalingPolicy#scalable_dimension}.
	ScalableDimension *string `json:"scalableDimension" yaml:"scalableDimension"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#service_namespace AppautoscalingPolicy#service_namespace}.
	ServiceNamespace *string `json:"serviceNamespace" yaml:"serviceNamespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#policy_type AppautoscalingPolicy#policy_type}.
	PolicyType *string `json:"policyType" yaml:"policyType"`
	// step_scaling_policy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#step_scaling_policy_configuration AppautoscalingPolicy#step_scaling_policy_configuration}
	StepScalingPolicyConfiguration *AppautoscalingPolicyStepScalingPolicyConfiguration `json:"stepScalingPolicyConfiguration" yaml:"stepScalingPolicyConfiguration"`
	// target_tracking_scaling_policy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#target_tracking_scaling_policy_configuration AppautoscalingPolicy#target_tracking_scaling_policy_configuration}
	TargetTrackingScalingPolicyConfiguration *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration `json:"targetTrackingScalingPolicyConfiguration" yaml:"targetTrackingScalingPolicyConfiguration"`
}

type AppautoscalingPolicyStepScalingPolicyConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#adjustment_type AppautoscalingPolicy#adjustment_type}.
	AdjustmentType *string `json:"adjustmentType" yaml:"adjustmentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#cooldown AppautoscalingPolicy#cooldown}.
	Cooldown *float64 `json:"cooldown" yaml:"cooldown"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#metric_aggregation_type AppautoscalingPolicy#metric_aggregation_type}.
	MetricAggregationType *string `json:"metricAggregationType" yaml:"metricAggregationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#min_adjustment_magnitude AppautoscalingPolicy#min_adjustment_magnitude}.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// step_adjustment block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#step_adjustment AppautoscalingPolicy#step_adjustment}
	StepAdjustment interface{} `json:"stepAdjustment" yaml:"stepAdjustment"`
}

type AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference interface {
	cdktf.ComplexObject
	AdjustmentType() *string
	SetAdjustmentType(val *string)
	AdjustmentTypeInput() *string
	Cooldown() *float64
	SetCooldown(val *float64)
	CooldownInput() *float64
	InternalValue() *AppautoscalingPolicyStepScalingPolicyConfiguration
	SetInternalValue(val *AppautoscalingPolicyStepScalingPolicyConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MetricAggregationType() *string
	SetMetricAggregationType(val *string)
	MetricAggregationTypeInput() *string
	MinAdjustmentMagnitude() *float64
	SetMinAdjustmentMagnitude(val *float64)
	MinAdjustmentMagnitudeInput() *float64
	StepAdjustment() interface{}
	SetStepAdjustment(val interface{})
	StepAdjustmentInput() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAdjustmentType()
	ResetCooldown()
	ResetMetricAggregationType()
	ResetMinAdjustmentMagnitude()
	ResetStepAdjustment()
}

// The jsii proxy struct for AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference
type jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) AdjustmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) AdjustmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) Cooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) CooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) InternalValue() *AppautoscalingPolicyStepScalingPolicyConfiguration {
	var returns *AppautoscalingPolicyStepScalingPolicyConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) MetricAggregationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricAggregationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) MetricAggregationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricAggregationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) MinAdjustmentMagnitude() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAdjustmentMagnitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) MinAdjustmentMagnitudeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAdjustmentMagnitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) StepAdjustment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepAdjustment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) StepAdjustmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepAdjustmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppautoscalingPolicyStepScalingPolicyConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppautoscalingPolicyStepScalingPolicyConfigurationOutputReference_Override(a AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) SetAdjustmentType(val *string) {
	_jsii_.Set(
		j,
		"adjustmentType",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) SetCooldown(val *float64) {
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) SetInternalValue(val *AppautoscalingPolicyStepScalingPolicyConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) SetMetricAggregationType(val *string) {
	_jsii_.Set(
		j,
		"metricAggregationType",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) SetMinAdjustmentMagnitude(val *float64) {
	_jsii_.Set(
		j,
		"minAdjustmentMagnitude",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) SetStepAdjustment(val interface{}) {
	_jsii_.Set(
		j,
		"stepAdjustment",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) ResetAdjustmentType() {
	_jsii_.InvokeVoid(
		a,
		"resetAdjustmentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) ResetCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) ResetMetricAggregationType() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricAggregationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) ResetMinAdjustmentMagnitude() {
	_jsii_.InvokeVoid(
		a,
		"resetMinAdjustmentMagnitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyStepScalingPolicyConfigurationOutputReference) ResetStepAdjustment() {
	_jsii_.InvokeVoid(
		a,
		"resetStepAdjustment",
		nil, // no parameters
	)
}

type AppautoscalingPolicyStepScalingPolicyConfigurationStepAdjustment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#scaling_adjustment AppautoscalingPolicy#scaling_adjustment}.
	ScalingAdjustment *float64 `json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#metric_interval_lower_bound AppautoscalingPolicy#metric_interval_lower_bound}.
	MetricIntervalLowerBound *string `json:"metricIntervalLowerBound" yaml:"metricIntervalLowerBound"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#metric_interval_upper_bound AppautoscalingPolicy#metric_interval_upper_bound}.
	MetricIntervalUpperBound *string `json:"metricIntervalUpperBound" yaml:"metricIntervalUpperBound"`
}

type AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#target_value AppautoscalingPolicy#target_value}.
	TargetValue *float64 `json:"targetValue" yaml:"targetValue"`
	// customized_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#customized_metric_specification AppautoscalingPolicy#customized_metric_specification}
	CustomizedMetricSpecification *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification `json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#disable_scale_in AppautoscalingPolicy#disable_scale_in}.
	DisableScaleIn interface{} `json:"disableScaleIn" yaml:"disableScaleIn"`
	// predefined_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#predefined_metric_specification AppautoscalingPolicy#predefined_metric_specification}
	PredefinedMetricSpecification *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification `json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#scale_in_cooldown AppautoscalingPolicy#scale_in_cooldown}.
	ScaleInCooldown *float64 `json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#scale_out_cooldown AppautoscalingPolicy#scale_out_cooldown}.
	ScaleOutCooldown *float64 `json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#metric_name AppautoscalingPolicy#metric_name}.
	MetricName *string `json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#namespace AppautoscalingPolicy#namespace}.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#statistic AppautoscalingPolicy#statistic}.
	Statistic *string `json:"statistic" yaml:"statistic"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#dimensions AppautoscalingPolicy#dimensions}
	Dimensions interface{} `json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#unit AppautoscalingPolicy#unit}.
	Unit *string `json:"unit" yaml:"unit"`
}

type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#name AppautoscalingPolicy#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#value AppautoscalingPolicy#value}.
	Value *string `json:"value" yaml:"value"`
}

type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference interface {
	cdktf.ComplexObject
	Dimensions() interface{}
	SetDimensions(val interface{})
	DimensionsInput() interface{}
	InternalValue() *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification
	SetInternalValue(val *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MetricName() *string
	SetMetricName(val *string)
	MetricNameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	Statistic() *string
	SetStatistic(val *string)
	StatisticInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDimensions()
	ResetUnit()
}

// The jsii proxy struct for AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference
type jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) Dimensions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) DimensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) InternalValue() *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification {
	var returns *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) StatisticInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statisticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func NewAppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference_Override(a AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) SetDimensions(val interface{}) {
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) SetInternalValue(val *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) SetMetricName(val *string) {
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) SetStatistic(val *string) {
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		a,
		"resetDimensions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference) ResetUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetUnit",
		nil, // no parameters
	)
}

type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference interface {
	cdktf.ComplexObject
	CustomizedMetricSpecification() AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference
	CustomizedMetricSpecificationInput() *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification
	DisableScaleIn() interface{}
	SetDisableScaleIn(val interface{})
	DisableScaleInInput() interface{}
	InternalValue() *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration
	SetInternalValue(val *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	PredefinedMetricSpecification() AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference
	PredefinedMetricSpecificationInput() *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification
	ScaleInCooldown() *float64
	SetScaleInCooldown(val *float64)
	ScaleInCooldownInput() *float64
	ScaleOutCooldown() *float64
	SetScaleOutCooldown(val *float64)
	ScaleOutCooldownInput() *float64
	TargetValue() *float64
	SetTargetValue(val *float64)
	TargetValueInput() *float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutCustomizedMetricSpecification(value *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification)
	PutPredefinedMetricSpecification(value *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification)
	ResetCustomizedMetricSpecification()
	ResetDisableScaleIn()
	ResetPredefinedMetricSpecification()
	ResetScaleInCooldown()
	ResetScaleOutCooldown()
}

// The jsii proxy struct for AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference
type jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) CustomizedMetricSpecification() AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference {
	var returns AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"customizedMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) CustomizedMetricSpecificationInput() *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification {
	var returns *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification
	_jsii_.Get(
		j,
		"customizedMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) DisableScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) DisableScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) InternalValue() *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration {
	var returns *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) PredefinedMetricSpecification() AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference {
	var returns AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference
	_jsii_.Get(
		j,
		"predefinedMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) PredefinedMetricSpecificationInput() *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification {
	var returns *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification
	_jsii_.Get(
		j,
		"predefinedMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ScaleInCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ScaleInCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ScaleOutCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ScaleOutCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) TargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference_Override(a AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) SetDisableScaleIn(val interface{}) {
	_jsii_.Set(
		j,
		"disableScaleIn",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) SetInternalValue(val *AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) SetScaleInCooldown(val *float64) {
	_jsii_.Set(
		j,
		"scaleInCooldown",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) SetScaleOutCooldown(val *float64) {
	_jsii_.Set(
		j,
		"scaleOutCooldown",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) SetTargetValue(val *float64) {
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) PutCustomizedMetricSpecification(value *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) {
	_jsii_.InvokeVoid(
		a,
		"putCustomizedMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) PutPredefinedMetricSpecification(value *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification) {
	_jsii_.InvokeVoid(
		a,
		"putPredefinedMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ResetCustomizedMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ResetDisableScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ResetPredefinedMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ResetScaleInCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleInCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationOutputReference) ResetScaleOutCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleOutCooldown",
		nil, // no parameters
	)
}

type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#predefined_metric_type AppautoscalingPolicy#predefined_metric_type}.
	PredefinedMetricType *string `json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy#resource_label AppautoscalingPolicy#resource_label}.
	ResourceLabel *string `json:"resourceLabel" yaml:"resourceLabel"`
}

type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification
	SetInternalValue(val *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	PredefinedMetricType() *string
	SetPredefinedMetricType(val *string)
	PredefinedMetricTypeInput() *string
	ResourceLabel() *string
	SetResourceLabel(val *string)
	ResourceLabelInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetResourceLabel()
}

// The jsii proxy struct for AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference
type jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) InternalValue() *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification {
	var returns *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) PredefinedMetricType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedMetricType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) PredefinedMetricTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedMetricTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) ResourceLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) ResourceLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference_Override(a AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) SetInternalValue(val *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) SetPredefinedMetricType(val *string) {
	_jsii_.Set(
		j,
		"predefinedMetricType",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) SetResourceLabel(val *string) {
	_jsii_.Set(
		j,
		"resourceLabel",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationOutputReference) ResetResourceLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceLabel",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action aws_appautoscaling_scheduled_action}.
type AppautoscalingScheduledAction interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ScalableDimension() *string
	SetScalableDimension(val *string)
	ScalableDimensionInput() *string
	ScalableTargetAction() AppautoscalingScheduledActionScalableTargetActionOutputReference
	ScalableTargetActionInput() *AppautoscalingScheduledActionScalableTargetAction
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	ServiceNamespace() *string
	SetServiceNamespace(val *string)
	ServiceNamespaceInput() *string
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutScalableTargetAction(value *AppautoscalingScheduledActionScalableTargetAction)
	ResetEndTime()
	ResetOverrideLogicalId()
	ResetStartTime()
	ResetTimezone()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppautoscalingScheduledAction
type jsiiProxy_AppautoscalingScheduledAction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ScalableDimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ScalableTargetAction() AppautoscalingScheduledActionScalableTargetActionOutputReference {
	var returns AppautoscalingScheduledActionScalableTargetActionOutputReference
	_jsii_.Get(
		j,
		"scalableTargetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ScalableTargetActionInput() *AppautoscalingScheduledActionScalableTargetAction {
	var returns *AppautoscalingScheduledActionScalableTargetAction
	_jsii_.Get(
		j,
		"scalableTargetActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) ServiceNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledAction) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action aws_appautoscaling_scheduled_action} Resource.
func NewAppautoscalingScheduledAction(scope constructs.Construct, id *string, config *AppautoscalingScheduledActionConfig) AppautoscalingScheduledAction {
	_init_.Initialize()

	j := jsiiProxy_AppautoscalingScheduledAction{}

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingScheduledAction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action aws_appautoscaling_scheduled_action} Resource.
func NewAppautoscalingScheduledAction_Override(a AppautoscalingScheduledAction, scope constructs.Construct, id *string, config *AppautoscalingScheduledActionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingScheduledAction",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetEndTime(val *string) {
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetScalableDimension(val *string) {
	_jsii_.Set(
		j,
		"scalableDimension",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetServiceNamespace(val *string) {
	_jsii_.Set(
		j,
		"serviceNamespace",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledAction) SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AppautoscalingScheduledAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appautoscaling.AppautoscalingScheduledAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppautoscalingScheduledAction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appautoscaling.AppautoscalingScheduledAction",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppautoscalingScheduledAction) PutScalableTargetAction(value *AppautoscalingScheduledActionScalableTargetAction) {
	_jsii_.InvokeVoid(
		a,
		"putScalableTargetAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppautoscalingScheduledAction) ResetEndTime() {
	_jsii_.InvokeVoid(
		a,
		"resetEndTime",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingScheduledAction) ResetStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingScheduledAction) ResetTimezone() {
	_jsii_.InvokeVoid(
		a,
		"resetTimezone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingScheduledAction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppautoscalingScheduledAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledAction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS AppAutoScaling.
type AppautoscalingScheduledActionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#name AppautoscalingScheduledAction#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#resource_id AppautoscalingScheduledAction#resource_id}.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#scalable_dimension AppautoscalingScheduledAction#scalable_dimension}.
	ScalableDimension *string `json:"scalableDimension" yaml:"scalableDimension"`
	// scalable_target_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#scalable_target_action AppautoscalingScheduledAction#scalable_target_action}
	ScalableTargetAction *AppautoscalingScheduledActionScalableTargetAction `json:"scalableTargetAction" yaml:"scalableTargetAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#schedule AppautoscalingScheduledAction#schedule}.
	Schedule *string `json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#service_namespace AppautoscalingScheduledAction#service_namespace}.
	ServiceNamespace *string `json:"serviceNamespace" yaml:"serviceNamespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#end_time AppautoscalingScheduledAction#end_time}.
	EndTime *string `json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#start_time AppautoscalingScheduledAction#start_time}.
	StartTime *string `json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#timezone AppautoscalingScheduledAction#timezone}.
	Timezone *string `json:"timezone" yaml:"timezone"`
}

type AppautoscalingScheduledActionScalableTargetAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#max_capacity AppautoscalingScheduledAction#max_capacity}.
	MaxCapacity *string `json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#min_capacity AppautoscalingScheduledAction#min_capacity}.
	MinCapacity *string `json:"minCapacity" yaml:"minCapacity"`
}

type AppautoscalingScheduledActionScalableTargetActionOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *AppautoscalingScheduledActionScalableTargetAction
	SetInternalValue(val *AppautoscalingScheduledActionScalableTargetAction)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxCapacity() *string
	SetMaxCapacity(val *string)
	MaxCapacityInput() *string
	MinCapacity() *string
	SetMinCapacity(val *string)
	MinCapacityInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetMaxCapacity()
	ResetMinCapacity()
}

// The jsii proxy struct for AppautoscalingScheduledActionScalableTargetActionOutputReference
type jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) InternalValue() *AppautoscalingScheduledActionScalableTargetAction {
	var returns *AppautoscalingScheduledActionScalableTargetAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) MaxCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) MaxCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) MinCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) MinCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppautoscalingScheduledActionScalableTargetActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppautoscalingScheduledActionScalableTargetActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingScheduledActionScalableTargetActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppautoscalingScheduledActionScalableTargetActionOutputReference_Override(a AppautoscalingScheduledActionScalableTargetActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingScheduledActionScalableTargetActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) SetInternalValue(val *AppautoscalingScheduledActionScalableTargetAction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) SetMaxCapacity(val *string) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) SetMinCapacity(val *string) {
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingScheduledActionScalableTargetActionOutputReference) ResetMinCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMinCapacity",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_target aws_appautoscaling_target}.
type AppautoscalingTarget interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxCapacityInput() *float64
	MinCapacity() *float64
	SetMinCapacity(val *float64)
	MinCapacityInput() *float64
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	ScalableDimension() *string
	SetScalableDimension(val *string)
	ScalableDimensionInput() *string
	ServiceNamespace() *string
	SetServiceNamespace(val *string)
	ServiceNamespaceInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetRoleArn()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppautoscalingTarget
type jsiiProxy_AppautoscalingTarget struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppautoscalingTarget) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) MinCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) ScalableDimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) ServiceNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppautoscalingTarget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_target aws_appautoscaling_target} Resource.
func NewAppautoscalingTarget(scope constructs.Construct, id *string, config *AppautoscalingTargetConfig) AppautoscalingTarget {
	_init_.Initialize()

	j := jsiiProxy_AppautoscalingTarget{}

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_target aws_appautoscaling_target} Resource.
func NewAppautoscalingTarget_Override(a AppautoscalingTarget, scope constructs.Construct, id *string, config *AppautoscalingTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appautoscaling.AppautoscalingTarget",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetMaxCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetMinCapacity(val *float64) {
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetScalableDimension(val *string) {
	_jsii_.Set(
		j,
		"scalableDimension",
		val,
	)
}

func (j *jsiiProxy_AppautoscalingTarget) SetServiceNamespace(val *string) {
	_jsii_.Set(
		j,
		"serviceNamespace",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AppautoscalingTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appautoscaling.AppautoscalingTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppautoscalingTarget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appautoscaling.AppautoscalingTarget",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingTarget) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppautoscalingTarget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppautoscalingTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppautoscalingTarget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS AppAutoScaling.
type AppautoscalingTargetConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_target#max_capacity AppautoscalingTarget#max_capacity}.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_target#min_capacity AppautoscalingTarget#min_capacity}.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_target#resource_id AppautoscalingTarget#resource_id}.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_target#scalable_dimension AppautoscalingTarget#scalable_dimension}.
	ScalableDimension *string `json:"scalableDimension" yaml:"scalableDimension"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_target#service_namespace AppautoscalingTarget#service_namespace}.
	ServiceNamespace *string `json:"serviceNamespace" yaml:"serviceNamespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_target#role_arn AppautoscalingTarget#role_arn}.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}
