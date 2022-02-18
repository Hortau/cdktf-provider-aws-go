package batch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/batch/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment aws_batch_compute_environment}.
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
	Count() *float64
	SetCount(val *float64)
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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

func (j *jsiiProxy_BatchComputeEnvironment) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_BatchComputeEnvironment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironment) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment aws_batch_compute_environment} Resource.
func NewBatchComputeEnvironment(scope constructs.Construct, id *string, config *BatchComputeEnvironmentConfig) BatchComputeEnvironment {
	_init_.Initialize()

	j := jsiiProxy_BatchComputeEnvironment{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchComputeEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment aws_batch_compute_environment} Resource.
func NewBatchComputeEnvironment_Override(b BatchComputeEnvironment, scope constructs.Construct, id *string, config *BatchComputeEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchComputeEnvironment",
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

func (j *jsiiProxy_BatchComputeEnvironment) SetCount(val *float64) {
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

func (j *jsiiProxy_BatchComputeEnvironment) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchComputeEnvironment) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.batch.BatchComputeEnvironment",
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
		"hashicorp_aws.batch.BatchComputeEnvironment",
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
func (b *jsiiProxy_BatchComputeEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BatchComputeEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
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
func (b *jsiiProxy_BatchComputeEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
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
func (b *jsiiProxy_BatchComputeEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#max_vcpus BatchComputeEnvironment#max_vcpus}.
	MaxVcpus *float64 `json:"maxVcpus" yaml:"maxVcpus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#security_group_ids BatchComputeEnvironment#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#subnets BatchComputeEnvironment#subnets}.
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#type BatchComputeEnvironment#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#allocation_strategy BatchComputeEnvironment#allocation_strategy}.
	AllocationStrategy *string `json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#bid_percentage BatchComputeEnvironment#bid_percentage}.
	BidPercentage *float64 `json:"bidPercentage" yaml:"bidPercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#desired_vcpus BatchComputeEnvironment#desired_vcpus}.
	DesiredVcpus *float64 `json:"desiredVcpus" yaml:"desiredVcpus"`
	// ec2_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#ec2_configuration BatchComputeEnvironment#ec2_configuration}
	Ec2Configuration *BatchComputeEnvironmentComputeResourcesEc2Configuration `json:"ec2Configuration" yaml:"ec2Configuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#ec2_key_pair BatchComputeEnvironment#ec2_key_pair}.
	Ec2KeyPair *string `json:"ec2KeyPair" yaml:"ec2KeyPair"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#image_id BatchComputeEnvironment#image_id}.
	ImageId *string `json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#instance_role BatchComputeEnvironment#instance_role}.
	InstanceRole *string `json:"instanceRole" yaml:"instanceRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#instance_type BatchComputeEnvironment#instance_type}.
	InstanceType *[]*string `json:"instanceType" yaml:"instanceType"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#launch_template BatchComputeEnvironment#launch_template}
	LaunchTemplate *BatchComputeEnvironmentComputeResourcesLaunchTemplate `json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#min_vcpus BatchComputeEnvironment#min_vcpus}.
	MinVcpus *float64 `json:"minVcpus" yaml:"minVcpus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#spot_iam_fleet_role BatchComputeEnvironment#spot_iam_fleet_role}.
	SpotIamFleetRole *string `json:"spotIamFleetRole" yaml:"spotIamFleetRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#tags BatchComputeEnvironment#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type BatchComputeEnvironmentComputeResourcesEc2Configuration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#image_id_override BatchComputeEnvironment#image_id_override}.
	ImageIdOverride *string `json:"imageIdOverride" yaml:"imageIdOverride"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#image_type BatchComputeEnvironment#image_type}.
	ImageType *string `json:"imageType" yaml:"imageType"`
}

type BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference interface {
	cdktf.ComplexObject
	ImageIdOverride() *string
	SetImageIdOverride(val *string)
	ImageIdOverrideInput() *string
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	InternalValue() *BatchComputeEnvironmentComputeResourcesEc2Configuration
	SetInternalValue(val *BatchComputeEnvironmentComputeResourcesEc2Configuration)
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) InternalValue() *BatchComputeEnvironmentComputeResourcesEc2Configuration {
	var returns *BatchComputeEnvironmentComputeResourcesEc2Configuration
	_jsii_.Get(
		j,
		"internalValue",
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference_Override(b BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference",
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) SetInternalValue(val *BatchComputeEnvironmentComputeResourcesEc2Configuration) {
	_jsii_.Set(
		j,
		"internalValue",
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesEc2ConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#launch_template_id BatchComputeEnvironment#launch_template_id}.
	LaunchTemplateId *string `json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#launch_template_name BatchComputeEnvironment#launch_template_name}.
	LaunchTemplateName *string `json:"launchTemplateName" yaml:"launchTemplateName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#version BatchComputeEnvironment#version}.
	Version *string `json:"version" yaml:"version"`
}

type BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *BatchComputeEnvironmentComputeResourcesLaunchTemplate
	SetInternalValue(val *BatchComputeEnvironmentComputeResourcesLaunchTemplate)
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
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	ResetLaunchTemplateId()
	ResetLaunchTemplateName()
	ResetVersion()
}

// The jsii proxy struct for BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference
type jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) InternalValue() *BatchComputeEnvironmentComputeResourcesLaunchTemplate {
	var returns *BatchComputeEnvironmentComputeResourcesLaunchTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
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

func NewBatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference_Override(b BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) SetInternalValue(val *BatchComputeEnvironmentComputeResourcesLaunchTemplate) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesLaunchTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
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
	InternalValue() *BatchComputeEnvironmentComputeResources
	SetInternalValue(val *BatchComputeEnvironmentComputeResources)
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) InternalValue() *BatchComputeEnvironmentComputeResources {
	var returns *BatchComputeEnvironmentComputeResources
	_jsii_.Get(
		j,
		"internalValue",
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
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

func NewBatchComputeEnvironmentComputeResourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) BatchComputeEnvironmentComputeResourcesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchComputeEnvironmentComputeResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchComputeEnvironmentComputeResourcesOutputReference_Override(b BatchComputeEnvironmentComputeResourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchComputeEnvironmentComputeResourcesOutputReference",
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetInternalValue(val *BatchComputeEnvironmentComputeResources) {
	_jsii_.Set(
		j,
		"internalValue",
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetTags(val *map[string]*string) {
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

func (j *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
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
func (b *jsiiProxy_BatchComputeEnvironmentComputeResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
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

// AWS Batch.
type BatchComputeEnvironmentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#type BatchComputeEnvironment#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#compute_environment_name BatchComputeEnvironment#compute_environment_name}.
	ComputeEnvironmentName *string `json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#compute_environment_name_prefix BatchComputeEnvironment#compute_environment_name_prefix}.
	ComputeEnvironmentNamePrefix *string `json:"computeEnvironmentNamePrefix" yaml:"computeEnvironmentNamePrefix"`
	// compute_resources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#compute_resources BatchComputeEnvironment#compute_resources}
	ComputeResources *BatchComputeEnvironmentComputeResources `json:"computeResources" yaml:"computeResources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#service_role BatchComputeEnvironment#service_role}.
	ServiceRole *string `json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#state BatchComputeEnvironment#state}.
	State *string `json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#tags BatchComputeEnvironment#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#tags_all BatchComputeEnvironment#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition aws_batch_job_definition}.
type BatchJobDefinition interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ContainerProperties() *string
	SetContainerProperties(val *string)
	ContainerPropertiesInput() *string
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
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeout() BatchJobDefinitionTimeoutOutputReference
	TimeoutInput() *BatchJobDefinitionTimeout
	Type() *string
	SetType(val *string)
	TypeInput() *string
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

func (j *jsiiProxy_BatchJobDefinition) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_BatchJobDefinition) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) ParametersInput() *map[string]*string {
	var returns *map[string]*string
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

func (j *jsiiProxy_BatchJobDefinition) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinition) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition aws_batch_job_definition} Resource.
func NewBatchJobDefinition(scope constructs.Construct, id *string, config *BatchJobDefinitionConfig) BatchJobDefinition {
	_init_.Initialize()

	j := jsiiProxy_BatchJobDefinition{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchJobDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition aws_batch_job_definition} Resource.
func NewBatchJobDefinition_Override(b BatchJobDefinition, scope constructs.Construct, id *string, config *BatchJobDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchJobDefinition",
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

func (j *jsiiProxy_BatchJobDefinition) SetCount(val *float64) {
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

func (j *jsiiProxy_BatchJobDefinition) SetParameters(val *map[string]*string) {
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

func (j *jsiiProxy_BatchJobDefinition) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinition) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.batch.BatchJobDefinition",
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
		"hashicorp_aws.batch.BatchJobDefinition",
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
func (b *jsiiProxy_BatchJobDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BatchJobDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
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
func (b *jsiiProxy_BatchJobDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
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
func (b *jsiiProxy_BatchJobDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
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

// AWS Batch.
type BatchJobDefinitionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#type BatchJobDefinition#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#container_properties BatchJobDefinition#container_properties}.
	ContainerProperties *string `json:"containerProperties" yaml:"containerProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#parameters BatchJobDefinition#parameters}.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#platform_capabilities BatchJobDefinition#platform_capabilities}.
	PlatformCapabilities *[]*string `json:"platformCapabilities" yaml:"platformCapabilities"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#propagate_tags BatchJobDefinition#propagate_tags}.
	PropagateTags interface{} `json:"propagateTags" yaml:"propagateTags"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#retry_strategy BatchJobDefinition#retry_strategy}
	RetryStrategy *BatchJobDefinitionRetryStrategy `json:"retryStrategy" yaml:"retryStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#tags BatchJobDefinition#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#tags_all BatchJobDefinition#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#timeout BatchJobDefinition#timeout}
	Timeout *BatchJobDefinitionTimeout `json:"timeout" yaml:"timeout"`
}

type BatchJobDefinitionRetryStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#attempts BatchJobDefinition#attempts}.
	Attempts *float64 `json:"attempts" yaml:"attempts"`
	// evaluate_on_exit block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#evaluate_on_exit BatchJobDefinition#evaluate_on_exit}
	EvaluateOnExit interface{} `json:"evaluateOnExit" yaml:"evaluateOnExit"`
}

type BatchJobDefinitionRetryStrategyEvaluateOnExit struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#action BatchJobDefinition#action}.
	Action *string `json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#on_exit_code BatchJobDefinition#on_exit_code}.
	OnExitCode *string `json:"onExitCode" yaml:"onExitCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#on_reason BatchJobDefinition#on_reason}.
	OnReason *string `json:"onReason" yaml:"onReason"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#on_status_reason BatchJobDefinition#on_status_reason}.
	OnStatusReason *string `json:"onStatusReason" yaml:"onStatusReason"`
}

type BatchJobDefinitionRetryStrategyOutputReference interface {
	cdktf.ComplexObject
	Attempts() *float64
	SetAttempts(val *float64)
	AttemptsInput() *float64
	EvaluateOnExit() interface{}
	SetEvaluateOnExit(val interface{})
	EvaluateOnExitInput() interface{}
	InternalValue() *BatchJobDefinitionRetryStrategy
	SetInternalValue(val *BatchJobDefinitionRetryStrategy)
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

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) EvaluateOnExit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateOnExit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) EvaluateOnExitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluateOnExitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) InternalValue() *BatchJobDefinitionRetryStrategy {
	var returns *BatchJobDefinitionRetryStrategy
	_jsii_.Get(
		j,
		"internalValue",
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

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBatchJobDefinitionRetryStrategyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) BatchJobDefinitionRetryStrategyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchJobDefinitionRetryStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionRetryStrategyOutputReference_Override(b BatchJobDefinitionRetryStrategyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchJobDefinitionRetryStrategyOutputReference",
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

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) SetEvaluateOnExit(val interface{}) {
	_jsii_.Set(
		j,
		"evaluateOnExit",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) SetInternalValue(val *BatchJobDefinitionRetryStrategy) {
	_jsii_.Set(
		j,
		"internalValue",
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

func (j *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
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
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
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
func (b *jsiiProxy_BatchJobDefinitionRetryStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_definition#attempt_duration_seconds BatchJobDefinition#attempt_duration_seconds}.
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds" yaml:"attemptDurationSeconds"`
}

type BatchJobDefinitionTimeoutOutputReference interface {
	cdktf.ComplexObject
	AttemptDurationSeconds() *float64
	SetAttemptDurationSeconds(val *float64)
	AttemptDurationSecondsInput() *float64
	InternalValue() *BatchJobDefinitionTimeout
	SetInternalValue(val *BatchJobDefinitionTimeout)
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

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) InternalValue() *BatchJobDefinitionTimeout {
	var returns *BatchJobDefinitionTimeout
	_jsii_.Get(
		j,
		"internalValue",
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

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBatchJobDefinitionTimeoutOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) BatchJobDefinitionTimeoutOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchJobDefinitionTimeoutOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchJobDefinitionTimeoutOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionTimeoutOutputReference_Override(b BatchJobDefinitionTimeoutOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchJobDefinitionTimeoutOutputReference",
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

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) SetInternalValue(val *BatchJobDefinitionTimeout) {
	_jsii_.Set(
		j,
		"internalValue",
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

func (j *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
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
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
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
func (b *jsiiProxy_BatchJobDefinitionTimeoutOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
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

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue aws_batch_job_queue}.
type BatchJobQueue interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ComputeEnvironments() *[]*string
	SetComputeEnvironments(val *[]*string)
	ComputeEnvironmentsInput() *[]*string
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
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SchedulingPolicyArn() *string
	SetSchedulingPolicyArn(val *string)
	SchedulingPolicyArnInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
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
	ResetOverrideLogicalId()
	ResetSchedulingPolicyArn()
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

func (j *jsiiProxy_BatchJobQueue) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_BatchJobQueue) SchedulingPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) SchedulingPolicyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyArnInput",
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

func (j *jsiiProxy_BatchJobQueue) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobQueue) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue aws_batch_job_queue} Resource.
func NewBatchJobQueue(scope constructs.Construct, id *string, config *BatchJobQueueConfig) BatchJobQueue {
	_init_.Initialize()

	j := jsiiProxy_BatchJobQueue{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchJobQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue aws_batch_job_queue} Resource.
func NewBatchJobQueue_Override(b BatchJobQueue, scope constructs.Construct, id *string, config *BatchJobQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchJobQueue",
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

func (j *jsiiProxy_BatchJobQueue) SetCount(val *float64) {
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

func (j *jsiiProxy_BatchJobQueue) SetSchedulingPolicyArn(val *string) {
	_jsii_.Set(
		j,
		"schedulingPolicyArn",
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

func (j *jsiiProxy_BatchJobQueue) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchJobQueue) SetTagsAll(val *map[string]*string) {
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
		"hashicorp_aws.batch.BatchJobQueue",
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
		"hashicorp_aws.batch.BatchJobQueue",
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
func (b *jsiiProxy_BatchJobQueue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BatchJobQueue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
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
func (b *jsiiProxy_BatchJobQueue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchJobQueue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
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
func (b *jsiiProxy_BatchJobQueue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
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

func (b *jsiiProxy_BatchJobQueue) ResetSchedulingPolicyArn() {
	_jsii_.InvokeVoid(
		b,
		"resetSchedulingPolicyArn",
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

// AWS Batch.
type BatchJobQueueConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue#compute_environments BatchJobQueue#compute_environments}.
	ComputeEnvironments *[]*string `json:"computeEnvironments" yaml:"computeEnvironments"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue#name BatchJobQueue#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue#priority BatchJobQueue#priority}.
	Priority *float64 `json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue#state BatchJobQueue#state}.
	State *string `json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue#scheduling_policy_arn BatchJobQueue#scheduling_policy_arn}.
	SchedulingPolicyArn *string `json:"schedulingPolicyArn" yaml:"schedulingPolicyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue#tags BatchJobQueue#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_job_queue#tags_all BatchJobQueue#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy aws_batch_scheduling_policy}.
type BatchSchedulingPolicy interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FairSharePolicy() BatchSchedulingPolicyFairSharePolicyOutputReference
	FairSharePolicyInput() *BatchSchedulingPolicyFairSharePolicy
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
	PutFairSharePolicy(value *BatchSchedulingPolicyFairSharePolicy)
	ResetFairSharePolicy()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for BatchSchedulingPolicy
type jsiiProxy_BatchSchedulingPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BatchSchedulingPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) FairSharePolicy() BatchSchedulingPolicyFairSharePolicyOutputReference {
	var returns BatchSchedulingPolicyFairSharePolicyOutputReference
	_jsii_.Get(
		j,
		"fairSharePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) FairSharePolicyInput() *BatchSchedulingPolicyFairSharePolicy {
	var returns *BatchSchedulingPolicyFairSharePolicy
	_jsii_.Get(
		j,
		"fairSharePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy aws_batch_scheduling_policy} Resource.
func NewBatchSchedulingPolicy(scope constructs.Construct, id *string, config *BatchSchedulingPolicyConfig) BatchSchedulingPolicy {
	_init_.Initialize()

	j := jsiiProxy_BatchSchedulingPolicy{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchSchedulingPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy aws_batch_scheduling_policy} Resource.
func NewBatchSchedulingPolicy_Override(b BatchSchedulingPolicy, scope constructs.Construct, id *string, config *BatchSchedulingPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchSchedulingPolicy",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicy) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicy) SetTagsAll(val *map[string]*string) {
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
func BatchSchedulingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.batch.BatchSchedulingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BatchSchedulingPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.batch.BatchSchedulingPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchSchedulingPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchSchedulingPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchSchedulingPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchSchedulingPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchSchedulingPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BatchSchedulingPolicy) PutFairSharePolicy(value *BatchSchedulingPolicyFairSharePolicy) {
	_jsii_.InvokeVoid(
		b,
		"putFairSharePolicy",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchSchedulingPolicy) ResetFairSharePolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetFairSharePolicy",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchSchedulingPolicy) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchSchedulingPolicy) ResetTagsAll() {
	_jsii_.InvokeVoid(
		b,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchSchedulingPolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (b *jsiiProxy_BatchSchedulingPolicy) ToMetadata() interface{} {
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
func (b *jsiiProxy_BatchSchedulingPolicy) ToString() *string {
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
func (b *jsiiProxy_BatchSchedulingPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Batch.
type BatchSchedulingPolicyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#name BatchSchedulingPolicy#name}.
	Name *string `json:"name" yaml:"name"`
	// fair_share_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#fair_share_policy BatchSchedulingPolicy#fair_share_policy}
	FairSharePolicy *BatchSchedulingPolicyFairSharePolicy `json:"fairSharePolicy" yaml:"fairSharePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#tags BatchSchedulingPolicy#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#tags_all BatchSchedulingPolicy#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

type BatchSchedulingPolicyFairSharePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#compute_reservation BatchSchedulingPolicy#compute_reservation}.
	ComputeReservation *float64 `json:"computeReservation" yaml:"computeReservation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#share_decay_seconds BatchSchedulingPolicy#share_decay_seconds}.
	ShareDecaySeconds *float64 `json:"shareDecaySeconds" yaml:"shareDecaySeconds"`
	// share_distribution block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#share_distribution BatchSchedulingPolicy#share_distribution}
	ShareDistribution interface{} `json:"shareDistribution" yaml:"shareDistribution"`
}

type BatchSchedulingPolicyFairSharePolicyOutputReference interface {
	cdktf.ComplexObject
	ComputeReservation() *float64
	SetComputeReservation(val *float64)
	ComputeReservationInput() *float64
	InternalValue() *BatchSchedulingPolicyFairSharePolicy
	SetInternalValue(val *BatchSchedulingPolicyFairSharePolicy)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ShareDecaySeconds() *float64
	SetShareDecaySeconds(val *float64)
	ShareDecaySecondsInput() *float64
	ShareDistribution() interface{}
	SetShareDistribution(val interface{})
	ShareDistributionInput() interface{}
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
	ResetComputeReservation()
	ResetShareDecaySeconds()
	ResetShareDistribution()
}

// The jsii proxy struct for BatchSchedulingPolicyFairSharePolicyOutputReference
type jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ComputeReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ComputeReservationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) InternalValue() *BatchSchedulingPolicyFairSharePolicy {
	var returns *BatchSchedulingPolicyFairSharePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ShareDecaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shareDecaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ShareDecaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shareDecaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ShareDistribution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ShareDistributionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shareDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewBatchSchedulingPolicyFairSharePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) BatchSchedulingPolicyFairSharePolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.batch.BatchSchedulingPolicyFairSharePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewBatchSchedulingPolicyFairSharePolicyOutputReference_Override(b BatchSchedulingPolicyFairSharePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.BatchSchedulingPolicyFairSharePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		b,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) SetComputeReservation(val *float64) {
	_jsii_.Set(
		j,
		"computeReservation",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) SetInternalValue(val *BatchSchedulingPolicyFairSharePolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) SetShareDecaySeconds(val *float64) {
	_jsii_.Set(
		j,
		"shareDecaySeconds",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) SetShareDistribution(val interface{}) {
	_jsii_.Set(
		j,
		"shareDistribution",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ResetComputeReservation() {
	_jsii_.InvokeVoid(
		b,
		"resetComputeReservation",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ResetShareDecaySeconds() {
	_jsii_.InvokeVoid(
		b,
		"resetShareDecaySeconds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchSchedulingPolicyFairSharePolicyOutputReference) ResetShareDistribution() {
	_jsii_.InvokeVoid(
		b,
		"resetShareDistribution",
		nil, // no parameters
	)
}

type BatchSchedulingPolicyFairSharePolicyShareDistribution struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#share_identifier BatchSchedulingPolicy#share_identifier}.
	ShareIdentifier *string `json:"shareIdentifier" yaml:"shareIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_scheduling_policy#weight_factor BatchSchedulingPolicy#weight_factor}.
	WeightFactor *float64 `json:"weightFactor" yaml:"weightFactor"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment aws_batch_compute_environment}.
type DataAwsBatchComputeEnvironment interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ComputeEnvironmentName() *string
	SetComputeEnvironmentName(val *string)
	ComputeEnvironmentNameInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
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

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment aws_batch_compute_environment} Data Source.
func NewDataAwsBatchComputeEnvironment(scope constructs.Construct, id *string, config *DataAwsBatchComputeEnvironmentConfig) DataAwsBatchComputeEnvironment {
	_init_.Initialize()

	j := jsiiProxy_DataAwsBatchComputeEnvironment{}

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchComputeEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment aws_batch_compute_environment} Data Source.
func NewDataAwsBatchComputeEnvironment_Override(d DataAwsBatchComputeEnvironment, scope constructs.Construct, id *string, config *DataAwsBatchComputeEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchComputeEnvironment",
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

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) SetCount(val *float64) {
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

func (j *jsiiProxy_DataAwsBatchComputeEnvironment) SetTags(val *map[string]*string) {
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
		"hashicorp_aws.batch.DataAwsBatchComputeEnvironment",
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
		"hashicorp_aws.batch.DataAwsBatchComputeEnvironment",
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsBatchComputeEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

// AWS Batch.
type DataAwsBatchComputeEnvironmentConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment#compute_environment_name DataAwsBatchComputeEnvironment#compute_environment_name}.
	ComputeEnvironmentName *string `json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_compute_environment#tags DataAwsBatchComputeEnvironment#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue aws_batch_job_queue}.
type DataAwsBatchJobQueue interface {
	cdktf.TerraformDataSource
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
	Priority() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SchedulingPolicyArn() *string
	State() *string
	Status() *string
	StatusReason() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	ComputeEnvironmentOrder(index *string) DataAwsBatchJobQueueComputeEnvironmentOrder
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

func (j *jsiiProxy_DataAwsBatchJobQueue) Count() *float64 {
	var returns *float64
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

func (j *jsiiProxy_DataAwsBatchJobQueue) SchedulingPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingPolicyArn",
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

func (j *jsiiProxy_DataAwsBatchJobQueue) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueue) TagsInput() *map[string]*string {
	var returns *map[string]*string
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue aws_batch_job_queue} Data Source.
func NewDataAwsBatchJobQueue(scope constructs.Construct, id *string, config *DataAwsBatchJobQueueConfig) DataAwsBatchJobQueue {
	_init_.Initialize()

	j := jsiiProxy_DataAwsBatchJobQueue{}

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchJobQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue aws_batch_job_queue} Data Source.
func NewDataAwsBatchJobQueue_Override(d DataAwsBatchJobQueue, scope constructs.Construct, id *string, config *DataAwsBatchJobQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchJobQueue",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueue) SetCount(val *float64) {
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

func (j *jsiiProxy_DataAwsBatchJobQueue) SetTags(val *map[string]*string) {
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
		"hashicorp_aws.batch.DataAwsBatchJobQueue",
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
		"hashicorp_aws.batch.DataAwsBatchJobQueue",
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
func (d *jsiiProxy_DataAwsBatchJobQueue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsBatchJobQueue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsBatchJobQueueComputeEnvironmentOrder(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsBatchJobQueueComputeEnvironmentOrder {
	_init_.Initialize()

	j := jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder{}

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchJobQueueComputeEnvironmentOrder",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsBatchJobQueueComputeEnvironmentOrder_Override(d DataAwsBatchJobQueueComputeEnvironmentOrder, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchJobQueueComputeEnvironmentOrder",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
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

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsBatchJobQueueComputeEnvironmentOrder) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

// AWS Batch.
type DataAwsBatchJobQueueConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue#name DataAwsBatchJobQueue#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_job_queue#tags DataAwsBatchJobQueue#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/batch_scheduling_policy aws_batch_scheduling_policy}.
type DataAwsBatchSchedulingPolicy interface {
	cdktf.TerraformDataSource
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
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
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	FairSharePolicy(index *string) DataAwsBatchSchedulingPolicyFairSharePolicy
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
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsBatchSchedulingPolicy
type jsiiProxy_DataAwsBatchSchedulingPolicy struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_scheduling_policy aws_batch_scheduling_policy} Data Source.
func NewDataAwsBatchSchedulingPolicy(scope constructs.Construct, id *string, config *DataAwsBatchSchedulingPolicyConfig) DataAwsBatchSchedulingPolicy {
	_init_.Initialize()

	j := jsiiProxy_DataAwsBatchSchedulingPolicy{}

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchSchedulingPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/batch_scheduling_policy aws_batch_scheduling_policy} Data Source.
func NewDataAwsBatchSchedulingPolicy_Override(d DataAwsBatchSchedulingPolicy, scope constructs.Construct, id *string, config *DataAwsBatchSchedulingPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchSchedulingPolicy",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicy) SetTags(val *map[string]*string) {
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
func DataAwsBatchSchedulingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.batch.DataAwsBatchSchedulingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsBatchSchedulingPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.batch.DataAwsBatchSchedulingPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) FairSharePolicy(index *string) DataAwsBatchSchedulingPolicyFairSharePolicy {
	var returns DataAwsBatchSchedulingPolicyFairSharePolicy

	_jsii_.Invoke(
		d,
		"fairSharePolicy",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) ToString() *string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Batch.
type DataAwsBatchSchedulingPolicyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_scheduling_policy#arn DataAwsBatchSchedulingPolicy#arn}.
	Arn *string `json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/batch_scheduling_policy#tags DataAwsBatchSchedulingPolicy#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type DataAwsBatchSchedulingPolicyFairSharePolicy interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ComputeReservation() *float64
	ShareDecaySeconds() *float64
	ShareDistribution() cdktf.IResolvable
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

// The jsii proxy struct for DataAwsBatchSchedulingPolicyFairSharePolicy
type jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) ComputeReservation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeReservation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) ShareDecaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shareDecaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) ShareDistribution() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"shareDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsBatchSchedulingPolicyFairSharePolicy(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsBatchSchedulingPolicyFairSharePolicy {
	_init_.Initialize()

	j := jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy{}

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchSchedulingPolicyFairSharePolicy",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsBatchSchedulingPolicyFairSharePolicy_Override(d DataAwsBatchSchedulingPolicyFairSharePolicy, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchSchedulingPolicyFairSharePolicy",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicy) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ShareIdentifier() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeightFactor() *float64
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

// The jsii proxy struct for DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution
type jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) ShareIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) WeightFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"weightFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution {
	_init_.Initialize()

	j := jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution{}

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution_Override(d DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.batch.DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsBatchSchedulingPolicyFairSharePolicyShareDistribution) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}
