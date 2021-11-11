package fsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/fsx/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup.html aws_fsx_backup}.
type FsxBackup interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() FsxBackupTimeoutsOutputReference
	TimeoutsInput() *FsxBackupTimeouts
	Type() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *FsxBackupTimeouts)
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
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

func (j *jsiiProxy_FsxBackup) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_FsxBackup) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxBackup) TagsInput() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup.html aws_fsx_backup} Resource.
func NewFsxBackup(scope constructs.Construct, id *string, config *FsxBackupConfig) FsxBackup {
	_init_.Initialize()

	j := jsiiProxy_FsxBackup{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxBackup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup.html aws_fsx_backup} Resource.
func NewFsxBackup_Override(f FsxBackup, scope constructs.Construct, id *string, config *FsxBackupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxBackup",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxBackup) SetCount(val interface{}) {
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

func (j *jsiiProxy_FsxBackup) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxBackup) SetTagsAll(val interface{}) {
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
func FsxBackup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.FSx.FsxBackup",
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
		"hashicorp_aws.FSx.FsxBackup",
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

type FsxBackupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup.html#file_system_id FsxBackup#file_system_id}.
	FileSystemId *string `json:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup.html#tags FsxBackup#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup.html#tags_all FsxBackup#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup.html#timeouts FsxBackup#timeouts}
	Timeouts *FsxBackupTimeouts `json:"timeouts"`
}

type FsxBackupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup.html#create FsxBackup#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_backup.html#delete FsxBackup#delete}.
	Delete *string `json:"delete"`
}

type FsxBackupTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
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

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxBackupTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FsxBackupTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxBackupTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxBackupTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxBackupTimeoutsOutputReference_Override(f FsxBackupTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxBackupTimeoutsOutputReference",
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

func (j *jsiiProxy_FsxBackupTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxBackupTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html aws_fsx_lustre_file_system}.
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
	Count() interface{}
	SetCount(val interface{})
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
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
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
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_FsxLustreFileSystem) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_FsxLustreFileSystem) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxLustreFileSystem) TagsInput() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html aws_fsx_lustre_file_system} Resource.
func NewFsxLustreFileSystem(scope constructs.Construct, id *string, config *FsxLustreFileSystemConfig) FsxLustreFileSystem {
	_init_.Initialize()

	j := jsiiProxy_FsxLustreFileSystem{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxLustreFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html aws_fsx_lustre_file_system} Resource.
func NewFsxLustreFileSystem_Override(f FsxLustreFileSystem, scope constructs.Construct, id *string, config *FsxLustreFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxLustreFileSystem",
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

func (j *jsiiProxy_FsxLustreFileSystem) SetCount(val interface{}) {
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

func (j *jsiiProxy_FsxLustreFileSystem) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxLustreFileSystem) SetTagsAll(val interface{}) {
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
		"hashicorp_aws.FSx.FsxLustreFileSystem",
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
		"hashicorp_aws.FSx.FsxLustreFileSystem",
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

type FsxLustreFileSystemConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#subnet_ids FsxLustreFileSystem#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#auto_import_policy FsxLustreFileSystem#auto_import_policy}.
	AutoImportPolicy *string `json:"autoImportPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#automatic_backup_retention_days FsxLustreFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#backup_id FsxLustreFileSystem#backup_id}.
	BackupId *string `json:"backupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#copy_tags_to_backups FsxLustreFileSystem#copy_tags_to_backups}.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#daily_automatic_backup_start_time FsxLustreFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#data_compression_type FsxLustreFileSystem#data_compression_type}.
	DataCompressionType *string `json:"dataCompressionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#deployment_type FsxLustreFileSystem#deployment_type}.
	DeploymentType *string `json:"deploymentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#drive_cache_type FsxLustreFileSystem#drive_cache_type}.
	DriveCacheType *string `json:"driveCacheType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#export_path FsxLustreFileSystem#export_path}.
	ExportPath *string `json:"exportPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#imported_file_chunk_size FsxLustreFileSystem#imported_file_chunk_size}.
	ImportedFileChunkSize *float64 `json:"importedFileChunkSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#import_path FsxLustreFileSystem#import_path}.
	ImportPath *string `json:"importPath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#kms_key_id FsxLustreFileSystem#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#per_unit_storage_throughput FsxLustreFileSystem#per_unit_storage_throughput}.
	PerUnitStorageThroughput *float64 `json:"perUnitStorageThroughput"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#security_group_ids FsxLustreFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#storage_capacity FsxLustreFileSystem#storage_capacity}.
	StorageCapacity *float64 `json:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#storage_type FsxLustreFileSystem#storage_type}.
	StorageType *string `json:"storageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#tags FsxLustreFileSystem#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#tags_all FsxLustreFileSystem#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#timeouts FsxLustreFileSystem#timeouts}
	Timeouts *FsxLustreFileSystemTimeouts `json:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#weekly_maintenance_start_time FsxLustreFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

type FsxLustreFileSystemTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#create FsxLustreFileSystem#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#delete FsxLustreFileSystem#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_lustre_file_system.html#update FsxLustreFileSystem#update}.
	Update *string `json:"update"`
}

type FsxLustreFileSystemTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func NewFsxLustreFileSystemTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FsxLustreFileSystemTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxLustreFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxLustreFileSystemTimeoutsOutputReference_Override(f FsxLustreFileSystemTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxLustreFileSystemTimeoutsOutputReference",
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

func (j *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (f *jsiiProxy_FsxLustreFileSystemTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html aws_fsx_ontap_file_system}.
type FsxOntapFileSystem interface {
	cdktf.TerraformResource
	Arn() *string
	AutomaticBackupRetentionDays() *float64
	SetAutomaticBackupRetentionDays(val *float64)
	AutomaticBackupRetentionDaysInput() *float64
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
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
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_FsxOntapFileSystem) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_FsxOntapFileSystem) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystem) TagsInput() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html aws_fsx_ontap_file_system} Resource.
func NewFsxOntapFileSystem(scope constructs.Construct, id *string, config *FsxOntapFileSystemConfig) FsxOntapFileSystem {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystem{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html aws_fsx_ontap_file_system} Resource.
func NewFsxOntapFileSystem_Override(f FsxOntapFileSystem, scope constructs.Construct, id *string, config *FsxOntapFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystem",
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

func (j *jsiiProxy_FsxOntapFileSystem) SetCount(val interface{}) {
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

func (j *jsiiProxy_FsxOntapFileSystem) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxOntapFileSystem) SetTagsAll(val interface{}) {
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
		"hashicorp_aws.FSx.FsxOntapFileSystem",
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
		"hashicorp_aws.FSx.FsxOntapFileSystem",
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

type FsxOntapFileSystemConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#deployment_type FsxOntapFileSystem#deployment_type}.
	DeploymentType *string `json:"deploymentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#preferred_subnet_id FsxOntapFileSystem#preferred_subnet_id}.
	PreferredSubnetId *string `json:"preferredSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#subnet_ids FsxOntapFileSystem#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#throughput_capacity FsxOntapFileSystem#throughput_capacity}.
	ThroughputCapacity *float64 `json:"throughputCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#automatic_backup_retention_days FsxOntapFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#daily_automatic_backup_start_time FsxOntapFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// disk_iops_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#disk_iops_configuration FsxOntapFileSystem#disk_iops_configuration}
	DiskIopsConfiguration *FsxOntapFileSystemDiskIopsConfiguration `json:"diskIopsConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#endpoint_ip_address_range FsxOntapFileSystem#endpoint_ip_address_range}.
	EndpointIpAddressRange *string `json:"endpointIpAddressRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#fsx_admin_password FsxOntapFileSystem#fsx_admin_password}.
	FsxAdminPassword *string `json:"fsxAdminPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#kms_key_id FsxOntapFileSystem#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#route_table_ids FsxOntapFileSystem#route_table_ids}.
	RouteTableIds *[]*string `json:"routeTableIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#security_group_ids FsxOntapFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#storage_capacity FsxOntapFileSystem#storage_capacity}.
	StorageCapacity *float64 `json:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#storage_type FsxOntapFileSystem#storage_type}.
	StorageType *string `json:"storageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#tags FsxOntapFileSystem#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#tags_all FsxOntapFileSystem#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#timeouts FsxOntapFileSystem#timeouts}
	Timeouts *FsxOntapFileSystemTimeouts `json:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#weekly_maintenance_start_time FsxOntapFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

type FsxOntapFileSystemDiskIopsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#iops FsxOntapFileSystem#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#mode FsxOntapFileSystem#mode}.
	Mode *string `json:"mode"`
}

type FsxOntapFileSystemDiskIopsConfigurationOutputReference interface {
	cdktf.ComplexObject
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIops()
	ResetMode()
}

// The jsii proxy struct for FsxOntapFileSystemDiskIopsConfigurationOutputReference
type jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxOntapFileSystemDiskIopsConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FsxOntapFileSystemDiskIopsConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemDiskIopsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOntapFileSystemDiskIopsConfigurationOutputReference_Override(f FsxOntapFileSystemDiskIopsConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemDiskIopsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		f,
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

func (j *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemDiskIopsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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
	Intercluster() interface{}
	Management() interface{}
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

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) Intercluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intercluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) Management() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapFileSystemEndpoints(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) FsxOntapFileSystemEndpoints {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemEndpoints{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemEndpoints",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapFileSystemEndpoints_Override(f FsxOntapFileSystemEndpoints, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemEndpoints",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_FsxOntapFileSystemEndpoints) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpoints) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapFileSystemEndpointsIntercluster(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) FsxOntapFileSystemEndpointsIntercluster {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemEndpointsIntercluster{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemEndpointsIntercluster",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapFileSystemEndpointsIntercluster_Override(f FsxOntapFileSystemEndpointsIntercluster, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemEndpointsIntercluster",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsIntercluster) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewFsxOntapFileSystemEndpointsManagement(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) FsxOntapFileSystemEndpointsManagement {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemEndpointsManagement{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemEndpointsManagement",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewFsxOntapFileSystemEndpointsManagement_Override(f FsxOntapFileSystemEndpointsManagement, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemEndpointsManagement",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_FsxOntapFileSystemEndpointsManagement) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxOntapFileSystemEndpointsManagement) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#create FsxOntapFileSystem#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#delete FsxOntapFileSystem#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_ontap_file_system.html#update FsxOntapFileSystem#update}.
	Update *string `json:"update"`
}

type FsxOntapFileSystemTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func NewFsxOntapFileSystemTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FsxOntapFileSystemTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxOntapFileSystemTimeoutsOutputReference_Override(f FsxOntapFileSystemTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxOntapFileSystemTimeoutsOutputReference",
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

func (j *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (f *jsiiProxy_FsxOntapFileSystemTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html aws_fsx_windows_file_system}.
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
	Count() interface{}
	SetCount(val interface{})
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
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
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
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_FsxWindowsFileSystem) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_FsxWindowsFileSystem) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxWindowsFileSystem) TagsInput() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html aws_fsx_windows_file_system} Resource.
func NewFsxWindowsFileSystem(scope constructs.Construct, id *string, config *FsxWindowsFileSystemConfig) FsxWindowsFileSystem {
	_init_.Initialize()

	j := jsiiProxy_FsxWindowsFileSystem{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxWindowsFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html aws_fsx_windows_file_system} Resource.
func NewFsxWindowsFileSystem_Override(f FsxWindowsFileSystem, scope constructs.Construct, id *string, config *FsxWindowsFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxWindowsFileSystem",
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

func (j *jsiiProxy_FsxWindowsFileSystem) SetCount(val interface{}) {
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

func (j *jsiiProxy_FsxWindowsFileSystem) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxWindowsFileSystem) SetTagsAll(val interface{}) {
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
		"hashicorp_aws.FSx.FsxWindowsFileSystem",
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
		"hashicorp_aws.FSx.FsxWindowsFileSystem",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#audit_log_destination FsxWindowsFileSystem#audit_log_destination}.
	AuditLogDestination *string `json:"auditLogDestination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#file_access_audit_log_level FsxWindowsFileSystem#file_access_audit_log_level}.
	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#file_share_access_audit_log_level FsxWindowsFileSystem#file_share_access_audit_log_level}.
	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel"`
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

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewFsxWindowsFileSystemAuditLogConfigurationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FsxWindowsFileSystemAuditLogConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxWindowsFileSystemAuditLogConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxWindowsFileSystemAuditLogConfigurationOutputReference_Override(f FsxWindowsFileSystemAuditLogConfigurationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxWindowsFileSystemAuditLogConfigurationOutputReference",
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

func (j *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (f *jsiiProxy_FsxWindowsFileSystemAuditLogConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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

type FsxWindowsFileSystemConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#subnet_ids FsxWindowsFileSystem#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#throughput_capacity FsxWindowsFileSystem#throughput_capacity}.
	ThroughputCapacity *float64 `json:"throughputCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#active_directory_id FsxWindowsFileSystem#active_directory_id}.
	ActiveDirectoryId *string `json:"activeDirectoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#aliases FsxWindowsFileSystem#aliases}.
	Aliases *[]*string `json:"aliases"`
	// audit_log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#audit_log_configuration FsxWindowsFileSystem#audit_log_configuration}
	AuditLogConfiguration *FsxWindowsFileSystemAuditLogConfiguration `json:"auditLogConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#automatic_backup_retention_days FsxWindowsFileSystem#automatic_backup_retention_days}.
	AutomaticBackupRetentionDays *float64 `json:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#backup_id FsxWindowsFileSystem#backup_id}.
	BackupId *string `json:"backupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#copy_tags_to_backups FsxWindowsFileSystem#copy_tags_to_backups}.
	CopyTagsToBackups interface{} `json:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#daily_automatic_backup_start_time FsxWindowsFileSystem#daily_automatic_backup_start_time}.
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#deployment_type FsxWindowsFileSystem#deployment_type}.
	DeploymentType *string `json:"deploymentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#kms_key_id FsxWindowsFileSystem#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#preferred_subnet_id FsxWindowsFileSystem#preferred_subnet_id}.
	PreferredSubnetId *string `json:"preferredSubnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#security_group_ids FsxWindowsFileSystem#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// self_managed_active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#self_managed_active_directory FsxWindowsFileSystem#self_managed_active_directory}
	SelfManagedActiveDirectory *FsxWindowsFileSystemSelfManagedActiveDirectory `json:"selfManagedActiveDirectory"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#skip_final_backup FsxWindowsFileSystem#skip_final_backup}.
	SkipFinalBackup interface{} `json:"skipFinalBackup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#storage_capacity FsxWindowsFileSystem#storage_capacity}.
	StorageCapacity *float64 `json:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#storage_type FsxWindowsFileSystem#storage_type}.
	StorageType *string `json:"storageType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#tags FsxWindowsFileSystem#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#tags_all FsxWindowsFileSystem#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#timeouts FsxWindowsFileSystem#timeouts}
	Timeouts *FsxWindowsFileSystemTimeouts `json:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#weekly_maintenance_start_time FsxWindowsFileSystem#weekly_maintenance_start_time}.
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime"`
}

type FsxWindowsFileSystemSelfManagedActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#dns_ips FsxWindowsFileSystem#dns_ips}.
	DnsIps *[]*string `json:"dnsIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#domain_name FsxWindowsFileSystem#domain_name}.
	DomainName *string `json:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#password FsxWindowsFileSystem#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#username FsxWindowsFileSystem#username}.
	Username *string `json:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#file_system_administrators_group FsxWindowsFileSystem#file_system_administrators_group}.
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#organizational_unit_distinguished_name FsxWindowsFileSystem#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName"`
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func NewFsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference_Override(f FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference",
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

func (j *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (f *jsiiProxy_FsxWindowsFileSystemSelfManagedActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#create FsxWindowsFileSystem#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#delete FsxWindowsFileSystem#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/fsx_windows_file_system.html#update FsxWindowsFileSystem#update}.
	Update *string `json:"update"`
}

type FsxWindowsFileSystemTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func NewFsxWindowsFileSystemTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) FsxWindowsFileSystemTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxWindowsFileSystemTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewFsxWindowsFileSystemTimeoutsOutputReference_Override(f FsxWindowsFileSystemTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.FSx.FsxWindowsFileSystemTimeoutsOutputReference",
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

func (j *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (f *jsiiProxy_FsxWindowsFileSystemTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
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
