package organizations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/organizations/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/organizations_delegated_administrators.html aws_organizations_delegated_administrators}.
type DataAwsOrganizationsDelegatedAdministrators interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServicePrincipal() *string
	SetServicePrincipal(val *string)
	ServicePrincipalInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	DelegatedAdministrators(index *string) DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetServicePrincipal()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsOrganizationsDelegatedAdministrators
type jsiiProxy_DataAwsOrganizationsDelegatedAdministrators struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) ServicePrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) ServicePrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/organizations_delegated_administrators.html aws_organizations_delegated_administrators} Data Source.
func NewDataAwsOrganizationsDelegatedAdministrators(scope constructs.Construct, id *string, config *DataAwsOrganizationsDelegatedAdministratorsConfig) DataAwsOrganizationsDelegatedAdministrators {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsDelegatedAdministrators{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedAdministrators",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/organizations_delegated_administrators.html aws_organizations_delegated_administrators} Data Source.
func NewDataAwsOrganizationsDelegatedAdministrators_Override(d DataAwsOrganizationsDelegatedAdministrators, scope constructs.Construct, id *string, config *DataAwsOrganizationsDelegatedAdministratorsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedAdministrators",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) SetServicePrincipal(val *string) {
	_jsii_.Set(
		j,
		"servicePrincipal",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsOrganizationsDelegatedAdministrators_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedAdministrators",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsOrganizationsDelegatedAdministrators_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedAdministrators",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) DelegatedAdministrators(index *string) DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators {
	var returns DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators

	_jsii_.Invoke(
		d,
		"delegatedAdministrators",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) ResetServicePrincipal() {
	_jsii_.InvokeVoid(
		d,
		"resetServicePrincipal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministrators) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsOrganizationsDelegatedAdministratorsConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/organizations_delegated_administrators.html#service_principal DataAwsOrganizationsDelegatedAdministrators#service_principal}.
	ServicePrincipal *string `json:"servicePrincipal"`
}

type DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DelegationEnabledDate() *string
	Email() *string
	Id() *string
	JoinedMethod() *string
	JoinedTimestamp() *string
	Name() *string
	Status() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators
type jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) DelegationEnabledDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegationEnabledDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) JoinedMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinedMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) JoinedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators_Override(d DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedAdministratorsDelegatedAdministrators) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/organizations_delegated_services.html aws_organizations_delegated_services}.
type DataAwsOrganizationsDelegatedServices interface {
	cdktf.TerraformDataSource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	DelegatedServices(index *string) DataAwsOrganizationsDelegatedServicesDelegatedServices
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsOrganizationsDelegatedServices
type jsiiProxy_DataAwsOrganizationsDelegatedServices struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/organizations_delegated_services.html aws_organizations_delegated_services} Data Source.
func NewDataAwsOrganizationsDelegatedServices(scope constructs.Construct, id *string, config *DataAwsOrganizationsDelegatedServicesConfig) DataAwsOrganizationsDelegatedServices {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsDelegatedServices{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedServices",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/organizations_delegated_services.html aws_organizations_delegated_services} Data Source.
func NewDataAwsOrganizationsDelegatedServices_Override(d DataAwsOrganizationsDelegatedServices, scope constructs.Construct, id *string, config *DataAwsOrganizationsDelegatedServicesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedServices",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServices) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsOrganizationsDelegatedServices_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedServices",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsOrganizationsDelegatedServices_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedServices",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) DelegatedServices(index *string) DataAwsOrganizationsDelegatedServicesDelegatedServices {
	var returns DataAwsOrganizationsDelegatedServicesDelegatedServices

	_jsii_.Invoke(
		d,
		"delegatedServices",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServices) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsOrganizationsDelegatedServicesConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/organizations_delegated_services.html#account_id DataAwsOrganizationsDelegatedServices#account_id}.
	AccountId *string `json:"accountId"`
}

type DataAwsOrganizationsDelegatedServicesDelegatedServices interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DelegationEnabledDate() *string
	ServicePrincipal() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsOrganizationsDelegatedServicesDelegatedServices
type jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) DelegationEnabledDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegationEnabledDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) ServicePrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsOrganizationsDelegatedServicesDelegatedServices(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsOrganizationsDelegatedServicesDelegatedServices {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedServicesDelegatedServices",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsOrganizationsDelegatedServicesDelegatedServices_Override(d DataAwsOrganizationsDelegatedServicesDelegatedServices, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsDelegatedServicesDelegatedServices",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsDelegatedServicesDelegatedServices) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/organizations_organization.html aws_organizations_organization}.
type DataAwsOrganizationsOrganization interface {
	cdktf.TerraformDataSource
	Arn() *string
	AwsServiceAccessPrincipals() *[]*string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EnabledPolicyTypes() *[]*string
	FeatureSet() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MasterAccountArn() *string
	MasterAccountEmail() *string
	MasterAccountId() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Accounts(index *string) DataAwsOrganizationsOrganizationAccounts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	NonMasterAccounts(index *string) DataAwsOrganizationsOrganizationNonMasterAccounts
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	Roots(index *string) DataAwsOrganizationsOrganizationRoots
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsOrganizationsOrganization
type jsiiProxy_DataAwsOrganizationsOrganization struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) AwsServiceAccessPrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"awsServiceAccessPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) EnabledPolicyTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledPolicyTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) FeatureSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) MasterAccountArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAccountArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) MasterAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) MasterAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/organizations_organization.html aws_organizations_organization} Data Source.
func NewDataAwsOrganizationsOrganization(scope constructs.Construct, id *string, config *DataAwsOrganizationsOrganizationConfig) DataAwsOrganizationsOrganization {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsOrganization{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganization",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/organizations_organization.html aws_organizations_organization} Data Source.
func NewDataAwsOrganizationsOrganization_Override(d DataAwsOrganizationsOrganization, scope constructs.Construct, id *string, config *DataAwsOrganizationsOrganizationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganization",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganization) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsOrganizationsOrganization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsOrganizationsOrganization_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganization",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsOrganizationsOrganization) Accounts(index *string) DataAwsOrganizationsOrganizationAccounts {
	var returns DataAwsOrganizationsOrganizationAccounts

	_jsii_.Invoke(
		d,
		"accounts",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOrganizationsOrganization) NonMasterAccounts(index *string) DataAwsOrganizationsOrganizationNonMasterAccounts {
	var returns DataAwsOrganizationsOrganizationNonMasterAccounts

	_jsii_.Invoke(
		d,
		"nonMasterAccounts",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrganizationsOrganization) Roots(index *string) DataAwsOrganizationsOrganizationRoots {
	var returns DataAwsOrganizationsOrganizationRoots

	_jsii_.Invoke(
		d,
		"roots",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsOrganizationsOrganization) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganization) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsOrganizationsOrganizationAccounts interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Email() *string
	Id() *string
	Name() *string
	Status() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsOrganizationsOrganizationAccounts
type jsiiProxy_DataAwsOrganizationsOrganizationAccounts struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsOrganizationsOrganizationAccounts(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsOrganizationsOrganizationAccounts {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsOrganizationAccounts{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsOrganizationsOrganizationAccounts_Override(d DataAwsOrganizationsOrganizationAccounts, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationAccounts) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsOrganizationsOrganizationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
}

type DataAwsOrganizationsOrganizationNonMasterAccounts interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Email() *string
	Id() *string
	Name() *string
	Status() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsOrganizationsOrganizationNonMasterAccounts
type jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsOrganizationsOrganizationNonMasterAccounts(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsOrganizationsOrganizationNonMasterAccounts {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationNonMasterAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsOrganizationsOrganizationNonMasterAccounts_Override(d DataAwsOrganizationsOrganizationNonMasterAccounts, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationNonMasterAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationNonMasterAccounts) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsOrganizationsOrganizationRoots interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Id() *string
	Name() *string
	PolicyTypes() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsOrganizationsOrganizationRoots
type jsiiProxy_DataAwsOrganizationsOrganizationRoots struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) PolicyTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsOrganizationsOrganizationRoots(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsOrganizationsOrganizationRoots {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsOrganizationRoots{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationRoots",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsOrganizationsOrganizationRoots_Override(d DataAwsOrganizationsOrganizationRoots, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationRoots",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRoots) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRoots) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRoots) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRoots) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRoots) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRoots) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsOrganizationsOrganizationRootsPolicyTypes interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Status() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Type() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsOrganizationsOrganizationRootsPolicyTypes
type jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsOrganizationsOrganizationRootsPolicyTypes(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsOrganizationsOrganizationRootsPolicyTypes {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationRootsPolicyTypes",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsOrganizationsOrganizationRootsPolicyTypes_Override(d DataAwsOrganizationsOrganizationRootsPolicyTypes, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationRootsPolicyTypes",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationRootsPolicyTypes) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/organizations_organizational_units.html aws_organizations_organizational_units}.
type DataAwsOrganizationsOrganizationalUnits interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	ParentId() *string
	SetParentId(val *string)
	ParentIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	Children(index *string) DataAwsOrganizationsOrganizationalUnitsChildren
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsOrganizationsOrganizationalUnits
type jsiiProxy_DataAwsOrganizationsOrganizationalUnits struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/organizations_organizational_units.html aws_organizations_organizational_units} Data Source.
func NewDataAwsOrganizationsOrganizationalUnits(scope constructs.Construct, id *string, config *DataAwsOrganizationsOrganizationalUnitsConfig) DataAwsOrganizationsOrganizationalUnits {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsOrganizationalUnits{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationalUnits",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/organizations_organizational_units.html aws_organizations_organizational_units} Data Source.
func NewDataAwsOrganizationsOrganizationalUnits_Override(d DataAwsOrganizationsOrganizationalUnits, scope constructs.Construct, id *string, config *DataAwsOrganizationsOrganizationalUnitsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationalUnits",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) SetParentId(val *string) {
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsOrganizationsOrganizationalUnits_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationalUnits",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsOrganizationsOrganizationalUnits_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationalUnits",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) Children(index *string) DataAwsOrganizationsOrganizationalUnitsChildren {
	var returns DataAwsOrganizationsOrganizationalUnitsChildren

	_jsii_.Invoke(
		d,
		"children",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnits) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsOrganizationsOrganizationalUnitsChildren interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Id() *string
	Name() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsOrganizationsOrganizationalUnitsChildren
type jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsOrganizationsOrganizationalUnitsChildren(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsOrganizationsOrganizationalUnitsChildren {
	_init_.Initialize()

	j := jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationalUnitsChildren",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsOrganizationsOrganizationalUnitsChildren_Override(d DataAwsOrganizationsOrganizationalUnitsChildren, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.DataAwsOrganizationsOrganizationalUnitsChildren",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsOrganizationsOrganizationalUnitsChildren) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsOrganizationsOrganizationalUnitsConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/organizations_organizational_units.html#parent_id DataAwsOrganizationsOrganizationalUnits#parent_id}.
	ParentId *string `json:"parentId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html aws_organizations_account}.
type OrganizationsAccount interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	IamUserAccessToBilling() *string
	SetIamUserAccessToBilling(val *string)
	IamUserAccessToBillingInput() *string
	Id() *string
	JoinedMethod() *string
	JoinedTimestamp() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	ParentId() *string
	SetParentId(val *string)
	ParentIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	Status() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
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
	ResetIamUserAccessToBilling()
	ResetOverrideLogicalId()
	ResetParentId()
	ResetRoleName()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OrganizationsAccount
type jsiiProxy_OrganizationsAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrganizationsAccount) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) IamUserAccessToBilling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamUserAccessToBilling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) IamUserAccessToBillingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamUserAccessToBillingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) JoinedMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinedMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) JoinedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html aws_organizations_account} Resource.
func NewOrganizationsAccount(scope constructs.Construct, id *string, config *OrganizationsAccountConfig) OrganizationsAccount {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsAccount{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html aws_organizations_account} Resource.
func NewOrganizationsAccount_Override(o OrganizationsAccount, scope constructs.Construct, id *string, config *OrganizationsAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsAccount",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetEmail(val *string) {
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetIamUserAccessToBilling(val *string) {
	_jsii_.Set(
		j,
		"iamUserAccessToBilling",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetParentId(val *string) {
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetRoleName(val *string) {
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OrganizationsAccount) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OrganizationsAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.OrganizationsAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrganizationsAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.OrganizationsAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsAccount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsAccount) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OrganizationsAccount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OrganizationsAccount) ResetIamUserAccessToBilling() {
	_jsii_.InvokeVoid(
		o,
		"resetIamUserAccessToBilling",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OrganizationsAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsAccount) ResetParentId() {
	_jsii_.InvokeVoid(
		o,
		"resetParentId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsAccount) ResetRoleName() {
	_jsii_.InvokeVoid(
		o,
		"resetRoleName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsAccount) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsAccount) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OrganizationsAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OrganizationsAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OrganizationsAccountConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html#email OrganizationsAccount#email}.
	Email *string `json:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html#name OrganizationsAccount#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html#iam_user_access_to_billing OrganizationsAccount#iam_user_access_to_billing}.
	IamUserAccessToBilling *string `json:"iamUserAccessToBilling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html#parent_id OrganizationsAccount#parent_id}.
	ParentId *string `json:"parentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html#role_name OrganizationsAccount#role_name}.
	RoleName *string `json:"roleName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html#tags OrganizationsAccount#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_account.html#tags_all OrganizationsAccount#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/organizations_delegated_administrator.html aws_organizations_delegated_administrator}.
type OrganizationsDelegatedAdministrator interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DelegationEnabledDate() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Email() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	JoinedMethod() *string
	JoinedTimestamp() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServicePrincipal() *string
	SetServicePrincipal(val *string)
	ServicePrincipalInput() *string
	Status() *string
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
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OrganizationsDelegatedAdministrator
type jsiiProxy_OrganizationsDelegatedAdministrator struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) DelegationEnabledDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegationEnabledDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) JoinedMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinedMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) JoinedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"joinedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) ServicePrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) ServicePrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_delegated_administrator.html aws_organizations_delegated_administrator} Resource.
func NewOrganizationsDelegatedAdministrator(scope constructs.Construct, id *string, config *OrganizationsDelegatedAdministratorConfig) OrganizationsDelegatedAdministrator {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsDelegatedAdministrator{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsDelegatedAdministrator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_delegated_administrator.html aws_organizations_delegated_administrator} Resource.
func NewOrganizationsDelegatedAdministrator_Override(o OrganizationsDelegatedAdministrator, scope constructs.Construct, id *string, config *OrganizationsDelegatedAdministratorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsDelegatedAdministrator",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) SetAccountId(val *string) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrganizationsDelegatedAdministrator) SetServicePrincipal(val *string) {
	_jsii_.Set(
		j,
		"servicePrincipal",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OrganizationsDelegatedAdministrator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.OrganizationsDelegatedAdministrator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrganizationsDelegatedAdministrator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.OrganizationsDelegatedAdministrator",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsDelegatedAdministrator) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OrganizationsDelegatedAdministrator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OrganizationsDelegatedAdministratorConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_delegated_administrator.html#account_id OrganizationsDelegatedAdministrator#account_id}.
	AccountId *string `json:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_delegated_administrator.html#service_principal OrganizationsDelegatedAdministrator#service_principal}.
	ServicePrincipal *string `json:"servicePrincipal"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/organizations_organization.html aws_organizations_organization}.
type OrganizationsOrganization interface {
	cdktf.TerraformResource
	Arn() *string
	AwsServiceAccessPrincipals() *[]*string
	SetAwsServiceAccessPrincipals(val *[]*string)
	AwsServiceAccessPrincipalsInput() *[]*string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EnabledPolicyTypes() *[]*string
	SetEnabledPolicyTypes(val *[]*string)
	EnabledPolicyTypesInput() *[]*string
	FeatureSet() *string
	SetFeatureSet(val *string)
	FeatureSetInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MasterAccountArn() *string
	MasterAccountEmail() *string
	MasterAccountId() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Accounts(index *string) OrganizationsOrganizationAccounts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	NonMasterAccounts(index *string) OrganizationsOrganizationNonMasterAccounts
	OverrideLogicalId(newLogicalId *string)
	ResetAwsServiceAccessPrincipals()
	ResetEnabledPolicyTypes()
	ResetFeatureSet()
	ResetOverrideLogicalId()
	Roots(index *string) OrganizationsOrganizationRoots
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OrganizationsOrganization
type jsiiProxy_OrganizationsOrganization struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrganizationsOrganization) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) AwsServiceAccessPrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"awsServiceAccessPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) AwsServiceAccessPrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"awsServiceAccessPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) EnabledPolicyTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledPolicyTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) EnabledPolicyTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledPolicyTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) FeatureSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) FeatureSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) MasterAccountArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAccountArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) MasterAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) MasterAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganization) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_organization.html aws_organizations_organization} Resource.
func NewOrganizationsOrganization(scope constructs.Construct, id *string, config *OrganizationsOrganizationConfig) OrganizationsOrganization {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsOrganization{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganization",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_organization.html aws_organizations_organization} Resource.
func NewOrganizationsOrganization_Override(o OrganizationsOrganization, scope constructs.Construct, id *string, config *OrganizationsOrganizationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganization",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrganizationsOrganization) SetAwsServiceAccessPrincipals(val *[]*string) {
	_jsii_.Set(
		j,
		"awsServiceAccessPrincipals",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganization) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganization) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganization) SetEnabledPolicyTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"enabledPolicyTypes",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganization) SetFeatureSet(val *string) {
	_jsii_.Set(
		j,
		"featureSet",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganization) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganization) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OrganizationsOrganization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.OrganizationsOrganization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrganizationsOrganization_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.OrganizationsOrganization",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OrganizationsOrganization) Accounts(index *string) OrganizationsOrganizationAccounts {
	var returns OrganizationsOrganizationAccounts

	_jsii_.Invoke(
		o,
		"accounts",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationsOrganization) NonMasterAccounts(index *string) OrganizationsOrganizationNonMasterAccounts {
	var returns OrganizationsOrganizationNonMasterAccounts

	_jsii_.Invoke(
		o,
		"nonMasterAccounts",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OrganizationsOrganization) ResetAwsServiceAccessPrincipals() {
	_jsii_.InvokeVoid(
		o,
		"resetAwsServiceAccessPrincipals",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsOrganization) ResetEnabledPolicyTypes() {
	_jsii_.InvokeVoid(
		o,
		"resetEnabledPolicyTypes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsOrganization) ResetFeatureSet() {
	_jsii_.InvokeVoid(
		o,
		"resetFeatureSet",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsOrganization) Roots(index *string) OrganizationsOrganizationRoots {
	var returns OrganizationsOrganizationRoots

	_jsii_.Invoke(
		o,
		"roots",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationsOrganization) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OrganizationsOrganization) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OrganizationsOrganization) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OrganizationsOrganizationAccounts interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Email() *string
	Id() *string
	Name() *string
	Status() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for OrganizationsOrganizationAccounts
type jsiiProxy_OrganizationsOrganizationAccounts struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewOrganizationsOrganizationAccounts(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) OrganizationsOrganizationAccounts {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsOrganizationAccounts{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewOrganizationsOrganizationAccounts_Override(o OrganizationsOrganizationAccounts, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		o,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationAccounts) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationAccounts) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationAccounts) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationAccounts) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationAccounts) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationAccounts) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type OrganizationsOrganizationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_organization.html#aws_service_access_principals OrganizationsOrganization#aws_service_access_principals}.
	AwsServiceAccessPrincipals *[]*string `json:"awsServiceAccessPrincipals"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_organization.html#enabled_policy_types OrganizationsOrganization#enabled_policy_types}.
	EnabledPolicyTypes *[]*string `json:"enabledPolicyTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_organization.html#feature_set OrganizationsOrganization#feature_set}.
	FeatureSet *string `json:"featureSet"`
}

type OrganizationsOrganizationNonMasterAccounts interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Email() *string
	Id() *string
	Name() *string
	Status() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for OrganizationsOrganizationNonMasterAccounts
type jsiiProxy_OrganizationsOrganizationNonMasterAccounts struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewOrganizationsOrganizationNonMasterAccounts(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) OrganizationsOrganizationNonMasterAccounts {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsOrganizationNonMasterAccounts{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationNonMasterAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewOrganizationsOrganizationNonMasterAccounts_Override(o OrganizationsOrganizationNonMasterAccounts, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationNonMasterAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		o,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationNonMasterAccounts) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type OrganizationsOrganizationRoots interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Id() *string
	Name() *string
	PolicyTypes() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for OrganizationsOrganizationRoots
type jsiiProxy_OrganizationsOrganizationRoots struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) PolicyTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewOrganizationsOrganizationRoots(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) OrganizationsOrganizationRoots {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsOrganizationRoots{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationRoots",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewOrganizationsOrganizationRoots_Override(o OrganizationsOrganizationRoots, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationRoots",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		o,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationRoots) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRoots) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRoots) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRoots) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRoots) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRoots) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type OrganizationsOrganizationRootsPolicyTypes interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Status() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Type() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for OrganizationsOrganizationRootsPolicyTypes
type jsiiProxy_OrganizationsOrganizationRootsPolicyTypes struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Experimental.
func NewOrganizationsOrganizationRootsPolicyTypes(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) OrganizationsOrganizationRootsPolicyTypes {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsOrganizationRootsPolicyTypes{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationRootsPolicyTypes",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewOrganizationsOrganizationRootsPolicyTypes_Override(o OrganizationsOrganizationRootsPolicyTypes, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationRootsPolicyTypes",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		o,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationRootsPolicyTypes) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/organizations_organizational_unit.html aws_organizations_organizational_unit}.
type OrganizationsOrganizationalUnit interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	ParentId() *string
	SetParentId(val *string)
	ParentIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Accounts(index *string) OrganizationsOrganizationalUnitAccounts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OrganizationsOrganizationalUnit
type jsiiProxy_OrganizationsOrganizationalUnit struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_organizational_unit.html aws_organizations_organizational_unit} Resource.
func NewOrganizationsOrganizationalUnit(scope constructs.Construct, id *string, config *OrganizationsOrganizationalUnitConfig) OrganizationsOrganizationalUnit {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsOrganizationalUnit{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationalUnit",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_organizational_unit.html aws_organizations_organizational_unit} Resource.
func NewOrganizationsOrganizationalUnit_Override(o OrganizationsOrganizationalUnit, scope constructs.Construct, id *string, config *OrganizationsOrganizationalUnitConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationalUnit",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) SetParentId(val *string) {
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnit) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OrganizationsOrganizationalUnit_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.OrganizationsOrganizationalUnit",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrganizationsOrganizationalUnit_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.OrganizationsOrganizationalUnit",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OrganizationsOrganizationalUnit) Accounts(index *string) OrganizationsOrganizationalUnitAccounts {
	var returns OrganizationsOrganizationalUnitAccounts

	_jsii_.Invoke(
		o,
		"accounts",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsOrganizationalUnit) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsOrganizationalUnit) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsOrganizationalUnit) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnit) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OrganizationsOrganizationalUnitAccounts interface {
	cdktf.ComplexComputedList
	Arn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Email() *string
	Id() *string
	Name() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for OrganizationsOrganizationalUnitAccounts
type jsiiProxy_OrganizationsOrganizationalUnitAccounts struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewOrganizationsOrganizationalUnitAccounts(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) OrganizationsOrganizationalUnitAccounts {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsOrganizationalUnitAccounts{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationalUnitAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewOrganizationsOrganizationalUnitAccounts_Override(o OrganizationsOrganizationalUnitAccounts, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsOrganizationalUnitAccounts",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		o,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationsOrganizationalUnitAccounts) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnitAccounts) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnitAccounts) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnitAccounts) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnitAccounts) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsOrganizationalUnitAccounts) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type OrganizationsOrganizationalUnitConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_organizational_unit.html#name OrganizationsOrganizationalUnit#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_organizational_unit.html#parent_id OrganizationsOrganizationalUnit#parent_id}.
	ParentId *string `json:"parentId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_organizational_unit.html#tags OrganizationsOrganizationalUnit#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_organizational_unit.html#tags_all OrganizationsOrganizationalUnit#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy.html aws_organizations_policy}.
type OrganizationsPolicy interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OrganizationsPolicy
type jsiiProxy_OrganizationsPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrganizationsPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicy) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy.html aws_organizations_policy} Resource.
func NewOrganizationsPolicy(scope constructs.Construct, id *string, config *OrganizationsPolicyConfig) OrganizationsPolicy {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsPolicy{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy.html aws_organizations_policy} Resource.
func NewOrganizationsPolicy_Override(o OrganizationsPolicy, scope constructs.Construct, id *string, config *OrganizationsPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsPolicy",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetContent(val *string) {
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicy) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OrganizationsPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.OrganizationsPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrganizationsPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.OrganizationsPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OrganizationsPolicy) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsPolicy) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsPolicy) ResetTagsAll() {
	_jsii_.InvokeVoid(
		o,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsPolicy) ResetType() {
	_jsii_.InvokeVoid(
		o,
		"resetType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OrganizationsPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OrganizationsPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy_attachment.html aws_organizations_policy_attachment}.
type OrganizationsPolicyAttachment interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PolicyId() *string
	SetPolicyId(val *string)
	PolicyIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TargetId() *string
	SetTargetId(val *string)
	TargetIdInput() *string
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
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for OrganizationsPolicyAttachment
type jsiiProxy_OrganizationsPolicyAttachment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) PolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) TargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) TargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy_attachment.html aws_organizations_policy_attachment} Resource.
func NewOrganizationsPolicyAttachment(scope constructs.Construct, id *string, config *OrganizationsPolicyAttachmentConfig) OrganizationsPolicyAttachment {
	_init_.Initialize()

	j := jsiiProxy_OrganizationsPolicyAttachment{}

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsPolicyAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy_attachment.html aws_organizations_policy_attachment} Resource.
func NewOrganizationsPolicyAttachment_Override(o OrganizationsPolicyAttachment, scope constructs.Construct, id *string, config *OrganizationsPolicyAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Organizations.OrganizationsPolicyAttachment",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) SetPolicyId(val *string) {
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrganizationsPolicyAttachment) SetTargetId(val *string) {
	_jsii_.Set(
		j,
		"targetId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func OrganizationsPolicyAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Organizations.OrganizationsPolicyAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrganizationsPolicyAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Organizations.OrganizationsPolicyAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationsPolicyAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (o *jsiiProxy_OrganizationsPolicyAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (o *jsiiProxy_OrganizationsPolicyAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type OrganizationsPolicyAttachmentConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy_attachment.html#policy_id OrganizationsPolicyAttachment#policy_id}.
	PolicyId *string `json:"policyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy_attachment.html#target_id OrganizationsPolicyAttachment#target_id}.
	TargetId *string `json:"targetId"`
}

type OrganizationsPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy.html#content OrganizationsPolicy#content}.
	Content *string `json:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy.html#name OrganizationsPolicy#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy.html#description OrganizationsPolicy#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy.html#tags OrganizationsPolicy#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy.html#tags_all OrganizationsPolicy#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/organizations_policy.html#type OrganizationsPolicy#type}.
	Type *string `json:"type"`
}
