package batch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/batch/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html aws_batch_compute_environment}.
type BatchComputeEnvironment interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ComputeEnvironmentName() *string
	SetComputeEnvironmentName(val *string)
	ComputeEnvironmentNameInput() *string
	ComputeEnvironmentNamePrefix() *string
	SetComputeEnvironmentNamePrefix(val *string)
	ComputeEnvironmentNamePrefixInput() *string
	ComputeResources() BatchComputeEnvironmentComputeResourcesOutputReference
	ComputeResourcesInput() *BatchComputeEnvironmentComputeResources
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EcsClusterArn() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServiceRole() *string
	SetServiceRole(val *string)
	ServiceRoleInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	Status() *string
	StatusReason() *string
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
	PutComputeResources(value *BatchComputeEnvironmentComputeResources)
	ResetComputeEnvironmentName()
	ResetComputeEnvironmentNamePrefix()
	ResetComputeResources()
	ResetOverrideLogicalId()
	ResetServiceRole()
	ResetState()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for BatchComputeEnvironment
type jsiiProxy_BatchComputeEnvironment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BatchComputeEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) ComputeEnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) ComputeEnvironmentNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) ComputeEnvironmentNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) ComputeResources() BatchComputeEnvironmentComputeResourcesOutputReference {
	var returns BatchComputeEnvironmentComputeResourcesOutputReference
	_jsii_.Get(
		j,
		"computeResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) ComputeResourcesInput() *BatchComputeEnvironmentComputeResources {
	var returns *BatchComputeEnvironmentComputeResources
	_jsii_.Get(
		j,
		"computeResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) EcsClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html aws_batch_compute_environment} Resource.
func NewBatchComputeEnvironment(scope constructs.Construct, id *string, config *BatchComputeEnvironmentConfig) BatchComputeEnvironment {
	_init_.Initialize()

	j := jsiiProxy_BatchComputeEnvironment{}

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchComputeEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html aws_batch_compute_environment} Resource.
func NewBatchComputeEnvironment_Override(b BatchComputeEnvironment, scope constructs.Construct, id *string, config *BatchComputeEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchComputeEnvironment",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetComputeEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"computeEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetComputeEnvironmentNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"computeEnvironmentNamePrefix",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetType(val *string) {
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
func BatchComputeEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Batch.BatchComputeEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BatchComputeEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Batch.BatchComputeEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchComputeEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchComputeEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchComputeEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchComputeEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchComputeEnvironment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BatchComputeEnvironment) PutComputeResources(value *BatchComputeEnvironmentComputeResources) {
	_jsii_.InvokeVoid(
		b,
		"putComputeResources",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchComputeEnvironment) ResetComputeEnvironmentName() {
	_jsii_.InvokeVoid(
		b,
		"resetComputeEnvironmentName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironment) ResetComputeEnvironmentNamePrefix() {
	_jsii_.InvokeVoid(
		b,
		"resetComputeEnvironmentNamePrefix",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironment) ResetComputeResources() {
	_jsii_.InvokeVoid(
		b,
		"resetComputeResources",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (b *jsiiProxy_BatchComputeEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironment) ResetServiceRole() {
	_jsii_.InvokeVoid(
		b,
		"resetServiceRole",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironment) ResetState() {
	_jsii_.InvokeVoid(
		b,
		"resetState",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironment) ResetTagsAll() {
	_jsii_.InvokeVoid(
		b,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironment) SynthesizeAttributes() *map[string]interface{} {
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
func (b *jsiiProxy_BatchComputeEnvironment) ToMetadata() interface{} {
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
func (b *jsiiProxy_BatchComputeEnvironment) ToString() *string {
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
func (b *jsiiProxy_BatchComputeEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchComputeEnvironmentComputeResources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#max_vcpus BatchComputeEnvironment#max_vcpus}.
	MaxVcpus *float64 `json:"maxVcpus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#security_group_ids BatchComputeEnvironment#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#subnets BatchComputeEnvironment#subnets}.
	Subnets *[]*string `json:"subnets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#type BatchComputeEnvironment#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#allocation_strategy BatchComputeEnvironment#allocation_strategy}.
	AllocationStrategy *string `json:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#bid_percentage BatchComputeEnvironment#bid_percentage}.
	BidPercentage *float64 `json:"bidPercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#desired_vcpus BatchComputeEnvironment#desired_vcpus}.
	DesiredVcpus *float64 `json:"desiredVcpus"`
	// ec2_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#ec2_configuration BatchComputeEnvironment#ec2_configuration}
	Ec2Configuration *BatchComputeEnvironmentComputeResourcesEc2Configuration `json:"ec2Configuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#ec2_key_pair BatchComputeEnvironment#ec2_key_pair}.
	Ec2KeyPair *string `json:"ec2KeyPair"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#image_id BatchComputeEnvironment#image_id}.
	ImageId *string `json:"imageId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#instance_role BatchComputeEnvironment#instance_role}.
	InstanceRole *string `json:"instanceRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#instance_type BatchComputeEnvironment#instance_type}.
	InstanceType *[]*string `json:"instanceType"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#launch_template BatchComputeEnvironment#launch_template}
	LaunchTemplate *BatchComputeEnvironmentComputeResourcesLaunchTemplate `json:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#min_vcpus BatchComputeEnvironment#min_vcpus}.
	MinVcpus *float64 `json:"minVcpus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#spot_iam_fleet_role BatchComputeEnvironment#spot_iam_fleet_role}.
	SpotIamFleetRole *string `json:"spotIamFleetRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#tags BatchComputeEnvironment#tags}.
	Tags interface{} `json:"tags"`
}

type BatchComputeEnvironmentComputeResourcesEc2Configuration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#image_id_override BatchComputeEnvironment#image_id_override}.
	ImageIdOverride *string `json:"imageIdOverride"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#image_type BatchComputeEnvironment#image_type}.
	ImageType *string `json:"imageType"`
}

type BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference interface {
	cdktf.ComplexObject
	ImageIdOverride() *string
	SetImageIdOverride(val *string)
	ImageIdOverrideInput() *string
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
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
	ResetImageIdOverride()
	ResetImageType()
}

// The jsii proxy struct for BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference
type jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) ImageIdOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) ImageIdOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference_Override(b BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) SetImageIdOverride(val *string) {
	_jsii_.Set(
		j,
		"imageIdOverride",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) SetImageType(val *string) {
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) ResetImageIdOverride() {
	_jsii_.InvokeVoid(
		b,
		"resetImageIdOverride",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) ResetImageType() {
	_jsii_.InvokeVoid(
		b,
		"resetImageType",
		nil, // no parameters
	)
}

type BatchComputeEnvironmentComputeResourcesLaunchTemplate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#launch_template_id BatchComputeEnvironment#launch_template_id}.
	LaunchTemplateId *string `json:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#launch_template_name BatchComputeEnvironment#launch_template_name}.
	LaunchTemplateName *string `json:"launchTemplateName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#version BatchComputeEnvironment#version}.
	Version *string `json:"version"`
}

type BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LaunchTemplateId() *string
	SetLaunchTemplateId(val *string)
	LaunchTemplateIdInput() *string
	LaunchTemplateName() *string
	SetLaunchTemplateName(val *string)
	LaunchTemplateNameInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetLaunchTemplateId()
	ResetLaunchTemplateName()
	ResetVersion()
}

// The jsii proxy struct for BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference
type jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) LaunchTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) LaunchTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) LaunchTemplateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) LaunchTemplateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTemplateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func NewBatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference_Override(b BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) SetLaunchTemplateId(val *string) {
	_jsii_.Set(
		j,
		"launchTemplateId",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) SetLaunchTemplateName(val *string) {
	_jsii_.Set(
		j,
		"launchTemplateName",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ResetLaunchTemplateId() {
	_jsii_.InvokeVoid(
		b,
		"resetLaunchTemplateId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ResetLaunchTemplateName() {
	_jsii_.InvokeVoid(
		b,
		"resetLaunchTemplateName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetVersion",
		nil, // no parameters
	)
}

type BatchComputeEnvironmentComputeResourcesOutputReference interface {
	cdktf.ComplexObject
	AllocationStrategy() *string
	SetAllocationStrategy(val *string)
	AllocationStrategyInput() *string
	BidPercentage() *float64
	SetBidPercentage(val *float64)
	BidPercentageInput() *float64
	DesiredVcpus() *float64
	SetDesiredVcpus(val *float64)
	DesiredVcpusInput() *float64
	Ec2Configuration() BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference
	Ec2ConfigurationInput() *BatchComputeEnvironmentComputeResourcesEc2Configuration
	Ec2KeyPair() *string
	SetEc2KeyPair(val *string)
	Ec2KeyPairInput() *string
	ImageId() *string
	SetImageId(val *string)
	ImageIdInput() *string
	InstanceRole() *string
	SetInstanceRole(val *string)
	InstanceRoleInput() *string
	InstanceType() *[]*string
	SetInstanceType(val *[]*string)
	InstanceTypeInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LaunchTemplate() BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference
	LaunchTemplateInput() *BatchComputeEnvironmentComputeResourcesLaunchTemplate
	MaxVcpus() *float64
	SetMaxVcpus(val *float64)
	MaxVcpusInput() *float64
	MinVcpus() *float64
	SetMinVcpus(val *float64)
	MinVcpusInput() *float64
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SpotIamFleetRole() *string
	SetSpotIamFleetRole(val *string)
	SpotIamFleetRoleInput() *string
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	SubnetsInput() *[]*string
	Tags() interface{}
	SetTags(val interface{})
	TagsInput() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutEc2Configuration(value *BatchComputeEnvironmentComputeResourcesEc2Configuration)
	PutLaunchTemplate(value *BatchComputeEnvironmentComputeResourcesLaunchTemplate)
	ResetAllocationStrategy()
	ResetBidPercentage()
	ResetDesiredVcpus()
	ResetEc2Configuration()
	ResetEc2KeyPair()
	ResetImageId()
	ResetInstanceRole()
	ResetInstanceType()
	ResetLaunchTemplate()
	ResetMinVcpus()
	ResetSpotIamFleetRole()
	ResetTags()
}

// The jsii proxy struct for BatchComputeEnvironmentComputeResourcesOutputReference
type jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) BidPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) BidPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bidPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) DesiredVcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredVcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) DesiredVcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredVcpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Ec2Configuration() BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference {
	var returns BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference
	_jsii_.Get(
		j,
		"ec2Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Ec2ConfigurationInput() *BatchComputeEnvironmentComputeResourcesEc2Configuration {
	var returns *BatchComputeEnvironmentComputeResourcesEc2Configuration
	_jsii_.Get(
		j,
		"ec2ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Ec2KeyPair() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2KeyPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Ec2KeyPairInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2KeyPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InstanceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InstanceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InstanceType() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InstanceTypeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) LaunchTemplate() BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference {
	var returns BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) LaunchTemplateInput() *BatchComputeEnvironmentComputeResourcesLaunchTemplate {
	var returns *BatchComputeEnvironmentComputeResourcesLaunchTemplate
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) MaxVcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) MaxVcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxVcpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) MinVcpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) MinVcpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minVcpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SpotIamFleetRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotIamFleetRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SpotIamFleetRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotIamFleetRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewBatchComputeEnvironmentComputeResourcesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BatchComputeEnvironmentComputeResourcesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchComputeEnvironmentComputeResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchComputeEnvironmentComputeResourcesOutputReference_Override(b BatchComputeEnvironmentComputeResourcesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchComputeEnvironmentComputeResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetAllocationStrategy(val *string) {
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetBidPercentage(val *float64) {
	_jsii_.Set(
		j,
		"bidPercentage",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetDesiredVcpus(val *float64) {
	_jsii_.Set(
		j,
		"desiredVcpus",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetEc2KeyPair(val *string) {
	_jsii_.Set(
		j,
		"ec2KeyPair",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetImageId(val *string) {
	_jsii_.Set(
		j,
		"imageId",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetInstanceRole(val *string) {
	_jsii_.Set(
		j,
		"instanceRole",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetInstanceType(val *[]*string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetMaxVcpus(val *float64) {
	_jsii_.Set(
		j,
		"maxVcpus",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetMinVcpus(val *float64) {
	_jsii_.Set(
		j,
		"minVcpus",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetSpotIamFleetRole(val *string) {
	_jsii_.Set(
		j,
		"spotIamFleetRole",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) PutEc2Configuration(value *BatchComputeEnvironmentComputeResourcesEc2Configuration) {
	_jsii_.InvokeVoid(
		b,
		"putEc2Configuration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) PutLaunchTemplate(value *BatchComputeEnvironmentComputeResourcesLaunchTemplate) {
	_jsii_.InvokeVoid(
		b,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetBidPercentage() {
	_jsii_.InvokeVoid(
		b,
		"resetBidPercentage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetDesiredVcpus() {
	_jsii_.InvokeVoid(
		b,
		"resetDesiredVcpus",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetEc2Configuration() {
	_jsii_.InvokeVoid(
		b,
		"resetEc2Configuration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetEc2KeyPair() {
	_jsii_.InvokeVoid(
		b,
		"resetEc2KeyPair",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetImageId() {
	_jsii_.InvokeVoid(
		b,
		"resetImageId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetInstanceRole() {
	_jsii_.InvokeVoid(
		b,
		"resetInstanceRole",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		b,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		b,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetMinVcpus() {
	_jsii_.InvokeVoid(
		b,
		"resetMinVcpus",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetSpotIamFleetRole() {
	_jsii_.InvokeVoid(
		b,
		"resetSpotIamFleetRole",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

type BatchComputeEnvironmentConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#type BatchComputeEnvironment#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#compute_environment_name BatchComputeEnvironment#compute_environment_name}.
	ComputeEnvironmentName *string `json:"computeEnvironmentName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#compute_environment_name_prefix BatchComputeEnvironment#compute_environment_name_prefix}.
	ComputeEnvironmentNamePrefix *string `json:"computeEnvironmentNamePrefix"`
	// compute_resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#compute_resources BatchComputeEnvironment#compute_resources}
	ComputeResources *BatchComputeEnvironmentComputeResources `json:"computeResources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#service_role BatchComputeEnvironment#service_role}.
	ServiceRole *string `json:"serviceRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#state BatchComputeEnvironment#state}.
	State *string `json:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#tags BatchComputeEnvironment#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment.html#tags_all BatchComputeEnvironment#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html aws_batch_job_definition}.
type BatchJobDefinition interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ContainerProperties() *string
	SetContainerProperties(val *string)
	ContainerPropertiesInput() *string
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
	Parameters() interface{}
	SetParameters(val interface{})
	ParametersInput() interface{}
	PlatformCapabilities() *[]*string
	SetPlatformCapabilities(val *[]*string)
	PlatformCapabilitiesInput() *[]*string
	PropagateTags() interface{}
	SetPropagateTags(val interface{})
	PropagateTagsInput() interface{}
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RetryStrategy() BatchJobDefinitionRetryStrategyOutputReference
	RetryStrategyInput() *BatchJobDefinitionRetryStrategy
	Revision() *float64
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeout() BatchJobDefinitionTimeoutOutputReference
	TimeoutInput() *BatchJobDefinitionTimeout
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
	PutRetryStrategy(value *BatchJobDefinitionRetryStrategy)
	PutTimeout(value *BatchJobDefinitionTimeout)
	ResetContainerProperties()
	ResetOverrideLogicalId()
	ResetParameters()
	ResetPlatformCapabilities()
	ResetPropagateTags()
	ResetRetryStrategy()
	ResetTags()
	ResetTagsAll()
	ResetTimeout()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for BatchJobDefinition
type jsiiProxy_BatchJobDefinition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BatchJobDefinition) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ContainerProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ContainerPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) PlatformCapabilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"platformCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) PlatformCapabilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"platformCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) PropagateTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) PropagateTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) RetryStrategy() BatchJobDefinitionRetryStrategyOutputReference {
	var returns BatchJobDefinitionRetryStrategyOutputReference
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) RetryStrategyInput() *BatchJobDefinitionRetryStrategy {
	var returns *BatchJobDefinitionRetryStrategy
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Revision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Timeout() BatchJobDefinitionTimeoutOutputReference {
	var returns BatchJobDefinitionTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TimeoutInput() *BatchJobDefinitionTimeout {
	var returns *BatchJobDefinitionTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html aws_batch_job_definition} Resource.
func NewBatchJobDefinition(scope constructs.Construct, id *string, config *BatchJobDefinitionConfig) BatchJobDefinition {
	_init_.Initialize()

	j := jsiiProxy_BatchJobDefinition{}

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchJobDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html aws_batch_job_definition} Resource.
func NewBatchJobDefinition_Override(b BatchJobDefinition, scope constructs.Construct, id *string, config *BatchJobDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchJobDefinition",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetContainerProperties(val *string) {
	_jsii_.Set(
		j,
		"containerProperties",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetPlatformCapabilities(val *[]*string) {
	_jsii_.Set(
		j,
		"platformCapabilities",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetPropagateTags(val interface{}) {
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetType(val *string) {
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
func BatchJobDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Batch.BatchJobDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BatchJobDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Batch.BatchJobDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchJobDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchJobDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchJobDefinition) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchJobDefinition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchJobDefinition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BatchJobDefinition) PutRetryStrategy(value *BatchJobDefinitionRetryStrategy) {
	_jsii_.InvokeVoid(
		b,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinition) PutTimeout(value *BatchJobDefinitionTimeout) {
	_jsii_.InvokeVoid(
		b,
		"putTimeout",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetContainerProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetContainerProperties",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (b *jsiiProxy_BatchJobDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetPlatformCapabilities() {
	_jsii_.InvokeVoid(
		b,
		"resetPlatformCapabilities",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		b,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		b,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetTagsAll() {
	_jsii_.InvokeVoid(
		b,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) ResetTimeout() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeout",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinition) SynthesizeAttributes() *map[string]interface{} {
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
func (b *jsiiProxy_BatchJobDefinition) ToMetadata() interface{} {
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
func (b *jsiiProxy_BatchJobDefinition) ToString() *string {
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
func (b *jsiiProxy_BatchJobDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchJobDefinitionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#name BatchJobDefinition#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#type BatchJobDefinition#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#container_properties BatchJobDefinition#container_properties}.
	ContainerProperties *string `json:"containerProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#parameters BatchJobDefinition#parameters}.
	Parameters interface{} `json:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#platform_capabilities BatchJobDefinition#platform_capabilities}.
	PlatformCapabilities *[]*string `json:"platformCapabilities"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#propagate_tags BatchJobDefinition#propagate_tags}.
	PropagateTags interface{} `json:"propagateTags"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#retry_strategy BatchJobDefinition#retry_strategy}
	RetryStrategy *BatchJobDefinitionRetryStrategy `json:"retryStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#tags BatchJobDefinition#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#tags_all BatchJobDefinition#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#timeout BatchJobDefinition#timeout}
	Timeout *BatchJobDefinitionTimeout `json:"timeout"`
}

type BatchJobDefinitionRetryStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#attempts BatchJobDefinition#attempts}.
	Attempts *float64 `json:"attempts"`
	// evaluate_on_exit block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#evaluate_on_exit BatchJobDefinition#evaluate_on_exit}
	EvaluateOnExit *[]*BatchJobDefinitionRetryStrategyEvaluateOnExit `json:"evaluateOnExit"`
}

type BatchJobDefinitionRetryStrategyEvaluateOnExit struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#action BatchJobDefinition#action}.
	Action *string `json:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#on_exit_code BatchJobDefinition#on_exit_code}.
	OnExitCode *string `json:"onExitCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#on_reason BatchJobDefinition#on_reason}.
	OnReason *string `json:"onReason"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#on_status_reason BatchJobDefinition#on_status_reason}.
	OnStatusReason *string `json:"onStatusReason"`
}

type BatchJobDefinitionRetryStrategyOutputReference interface {
	cdktf.ComplexObject
	Attempts() *float64
	SetAttempts(val *float64)
	AttemptsInput() *float64
	EvaluateOnExit() *[]*BatchJobDefinitionRetryStrategyEvaluateOnExit
	SetEvaluateOnExit(val *[]*BatchJobDefinitionRetryStrategyEvaluateOnExit)
	EvaluateOnExitInput() *[]*BatchJobDefinitionRetryStrategyEvaluateOnExit
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
	ResetAttempts()
	ResetEvaluateOnExit()
}

// The jsii proxy struct for BatchJobDefinitionRetryStrategyOutputReference
type jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) Attempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) AttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) EvaluateOnExit() *[]*BatchJobDefinitionRetryStrategyEvaluateOnExit {
	var returns *[]*BatchJobDefinitionRetryStrategyEvaluateOnExit
	_jsii_.Get(
		j,
		"evaluateOnExit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) EvaluateOnExitInput() *[]*BatchJobDefinitionRetryStrategyEvaluateOnExit {
	var returns *[]*BatchJobDefinitionRetryStrategyEvaluateOnExit
	_jsii_.Get(
		j,
		"evaluateOnExitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBatchJobDefinitionRetryStrategyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BatchJobDefinitionRetryStrategyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchJobDefinitionRetryStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionRetryStrategyOutputReference_Override(b BatchJobDefinitionRetryStrategyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchJobDefinitionRetryStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) SetAttempts(val *float64) {
	_jsii_.Set(
		j,
		"attempts",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) SetEvaluateOnExit(val *[]*BatchJobDefinitionRetryStrategyEvaluateOnExit) {
	_jsii_.Set(
		j,
		"evaluateOnExit",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) ResetAttempts() {
	_jsii_.InvokeVoid(
		b,
		"resetAttempts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) ResetEvaluateOnExit() {
	_jsii_.InvokeVoid(
		b,
		"resetEvaluateOnExit",
		nil, // no parameters
	)
}

type BatchJobDefinitionTimeout struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition.html#attempt_duration_seconds BatchJobDefinition#attempt_duration_seconds}.
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds"`
}

type BatchJobDefinitionTimeoutOutputReference interface {
	cdktf.ComplexObject
	AttemptDurationSeconds() *float64
	SetAttemptDurationSeconds(val *float64)
	AttemptDurationSecondsInput() *float64
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
	ResetAttemptDurationSeconds()
}

// The jsii proxy struct for BatchJobDefinitionTimeoutOutputReference
type jsiiProxy_BatchJobDefinitionTimeoutOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) AttemptDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attemptDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) AttemptDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attemptDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBatchJobDefinitionTimeoutOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) BatchJobDefinitionTimeoutOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchJobDefinitionTimeoutOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchJobDefinitionTimeoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionTimeoutOutputReference_Override(b BatchJobDefinitionTimeoutOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchJobDefinitionTimeoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) SetAttemptDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"attemptDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) ResetAttemptDurationSeconds() {
	_jsii_.InvokeVoid(
		b,
		"resetAttemptDurationSeconds",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue.html aws_batch_job_queue}.
type BatchJobQueue interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ComputeEnvironments() *[]*string
	SetComputeEnvironments(val *[]*string)
	ComputeEnvironmentsInput() *[]*string
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
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	State() *string
	SetState(val *string)
	StateInput() *string
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
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for BatchJobQueue
type jsiiProxy_BatchJobQueue struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BatchJobQueue) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) ComputeEnvironments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"computeEnvironments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) ComputeEnvironmentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"computeEnvironmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue.html aws_batch_job_queue} Resource.
func NewBatchJobQueue(scope constructs.Construct, id *string, config *BatchJobQueueConfig) BatchJobQueue {
	_init_.Initialize()

	j := jsiiProxy_BatchJobQueue{}

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchJobQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue.html aws_batch_job_queue} Resource.
func NewBatchJobQueue_Override(b BatchJobQueue, scope constructs.Construct, id *string, config *BatchJobQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.BatchJobQueue",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetComputeEnvironments(val *[]*string) {
	_jsii_.Set(
		j,
		"computeEnvironments",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetTagsAll(val interface{}) {
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
func BatchJobQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Batch.BatchJobQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BatchJobQueue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Batch.BatchJobQueue",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (b *jsiiProxy_BatchJobQueue) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (b *jsiiProxy_BatchJobQueue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchJobQueue) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchJobQueue) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchJobQueue) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchJobQueue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchJobQueue) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (b *jsiiProxy_BatchJobQueue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobQueue) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobQueue) ResetTagsAll() {
	_jsii_.InvokeVoid(
		b,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobQueue) SynthesizeAttributes() *map[string]interface{} {
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
func (b *jsiiProxy_BatchJobQueue) ToMetadata() interface{} {
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
func (b *jsiiProxy_BatchJobQueue) ToString() *string {
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
func (b *jsiiProxy_BatchJobQueue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type BatchJobQueueConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue.html#compute_environments BatchJobQueue#compute_environments}.
	ComputeEnvironments *[]*string `json:"computeEnvironments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue.html#name BatchJobQueue#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue.html#priority BatchJobQueue#priority}.
	Priority *float64 `json:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue.html#state BatchJobQueue#state}.
	State *string `json:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue.html#tags BatchJobQueue#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue.html#tags_all BatchJobQueue#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment.html aws_batch_compute_environment}.
type DataAwsBatchComputeEnvironment interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ComputeEnvironmentName() *string
	SetComputeEnvironmentName(val *string)
	ComputeEnvironmentNameInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EcsClusterArn() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServiceRole() *string
	State() *string
	Status() *string
	StatusReason() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsBatchComputeEnvironment
type jsiiProxy_DataAwsBatchComputeEnvironment struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) ComputeEnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) ComputeEnvironmentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironmentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) EcsClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment.html aws_batch_compute_environment} Data Source.
func NewDataAwsBatchComputeEnvironment(scope constructs.Construct, id *string, config *DataAwsBatchComputeEnvironmentConfig) DataAwsBatchComputeEnvironment {
	_init_.Initialize()

	j := jsiiProxy_DataAwsBatchComputeEnvironment{}

	_jsii_.Create(
		"hashicorp_aws.Batch.DataAwsBatchComputeEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment.html aws_batch_compute_environment} Data Source.
func NewDataAwsBatchComputeEnvironment_Override(d DataAwsBatchComputeEnvironment, scope constructs.Construct, id *string, config *DataAwsBatchComputeEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.DataAwsBatchComputeEnvironment",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) SetComputeEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"computeEnvironmentName",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) SetTags(val interface{}) {
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
func DataAwsBatchComputeEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Batch.DataAwsBatchComputeEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsBatchComputeEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Batch.DataAwsBatchComputeEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBatchComputeEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBatchComputeEnvironment) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) ToString() *string {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsBatchComputeEnvironmentConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment.html#compute_environment_name DataAwsBatchComputeEnvironment#compute_environment_name}.
	ComputeEnvironmentName *string `json:"computeEnvironmentName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment.html#tags DataAwsBatchComputeEnvironment#tags}.
	Tags interface{} `json:"tags"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue.html aws_batch_job_queue}.
type DataAwsBatchJobQueue interface {
	cdktf.TerraformDataSource
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
	Priority() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	State() *string
	Status() *string
	StatusReason() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	ComputeEnvironmentOrder(index *string) DataAwsBatchJobQueueComputeEnvironmentOrder
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsBatchJobQueue
type jsiiProxy_DataAwsBatchJobQueue struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue.html aws_batch_job_queue} Data Source.
func NewDataAwsBatchJobQueue(scope constructs.Construct, id *string, config *DataAwsBatchJobQueueConfig) DataAwsBatchJobQueue {
	_init_.Initialize()

	j := jsiiProxy_DataAwsBatchJobQueue{}

	_jsii_.Create(
		"hashicorp_aws.Batch.DataAwsBatchJobQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue.html aws_batch_job_queue} Data Source.
func NewDataAwsBatchJobQueue_Override(d DataAwsBatchJobQueue, scope constructs.Construct, id *string, config *DataAwsBatchJobQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.DataAwsBatchJobQueue",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueue) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueue) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueue) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueue) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueue) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueue) SetTags(val interface{}) {
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
func DataAwsBatchJobQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Batch.DataAwsBatchJobQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsBatchJobQueue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Batch.DataAwsBatchJobQueue",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchJobQueue) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsBatchJobQueue) ComputeEnvironmentOrder(index *string) DataAwsBatchJobQueueComputeEnvironmentOrder {
	var returns DataAwsBatchJobQueueComputeEnvironmentOrder

	_jsii_.Invoke(
		d,
		"computeEnvironmentOrder",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchJobQueue) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsBatchJobQueue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBatchJobQueue) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBatchJobQueue) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) ToString() *string {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsBatchJobQueueComputeEnvironmentOrder interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ComputeEnvironment() *string
	Order() *float64
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

// The jsii proxy struct for DataAwsBatchJobQueueComputeEnvironmentOrder
type jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) ComputeEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsBatchJobQueueComputeEnvironmentOrder(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsBatchJobQueueComputeEnvironmentOrder {
	_init_.Initialize()

	j := jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder{}

	_jsii_.Create(
		"hashicorp_aws.Batch.DataAwsBatchJobQueueComputeEnvironmentOrder",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsBatchJobQueueComputeEnvironmentOrder_Override(d DataAwsBatchJobQueueComputeEnvironmentOrder, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Batch.DataAwsBatchJobQueueComputeEnvironmentOrder",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsBatchJobQueueConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue.html#name DataAwsBatchJobQueue#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue.html#tags DataAwsBatchJobQueue#tags}.
	Tags interface{} `json:"tags"`
}
