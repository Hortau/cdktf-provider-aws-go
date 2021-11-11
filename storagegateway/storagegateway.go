package storagegateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/storagegateway/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/storagegateway_local_disk.html aws_storagegateway_local_disk}.
type DataAwsStoragegatewayLocalDisk interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DiskId() *string
	DiskNode() *string
	SetDiskNode(val *string)
	DiskNodeInput() *string
	DiskPath() *string
	SetDiskPath(val *string)
	DiskPathInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
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
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDiskNode()
	ResetDiskPath()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsStoragegatewayLocalDisk
type jsiiProxy_DataAwsStoragegatewayLocalDisk struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) DiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) DiskNode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) DiskNodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) DiskPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) DiskPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/storagegateway_local_disk.html aws_storagegateway_local_disk} Data Source.
func NewDataAwsStoragegatewayLocalDisk(scope constructs.Construct, id *string, config *DataAwsStoragegatewayLocalDiskConfig) DataAwsStoragegatewayLocalDisk {
	_init_.Initialize()

	j := jsiiProxy_DataAwsStoragegatewayLocalDisk{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.DataAwsStoragegatewayLocalDisk",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/storagegateway_local_disk.html aws_storagegateway_local_disk} Data Source.
func NewDataAwsStoragegatewayLocalDisk_Override(d DataAwsStoragegatewayLocalDisk, scope constructs.Construct, id *string, config *DataAwsStoragegatewayLocalDiskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.DataAwsStoragegatewayLocalDisk",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) SetDiskNode(val *string) {
	_jsii_.Set(
		j,
		"diskNode",
		val,
	)
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) SetDiskPath(val *string) {
	_jsii_.Set(
		j,
		"diskPath",
		val,
	)
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) SetGatewayArn(val *string) {
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsStoragegatewayLocalDisk) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsStoragegatewayLocalDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.DataAwsStoragegatewayLocalDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsStoragegatewayLocalDisk_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.DataAwsStoragegatewayLocalDisk",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) ResetDiskNode() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskNode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) ResetDiskPath() {
	_jsii_.InvokeVoid(
		d,
		"resetDiskPath",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) ToString() *string {
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
func (d *jsiiProxy_DataAwsStoragegatewayLocalDisk) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsStoragegatewayLocalDiskConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/storagegateway_local_disk.html#gateway_arn DataAwsStoragegatewayLocalDisk#gateway_arn}.
	GatewayArn *string `json:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/storagegateway_local_disk.html#disk_node DataAwsStoragegatewayLocalDisk#disk_node}.
	DiskNode *string `json:"diskNode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/storagegateway_local_disk.html#disk_path DataAwsStoragegatewayLocalDisk#disk_path}.
	DiskPath *string `json:"diskPath"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cache.html aws_storagegateway_cache}.
type StoragegatewayCache interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DiskId() *string
	SetDiskId(val *string)
	DiskIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
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

// The jsii proxy struct for StoragegatewayCache
type jsiiProxy_StoragegatewayCache struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayCache) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) DiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) DiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCache) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cache.html aws_storagegateway_cache} Resource.
func NewStoragegatewayCache(scope constructs.Construct, id *string, config *StoragegatewayCacheConfig) StoragegatewayCache {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayCache{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayCache",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cache.html aws_storagegateway_cache} Resource.
func NewStoragegatewayCache_Override(s StoragegatewayCache, scope constructs.Construct, id *string, config *StoragegatewayCacheConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayCache",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayCache) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCache) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCache) SetDiskId(val *string) {
	_jsii_.Set(
		j,
		"diskId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCache) SetGatewayArn(val *string) {
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCache) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCache) SetProvider(val cdktf.TerraformProvider) {
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
func StoragegatewayCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewayCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayCache_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewayCache",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCache) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCache) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCache) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCache) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCache) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCache) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewayCache) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewayCache) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCache) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCache) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewayCache) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewayCache) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewayCacheConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cache.html#disk_id StoragegatewayCache#disk_id}.
	DiskId *string `json:"diskId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cache.html#gateway_arn StoragegatewayCache#gateway_arn}.
	GatewayArn *string `json:"gatewayArn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html aws_storagegateway_cached_iscsi_volume}.
type StoragegatewayCachedIscsiVolume interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ChapEnabled() interface{}
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
	Id() *string
	KmsEncrypted() interface{}
	SetKmsEncrypted(val interface{})
	KmsEncryptedInput() interface{}
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LunNumber() *float64
	NetworkInterfaceId() *string
	SetNetworkInterfaceId(val *string)
	NetworkInterfaceIdInput() *string
	NetworkInterfacePort() *float64
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	SourceVolumeArn() *string
	SetSourceVolumeArn(val *string)
	SourceVolumeArnInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TargetArn() *string
	TargetName() *string
	SetTargetName(val *string)
	TargetNameInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VolumeArn() *string
	VolumeId() *string
	VolumeSizeInBytes() *float64
	SetVolumeSizeInBytes(val *float64)
	VolumeSizeInBytesInput() *float64
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetKmsEncrypted()
	ResetKmsKey()
	ResetOverrideLogicalId()
	ResetSnapshotId()
	ResetSourceVolumeArn()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewayCachedIscsiVolume
type jsiiProxy_StoragegatewayCachedIscsiVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) ChapEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chapEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) LunNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lunNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) NetworkInterfacePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkInterfacePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SourceVolumeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVolumeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SourceVolumeArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVolumeArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TargetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TargetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) VolumeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) VolumeSizeInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) VolumeSizeInBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInBytesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html aws_storagegateway_cached_iscsi_volume} Resource.
func NewStoragegatewayCachedIscsiVolume(scope constructs.Construct, id *string, config *StoragegatewayCachedIscsiVolumeConfig) StoragegatewayCachedIscsiVolume {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayCachedIscsiVolume{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayCachedIscsiVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html aws_storagegateway_cached_iscsi_volume} Resource.
func NewStoragegatewayCachedIscsiVolume_Override(s StoragegatewayCachedIscsiVolume, scope constructs.Construct, id *string, config *StoragegatewayCachedIscsiVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayCachedIscsiVolume",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetGatewayArn(val *string) {
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetKmsEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetKmsKey(val *string) {
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetNetworkInterfaceId(val *string) {
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetSnapshotId(val *string) {
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetSourceVolumeArn(val *string) {
	_jsii_.Set(
		j,
		"sourceVolumeArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetTargetName(val *string) {
	_jsii_.Set(
		j,
		"targetName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SetVolumeSizeInBytes(val *float64) {
	_jsii_.Set(
		j,
		"volumeSizeInBytes",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func StoragegatewayCachedIscsiVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewayCachedIscsiVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayCachedIscsiVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewayCachedIscsiVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetKmsKey() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKey",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		s,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetSourceVolumeArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceVolumeArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewayCachedIscsiVolumeConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#gateway_arn StoragegatewayCachedIscsiVolume#gateway_arn}.
	GatewayArn *string `json:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#network_interface_id StoragegatewayCachedIscsiVolume#network_interface_id}.
	NetworkInterfaceId *string `json:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#target_name StoragegatewayCachedIscsiVolume#target_name}.
	TargetName *string `json:"targetName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#volume_size_in_bytes StoragegatewayCachedIscsiVolume#volume_size_in_bytes}.
	VolumeSizeInBytes *float64 `json:"volumeSizeInBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#kms_encrypted StoragegatewayCachedIscsiVolume#kms_encrypted}.
	KmsEncrypted interface{} `json:"kmsEncrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#kms_key StoragegatewayCachedIscsiVolume#kms_key}.
	KmsKey *string `json:"kmsKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#snapshot_id StoragegatewayCachedIscsiVolume#snapshot_id}.
	SnapshotId *string `json:"snapshotId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#source_volume_arn StoragegatewayCachedIscsiVolume#source_volume_arn}.
	SourceVolumeArn *string `json:"sourceVolumeArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#tags StoragegatewayCachedIscsiVolume#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_cached_iscsi_volume.html#tags_all StoragegatewayCachedIscsiVolume#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html aws_storagegateway_file_system_association}.
type StoragegatewayFileSystemAssociation interface {
	cdktf.TerraformResource
	Arn() *string
	AuditDestinationArn() *string
	SetAuditDestinationArn(val *string)
	AuditDestinationArnInput() *string
	CacheAttributes() StoragegatewayFileSystemAssociationCacheAttributesOutputReference
	CacheAttributesInput() *StoragegatewayFileSystemAssociationCacheAttributes
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationArn() *string
	SetLocationArn(val *string)
	LocationArnInput() *string
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutCacheAttributes(value *StoragegatewayFileSystemAssociationCacheAttributes)
	ResetAuditDestinationArn()
	ResetCacheAttributes()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewayFileSystemAssociation
type jsiiProxy_StoragegatewayFileSystemAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) AuditDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) AuditDestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) CacheAttributes() StoragegatewayFileSystemAssociationCacheAttributesOutputReference {
	var returns StoragegatewayFileSystemAssociationCacheAttributesOutputReference
	_jsii_.Get(
		j,
		"cacheAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) CacheAttributesInput() *StoragegatewayFileSystemAssociationCacheAttributes {
	var returns *StoragegatewayFileSystemAssociationCacheAttributes
	_jsii_.Get(
		j,
		"cacheAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) LocationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html aws_storagegateway_file_system_association} Resource.
func NewStoragegatewayFileSystemAssociation(scope constructs.Construct, id *string, config *StoragegatewayFileSystemAssociationConfig) StoragegatewayFileSystemAssociation {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayFileSystemAssociation{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayFileSystemAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html aws_storagegateway_file_system_association} Resource.
func NewStoragegatewayFileSystemAssociation_Override(s StoragegatewayFileSystemAssociation, scope constructs.Construct, id *string, config *StoragegatewayFileSystemAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayFileSystemAssociation",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetAuditDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"auditDestinationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetGatewayArn(val *string) {
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetLocationArn(val *string) {
	_jsii_.Set(
		j,
		"locationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociation) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func StoragegatewayFileSystemAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewayFileSystemAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayFileSystemAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewayFileSystemAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayFileSystemAssociation) PutCacheAttributes(value *StoragegatewayFileSystemAssociationCacheAttributes) {
	_jsii_.InvokeVoid(
		s,
		"putCacheAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayFileSystemAssociation) ResetAuditDestinationArn() {
	_jsii_.InvokeVoid(
		s,
		"resetAuditDestinationArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayFileSystemAssociation) ResetCacheAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheAttributes",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayFileSystemAssociation) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayFileSystemAssociation) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayFileSystemAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewayFileSystemAssociationCacheAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html#cache_stale_timeout_in_seconds StoragegatewayFileSystemAssociation#cache_stale_timeout_in_seconds}.
	CacheStaleTimeoutInSeconds *float64 `json:"cacheStaleTimeoutInSeconds"`
}

type StoragegatewayFileSystemAssociationCacheAttributesOutputReference interface {
	cdktf.ComplexObject
	CacheStaleTimeoutInSeconds() *float64
	SetCacheStaleTimeoutInSeconds(val *float64)
	CacheStaleTimeoutInSecondsInput() *float64
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
	ResetCacheStaleTimeoutInSeconds()
}

// The jsii proxy struct for StoragegatewayFileSystemAssociationCacheAttributesOutputReference
type jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) CacheStaleTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheStaleTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) CacheStaleTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheStaleTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewStoragegatewayFileSystemAssociationCacheAttributesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) StoragegatewayFileSystemAssociationCacheAttributesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayFileSystemAssociationCacheAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewStoragegatewayFileSystemAssociationCacheAttributesOutputReference_Override(s StoragegatewayFileSystemAssociationCacheAttributesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayFileSystemAssociationCacheAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) SetCacheStaleTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"cacheStaleTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayFileSystemAssociationCacheAttributesOutputReference) ResetCacheStaleTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheStaleTimeoutInSeconds",
		nil, // no parameters
	)
}

type StoragegatewayFileSystemAssociationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html#gateway_arn StoragegatewayFileSystemAssociation#gateway_arn}.
	GatewayArn *string `json:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html#location_arn StoragegatewayFileSystemAssociation#location_arn}.
	LocationArn *string `json:"locationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html#password StoragegatewayFileSystemAssociation#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html#username StoragegatewayFileSystemAssociation#username}.
	Username *string `json:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html#audit_destination_arn StoragegatewayFileSystemAssociation#audit_destination_arn}.
	AuditDestinationArn *string `json:"auditDestinationArn"`
	// cache_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html#cache_attributes StoragegatewayFileSystemAssociation#cache_attributes}
	CacheAttributes *StoragegatewayFileSystemAssociationCacheAttributes `json:"cacheAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html#tags StoragegatewayFileSystemAssociation#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_file_system_association.html#tags_all StoragegatewayFileSystemAssociation#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html aws_storagegateway_gateway}.
type StoragegatewayGateway interface {
	cdktf.TerraformResource
	ActivationKey() *string
	SetActivationKey(val *string)
	ActivationKeyInput() *string
	Arn() *string
	AverageDownloadRateLimitInBitsPerSec() *float64
	SetAverageDownloadRateLimitInBitsPerSec(val *float64)
	AverageDownloadRateLimitInBitsPerSecInput() *float64
	AverageUploadRateLimitInBitsPerSec() *float64
	SetAverageUploadRateLimitInBitsPerSec(val *float64)
	AverageUploadRateLimitInBitsPerSecInput() *float64
	CdktfStack() cdktf.TerraformStack
	CloudwatchLogGroupArn() *string
	SetCloudwatchLogGroupArn(val *string)
	CloudwatchLogGroupArnInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Ec2InstanceId() *string
	EndpointType() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayId() *string
	GatewayIpAddress() *string
	SetGatewayIpAddress(val *string)
	GatewayIpAddressInput() *string
	GatewayName() *string
	SetGatewayName(val *string)
	GatewayNameInput() *string
	GatewayTimezone() *string
	SetGatewayTimezone(val *string)
	GatewayTimezoneInput() *string
	GatewayType() *string
	SetGatewayType(val *string)
	GatewayTypeInput() *string
	GatewayVpcEndpoint() *string
	SetGatewayVpcEndpoint(val *string)
	GatewayVpcEndpointInput() *string
	HostEnvironment() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MediumChangerType() *string
	SetMediumChangerType(val *string)
	MediumChangerTypeInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SmbActiveDirectorySettings() StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference
	SmbActiveDirectorySettingsInput() *StoragegatewayGatewaySmbActiveDirectorySettings
	SmbFileShareVisibility() interface{}
	SetSmbFileShareVisibility(val interface{})
	SmbFileShareVisibilityInput() interface{}
	SmbGuestPassword() *string
	SetSmbGuestPassword(val *string)
	SmbGuestPasswordInput() *string
	SmbSecurityStrategy() *string
	SetSmbSecurityStrategy(val *string)
	SmbSecurityStrategyInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TapeDriveType() *string
	SetTapeDriveType(val *string)
	TapeDriveTypeInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() StoragegatewayGatewayTimeoutsOutputReference
	TimeoutsInput() *StoragegatewayGatewayTimeouts
	AddOverride(path *string, value interface{})
	GatewayNetworkInterface(index *string) StoragegatewayGatewayGatewayNetworkInterface
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutSmbActiveDirectorySettings(value *StoragegatewayGatewaySmbActiveDirectorySettings)
	PutTimeouts(value *StoragegatewayGatewayTimeouts)
	ResetActivationKey()
	ResetAverageDownloadRateLimitInBitsPerSec()
	ResetAverageUploadRateLimitInBitsPerSec()
	ResetCloudwatchLogGroupArn()
	ResetGatewayIpAddress()
	ResetGatewayType()
	ResetGatewayVpcEndpoint()
	ResetMediumChangerType()
	ResetOverrideLogicalId()
	ResetSmbActiveDirectorySettings()
	ResetSmbFileShareVisibility()
	ResetSmbGuestPassword()
	ResetSmbSecurityStrategy()
	ResetTags()
	ResetTagsAll()
	ResetTapeDriveType()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewayGateway
type jsiiProxy_StoragegatewayGateway struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayGateway) ActivationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) ActivationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) AverageDownloadRateLimitInBitsPerSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageDownloadRateLimitInBitsPerSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) AverageDownloadRateLimitInBitsPerSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageDownloadRateLimitInBitsPerSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) AverageUploadRateLimitInBitsPerSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageUploadRateLimitInBitsPerSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) AverageUploadRateLimitInBitsPerSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"averageUploadRateLimitInBitsPerSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) CloudwatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) CloudwatchLogGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Ec2InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayVpcEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayVpcEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) GatewayVpcEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayVpcEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) HostEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) MediumChangerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediumChangerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) MediumChangerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediumChangerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbActiveDirectorySettings() StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference {
	var returns StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference
	_jsii_.Get(
		j,
		"smbActiveDirectorySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbActiveDirectorySettingsInput() *StoragegatewayGatewaySmbActiveDirectorySettings {
	var returns *StoragegatewayGatewaySmbActiveDirectorySettings
	_jsii_.Get(
		j,
		"smbActiveDirectorySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbFileShareVisibility() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbFileShareVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbFileShareVisibilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbFileShareVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbGuestPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbGuestPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbGuestPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbGuestPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbSecurityStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbSecurityStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) SmbSecurityStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbSecurityStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TapeDriveType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tapeDriveType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TapeDriveTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tapeDriveTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) Timeouts() StoragegatewayGatewayTimeoutsOutputReference {
	var returns StoragegatewayGatewayTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGateway) TimeoutsInput() *StoragegatewayGatewayTimeouts {
	var returns *StoragegatewayGatewayTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html aws_storagegateway_gateway} Resource.
func NewStoragegatewayGateway(scope constructs.Construct, id *string, config *StoragegatewayGatewayConfig) StoragegatewayGateway {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayGateway{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayGateway",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html aws_storagegateway_gateway} Resource.
func NewStoragegatewayGateway_Override(s StoragegatewayGateway, scope constructs.Construct, id *string, config *StoragegatewayGatewayConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayGateway",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetActivationKey(val *string) {
	_jsii_.Set(
		j,
		"activationKey",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetAverageDownloadRateLimitInBitsPerSec(val *float64) {
	_jsii_.Set(
		j,
		"averageDownloadRateLimitInBitsPerSec",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetAverageUploadRateLimitInBitsPerSec(val *float64) {
	_jsii_.Set(
		j,
		"averageUploadRateLimitInBitsPerSec",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetCloudwatchLogGroupArn(val *string) {
	_jsii_.Set(
		j,
		"cloudwatchLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetGatewayIpAddress(val *string) {
	_jsii_.Set(
		j,
		"gatewayIpAddress",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetGatewayName(val *string) {
	_jsii_.Set(
		j,
		"gatewayName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetGatewayTimezone(val *string) {
	_jsii_.Set(
		j,
		"gatewayTimezone",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetGatewayType(val *string) {
	_jsii_.Set(
		j,
		"gatewayType",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetGatewayVpcEndpoint(val *string) {
	_jsii_.Set(
		j,
		"gatewayVpcEndpoint",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetMediumChangerType(val *string) {
	_jsii_.Set(
		j,
		"mediumChangerType",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetSmbFileShareVisibility(val interface{}) {
	_jsii_.Set(
		j,
		"smbFileShareVisibility",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetSmbGuestPassword(val *string) {
	_jsii_.Set(
		j,
		"smbGuestPassword",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetSmbSecurityStrategy(val *string) {
	_jsii_.Set(
		j,
		"smbSecurityStrategy",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGateway) SetTapeDriveType(val *string) {
	_jsii_.Set(
		j,
		"tapeDriveType",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func StoragegatewayGateway_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewayGateway",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayGateway_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewayGateway",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StoragegatewayGateway) GatewayNetworkInterface(index *string) StoragegatewayGatewayGatewayNetworkInterface {
	var returns StoragegatewayGatewayGatewayNetworkInterface

	_jsii_.Invoke(
		s,
		"gatewayNetworkInterface",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayGateway) PutSmbActiveDirectorySettings(value *StoragegatewayGatewaySmbActiveDirectorySettings) {
	_jsii_.InvokeVoid(
		s,
		"putSmbActiveDirectorySettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayGateway) PutTimeouts(value *StoragegatewayGatewayTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetActivationKey() {
	_jsii_.InvokeVoid(
		s,
		"resetActivationKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetAverageDownloadRateLimitInBitsPerSec() {
	_jsii_.InvokeVoid(
		s,
		"resetAverageDownloadRateLimitInBitsPerSec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetAverageUploadRateLimitInBitsPerSec() {
	_jsii_.InvokeVoid(
		s,
		"resetAverageUploadRateLimitInBitsPerSec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetCloudwatchLogGroupArn() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatchLogGroupArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetGatewayIpAddress() {
	_jsii_.InvokeVoid(
		s,
		"resetGatewayIpAddress",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetGatewayType() {
	_jsii_.InvokeVoid(
		s,
		"resetGatewayType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetGatewayVpcEndpoint() {
	_jsii_.InvokeVoid(
		s,
		"resetGatewayVpcEndpoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetMediumChangerType() {
	_jsii_.InvokeVoid(
		s,
		"resetMediumChangerType",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetSmbActiveDirectorySettings() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbActiveDirectorySettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetSmbFileShareVisibility() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbFileShareVisibility",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetSmbGuestPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbGuestPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetSmbSecurityStrategy() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbSecurityStrategy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetTapeDriveType() {
	_jsii_.InvokeVoid(
		s,
		"resetTapeDriveType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGateway) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewayGateway) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewayGateway) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewayGatewayConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#gateway_name StoragegatewayGateway#gateway_name}.
	GatewayName *string `json:"gatewayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#gateway_timezone StoragegatewayGateway#gateway_timezone}.
	GatewayTimezone *string `json:"gatewayTimezone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#activation_key StoragegatewayGateway#activation_key}.
	ActivationKey *string `json:"activationKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#average_download_rate_limit_in_bits_per_sec StoragegatewayGateway#average_download_rate_limit_in_bits_per_sec}.
	AverageDownloadRateLimitInBitsPerSec *float64 `json:"averageDownloadRateLimitInBitsPerSec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#average_upload_rate_limit_in_bits_per_sec StoragegatewayGateway#average_upload_rate_limit_in_bits_per_sec}.
	AverageUploadRateLimitInBitsPerSec *float64 `json:"averageUploadRateLimitInBitsPerSec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#cloudwatch_log_group_arn StoragegatewayGateway#cloudwatch_log_group_arn}.
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#gateway_ip_address StoragegatewayGateway#gateway_ip_address}.
	GatewayIpAddress *string `json:"gatewayIpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#gateway_type StoragegatewayGateway#gateway_type}.
	GatewayType *string `json:"gatewayType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#gateway_vpc_endpoint StoragegatewayGateway#gateway_vpc_endpoint}.
	GatewayVpcEndpoint *string `json:"gatewayVpcEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#medium_changer_type StoragegatewayGateway#medium_changer_type}.
	MediumChangerType *string `json:"mediumChangerType"`
	// smb_active_directory_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#smb_active_directory_settings StoragegatewayGateway#smb_active_directory_settings}
	SmbActiveDirectorySettings *StoragegatewayGatewaySmbActiveDirectorySettings `json:"smbActiveDirectorySettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#smb_file_share_visibility StoragegatewayGateway#smb_file_share_visibility}.
	SmbFileShareVisibility interface{} `json:"smbFileShareVisibility"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#smb_guest_password StoragegatewayGateway#smb_guest_password}.
	SmbGuestPassword *string `json:"smbGuestPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#smb_security_strategy StoragegatewayGateway#smb_security_strategy}.
	SmbSecurityStrategy *string `json:"smbSecurityStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#tags StoragegatewayGateway#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#tags_all StoragegatewayGateway#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#tape_drive_type StoragegatewayGateway#tape_drive_type}.
	TapeDriveType *string `json:"tapeDriveType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#timeouts StoragegatewayGateway#timeouts}
	Timeouts *StoragegatewayGatewayTimeouts `json:"timeouts"`
}

type StoragegatewayGatewayGatewayNetworkInterface interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Ipv4Address() *string
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

// The jsii proxy struct for StoragegatewayGatewayGatewayNetworkInterface
type jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) Ipv4Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewStoragegatewayGatewayGatewayNetworkInterface(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) StoragegatewayGatewayGatewayNetworkInterface {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayGatewayGatewayNetworkInterface",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewStoragegatewayGatewayGatewayNetworkInterface_Override(s StoragegatewayGatewayGatewayNetworkInterface, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayGatewayGatewayNetworkInterface",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayGatewayNetworkInterface) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type StoragegatewayGatewaySmbActiveDirectorySettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#domain_name StoragegatewayGateway#domain_name}.
	DomainName *string `json:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#password StoragegatewayGateway#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#username StoragegatewayGateway#username}.
	Username *string `json:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#domain_controllers StoragegatewayGateway#domain_controllers}.
	DomainControllers *[]*string `json:"domainControllers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#organizational_unit StoragegatewayGateway#organizational_unit}.
	OrganizationalUnit *string `json:"organizationalUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#timeout_in_seconds StoragegatewayGateway#timeout_in_seconds}.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds"`
}

type StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference interface {
	cdktf.ComplexObject
	DomainControllers() *[]*string
	SetDomainControllers(val *[]*string)
	DomainControllersInput() *[]*string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OrganizationalUnit() *string
	SetOrganizationalUnit(val *string)
	OrganizationalUnitInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDomainControllers()
	ResetOrganizationalUnit()
	ResetTimeoutInSeconds()
}

// The jsii proxy struct for StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference
type jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) DomainControllers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainControllers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) DomainControllersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainControllersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) OrganizationalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) OrganizationalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func NewStoragegatewayGatewaySmbActiveDirectorySettingsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewStoragegatewayGatewaySmbActiveDirectorySettingsOutputReference_Override(s StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) SetDomainControllers(val *[]*string) {
	_jsii_.Set(
		j,
		"domainControllers",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) SetOrganizationalUnit(val *string) {
	_jsii_.Set(
		j,
		"organizationalUnit",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) SetTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) ResetDomainControllers() {
	_jsii_.InvokeVoid(
		s,
		"resetDomainControllers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) ResetOrganizationalUnit() {
	_jsii_.InvokeVoid(
		s,
		"resetOrganizationalUnit",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayGatewaySmbActiveDirectorySettingsOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

type StoragegatewayGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_gateway.html#create StoragegatewayGateway#create}.
	Create *string `json:"create"`
}

type StoragegatewayGatewayTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
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
}

// The jsii proxy struct for StoragegatewayGatewayTimeoutsOutputReference
type jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewStoragegatewayGatewayTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) StoragegatewayGatewayTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayGatewayTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewStoragegatewayGatewayTimeoutsOutputReference_Override(s StoragegatewayGatewayTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayGatewayTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayGatewayTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html aws_storagegateway_nfs_file_share}.
type StoragegatewayNfsFileShare interface {
	cdktf.TerraformResource
	Arn() *string
	AuditDestinationArn() *string
	SetAuditDestinationArn(val *string)
	AuditDestinationArnInput() *string
	CacheAttributes() StoragegatewayNfsFileShareCacheAttributesOutputReference
	CacheAttributesInput() *StoragegatewayNfsFileShareCacheAttributes
	CdktfStack() cdktf.TerraformStack
	ClientList() *[]*string
	SetClientList(val *[]*string)
	ClientListInput() *[]*string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultStorageClass() *string
	SetDefaultStorageClass(val *string)
	DefaultStorageClassInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FileshareId() *string
	FileShareName() *string
	SetFileShareName(val *string)
	FileShareNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
	GuessMimeTypeEnabled() interface{}
	SetGuessMimeTypeEnabled(val interface{})
	GuessMimeTypeEnabledInput() interface{}
	Id() *string
	KmsEncrypted() interface{}
	SetKmsEncrypted(val interface{})
	KmsEncryptedInput() interface{}
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationArn() *string
	SetLocationArn(val *string)
	LocationArnInput() *string
	NfsFileShareDefaults() StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference
	NfsFileShareDefaultsInput() *StoragegatewayNfsFileShareNfsFileShareDefaults
	Node() constructs.Node
	NotificationPolicy() *string
	SetNotificationPolicy(val *string)
	NotificationPolicyInput() *string
	ObjectAcl() *string
	SetObjectAcl(val *string)
	ObjectAclInput() *string
	Path() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	RequesterPays() interface{}
	SetRequesterPays(val interface{})
	RequesterPaysInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Squash() *string
	SetSquash(val *string)
	SquashInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() StoragegatewayNfsFileShareTimeoutsOutputReference
	TimeoutsInput() *StoragegatewayNfsFileShareTimeouts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutCacheAttributes(value *StoragegatewayNfsFileShareCacheAttributes)
	PutNfsFileShareDefaults(value *StoragegatewayNfsFileShareNfsFileShareDefaults)
	PutTimeouts(value *StoragegatewayNfsFileShareTimeouts)
	ResetAuditDestinationArn()
	ResetCacheAttributes()
	ResetDefaultStorageClass()
	ResetFileShareName()
	ResetGuessMimeTypeEnabled()
	ResetKmsEncrypted()
	ResetKmsKeyArn()
	ResetNfsFileShareDefaults()
	ResetNotificationPolicy()
	ResetObjectAcl()
	ResetOverrideLogicalId()
	ResetReadOnly()
	ResetRequesterPays()
	ResetSquash()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewayNfsFileShare
type jsiiProxy_StoragegatewayNfsFileShare struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) AuditDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) AuditDestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) CacheAttributes() StoragegatewayNfsFileShareCacheAttributesOutputReference {
	var returns StoragegatewayNfsFileShareCacheAttributesOutputReference
	_jsii_.Get(
		j,
		"cacheAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) CacheAttributesInput() *StoragegatewayNfsFileShareCacheAttributes {
	var returns *StoragegatewayNfsFileShareCacheAttributes
	_jsii_.Get(
		j,
		"cacheAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ClientList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ClientListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) DefaultStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) DefaultStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) FileshareId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileshareId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) FileShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) FileShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) GuessMimeTypeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) GuessMimeTypeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) LocationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) NfsFileShareDefaults() StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference {
	var returns StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference
	_jsii_.Get(
		j,
		"nfsFileShareDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) NfsFileShareDefaultsInput() *StoragegatewayNfsFileShareNfsFileShareDefaults {
	var returns *StoragegatewayNfsFileShareNfsFileShareDefaults
	_jsii_.Get(
		j,
		"nfsFileShareDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) NotificationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) NotificationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RequesterPays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RequesterPaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Squash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SquashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) Timeouts() StoragegatewayNfsFileShareTimeoutsOutputReference {
	var returns StoragegatewayNfsFileShareTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) TimeoutsInput() *StoragegatewayNfsFileShareTimeouts {
	var returns *StoragegatewayNfsFileShareTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html aws_storagegateway_nfs_file_share} Resource.
func NewStoragegatewayNfsFileShare(scope constructs.Construct, id *string, config *StoragegatewayNfsFileShareConfig) StoragegatewayNfsFileShare {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayNfsFileShare{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html aws_storagegateway_nfs_file_share} Resource.
func NewStoragegatewayNfsFileShare_Override(s StoragegatewayNfsFileShare, scope constructs.Construct, id *string, config *StoragegatewayNfsFileShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShare",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetAuditDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"auditDestinationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetClientList(val *[]*string) {
	_jsii_.Set(
		j,
		"clientList",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetDefaultStorageClass(val *string) {
	_jsii_.Set(
		j,
		"defaultStorageClass",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetFileShareName(val *string) {
	_jsii_.Set(
		j,
		"fileShareName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetGatewayArn(val *string) {
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetGuessMimeTypeEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"guessMimeTypeEnabled",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetKmsEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetLocationArn(val *string) {
	_jsii_.Set(
		j,
		"locationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetNotificationPolicy(val *string) {
	_jsii_.Set(
		j,
		"notificationPolicy",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetObjectAcl(val *string) {
	_jsii_.Set(
		j,
		"objectAcl",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetReadOnly(val interface{}) {
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetRequesterPays(val interface{}) {
	_jsii_.Set(
		j,
		"requesterPays",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetSquash(val *string) {
	_jsii_.Set(
		j,
		"squash",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShare) SetTagsAll(val interface{}) {
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
func StoragegatewayNfsFileShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayNfsFileShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) PutCacheAttributes(value *StoragegatewayNfsFileShareCacheAttributes) {
	_jsii_.InvokeVoid(
		s,
		"putCacheAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) PutNfsFileShareDefaults(value *StoragegatewayNfsFileShareNfsFileShareDefaults) {
	_jsii_.InvokeVoid(
		s,
		"putNfsFileShareDefaults",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) PutTimeouts(value *StoragegatewayNfsFileShareTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetAuditDestinationArn() {
	_jsii_.InvokeVoid(
		s,
		"resetAuditDestinationArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetCacheAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetDefaultStorageClass() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultStorageClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetFileShareName() {
	_jsii_.InvokeVoid(
		s,
		"resetFileShareName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetGuessMimeTypeEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetGuessMimeTypeEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetNfsFileShareDefaults() {
	_jsii_.InvokeVoid(
		s,
		"resetNfsFileShareDefaults",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetNotificationPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetObjectAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectAcl",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetReadOnly() {
	_jsii_.InvokeVoid(
		s,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetRequesterPays() {
	_jsii_.InvokeVoid(
		s,
		"resetRequesterPays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetSquash() {
	_jsii_.InvokeVoid(
		s,
		"resetSquash",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewayNfsFileShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewayNfsFileShareCacheAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#cache_stale_timeout_in_seconds StoragegatewayNfsFileShare#cache_stale_timeout_in_seconds}.
	CacheStaleTimeoutInSeconds *float64 `json:"cacheStaleTimeoutInSeconds"`
}

type StoragegatewayNfsFileShareCacheAttributesOutputReference interface {
	cdktf.ComplexObject
	CacheStaleTimeoutInSeconds() *float64
	SetCacheStaleTimeoutInSeconds(val *float64)
	CacheStaleTimeoutInSecondsInput() *float64
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
	ResetCacheStaleTimeoutInSeconds()
}

// The jsii proxy struct for StoragegatewayNfsFileShareCacheAttributesOutputReference
type jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) CacheStaleTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheStaleTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) CacheStaleTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheStaleTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewStoragegatewayNfsFileShareCacheAttributesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) StoragegatewayNfsFileShareCacheAttributesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShareCacheAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewStoragegatewayNfsFileShareCacheAttributesOutputReference_Override(s StoragegatewayNfsFileShareCacheAttributesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShareCacheAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) SetCacheStaleTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"cacheStaleTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShareCacheAttributesOutputReference) ResetCacheStaleTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheStaleTimeoutInSeconds",
		nil, // no parameters
	)
}

type StoragegatewayNfsFileShareConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#client_list StoragegatewayNfsFileShare#client_list}.
	ClientList *[]*string `json:"clientList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#gateway_arn StoragegatewayNfsFileShare#gateway_arn}.
	GatewayArn *string `json:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#location_arn StoragegatewayNfsFileShare#location_arn}.
	LocationArn *string `json:"locationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#role_arn StoragegatewayNfsFileShare#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#audit_destination_arn StoragegatewayNfsFileShare#audit_destination_arn}.
	AuditDestinationArn *string `json:"auditDestinationArn"`
	// cache_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#cache_attributes StoragegatewayNfsFileShare#cache_attributes}
	CacheAttributes *StoragegatewayNfsFileShareCacheAttributes `json:"cacheAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#default_storage_class StoragegatewayNfsFileShare#default_storage_class}.
	DefaultStorageClass *string `json:"defaultStorageClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#file_share_name StoragegatewayNfsFileShare#file_share_name}.
	FileShareName *string `json:"fileShareName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#guess_mime_type_enabled StoragegatewayNfsFileShare#guess_mime_type_enabled}.
	GuessMimeTypeEnabled interface{} `json:"guessMimeTypeEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#kms_encrypted StoragegatewayNfsFileShare#kms_encrypted}.
	KmsEncrypted interface{} `json:"kmsEncrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#kms_key_arn StoragegatewayNfsFileShare#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn"`
	// nfs_file_share_defaults block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#nfs_file_share_defaults StoragegatewayNfsFileShare#nfs_file_share_defaults}
	NfsFileShareDefaults *StoragegatewayNfsFileShareNfsFileShareDefaults `json:"nfsFileShareDefaults"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#notification_policy StoragegatewayNfsFileShare#notification_policy}.
	NotificationPolicy *string `json:"notificationPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#object_acl StoragegatewayNfsFileShare#object_acl}.
	ObjectAcl *string `json:"objectAcl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#read_only StoragegatewayNfsFileShare#read_only}.
	ReadOnly interface{} `json:"readOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#requester_pays StoragegatewayNfsFileShare#requester_pays}.
	RequesterPays interface{} `json:"requesterPays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#squash StoragegatewayNfsFileShare#squash}.
	Squash *string `json:"squash"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#tags StoragegatewayNfsFileShare#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#tags_all StoragegatewayNfsFileShare#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#timeouts StoragegatewayNfsFileShare#timeouts}
	Timeouts *StoragegatewayNfsFileShareTimeouts `json:"timeouts"`
}

