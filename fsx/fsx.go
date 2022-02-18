package fsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/fsx/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup aws_fsx_backup}.
type FsxBackup interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FileSystemId() *string
	SetFileSystemId(val *string)
	FileSystemIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KmsKeyId() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	OwnerId() *string
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
	Timeouts() FsxBackupTimeoutsOutputReference
	TimeoutsInput() *FsxBackupTimeouts
	Type() *string
	VolumeId() *string
	SetVolumeId(val *string)
	VolumeIdInput() *string
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
	PutTimeouts(value *FsxBackupTimeouts)
	ResetFileSystemId()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVolumeId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxBackup
type jsiiProxy_FsxBackup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxBackup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) FileSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) Timeouts() FsxBackupTimeoutsOutputReference {
	var returns FsxBackupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TimeoutsInput() *FsxBackupTimeouts {
	var returns *FsxBackupTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) VolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup aws_fsx_backup} Resource.
func NewFsxBackup(scope constructs.Construct, id *string, config *FsxBackupConfig) FsxBackup {
	_init_.Initialize()

	j := jsiiProxy_FsxBackup{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxBackup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup aws_fsx_backup} Resource.
func NewFsxBackup_Override(f FsxBackup, scope constructs.Construct, id *string, config *FsxBackupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxBackup",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxBackup) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxBackup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxBackup) SetFileSystemId(val *string) {
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_FsxBackup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxBackup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxBackup) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxBackup) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxBackup) SetVolumeId(val *string) {
	_jsii_.Set(
		j,
		"volumeId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FsxBackup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxBackup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxBackup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxBackup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FsxBackup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxBackup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxBackup) PutTimeouts(value *FsxBackupTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxBackup) ResetFileSystemId() {
	_jsii_.InvokeVoid(
		f,
		"resetFileSystemId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxBackup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxBackup) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxBackup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxBackup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxBackup) ResetVolumeId() {
	_jsii_.InvokeVoid(
		f,
		"resetVolumeId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxBackup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxBackup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxBackup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS File System FSx.
type FsxBackupConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup#file_system_id FsxBackup#file_system_id}.
	FileSystemId *string `json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup#tags FsxBackup#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup#tags_all FsxBackup#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup#timeouts FsxBackup#timeouts}
	Timeouts *FsxBackupTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup#volume_id FsxBackup#volume_id}.
	VolumeId *string `json:"volumeId" yaml:"volumeId"`
}

type FsxBackupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup#create FsxBackup#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup#delete FsxBackup#delete}.
	Delete *string `json:"delete" yaml:"delete"`
}

type FsxBackupTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxBackupTimeouts
	SetInternalValue(val *FsxBackupTimeouts)
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
	ResetCreate()
	ResetDelete()
}

// The jsii proxy struct for FsxBackupTimeoutsOutputReference
type jsiiProxy_FsxBackupTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) InternalValue() *FsxBackupTimeouts {
	var returns *FsxBackupTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxBackupTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxBackupTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxBackupTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxBackupTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxBackupTimeoutsOutputReference_Override(f FsxBackupTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxBackupTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) SetInternalValue(val *FsxBackupTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association aws_fsx_data_repository_association}.
type FsxDataRepositoryAssociation interface {
	cdktf.TerraformResource
	Arn() *string
	AssociationId() *string
	BatchImportMetaDataOnCreate() interface{}
	SetBatchImportMetaDataOnCreate(val interface{})
	BatchImportMetaDataOnCreateInput() interface{}
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DataRepositoryPath() *string
	SetDataRepositoryPath(val *string)
	DataRepositoryPathInput() *string
	DeleteDataInFilesystem() interface{}
	SetDeleteDataInFilesystem(val interface{})
	DeleteDataInFilesystemInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FileSystemId() *string
	SetFileSystemId(val *string)
	FileSystemIdInput() *string
	FileSystemPath() *string
	SetFileSystemPath(val *string)
	FileSystemPathInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	ImportedFileChunkSize() *float64
	SetImportedFileChunkSize(val *float64)
	ImportedFileChunkSizeInput() *float64
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	S3() FsxDataRepositoryAssociationS3OutputReference
	S3Input() *FsxDataRepositoryAssociationS3
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() FsxDataRepositoryAssociationTimeoutsOutputReference
	TimeoutsInput() *FsxDataRepositoryAssociationTimeouts
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
	PutS3(value *FsxDataRepositoryAssociationS3)
	PutTimeouts(value *FsxDataRepositoryAssociationTimeouts)
	ResetBatchImportMetaDataOnCreate()
	ResetDeleteDataInFilesystem()
	ResetImportedFileChunkSize()
	ResetOverrideLogicalId()
	ResetS3()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxDataRepositoryAssociation
type jsiiProxy_FsxDataRepositoryAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) BatchImportMetaDataOnCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchImportMetaDataOnCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) BatchImportMetaDataOnCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"batchImportMetaDataOnCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) DataRepositoryPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataRepositoryPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) DataRepositoryPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataRepositoryPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) DeleteDataInFilesystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteDataInFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) DeleteDataInFilesystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteDataInFilesystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FileSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FileSystemPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FileSystemPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) ImportedFileChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) ImportedFileChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) S3() FsxDataRepositoryAssociationS3OutputReference {
	var returns FsxDataRepositoryAssociationS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) S3Input() *FsxDataRepositoryAssociationS3 {
	var returns *FsxDataRepositoryAssociationS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) Timeouts() FsxDataRepositoryAssociationTimeoutsOutputReference {
	var returns FsxDataRepositoryAssociationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) TimeoutsInput() *FsxDataRepositoryAssociationTimeouts {
	var returns *FsxDataRepositoryAssociationTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association aws_fsx_data_repository_association} Resource.
func NewFsxDataRepositoryAssociation(scope constructs.Construct, id *string, config *FsxDataRepositoryAssociationConfig) FsxDataRepositoryAssociation {
	_init_.Initialize()

	j := jsiiProxy_FsxDataRepositoryAssociation{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association aws_fsx_data_repository_association} Resource.
func NewFsxDataRepositoryAssociation_Override(f FsxDataRepositoryAssociation, scope constructs.Construct, id *string, config *FsxDataRepositoryAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociation",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetBatchImportMetaDataOnCreate(val interface{}) {
	_jsii_.Set(
		j,
		"batchImportMetaDataOnCreate",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetDataRepositoryPath(val *string) {
	_jsii_.Set(
		j,
		"dataRepositoryPath",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetDeleteDataInFilesystem(val interface{}) {
	_jsii_.Set(
		j,
		"deleteDataInFilesystem",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetFileSystemId(val *string) {
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetFileSystemPath(val *string) {
	_jsii_.Set(
		j,
		"fileSystemPath",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetImportedFileChunkSize(val *float64) {
	_jsii_.Set(
		j,
		"importedFileChunkSize",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociation) SetTagsAll(val *map[string]*string) {
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
func FsxDataRepositoryAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxDataRepositoryAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) PutS3(value *FsxDataRepositoryAssociationS3) {
	_jsii_.InvokeVoid(
		f,
		"putS3",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) PutTimeouts(value *FsxDataRepositoryAssociationTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetBatchImportMetaDataOnCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetBatchImportMetaDataOnCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetDeleteDataInFilesystem() {
	_jsii_.InvokeVoid(
		f,
		"resetDeleteDataInFilesystem",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetImportedFileChunkSize() {
	_jsii_.InvokeVoid(
		f,
		"resetImportedFileChunkSize",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetS3() {
	_jsii_.InvokeVoid(
		f,
		"resetS3",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxDataRepositoryAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS File System FSx.
type FsxDataRepositoryAssociationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#data_repository_path FsxDataRepositoryAssociation#data_repository_path}.
	DataRepositoryPath *string `json:"dataRepositoryPath" yaml:"dataRepositoryPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#file_system_id FsxDataRepositoryAssociation#file_system_id}.
	FileSystemId *string `json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#file_system_path FsxDataRepositoryAssociation#file_system_path}.
	FileSystemPath *string `json:"fileSystemPath" yaml:"fileSystemPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#batch_import_meta_data_on_create FsxDataRepositoryAssociation#batch_import_meta_data_on_create}.
	BatchImportMetaDataOnCreate interface{} `json:"batchImportMetaDataOnCreate" yaml:"batchImportMetaDataOnCreate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#delete_data_in_filesystem FsxDataRepositoryAssociation#delete_data_in_filesystem}.
	DeleteDataInFilesystem interface{} `json:"deleteDataInFilesystem" yaml:"deleteDataInFilesystem"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#imported_file_chunk_size FsxDataRepositoryAssociation#imported_file_chunk_size}.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#s3 FsxDataRepositoryAssociation#s3}
	S3 *FsxDataRepositoryAssociationS3 `json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#tags FsxDataRepositoryAssociation#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#tags_all FsxDataRepositoryAssociation#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#timeouts FsxDataRepositoryAssociation#timeouts}
	Timeouts *FsxDataRepositoryAssociationTimeouts `json:"timeouts" yaml:"timeouts"`
}

type FsxDataRepositoryAssociationS3 struct {
	// auto_export_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#auto_export_policy FsxDataRepositoryAssociation#auto_export_policy}
	AutoExportPolicy *FsxDataRepositoryAssociationS3AutoExportPolicy `json:"autoExportPolicy" yaml:"autoExportPolicy"`
	// auto_import_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#auto_import_policy FsxDataRepositoryAssociation#auto_import_policy}
	AutoImportPolicy *FsxDataRepositoryAssociationS3AutoImportPolicy `json:"autoImportPolicy" yaml:"autoImportPolicy"`
}

type FsxDataRepositoryAssociationS3AutoExportPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#events FsxDataRepositoryAssociation#events}.
	Events *[]*string `json:"events" yaml:"events"`
}

type FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference interface {
	cdktf.ComplexObject
	Events() *[]*string
	SetEvents(val *[]*string)
	EventsInput() *[]*string
	InternalValue() *FsxDataRepositoryAssociationS3AutoExportPolicy
	SetInternalValue(val *FsxDataRepositoryAssociationS3AutoExportPolicy)
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
	ResetEvents()
}

// The jsii proxy struct for FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference
type jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) Events() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) EventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) InternalValue() *FsxDataRepositoryAssociationS3AutoExportPolicy {
	var returns *FsxDataRepositoryAssociationS3AutoExportPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxDataRepositoryAssociationS3AutoExportPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxDataRepositoryAssociationS3AutoExportPolicyOutputReference_Override(f FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) SetEvents(val *[]*string) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) SetInternalValue(val *FsxDataRepositoryAssociationS3AutoExportPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference) ResetEvents() {
	_jsii_.InvokeVoid(
		f,
		"resetEvents",
		nil, // no parameters
	)
}

type FsxDataRepositoryAssociationS3AutoImportPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#events FsxDataRepositoryAssociation#events}.
	Events *[]*string `json:"events" yaml:"events"`
}

type FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference interface {
	cdktf.ComplexObject
	Events() *[]*string
	SetEvents(val *[]*string)
	EventsInput() *[]*string
	InternalValue() *FsxDataRepositoryAssociationS3AutoImportPolicy
	SetInternalValue(val *FsxDataRepositoryAssociationS3AutoImportPolicy)
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
	ResetEvents()
}

// The jsii proxy struct for FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference
type jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) Events() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) EventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) InternalValue() *FsxDataRepositoryAssociationS3AutoImportPolicy {
	var returns *FsxDataRepositoryAssociationS3AutoImportPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxDataRepositoryAssociationS3AutoImportPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxDataRepositoryAssociationS3AutoImportPolicyOutputReference_Override(f FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) SetEvents(val *[]*string) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) SetInternalValue(val *FsxDataRepositoryAssociationS3AutoImportPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference) ResetEvents() {
	_jsii_.InvokeVoid(
		f,
		"resetEvents",
		nil, // no parameters
	)
}

type FsxDataRepositoryAssociationS3OutputReference interface {
	cdktf.ComplexObject
	AutoExportPolicy() FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference
	AutoExportPolicyInput() *FsxDataRepositoryAssociationS3AutoExportPolicy
	AutoImportPolicy() FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference
	AutoImportPolicyInput() *FsxDataRepositoryAssociationS3AutoImportPolicy
	InternalValue() *FsxDataRepositoryAssociationS3
	SetInternalValue(val *FsxDataRepositoryAssociationS3)
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
	PutAutoExportPolicy(value *FsxDataRepositoryAssociationS3AutoExportPolicy)
	PutAutoImportPolicy(value *FsxDataRepositoryAssociationS3AutoImportPolicy)
	ResetAutoExportPolicy()
	ResetAutoImportPolicy()
}

// The jsii proxy struct for FsxDataRepositoryAssociationS3OutputReference
type jsiiProxy_FsxDataRepositoryAssociationS3OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) AutoExportPolicy() FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference {
	var returns FsxDataRepositoryAssociationS3AutoExportPolicyOutputReference
	_jsii_.Get(
		j,
		"autoExportPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) AutoExportPolicyInput() *FsxDataRepositoryAssociationS3AutoExportPolicy {
	var returns *FsxDataRepositoryAssociationS3AutoExportPolicy
	_jsii_.Get(
		j,
		"autoExportPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) AutoImportPolicy() FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference {
	var returns FsxDataRepositoryAssociationS3AutoImportPolicyOutputReference
	_jsii_.Get(
		j,
		"autoImportPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) AutoImportPolicyInput() *FsxDataRepositoryAssociationS3AutoImportPolicy {
	var returns *FsxDataRepositoryAssociationS3AutoImportPolicy
	_jsii_.Get(
		j,
		"autoImportPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) InternalValue() *FsxDataRepositoryAssociationS3 {
	var returns *FsxDataRepositoryAssociationS3
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxDataRepositoryAssociationS3OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxDataRepositoryAssociationS3OutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxDataRepositoryAssociationS3OutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociationS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxDataRepositoryAssociationS3OutputReference_Override(f FsxDataRepositoryAssociationS3OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociationS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) SetInternalValue(val *FsxDataRepositoryAssociationS3) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) PutAutoExportPolicy(value *FsxDataRepositoryAssociationS3AutoExportPolicy) {
	_jsii_.InvokeVoid(
		f,
		"putAutoExportPolicy",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) PutAutoImportPolicy(value *FsxDataRepositoryAssociationS3AutoImportPolicy) {
	_jsii_.InvokeVoid(
		f,
		"putAutoImportPolicy",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) ResetAutoExportPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetAutoExportPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociationS3OutputReference) ResetAutoImportPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetAutoImportPolicy",
		nil, // no parameters
	)
}

type FsxDataRepositoryAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#create FsxDataRepositoryAssociation#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#delete FsxDataRepositoryAssociation#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_data_repository_association#update FsxDataRepositoryAssociation#update}.
	Update *string `json:"update" yaml:"update"`
}

type FsxDataRepositoryAssociationTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxDataRepositoryAssociationTimeouts
	SetInternalValue(val *FsxDataRepositoryAssociationTimeouts)
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

// The jsii proxy struct for FsxDataRepositoryAssociationTimeoutsOutputReference
type jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) InternalValue() *FsxDataRepositoryAssociationTimeouts {
	var returns *FsxDataRepositoryAssociationTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewFsxDataRepositoryAssociationTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxDataRepositoryAssociationTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxDataRepositoryAssociationTimeoutsOutputReference_Override(f FsxDataRepositoryAssociationTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxDataRepositoryAssociationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) SetInternalValue(val *FsxDataRepositoryAssociationTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxDataRepositoryAssociationTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system aws_fsx_lustre_file_system}.
type FsxLustreFileSystem interface {
	cdktf.TerraformResource
	Arn() *string
	AutoImportPolicy() *string
	SetAutoImportPolicy(val *string)
	AutoImportPolicyInput() *string
	AutomaticBackupRetentionDays() *float64
	SetAutomaticBackupRetentionDays(val *float64)
	AutomaticBackupRetentionDaysInput() *float64
	BackupId() *string
	SetBackupId(val *string)
	BackupIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToBackups() interface{}
	SetCopyTagsToBackups(val interface{})
	CopyTagsToBackupsInput() interface{}
	Count() *float64
	SetCount(val *float64)
	DailyAutomaticBackupStartTime() *string
	SetDailyAutomaticBackupStartTime(val *string)
	DailyAutomaticBackupStartTimeInput() *string
	DataCompressionType() *string
	SetDataCompressionType(val *string)
	DataCompressionTypeInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
	DnsName() *string
	DriveCacheType() *string
	SetDriveCacheType(val *string)
	DriveCacheTypeInput() *string
	ExportPath() *string
	SetExportPath(val *string)
	ExportPathInput() *string
	FileSystemTypeVersion() *string
	SetFileSystemTypeVersion(val *string)
	FileSystemTypeVersionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	ImportedFileChunkSize() *float64
	SetImportedFileChunkSize(val *float64)
	ImportedFileChunkSizeInput() *float64
	ImportPath() *string
	SetImportPath(val *string)
	ImportPathInput() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MountName() *string
	NetworkInterfaceIds() *[]*string
	Node() constructs.Node
	OwnerId() *string
	PerUnitStorageThroughput() *float64
	SetPerUnitStorageThroughput(val *float64)
	PerUnitStorageThroughputInput() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	StorageCapacity() *float64
	SetStorageCapacity(val *float64)
	StorageCapacityInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
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
	Timeouts() FsxLustreFileSystemTimeoutsOutputReference
	TimeoutsInput() *FsxLustreFileSystemTimeouts
	VpcId() *string
	WeeklyMaintenanceStartTime() *string
	SetWeeklyMaintenanceStartTime(val *string)
	WeeklyMaintenanceStartTimeInput() *string
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
	PutTimeouts(value *FsxLustreFileSystemTimeouts)
	ResetAutoImportPolicy()
	ResetAutomaticBackupRetentionDays()
	ResetBackupId()
	ResetCopyTagsToBackups()
	ResetDailyAutomaticBackupStartTime()
	ResetDataCompressionType()
	ResetDeploymentType()
	ResetDriveCacheType()
	ResetExportPath()
	ResetFileSystemTypeVersion()
	ResetImportedFileChunkSize()
	ResetImportPath()
	ResetKmsKeyId()
	ResetOverrideLogicalId()
	ResetPerUnitStorageThroughput()
	ResetSecurityGroupIds()
	ResetStorageCapacity()
	ResetStorageType()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetWeeklyMaintenanceStartTime()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxLustreFileSystem
type jsiiProxy_FsxLustreFileSystem struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxLustreFileSystem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) AutoImportPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) AutoImportPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) AutomaticBackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) AutomaticBackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) BackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) CopyTagsToBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) CopyTagsToBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DailyAutomaticBackupStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DailyAutomaticBackupStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DriveCacheType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) DriveCacheTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ExportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ExportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) FileSystemTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) FileSystemTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ImportedFileChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ImportedFileChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ImportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) ImportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) MountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) PerUnitStorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) PerUnitStorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) StorageCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) Timeouts() FsxLustreFileSystemTimeoutsOutputReference {
	var returns FsxLustreFileSystemTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TimeoutsInput() *FsxLustreFileSystemTimeouts {
	var returns *FsxLustreFileSystemTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) WeeklyMaintenanceStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTimeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system aws_fsx_lustre_file_system} Resource.
func NewFsxLustreFileSystem(scope constructs.Construct, id *string, config *FsxLustreFileSystemConfig) FsxLustreFileSystem {
	_init_.Initialize()

	j := jsiiProxy_FsxLustreFileSystem{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxLustreFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system aws_fsx_lustre_file_system} Resource.
func NewFsxLustreFileSystem_Override(f FsxLustreFileSystem, scope constructs.Construct, id *string, config *FsxLustreFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxLustreFileSystem",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetAutoImportPolicy(val *string) {
	_jsii_.Set(
		j,
		"autoImportPolicy",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetAutomaticBackupRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"automaticBackupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetBackupId(val *string) {
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetCopyTagsToBackups(val interface{}) {
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetDailyAutomaticBackupStartTime(val *string) {
	_jsii_.Set(
		j,
		"dailyAutomaticBackupStartTime",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetDataCompressionType(val *string) {
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetDeploymentType(val *string) {
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetDriveCacheType(val *string) {
	_jsii_.Set(
		j,
		"driveCacheType",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetExportPath(val *string) {
	_jsii_.Set(
		j,
		"exportPath",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetFileSystemTypeVersion(val *string) {
	_jsii_.Set(
		j,
		"fileSystemTypeVersion",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetImportedFileChunkSize(val *float64) {
	_jsii_.Set(
		j,
		"importedFileChunkSize",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetImportPath(val *string) {
	_jsii_.Set(
		j,
		"importPath",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetPerUnitStorageThroughput(val *float64) {
	_jsii_.Set(
		j,
		"perUnitStorageThroughput",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetStorageCapacity(val *float64) {
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetWeeklyMaintenanceStartTime(val *string) {
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FsxLustreFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxLustreFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxLustreFileSystem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxLustreFileSystem",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) PutTimeouts(value *FsxLustreFileSystemTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetAutoImportPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetAutoImportPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetAutomaticBackupRetentionDays() {
	_jsii_.InvokeVoid(
		f,
		"resetAutomaticBackupRetentionDays",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetBackupId() {
	_jsii_.InvokeVoid(
		f,
		"resetBackupId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetDailyAutomaticBackupStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetDailyAutomaticBackupStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		f,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetDeploymentType() {
	_jsii_.InvokeVoid(
		f,
		"resetDeploymentType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetDriveCacheType() {
	_jsii_.InvokeVoid(
		f,
		"resetDriveCacheType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetExportPath() {
	_jsii_.InvokeVoid(
		f,
		"resetExportPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetFileSystemTypeVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetFileSystemTypeVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetImportedFileChunkSize() {
	_jsii_.InvokeVoid(
		f,
		"resetImportedFileChunkSize",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetImportPath() {
	_jsii_.InvokeVoid(
		f,
		"resetImportPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		f,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetPerUnitStorageThroughput() {
	_jsii_.InvokeVoid(
		f,
		"resetPerUnitStorageThroughput",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		f,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetStorageCapacity() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetStorageType() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxLustreFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxLustreFileSystem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS File System FSx.
type FsxLustreFileSystemConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#subnet_ids FsxLustreFileSystem#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#auto_import_policy FsxLustreFileSystem#auto_import_policy}.
	AutoImportPolicy *string `json:"autoImportPolicy" yaml:"autoImportPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#automatic_backup_retention_days FsxLustreFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#backup_id FsxLustreFileSystem#backup_id}.
	BackupId *string `json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#copy_tags_to_backups FsxLustreFileSystem#copy_tags_to_backups}.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#daily_automatic_backup_start_time FsxLustreFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#data_compression_type FsxLustreFileSystem#data_compression_type}.
	DataCompressionType *string `json:"dataCompressionType" yaml:"dataCompressionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#deployment_type FsxLustreFileSystem#deployment_type}.
	DeploymentType *string `json:"deploymentType" yaml:"deploymentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#drive_cache_type FsxLustreFileSystem#drive_cache_type}.
	DriveCacheType *string `json:"driveCacheType" yaml:"driveCacheType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#export_path FsxLustreFileSystem#export_path}.
	ExportPath *string `json:"exportPath" yaml:"exportPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#file_system_type_version FsxLustreFileSystem#file_system_type_version}.
	FileSystemTypeVersion *string `json:"fileSystemTypeVersion" yaml:"fileSystemTypeVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#imported_file_chunk_size FsxLustreFileSystem#imported_file_chunk_size}.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#import_path FsxLustreFileSystem#import_path}.
	ImportPath *string `json:"importPath" yaml:"importPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#kms_key_id FsxLustreFileSystem#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#per_unit_storage_throughput FsxLustreFileSystem#per_unit_storage_throughput}.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#security_group_ids FsxLustreFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#storage_capacity FsxLustreFileSystem#storage_capacity}.
	StorageCapacity *float64 `json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#storage_type FsxLustreFileSystem#storage_type}.
	StorageType *string `json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#tags FsxLustreFileSystem#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#tags_all FsxLustreFileSystem#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#timeouts FsxLustreFileSystem#timeouts}
	Timeouts *FsxLustreFileSystemTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#weekly_maintenance_start_time FsxLustreFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

type FsxLustreFileSystemTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#create FsxLustreFileSystem#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#delete FsxLustreFileSystem#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system#update FsxLustreFileSystem#update}.
	Update *string `json:"update" yaml:"update"`
}

type FsxLustreFileSystemTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxLustreFileSystemTimeouts
	SetInternalValue(val *FsxLustreFileSystemTimeouts)
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

// The jsii proxy struct for FsxLustreFileSystemTimeoutsOutputReference
type jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) InternalValue() *FsxLustreFileSystemTimeouts {
	var returns *FsxLustreFileSystemTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewFsxLustreFileSystemTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxLustreFileSystemTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxLustreFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxLustreFileSystemTimeoutsOutputReference_Override(f FsxLustreFileSystemTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxLustreFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) SetInternalValue(val *FsxLustreFileSystemTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system aws_fsx_ontap_file_system}.
type FsxOntapFileSystem interface {
	cdktf.TerraformResource
	Arn() *string
	AutomaticBackupRetentionDays() *float64
	SetAutomaticBackupRetentionDays(val *float64)
	AutomaticBackupRetentionDaysInput() *float64
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DailyAutomaticBackupStartTime() *string
	SetDailyAutomaticBackupStartTime(val *string)
	DailyAutomaticBackupStartTimeInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
	DiskIopsConfiguration() FsxOntapFileSystemDiskIopsConfigurationOutputReference
	DiskIopsConfigurationInput() *FsxOntapFileSystemDiskIopsConfiguration
	DnsName() *string
	EndpointIpAddressRange() *string
	SetEndpointIpAddressRange(val *string)
	EndpointIpAddressRangeInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	FsxAdminPassword() *string
	SetFsxAdminPassword(val *string)
	FsxAdminPasswordInput() *string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkInterfaceIds() *[]*string
	Node() constructs.Node
	OwnerId() *string
	PreferredSubnetId() *string
	SetPreferredSubnetId(val *string)
	PreferredSubnetIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RouteTableIds() *[]*string
	SetRouteTableIds(val *[]*string)
	RouteTableIdsInput() *[]*string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	StorageCapacity() *float64
	SetStorageCapacity(val *float64)
	StorageCapacityInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
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
	ThroughputCapacity() *float64
	SetThroughputCapacity(val *float64)
	ThroughputCapacityInput() *float64
	Timeouts() FsxOntapFileSystemTimeoutsOutputReference
	TimeoutsInput() *FsxOntapFileSystemTimeouts
	VpcId() *string
	WeeklyMaintenanceStartTime() *string
	SetWeeklyMaintenanceStartTime(val *string)
	WeeklyMaintenanceStartTimeInput() *string
	AddOverride(path *string, value interface{})
	Endpoints(index *string) FsxOntapFileSystemEndpoints
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
	PutDiskIopsConfiguration(value *FsxOntapFileSystemDiskIopsConfiguration)
	PutTimeouts(value *FsxOntapFileSystemTimeouts)
	ResetAutomaticBackupRetentionDays()
	ResetDailyAutomaticBackupStartTime()
	ResetDiskIopsConfiguration()
	ResetEndpointIpAddressRange()
	ResetFsxAdminPassword()
	ResetKmsKeyId()
	ResetOverrideLogicalId()
	ResetRouteTableIds()
	ResetSecurityGroupIds()
	ResetStorageCapacity()
	ResetStorageType()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetWeeklyMaintenanceStartTime()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxOntapFileSystem
type jsiiProxy_FsxOntapFileSystem struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxOntapFileSystem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) AutomaticBackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) AutomaticBackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) DailyAutomaticBackupStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) DailyAutomaticBackupStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) DiskIopsConfiguration() FsxOntapFileSystemDiskIopsConfigurationOutputReference {
	var returns FsxOntapFileSystemDiskIopsConfigurationOutputReference
	_jsii_.Get(
		j,
		"diskIopsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) DiskIopsConfigurationInput() *FsxOntapFileSystemDiskIopsConfiguration {
	var returns *FsxOntapFileSystemDiskIopsConfiguration
	_jsii_.Get(
		j,
		"diskIopsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) EndpointIpAddressRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) EndpointIpAddressRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) FsxAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsxAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) FsxAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsxAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) PreferredSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) PreferredSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) RouteTableIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTableIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) RouteTableIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routeTableIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) StorageCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) ThroughputCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) ThroughputCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) Timeouts() FsxOntapFileSystemTimeoutsOutputReference {
	var returns FsxOntapFileSystemTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TimeoutsInput() *FsxOntapFileSystemTimeouts {
	var returns *FsxOntapFileSystemTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) WeeklyMaintenanceStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTimeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system aws_fsx_ontap_file_system} Resource.
func NewFsxOntapFileSystem(scope constructs.Construct, id *string, config *FsxOntapFileSystemConfig) FsxOntapFileSystem {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystem{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system aws_fsx_ontap_file_system} Resource.
func NewFsxOntapFileSystem_Override(f FsxOntapFileSystem, scope constructs.Construct, id *string, config *FsxOntapFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystem",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetAutomaticBackupRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"automaticBackupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetDailyAutomaticBackupStartTime(val *string) {
	_jsii_.Set(
		j,
		"dailyAutomaticBackupStartTime",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetDeploymentType(val *string) {
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetEndpointIpAddressRange(val *string) {
	_jsii_.Set(
		j,
		"endpointIpAddressRange",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetFsxAdminPassword(val *string) {
	_jsii_.Set(
		j,
		"fsxAdminPassword",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetPreferredSubnetId(val *string) {
	_jsii_.Set(
		j,
		"preferredSubnetId",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetRouteTableIds(val *[]*string) {
	_jsii_.Set(
		j,
		"routeTableIds",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetStorageCapacity(val *float64) {
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetThroughputCapacity(val *float64) {
	_jsii_.Set(
		j,
		"throughputCapacity",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetWeeklyMaintenanceStartTime(val *string) {
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FsxOntapFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxOntapFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxOntapFileSystem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxOntapFileSystem",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) Endpoints(index *string) FsxOntapFileSystemEndpoints {
	var returns FsxOntapFileSystemEndpoints

	_jsii_.Invoke(
		f,
		"endpoints",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) PutDiskIopsConfiguration(value *FsxOntapFileSystemDiskIopsConfiguration) {
	_jsii_.InvokeVoid(
		f,
		"putDiskIopsConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) PutTimeouts(value *FsxOntapFileSystemTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetAutomaticBackupRetentionDays() {
	_jsii_.InvokeVoid(
		f,
		"resetAutomaticBackupRetentionDays",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetDailyAutomaticBackupStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetDailyAutomaticBackupStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetDiskIopsConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetDiskIopsConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetEndpointIpAddressRange() {
	_jsii_.InvokeVoid(
		f,
		"resetEndpointIpAddressRange",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetFsxAdminPassword() {
	_jsii_.InvokeVoid(
		f,
		"resetFsxAdminPassword",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		f,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetRouteTableIds() {
	_jsii_.InvokeVoid(
		f,
		"resetRouteTableIds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		f,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetStorageCapacity() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetStorageType() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxOntapFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxOntapFileSystem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS File System FSx.
type FsxOntapFileSystemConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#deployment_type FsxOntapFileSystem#deployment_type}.
	DeploymentType *string `json:"deploymentType" yaml:"deploymentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#preferred_subnet_id FsxOntapFileSystem#preferred_subnet_id}.
	PreferredSubnetId *string `json:"preferredSubnetId" yaml:"preferredSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#subnet_ids FsxOntapFileSystem#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#throughput_capacity FsxOntapFileSystem#throughput_capacity}.
	ThroughputCapacity *float64 `json:"throughputCapacity" yaml:"throughputCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#automatic_backup_retention_days FsxOntapFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#daily_automatic_backup_start_time FsxOntapFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// disk_iops_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#disk_iops_configuration FsxOntapFileSystem#disk_iops_configuration}
	DiskIopsConfiguration *FsxOntapFileSystemDiskIopsConfiguration `json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#endpoint_ip_address_range FsxOntapFileSystem#endpoint_ip_address_range}.
	EndpointIpAddressRange *string `json:"endpointIpAddressRange" yaml:"endpointIpAddressRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#fsx_admin_password FsxOntapFileSystem#fsx_admin_password}.
	FsxAdminPassword *string `json:"fsxAdminPassword" yaml:"fsxAdminPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#kms_key_id FsxOntapFileSystem#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#route_table_ids FsxOntapFileSystem#route_table_ids}.
	RouteTableIds *[]*string `json:"routeTableIds" yaml:"routeTableIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#security_group_ids FsxOntapFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#storage_capacity FsxOntapFileSystem#storage_capacity}.
	StorageCapacity *float64 `json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#storage_type FsxOntapFileSystem#storage_type}.
	StorageType *string `json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#tags FsxOntapFileSystem#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#tags_all FsxOntapFileSystem#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#timeouts FsxOntapFileSystem#timeouts}
	Timeouts *FsxOntapFileSystemTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#weekly_maintenance_start_time FsxOntapFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

type FsxOntapFileSystemDiskIopsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#iops FsxOntapFileSystem#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#mode FsxOntapFileSystem#mode}.
	Mode *string `json:"mode" yaml:"mode"`
}

type FsxOntapFileSystemDiskIopsConfigurationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *FsxOntapFileSystemDiskIopsConfiguration
	SetInternalValue(val *FsxOntapFileSystemDiskIopsConfiguration)
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
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
	ResetIops()
	ResetMode()
}

// The jsii proxy struct for FsxOntapFileSystemDiskIopsConfigurationOutputReference
type jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) InternalValue() *FsxOntapFileSystemDiskIopsConfiguration {
	var returns *FsxOntapFileSystemDiskIopsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxOntapFileSystemDiskIopsConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOntapFileSystemDiskIopsConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemDiskIopsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOntapFileSystemDiskIopsConfigurationOutputReference_Override(f FsxOntapFileSystemDiskIopsConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemDiskIopsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) SetInternalValue(val *FsxOntapFileSystemDiskIopsConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) SetIops(val *float64) {
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		f,
		"resetIops",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		f,
		"resetMode",
		nil, // no parameters
	)
}

type FsxOntapFileSystemEndpoints interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Intercluster() cdktf.IResolvable
	Management() cdktf.IResolvable
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

// The jsii proxy struct for FsxOntapFileSystemEndpoints
type jsiiProxy_FsxOntapFileSystemEndpoints struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) Intercluster() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"intercluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) Management() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"management",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapFileSystemEndpoints(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) FsxOntapFileSystemEndpoints {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemEndpoints{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemEndpoints",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapFileSystemEndpoints_Override(f FsxOntapFileSystemEndpoints, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemEndpoints",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOntapFileSystemEndpointsIntercluster interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DnsName() *string
	IpAddresses() *[]*string
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

// The jsii proxy struct for FsxOntapFileSystemEndpointsIntercluster
type jsiiProxy_FsxOntapFileSystemEndpointsIntercluster struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapFileSystemEndpointsIntercluster(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) FsxOntapFileSystemEndpointsIntercluster {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemEndpointsIntercluster{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemEndpointsIntercluster",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapFileSystemEndpointsIntercluster_Override(f FsxOntapFileSystemEndpointsIntercluster, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemEndpointsIntercluster",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOntapFileSystemEndpointsManagement interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DnsName() *string
	IpAddresses() *[]*string
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

// The jsii proxy struct for FsxOntapFileSystemEndpointsManagement
type jsiiProxy_FsxOntapFileSystemEndpointsManagement struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapFileSystemEndpointsManagement(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) FsxOntapFileSystemEndpointsManagement {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemEndpointsManagement{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemEndpointsManagement",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapFileSystemEndpointsManagement_Override(f FsxOntapFileSystemEndpointsManagement, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemEndpointsManagement",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOntapFileSystemTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#create FsxOntapFileSystem#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#delete FsxOntapFileSystem#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system#update FsxOntapFileSystem#update}.
	Update *string `json:"update" yaml:"update"`
}

type FsxOntapFileSystemTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxOntapFileSystemTimeouts
	SetInternalValue(val *FsxOntapFileSystemTimeouts)
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

// The jsii proxy struct for FsxOntapFileSystemTimeoutsOutputReference
type jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) InternalValue() *FsxOntapFileSystemTimeouts {
	var returns *FsxOntapFileSystemTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewFsxOntapFileSystemTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOntapFileSystemTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOntapFileSystemTimeoutsOutputReference_Override(f FsxOntapFileSystemTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) SetInternalValue(val *FsxOntapFileSystemTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine aws_fsx_ontap_storage_virtual_machine}.
type FsxOntapStorageVirtualMachine interface {
	cdktf.TerraformResource
	ActiveDirectoryConfiguration() FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference
	ActiveDirectoryConfigurationInput() *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FileSystemId() *string
	SetFileSystemId(val *string)
	FileSystemIdInput() *string
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
	RootVolumeSecurityStyle() *string
	SetRootVolumeSecurityStyle(val *string)
	RootVolumeSecurityStyleInput() *string
	Subtype() *string
	SvmAdminPassword() *string
	SetSvmAdminPassword(val *string)
	SvmAdminPasswordInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() FsxOntapStorageVirtualMachineTimeoutsOutputReference
	TimeoutsInput() *FsxOntapStorageVirtualMachineTimeouts
	Uuid() *string
	AddOverride(path *string, value interface{})
	Endpoints(index *string) FsxOntapStorageVirtualMachineEndpoints
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
	PutActiveDirectoryConfiguration(value *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration)
	PutTimeouts(value *FsxOntapStorageVirtualMachineTimeouts)
	ResetActiveDirectoryConfiguration()
	ResetOverrideLogicalId()
	ResetRootVolumeSecurityStyle()
	ResetSvmAdminPassword()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxOntapStorageVirtualMachine
type jsiiProxy_FsxOntapStorageVirtualMachine struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) ActiveDirectoryConfiguration() FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference {
	var returns FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference
	_jsii_.Get(
		j,
		"activeDirectoryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) ActiveDirectoryConfigurationInput() *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration {
	var returns *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration
	_jsii_.Get(
		j,
		"activeDirectoryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) FileSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) RootVolumeSecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootVolumeSecurityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) RootVolumeSecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootVolumeSecurityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Subtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SvmAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"svmAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SvmAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"svmAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Timeouts() FsxOntapStorageVirtualMachineTimeoutsOutputReference {
	var returns FsxOntapStorageVirtualMachineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) TimeoutsInput() *FsxOntapStorageVirtualMachineTimeouts {
	var returns *FsxOntapStorageVirtualMachineTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine aws_fsx_ontap_storage_virtual_machine} Resource.
func NewFsxOntapStorageVirtualMachine(scope constructs.Construct, id *string, config *FsxOntapStorageVirtualMachineConfig) FsxOntapStorageVirtualMachine {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapStorageVirtualMachine{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine aws_fsx_ontap_storage_virtual_machine} Resource.
func NewFsxOntapStorageVirtualMachine_Override(f FsxOntapStorageVirtualMachine, scope constructs.Construct, id *string, config *FsxOntapStorageVirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachine",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetFileSystemId(val *string) {
	_jsii_.Set(
		j,
		"fileSystemId",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetRootVolumeSecurityStyle(val *string) {
	_jsii_.Set(
		j,
		"rootVolumeSecurityStyle",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetSvmAdminPassword(val *string) {
	_jsii_.Set(
		j,
		"svmAdminPassword",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachine) SetTagsAll(val *map[string]*string) {
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
func FsxOntapStorageVirtualMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxOntapStorageVirtualMachine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachine",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) Endpoints(index *string) FsxOntapStorageVirtualMachineEndpoints {
	var returns FsxOntapStorageVirtualMachineEndpoints

	_jsii_.Invoke(
		f,
		"endpoints",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) PutActiveDirectoryConfiguration(value *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration) {
	_jsii_.InvokeVoid(
		f,
		"putActiveDirectoryConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) PutTimeouts(value *FsxOntapStorageVirtualMachineTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ResetActiveDirectoryConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetActiveDirectoryConfiguration",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ResetRootVolumeSecurityStyle() {
	_jsii_.InvokeVoid(
		f,
		"resetRootVolumeSecurityStyle",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ResetSvmAdminPassword() {
	_jsii_.InvokeVoid(
		f,
		"resetSvmAdminPassword",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FsxOntapStorageVirtualMachineActiveDirectoryConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#netbios_name FsxOntapStorageVirtualMachine#netbios_name}.
	NetbiosName *string `json:"netbiosName" yaml:"netbiosName"`
	// self_managed_active_directory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#self_managed_active_directory_configuration FsxOntapStorageVirtualMachine#self_managed_active_directory_configuration}
	SelfManagedActiveDirectoryConfiguration *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration `json:"selfManagedActiveDirectoryConfiguration" yaml:"selfManagedActiveDirectoryConfiguration"`
}

type FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration
	SetInternalValue(val *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	NetbiosName() *string
	SetNetbiosName(val *string)
	NetbiosNameInput() *string
	SelfManagedActiveDirectoryConfiguration() FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference
	SelfManagedActiveDirectoryConfigurationInput() *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration
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
	PutSelfManagedActiveDirectoryConfiguration(value *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration)
	ResetNetbiosName()
	ResetSelfManagedActiveDirectoryConfiguration()
}

// The jsii proxy struct for FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference
type jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) InternalValue() *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration {
	var returns *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) NetbiosName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) NetbiosNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) SelfManagedActiveDirectoryConfiguration() FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference {
	var returns FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference
	_jsii_.Get(
		j,
		"selfManagedActiveDirectoryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) SelfManagedActiveDirectoryConfigurationInput() *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration {
	var returns *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration
	_jsii_.Get(
		j,
		"selfManagedActiveDirectoryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference_Override(f FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) SetInternalValue(val *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) SetNetbiosName(val *string) {
	_jsii_.Set(
		j,
		"netbiosName",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) PutSelfManagedActiveDirectoryConfiguration(value *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration) {
	_jsii_.InvokeVoid(
		f,
		"putSelfManagedActiveDirectoryConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) ResetNetbiosName() {
	_jsii_.InvokeVoid(
		f,
		"resetNetbiosName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationOutputReference) ResetSelfManagedActiveDirectoryConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetSelfManagedActiveDirectoryConfiguration",
		nil, // no parameters
	)
}

type FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#dns_ips FsxOntapStorageVirtualMachine#dns_ips}.
	DnsIps *[]*string `json:"dnsIps" yaml:"dnsIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#domain_name FsxOntapStorageVirtualMachine#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#password FsxOntapStorageVirtualMachine#password}.
	Password *string `json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#username FsxOntapStorageVirtualMachine#username}.
	Username *string `json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#file_system_administrators_group FsxOntapStorageVirtualMachine#file_system_administrators_group}.
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#organizational_unit_distinguidshed_name FsxOntapStorageVirtualMachine#organizational_unit_distinguidshed_name}.
	OrganizationalUnitDistinguidshedName *string `json:"organizationalUnitDistinguidshedName" yaml:"organizationalUnitDistinguidshedName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#organizational_unit_distinguished_name FsxOntapStorageVirtualMachine#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

type FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference interface {
	cdktf.ComplexObject
	DnsIps() *[]*string
	SetDnsIps(val *[]*string)
	DnsIpsInput() *[]*string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	FileSystemAdministratorsGroup() *string
	SetFileSystemAdministratorsGroup(val *string)
	FileSystemAdministratorsGroupInput() *string
	InternalValue() *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration
	SetInternalValue(val *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OrganizationalUnitDistinguidshedName() *string
	SetOrganizationalUnitDistinguidshedName(val *string)
	OrganizationalUnitDistinguidshedNameInput() *string
	OrganizationalUnitDistinguishedName() *string
	SetOrganizationalUnitDistinguishedName(val *string)
	OrganizationalUnitDistinguishedNameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetFileSystemAdministratorsGroup()
	ResetOrganizationalUnitDistinguidshedName()
	ResetOrganizationalUnitDistinguishedName()
}

// The jsii proxy struct for FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference
type jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) DnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) DnsIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) FileSystemAdministratorsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemAdministratorsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) FileSystemAdministratorsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemAdministratorsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) InternalValue() *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration {
	var returns *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) OrganizationalUnitDistinguidshedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguidshedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) OrganizationalUnitDistinguidshedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguidshedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) OrganizationalUnitDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) OrganizationalUnitDistinguishedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func NewFsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference_Override(f FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetDnsIps(val *[]*string) {
	_jsii_.Set(
		j,
		"dnsIps",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetFileSystemAdministratorsGroup(val *string) {
	_jsii_.Set(
		j,
		"fileSystemAdministratorsGroup",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetInternalValue(val *FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetOrganizationalUnitDistinguidshedName(val *string) {
	_jsii_.Set(
		j,
		"organizationalUnitDistinguidshedName",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetOrganizationalUnitDistinguishedName(val *string) {
	_jsii_.Set(
		j,
		"organizationalUnitDistinguishedName",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) ResetFileSystemAdministratorsGroup() {
	_jsii_.InvokeVoid(
		f,
		"resetFileSystemAdministratorsGroup",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) ResetOrganizationalUnitDistinguidshedName() {
	_jsii_.InvokeVoid(
		f,
		"resetOrganizationalUnitDistinguidshedName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationOutputReference) ResetOrganizationalUnitDistinguishedName() {
	_jsii_.InvokeVoid(
		f,
		"resetOrganizationalUnitDistinguishedName",
		nil, // no parameters
	)
}

// AWS File System FSx.
type FsxOntapStorageVirtualMachineConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#file_system_id FsxOntapStorageVirtualMachine#file_system_id}.
	FileSystemId *string `json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#name FsxOntapStorageVirtualMachine#name}.
	Name *string `json:"name" yaml:"name"`
	// active_directory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#active_directory_configuration FsxOntapStorageVirtualMachine#active_directory_configuration}
	ActiveDirectoryConfiguration *FsxOntapStorageVirtualMachineActiveDirectoryConfiguration `json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#root_volume_security_style FsxOntapStorageVirtualMachine#root_volume_security_style}.
	RootVolumeSecurityStyle *string `json:"rootVolumeSecurityStyle" yaml:"rootVolumeSecurityStyle"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#svm_admin_password FsxOntapStorageVirtualMachine#svm_admin_password}.
	SvmAdminPassword *string `json:"svmAdminPassword" yaml:"svmAdminPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#tags FsxOntapStorageVirtualMachine#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#tags_all FsxOntapStorageVirtualMachine#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#timeouts FsxOntapStorageVirtualMachine#timeouts}
	Timeouts *FsxOntapStorageVirtualMachineTimeouts `json:"timeouts" yaml:"timeouts"`
}

type FsxOntapStorageVirtualMachineEndpoints interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Iscsi() cdktf.IResolvable
	Management() cdktf.IResolvable
	Nfs() cdktf.IResolvable
	Smb() cdktf.IResolvable
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

// The jsii proxy struct for FsxOntapStorageVirtualMachineEndpoints
type jsiiProxy_FsxOntapStorageVirtualMachineEndpoints struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) Iscsi() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) Management() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"management",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) Nfs() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) Smb() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"smb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpoints(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) FsxOntapStorageVirtualMachineEndpoints {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapStorageVirtualMachineEndpoints{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpoints",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpoints_Override(f FsxOntapStorageVirtualMachineEndpoints, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpoints",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpoints) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOntapStorageVirtualMachineEndpointsIscsi interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DnsName() *string
	IpAddresses() *[]*string
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

// The jsii proxy struct for FsxOntapStorageVirtualMachineEndpointsIscsi
type jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpointsIscsi(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) FsxOntapStorageVirtualMachineEndpointsIscsi {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpointsIscsi",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpointsIscsi_Override(f FsxOntapStorageVirtualMachineEndpointsIscsi, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpointsIscsi",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsIscsi) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOntapStorageVirtualMachineEndpointsManagement interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DnsName() *string
	IpAddresses() *[]*string
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

// The jsii proxy struct for FsxOntapStorageVirtualMachineEndpointsManagement
type jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpointsManagement(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) FsxOntapStorageVirtualMachineEndpointsManagement {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpointsManagement",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpointsManagement_Override(f FsxOntapStorageVirtualMachineEndpointsManagement, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpointsManagement",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsManagement) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOntapStorageVirtualMachineEndpointsNfs interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DnsName() *string
	IpAddresses() *[]*string
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

// The jsii proxy struct for FsxOntapStorageVirtualMachineEndpointsNfs
type jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpointsNfs(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) FsxOntapStorageVirtualMachineEndpointsNfs {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpointsNfs",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpointsNfs_Override(f FsxOntapStorageVirtualMachineEndpointsNfs, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpointsNfs",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsNfs) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOntapStorageVirtualMachineEndpointsSmb interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DnsName() *string
	IpAddresses() *[]*string
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

// The jsii proxy struct for FsxOntapStorageVirtualMachineEndpointsSmb
type jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpointsSmb(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) FsxOntapStorageVirtualMachineEndpointsSmb {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpointsSmb",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapStorageVirtualMachineEndpointsSmb_Override(f FsxOntapStorageVirtualMachineEndpointsSmb, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineEndpointsSmb",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		f,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineEndpointsSmb) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOntapStorageVirtualMachineTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#create FsxOntapStorageVirtualMachine#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#delete FsxOntapStorageVirtualMachine#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_storage_virtual_machine#update FsxOntapStorageVirtualMachine#update}.
	Update *string `json:"update" yaml:"update"`
}

type FsxOntapStorageVirtualMachineTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxOntapStorageVirtualMachineTimeouts
	SetInternalValue(val *FsxOntapStorageVirtualMachineTimeouts)
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

// The jsii proxy struct for FsxOntapStorageVirtualMachineTimeoutsOutputReference
type jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) InternalValue() *FsxOntapStorageVirtualMachineTimeouts {
	var returns *FsxOntapStorageVirtualMachineTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewFsxOntapStorageVirtualMachineTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOntapStorageVirtualMachineTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOntapStorageVirtualMachineTimeoutsOutputReference_Override(f FsxOntapStorageVirtualMachineTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapStorageVirtualMachineTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) SetInternalValue(val *FsxOntapStorageVirtualMachineTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapStorageVirtualMachineTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume aws_fsx_ontap_volume}.
type FsxOntapVolume interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FileSystemId() *string
	FlexcacheEndpointType() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	JunctionPath() *string
	SetJunctionPath(val *string)
	JunctionPathInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	OntapVolumeType() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SecurityStyle() *string
	SetSecurityStyle(val *string)
	SecurityStyleInput() *string
	SizeInMegabytes() *float64
	SetSizeInMegabytes(val *float64)
	SizeInMegabytesInput() *float64
	StorageEfficiencyEnabled() interface{}
	SetStorageEfficiencyEnabled(val interface{})
	StorageEfficiencyEnabledInput() interface{}
	StorageVirtualMachineId() *string
	SetStorageVirtualMachineId(val *string)
	StorageVirtualMachineIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TieringPolicy() FsxOntapVolumeTieringPolicyOutputReference
	TieringPolicyInput() *FsxOntapVolumeTieringPolicy
	Timeouts() FsxOntapVolumeTimeoutsOutputReference
	TimeoutsInput() *FsxOntapVolumeTimeouts
	Uuid() *string
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
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
	PutTieringPolicy(value *FsxOntapVolumeTieringPolicy)
	PutTimeouts(value *FsxOntapVolumeTimeouts)
	ResetOverrideLogicalId()
	ResetSecurityStyle()
	ResetTags()
	ResetTagsAll()
	ResetTieringPolicy()
	ResetTimeouts()
	ResetVolumeType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxOntapVolume
type jsiiProxy_FsxOntapVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxOntapVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) FlexcacheEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flexcacheEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) JunctionPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"junctionPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) JunctionPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"junctionPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) OntapVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ontapVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SizeInMegabytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInMegabytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SizeInMegabytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInMegabytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) StorageEfficiencyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEfficiencyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) StorageEfficiencyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEfficiencyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) StorageVirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageVirtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) StorageVirtualMachineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageVirtualMachineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TieringPolicy() FsxOntapVolumeTieringPolicyOutputReference {
	var returns FsxOntapVolumeTieringPolicyOutputReference
	_jsii_.Get(
		j,
		"tieringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TieringPolicyInput() *FsxOntapVolumeTieringPolicy {
	var returns *FsxOntapVolumeTieringPolicy
	_jsii_.Get(
		j,
		"tieringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Timeouts() FsxOntapVolumeTimeoutsOutputReference {
	var returns FsxOntapVolumeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TimeoutsInput() *FsxOntapVolumeTimeouts {
	var returns *FsxOntapVolumeTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume aws_fsx_ontap_volume} Resource.
func NewFsxOntapVolume(scope constructs.Construct, id *string, config *FsxOntapVolumeConfig) FsxOntapVolume {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapVolume{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume aws_fsx_ontap_volume} Resource.
func NewFsxOntapVolume_Override(f FsxOntapVolume, scope constructs.Construct, id *string, config *FsxOntapVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapVolume",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetJunctionPath(val *string) {
	_jsii_.Set(
		j,
		"junctionPath",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetSecurityStyle(val *string) {
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetSizeInMegabytes(val *float64) {
	_jsii_.Set(
		j,
		"sizeInMegabytes",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetStorageEfficiencyEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"storageEfficiencyEnabled",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetStorageVirtualMachineId(val *string) {
	_jsii_.Set(
		j,
		"storageVirtualMachineId",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume) SetVolumeType(val *string) {
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FsxOntapVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxOntapVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxOntapVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxOntapVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxOntapVolume) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxOntapVolume) PutTieringPolicy(value *FsxOntapVolumeTieringPolicy) {
	_jsii_.InvokeVoid(
		f,
		"putTieringPolicy",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolume) PutTimeouts(value *FsxOntapVolumeTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxOntapVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetSecurityStyle() {
	_jsii_.InvokeVoid(
		f,
		"resetSecurityStyle",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetTieringPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetTieringPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetVolumeType() {
	_jsii_.InvokeVoid(
		f,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxOntapVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxOntapVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS File System FSx.
type FsxOntapVolumeConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#junction_path FsxOntapVolume#junction_path}.
	JunctionPath *string `json:"junctionPath" yaml:"junctionPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#name FsxOntapVolume#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#size_in_megabytes FsxOntapVolume#size_in_megabytes}.
	SizeInMegabytes *float64 `json:"sizeInMegabytes" yaml:"sizeInMegabytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#storage_efficiency_enabled FsxOntapVolume#storage_efficiency_enabled}.
	StorageEfficiencyEnabled interface{} `json:"storageEfficiencyEnabled" yaml:"storageEfficiencyEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#storage_virtual_machine_id FsxOntapVolume#storage_virtual_machine_id}.
	StorageVirtualMachineId *string `json:"storageVirtualMachineId" yaml:"storageVirtualMachineId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#security_style FsxOntapVolume#security_style}.
	SecurityStyle *string `json:"securityStyle" yaml:"securityStyle"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#tags FsxOntapVolume#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#tags_all FsxOntapVolume#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// tiering_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#tiering_policy FsxOntapVolume#tiering_policy}
	TieringPolicy *FsxOntapVolumeTieringPolicy `json:"tieringPolicy" yaml:"tieringPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#timeouts FsxOntapVolume#timeouts}
	Timeouts *FsxOntapVolumeTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#volume_type FsxOntapVolume#volume_type}.
	VolumeType *string `json:"volumeType" yaml:"volumeType"`
}

type FsxOntapVolumeTieringPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#cooling_period FsxOntapVolume#cooling_period}.
	CoolingPeriod *float64 `json:"coolingPeriod" yaml:"coolingPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#name FsxOntapVolume#name}.
	Name *string `json:"name" yaml:"name"`
}

type FsxOntapVolumeTieringPolicyOutputReference interface {
	cdktf.ComplexObject
	CoolingPeriod() *float64
	SetCoolingPeriod(val *float64)
	CoolingPeriodInput() *float64
	InternalValue() *FsxOntapVolumeTieringPolicy
	SetInternalValue(val *FsxOntapVolumeTieringPolicy)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	ResetCoolingPeriod()
	ResetName()
}

// The jsii proxy struct for FsxOntapVolumeTieringPolicyOutputReference
type jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) CoolingPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coolingPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) CoolingPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coolingPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) InternalValue() *FsxOntapVolumeTieringPolicy {
	var returns *FsxOntapVolumeTieringPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxOntapVolumeTieringPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOntapVolumeTieringPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapVolumeTieringPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOntapVolumeTieringPolicyOutputReference_Override(f FsxOntapVolumeTieringPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapVolumeTieringPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) SetCoolingPeriod(val *float64) {
	_jsii_.Set(
		j,
		"coolingPeriod",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) SetInternalValue(val *FsxOntapVolumeTieringPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) ResetCoolingPeriod() {
	_jsii_.InvokeVoid(
		f,
		"resetCoolingPeriod",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeTieringPolicyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		f,
		"resetName",
		nil, // no parameters
	)
}

type FsxOntapVolumeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#create FsxOntapVolume#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#delete FsxOntapVolume#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_volume#update FsxOntapVolume#update}.
	Update *string `json:"update" yaml:"update"`
}

type FsxOntapVolumeTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxOntapVolumeTimeouts
	SetInternalValue(val *FsxOntapVolumeTimeouts)
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

// The jsii proxy struct for FsxOntapVolumeTimeoutsOutputReference
type jsiiProxy_FsxOntapVolumeTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) InternalValue() *FsxOntapVolumeTimeouts {
	var returns *FsxOntapVolumeTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewFsxOntapVolumeTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOntapVolumeTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapVolumeTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapVolumeTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOntapVolumeTimeoutsOutputReference_Override(f FsxOntapVolumeTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOntapVolumeTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) SetInternalValue(val *FsxOntapVolumeTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolumeTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system aws_fsx_openzfs_file_system}.
type FsxOpenzfsFileSystem interface {
	cdktf.TerraformResource
	Arn() *string
	AutomaticBackupRetentionDays() *float64
	SetAutomaticBackupRetentionDays(val *float64)
	AutomaticBackupRetentionDaysInput() *float64
	BackupId() *string
	SetBackupId(val *string)
	BackupIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToBackups() interface{}
	SetCopyTagsToBackups(val interface{})
	CopyTagsToBackupsInput() interface{}
	CopyTagsToVolumes() interface{}
	SetCopyTagsToVolumes(val interface{})
	CopyTagsToVolumesInput() interface{}
	Count() *float64
	SetCount(val *float64)
	DailyAutomaticBackupStartTime() *string
	SetDailyAutomaticBackupStartTime(val *string)
	DailyAutomaticBackupStartTimeInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
	DiskIopsConfiguration() FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference
	DiskIopsConfigurationInput() *FsxOpenzfsFileSystemDiskIopsConfiguration
	DnsName() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkInterfaceIds() *[]*string
	Node() constructs.Node
	OwnerId() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RootVolumeConfiguration() FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference
	RootVolumeConfigurationInput() *FsxOpenzfsFileSystemRootVolumeConfiguration
	RootVolumeId() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	StorageCapacity() *float64
	SetStorageCapacity(val *float64)
	StorageCapacityInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
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
	ThroughputCapacity() *float64
	SetThroughputCapacity(val *float64)
	ThroughputCapacityInput() *float64
	Timeouts() FsxOpenzfsFileSystemTimeoutsOutputReference
	TimeoutsInput() *FsxOpenzfsFileSystemTimeouts
	VpcId() *string
	WeeklyMaintenanceStartTime() *string
	SetWeeklyMaintenanceStartTime(val *string)
	WeeklyMaintenanceStartTimeInput() *string
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
	PutDiskIopsConfiguration(value *FsxOpenzfsFileSystemDiskIopsConfiguration)
	PutRootVolumeConfiguration(value *FsxOpenzfsFileSystemRootVolumeConfiguration)
	PutTimeouts(value *FsxOpenzfsFileSystemTimeouts)
	ResetAutomaticBackupRetentionDays()
	ResetBackupId()
	ResetCopyTagsToBackups()
	ResetCopyTagsToVolumes()
	ResetDailyAutomaticBackupStartTime()
	ResetDiskIopsConfiguration()
	ResetKmsKeyId()
	ResetOverrideLogicalId()
	ResetRootVolumeConfiguration()
	ResetSecurityGroupIds()
	ResetStorageCapacity()
	ResetStorageType()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetWeeklyMaintenanceStartTime()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxOpenzfsFileSystem
type jsiiProxy_FsxOpenzfsFileSystem struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) AutomaticBackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) AutomaticBackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) BackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) CopyTagsToBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) CopyTagsToBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) CopyTagsToVolumes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) CopyTagsToVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) DailyAutomaticBackupStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) DailyAutomaticBackupStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) DiskIopsConfiguration() FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference {
	var returns FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference
	_jsii_.Get(
		j,
		"diskIopsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) DiskIopsConfigurationInput() *FsxOpenzfsFileSystemDiskIopsConfiguration {
	var returns *FsxOpenzfsFileSystemDiskIopsConfiguration
	_jsii_.Get(
		j,
		"diskIopsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) RootVolumeConfiguration() FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference {
	var returns FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"rootVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) RootVolumeConfigurationInput() *FsxOpenzfsFileSystemRootVolumeConfiguration {
	var returns *FsxOpenzfsFileSystemRootVolumeConfiguration
	_jsii_.Get(
		j,
		"rootVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) RootVolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootVolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) StorageCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) ThroughputCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) ThroughputCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) Timeouts() FsxOpenzfsFileSystemTimeoutsOutputReference {
	var returns FsxOpenzfsFileSystemTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) TimeoutsInput() *FsxOpenzfsFileSystemTimeouts {
	var returns *FsxOpenzfsFileSystemTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) WeeklyMaintenanceStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTimeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system aws_fsx_openzfs_file_system} Resource.
func NewFsxOpenzfsFileSystem(scope constructs.Construct, id *string, config *FsxOpenzfsFileSystemConfig) FsxOpenzfsFileSystem {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsFileSystem{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system aws_fsx_openzfs_file_system} Resource.
func NewFsxOpenzfsFileSystem_Override(f FsxOpenzfsFileSystem, scope constructs.Construct, id *string, config *FsxOpenzfsFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystem",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetAutomaticBackupRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"automaticBackupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetBackupId(val *string) {
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetCopyTagsToBackups(val interface{}) {
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetCopyTagsToVolumes(val interface{}) {
	_jsii_.Set(
		j,
		"copyTagsToVolumes",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetDailyAutomaticBackupStartTime(val *string) {
	_jsii_.Set(
		j,
		"dailyAutomaticBackupStartTime",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetDeploymentType(val *string) {
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetStorageCapacity(val *float64) {
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetThroughputCapacity(val *float64) {
	_jsii_.Set(
		j,
		"throughputCapacity",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystem) SetWeeklyMaintenanceStartTime(val *string) {
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FsxOpenzfsFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxOpenzfsFileSystem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystem",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) PutDiskIopsConfiguration(value *FsxOpenzfsFileSystemDiskIopsConfiguration) {
	_jsii_.InvokeVoid(
		f,
		"putDiskIopsConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) PutRootVolumeConfiguration(value *FsxOpenzfsFileSystemRootVolumeConfiguration) {
	_jsii_.InvokeVoid(
		f,
		"putRootVolumeConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) PutTimeouts(value *FsxOpenzfsFileSystemTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetAutomaticBackupRetentionDays() {
	_jsii_.InvokeVoid(
		f,
		"resetAutomaticBackupRetentionDays",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetBackupId() {
	_jsii_.InvokeVoid(
		f,
		"resetBackupId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetCopyTagsToVolumes() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToVolumes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetDailyAutomaticBackupStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetDailyAutomaticBackupStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetDiskIopsConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetDiskIopsConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		f,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetRootVolumeConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetRootVolumeConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		f,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetStorageCapacity() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetStorageType() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxOpenzfsFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS File System FSx.
type FsxOpenzfsFileSystemConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#deployment_type FsxOpenzfsFileSystem#deployment_type}.
	DeploymentType *string `json:"deploymentType" yaml:"deploymentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#subnet_ids FsxOpenzfsFileSystem#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#throughput_capacity FsxOpenzfsFileSystem#throughput_capacity}.
	ThroughputCapacity *float64 `json:"throughputCapacity" yaml:"throughputCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#automatic_backup_retention_days FsxOpenzfsFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#backup_id FsxOpenzfsFileSystem#backup_id}.
	BackupId *string `json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#copy_tags_to_backups FsxOpenzfsFileSystem#copy_tags_to_backups}.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#copy_tags_to_volumes FsxOpenzfsFileSystem#copy_tags_to_volumes}.
	CopyTagsToVolumes interface{} `json:"copyTagsToVolumes" yaml:"copyTagsToVolumes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#daily_automatic_backup_start_time FsxOpenzfsFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// disk_iops_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#disk_iops_configuration FsxOpenzfsFileSystem#disk_iops_configuration}
	DiskIopsConfiguration *FsxOpenzfsFileSystemDiskIopsConfiguration `json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#kms_key_id FsxOpenzfsFileSystem#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// root_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#root_volume_configuration FsxOpenzfsFileSystem#root_volume_configuration}
	RootVolumeConfiguration *FsxOpenzfsFileSystemRootVolumeConfiguration `json:"rootVolumeConfiguration" yaml:"rootVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#security_group_ids FsxOpenzfsFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#storage_capacity FsxOpenzfsFileSystem#storage_capacity}.
	StorageCapacity *float64 `json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#storage_type FsxOpenzfsFileSystem#storage_type}.
	StorageType *string `json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#tags FsxOpenzfsFileSystem#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#tags_all FsxOpenzfsFileSystem#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#timeouts FsxOpenzfsFileSystem#timeouts}
	Timeouts *FsxOpenzfsFileSystemTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#weekly_maintenance_start_time FsxOpenzfsFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

type FsxOpenzfsFileSystemDiskIopsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#iops FsxOpenzfsFileSystem#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#mode FsxOpenzfsFileSystem#mode}.
	Mode *string `json:"mode" yaml:"mode"`
}

type FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *FsxOpenzfsFileSystemDiskIopsConfiguration
	SetInternalValue(val *FsxOpenzfsFileSystemDiskIopsConfiguration)
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
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
	ResetIops()
	ResetMode()
}

// The jsii proxy struct for FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference
type jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) InternalValue() *FsxOpenzfsFileSystemDiskIopsConfiguration {
	var returns *FsxOpenzfsFileSystemDiskIopsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxOpenzfsFileSystemDiskIopsConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOpenzfsFileSystemDiskIopsConfigurationOutputReference_Override(f FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) SetInternalValue(val *FsxOpenzfsFileSystemDiskIopsConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) SetIops(val *float64) {
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		f,
		"resetIops",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemDiskIopsConfigurationOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		f,
		"resetMode",
		nil, // no parameters
	)
}

type FsxOpenzfsFileSystemRootVolumeConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#copy_tags_to_snapshots FsxOpenzfsFileSystem#copy_tags_to_snapshots}.
	CopyTagsToSnapshots interface{} `json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#data_compression_type FsxOpenzfsFileSystem#data_compression_type}.
	DataCompressionType *string `json:"dataCompressionType" yaml:"dataCompressionType"`
	// nfs_exports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#nfs_exports FsxOpenzfsFileSystem#nfs_exports}
	NfsExports *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports `json:"nfsExports" yaml:"nfsExports"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#read_only FsxOpenzfsFileSystem#read_only}.
	ReadOnly interface{} `json:"readOnly" yaml:"readOnly"`
	// user_and_group_quotas block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#user_and_group_quotas FsxOpenzfsFileSystem#user_and_group_quotas}
	UserAndGroupQuotas interface{} `json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
}

type FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports struct {
	// client_configurations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#client_configurations FsxOpenzfsFileSystem#client_configurations}
	ClientConfigurations interface{} `json:"clientConfigurations" yaml:"clientConfigurations"`
}

type FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsClientConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#clients FsxOpenzfsFileSystem#clients}.
	Clients *string `json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#options FsxOpenzfsFileSystem#options}.
	Options *[]*string `json:"options" yaml:"options"`
}

type FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference interface {
	cdktf.ComplexObject
	ClientConfigurations() interface{}
	SetClientConfigurations(val interface{})
	ClientConfigurationsInput() interface{}
	InternalValue() *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports
	SetInternalValue(val *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports)
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
}

// The jsii proxy struct for FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference
type jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) ClientConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) ClientConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) InternalValue() *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports {
	var returns *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference_Override(f FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) SetClientConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"clientConfigurations",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) SetInternalValue(val *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference interface {
	cdktf.ComplexObject
	CopyTagsToSnapshots() interface{}
	SetCopyTagsToSnapshots(val interface{})
	CopyTagsToSnapshotsInput() interface{}
	DataCompressionType() *string
	SetDataCompressionType(val *string)
	DataCompressionTypeInput() *string
	InternalValue() *FsxOpenzfsFileSystemRootVolumeConfiguration
	SetInternalValue(val *FsxOpenzfsFileSystemRootVolumeConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	NfsExports() FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference
	NfsExportsInput() *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAndGroupQuotas() interface{}
	SetUserAndGroupQuotas(val interface{})
	UserAndGroupQuotasInput() interface{}
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
	PutNfsExports(value *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports)
	ResetCopyTagsToSnapshots()
	ResetDataCompressionType()
	ResetNfsExports()
	ResetReadOnly()
	ResetUserAndGroupQuotas()
}

// The jsii proxy struct for FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference
type jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) CopyTagsToSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) CopyTagsToSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) InternalValue() *FsxOpenzfsFileSystemRootVolumeConfiguration {
	var returns *FsxOpenzfsFileSystemRootVolumeConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) NfsExports() FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference {
	var returns FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsOutputReference
	_jsii_.Get(
		j,
		"nfsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) NfsExportsInput() *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports {
	var returns *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports
	_jsii_.Get(
		j,
		"nfsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) UserAndGroupQuotas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) UserAndGroupQuotasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotasInput",
		&returns,
	)
	return returns
}

func NewFsxOpenzfsFileSystemRootVolumeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOpenzfsFileSystemRootVolumeConfigurationOutputReference_Override(f FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) SetCopyTagsToSnapshots(val interface{}) {
	_jsii_.Set(
		j,
		"copyTagsToSnapshots",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) SetDataCompressionType(val *string) {
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) SetInternalValue(val *FsxOpenzfsFileSystemRootVolumeConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) SetReadOnly(val interface{}) {
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) SetUserAndGroupQuotas(val interface{}) {
	_jsii_.Set(
		j,
		"userAndGroupQuotas",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) PutNfsExports(value *FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports) {
	_jsii_.InvokeVoid(
		f,
		"putNfsExports",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetCopyTagsToSnapshots() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToSnapshots",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		f,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetNfsExports() {
	_jsii_.InvokeVoid(
		f,
		"resetNfsExports",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		f,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemRootVolumeConfigurationOutputReference) ResetUserAndGroupQuotas() {
	_jsii_.InvokeVoid(
		f,
		"resetUserAndGroupQuotas",
		nil, // no parameters
	)
}

type FsxOpenzfsFileSystemRootVolumeConfigurationUserAndGroupQuotas struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#id FsxOpenzfsFileSystem#id}.
	Id *float64 `json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#storage_capacity_quota_gib FsxOpenzfsFileSystem#storage_capacity_quota_gib}.
	StorageCapacityQuotaGib *float64 `json:"storageCapacityQuotaGib" yaml:"storageCapacityQuotaGib"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#type FsxOpenzfsFileSystem#type}.
	Type *string `json:"type" yaml:"type"`
}

type FsxOpenzfsFileSystemTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#create FsxOpenzfsFileSystem#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#delete FsxOpenzfsFileSystem#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_file_system#update FsxOpenzfsFileSystem#update}.
	Update *string `json:"update" yaml:"update"`
}

type FsxOpenzfsFileSystemTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxOpenzfsFileSystemTimeouts
	SetInternalValue(val *FsxOpenzfsFileSystemTimeouts)
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

// The jsii proxy struct for FsxOpenzfsFileSystemTimeoutsOutputReference
type jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) InternalValue() *FsxOpenzfsFileSystemTimeouts {
	var returns *FsxOpenzfsFileSystemTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewFsxOpenzfsFileSystemTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOpenzfsFileSystemTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOpenzfsFileSystemTimeoutsOutputReference_Override(f FsxOpenzfsFileSystemTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) SetInternalValue(val *FsxOpenzfsFileSystemTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsFileSystemTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot aws_fsx_openzfs_snapshot}.
type FsxOpenzfsSnapshot interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreationTime() *string
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
	Timeouts() FsxOpenzfsSnapshotTimeoutsOutputReference
	TimeoutsInput() *FsxOpenzfsSnapshotTimeouts
	VolumeId() *string
	SetVolumeId(val *string)
	VolumeIdInput() *string
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
	PutTimeouts(value *FsxOpenzfsSnapshotTimeouts)
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxOpenzfsSnapshot
type jsiiProxy_FsxOpenzfsSnapshot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) Timeouts() FsxOpenzfsSnapshotTimeoutsOutputReference {
	var returns FsxOpenzfsSnapshotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) TimeoutsInput() *FsxOpenzfsSnapshotTimeouts {
	var returns *FsxOpenzfsSnapshotTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) VolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot aws_fsx_openzfs_snapshot} Resource.
func NewFsxOpenzfsSnapshot(scope constructs.Construct, id *string, config *FsxOpenzfsSnapshotConfig) FsxOpenzfsSnapshot {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsSnapshot{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsSnapshot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot aws_fsx_openzfs_snapshot} Resource.
func NewFsxOpenzfsSnapshot_Override(f FsxOpenzfsSnapshot, scope constructs.Construct, id *string, config *FsxOpenzfsSnapshotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsSnapshot",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshot) SetVolumeId(val *string) {
	_jsii_.Set(
		j,
		"volumeId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FsxOpenzfsSnapshot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxOpenzfsSnapshot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxOpenzfsSnapshot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxOpenzfsSnapshot",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxOpenzfsSnapshot) PutTimeouts(value *FsxOpenzfsSnapshotTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsSnapshot) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsSnapshot) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsSnapshot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsSnapshot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxOpenzfsSnapshot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS File System FSx.
type FsxOpenzfsSnapshotConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot#name FsxOpenzfsSnapshot#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot#volume_id FsxOpenzfsSnapshot#volume_id}.
	VolumeId *string `json:"volumeId" yaml:"volumeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot#tags FsxOpenzfsSnapshot#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot#tags_all FsxOpenzfsSnapshot#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot#timeouts FsxOpenzfsSnapshot#timeouts}
	Timeouts *FsxOpenzfsSnapshotTimeouts `json:"timeouts" yaml:"timeouts"`
}

type FsxOpenzfsSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot#create FsxOpenzfsSnapshot#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot#delete FsxOpenzfsSnapshot#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot#read FsxOpenzfsSnapshot#read}.
	Read *string `json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_snapshot#update FsxOpenzfsSnapshot#update}.
	Update *string `json:"update" yaml:"update"`
}

type FsxOpenzfsSnapshotTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxOpenzfsSnapshotTimeouts
	SetInternalValue(val *FsxOpenzfsSnapshotTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Read() *string
	SetRead(val *string)
	ReadInput() *string
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
	ResetRead()
	ResetUpdate()
}

// The jsii proxy struct for FsxOpenzfsSnapshotTimeoutsOutputReference
type jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) InternalValue() *FsxOpenzfsSnapshotTimeouts {
	var returns *FsxOpenzfsSnapshotTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewFsxOpenzfsSnapshotTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOpenzfsSnapshotTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsSnapshotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOpenzfsSnapshotTimeoutsOutputReference_Override(f FsxOpenzfsSnapshotTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsSnapshotTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) SetInternalValue(val *FsxOpenzfsSnapshotTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		f,
		"resetRead",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsSnapshotTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume aws_fsx_openzfs_volume}.
type FsxOpenzfsVolume interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToSnapshots() interface{}
	SetCopyTagsToSnapshots(val interface{})
	CopyTagsToSnapshotsInput() interface{}
	Count() *float64
	SetCount(val *float64)
	DataCompressionType() *string
	SetDataCompressionType(val *string)
	DataCompressionTypeInput() *string
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
	NfsExports() FsxOpenzfsVolumeNfsExportsOutputReference
	NfsExportsInput() *FsxOpenzfsVolumeNfsExports
	Node() constructs.Node
	OriginSnapshot() FsxOpenzfsVolumeOriginSnapshotOutputReference
	OriginSnapshotInput() *FsxOpenzfsVolumeOriginSnapshot
	ParentVolumeId() *string
	SetParentVolumeId(val *string)
	ParentVolumeIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	StorageCapacityQuotaGib() *float64
	SetStorageCapacityQuotaGib(val *float64)
	StorageCapacityQuotaGibInput() *float64
	StorageCapacityReservationGib() *float64
	SetStorageCapacityReservationGib(val *float64)
	StorageCapacityReservationGibInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() FsxOpenzfsVolumeTimeoutsOutputReference
	TimeoutsInput() *FsxOpenzfsVolumeTimeouts
	UserAndGroupQuotas() interface{}
	SetUserAndGroupQuotas(val interface{})
	UserAndGroupQuotasInput() interface{}
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
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
	PutNfsExports(value *FsxOpenzfsVolumeNfsExports)
	PutOriginSnapshot(value *FsxOpenzfsVolumeOriginSnapshot)
	PutTimeouts(value *FsxOpenzfsVolumeTimeouts)
	ResetCopyTagsToSnapshots()
	ResetDataCompressionType()
	ResetNfsExports()
	ResetOriginSnapshot()
	ResetOverrideLogicalId()
	ResetReadOnly()
	ResetStorageCapacityQuotaGib()
	ResetStorageCapacityReservationGib()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetUserAndGroupQuotas()
	ResetVolumeType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxOpenzfsVolume
type jsiiProxy_FsxOpenzfsVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxOpenzfsVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) CopyTagsToSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) CopyTagsToSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) NfsExports() FsxOpenzfsVolumeNfsExportsOutputReference {
	var returns FsxOpenzfsVolumeNfsExportsOutputReference
	_jsii_.Get(
		j,
		"nfsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) NfsExportsInput() *FsxOpenzfsVolumeNfsExports {
	var returns *FsxOpenzfsVolumeNfsExports
	_jsii_.Get(
		j,
		"nfsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) OriginSnapshot() FsxOpenzfsVolumeOriginSnapshotOutputReference {
	var returns FsxOpenzfsVolumeOriginSnapshotOutputReference
	_jsii_.Get(
		j,
		"originSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) OriginSnapshotInput() *FsxOpenzfsVolumeOriginSnapshot {
	var returns *FsxOpenzfsVolumeOriginSnapshot
	_jsii_.Get(
		j,
		"originSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ParentVolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentVolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ParentVolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentVolumeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) StorageCapacityQuotaGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityQuotaGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) StorageCapacityQuotaGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityQuotaGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) StorageCapacityReservationGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityReservationGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) StorageCapacityReservationGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityReservationGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Timeouts() FsxOpenzfsVolumeTimeoutsOutputReference {
	var returns FsxOpenzfsVolumeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TimeoutsInput() *FsxOpenzfsVolumeTimeouts {
	var returns *FsxOpenzfsVolumeTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) UserAndGroupQuotas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) UserAndGroupQuotasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume aws_fsx_openzfs_volume} Resource.
func NewFsxOpenzfsVolume(scope constructs.Construct, id *string, config *FsxOpenzfsVolumeConfig) FsxOpenzfsVolume {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsVolume{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume aws_fsx_openzfs_volume} Resource.
func NewFsxOpenzfsVolume_Override(f FsxOpenzfsVolume, scope constructs.Construct, id *string, config *FsxOpenzfsVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsVolume",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetCopyTagsToSnapshots(val interface{}) {
	_jsii_.Set(
		j,
		"copyTagsToSnapshots",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetDataCompressionType(val *string) {
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetParentVolumeId(val *string) {
	_jsii_.Set(
		j,
		"parentVolumeId",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetReadOnly(val interface{}) {
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetStorageCapacityQuotaGib(val *float64) {
	_jsii_.Set(
		j,
		"storageCapacityQuotaGib",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetStorageCapacityReservationGib(val *float64) {
	_jsii_.Set(
		j,
		"storageCapacityReservationGib",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetUserAndGroupQuotas(val interface{}) {
	_jsii_.Set(
		j,
		"userAndGroupQuotas",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume) SetVolumeType(val *string) {
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FsxOpenzfsVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxOpenzfsVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxOpenzfsVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxOpenzfsVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) PutNfsExports(value *FsxOpenzfsVolumeNfsExports) {
	_jsii_.InvokeVoid(
		f,
		"putNfsExports",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) PutOriginSnapshot(value *FsxOpenzfsVolumeOriginSnapshot) {
	_jsii_.InvokeVoid(
		f,
		"putOriginSnapshot",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) PutTimeouts(value *FsxOpenzfsVolumeTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetCopyTagsToSnapshots() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToSnapshots",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		f,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetNfsExports() {
	_jsii_.InvokeVoid(
		f,
		"resetNfsExports",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetOriginSnapshot() {
	_jsii_.InvokeVoid(
		f,
		"resetOriginSnapshot",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetReadOnly() {
	_jsii_.InvokeVoid(
		f,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetStorageCapacityQuotaGib() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacityQuotaGib",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetStorageCapacityReservationGib() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacityReservationGib",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetUserAndGroupQuotas() {
	_jsii_.InvokeVoid(
		f,
		"resetUserAndGroupQuotas",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetVolumeType() {
	_jsii_.InvokeVoid(
		f,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxOpenzfsVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS File System FSx.
type FsxOpenzfsVolumeConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#name FsxOpenzfsVolume#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#parent_volume_id FsxOpenzfsVolume#parent_volume_id}.
	ParentVolumeId *string `json:"parentVolumeId" yaml:"parentVolumeId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#copy_tags_to_snapshots FsxOpenzfsVolume#copy_tags_to_snapshots}.
	CopyTagsToSnapshots interface{} `json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#data_compression_type FsxOpenzfsVolume#data_compression_type}.
	DataCompressionType *string `json:"dataCompressionType" yaml:"dataCompressionType"`
	// nfs_exports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#nfs_exports FsxOpenzfsVolume#nfs_exports}
	NfsExports *FsxOpenzfsVolumeNfsExports `json:"nfsExports" yaml:"nfsExports"`
	// origin_snapshot block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#origin_snapshot FsxOpenzfsVolume#origin_snapshot}
	OriginSnapshot *FsxOpenzfsVolumeOriginSnapshot `json:"originSnapshot" yaml:"originSnapshot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#read_only FsxOpenzfsVolume#read_only}.
	ReadOnly interface{} `json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#storage_capacity_quota_gib FsxOpenzfsVolume#storage_capacity_quota_gib}.
	StorageCapacityQuotaGib *float64 `json:"storageCapacityQuotaGib" yaml:"storageCapacityQuotaGib"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#storage_capacity_reservation_gib FsxOpenzfsVolume#storage_capacity_reservation_gib}.
	StorageCapacityReservationGib *float64 `json:"storageCapacityReservationGib" yaml:"storageCapacityReservationGib"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#tags FsxOpenzfsVolume#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#tags_all FsxOpenzfsVolume#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#timeouts FsxOpenzfsVolume#timeouts}
	Timeouts *FsxOpenzfsVolumeTimeouts `json:"timeouts" yaml:"timeouts"`
	// user_and_group_quotas block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#user_and_group_quotas FsxOpenzfsVolume#user_and_group_quotas}
	UserAndGroupQuotas interface{} `json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#volume_type FsxOpenzfsVolume#volume_type}.
	VolumeType *string `json:"volumeType" yaml:"volumeType"`
}

type FsxOpenzfsVolumeNfsExports struct {
	// client_configurations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#client_configurations FsxOpenzfsVolume#client_configurations}
	ClientConfigurations interface{} `json:"clientConfigurations" yaml:"clientConfigurations"`
}

type FsxOpenzfsVolumeNfsExportsClientConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#clients FsxOpenzfsVolume#clients}.
	Clients *string `json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#options FsxOpenzfsVolume#options}.
	Options *[]*string `json:"options" yaml:"options"`
}

type FsxOpenzfsVolumeNfsExportsOutputReference interface {
	cdktf.ComplexObject
	ClientConfigurations() interface{}
	SetClientConfigurations(val interface{})
	ClientConfigurationsInput() interface{}
	InternalValue() *FsxOpenzfsVolumeNfsExports
	SetInternalValue(val *FsxOpenzfsVolumeNfsExports)
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
}

// The jsii proxy struct for FsxOpenzfsVolumeNfsExportsOutputReference
type jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) ClientConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) ClientConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) InternalValue() *FsxOpenzfsVolumeNfsExports {
	var returns *FsxOpenzfsVolumeNfsExports
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxOpenzfsVolumeNfsExportsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOpenzfsVolumeNfsExportsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsVolumeNfsExportsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOpenzfsVolumeNfsExportsOutputReference_Override(f FsxOpenzfsVolumeNfsExportsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsVolumeNfsExportsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) SetClientConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"clientConfigurations",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) SetInternalValue(val *FsxOpenzfsVolumeNfsExports) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeNfsExportsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOpenzfsVolumeOriginSnapshot struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#copy_strategy FsxOpenzfsVolume#copy_strategy}.
	CopyStrategy *string `json:"copyStrategy" yaml:"copyStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#snapshot_arn FsxOpenzfsVolume#snapshot_arn}.
	SnapshotArn *string `json:"snapshotArn" yaml:"snapshotArn"`
}

type FsxOpenzfsVolumeOriginSnapshotOutputReference interface {
	cdktf.ComplexObject
	CopyStrategy() *string
	SetCopyStrategy(val *string)
	CopyStrategyInput() *string
	InternalValue() *FsxOpenzfsVolumeOriginSnapshot
	SetInternalValue(val *FsxOpenzfsVolumeOriginSnapshot)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SnapshotArn() *string
	SetSnapshotArn(val *string)
	SnapshotArnInput() *string
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
}

// The jsii proxy struct for FsxOpenzfsVolumeOriginSnapshotOutputReference
type jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) CopyStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) CopyStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) InternalValue() *FsxOpenzfsVolumeOriginSnapshot {
	var returns *FsxOpenzfsVolumeOriginSnapshot
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) SnapshotArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) SnapshotArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxOpenzfsVolumeOriginSnapshotOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOpenzfsVolumeOriginSnapshotOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsVolumeOriginSnapshotOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOpenzfsVolumeOriginSnapshotOutputReference_Override(f FsxOpenzfsVolumeOriginSnapshotOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsVolumeOriginSnapshotOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) SetCopyStrategy(val *string) {
	_jsii_.Set(
		j,
		"copyStrategy",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) SetInternalValue(val *FsxOpenzfsVolumeOriginSnapshot) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) SetSnapshotArn(val *string) {
	_jsii_.Set(
		j,
		"snapshotArn",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeOriginSnapshotOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type FsxOpenzfsVolumeTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#create FsxOpenzfsVolume#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#delete FsxOpenzfsVolume#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#update FsxOpenzfsVolume#update}.
	Update *string `json:"update" yaml:"update"`
}

type FsxOpenzfsVolumeTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxOpenzfsVolumeTimeouts
	SetInternalValue(val *FsxOpenzfsVolumeTimeouts)
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

// The jsii proxy struct for FsxOpenzfsVolumeTimeoutsOutputReference
type jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) InternalValue() *FsxOpenzfsVolumeTimeouts {
	var returns *FsxOpenzfsVolumeTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewFsxOpenzfsVolumeTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxOpenzfsVolumeTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsVolumeTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOpenzfsVolumeTimeoutsOutputReference_Override(f FsxOpenzfsVolumeTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxOpenzfsVolumeTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) SetInternalValue(val *FsxOpenzfsVolumeTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolumeTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

type FsxOpenzfsVolumeUserAndGroupQuotas struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#id FsxOpenzfsVolume#id}.
	Id *float64 `json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#storage_capacity_quota_gib FsxOpenzfsVolume#storage_capacity_quota_gib}.
	StorageCapacityQuotaGib *float64 `json:"storageCapacityQuotaGib" yaml:"storageCapacityQuotaGib"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_openzfs_volume#type FsxOpenzfsVolume#type}.
	Type *string `json:"type" yaml:"type"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system aws_fsx_windows_file_system}.
type FsxWindowsFileSystem interface {
	cdktf.TerraformResource
	ActiveDirectoryId() *string
	SetActiveDirectoryId(val *string)
	ActiveDirectoryIdInput() *string
	Aliases() *[]*string
	SetAliases(val *[]*string)
	AliasesInput() *[]*string
	Arn() *string
	AuditLogConfiguration() FsxWindowsFileSystemAuditLogConfigurationOutputReference
	AuditLogConfigurationInput() *FsxWindowsFileSystemAuditLogConfiguration
	AutomaticBackupRetentionDays() *float64
	SetAutomaticBackupRetentionDays(val *float64)
	AutomaticBackupRetentionDaysInput() *float64
	BackupId() *string
	SetBackupId(val *string)
	BackupIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToBackups() interface{}
	SetCopyTagsToBackups(val interface{})
	CopyTagsToBackupsInput() interface{}
	Count() *float64
	SetCount(val *float64)
	DailyAutomaticBackupStartTime() *string
	SetDailyAutomaticBackupStartTime(val *string)
	DailyAutomaticBackupStartTimeInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DeploymentType() *string
	SetDeploymentType(val *string)
	DeploymentTypeInput() *string
	DnsName() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	NetworkInterfaceIds() *[]*string
	Node() constructs.Node
	OwnerId() *string
	PreferredFileServerIp() *string
	PreferredSubnetId() *string
	SetPreferredSubnetId(val *string)
	PreferredSubnetIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RemoteAdministrationEndpoint() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SelfManagedActiveDirectory() FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference
	SelfManagedActiveDirectoryInput() *FsxWindowsFileSystemSelfManagedActiveDirectory
	SkipFinalBackup() interface{}
	SetSkipFinalBackup(val interface{})
	SkipFinalBackupInput() interface{}
	StorageCapacity() *float64
	SetStorageCapacity(val *float64)
	StorageCapacityInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
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
	ThroughputCapacity() *float64
	SetThroughputCapacity(val *float64)
	ThroughputCapacityInput() *float64
	Timeouts() FsxWindowsFileSystemTimeoutsOutputReference
	TimeoutsInput() *FsxWindowsFileSystemTimeouts
	VpcId() *string
	WeeklyMaintenanceStartTime() *string
	SetWeeklyMaintenanceStartTime(val *string)
	WeeklyMaintenanceStartTimeInput() *string
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
	PutAuditLogConfiguration(value *FsxWindowsFileSystemAuditLogConfiguration)
	PutSelfManagedActiveDirectory(value *FsxWindowsFileSystemSelfManagedActiveDirectory)
	PutTimeouts(value *FsxWindowsFileSystemTimeouts)
	ResetActiveDirectoryId()
	ResetAliases()
	ResetAuditLogConfiguration()
	ResetAutomaticBackupRetentionDays()
	ResetBackupId()
	ResetCopyTagsToBackups()
	ResetDailyAutomaticBackupStartTime()
	ResetDeploymentType()
	ResetKmsKeyId()
	ResetOverrideLogicalId()
	ResetPreferredSubnetId()
	ResetSecurityGroupIds()
	ResetSelfManagedActiveDirectory()
	ResetSkipFinalBackup()
	ResetStorageCapacity()
	ResetStorageType()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetWeeklyMaintenanceStartTime()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxWindowsFileSystem
type jsiiProxy_FsxWindowsFileSystem struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxWindowsFileSystem) ActiveDirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) ActiveDirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) AuditLogConfiguration() FsxWindowsFileSystemAuditLogConfigurationOutputReference {
	var returns FsxWindowsFileSystemAuditLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"auditLogConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) AuditLogConfigurationInput() *FsxWindowsFileSystemAuditLogConfiguration {
	var returns *FsxWindowsFileSystemAuditLogConfiguration
	_jsii_.Get(
		j,
		"auditLogConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) AutomaticBackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) AutomaticBackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) BackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) CopyTagsToBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) CopyTagsToBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) DailyAutomaticBackupStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) DailyAutomaticBackupStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) PreferredFileServerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredFileServerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) PreferredSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) PreferredSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) RemoteAdministrationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteAdministrationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) SelfManagedActiveDirectory() FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference {
	var returns FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"selfManagedActiveDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) SelfManagedActiveDirectoryInput() *FsxWindowsFileSystemSelfManagedActiveDirectory {
	var returns *FsxWindowsFileSystemSelfManagedActiveDirectory
	_jsii_.Get(
		j,
		"selfManagedActiveDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) SkipFinalBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) SkipFinalBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) StorageCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) ThroughputCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) ThroughputCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) Timeouts() FsxWindowsFileSystemTimeoutsOutputReference {
	var returns FsxWindowsFileSystemTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TimeoutsInput() *FsxWindowsFileSystemTimeouts {
	var returns *FsxWindowsFileSystemTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) WeeklyMaintenanceStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTimeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system aws_fsx_windows_file_system} Resource.
func NewFsxWindowsFileSystem(scope constructs.Construct, id *string, config *FsxWindowsFileSystemConfig) FsxWindowsFileSystem {
	_init_.Initialize()

	j := jsiiProxy_FsxWindowsFileSystem{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxWindowsFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system aws_fsx_windows_file_system} Resource.
func NewFsxWindowsFileSystem_Override(f FsxWindowsFileSystem, scope constructs.Construct, id *string, config *FsxWindowsFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxWindowsFileSystem",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetActiveDirectoryId(val *string) {
	_jsii_.Set(
		j,
		"activeDirectoryId",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetAliases(val *[]*string) {
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetAutomaticBackupRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"automaticBackupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetBackupId(val *string) {
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetCopyTagsToBackups(val interface{}) {
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetDailyAutomaticBackupStartTime(val *string) {
	_jsii_.Set(
		j,
		"dailyAutomaticBackupStartTime",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetDeploymentType(val *string) {
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetPreferredSubnetId(val *string) {
	_jsii_.Set(
		j,
		"preferredSubnetId",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetSkipFinalBackup(val interface{}) {
	_jsii_.Set(
		j,
		"skipFinalBackup",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetStorageCapacity(val *float64) {
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetStorageType(val *string) {
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetThroughputCapacity(val *float64) {
	_jsii_.Set(
		j,
		"throughputCapacity",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetWeeklyMaintenanceStartTime(val *string) {
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FsxWindowsFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.fsx.FsxWindowsFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxWindowsFileSystem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.fsx.FsxWindowsFileSystem",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) PutAuditLogConfiguration(value *FsxWindowsFileSystemAuditLogConfiguration) {
	_jsii_.InvokeVoid(
		f,
		"putAuditLogConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) PutSelfManagedActiveDirectory(value *FsxWindowsFileSystemSelfManagedActiveDirectory) {
	_jsii_.InvokeVoid(
		f,
		"putSelfManagedActiveDirectory",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) PutTimeouts(value *FsxWindowsFileSystemTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetActiveDirectoryId() {
	_jsii_.InvokeVoid(
		f,
		"resetActiveDirectoryId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetAliases() {
	_jsii_.InvokeVoid(
		f,
		"resetAliases",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetAuditLogConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetAuditLogConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetAutomaticBackupRetentionDays() {
	_jsii_.InvokeVoid(
		f,
		"resetAutomaticBackupRetentionDays",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetBackupId() {
	_jsii_.InvokeVoid(
		f,
		"resetBackupId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetDailyAutomaticBackupStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetDailyAutomaticBackupStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetDeploymentType() {
	_jsii_.InvokeVoid(
		f,
		"resetDeploymentType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		f,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetPreferredSubnetId() {
	_jsii_.InvokeVoid(
		f,
		"resetPreferredSubnetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		f,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetSelfManagedActiveDirectory() {
	_jsii_.InvokeVoid(
		f,
		"resetSelfManagedActiveDirectory",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetSkipFinalBackup() {
	_jsii_.InvokeVoid(
		f,
		"resetSkipFinalBackup",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetStorageCapacity() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacity",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetStorageType() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		f,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FsxWindowsFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FsxWindowsFileSystemAuditLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#audit_log_destination FsxWindowsFileSystem#audit_log_destination}.
	AuditLogDestination *string `json:"auditLogDestination" yaml:"auditLogDestination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#file_access_audit_log_level FsxWindowsFileSystem#file_access_audit_log_level}.
	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel" yaml:"fileAccessAuditLogLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#file_share_access_audit_log_level FsxWindowsFileSystem#file_share_access_audit_log_level}.
	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel" yaml:"fileShareAccessAuditLogLevel"`
}

type FsxWindowsFileSystemAuditLogConfigurationOutputReference interface {
	cdktf.ComplexObject
	AuditLogDestination() *string
	SetAuditLogDestination(val *string)
	AuditLogDestinationInput() *string
	FileAccessAuditLogLevel() *string
	SetFileAccessAuditLogLevel(val *string)
	FileAccessAuditLogLevelInput() *string
	FileShareAccessAuditLogLevel() *string
	SetFileShareAccessAuditLogLevel(val *string)
	FileShareAccessAuditLogLevelInput() *string
	InternalValue() *FsxWindowsFileSystemAuditLogConfiguration
	SetInternalValue(val *FsxWindowsFileSystemAuditLogConfiguration)
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
	ResetAuditLogDestination()
	ResetFileAccessAuditLogLevel()
	ResetFileShareAccessAuditLogLevel()
}

// The jsii proxy struct for FsxWindowsFileSystemAuditLogConfigurationOutputReference
type jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) AuditLogDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) AuditLogDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) FileAccessAuditLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAccessAuditLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) FileAccessAuditLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAccessAuditLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) FileShareAccessAuditLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareAccessAuditLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) FileShareAccessAuditLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareAccessAuditLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) InternalValue() *FsxWindowsFileSystemAuditLogConfiguration {
	var returns *FsxWindowsFileSystemAuditLogConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxWindowsFileSystemAuditLogConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxWindowsFileSystemAuditLogConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxWindowsFileSystemAuditLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxWindowsFileSystemAuditLogConfigurationOutputReference_Override(f FsxWindowsFileSystemAuditLogConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxWindowsFileSystemAuditLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) SetAuditLogDestination(val *string) {
	_jsii_.Set(
		j,
		"auditLogDestination",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) SetFileAccessAuditLogLevel(val *string) {
	_jsii_.Set(
		j,
		"fileAccessAuditLogLevel",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) SetFileShareAccessAuditLogLevel(val *string) {
	_jsii_.Set(
		j,
		"fileShareAccessAuditLogLevel",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) SetInternalValue(val *FsxWindowsFileSystemAuditLogConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ResetAuditLogDestination() {
	_jsii_.InvokeVoid(
		f,
		"resetAuditLogDestination",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ResetFileAccessAuditLogLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetFileAccessAuditLogLevel",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) ResetFileShareAccessAuditLogLevel() {
	_jsii_.InvokeVoid(
		f,
		"resetFileShareAccessAuditLogLevel",
		nil, // no parameters
	)
}

// AWS File System FSx.
type FsxWindowsFileSystemConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#subnet_ids FsxWindowsFileSystem#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#throughput_capacity FsxWindowsFileSystem#throughput_capacity}.
	ThroughputCapacity *float64 `json:"throughputCapacity" yaml:"throughputCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#active_directory_id FsxWindowsFileSystem#active_directory_id}.
	ActiveDirectoryId *string `json:"activeDirectoryId" yaml:"activeDirectoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#aliases FsxWindowsFileSystem#aliases}.
	Aliases *[]*string `json:"aliases" yaml:"aliases"`
	// audit_log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#audit_log_configuration FsxWindowsFileSystem#audit_log_configuration}
	AuditLogConfiguration *FsxWindowsFileSystemAuditLogConfiguration `json:"auditLogConfiguration" yaml:"auditLogConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#automatic_backup_retention_days FsxWindowsFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#backup_id FsxWindowsFileSystem#backup_id}.
	BackupId *string `json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#copy_tags_to_backups FsxWindowsFileSystem#copy_tags_to_backups}.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#daily_automatic_backup_start_time FsxWindowsFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#deployment_type FsxWindowsFileSystem#deployment_type}.
	DeploymentType *string `json:"deploymentType" yaml:"deploymentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#kms_key_id FsxWindowsFileSystem#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#preferred_subnet_id FsxWindowsFileSystem#preferred_subnet_id}.
	PreferredSubnetId *string `json:"preferredSubnetId" yaml:"preferredSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#security_group_ids FsxWindowsFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// self_managed_active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#self_managed_active_directory FsxWindowsFileSystem#self_managed_active_directory}
	SelfManagedActiveDirectory *FsxWindowsFileSystemSelfManagedActiveDirectory `json:"selfManagedActiveDirectory" yaml:"selfManagedActiveDirectory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#skip_final_backup FsxWindowsFileSystem#skip_final_backup}.
	SkipFinalBackup interface{} `json:"skipFinalBackup" yaml:"skipFinalBackup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#storage_capacity FsxWindowsFileSystem#storage_capacity}.
	StorageCapacity *float64 `json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#storage_type FsxWindowsFileSystem#storage_type}.
	StorageType *string `json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#tags FsxWindowsFileSystem#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#tags_all FsxWindowsFileSystem#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#timeouts FsxWindowsFileSystem#timeouts}
	Timeouts *FsxWindowsFileSystemTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#weekly_maintenance_start_time FsxWindowsFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

type FsxWindowsFileSystemSelfManagedActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#dns_ips FsxWindowsFileSystem#dns_ips}.
	DnsIps *[]*string `json:"dnsIps" yaml:"dnsIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#domain_name FsxWindowsFileSystem#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#password FsxWindowsFileSystem#password}.
	Password *string `json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#username FsxWindowsFileSystem#username}.
	Username *string `json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#file_system_administrators_group FsxWindowsFileSystem#file_system_administrators_group}.
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#organizational_unit_distinguished_name FsxWindowsFileSystem#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

type FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference interface {
	cdktf.ComplexObject
	DnsIps() *[]*string
	SetDnsIps(val *[]*string)
	DnsIpsInput() *[]*string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	FileSystemAdministratorsGroup() *string
	SetFileSystemAdministratorsGroup(val *string)
	FileSystemAdministratorsGroupInput() *string
	InternalValue() *FsxWindowsFileSystemSelfManagedActiveDirectory
	SetInternalValue(val *FsxWindowsFileSystemSelfManagedActiveDirectory)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OrganizationalUnitDistinguishedName() *string
	SetOrganizationalUnitDistinguishedName(val *string)
	OrganizationalUnitDistinguishedNameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetFileSystemAdministratorsGroup()
	ResetOrganizationalUnitDistinguishedName()
}

// The jsii proxy struct for FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference
type jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) DnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) DnsIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) FileSystemAdministratorsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemAdministratorsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) FileSystemAdministratorsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemAdministratorsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) InternalValue() *FsxWindowsFileSystemSelfManagedActiveDirectory {
	var returns *FsxWindowsFileSystemSelfManagedActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) OrganizationalUnitDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) OrganizationalUnitDistinguishedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func NewFsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference_Override(f FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetDnsIps(val *[]*string) {
	_jsii_.Set(
		j,
		"dnsIps",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetFileSystemAdministratorsGroup(val *string) {
	_jsii_.Set(
		j,
		"fileSystemAdministratorsGroup",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetInternalValue(val *FsxWindowsFileSystemSelfManagedActiveDirectory) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetOrganizationalUnitDistinguishedName(val *string) {
	_jsii_.Set(
		j,
		"organizationalUnitDistinguishedName",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) ResetFileSystemAdministratorsGroup() {
	_jsii_.InvokeVoid(
		f,
		"resetFileSystemAdministratorsGroup",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) ResetOrganizationalUnitDistinguishedName() {
	_jsii_.InvokeVoid(
		f,
		"resetOrganizationalUnitDistinguishedName",
		nil, // no parameters
	)
}

type FsxWindowsFileSystemTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#create FsxWindowsFileSystem#create}.
	Create *string `json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#delete FsxWindowsFileSystem#delete}.
	Delete *string `json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system#update FsxWindowsFileSystem#update}.
	Update *string `json:"update" yaml:"update"`
}

type FsxWindowsFileSystemTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	InternalValue() *FsxWindowsFileSystemTimeouts
	SetInternalValue(val *FsxWindowsFileSystemTimeouts)
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

// The jsii proxy struct for FsxWindowsFileSystemTimeoutsOutputReference
type jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) InternalValue() *FsxWindowsFileSystemTimeouts {
	var returns *FsxWindowsFileSystemTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewFsxWindowsFileSystemTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) FsxWindowsFileSystemTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxWindowsFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxWindowsFileSystemTimeoutsOutputReference_Override(f FsxWindowsFileSystemTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.fsx.FsxWindowsFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) SetInternalValue(val *FsxWindowsFileSystemTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}
