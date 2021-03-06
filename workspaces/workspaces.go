package workspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/workspaces/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/workspaces_bundle aws_workspaces_bundle}.
type DataAwsWorkspacesBundle interface {
	cdktf.TerraformDataSource
	BundleId() *string
	SetBundleId(val *string)
	BundleIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	ComputeType(index *string) DataAwsWorkspacesBundleComputeType
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
	ResetBundleId()
	ResetName()
	ResetOverrideLogicalId()
	ResetOwner()
	RootStorage(index *string) DataAwsWorkspacesBundleRootStorage
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
	UserStorage(index *string) DataAwsWorkspacesBundleUserStorage
}

// The jsii proxy struct for DataAwsWorkspacesBundle
type jsiiProxy_DataAwsWorkspacesBundle struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) BundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/workspaces_bundle aws_workspaces_bundle} Data Source.
func NewDataAwsWorkspacesBundle(scope constructs.Construct, id *string, config *DataAwsWorkspacesBundleConfig) DataAwsWorkspacesBundle {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesBundle{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundle",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/workspaces_bundle aws_workspaces_bundle} Data Source.
func NewDataAwsWorkspacesBundle_Override(d DataAwsWorkspacesBundle, scope constructs.Construct, id *string, config *DataAwsWorkspacesBundleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundle",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundle) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsWorkspacesBundle_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundle",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsWorkspacesBundle_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundle",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundle) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsWorkspacesBundle) ComputeType(index *string) DataAwsWorkspacesBundleComputeType {
	var returns DataAwsWorkspacesBundleComputeType

	_jsii_.Invoke(
		d,
		"computeType",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundle) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundle) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesBundle) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundle) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesBundle) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesBundle) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundle) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundle) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesBundle) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundle) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesBundle) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsWorkspacesBundle) ResetBundleId() {
	_jsii_.InvokeVoid(
		d,
		"resetBundleId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesBundle) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundle) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesBundle) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesBundle) RootStorage(index *string) DataAwsWorkspacesBundleRootStorage {
	var returns DataAwsWorkspacesBundleRootStorage

	_jsii_.Invoke(
		d,
		"rootStorage",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsWorkspacesBundle) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsWorkspacesBundle) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsWorkspacesBundle) ToString() *string {
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
func (d *jsiiProxy_DataAwsWorkspacesBundle) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsWorkspacesBundle) UserStorage(index *string) DataAwsWorkspacesBundleUserStorage {
	var returns DataAwsWorkspacesBundleUserStorage

	_jsii_.Invoke(
		d,
		"userStorage",
		[]interface{}{index},
		&returns,
	)

	return returns
}

