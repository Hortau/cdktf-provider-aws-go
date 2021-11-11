package budgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/budgets/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html aws_budgets_budget}.
type BudgetsBudget interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Arn() *string
	BudgetType() *string
	SetBudgetType(val *string)
	BudgetTypeInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	CostFilter() *[]*BudgetsBudgetCostFilter
	SetCostFilter(val *[]*BudgetsBudgetCostFilter)
	CostFilterInput() *[]*BudgetsBudgetCostFilter
	CostFilters() interface{}
	SetCostFilters(val interface{})
	CostFiltersInput() interface{}
	CostTypes() BudgetsBudgetCostTypesOutputReference
	CostTypesInput() *BudgetsBudgetCostTypes
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LimitAmount() *string
	SetLimitAmount(val *string)
	LimitAmountInput() *string
	LimitUnit() *string
	SetLimitUnit(val *string)
	LimitUnitInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	Node() constructs.Node
	Notification() *[]*BudgetsBudgetNotification
	SetNotification(val *[]*BudgetsBudgetNotification)
	NotificationInput() *[]*BudgetsBudgetNotification
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TimePeriodEnd() *string
	SetTimePeriodEnd(val *string)
	TimePeriodEndInput() *string
	TimePeriodStart() *string
	SetTimePeriodStart(val *string)
	TimePeriodStartInput() *string
	TimeUnit() *string
	SetTimeUnit(val *string)
	TimeUnitInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutCostTypes(value *BudgetsBudgetCostTypes)
	ResetAccountId()
	ResetCostFilter()
	ResetCostFilters()
	ResetCostTypes()
	ResetName()
	ResetNamePrefix()
	ResetNotification()
	ResetOverrideLogicalId()
	ResetTimePeriodEnd()
	ResetTimePeriodStart()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for BudgetsBudget