type StoragegatewayNfsFileShareNfsFileShareDefaults struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#directory_mode StoragegatewayNfsFileShare#directory_mode}.
	DirectoryMode *string `json:"directoryMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#file_mode StoragegatewayNfsFileShare#file_mode}.
	FileMode *string `json:"fileMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#group_id StoragegatewayNfsFileShare#group_id}.
	GroupId *string `json:"groupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#owner_id StoragegatewayNfsFileShare#owner_id}.
	OwnerId *string `json:"ownerId"`
}

type StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference interface {
	cdktf.ComplexObject
	DirectoryMode() *string
	SetDirectoryMode(val *string)
	DirectoryModeInput() *string
	FileMode() *string
	SetFileMode(val *string)
	FileModeInput() *string
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OwnerId() *string
	SetOwnerId(val *string)
	OwnerIdInput() *string
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
	ResetDirectoryMode()
	ResetFileMode()
	ResetGroupId()
	ResetOwnerId()
}

// The jsii proxy struct for StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference
type jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) DirectoryMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) DirectoryModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) FileMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) FileModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) OwnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewStoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewStoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference_Override(s StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) SetDirectoryMode(val *string) {
	_jsii_.Set(
		j,
		"directoryMode",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) SetFileMode(val *string) {
	_jsii_.Set(
		j,
		"fileMode",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) SetGroupId(val *string) {
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) SetOwnerId(val *string) {
	_jsii_.Set(
		j,
		"ownerId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ResetDirectoryMode() {
	_jsii_.InvokeVoid(
		s,
		"resetDirectoryMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ResetFileMode() {
	_jsii_.InvokeVoid(
		s,
		"resetFileMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ResetGroupId() {
	_jsii_.InvokeVoid(
		s,
		"resetGroupId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareNfsFileShareDefaultsOutputReference) ResetOwnerId() {
	_jsii_.InvokeVoid(
		s,
		"resetOwnerId",
		nil, // no parameters
	)
}

type StoragegatewayNfsFileShareTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#create StoragegatewayNfsFileShare#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#delete StoragegatewayNfsFileShare#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_nfs_file_share.html#update StoragegatewayNfsFileShare#update}.
	Update *string `json:"update"`
}

type StoragegatewayNfsFileShareTimeoutsOutputReference interface {
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

// The jsii proxy struct for StoragegatewayNfsFileShareTimeoutsOutputReference
type jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewStoragegatewayNfsFileShareTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) StoragegatewayNfsFileShareTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShareTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewStoragegatewayNfsFileShareTimeoutsOutputReference_Override(s StoragegatewayNfsFileShareTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayNfsFileShareTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayNfsFileShareTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html aws_storagegateway_smb_file_share}.
type StoragegatewaySmbFileShare interface {
	cdktf.TerraformResource
	AccessBasedEnumeration() interface{}
	SetAccessBasedEnumeration(val interface{})
	AccessBasedEnumerationInput() interface{}
	AdminUserList() *[]*string
	SetAdminUserList(val *[]*string)
	AdminUserListInput() *[]*string
	Arn() *string
	AuditDestinationArn() *string
	SetAuditDestinationArn(val *string)
	AuditDestinationArnInput() *string
	Authentication() *string
	SetAuthentication(val *string)
	AuthenticationInput() *string
	BucketRegion() *string
	SetBucketRegion(val *string)
	BucketRegionInput() *string
	CacheAttributes() StoragegatewaySmbFileShareCacheAttributesOutputReference
	CacheAttributesInput() *StoragegatewaySmbFileShareCacheAttributes
	CaseSensitivity() *string
	SetCaseSensitivity(val *string)
	CaseSensitivityInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultStorageClass() *string
	SetDefaultStorageClass(val *string)
	DefaultStorageClassInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FileshareId() *string
	FileShareName() *string
	SetFileShareName(val *string)
	FileShareNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
	GuessMimeTypeEnabled() interface{}
	SetGuessMimeTypeEnabled(val interface{})
	GuessMimeTypeEnabledInput() interface{}
	Id() *string
	InvalidUserList() *[]*string
	SetInvalidUserList(val *[]*string)
	InvalidUserListInput() *[]*string
	KmsEncrypted() interface{}
	SetKmsEncrypted(val interface{})
	KmsEncryptedInput() interface{}
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationArn() *string
	SetLocationArn(val *string)
	LocationArnInput() *string
	Node() constructs.Node
	NotificationPolicy() *string
	SetNotificationPolicy(val *string)
	NotificationPolicyInput() *string
	ObjectAcl() *string
	SetObjectAcl(val *string)
	ObjectAclInput() *string
	OplocksEnabled() interface{}
	SetOplocksEnabled(val interface{})
	OplocksEnabledInput() interface{}
	Path() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	RequesterPays() interface{}
	SetRequesterPays(val interface{})
	RequesterPaysInput() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SmbAclEnabled() interface{}
	SetSmbAclEnabled(val interface{})
	SmbAclEnabledInput() interface{}
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() StoragegatewaySmbFileShareTimeoutsOutputReference
	TimeoutsInput() *StoragegatewaySmbFileShareTimeouts
	ValidUserList() *[]*string
	SetValidUserList(val *[]*string)
	ValidUserListInput() *[]*string
	VpcEndpointDnsName() *string
	SetVpcEndpointDnsName(val *string)
	VpcEndpointDnsNameInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutCacheAttributes(value *StoragegatewaySmbFileShareCacheAttributes)
	PutTimeouts(value *StoragegatewaySmbFileShareTimeouts)
	ResetAccessBasedEnumeration()
	ResetAdminUserList()
	ResetAuditDestinationArn()
	ResetAuthentication()
	ResetBucketRegion()
	ResetCacheAttributes()
	ResetCaseSensitivity()
	ResetDefaultStorageClass()
	ResetFileShareName()
	ResetGuessMimeTypeEnabled()
	ResetInvalidUserList()
	ResetKmsEncrypted()
	ResetKmsKeyArn()
	ResetNotificationPolicy()
	ResetObjectAcl()
	ResetOplocksEnabled()
	ResetOverrideLogicalId()
	ResetReadOnly()
	ResetRequesterPays()
	ResetSmbAclEnabled()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetValidUserList()
	ResetVpcEndpointDnsName()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewaySmbFileShare
type jsiiProxy_StoragegatewaySmbFileShare struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AccessBasedEnumeration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessBasedEnumeration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AccessBasedEnumerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessBasedEnumerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AdminUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AdminUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AuditDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AuditDestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Authentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) AuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) BucketRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CacheAttributes() StoragegatewaySmbFileShareCacheAttributesOutputReference {
	var returns StoragegatewaySmbFileShareCacheAttributesOutputReference
	_jsii_.Get(
		j,
		"cacheAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CacheAttributesInput() *StoragegatewaySmbFileShareCacheAttributes {
	var returns *StoragegatewaySmbFileShareCacheAttributes
	_jsii_.Get(
		j,
		"cacheAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CaseSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caseSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CaseSensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caseSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) DefaultStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) DefaultStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) FileshareId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileshareId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) FileShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) FileShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) GuessMimeTypeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) GuessMimeTypeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) InvalidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) InvalidUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) LocationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) NotificationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) NotificationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) OplocksEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oplocksEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) OplocksEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oplocksEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RequesterPays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RequesterPaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SmbAclEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbAclEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SmbAclEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbAclEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) Timeouts() StoragegatewaySmbFileShareTimeoutsOutputReference {
	var returns StoragegatewaySmbFileShareTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) TimeoutsInput() *StoragegatewaySmbFileShareTimeouts {
	var returns *StoragegatewaySmbFileShareTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ValidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) ValidUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) VpcEndpointDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) VpcEndpointDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsNameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html aws_storagegateway_smb_file_share} Resource.
func NewStoragegatewaySmbFileShare(scope constructs.Construct, id *string, config *StoragegatewaySmbFileShareConfig) StoragegatewaySmbFileShare {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewaySmbFileShare{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewaySmbFileShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html aws_storagegateway_smb_file_share} Resource.
func NewStoragegatewaySmbFileShare_Override(s StoragegatewaySmbFileShare, scope constructs.Construct, id *string, config *StoragegatewaySmbFileShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewaySmbFileShare",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetAccessBasedEnumeration(val interface{}) {
	_jsii_.Set(
		j,
		"accessBasedEnumeration",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetAdminUserList(val *[]*string) {
	_jsii_.Set(
		j,
		"adminUserList",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetAuditDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"auditDestinationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetAuthentication(val *string) {
	_jsii_.Set(
		j,
		"authentication",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetBucketRegion(val *string) {
	_jsii_.Set(
		j,
		"bucketRegion",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetCaseSensitivity(val *string) {
	_jsii_.Set(
		j,
		"caseSensitivity",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetDefaultStorageClass(val *string) {
	_jsii_.Set(
		j,
		"defaultStorageClass",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetFileShareName(val *string) {
	_jsii_.Set(
		j,
		"fileShareName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetGatewayArn(val *string) {
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetGuessMimeTypeEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"guessMimeTypeEnabled",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetInvalidUserList(val *[]*string) {
	_jsii_.Set(
		j,
		"invalidUserList",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetKmsEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetLocationArn(val *string) {
	_jsii_.Set(
		j,
		"locationArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetNotificationPolicy(val *string) {
	_jsii_.Set(
		j,
		"notificationPolicy",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetObjectAcl(val *string) {
	_jsii_.Set(
		j,
		"objectAcl",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetOplocksEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"oplocksEnabled",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetReadOnly(val interface{}) {
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetRequesterPays(val interface{}) {
	_jsii_.Set(
		j,
		"requesterPays",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetSmbAclEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"smbAclEnabled",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetValidUserList(val *[]*string) {
	_jsii_.Set(
		j,
		"validUserList",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShare) SetVpcEndpointDnsName(val *string) {
	_jsii_.Set(
		j,
		"vpcEndpointDnsName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func StoragegatewaySmbFileShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewaySmbFileShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewaySmbFileShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewaySmbFileShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) PutCacheAttributes(value *StoragegatewaySmbFileShareCacheAttributes) {
	_jsii_.InvokeVoid(
		s,
		"putCacheAttributes",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) PutTimeouts(value *StoragegatewaySmbFileShareTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetAccessBasedEnumeration() {
	_jsii_.InvokeVoid(
		s,
		"resetAccessBasedEnumeration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetAdminUserList() {
	_jsii_.InvokeVoid(
		s,
		"resetAdminUserList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetAuditDestinationArn() {
	_jsii_.InvokeVoid(
		s,
		"resetAuditDestinationArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetAuthentication() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetBucketRegion() {
	_jsii_.InvokeVoid(
		s,
		"resetBucketRegion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetCacheAttributes() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheAttributes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetCaseSensitivity() {
	_jsii_.InvokeVoid(
		s,
		"resetCaseSensitivity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetDefaultStorageClass() {
	_jsii_.InvokeVoid(
		s,
		"resetDefaultStorageClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetFileShareName() {
	_jsii_.InvokeVoid(
		s,
		"resetFileShareName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetGuessMimeTypeEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetGuessMimeTypeEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetInvalidUserList() {
	_jsii_.InvokeVoid(
		s,
		"resetInvalidUserList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetNotificationPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetNotificationPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetObjectAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectAcl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetOplocksEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetOplocksEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetReadOnly() {
	_jsii_.InvokeVoid(
		s,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetRequesterPays() {
	_jsii_.InvokeVoid(
		s,
		"resetRequesterPays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetSmbAclEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbAclEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetValidUserList() {
	_jsii_.InvokeVoid(
		s,
		"resetValidUserList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) ResetVpcEndpointDnsName() {
	_jsii_.InvokeVoid(
		s,
		"resetVpcEndpointDnsName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewaySmbFileShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewaySmbFileShareCacheAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#cache_stale_timeout_in_seconds StoragegatewaySmbFileShare#cache_stale_timeout_in_seconds}.
	CacheStaleTimeoutInSeconds *float64 `json:"cacheStaleTimeoutInSeconds"`
}

type StoragegatewaySmbFileShareCacheAttributesOutputReference interface {
	cdktf.ComplexObject
	CacheStaleTimeoutInSeconds() *float64
	SetCacheStaleTimeoutInSeconds(val *float64)
	CacheStaleTimeoutInSecondsInput() *float64
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
	ResetCacheStaleTimeoutInSeconds()
}

// The jsii proxy struct for StoragegatewaySmbFileShareCacheAttributesOutputReference
type jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) CacheStaleTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheStaleTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) CacheStaleTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheStaleTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewStoragegatewaySmbFileShareCacheAttributesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) StoragegatewaySmbFileShareCacheAttributesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewaySmbFileShareCacheAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewStoragegatewaySmbFileShareCacheAttributesOutputReference_Override(s StoragegatewaySmbFileShareCacheAttributesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewaySmbFileShareCacheAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) SetCacheStaleTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"cacheStaleTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewaySmbFileShareCacheAttributesOutputReference) ResetCacheStaleTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetCacheStaleTimeoutInSeconds",
		nil, // no parameters
	)
}

type StoragegatewaySmbFileShareConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#gateway_arn StoragegatewaySmbFileShare#gateway_arn}.
	GatewayArn *string `json:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#location_arn StoragegatewaySmbFileShare#location_arn}.
	LocationArn *string `json:"locationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#role_arn StoragegatewaySmbFileShare#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#access_based_enumeration StoragegatewaySmbFileShare#access_based_enumeration}.
	AccessBasedEnumeration interface{} `json:"accessBasedEnumeration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#admin_user_list StoragegatewaySmbFileShare#admin_user_list}.
	AdminUserList *[]*string `json:"adminUserList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#audit_destination_arn StoragegatewaySmbFileShare#audit_destination_arn}.
	AuditDestinationArn *string `json:"auditDestinationArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#authentication StoragegatewaySmbFileShare#authentication}.
	Authentication *string `json:"authentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#bucket_region StoragegatewaySmbFileShare#bucket_region}.
	BucketRegion *string `json:"bucketRegion"`
	// cache_attributes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#cache_attributes StoragegatewaySmbFileShare#cache_attributes}
	CacheAttributes *StoragegatewaySmbFileShareCacheAttributes `json:"cacheAttributes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#case_sensitivity StoragegatewaySmbFileShare#case_sensitivity}.
	CaseSensitivity *string `json:"caseSensitivity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#default_storage_class StoragegatewaySmbFileShare#default_storage_class}.
	DefaultStorageClass *string `json:"defaultStorageClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#file_share_name StoragegatewaySmbFileShare#file_share_name}.
	FileShareName *string `json:"fileShareName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#guess_mime_type_enabled StoragegatewaySmbFileShare#guess_mime_type_enabled}.
	GuessMimeTypeEnabled interface{} `json:"guessMimeTypeEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#invalid_user_list StoragegatewaySmbFileShare#invalid_user_list}.
	InvalidUserList *[]*string `json:"invalidUserList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#kms_encrypted StoragegatewaySmbFileShare#kms_encrypted}.
	KmsEncrypted interface{} `json:"kmsEncrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#kms_key_arn StoragegatewaySmbFileShare#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#notification_policy StoragegatewaySmbFileShare#notification_policy}.
	NotificationPolicy *string `json:"notificationPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#object_acl StoragegatewaySmbFileShare#object_acl}.
	ObjectAcl *string `json:"objectAcl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#oplocks_enabled StoragegatewaySmbFileShare#oplocks_enabled}.
	OplocksEnabled interface{} `json:"oplocksEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#read_only StoragegatewaySmbFileShare#read_only}.
	ReadOnly interface{} `json:"readOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#requester_pays StoragegatewaySmbFileShare#requester_pays}.
	RequesterPays interface{} `json:"requesterPays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#smb_acl_enabled StoragegatewaySmbFileShare#smb_acl_enabled}.
	SmbAclEnabled interface{} `json:"smbAclEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#tags StoragegatewaySmbFileShare#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#tags_all StoragegatewaySmbFileShare#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#timeouts StoragegatewaySmbFileShare#timeouts}
	Timeouts *StoragegatewaySmbFileShareTimeouts `json:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#valid_user_list StoragegatewaySmbFileShare#valid_user_list}.
	ValidUserList *[]*string `json:"validUserList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#vpc_endpoint_dns_name StoragegatewaySmbFileShare#vpc_endpoint_dns_name}.
	VpcEndpointDnsName *string `json:"vpcEndpointDnsName"`
}

type StoragegatewaySmbFileShareTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#create StoragegatewaySmbFileShare#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#delete StoragegatewaySmbFileShare#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_smb_file_share.html#update StoragegatewaySmbFileShare#update}.
	Update *string `json:"update"`
}

type StoragegatewaySmbFileShareTimeoutsOutputReference interface {
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

// The jsii proxy struct for StoragegatewaySmbFileShareTimeoutsOutputReference
type jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewStoragegatewaySmbFileShareTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) StoragegatewaySmbFileShareTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewaySmbFileShareTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewStoragegatewaySmbFileShareTimeoutsOutputReference_Override(s StoragegatewaySmbFileShareTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewaySmbFileShareTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewaySmbFileShareTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html aws_storagegateway_stored_iscsi_volume}.
type StoragegatewayStoredIscsiVolume interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ChapEnabled() interface{}
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DiskId() *string
	SetDiskId(val *string)
	DiskIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
	Id() *string
	KmsEncrypted() interface{}
	SetKmsEncrypted(val interface{})
	KmsEncryptedInput() interface{}
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LunNumber() *float64
	NetworkInterfaceId() *string
	SetNetworkInterfaceId(val *string)
	NetworkInterfaceIdInput() *string
	NetworkInterfacePort() *float64
	Node() constructs.Node
	PreserveExistingData() interface{}
	SetPreserveExistingData(val interface{})
	PreserveExistingDataInput() interface{}
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TargetArn() *string
	TargetName() *string
	SetTargetName(val *string)
	TargetNameInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VolumeAttachmentStatus() *string
	VolumeId() *string
	VolumeSizeInBytes() *float64
	VolumeStatus() *string
	VolumeType() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetKmsEncrypted()
	ResetKmsKey()
	ResetOverrideLogicalId()
	ResetSnapshotId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewayStoredIscsiVolume
type jsiiProxy_StoragegatewayStoredIscsiVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) ChapEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"chapEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) DiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) DiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) LunNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lunNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) NetworkInterfacePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkInterfacePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) PreserveExistingData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveExistingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) PreserveExistingDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveExistingDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TargetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TargetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeAttachmentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeAttachmentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeSizeInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html aws_storagegateway_stored_iscsi_volume} Resource.
func NewStoragegatewayStoredIscsiVolume(scope constructs.Construct, id *string, config *StoragegatewayStoredIscsiVolumeConfig) StoragegatewayStoredIscsiVolume {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayStoredIscsiVolume{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayStoredIscsiVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html aws_storagegateway_stored_iscsi_volume} Resource.
func NewStoragegatewayStoredIscsiVolume_Override(s StoragegatewayStoredIscsiVolume, scope constructs.Construct, id *string, config *StoragegatewayStoredIscsiVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayStoredIscsiVolume",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetDiskId(val *string) {
	_jsii_.Set(
		j,
		"diskId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetGatewayArn(val *string) {
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetKmsEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetKmsKey(val *string) {
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetNetworkInterfaceId(val *string) {
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetPreserveExistingData(val interface{}) {
	_jsii_.Set(
		j,
		"preserveExistingData",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetSnapshotId(val *string) {
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SetTargetName(val *string) {
	_jsii_.Set(
		j,
		"targetName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func StoragegatewayStoredIscsiVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewayStoredIscsiVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayStoredIscsiVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewayStoredIscsiVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetKmsKey() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKey",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		s,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewayStoredIscsiVolumeConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#disk_id StoragegatewayStoredIscsiVolume#disk_id}.
	DiskId *string `json:"diskId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#gateway_arn StoragegatewayStoredIscsiVolume#gateway_arn}.
	GatewayArn *string `json:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#network_interface_id StoragegatewayStoredIscsiVolume#network_interface_id}.
	NetworkInterfaceId *string `json:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#preserve_existing_data StoragegatewayStoredIscsiVolume#preserve_existing_data}.
	PreserveExistingData interface{} `json:"preserveExistingData"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#target_name StoragegatewayStoredIscsiVolume#target_name}.
	TargetName *string `json:"targetName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#kms_encrypted StoragegatewayStoredIscsiVolume#kms_encrypted}.
	KmsEncrypted interface{} `json:"kmsEncrypted"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#kms_key StoragegatewayStoredIscsiVolume#kms_key}.
	KmsKey *string `json:"kmsKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#snapshot_id StoragegatewayStoredIscsiVolume#snapshot_id}.
	SnapshotId *string `json:"snapshotId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#tags StoragegatewayStoredIscsiVolume#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume.html#tags_all StoragegatewayStoredIscsiVolume#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_tape_pool.html aws_storagegateway_tape_pool}.
type StoragegatewayTapePool interface {
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
	Node() constructs.Node
	PoolName() *string
	SetPoolName(val *string)
	PoolNameInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RetentionLockTimeInDays() *float64
	SetRetentionLockTimeInDays(val *float64)
	RetentionLockTimeInDaysInput() *float64
	RetentionLockType() *string
	SetRetentionLockType(val *string)
	RetentionLockTypeInput() *string
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
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
	ResetRetentionLockTimeInDays()
	ResetRetentionLockType()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewayTapePool
type jsiiProxy_StoragegatewayTapePool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayTapePool) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) PoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) PoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) RetentionLockTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionLockTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) RetentionLockTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionLockTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) RetentionLockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionLockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) RetentionLockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionLockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayTapePool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_tape_pool.html aws_storagegateway_tape_pool} Resource.
func NewStoragegatewayTapePool(scope constructs.Construct, id *string, config *StoragegatewayTapePoolConfig) StoragegatewayTapePool {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayTapePool{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayTapePool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_tape_pool.html aws_storagegateway_tape_pool} Resource.
func NewStoragegatewayTapePool_Override(s StoragegatewayTapePool, scope constructs.Construct, id *string, config *StoragegatewayTapePoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayTapePool",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetPoolName(val *string) {
	_jsii_.Set(
		j,
		"poolName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetRetentionLockTimeInDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionLockTimeInDays",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetRetentionLockType(val *string) {
	_jsii_.Set(
		j,
		"retentionLockType",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetStorageClass(val *string) {
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayTapePool) SetTagsAll(val interface{}) {
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
func StoragegatewayTapePool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewayTapePool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayTapePool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewayTapePool",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayTapePool) ResetRetentionLockTimeInDays() {
	_jsii_.InvokeVoid(
		s,
		"resetRetentionLockTimeInDays",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayTapePool) ResetRetentionLockType() {
	_jsii_.InvokeVoid(
		s,
		"resetRetentionLockType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayTapePool) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayTapePool) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayTapePool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewayTapePool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewayTapePool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewayTapePoolConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_tape_pool.html#pool_name StoragegatewayTapePool#pool_name}.
	PoolName *string `json:"poolName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_tape_pool.html#storage_class StoragegatewayTapePool#storage_class}.
	StorageClass *string `json:"storageClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_tape_pool.html#retention_lock_time_in_days StoragegatewayTapePool#retention_lock_time_in_days}.
	RetentionLockTimeInDays *float64 `json:"retentionLockTimeInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_tape_pool.html#retention_lock_type StoragegatewayTapePool#retention_lock_type}.
	RetentionLockType *string `json:"retentionLockType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_tape_pool.html#tags StoragegatewayTapePool#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_tape_pool.html#tags_all StoragegatewayTapePool#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_upload_buffer.html aws_storagegateway_upload_buffer}.
type StoragegatewayUploadBuffer interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DiskId() *string
	SetDiskId(val *string)
	DiskIdInput() *string
	DiskPath() *string
	SetDiskPath(val *string)
	DiskPathInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
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
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetDiskId()
	ResetDiskPath()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewayUploadBuffer
type jsiiProxy_StoragegatewayUploadBuffer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) DiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) DiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) DiskPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) DiskPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_upload_buffer.html aws_storagegateway_upload_buffer} Resource.
func NewStoragegatewayUploadBuffer(scope constructs.Construct, id *string, config *StoragegatewayUploadBufferConfig) StoragegatewayUploadBuffer {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayUploadBuffer{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayUploadBuffer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_upload_buffer.html aws_storagegateway_upload_buffer} Resource.
func NewStoragegatewayUploadBuffer_Override(s StoragegatewayUploadBuffer, scope constructs.Construct, id *string, config *StoragegatewayUploadBufferConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayUploadBuffer",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) SetDiskId(val *string) {
	_jsii_.Set(
		j,
		"diskId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) SetDiskPath(val *string) {
	_jsii_.Set(
		j,
		"diskPath",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) SetGatewayArn(val *string) {
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayUploadBuffer) SetProvider(val cdktf.TerraformProvider) {
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
func StoragegatewayUploadBuffer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewayUploadBuffer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayUploadBuffer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewayUploadBuffer",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayUploadBuffer) ResetDiskId() {
	_jsii_.InvokeVoid(
		s,
		"resetDiskId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayUploadBuffer) ResetDiskPath() {
	_jsii_.InvokeVoid(
		s,
		"resetDiskPath",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayUploadBuffer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewayUploadBuffer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewayUploadBuffer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewayUploadBufferConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_upload_buffer.html#gateway_arn StoragegatewayUploadBuffer#gateway_arn}.
	GatewayArn *string `json:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_upload_buffer.html#disk_id StoragegatewayUploadBuffer#disk_id}.
	DiskId *string `json:"diskId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_upload_buffer.html#disk_path StoragegatewayUploadBuffer#disk_path}.
	DiskPath *string `json:"diskPath"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_working_storage.html aws_storagegateway_working_storage}.
type StoragegatewayWorkingStorage interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DiskId() *string
	SetDiskId(val *string)
	DiskIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
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

// The jsii proxy struct for StoragegatewayWorkingStorage
type jsiiProxy_StoragegatewayWorkingStorage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) DiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) DiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_working_storage.html aws_storagegateway_working_storage} Resource.
func NewStoragegatewayWorkingStorage(scope constructs.Construct, id *string, config *StoragegatewayWorkingStorageConfig) StoragegatewayWorkingStorage {
	_init_.Initialize()

	j := jsiiProxy_StoragegatewayWorkingStorage{}

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayWorkingStorage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_working_storage.html aws_storagegateway_working_storage} Resource.
func NewStoragegatewayWorkingStorage_Override(s StoragegatewayWorkingStorage, scope constructs.Construct, id *string, config *StoragegatewayWorkingStorageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.StorageGateway.StoragegatewayWorkingStorage",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) SetDiskId(val *string) {
	_jsii_.Set(
		j,
		"diskId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) SetGatewayArn(val *string) {
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayWorkingStorage) SetProvider(val cdktf.TerraformProvider) {
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
func StoragegatewayWorkingStorage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.StorageGateway.StoragegatewayWorkingStorage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayWorkingStorage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.StorageGateway.StoragegatewayWorkingStorage",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayWorkingStorage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_StoragegatewayWorkingStorage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_StoragegatewayWorkingStorage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type StoragegatewayWorkingStorageConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_working_storage.html#disk_id StoragegatewayWorkingStorage#disk_id}.
	DiskId *string `json:"diskId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_working_storage.html#gateway_arn StoragegatewayWorkingStorage#gateway_arn}.
	GatewayArn *string `json:"gatewayArn"`
}