type DataAwsWorkspacesBundleComputeType interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Name() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsWorkspacesBundleComputeType
type jsiiProxy_DataAwsWorkspacesBundleComputeType struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsWorkspacesBundleComputeType) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleComputeType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleComputeType) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleComputeType) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleComputeType) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsWorkspacesBundleComputeType(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsWorkspacesBundleComputeType {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesBundleComputeType{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundleComputeType",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsWorkspacesBundleComputeType_Override(d DataAwsWorkspacesBundleComputeType, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundleComputeType",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleComputeType) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleComputeType) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleComputeType) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleComputeType) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleComputeType) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// AWS WorkSpaces.
type DataAwsWorkspacesBundleConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_bundle#bundle_id DataAwsWorkspacesBundle#bundle_id}.
	BundleId *string `json:"bundleId" yaml:"bundleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_bundle#name DataAwsWorkspacesBundle#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_bundle#owner DataAwsWorkspacesBundle#owner}.
	Owner *string `json:"owner" yaml:"owner"`
}

type DataAwsWorkspacesBundleRootStorage interface {
	cdktf.ComplexComputedList
	Capacity() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsWorkspacesBundleRootStorage
type jsiiProxy_DataAwsWorkspacesBundleRootStorage struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsWorkspacesBundleRootStorage) Capacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleRootStorage) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleRootStorage) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleRootStorage) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleRootStorage) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsWorkspacesBundleRootStorage(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsWorkspacesBundleRootStorage {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesBundleRootStorage{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundleRootStorage",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsWorkspacesBundleRootStorage_Override(d DataAwsWorkspacesBundleRootStorage, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundleRootStorage",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleRootStorage) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleRootStorage) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleRootStorage) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleRootStorage) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleRootStorage) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsWorkspacesBundleUserStorage interface {
	cdktf.ComplexComputedList
	Capacity() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsWorkspacesBundleUserStorage
type jsiiProxy_DataAwsWorkspacesBundleUserStorage struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsWorkspacesBundleUserStorage) Capacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleUserStorage) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleUserStorage) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleUserStorage) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesBundleUserStorage) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsWorkspacesBundleUserStorage(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsWorkspacesBundleUserStorage {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesBundleUserStorage{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundleUserStorage",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsWorkspacesBundleUserStorage_Override(d DataAwsWorkspacesBundleUserStorage, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesBundleUserStorage",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleUserStorage) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleUserStorage) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleUserStorage) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesBundleUserStorage) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesBundleUserStorage) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/workspaces_directory aws_workspaces_directory}.
type DataAwsWorkspacesDirectory interface {
	cdktf.TerraformDataSource
	Alias() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CustomerUserName() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	DirectoryName() *string
	DirectoryType() *string
	DnsIpAddresses() *[]*string
	Fqn() *string
	FriendlyUniqueId() *string
	IamRoleId() *string
	Id() *string
	IpGroupIds() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RegistrationCode() *string
	SubnetIds() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	WorkspaceSecurityGroupId() *string
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
	ResetTags()
	SelfServicePermissions(index *string) DataAwsWorkspacesDirectorySelfServicePermissions
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
	WorkspaceAccessProperties(index *string) DataAwsWorkspacesDirectoryWorkspaceAccessProperties
	WorkspaceCreationProperties(index *string) DataAwsWorkspacesDirectoryWorkspaceCreationProperties
}

// The jsii proxy struct for DataAwsWorkspacesDirectory
type jsiiProxy_DataAwsWorkspacesDirectory struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) CustomerUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) DirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) DnsIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) IamRoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) IpGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) RegistrationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) WorkspaceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupId",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/workspaces_directory aws_workspaces_directory} Data Source.
func NewDataAwsWorkspacesDirectory(scope constructs.Construct, id *string, config *DataAwsWorkspacesDirectoryConfig) DataAwsWorkspacesDirectory {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesDirectory{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectory",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/workspaces_directory aws_workspaces_directory} Data Source.
func NewDataAwsWorkspacesDirectory_Override(d DataAwsWorkspacesDirectory, scope constructs.Construct, id *string, config *DataAwsWorkspacesDirectoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectory",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) SetDirectoryId(val *string) {
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectory) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsWorkspacesDirectory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsWorkspacesDirectory_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectory",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectory) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectory) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectory) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectory) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectory) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectory) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectory) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectory) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectory) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectory) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectory) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectory) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectory) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesDirectory) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesDirectory) SelfServicePermissions(index *string) DataAwsWorkspacesDirectorySelfServicePermissions {
	var returns DataAwsWorkspacesDirectorySelfServicePermissions

	_jsii_.Invoke(
		d,
		"selfServicePermissions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsWorkspacesDirectory) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectory) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectory) ToString() *string {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectory) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsWorkspacesDirectory) WorkspaceAccessProperties(index *string) DataAwsWorkspacesDirectoryWorkspaceAccessProperties {
	var returns DataAwsWorkspacesDirectoryWorkspaceAccessProperties

	_jsii_.Invoke(
		d,
		"workspaceAccessProperties",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsWorkspacesDirectory) WorkspaceCreationProperties(index *string) DataAwsWorkspacesDirectoryWorkspaceCreationProperties {
	var returns DataAwsWorkspacesDirectoryWorkspaceCreationProperties

	_jsii_.Invoke(
		d,
		"workspaceCreationProperties",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// AWS WorkSpaces.
type DataAwsWorkspacesDirectoryConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_directory#directory_id DataAwsWorkspacesDirectory#directory_id}.
	DirectoryId *string `json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_directory#tags DataAwsWorkspacesDirectory#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type DataAwsWorkspacesDirectorySelfServicePermissions interface {
	cdktf.ComplexComputedList
	ChangeComputeType() cdktf.IResolvable
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	IncreaseVolumeSize() cdktf.IResolvable
	RebuildWorkspace() cdktf.IResolvable
	RestartWorkspace() cdktf.IResolvable
	SwitchRunningMode() cdktf.IResolvable
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsWorkspacesDirectorySelfServicePermissions
type jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) ChangeComputeType() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"changeComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) IncreaseVolumeSize() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"increaseVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) RebuildWorkspace() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rebuildWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) RestartWorkspace() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"restartWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) SwitchRunningMode() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"switchRunningMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsWorkspacesDirectorySelfServicePermissions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsWorkspacesDirectorySelfServicePermissions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectorySelfServicePermissions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsWorkspacesDirectorySelfServicePermissions_Override(d DataAwsWorkspacesDirectorySelfServicePermissions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectorySelfServicePermissions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectorySelfServicePermissions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsWorkspacesDirectoryWorkspaceAccessProperties interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DeviceTypeAndroid() *string
	DeviceTypeChromeos() *string
	DeviceTypeIos() *string
	DeviceTypeLinux() *string
	DeviceTypeOsx() *string
	DeviceTypeWeb() *string
	DeviceTypeWindows() *string
	DeviceTypeZeroclient() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsWorkspacesDirectoryWorkspaceAccessProperties
type jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) DeviceTypeAndroid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeAndroid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) DeviceTypeChromeos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeChromeos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) DeviceTypeIos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) DeviceTypeLinux() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) DeviceTypeOsx() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeOsx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) DeviceTypeWeb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) DeviceTypeWindows() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) DeviceTypeZeroclient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeZeroclient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsWorkspacesDirectoryWorkspaceAccessProperties(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsWorkspacesDirectoryWorkspaceAccessProperties {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectoryWorkspaceAccessProperties",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsWorkspacesDirectoryWorkspaceAccessProperties_Override(d DataAwsWorkspacesDirectoryWorkspaceAccessProperties, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectoryWorkspaceAccessProperties",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceAccessProperties) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsWorkspacesDirectoryWorkspaceCreationProperties interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	CustomSecurityGroupId() *string
	DefaultOu() *string
	EnableInternetAccess() cdktf.IResolvable
	EnableMaintenanceMode() cdktf.IResolvable
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserEnabledAsLocalAdministrator() cdktf.IResolvable
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsWorkspacesDirectoryWorkspaceCreationProperties
type jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) CustomSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) DefaultOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) EnableInternetAccess() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) EnableMaintenanceMode() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableMaintenanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) UserEnabledAsLocalAdministrator() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministrator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsWorkspacesDirectoryWorkspaceCreationProperties(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsWorkspacesDirectoryWorkspaceCreationProperties {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectoryWorkspaceCreationProperties",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsWorkspacesDirectoryWorkspaceCreationProperties_Override(d DataAwsWorkspacesDirectoryWorkspaceCreationProperties, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesDirectoryWorkspaceCreationProperties",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationProperties) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/workspaces_image aws_workspaces_image}.
type DataAwsWorkspacesImage interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	Node() constructs.Node
	OperatingSystemType() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequiredTenancy() *string
	State() *string
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
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsWorkspacesImage
type jsiiProxy_DataAwsWorkspacesImage struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsWorkspacesImage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) OperatingSystemType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatingSystemType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) RequiredTenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredTenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesImage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/workspaces_image aws_workspaces_image} Data Source.
func NewDataAwsWorkspacesImage(scope constructs.Construct, id *string, config *DataAwsWorkspacesImageConfig) DataAwsWorkspacesImage {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesImage{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesImage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/workspaces_image aws_workspaces_image} Data Source.
func NewDataAwsWorkspacesImage_Override(d DataAwsWorkspacesImage, scope constructs.Construct, id *string, config *DataAwsWorkspacesImageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesImage",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesImage) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesImage) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesImage) SetImageId(val *string) {
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesImage) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesImage) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsWorkspacesImage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.workspaces.DataAwsWorkspacesImage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsWorkspacesImage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.workspaces.DataAwsWorkspacesImage",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesImage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesImage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesImage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesImage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesImage) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesImage) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesImage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesImage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesImage) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesImage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesImage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesImage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesImage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesImage) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsWorkspacesImage) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsWorkspacesImage) ToString() *string {
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
func (d *jsiiProxy_DataAwsWorkspacesImage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS WorkSpaces.
type DataAwsWorkspacesImageConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_image#image_id DataAwsWorkspacesImage#image_id}.
	ImageId *string `json:"imageId" yaml:"imageId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/workspaces_workspace aws_workspaces_workspace}.
type DataAwsWorkspacesWorkspace interface {
	cdktf.TerraformDataSource
	BundleId() *string
	CdktfStack() cdktf.TerraformStack
	ComputerName() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IpAddress() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RootVolumeEncryptionEnabled() cdktf.IResolvable
	State() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
	UserVolumeEncryptionEnabled() cdktf.IResolvable
	VolumeEncryptionKey() *string
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	ResetDirectoryId()
	ResetOverrideLogicalId()
	ResetTags()
	ResetUserName()
	ResetWorkspaceId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
	WorkspaceProperties(index *string) DataAwsWorkspacesWorkspaceWorkspaceProperties
}

// The jsii proxy struct for DataAwsWorkspacesWorkspace
type jsiiProxy_DataAwsWorkspacesWorkspace struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) ComputerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) RootVolumeEncryptionEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rootVolumeEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) UserVolumeEncryptionEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"userVolumeEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) VolumeEncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/workspaces_workspace aws_workspaces_workspace} Data Source.
func NewDataAwsWorkspacesWorkspace(scope constructs.Construct, id *string, config *DataAwsWorkspacesWorkspaceConfig) DataAwsWorkspacesWorkspace {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesWorkspace{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/workspaces_workspace aws_workspaces_workspace} Data Source.
func NewDataAwsWorkspacesWorkspace_Override(d DataAwsWorkspacesWorkspace, scope constructs.Construct, id *string, config *DataAwsWorkspacesWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesWorkspace",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) SetDirectoryId(val *string) {
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspace) SetWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsWorkspacesWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.workspaces.DataAwsWorkspacesWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsWorkspacesWorkspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.workspaces.DataAwsWorkspacesWorkspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsWorkspacesWorkspace) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		d,
		"resetDirectoryId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesWorkspace) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesWorkspace) ResetUserName() {
	_jsii_.InvokeVoid(
		d,
		"resetUserName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesWorkspace) ResetWorkspaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkspaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsWorkspacesWorkspace) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) ToString() *string {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsWorkspacesWorkspace) WorkspaceProperties(index *string) DataAwsWorkspacesWorkspaceWorkspaceProperties {
	var returns DataAwsWorkspacesWorkspaceWorkspaceProperties

	_jsii_.Invoke(
		d,
		"workspaceProperties",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// AWS WorkSpaces.
type DataAwsWorkspacesWorkspaceConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_workspace#directory_id DataAwsWorkspacesWorkspace#directory_id}.
	DirectoryId *string `json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_workspace#tags DataAwsWorkspacesWorkspace#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_workspace#user_name DataAwsWorkspacesWorkspace#user_name}.
	UserName *string `json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/workspaces_workspace#workspace_id DataAwsWorkspacesWorkspace#workspace_id}.
	WorkspaceId *string `json:"workspaceId" yaml:"workspaceId"`
}

type DataAwsWorkspacesWorkspaceWorkspaceProperties interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ComputeTypeName() *string
	RootVolumeSizeGib() *float64
	RunningMode() *string
	RunningModeAutoStopTimeoutInMinutes() *float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserVolumeSizeGib() *float64
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	GetStringAttribute(terraformAttribute *string) *string
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsWorkspacesWorkspaceWorkspaceProperties
type jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) ComputeTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) RootVolumeSizeGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) RunningMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runningMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) RunningModeAutoStopTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningModeAutoStopTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) UserVolumeSizeGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userVolumeSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsWorkspacesWorkspaceWorkspaceProperties(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsWorkspacesWorkspaceWorkspaceProperties {
	_init_.Initialize()

	j := jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesWorkspaceWorkspaceProperties",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsWorkspacesWorkspaceWorkspaceProperties_Override(d DataAwsWorkspacesWorkspaceWorkspaceProperties, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.DataAwsWorkspacesWorkspaceWorkspaceProperties",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspaceProperties) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory aws_workspaces_directory}.
type WorkspacesDirectory interface {
	cdktf.TerraformResource
	Alias() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CustomerUserName() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	DirectoryName() *string
	DirectoryType() *string
	DnsIpAddresses() *[]*string
	Fqn() *string
	FriendlyUniqueId() *string
	IamRoleId() *string
	Id() *string
	IpGroupIds() *[]*string
	SetIpGroupIds(val *[]*string)
	IpGroupIdsInput() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RegistrationCode() *string
	SelfServicePermissions() WorkspacesDirectorySelfServicePermissionsOutputReference
	SelfServicePermissionsInput() *WorkspacesDirectorySelfServicePermissions
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	WorkspaceAccessProperties() WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference
	WorkspaceAccessPropertiesInput() *WorkspacesDirectoryWorkspaceAccessProperties
	WorkspaceCreationProperties() WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference
	WorkspaceCreationPropertiesInput() *WorkspacesDirectoryWorkspaceCreationProperties
	WorkspaceSecurityGroupId() *string
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
	PutSelfServicePermissions(value *WorkspacesDirectorySelfServicePermissions)
	PutWorkspaceAccessProperties(value *WorkspacesDirectoryWorkspaceAccessProperties)
	PutWorkspaceCreationProperties(value *WorkspacesDirectoryWorkspaceCreationProperties)
	ResetIpGroupIds()
	ResetOverrideLogicalId()
	ResetSelfServicePermissions()
	ResetSubnetIds()
	ResetTags()
	ResetTagsAll()
	ResetWorkspaceAccessProperties()
	ResetWorkspaceCreationProperties()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for WorkspacesDirectory
type jsiiProxy_WorkspacesDirectory struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorkspacesDirectory) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) CustomerUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) DirectoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) DnsIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) IamRoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) IpGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) IpGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) RegistrationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) SelfServicePermissions() WorkspacesDirectorySelfServicePermissionsOutputReference {
	var returns WorkspacesDirectorySelfServicePermissionsOutputReference
	_jsii_.Get(
		j,
		"selfServicePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) SelfServicePermissionsInput() *WorkspacesDirectorySelfServicePermissions {
	var returns *WorkspacesDirectorySelfServicePermissions
	_jsii_.Get(
		j,
		"selfServicePermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) WorkspaceAccessProperties() WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference {
	var returns WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference
	_jsii_.Get(
		j,
		"workspaceAccessProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) WorkspaceAccessPropertiesInput() *WorkspacesDirectoryWorkspaceAccessProperties {
	var returns *WorkspacesDirectoryWorkspaceAccessProperties
	_jsii_.Get(
		j,
		"workspaceAccessPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) WorkspaceCreationProperties() WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference {
	var returns WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference
	_jsii_.Get(
		j,
		"workspaceCreationProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) WorkspaceCreationPropertiesInput() *WorkspacesDirectoryWorkspaceCreationProperties {
	var returns *WorkspacesDirectoryWorkspaceCreationProperties
	_jsii_.Get(
		j,
		"workspaceCreationPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectory) WorkspaceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupId",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory aws_workspaces_directory} Resource.
func NewWorkspacesDirectory(scope constructs.Construct, id *string, config *WorkspacesDirectoryConfig) WorkspacesDirectory {
	_init_.Initialize()

	j := jsiiProxy_WorkspacesDirectory{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesDirectory",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory aws_workspaces_directory} Resource.
func NewWorkspacesDirectory_Override(w WorkspacesDirectory, scope constructs.Construct, id *string, config *WorkspacesDirectoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesDirectory",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkspacesDirectory) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectory) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectory) SetDirectoryId(val *string) {
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectory) SetIpGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"ipGroupIds",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectory) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectory) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectory) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectory) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectory) SetTagsAll(val *map[string]*string) {
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
func WorkspacesDirectory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.workspaces.WorkspacesDirectory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkspacesDirectory_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.workspaces.WorkspacesDirectory",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkspacesDirectory) PutSelfServicePermissions(value *WorkspacesDirectorySelfServicePermissions) {
	_jsii_.InvokeVoid(
		w,
		"putSelfServicePermissions",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesDirectory) PutWorkspaceAccessProperties(value *WorkspacesDirectoryWorkspaceAccessProperties) {
	_jsii_.InvokeVoid(
		w,
		"putWorkspaceAccessProperties",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesDirectory) PutWorkspaceCreationProperties(value *WorkspacesDirectoryWorkspaceCreationProperties) {
	_jsii_.InvokeVoid(
		w,
		"putWorkspaceCreationProperties",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesDirectory) ResetIpGroupIds() {
	_jsii_.InvokeVoid(
		w,
		"resetIpGroupIds",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectory) ResetSelfServicePermissions() {
	_jsii_.InvokeVoid(
		w,
		"resetSelfServicePermissions",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectory) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		w,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectory) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectory) ResetTagsAll() {
	_jsii_.InvokeVoid(
		w,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectory) ResetWorkspaceAccessProperties() {
	_jsii_.InvokeVoid(
		w,
		"resetWorkspaceAccessProperties",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectory) ResetWorkspaceCreationProperties() {
	_jsii_.InvokeVoid(
		w,
		"resetWorkspaceCreationProperties",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectory) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WorkspacesDirectory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (w *jsiiProxy_WorkspacesDirectory) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS WorkSpaces.
type WorkspacesDirectoryConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#directory_id WorkspacesDirectory#directory_id}.
	DirectoryId *string `json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#ip_group_ids WorkspacesDirectory#ip_group_ids}.
	IpGroupIds *[]*string `json:"ipGroupIds" yaml:"ipGroupIds"`
	// self_service_permissions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#self_service_permissions WorkspacesDirectory#self_service_permissions}
	SelfServicePermissions *WorkspacesDirectorySelfServicePermissions `json:"selfServicePermissions" yaml:"selfServicePermissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#subnet_ids WorkspacesDirectory#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#tags WorkspacesDirectory#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#tags_all WorkspacesDirectory#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// workspace_access_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#workspace_access_properties WorkspacesDirectory#workspace_access_properties}
	WorkspaceAccessProperties *WorkspacesDirectoryWorkspaceAccessProperties `json:"workspaceAccessProperties" yaml:"workspaceAccessProperties"`
	// workspace_creation_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#workspace_creation_properties WorkspacesDirectory#workspace_creation_properties}
	WorkspaceCreationProperties *WorkspacesDirectoryWorkspaceCreationProperties `json:"workspaceCreationProperties" yaml:"workspaceCreationProperties"`
}

type WorkspacesDirectorySelfServicePermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#change_compute_type WorkspacesDirectory#change_compute_type}.
	ChangeComputeType interface{} `json:"changeComputeType" yaml:"changeComputeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#increase_volume_size WorkspacesDirectory#increase_volume_size}.
	IncreaseVolumeSize interface{} `json:"increaseVolumeSize" yaml:"increaseVolumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#rebuild_workspace WorkspacesDirectory#rebuild_workspace}.
	RebuildWorkspace interface{} `json:"rebuildWorkspace" yaml:"rebuildWorkspace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#restart_workspace WorkspacesDirectory#restart_workspace}.
	RestartWorkspace interface{} `json:"restartWorkspace" yaml:"restartWorkspace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#switch_running_mode WorkspacesDirectory#switch_running_mode}.
	SwitchRunningMode interface{} `json:"switchRunningMode" yaml:"switchRunningMode"`
}

type WorkspacesDirectorySelfServicePermissionsOutputReference interface {
	cdktf.ComplexObject
	ChangeComputeType() interface{}
	SetChangeComputeType(val interface{})
	ChangeComputeTypeInput() interface{}
	IncreaseVolumeSize() interface{}
	SetIncreaseVolumeSize(val interface{})
	IncreaseVolumeSizeInput() interface{}
	InternalValue() *WorkspacesDirectorySelfServicePermissions
	SetInternalValue(val *WorkspacesDirectorySelfServicePermissions)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RebuildWorkspace() interface{}
	SetRebuildWorkspace(val interface{})
	RebuildWorkspaceInput() interface{}
	RestartWorkspace() interface{}
	SetRestartWorkspace(val interface{})
	RestartWorkspaceInput() interface{}
	SwitchRunningMode() interface{}
	SetSwitchRunningMode(val interface{})
	SwitchRunningModeInput() interface{}
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
	ResetChangeComputeType()
	ResetIncreaseVolumeSize()
	ResetRebuildWorkspace()
	ResetRestartWorkspace()
	ResetSwitchRunningMode()
}

// The jsii proxy struct for WorkspacesDirectorySelfServicePermissionsOutputReference
type jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) ChangeComputeType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) ChangeComputeTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeComputeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) IncreaseVolumeSize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) IncreaseVolumeSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) InternalValue() *WorkspacesDirectorySelfServicePermissions {
	var returns *WorkspacesDirectorySelfServicePermissions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) RebuildWorkspace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rebuildWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) RebuildWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rebuildWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) RestartWorkspace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) RestartWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SwitchRunningMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchRunningMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SwitchRunningModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchRunningModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewWorkspacesDirectorySelfServicePermissionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) WorkspacesDirectorySelfServicePermissionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesDirectorySelfServicePermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewWorkspacesDirectorySelfServicePermissionsOutputReference_Override(w WorkspacesDirectorySelfServicePermissionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesDirectorySelfServicePermissionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		w,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SetChangeComputeType(val interface{}) {
	_jsii_.Set(
		j,
		"changeComputeType",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SetIncreaseVolumeSize(val interface{}) {
	_jsii_.Set(
		j,
		"increaseVolumeSize",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SetInternalValue(val *WorkspacesDirectorySelfServicePermissions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SetRebuildWorkspace(val interface{}) {
	_jsii_.Set(
		j,
		"rebuildWorkspace",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SetRestartWorkspace(val interface{}) {
	_jsii_.Set(
		j,
		"restartWorkspace",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SetSwitchRunningMode(val interface{}) {
	_jsii_.Set(
		j,
		"switchRunningMode",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) ResetChangeComputeType() {
	_jsii_.InvokeVoid(
		w,
		"resetChangeComputeType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) ResetIncreaseVolumeSize() {
	_jsii_.InvokeVoid(
		w,
		"resetIncreaseVolumeSize",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) ResetRebuildWorkspace() {
	_jsii_.InvokeVoid(
		w,
		"resetRebuildWorkspace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) ResetRestartWorkspace() {
	_jsii_.InvokeVoid(
		w,
		"resetRestartWorkspace",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectorySelfServicePermissionsOutputReference) ResetSwitchRunningMode() {
	_jsii_.InvokeVoid(
		w,
		"resetSwitchRunningMode",
		nil, // no parameters
	)
}

type WorkspacesDirectoryWorkspaceAccessProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#device_type_android WorkspacesDirectory#device_type_android}.
	DeviceTypeAndroid *string `json:"deviceTypeAndroid" yaml:"deviceTypeAndroid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#device_type_chromeos WorkspacesDirectory#device_type_chromeos}.
	DeviceTypeChromeos *string `json:"deviceTypeChromeos" yaml:"deviceTypeChromeos"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#device_type_ios WorkspacesDirectory#device_type_ios}.
	DeviceTypeIos *string `json:"deviceTypeIos" yaml:"deviceTypeIos"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#device_type_linux WorkspacesDirectory#device_type_linux}.
	DeviceTypeLinux *string `json:"deviceTypeLinux" yaml:"deviceTypeLinux"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#device_type_osx WorkspacesDirectory#device_type_osx}.
	DeviceTypeOsx *string `json:"deviceTypeOsx" yaml:"deviceTypeOsx"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#device_type_web WorkspacesDirectory#device_type_web}.
	DeviceTypeWeb *string `json:"deviceTypeWeb" yaml:"deviceTypeWeb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#device_type_windows WorkspacesDirectory#device_type_windows}.
	DeviceTypeWindows *string `json:"deviceTypeWindows" yaml:"deviceTypeWindows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#device_type_zeroclient WorkspacesDirectory#device_type_zeroclient}.
	DeviceTypeZeroclient *string `json:"deviceTypeZeroclient" yaml:"deviceTypeZeroclient"`
}

type WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference interface {
	cdktf.ComplexObject
	DeviceTypeAndroid() *string
	SetDeviceTypeAndroid(val *string)
	DeviceTypeAndroidInput() *string
	DeviceTypeChromeos() *string
	SetDeviceTypeChromeos(val *string)
	DeviceTypeChromeosInput() *string
	DeviceTypeIos() *string
	SetDeviceTypeIos(val *string)
	DeviceTypeIosInput() *string
	DeviceTypeLinux() *string
	SetDeviceTypeLinux(val *string)
	DeviceTypeLinuxInput() *string
	DeviceTypeOsx() *string
	SetDeviceTypeOsx(val *string)
	DeviceTypeOsxInput() *string
	DeviceTypeWeb() *string
	SetDeviceTypeWeb(val *string)
	DeviceTypeWebInput() *string
	DeviceTypeWindows() *string
	SetDeviceTypeWindows(val *string)
	DeviceTypeWindowsInput() *string
	DeviceTypeZeroclient() *string
	SetDeviceTypeZeroclient(val *string)
	DeviceTypeZeroclientInput() *string
	InternalValue() *WorkspacesDirectoryWorkspaceAccessProperties
	SetInternalValue(val *WorkspacesDirectoryWorkspaceAccessProperties)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
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
	ResetDeviceTypeAndroid()
	ResetDeviceTypeChromeos()
	ResetDeviceTypeIos()
	ResetDeviceTypeLinux()
	ResetDeviceTypeOsx()
	ResetDeviceTypeWeb()
	ResetDeviceTypeWindows()
	ResetDeviceTypeZeroclient()
}