type jsiiProxy_BudgetsBudget struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BudgetsBudget) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) BudgetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) BudgetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostFilter() *[]*BudgetsBudgetCostFilter {
	var returns *[]*BudgetsBudgetCostFilter
	_jsii_.Get(
		j,
		"costFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostFilterInput() *[]*BudgetsBudgetCostFilter {
	var returns *[]*BudgetsBudgetCostFilter
	_jsii_.Get(
		j,
		"costFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostFilters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"costFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostFiltersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"costFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostTypes() BudgetsBudgetCostTypesOutputReference {
	var returns BudgetsBudgetCostTypesOutputReference
	_jsii_.Get(
		j,
		"costTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) CostTypesInput() *BudgetsBudgetCostTypes {
	var returns *BudgetsBudgetCostTypes
	_jsii_.Get(
		j,
		"costTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) LimitAmount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitAmount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) LimitAmountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitAmountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) LimitUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) LimitUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Notification() *[]*BudgetsBudgetNotification {
	var returns *[]*BudgetsBudgetNotification
	_jsii_.Get(
		j,
		"notification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) NotificationInput() *[]*BudgetsBudgetNotification {
	var returns *[]*BudgetsBudgetNotification
	_jsii_.Get(
		j,
		"notificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimePeriodEnd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimePeriodEndInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimePeriodStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimePeriodStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timePeriodStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudget) TimeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeUnitInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html aws_budgets_budget} Resource.
func NewBudgetsBudget(scope constructs.Construct, id *string, config *BudgetsBudgetConfig) BudgetsBudget {
	_init_.Initialize()

	j := jsiiProxy_BudgetsBudget{}

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html aws_budgets_budget} Resource.
func NewBudgetsBudget_Override(b BudgetsBudget, scope constructs.Construct, id *string, config *BudgetsBudgetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudget",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetBudgetType(val *string) {
	_jsii_.Set(
		j,
		"budgetType",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetCostFilter(val *[]*BudgetsBudgetCostFilter) {
	_jsii_.Set(
		j,
		"costFilter",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetCostFilters(val interface{}) {
	_jsii_.Set(
		j,
		"costFilters",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetLimitAmount(val *string) {
	_jsii_.Set(
		j,
		"limitAmount",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetLimitUnit(val *string) {
	_jsii_.Set(
		j,
		"limitUnit",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetNotification(val *[]*BudgetsBudgetNotification) {
	_jsii_.Set(
		j,
		"notification",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetTimePeriodEnd(val *string) {
	_jsii_.Set(
		j,
		"timePeriodEnd",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetTimePeriodStart(val *string) {
	_jsii_.Set(
		j,
		"timePeriodStart",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudget) SetTimeUnit(val *string) {
	_jsii_.Set(
		j,
		"timeUnit",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func BudgetsBudget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Budgets.BudgetsBudget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BudgetsBudget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Budgets.BudgetsBudget",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudget) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (b *jsiiProxy_BudgetsBudget) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudget) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudget) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudget) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudget) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (b *jsiiProxy_BudgetsBudget) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BudgetsBudget) PutCostTypes(value *BudgetsBudgetCostTypes) {
	_jsii_.InvokeVoid(
		b,
		"putCostTypes",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetAccountId() {
	_jsii_.InvokeVoid(
		b,
		"resetAccountId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetCostFilter() {
	_jsii_.InvokeVoid(
		b,
		"resetCostFilter",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetCostFilters() {
	_jsii_.InvokeVoid(
		b,
		"resetCostFilters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetCostTypes() {
	_jsii_.InvokeVoid(
		b,
		"resetCostTypes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		b,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetNotification() {
	_jsii_.InvokeVoid(
		b,
		"resetNotification",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (b *jsiiProxy_BudgetsBudget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetTimePeriodEnd() {
	_jsii_.InvokeVoid(
		b,
		"resetTimePeriodEnd",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) ResetTimePeriodStart() {
	_jsii_.InvokeVoid(
		b,
		"resetTimePeriodStart",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (b *jsiiProxy_BudgetsBudget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (b *jsiiProxy_BudgetsBudget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html aws_budgets_budget_action}.
type BudgetsBudgetAction interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	ActionId() *string
	ActionThreshold() BudgetsBudgetActionActionThresholdOutputReference
	ActionThresholdInput() *BudgetsBudgetActionActionThreshold
	ActionType() *string
	SetActionType(val *string)
	ActionTypeInput() *string
	ApprovalModel() *string
	SetApprovalModel(val *string)
	ApprovalModelInput() *string
	Arn() *string
	BudgetName() *string
	SetBudgetName(val *string)
	BudgetNameInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	Definition() BudgetsBudgetActionDefinitionOutputReference
	DefinitionInput() *BudgetsBudgetActionDefinition
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	ExecutionRoleArnInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	NotificationType() *string
	SetNotificationType(val *string)
	NotificationTypeInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	Subscriber() *[]*BudgetsBudgetActionSubscriber
	SetSubscriber(val *[]*BudgetsBudgetActionSubscriber)
	SubscriberInput() *[]*BudgetsBudgetActionSubscriber
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutActionThreshold(value *BudgetsBudgetActionActionThreshold)
	PutDefinition(value *BudgetsBudgetActionDefinition)
	ResetAccountId()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for BudgetsBudgetAction
type jsiiProxy_BudgetsBudgetAction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BudgetsBudgetAction) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionThreshold() BudgetsBudgetActionActionThresholdOutputReference {
	var returns BudgetsBudgetActionActionThresholdOutputReference
	_jsii_.Get(
		j,
		"actionThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionThresholdInput() *BudgetsBudgetActionActionThreshold {
	var returns *BudgetsBudgetActionActionThreshold
	_jsii_.Get(
		j,
		"actionThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ActionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ApprovalModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ApprovalModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) BudgetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) BudgetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"budgetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Definition() BudgetsBudgetActionDefinitionOutputReference {
	var returns BudgetsBudgetActionDefinitionOutputReference
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) DefinitionInput() *BudgetsBudgetActionDefinition {
	var returns *BudgetsBudgetActionDefinition
	_jsii_.Get(
		j,
		"definitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) NotificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) NotificationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) Subscriber() *[]*BudgetsBudgetActionSubscriber {
	var returns *[]*BudgetsBudgetActionSubscriber
	_jsii_.Get(
		j,
		"subscriber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) SubscriberInput() *[]*BudgetsBudgetActionSubscriber {
	var returns *[]*BudgetsBudgetActionSubscriber
	_jsii_.Get(
		j,
		"subscriberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetAction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html aws_budgets_budget_action} Resource.
func NewBudgetsBudgetAction(scope constructs.Construct, id *string, config *BudgetsBudgetActionConfig) BudgetsBudgetAction {
	_init_.Initialize()

	j := jsiiProxy_BudgetsBudgetAction{}

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetAction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html aws_budgets_budget_action} Resource.
func NewBudgetsBudgetAction_Override(b BudgetsBudgetAction, scope constructs.Construct, id *string, config *BudgetsBudgetActionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetAction",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetActionType(val *string) {
	_jsii_.Set(
		j,
		"actionType",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetApprovalModel(val *string) {
	_jsii_.Set(
		j,
		"approvalModel",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetBudgetName(val *string) {
	_jsii_.Set(
		j,
		"budgetName",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetNotificationType(val *string) {
	_jsii_.Set(
		j,
		"notificationType",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetAction) SetSubscriber(val *[]*BudgetsBudgetActionSubscriber) {
	_jsii_.Set(
		j,
		"subscriber",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func BudgetsBudgetAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Budgets.BudgetsBudgetAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BudgetsBudgetAction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Budgets.BudgetsBudgetAction",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) PutActionThreshold(value *BudgetsBudgetActionActionThreshold) {
	_jsii_.InvokeVoid(
		b,
		"putActionThreshold",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) PutDefinition(value *BudgetsBudgetActionDefinition) {
	_jsii_.InvokeVoid(
		b,
		"putDefinition",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) ResetAccountId() {
	_jsii_.InvokeVoid(
		b,
		"resetAccountId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetAction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (b *jsiiProxy_BudgetsBudgetAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (b *jsiiProxy_BudgetsBudgetAction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BudgetsBudgetActionActionThreshold struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#action_threshold_type BudgetsBudgetAction#action_threshold_type}.
	ActionThresholdType *string `json:"actionThresholdType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#action_threshold_value BudgetsBudgetAction#action_threshold_value}.
	ActionThresholdValue *float64 `json:"actionThresholdValue"`
}

type BudgetsBudgetActionActionThresholdOutputReference interface {
	cdktf.ComplexObject
	ActionThresholdType() *string
	SetActionThresholdType(val *string)
	ActionThresholdTypeInput() *string
	ActionThresholdValue() *float64
	SetActionThresholdValue(val *float64)
	ActionThresholdValueInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for BudgetsBudgetActionActionThresholdOutputReference
type jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) ActionThresholdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionThresholdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) ActionThresholdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionThresholdTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) ActionThresholdValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actionThresholdValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) ActionThresholdValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actionThresholdValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBudgetsBudgetActionActionThresholdOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BudgetsBudgetActionActionThresholdOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionActionThresholdOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBudgetsBudgetActionActionThresholdOutputReference_Override(b BudgetsBudgetActionActionThresholdOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionActionThresholdOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) SetActionThresholdType(val *string) {
	_jsii_.Set(
		j,
		"actionThresholdType",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) SetActionThresholdValue(val *float64) {
	_jsii_.Set(
		j,
		"actionThresholdValue",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionActionThresholdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type BudgetsBudgetActionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// action_threshold block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#action_threshold BudgetsBudgetAction#action_threshold}
	ActionThreshold *BudgetsBudgetActionActionThreshold `json:"actionThreshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#action_type BudgetsBudgetAction#action_type}.
	ActionType *string `json:"actionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#approval_model BudgetsBudgetAction#approval_model}.
	ApprovalModel *string `json:"approvalModel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#budget_name BudgetsBudgetAction#budget_name}.
	BudgetName *string `json:"budgetName"`
	// definition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#definition BudgetsBudgetAction#definition}
	Definition *BudgetsBudgetActionDefinition `json:"definition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#execution_role_arn BudgetsBudgetAction#execution_role_arn}.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#notification_type BudgetsBudgetAction#notification_type}.
	NotificationType *string `json:"notificationType"`
	// subscriber block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#subscriber BudgetsBudgetAction#subscriber}
	Subscriber *[]*BudgetsBudgetActionSubscriber `json:"subscriber"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#account_id BudgetsBudgetAction#account_id}.
	AccountId *string `json:"accountId"`
}

type BudgetsBudgetActionDefinition struct {
	// iam_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#iam_action_definition BudgetsBudgetAction#iam_action_definition}
	IamActionDefinition *BudgetsBudgetActionDefinitionIamActionDefinition `json:"iamActionDefinition"`
	// scp_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#scp_action_definition BudgetsBudgetAction#scp_action_definition}
	ScpActionDefinition *BudgetsBudgetActionDefinitionScpActionDefinition `json:"scpActionDefinition"`
	// ssm_action_definition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#ssm_action_definition BudgetsBudgetAction#ssm_action_definition}
	SsmActionDefinition *BudgetsBudgetActionDefinitionSsmActionDefinition `json:"ssmActionDefinition"`
}

type BudgetsBudgetActionDefinitionIamActionDefinition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#policy_arn BudgetsBudgetAction#policy_arn}.
	PolicyArn *string `json:"policyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#groups BudgetsBudgetAction#groups}.
	Groups *[]*string `json:"groups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#roles BudgetsBudgetAction#roles}.
	Roles *[]*string `json:"roles"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#users BudgetsBudgetAction#users}.
	Users *[]*string `json:"users"`
}

type BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference interface {
	cdktf.ComplexObject
	Groups() *[]*string
	SetGroups(val *[]*string)
	GroupsInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	PolicyArn() *string
	SetPolicyArn(val *string)
	PolicyArnInput() *string
	Roles() *[]*string
	SetRoles(val *[]*string)
	RolesInput() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Users() *[]*string
	SetUsers(val *[]*string)
	UsersInput() *[]*string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetGroups()
	ResetRoles()
	ResetUsers()
}

// The jsii proxy struct for BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference
type jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) Groups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) GroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"groupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) PolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) PolicyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) Roles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) RolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) UsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}

func NewBudgetsBudgetActionDefinitionIamActionDefinitionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBudgetsBudgetActionDefinitionIamActionDefinitionOutputReference_Override(b BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) SetGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"groups",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) SetPolicyArn(val *string) {
	_jsii_.Set(
		j,
		"policyArn",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) SetRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"roles",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) SetUsers(val *[]*string) {
	_jsii_.Set(
		j,
		"users",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) ResetGroups() {
	_jsii_.InvokeVoid(
		b,
		"resetGroups",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) ResetRoles() {
	_jsii_.InvokeVoid(
		b,
		"resetRoles",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference) ResetUsers() {
	_jsii_.InvokeVoid(
		b,
		"resetUsers",
		nil, // no parameters
	)
}

type BudgetsBudgetActionDefinitionOutputReference interface {
	cdktf.ComplexObject
	IamActionDefinition() BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference
	IamActionDefinitionInput() *BudgetsBudgetActionDefinitionIamActionDefinition
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ScpActionDefinition() BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference
	ScpActionDefinitionInput() *BudgetsBudgetActionDefinitionScpActionDefinition
	SsmActionDefinition() BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference
	SsmActionDefinitionInput() *BudgetsBudgetActionDefinitionSsmActionDefinition
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutIamActionDefinition(value *BudgetsBudgetActionDefinitionIamActionDefinition)
	PutScpActionDefinition(value *BudgetsBudgetActionDefinitionScpActionDefinition)
	PutSsmActionDefinition(value *BudgetsBudgetActionDefinitionSsmActionDefinition)
	ResetIamActionDefinition()
	ResetScpActionDefinition()
	ResetSsmActionDefinition()
}

// The jsii proxy struct for BudgetsBudgetActionDefinitionOutputReference
type jsiiProxy_BudgetsBudgetActionDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) IamActionDefinition() BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference {
	var returns BudgetsBudgetActionDefinitionIamActionDefinitionOutputReference
	_jsii_.Get(
		j,
		"iamActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) IamActionDefinitionInput() *BudgetsBudgetActionDefinitionIamActionDefinition {
	var returns *BudgetsBudgetActionDefinitionIamActionDefinition
	_jsii_.Get(
		j,
		"iamActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ScpActionDefinition() BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference {
	var returns BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference
	_jsii_.Get(
		j,
		"scpActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ScpActionDefinitionInput() *BudgetsBudgetActionDefinitionScpActionDefinition {
	var returns *BudgetsBudgetActionDefinitionScpActionDefinition
	_jsii_.Get(
		j,
		"scpActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) SsmActionDefinition() BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference {
	var returns BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference
	_jsii_.Get(
		j,
		"ssmActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) SsmActionDefinitionInput() *BudgetsBudgetActionDefinitionSsmActionDefinition {
	var returns *BudgetsBudgetActionDefinitionSsmActionDefinition
	_jsii_.Get(
		j,
		"ssmActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBudgetsBudgetActionDefinitionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BudgetsBudgetActionDefinitionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BudgetsBudgetActionDefinitionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBudgetsBudgetActionDefinitionOutputReference_Override(b BudgetsBudgetActionDefinitionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) PutIamActionDefinition(value *BudgetsBudgetActionDefinitionIamActionDefinition) {
	_jsii_.InvokeVoid(
		b,
		"putIamActionDefinition",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) PutScpActionDefinition(value *BudgetsBudgetActionDefinitionScpActionDefinition) {
	_jsii_.InvokeVoid(
		b,
		"putScpActionDefinition",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) PutSsmActionDefinition(value *BudgetsBudgetActionDefinitionSsmActionDefinition) {
	_jsii_.InvokeVoid(
		b,
		"putSsmActionDefinition",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ResetIamActionDefinition() {
	_jsii_.InvokeVoid(
		b,
		"resetIamActionDefinition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ResetScpActionDefinition() {
	_jsii_.InvokeVoid(
		b,
		"resetScpActionDefinition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetActionDefinitionOutputReference) ResetSsmActionDefinition() {
	_jsii_.InvokeVoid(
		b,
		"resetSsmActionDefinition",
		nil, // no parameters
	)
}

type BudgetsBudgetActionDefinitionScpActionDefinition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#policy_id BudgetsBudgetAction#policy_id}.
	PolicyId *string `json:"policyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#target_ids BudgetsBudgetAction#target_ids}.
	TargetIds *[]*string `json:"targetIds"`
}

type BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	TargetIds() *[]*string
	SetTargetIds(val *[]*string)
	TargetIdsInput() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference
type jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) TargetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) TargetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBudgetsBudgetActionDefinitionScpActionDefinitionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBudgetsBudgetActionDefinitionScpActionDefinitionOutputReference_Override(b BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) SetPolicyId(val *string) {
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) SetTargetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"targetIds",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionScpActionDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type BudgetsBudgetActionDefinitionSsmActionDefinition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#action_sub_type BudgetsBudgetAction#action_sub_type}.
	ActionSubType *string `json:"actionSubType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#instance_ids BudgetsBudgetAction#instance_ids}.
	InstanceIds *[]*string `json:"instanceIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#region BudgetsBudgetAction#region}.
	Region *string `json:"region"`
}

type BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference interface {
	cdktf.ComplexObject
	ActionSubType() *string
	SetActionSubType(val *string)
	ActionSubTypeInput() *string
	InstanceIds() *[]*string
	SetInstanceIds(val *[]*string)
	InstanceIdsInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference
type jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) ActionSubType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionSubType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) ActionSubTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionSubTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) InstanceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) InstanceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference_Override(b BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) SetActionSubType(val *string) {
	_jsii_.Set(
		j,
		"actionSubType",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) SetInstanceIds(val *[]*string) {
	_jsii_.Set(
		j,
		"instanceIds",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetActionDefinitionSsmActionDefinitionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type BudgetsBudgetActionSubscriber struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#address BudgetsBudgetAction#address}.
	Address *string `json:"address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget_action.html#subscription_type BudgetsBudgetAction#subscription_type}.
	SubscriptionType *string `json:"subscriptionType"`
}

type BudgetsBudgetConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#budget_type BudgetsBudget#budget_type}.
	BudgetType *string `json:"budgetType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#limit_amount BudgetsBudget#limit_amount}.
	LimitAmount *string `json:"limitAmount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#limit_unit BudgetsBudget#limit_unit}.
	LimitUnit *string `json:"limitUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#time_unit BudgetsBudget#time_unit}.
	TimeUnit *string `json:"timeUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#account_id BudgetsBudget#account_id}.
	AccountId *string `json:"accountId"`
	// cost_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#cost_filter BudgetsBudget#cost_filter}
	CostFilter *[]*BudgetsBudgetCostFilter `json:"costFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#cost_filters BudgetsBudget#cost_filters}.
	CostFilters interface{} `json:"costFilters"`
	// cost_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#cost_types BudgetsBudget#cost_types}
	CostTypes *BudgetsBudgetCostTypes `json:"costTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#name BudgetsBudget#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#name_prefix BudgetsBudget#name_prefix}.
	NamePrefix *string `json:"namePrefix"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#notification BudgetsBudget#notification}
	Notification *[]*BudgetsBudgetNotification `json:"notification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#time_period_end BudgetsBudget#time_period_end}.
	TimePeriodEnd *string `json:"timePeriodEnd"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#time_period_start BudgetsBudget#time_period_start}.
	TimePeriodStart *string `json:"timePeriodStart"`
}

type BudgetsBudgetCostFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#name BudgetsBudget#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#values BudgetsBudget#values}.
	Values *[]*string `json:"values"`
}

type BudgetsBudgetCostTypes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#include_credit BudgetsBudget#include_credit}.
	IncludeCredit interface{} `json:"includeCredit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#include_discount BudgetsBudget#include_discount}.
	IncludeDiscount interface{} `json:"includeDiscount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#include_other_subscription BudgetsBudget#include_other_subscription}.
	IncludeOtherSubscription interface{} `json:"includeOtherSubscription"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#include_recurring BudgetsBudget#include_recurring}.
	IncludeRecurring interface{} `json:"includeRecurring"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#include_refund BudgetsBudget#include_refund}.
	IncludeRefund interface{} `json:"includeRefund"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#include_subscription BudgetsBudget#include_subscription}.
	IncludeSubscription interface{} `json:"includeSubscription"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#include_support BudgetsBudget#include_support}.
	IncludeSupport interface{} `json:"includeSupport"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#include_tax BudgetsBudget#include_tax}.
	IncludeTax interface{} `json:"includeTax"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#include_upfront BudgetsBudget#include_upfront}.
	IncludeUpfront interface{} `json:"includeUpfront"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#use_amortized BudgetsBudget#use_amortized}.
	UseAmortized interface{} `json:"useAmortized"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#use_blended BudgetsBudget#use_blended}.
	UseBlended interface{} `json:"useBlended"`
}

type BudgetsBudgetCostTypesOutputReference interface {
	cdktf.ComplexObject
	IncludeCredit() interface{}
	SetIncludeCredit(val interface{})
	IncludeCreditInput() interface{}
	IncludeDiscount() interface{}
	SetIncludeDiscount(val interface{})
	IncludeDiscountInput() interface{}
	IncludeOtherSubscription() interface{}
	SetIncludeOtherSubscription(val interface{})
	IncludeOtherSubscriptionInput() interface{}
	IncludeRecurring() interface{}
	SetIncludeRecurring(val interface{})
	IncludeRecurringInput() interface{}
	IncludeRefund() interface{}
	SetIncludeRefund(val interface{})
	IncludeRefundInput() interface{}
	IncludeSubscription() interface{}
	SetIncludeSubscription(val interface{})
	IncludeSubscriptionInput() interface{}
	IncludeSupport() interface{}
	SetIncludeSupport(val interface{})
	IncludeSupportInput() interface{}
	IncludeTax() interface{}
	SetIncludeTax(val interface{})
	IncludeTaxInput() interface{}
	IncludeUpfront() interface{}
	SetIncludeUpfront(val interface{})
	IncludeUpfrontInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UseAmortized() interface{}
	SetUseAmortized(val interface{})
	UseAmortizedInput() interface{}
	UseBlended() interface{}
	SetUseBlended(val interface{})
	UseBlendedInput() interface{}
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIncludeCredit()
	ResetIncludeDiscount()
	ResetIncludeOtherSubscription()
	ResetIncludeRecurring()
	ResetIncludeRefund()
	ResetIncludeSubscription()
	ResetIncludeSupport()
	ResetIncludeTax()
	ResetIncludeUpfront()
	ResetUseAmortized()
	ResetUseBlended()
}

// The jsii proxy struct for BudgetsBudgetCostTypesOutputReference
type jsiiProxy_BudgetsBudgetCostTypesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeCredit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCredit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeCreditInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCreditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeDiscount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDiscount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeDiscountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDiscountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeOtherSubscription() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOtherSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeOtherSubscriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOtherSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeRecurring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRecurring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeRecurringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRecurringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeRefund() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRefund",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeRefundInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRefundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeSubscription() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeSubscriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeTax() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeTaxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeUpfront() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeUpfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IncludeUpfrontInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeUpfrontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) UseAmortized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAmortized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) UseAmortizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAmortizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) UseBlended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBlended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) UseBlendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBlendedInput",
		&returns,
	)
	return returns
}

func NewBudgetsBudgetCostTypesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BudgetsBudgetCostTypesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BudgetsBudgetCostTypesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetCostTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBudgetsBudgetCostTypesOutputReference_Override(b BudgetsBudgetCostTypesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Budgets.BudgetsBudgetCostTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeCredit(val interface{}) {
	_jsii_.Set(
		j,
		"includeCredit",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeDiscount(val interface{}) {
	_jsii_.Set(
		j,
		"includeDiscount",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeOtherSubscription(val interface{}) {
	_jsii_.Set(
		j,
		"includeOtherSubscription",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeRecurring(val interface{}) {
	_jsii_.Set(
		j,
		"includeRecurring",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeRefund(val interface{}) {
	_jsii_.Set(
		j,
		"includeRefund",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeSubscription(val interface{}) {
	_jsii_.Set(
		j,
		"includeSubscription",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeSupport(val interface{}) {
	_jsii_.Set(
		j,
		"includeSupport",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeTax(val interface{}) {
	_jsii_.Set(
		j,
		"includeTax",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIncludeUpfront(val interface{}) {
	_jsii_.Set(
		j,
		"includeUpfront",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetUseAmortized(val interface{}) {
	_jsii_.Set(
		j,
		"useAmortized",
		val,
	)
}

func (j *jsiiProxy_BudgetsBudgetCostTypesOutputReference) SetUseBlended(val interface{}) {
	_jsii_.Set(
		j,
		"useBlended",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeCredit() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeCredit",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeDiscount() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeDiscount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeOtherSubscription() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeOtherSubscription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeRecurring() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeRecurring",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeRefund() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeRefund",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeSubscription() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeSubscription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeSupport() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeSupport",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeTax() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeTax",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetIncludeUpfront() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeUpfront",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetUseAmortized() {
	_jsii_.InvokeVoid(
		b,
		"resetUseAmortized",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BudgetsBudgetCostTypesOutputReference) ResetUseBlended() {
	_jsii_.InvokeVoid(
		b,
		"resetUseBlended",
		nil, // no parameters
	)
}

type BudgetsBudgetNotification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#comparison_operator BudgetsBudget#comparison_operator}.
	ComparisonOperator *string `json:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#notification_type BudgetsBudget#notification_type}.
	NotificationType *string `json:"notificationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#threshold BudgetsBudget#threshold}.
	Threshold *float64 `json:"threshold"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#threshold_type BudgetsBudget#threshold_type}.
	ThresholdType *string `json:"thresholdType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#subscriber_email_addresses BudgetsBudget#subscriber_email_addresses}.
	SubscriberEmailAddresses *[]*string `json:"subscriberEmailAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/budgets_budget.html#subscriber_sns_topic_arns BudgetsBudget#subscriber_sns_topic_arns}.
	SubscriberSnsTopicArns *[]*string `json:"subscriberSnsTopicArns"`
}