// The jsii proxy struct for WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference
type jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeAndroid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeAndroid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeAndroidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeAndroidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeChromeos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeChromeos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeChromeosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeChromeosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeIos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeIosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeLinux() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeLinuxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeLinuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeOsx() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeOsx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeOsxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeOsxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeWeb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeWebInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeWindows() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeWindowsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeZeroclient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeZeroclient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) DeviceTypeZeroclientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeZeroclientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) InternalValue() *WorkspacesDirectoryWorkspaceAccessProperties {
	var returns *WorkspacesDirectoryWorkspaceAccessProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewWorkspacesDirectoryWorkspaceAccessPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewWorkspacesDirectoryWorkspaceAccessPropertiesOutputReference_Override(w WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		w,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetDeviceTypeAndroid(val *string) {
	_jsii_.Set(
		j,
		"deviceTypeAndroid",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetDeviceTypeChromeos(val *string) {
	_jsii_.Set(
		j,
		"deviceTypeChromeos",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetDeviceTypeIos(val *string) {
	_jsii_.Set(
		j,
		"deviceTypeIos",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetDeviceTypeLinux(val *string) {
	_jsii_.Set(
		j,
		"deviceTypeLinux",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetDeviceTypeOsx(val *string) {
	_jsii_.Set(
		j,
		"deviceTypeOsx",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetDeviceTypeWeb(val *string) {
	_jsii_.Set(
		j,
		"deviceTypeWeb",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetDeviceTypeWindows(val *string) {
	_jsii_.Set(
		j,
		"deviceTypeWindows",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetDeviceTypeZeroclient(val *string) {
	_jsii_.Set(
		j,
		"deviceTypeZeroclient",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetInternalValue(val *WorkspacesDirectoryWorkspaceAccessProperties) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeAndroid() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeAndroid",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeChromeos() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeChromeos",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeIos() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeIos",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeLinux() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeLinux",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeOsx() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeOsx",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeWeb() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeWeb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeWindows() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeWindows",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceAccessPropertiesOutputReference) ResetDeviceTypeZeroclient() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceTypeZeroclient",
		nil, // no parameters
	)
}

type WorkspacesDirectoryWorkspaceCreationProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#custom_security_group_id WorkspacesDirectory#custom_security_group_id}.
	CustomSecurityGroupId *string `json:"customSecurityGroupId" yaml:"customSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#default_ou WorkspacesDirectory#default_ou}.
	DefaultOu *string `json:"defaultOu" yaml:"defaultOu"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#enable_internet_access WorkspacesDirectory#enable_internet_access}.
	EnableInternetAccess interface{} `json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#enable_maintenance_mode WorkspacesDirectory#enable_maintenance_mode}.
	EnableMaintenanceMode interface{} `json:"enableMaintenanceMode" yaml:"enableMaintenanceMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_directory#user_enabled_as_local_administrator WorkspacesDirectory#user_enabled_as_local_administrator}.
	UserEnabledAsLocalAdministrator interface{} `json:"userEnabledAsLocalAdministrator" yaml:"userEnabledAsLocalAdministrator"`
}

type WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference interface {
	cdktf.ComplexObject
	CustomSecurityGroupId() *string
	SetCustomSecurityGroupId(val *string)
	CustomSecurityGroupIdInput() *string
	DefaultOu() *string
	SetDefaultOu(val *string)
	DefaultOuInput() *string
	EnableInternetAccess() interface{}
	SetEnableInternetAccess(val interface{})
	EnableInternetAccessInput() interface{}
	EnableMaintenanceMode() interface{}
	SetEnableMaintenanceMode(val interface{})
	EnableMaintenanceModeInput() interface{}
	InternalValue() *WorkspacesDirectoryWorkspaceCreationProperties
	SetInternalValue(val *WorkspacesDirectoryWorkspaceCreationProperties)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserEnabledAsLocalAdministrator() interface{}
	SetUserEnabledAsLocalAdministrator(val interface{})
	UserEnabledAsLocalAdministratorInput() interface{}
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
	ResetCustomSecurityGroupId()
	ResetDefaultOu()
	ResetEnableInternetAccess()
	ResetEnableMaintenanceMode()
	ResetUserEnabledAsLocalAdministrator()
}

// The jsii proxy struct for WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference
type jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) CustomSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) CustomSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) DefaultOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) DefaultOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) EnableInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) EnableInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) EnableMaintenanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) EnableMaintenanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) InternalValue() *WorkspacesDirectoryWorkspaceCreationProperties {
	var returns *WorkspacesDirectoryWorkspaceCreationProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) UserEnabledAsLocalAdministrator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministrator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) UserEnabledAsLocalAdministratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministratorInput",
		&returns,
	)
	return returns
}

func NewWorkspacesDirectoryWorkspaceCreationPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewWorkspacesDirectoryWorkspaceCreationPropertiesOutputReference_Override(w WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		w,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) SetCustomSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"customSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) SetDefaultOu(val *string) {
	_jsii_.Set(
		j,
		"defaultOu",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) SetEnableInternetAccess(val interface{}) {
	_jsii_.Set(
		j,
		"enableInternetAccess",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) SetEnableMaintenanceMode(val interface{}) {
	_jsii_.Set(
		j,
		"enableMaintenanceMode",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) SetInternalValue(val *WorkspacesDirectoryWorkspaceCreationProperties) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) SetUserEnabledAsLocalAdministrator(val interface{}) {
	_jsii_.Set(
		j,
		"userEnabledAsLocalAdministrator",
		val,
	)
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetCustomSecurityGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomSecurityGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetDefaultOu() {
	_jsii_.InvokeVoid(
		w,
		"resetDefaultOu",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetEnableInternetAccess() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableInternetAccess",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetEnableMaintenanceMode() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableMaintenanceMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesDirectoryWorkspaceCreationPropertiesOutputReference) ResetUserEnabledAsLocalAdministrator() {
	_jsii_.InvokeVoid(
		w,
		"resetUserEnabledAsLocalAdministrator",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group aws_workspaces_ip_group}.
type WorkspacesIpGroup interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	Rules() interface{}
	SetRules(val interface{})
	RulesInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
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
	ResetDescription()
	ResetOverrideLogicalId()
	ResetRules()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for WorkspacesIpGroup
type jsiiProxy_WorkspacesIpGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorkspacesIpGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Rules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesIpGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group aws_workspaces_ip_group} Resource.
func NewWorkspacesIpGroup(scope constructs.Construct, id *string, config *WorkspacesIpGroupConfig) WorkspacesIpGroup {
	_init_.Initialize()

	j := jsiiProxy_WorkspacesIpGroup{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesIpGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group aws_workspaces_ip_group} Resource.
func NewWorkspacesIpGroup_Override(w WorkspacesIpGroup, scope constructs.Construct, id *string, config *WorkspacesIpGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesIpGroup",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkspacesIpGroup) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkspacesIpGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkspacesIpGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_WorkspacesIpGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkspacesIpGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkspacesIpGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkspacesIpGroup) SetRules(val interface{}) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_WorkspacesIpGroup) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WorkspacesIpGroup) SetTagsAll(val *map[string]*string) {
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
func WorkspacesIpGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.workspaces.WorkspacesIpGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkspacesIpGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.workspaces.WorkspacesIpGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkspacesIpGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesIpGroup) ResetRules() {
	_jsii_.InvokeVoid(
		w,
		"resetRules",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesIpGroup) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesIpGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		w,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesIpGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WorkspacesIpGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (w *jsiiProxy_WorkspacesIpGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS WorkSpaces.
type WorkspacesIpGroupConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group#name WorkspacesIpGroup#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group#description WorkspacesIpGroup#description}.
	Description *string `json:"description" yaml:"description"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group#rules WorkspacesIpGroup#rules}
	Rules interface{} `json:"rules" yaml:"rules"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group#tags WorkspacesIpGroup#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group#tags_all WorkspacesIpGroup#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

type WorkspacesIpGroupRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group#source WorkspacesIpGroup#source}.
	Source *string `json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_ip_group#description WorkspacesIpGroup#description}.
	Description *string `json:"description" yaml:"description"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace aws_workspaces_workspace}.
type WorkspacesWorkspace interface {
	cdktf.TerraformResource
	BundleId() *string
	SetBundleId(val *string)
	BundleIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ComputerName() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DirectoryId() *string
	SetDirectoryId(val *string)
	DirectoryIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IpAddress() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RootVolumeEncryptionEnabled() interface{}
	SetRootVolumeEncryptionEnabled(val interface{})
	RootVolumeEncryptionEnabledInput() interface{}
	State() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() WorkspacesWorkspaceTimeoutsOutputReference
	TimeoutsInput() *WorkspacesWorkspaceTimeouts
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
	UserVolumeEncryptionEnabled() interface{}
	SetUserVolumeEncryptionEnabled(val interface{})
	UserVolumeEncryptionEnabledInput() interface{}
	VolumeEncryptionKey() *string
	SetVolumeEncryptionKey(val *string)
	VolumeEncryptionKeyInput() *string
	WorkspaceProperties() WorkspacesWorkspaceWorkspacePropertiesOutputReference
	WorkspacePropertiesInput() *WorkspacesWorkspaceWorkspaceProperties
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
	PutTimeouts(value *WorkspacesWorkspaceTimeouts)
	PutWorkspaceProperties(value *WorkspacesWorkspaceWorkspaceProperties)
	ResetOverrideLogicalId()
	ResetRootVolumeEncryptionEnabled()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetUserVolumeEncryptionEnabled()
	ResetVolumeEncryptionKey()
	ResetWorkspaceProperties()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for WorkspacesWorkspace
type jsiiProxy_WorkspacesWorkspace struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WorkspacesWorkspace) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) BundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) ComputerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) RootVolumeEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootVolumeEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) RootVolumeEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootVolumeEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) Timeouts() WorkspacesWorkspaceTimeoutsOutputReference {
	var returns WorkspacesWorkspaceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) TimeoutsInput() *WorkspacesWorkspaceTimeouts {
	var returns *WorkspacesWorkspaceTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) UserVolumeEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userVolumeEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) UserVolumeEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userVolumeEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) VolumeEncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) VolumeEncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) WorkspaceProperties() WorkspacesWorkspaceWorkspacePropertiesOutputReference {
	var returns WorkspacesWorkspaceWorkspacePropertiesOutputReference
	_jsii_.Get(
		j,
		"workspaceProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspace) WorkspacePropertiesInput() *WorkspacesWorkspaceWorkspaceProperties {
	var returns *WorkspacesWorkspaceWorkspaceProperties
	_jsii_.Get(
		j,
		"workspacePropertiesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace aws_workspaces_workspace} Resource.
func NewWorkspacesWorkspace(scope constructs.Construct, id *string, config *WorkspacesWorkspaceConfig) WorkspacesWorkspace {
	_init_.Initialize()

	j := jsiiProxy_WorkspacesWorkspace{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesWorkspace",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace aws_workspaces_workspace} Resource.
func NewWorkspacesWorkspace_Override(w WorkspacesWorkspace, scope constructs.Construct, id *string, config *WorkspacesWorkspaceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesWorkspace",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetDirectoryId(val *string) {
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetRootVolumeEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"rootVolumeEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetUserVolumeEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"userVolumeEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspace) SetVolumeEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"volumeEncryptionKey",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func WorkspacesWorkspace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.workspaces.WorkspacesWorkspace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkspacesWorkspace_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.workspaces.WorkspacesWorkspace",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) PutTimeouts(value *WorkspacesWorkspaceTimeouts) {
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) PutWorkspaceProperties(value *WorkspacesWorkspaceWorkspaceProperties) {
	_jsii_.InvokeVoid(
		w,
		"putWorkspaceProperties",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) ResetRootVolumeEncryptionEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetRootVolumeEncryptionEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) ResetTagsAll() {
	_jsii_.InvokeVoid(
		w,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) ResetUserVolumeEncryptionEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetUserVolumeEncryptionEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) ResetVolumeEncryptionKey() {
	_jsii_.InvokeVoid(
		w,
		"resetVolumeEncryptionKey",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) ResetWorkspaceProperties() {
	_jsii_.InvokeVoid(
		w,
		"resetWorkspaceProperties",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspace) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (w *jsiiProxy_WorkspacesWorkspace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (w *jsiiProxy_WorkspacesWorkspace) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS WorkSpaces.
type WorkspacesWorkspaceConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#bundle_id WorkspacesWorkspace#bundle_id}.
	BundleId *string `json:"bundleId" yaml:"bundleId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#directory_id WorkspacesWorkspace#directory_id}.
	DirectoryId *string `json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#user_name WorkspacesWorkspace#user_name}.
	UserName *string `json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#root_volume_encryption_enabled WorkspacesWorkspace#root_volume_encryption_enabled}.
	RootVolumeEncryptionEnabled interface{} `json:"rootVolumeEncryptionEnabled" yaml:"rootVolumeEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#tags WorkspacesWorkspace#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#tags_all WorkspacesWorkspace#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#timeouts WorkspacesWorkspace#timeouts}
	Timeouts *WorkspacesWorkspaceTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#user_volume_encryption_enabled WorkspacesWorkspace#user_volume_encryption_enabled}.
	UserVolumeEncryptionEnabled interface{} `json:"userVolumeEncryptionEnabled" yaml:"userVolumeEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#volume_encryption_key WorkspacesWorkspace#volume_encryption_key}.
	VolumeEncryptionKey *string `json:"volumeEncryptionKey" yaml:"volumeEncryptionKey"`
	// workspace_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#workspace_properties WorkspacesWorkspace#workspace_properties}
	WorkspaceProperties *WorkspacesWorkspaceWorkspaceProperties `json:"workspaceProperties" yaml:"workspaceProperties"`
}

type WorkspacesWorkspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#create WorkspacesWorkspace#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#delete WorkspacesWorkspace#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#update WorkspacesWorkspace#update}.
	Update *string `json:"update" yaml:"update"`
}

type WorkspacesWorkspaceTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *WorkspacesWorkspaceTimeouts
	SetInternalValue(val *WorkspacesWorkspaceTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetUpdate()
}

// The jsii proxy struct for WorkspacesWorkspaceTimeoutsOutputReference
type jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) InternalValue() *WorkspacesWorkspaceTimeouts {
	var returns *WorkspacesWorkspaceTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewWorkspacesWorkspaceTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) WorkspacesWorkspaceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesWorkspaceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewWorkspacesWorkspaceTimeoutsOutputReference_Override(w WorkspacesWorkspaceTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesWorkspaceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		w,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) SetInternalValue(val *WorkspacesWorkspaceTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		w,
		"resetCreate",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		w,
		"resetDelete",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		w,
		"resetUpdate",
		nil, // no parameters
	)
}

type WorkspacesWorkspaceWorkspaceProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#compute_type_name WorkspacesWorkspace#compute_type_name}.
	ComputeTypeName *string `json:"computeTypeName" yaml:"computeTypeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#root_volume_size_gib WorkspacesWorkspace#root_volume_size_gib}.
	RootVolumeSizeGib *float64 `json:"rootVolumeSizeGib" yaml:"rootVolumeSizeGib"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#running_mode WorkspacesWorkspace#running_mode}.
	RunningMode *string `json:"runningMode" yaml:"runningMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#running_mode_auto_stop_timeout_in_minutes WorkspacesWorkspace#running_mode_auto_stop_timeout_in_minutes}.
	RunningModeAutoStopTimeoutInMinutes *float64 `json:"runningModeAutoStopTimeoutInMinutes" yaml:"runningModeAutoStopTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/workspaces_workspace#user_volume_size_gib WorkspacesWorkspace#user_volume_size_gib}.
	UserVolumeSizeGib *float64 `json:"userVolumeSizeGib" yaml:"userVolumeSizeGib"`
}

type WorkspacesWorkspaceWorkspacePropertiesOutputReference interface {
	cdktf.ComplexObject
	ComputeTypeName() *string
	SetComputeTypeName(val *string)
	ComputeTypeNameInput() *string
	InternalValue() *WorkspacesWorkspaceWorkspaceProperties
	SetInternalValue(val *WorkspacesWorkspaceWorkspaceProperties)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RootVolumeSizeGib() *float64
	SetRootVolumeSizeGib(val *float64)
	RootVolumeSizeGibInput() *float64
	RunningMode() *string
	SetRunningMode(val *string)
	RunningModeAutoStopTimeoutInMinutes() *float64
	SetRunningModeAutoStopTimeoutInMinutes(val *float64)
	RunningModeAutoStopTimeoutInMinutesInput() *float64
	RunningModeInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserVolumeSizeGib() *float64
	SetUserVolumeSizeGib(val *float64)
	UserVolumeSizeGibInput() *float64
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
	ResetComputeTypeName()
	ResetRootVolumeSizeGib()
	ResetRunningMode()
	ResetRunningModeAutoStopTimeoutInMinutes()
	ResetUserVolumeSizeGib()
}

// The jsii proxy struct for WorkspacesWorkspaceWorkspacePropertiesOutputReference
type jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ComputeTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ComputeTypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) InternalValue() *WorkspacesWorkspaceWorkspaceProperties {
	var returns *WorkspacesWorkspaceWorkspaceProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RootVolumeSizeGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RootVolumeSizeGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RunningMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runningMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RunningModeAutoStopTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningModeAutoStopTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RunningModeAutoStopTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningModeAutoStopTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) RunningModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runningModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) UserVolumeSizeGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userVolumeSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) UserVolumeSizeGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userVolumeSizeGibInput",
		&returns,
	)
	return returns
}

func NewWorkspacesWorkspaceWorkspacePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) WorkspacesWorkspaceWorkspacePropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesWorkspaceWorkspacePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewWorkspacesWorkspaceWorkspacePropertiesOutputReference_Override(w WorkspacesWorkspaceWorkspacePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.workspaces.WorkspacesWorkspaceWorkspacePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		w,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) SetComputeTypeName(val *string) {
	_jsii_.Set(
		j,
		"computeTypeName",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) SetInternalValue(val *WorkspacesWorkspaceWorkspaceProperties) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) SetRootVolumeSizeGib(val *float64) {
	_jsii_.Set(
		j,
		"rootVolumeSizeGib",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) SetRunningMode(val *string) {
	_jsii_.Set(
		j,
		"runningMode",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) SetRunningModeAutoStopTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"runningModeAutoStopTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) SetUserVolumeSizeGib(val *float64) {
	_jsii_.Set(
		j,
		"userVolumeSizeGib",
		val,
	)
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetComputeTypeName() {
	_jsii_.InvokeVoid(
		w,
		"resetComputeTypeName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetRootVolumeSizeGib() {
	_jsii_.InvokeVoid(
		w,
		"resetRootVolumeSizeGib",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetRunningMode() {
	_jsii_.InvokeVoid(
		w,
		"resetRunningMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetRunningModeAutoStopTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		w,
		"resetRunningModeAutoStopTimeoutInMinutes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesWorkspaceWorkspacePropertiesOutputReference) ResetUserVolumeSizeGib() {
	_jsii_.InvokeVoid(
		w,
		"resetUserVolumeSizeGib",
		nil, // no parameters
	)
}
