package elastictranscoder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/elastictranscoder/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html aws_elastictranscoder_pipeline}.
type ElastictranscoderPipeline interface {
	cdktf.TerraformResource
	Arn() *string
	AwsKmsKeyArn() *string
	SetAwsKmsKeyArn(val *string)
	AwsKmsKeyArnInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	ContentConfig() ElastictranscoderPipelineContentConfigOutputReference
	ContentConfigInput() *ElastictranscoderPipelineContentConfig
	ContentConfigPermissions() *[]*ElastictranscoderPipelineContentConfigPermissions
	SetContentConfigPermissions(val *[]*ElastictranscoderPipelineContentConfigPermissions)
	ContentConfigPermissionsInput() *[]*ElastictranscoderPipelineContentConfigPermissions
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	InputBucket() *string
	SetInputBucket(val *string)
	InputBucketInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Notifications() ElastictranscoderPipelineNotificationsOutputReference
	NotificationsInput() *ElastictranscoderPipelineNotifications
	OutputBucket() *string
	SetOutputBucket(val *string)
	OutputBucketInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	ThumbnailConfig() ElastictranscoderPipelineThumbnailConfigOutputReference
	ThumbnailConfigInput() *ElastictranscoderPipelineThumbnailConfig
	ThumbnailConfigPermissions() *[]*ElastictranscoderPipelineThumbnailConfigPermissions
	SetThumbnailConfigPermissions(val *[]*ElastictranscoderPipelineThumbnailConfigPermissions)
	ThumbnailConfigPermissionsInput() *[]*ElastictranscoderPipelineThumbnailConfigPermissions
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutContentConfig(value *ElastictranscoderPipelineContentConfig)
	PutNotifications(value *ElastictranscoderPipelineNotifications)
	PutThumbnailConfig(value *ElastictranscoderPipelineThumbnailConfig)
	ResetAwsKmsKeyArn()
	ResetContentConfig()
	ResetContentConfigPermissions()
	ResetName()
	ResetNotifications()
	ResetOutputBucket()
	ResetOverrideLogicalId()
	ResetThumbnailConfig()
	ResetThumbnailConfigPermissions()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ElastictranscoderPipeline
type jsiiProxy_ElastictranscoderPipeline struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElastictranscoderPipeline) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) AwsKmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsKmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) AwsKmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsKmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ContentConfig() ElastictranscoderPipelineContentConfigOutputReference {
	var returns ElastictranscoderPipelineContentConfigOutputReference
	_jsii_.Get(
		j,
		"contentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ContentConfigInput() *ElastictranscoderPipelineContentConfig {
	var returns *ElastictranscoderPipelineContentConfig
	_jsii_.Get(
		j,
		"contentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ContentConfigPermissions() *[]*ElastictranscoderPipelineContentConfigPermissions {
	var returns *[]*ElastictranscoderPipelineContentConfigPermissions
	_jsii_.Get(
		j,
		"contentConfigPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ContentConfigPermissionsInput() *[]*ElastictranscoderPipelineContentConfigPermissions {
	var returns *[]*ElastictranscoderPipelineContentConfigPermissions
	_jsii_.Get(
		j,
		"contentConfigPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) InputBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) InputBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Notifications() ElastictranscoderPipelineNotificationsOutputReference {
	var returns ElastictranscoderPipelineNotificationsOutputReference
	_jsii_.Get(
		j,
		"notifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) NotificationsInput() *ElastictranscoderPipelineNotifications {
	var returns *ElastictranscoderPipelineNotifications
	_jsii_.Get(
		j,
		"notificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) OutputBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) OutputBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ThumbnailConfig() ElastictranscoderPipelineThumbnailConfigOutputReference {
	var returns ElastictranscoderPipelineThumbnailConfigOutputReference
	_jsii_.Get(
		j,
		"thumbnailConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ThumbnailConfigInput() *ElastictranscoderPipelineThumbnailConfig {
	var returns *ElastictranscoderPipelineThumbnailConfig
	_jsii_.Get(
		j,
		"thumbnailConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ThumbnailConfigPermissions() *[]*ElastictranscoderPipelineThumbnailConfigPermissions {
	var returns *[]*ElastictranscoderPipelineThumbnailConfigPermissions
	_jsii_.Get(
		j,
		"thumbnailConfigPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipeline) ThumbnailConfigPermissionsInput() *[]*ElastictranscoderPipelineThumbnailConfigPermissions {
	var returns *[]*ElastictranscoderPipelineThumbnailConfigPermissions
	_jsii_.Get(
		j,
		"thumbnailConfigPermissionsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html aws_elastictranscoder_pipeline} Resource.
func NewElastictranscoderPipeline(scope constructs.Construct, id *string, config *ElastictranscoderPipelineConfig) ElastictranscoderPipeline {
	_init_.Initialize()

	j := jsiiProxy_ElastictranscoderPipeline{}

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipeline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html aws_elastictranscoder_pipeline} Resource.
func NewElastictranscoderPipeline_Override(e ElastictranscoderPipeline, scope constructs.Construct, id *string, config *ElastictranscoderPipelineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipeline",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetAwsKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"awsKmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetContentConfigPermissions(val *[]*ElastictranscoderPipelineContentConfigPermissions) {
	_jsii_.Set(
		j,
		"contentConfigPermissions",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetInputBucket(val *string) {
	_jsii_.Set(
		j,
		"inputBucket",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetOutputBucket(val *string) {
	_jsii_.Set(
		j,
		"outputBucket",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipeline) SetThumbnailConfigPermissions(val *[]*ElastictranscoderPipelineThumbnailConfigPermissions) {
	_jsii_.Set(
		j,
		"thumbnailConfigPermissions",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ElastictranscoderPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElastictranscoderPipeline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipeline",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) PutContentConfig(value *ElastictranscoderPipelineContentConfig) {
	_jsii_.InvokeVoid(
		e,
		"putContentConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) PutNotifications(value *ElastictranscoderPipelineNotifications) {
	_jsii_.InvokeVoid(
		e,
		"putNotifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) PutThumbnailConfig(value *ElastictranscoderPipelineThumbnailConfig) {
	_jsii_.InvokeVoid(
		e,
		"putThumbnailConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetAwsKmsKeyArn() {
	_jsii_.InvokeVoid(
		e,
		"resetAwsKmsKeyArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetContentConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetContentConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetContentConfigPermissions() {
	_jsii_.InvokeVoid(
		e,
		"resetContentConfigPermissions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetNotifications() {
	_jsii_.InvokeVoid(
		e,
		"resetNotifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetOutputBucket() {
	_jsii_.InvokeVoid(
		e,
		"resetOutputBucket",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetThumbnailConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetThumbnailConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) ResetThumbnailConfigPermissions() {
	_jsii_.InvokeVoid(
		e,
		"resetThumbnailConfigPermissions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipeline) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_ElastictranscoderPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_ElastictranscoderPipeline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ElastictranscoderPipelineConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#input_bucket ElastictranscoderPipeline#input_bucket}.
	InputBucket *string `json:"inputBucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#role ElastictranscoderPipeline#role}.
	Role *string `json:"role"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#aws_kms_key_arn ElastictranscoderPipeline#aws_kms_key_arn}.
	AwsKmsKeyArn *string `json:"awsKmsKeyArn"`
	// content_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#content_config ElastictranscoderPipeline#content_config}
	ContentConfig *ElastictranscoderPipelineContentConfig `json:"contentConfig"`
	// content_config_permissions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#content_config_permissions ElastictranscoderPipeline#content_config_permissions}
	ContentConfigPermissions *[]*ElastictranscoderPipelineContentConfigPermissions `json:"contentConfigPermissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#name ElastictranscoderPipeline#name}.
	Name *string `json:"name"`
	// notifications block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#notifications ElastictranscoderPipeline#notifications}
	Notifications *ElastictranscoderPipelineNotifications `json:"notifications"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#output_bucket ElastictranscoderPipeline#output_bucket}.
	OutputBucket *string `json:"outputBucket"`
	// thumbnail_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#thumbnail_config ElastictranscoderPipeline#thumbnail_config}
	ThumbnailConfig *ElastictranscoderPipelineThumbnailConfig `json:"thumbnailConfig"`
	// thumbnail_config_permissions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#thumbnail_config_permissions ElastictranscoderPipeline#thumbnail_config_permissions}
	ThumbnailConfigPermissions *[]*ElastictranscoderPipelineThumbnailConfigPermissions `json:"thumbnailConfigPermissions"`
}

type ElastictranscoderPipelineContentConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#bucket ElastictranscoderPipeline#bucket}.
	Bucket *string `json:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#storage_class ElastictranscoderPipeline#storage_class}.
	StorageClass *string `json:"storageClass"`
}

type ElastictranscoderPipelineContentConfigOutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
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
	ResetBucket()
	ResetStorageClass()
}

// The jsii proxy struct for ElastictranscoderPipelineContentConfigOutputReference
type jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElastictranscoderPipelineContentConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElastictranscoderPipelineContentConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineContentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElastictranscoderPipelineContentConfigOutputReference_Override(e ElastictranscoderPipelineContentConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineContentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) SetStorageClass(val *string) {
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) ResetBucket() {
	_jsii_.InvokeVoid(
		e,
		"resetBucket",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageClass",
		nil, // no parameters
	)
}

type ElastictranscoderPipelineContentConfigPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#access ElastictranscoderPipeline#access}.
	Access *[]*string `json:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#grantee ElastictranscoderPipeline#grantee}.
	Grantee *string `json:"grantee"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#grantee_type ElastictranscoderPipeline#grantee_type}.
	GranteeType *string `json:"granteeType"`
}

type ElastictranscoderPipelineNotifications struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#completed ElastictranscoderPipeline#completed}.
	Completed *string `json:"completed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#error ElastictranscoderPipeline#error}.
	Error *string `json:"error"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#progressing ElastictranscoderPipeline#progressing}.
	Progressing *string `json:"progressing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#warning ElastictranscoderPipeline#warning}.
	Warning *string `json:"warning"`
}

type ElastictranscoderPipelineNotificationsOutputReference interface {
	cdktf.ComplexObject
	Completed() *string
	SetCompleted(val *string)
	CompletedInput() *string
	Error() *string
	SetError(val *string)
	ErrorInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Progressing() *string
	SetProgressing(val *string)
	ProgressingInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Warning() *string
	SetWarning(val *string)
	WarningInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCompleted()
	ResetError()
	ResetProgressing()
	ResetWarning()
}

// The jsii proxy struct for ElastictranscoderPipelineNotificationsOutputReference
type jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) Completed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) CompletedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) Error() *string {
	var returns *string
	_jsii_.Get(
		j,
		"error",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) ErrorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) Progressing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"progressing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) ProgressingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"progressingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) Warning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) WarningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}

func NewElastictranscoderPipelineNotificationsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElastictranscoderPipelineNotificationsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElastictranscoderPipelineNotificationsOutputReference_Override(e ElastictranscoderPipelineNotificationsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineNotificationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) SetCompleted(val *string) {
	_jsii_.Set(
		j,
		"completed",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) SetError(val *string) {
	_jsii_.Set(
		j,
		"error",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) SetProgressing(val *string) {
	_jsii_.Set(
		j,
		"progressing",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) SetWarning(val *string) {
	_jsii_.Set(
		j,
		"warning",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) ResetCompleted() {
	_jsii_.InvokeVoid(
		e,
		"resetCompleted",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) ResetError() {
	_jsii_.InvokeVoid(
		e,
		"resetError",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) ResetProgressing() {
	_jsii_.InvokeVoid(
		e,
		"resetProgressing",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference) ResetWarning() {
	_jsii_.InvokeVoid(
		e,
		"resetWarning",
		nil, // no parameters
	)
}

type ElastictranscoderPipelineThumbnailConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#bucket ElastictranscoderPipeline#bucket}.
	Bucket *string `json:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#storage_class ElastictranscoderPipeline#storage_class}.
	StorageClass *string `json:"storageClass"`
}

type ElastictranscoderPipelineThumbnailConfigOutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
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
	ResetBucket()
	ResetStorageClass()
}

// The jsii proxy struct for ElastictranscoderPipelineThumbnailConfigOutputReference
type jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElastictranscoderPipelineThumbnailConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElastictranscoderPipelineThumbnailConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineThumbnailConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElastictranscoderPipelineThumbnailConfigOutputReference_Override(e ElastictranscoderPipelineThumbnailConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineThumbnailConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) SetStorageClass(val *string) {
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) ResetBucket() {
	_jsii_.InvokeVoid(
		e,
		"resetBucket",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageClass",
		nil, // no parameters
	)
}

type ElastictranscoderPipelineThumbnailConfigPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#access ElastictranscoderPipeline#access}.
	Access *[]*string `json:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#grantee ElastictranscoderPipeline#grantee}.
	Grantee *string `json:"grantee"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_pipeline.html#grantee_type ElastictranscoderPipeline#grantee_type}.
	GranteeType *string `json:"granteeType"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html aws_elastictranscoder_preset}.
type ElastictranscoderPreset interface {
	cdktf.TerraformResource
	Arn() *string
	Audio() ElastictranscoderPresetAudioOutputReference
	AudioCodecOptions() ElastictranscoderPresetAudioCodecOptionsOutputReference
	AudioCodecOptionsInput() *ElastictranscoderPresetAudioCodecOptions
	AudioInput() *ElastictranscoderPresetAudio
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Container() *string
	SetContainer(val *string)
	ContainerInput() *string
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
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Thumbnails() ElastictranscoderPresetThumbnailsOutputReference
	ThumbnailsInput() *ElastictranscoderPresetThumbnails
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Video() ElastictranscoderPresetVideoOutputReference
	VideoCodecOptions() interface{}
	SetVideoCodecOptions(val interface{})
	VideoCodecOptionsInput() interface{}
	VideoInput() *ElastictranscoderPresetVideo
	VideoWatermarks() *[]*ElastictranscoderPresetVideoWatermarks
	SetVideoWatermarks(val *[]*ElastictranscoderPresetVideoWatermarks)
	VideoWatermarksInput() *[]*ElastictranscoderPresetVideoWatermarks
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutAudio(value *ElastictranscoderPresetAudio)
	PutAudioCodecOptions(value *ElastictranscoderPresetAudioCodecOptions)
	PutThumbnails(value *ElastictranscoderPresetThumbnails)
	PutVideo(value *ElastictranscoderPresetVideo)
	ResetAudio()
	ResetAudioCodecOptions()
	ResetDescription()
	ResetName()
	ResetOverrideLogicalId()
	ResetThumbnails()
	ResetType()
	ResetVideo()
	ResetVideoCodecOptions()
	ResetVideoWatermarks()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ElastictranscoderPreset
type jsiiProxy_ElastictranscoderPreset struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElastictranscoderPreset) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Audio() ElastictranscoderPresetAudioOutputReference {
	var returns ElastictranscoderPresetAudioOutputReference
	_jsii_.Get(
		j,
		"audio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) AudioCodecOptions() ElastictranscoderPresetAudioCodecOptionsOutputReference {
	var returns ElastictranscoderPresetAudioCodecOptionsOutputReference
	_jsii_.Get(
		j,
		"audioCodecOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) AudioCodecOptionsInput() *ElastictranscoderPresetAudioCodecOptions {
	var returns *ElastictranscoderPresetAudioCodecOptions
	_jsii_.Get(
		j,
		"audioCodecOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) AudioInput() *ElastictranscoderPresetAudio {
	var returns *ElastictranscoderPresetAudio
	_jsii_.Get(
		j,
		"audioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Container() *string {
	var returns *string
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) ContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Thumbnails() ElastictranscoderPresetThumbnailsOutputReference {
	var returns ElastictranscoderPresetThumbnailsOutputReference
	_jsii_.Get(
		j,
		"thumbnails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) ThumbnailsInput() *ElastictranscoderPresetThumbnails {
	var returns *ElastictranscoderPresetThumbnails
	_jsii_.Get(
		j,
		"thumbnailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) Video() ElastictranscoderPresetVideoOutputReference {
	var returns ElastictranscoderPresetVideoOutputReference
	_jsii_.Get(
		j,
		"video",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) VideoCodecOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoCodecOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) VideoCodecOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoCodecOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) VideoInput() *ElastictranscoderPresetVideo {
	var returns *ElastictranscoderPresetVideo
	_jsii_.Get(
		j,
		"videoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) VideoWatermarks() *[]*ElastictranscoderPresetVideoWatermarks {
	var returns *[]*ElastictranscoderPresetVideoWatermarks
	_jsii_.Get(
		j,
		"videoWatermarks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPreset) VideoWatermarksInput() *[]*ElastictranscoderPresetVideoWatermarks {
	var returns *[]*ElastictranscoderPresetVideoWatermarks
	_jsii_.Get(
		j,
		"videoWatermarksInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html aws_elastictranscoder_preset} Resource.
func NewElastictranscoderPreset(scope constructs.Construct, id *string, config *ElastictranscoderPresetConfig) ElastictranscoderPreset {
	_init_.Initialize()

	j := jsiiProxy_ElastictranscoderPreset{}

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPreset",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html aws_elastictranscoder_preset} Resource.
func NewElastictranscoderPreset_Override(e ElastictranscoderPreset, scope constructs.Construct, id *string, config *ElastictranscoderPresetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPreset",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetContainer(val *string) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetVideoCodecOptions(val interface{}) {
	_jsii_.Set(
		j,
		"videoCodecOptions",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPreset) SetVideoWatermarks(val *[]*ElastictranscoderPresetVideoWatermarks) {
	_jsii_.Set(
		j,
		"videoWatermarks",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ElastictranscoderPreset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPreset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElastictranscoderPreset_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPreset",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) PutAudio(value *ElastictranscoderPresetAudio) {
	_jsii_.InvokeVoid(
		e,
		"putAudio",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) PutAudioCodecOptions(value *ElastictranscoderPresetAudioCodecOptions) {
	_jsii_.InvokeVoid(
		e,
		"putAudioCodecOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) PutThumbnails(value *ElastictranscoderPresetThumbnails) {
	_jsii_.InvokeVoid(
		e,
		"putThumbnails",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) PutVideo(value *ElastictranscoderPresetVideo) {
	_jsii_.InvokeVoid(
		e,
		"putVideo",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) ResetAudio() {
	_jsii_.InvokeVoid(
		e,
		"resetAudio",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) ResetAudioCodecOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetAudioCodecOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) ResetThumbnails() {
	_jsii_.InvokeVoid(
		e,
		"resetThumbnails",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) ResetVideo() {
	_jsii_.InvokeVoid(
		e,
		"resetVideo",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) ResetVideoCodecOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetVideoCodecOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) ResetVideoWatermarks() {
	_jsii_.InvokeVoid(
		e,
		"resetVideoWatermarks",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPreset) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_ElastictranscoderPreset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_ElastictranscoderPreset) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ElastictranscoderPresetAudio struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#audio_packing_mode ElastictranscoderPreset#audio_packing_mode}.
	AudioPackingMode *string `json:"audioPackingMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#bit_rate ElastictranscoderPreset#bit_rate}.
	BitRate *string `json:"bitRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#channels ElastictranscoderPreset#channels}.
	Channels *string `json:"channels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#codec ElastictranscoderPreset#codec}.
	Codec *string `json:"codec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#sample_rate ElastictranscoderPreset#sample_rate}.
	SampleRate *string `json:"sampleRate"`
}

type ElastictranscoderPresetAudioCodecOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#bit_depth ElastictranscoderPreset#bit_depth}.
	BitDepth *string `json:"bitDepth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#bit_order ElastictranscoderPreset#bit_order}.
	BitOrder *string `json:"bitOrder"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#profile ElastictranscoderPreset#profile}.
	Profile *string `json:"profile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#signed ElastictranscoderPreset#signed}.
	Signed *string `json:"signed"`
}

type ElastictranscoderPresetAudioCodecOptionsOutputReference interface {
	cdktf.ComplexObject
	BitDepth() *string
	SetBitDepth(val *string)
	BitDepthInput() *string
	BitOrder() *string
	SetBitOrder(val *string)
	BitOrderInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	Signed() *string
	SetSigned(val *string)
	SignedInput() *string
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
	ResetBitDepth()
	ResetBitOrder()
	ResetProfile()
	ResetSigned()
}

// The jsii proxy struct for ElastictranscoderPresetAudioCodecOptionsOutputReference
type jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) BitDepth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) BitDepthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) BitOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) BitOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) Signed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) SignedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElastictranscoderPresetAudioCodecOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElastictranscoderPresetAudioCodecOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetAudioCodecOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElastictranscoderPresetAudioCodecOptionsOutputReference_Override(e ElastictranscoderPresetAudioCodecOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetAudioCodecOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) SetBitDepth(val *string) {
	_jsii_.Set(
		j,
		"bitDepth",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) SetBitOrder(val *string) {
	_jsii_.Set(
		j,
		"bitOrder",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) SetProfile(val *string) {
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) SetSigned(val *string) {
	_jsii_.Set(
		j,
		"signed",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ResetBitDepth() {
	_jsii_.InvokeVoid(
		e,
		"resetBitDepth",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ResetBitOrder() {
	_jsii_.InvokeVoid(
		e,
		"resetBitOrder",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		e,
		"resetProfile",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference) ResetSigned() {
	_jsii_.InvokeVoid(
		e,
		"resetSigned",
		nil, // no parameters
	)
}

type ElastictranscoderPresetAudioOutputReference interface {
	cdktf.ComplexObject
	AudioPackingMode() *string
	SetAudioPackingMode(val *string)
	AudioPackingModeInput() *string
	BitRate() *string
	SetBitRate(val *string)
	BitRateInput() *string
	Channels() *string
	SetChannels(val *string)
	ChannelsInput() *string
	Codec() *string
	SetCodec(val *string)
	CodecInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SampleRate() *string
	SetSampleRate(val *string)
	SampleRateInput() *string
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
	ResetAudioPackingMode()
	ResetBitRate()
	ResetChannels()
	ResetCodec()
	ResetSampleRate()
}

// The jsii proxy struct for ElastictranscoderPresetAudioOutputReference
type jsiiProxy_ElastictranscoderPresetAudioOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) AudioPackingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioPackingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) AudioPackingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioPackingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) BitRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) BitRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) Channels() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) ChannelsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) Codec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) CodecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SampleRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SampleRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElastictranscoderPresetAudioOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElastictranscoderPresetAudioOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastictranscoderPresetAudioOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetAudioOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElastictranscoderPresetAudioOutputReference_Override(e ElastictranscoderPresetAudioOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetAudioOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SetAudioPackingMode(val *string) {
	_jsii_.Set(
		j,
		"audioPackingMode",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SetBitRate(val *string) {
	_jsii_.Set(
		j,
		"bitRate",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SetChannels(val *string) {
	_jsii_.Set(
		j,
		"channels",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SetCodec(val *string) {
	_jsii_.Set(
		j,
		"codec",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SetSampleRate(val *string) {
	_jsii_.Set(
		j,
		"sampleRate",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetAudioOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) ResetAudioPackingMode() {
	_jsii_.InvokeVoid(
		e,
		"resetAudioPackingMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) ResetBitRate() {
	_jsii_.InvokeVoid(
		e,
		"resetBitRate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) ResetChannels() {
	_jsii_.InvokeVoid(
		e,
		"resetChannels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) ResetCodec() {
	_jsii_.InvokeVoid(
		e,
		"resetCodec",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetAudioOutputReference) ResetSampleRate() {
	_jsii_.InvokeVoid(
		e,
		"resetSampleRate",
		nil, // no parameters
	)
}

type ElastictranscoderPresetConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#container ElastictranscoderPreset#container}.
	Container *string `json:"container"`
	// audio block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#audio ElastictranscoderPreset#audio}
	Audio *ElastictranscoderPresetAudio `json:"audio"`
	// audio_codec_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#audio_codec_options ElastictranscoderPreset#audio_codec_options}
	AudioCodecOptions *ElastictranscoderPresetAudioCodecOptions `json:"audioCodecOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#description ElastictranscoderPreset#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#name ElastictranscoderPreset#name}.
	Name *string `json:"name"`
	// thumbnails block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#thumbnails ElastictranscoderPreset#thumbnails}
	Thumbnails *ElastictranscoderPresetThumbnails `json:"thumbnails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#type ElastictranscoderPreset#type}.
	Type *string `json:"type"`
	// video block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#video ElastictranscoderPreset#video}
	Video *ElastictranscoderPresetVideo `json:"video"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#video_codec_options ElastictranscoderPreset#video_codec_options}.
	VideoCodecOptions interface{} `json:"videoCodecOptions"`
	// video_watermarks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#video_watermarks ElastictranscoderPreset#video_watermarks}
	VideoWatermarks *[]*ElastictranscoderPresetVideoWatermarks `json:"videoWatermarks"`
}

type ElastictranscoderPresetThumbnails struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#aspect_ratio ElastictranscoderPreset#aspect_ratio}.
	AspectRatio *string `json:"aspectRatio"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#format ElastictranscoderPreset#format}.
	Format *string `json:"format"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#interval ElastictranscoderPreset#interval}.
	Interval *string `json:"interval"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#max_height ElastictranscoderPreset#max_height}.
	MaxHeight *string `json:"maxHeight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#max_width ElastictranscoderPreset#max_width}.
	MaxWidth *string `json:"maxWidth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#padding_policy ElastictranscoderPreset#padding_policy}.
	PaddingPolicy *string `json:"paddingPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#resolution ElastictranscoderPreset#resolution}.
	Resolution *string `json:"resolution"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#sizing_policy ElastictranscoderPreset#sizing_policy}.
	SizingPolicy *string `json:"sizingPolicy"`
}

type ElastictranscoderPresetThumbnailsOutputReference interface {
	cdktf.ComplexObject
	AspectRatio() *string
	SetAspectRatio(val *string)
	AspectRatioInput() *string
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	Interval() *string
	SetInterval(val *string)
	IntervalInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxHeight() *string
	SetMaxHeight(val *string)
	MaxHeightInput() *string
	MaxWidth() *string
	SetMaxWidth(val *string)
	MaxWidthInput() *string
	PaddingPolicy() *string
	SetPaddingPolicy(val *string)
	PaddingPolicyInput() *string
	Resolution() *string
	SetResolution(val *string)
	ResolutionInput() *string
	SizingPolicy() *string
	SetSizingPolicy(val *string)
	SizingPolicyInput() *string
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
	ResetAspectRatio()
	ResetFormat()
	ResetInterval()
	ResetMaxHeight()
	ResetMaxWidth()
	ResetPaddingPolicy()
	ResetResolution()
	ResetSizingPolicy()
}

// The jsii proxy struct for ElastictranscoderPresetThumbnailsOutputReference
type jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) AspectRatio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) AspectRatioInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) Interval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) IntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) MaxHeight() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) MaxHeightInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) MaxWidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) MaxWidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) PaddingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) PaddingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) Resolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) ResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SizingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SizingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElastictranscoderPresetThumbnailsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElastictranscoderPresetThumbnailsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetThumbnailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElastictranscoderPresetThumbnailsOutputReference_Override(e ElastictranscoderPresetThumbnailsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetThumbnailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetAspectRatio(val *string) {
	_jsii_.Set(
		j,
		"aspectRatio",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetFormat(val *string) {
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetInterval(val *string) {
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetMaxHeight(val *string) {
	_jsii_.Set(
		j,
		"maxHeight",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetMaxWidth(val *string) {
	_jsii_.Set(
		j,
		"maxWidth",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetPaddingPolicy(val *string) {
	_jsii_.Set(
		j,
		"paddingPolicy",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetResolution(val *string) {
	_jsii_.Set(
		j,
		"resolution",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetSizingPolicy(val *string) {
	_jsii_.Set(
		j,
		"sizingPolicy",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) ResetAspectRatio() {
	_jsii_.InvokeVoid(
		e,
		"resetAspectRatio",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) ResetFormat() {
	_jsii_.InvokeVoid(
		e,
		"resetFormat",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		e,
		"resetInterval",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) ResetMaxHeight() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxHeight",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) ResetMaxWidth() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxWidth",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) ResetPaddingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetPaddingPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) ResetResolution() {
	_jsii_.InvokeVoid(
		e,
		"resetResolution",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference) ResetSizingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetSizingPolicy",
		nil, // no parameters
	)
}

type ElastictranscoderPresetVideo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#aspect_ratio ElastictranscoderPreset#aspect_ratio}.
	AspectRatio *string `json:"aspectRatio"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#bit_rate ElastictranscoderPreset#bit_rate}.
	BitRate *string `json:"bitRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#codec ElastictranscoderPreset#codec}.
	Codec *string `json:"codec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#display_aspect_ratio ElastictranscoderPreset#display_aspect_ratio}.
	DisplayAspectRatio *string `json:"displayAspectRatio"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#fixed_gop ElastictranscoderPreset#fixed_gop}.
	FixedGop *string `json:"fixedGop"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#frame_rate ElastictranscoderPreset#frame_rate}.
	FrameRate *string `json:"frameRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#keyframes_max_dist ElastictranscoderPreset#keyframes_max_dist}.
	KeyframesMaxDist *string `json:"keyframesMaxDist"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#max_frame_rate ElastictranscoderPreset#max_frame_rate}.
	MaxFrameRate *string `json:"maxFrameRate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#max_height ElastictranscoderPreset#max_height}.
	MaxHeight *string `json:"maxHeight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#max_width ElastictranscoderPreset#max_width}.
	MaxWidth *string `json:"maxWidth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#padding_policy ElastictranscoderPreset#padding_policy}.
	PaddingPolicy *string `json:"paddingPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#resolution ElastictranscoderPreset#resolution}.
	Resolution *string `json:"resolution"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#sizing_policy ElastictranscoderPreset#sizing_policy}.
	SizingPolicy *string `json:"sizingPolicy"`
}

type ElastictranscoderPresetVideoOutputReference interface {
	cdktf.ComplexObject
	AspectRatio() *string
	SetAspectRatio(val *string)
	AspectRatioInput() *string
	BitRate() *string
	SetBitRate(val *string)
	BitRateInput() *string
	Codec() *string
	SetCodec(val *string)
	CodecInput() *string
	DisplayAspectRatio() *string
	SetDisplayAspectRatio(val *string)
	DisplayAspectRatioInput() *string
	FixedGop() *string
	SetFixedGop(val *string)
	FixedGopInput() *string
	FrameRate() *string
	SetFrameRate(val *string)
	FrameRateInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KeyframesMaxDist() *string
	SetKeyframesMaxDist(val *string)
	KeyframesMaxDistInput() *string
	MaxFrameRate() *string
	SetMaxFrameRate(val *string)
	MaxFrameRateInput() *string
	MaxHeight() *string
	SetMaxHeight(val *string)
	MaxHeightInput() *string
	MaxWidth() *string
	SetMaxWidth(val *string)
	MaxWidthInput() *string
	PaddingPolicy() *string
	SetPaddingPolicy(val *string)
	PaddingPolicyInput() *string
	Resolution() *string
	SetResolution(val *string)
	ResolutionInput() *string
	SizingPolicy() *string
	SetSizingPolicy(val *string)
	SizingPolicyInput() *string
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
	ResetAspectRatio()
	ResetBitRate()
	ResetCodec()
	ResetDisplayAspectRatio()
	ResetFixedGop()
	ResetFrameRate()
	ResetKeyframesMaxDist()
	ResetMaxFrameRate()
	ResetMaxHeight()
	ResetMaxWidth()
	ResetPaddingPolicy()
	ResetResolution()
	ResetSizingPolicy()
}

// The jsii proxy struct for ElastictranscoderPresetVideoOutputReference
type jsiiProxy_ElastictranscoderPresetVideoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) AspectRatio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) AspectRatioInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aspectRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) BitRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) BitRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) Codec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) CodecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) DisplayAspectRatio() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayAspectRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) DisplayAspectRatioInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayAspectRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) FixedGop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedGop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) FixedGopInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fixedGopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) FrameRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) FrameRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) KeyframesMaxDist() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyframesMaxDist",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) KeyframesMaxDistInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyframesMaxDistInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxFrameRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxFrameRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxFrameRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxFrameRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxHeight() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxHeightInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxHeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxWidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) MaxWidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxWidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) PaddingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) PaddingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paddingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) Resolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SizingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SizingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElastictranscoderPresetVideoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElastictranscoderPresetVideoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElastictranscoderPresetVideoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetVideoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElastictranscoderPresetVideoOutputReference_Override(e ElastictranscoderPresetVideoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetVideoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetAspectRatio(val *string) {
	_jsii_.Set(
		j,
		"aspectRatio",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetBitRate(val *string) {
	_jsii_.Set(
		j,
		"bitRate",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetCodec(val *string) {
	_jsii_.Set(
		j,
		"codec",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetDisplayAspectRatio(val *string) {
	_jsii_.Set(
		j,
		"displayAspectRatio",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetFixedGop(val *string) {
	_jsii_.Set(
		j,
		"fixedGop",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetFrameRate(val *string) {
	_jsii_.Set(
		j,
		"frameRate",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetKeyframesMaxDist(val *string) {
	_jsii_.Set(
		j,
		"keyframesMaxDist",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetMaxFrameRate(val *string) {
	_jsii_.Set(
		j,
		"maxFrameRate",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetMaxHeight(val *string) {
	_jsii_.Set(
		j,
		"maxHeight",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetMaxWidth(val *string) {
	_jsii_.Set(
		j,
		"maxWidth",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetPaddingPolicy(val *string) {
	_jsii_.Set(
		j,
		"paddingPolicy",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetResolution(val *string) {
	_jsii_.Set(
		j,
		"resolution",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetSizingPolicy(val *string) {
	_jsii_.Set(
		j,
		"sizingPolicy",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElastictranscoderPresetVideoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetAspectRatio() {
	_jsii_.InvokeVoid(
		e,
		"resetAspectRatio",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetBitRate() {
	_jsii_.InvokeVoid(
		e,
		"resetBitRate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetCodec() {
	_jsii_.InvokeVoid(
		e,
		"resetCodec",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetDisplayAspectRatio() {
	_jsii_.InvokeVoid(
		e,
		"resetDisplayAspectRatio",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetFixedGop() {
	_jsii_.InvokeVoid(
		e,
		"resetFixedGop",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetFrameRate() {
	_jsii_.InvokeVoid(
		e,
		"resetFrameRate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetKeyframesMaxDist() {
	_jsii_.InvokeVoid(
		e,
		"resetKeyframesMaxDist",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetMaxFrameRate() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxFrameRate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetMaxHeight() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxHeight",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetMaxWidth() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxWidth",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetPaddingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetPaddingPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetResolution() {
	_jsii_.InvokeVoid(
		e,
		"resetResolution",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElastictranscoderPresetVideoOutputReference) ResetSizingPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetSizingPolicy",
		nil, // no parameters
	)
}

type ElastictranscoderPresetVideoWatermarks struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#horizontal_align ElastictranscoderPreset#horizontal_align}.
	HorizontalAlign *string `json:"horizontalAlign"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#horizontal_offset ElastictranscoderPreset#horizontal_offset}.
	HorizontalOffset *string `json:"horizontalOffset"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#id ElastictranscoderPreset#id}.
	Id *string `json:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#max_height ElastictranscoderPreset#max_height}.
	MaxHeight *string `json:"maxHeight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#max_width ElastictranscoderPreset#max_width}.
	MaxWidth *string `json:"maxWidth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#opacity ElastictranscoderPreset#opacity}.
	Opacity *string `json:"opacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#sizing_policy ElastictranscoderPreset#sizing_policy}.
	SizingPolicy *string `json:"sizingPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#target ElastictranscoderPreset#target}.
	Target *string `json:"target"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#vertical_align ElastictranscoderPreset#vertical_align}.
	VerticalAlign *string `json:"verticalAlign"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elastictranscoder_preset.html#vertical_offset ElastictranscoderPreset#vertical_offset}.
	VerticalOffset *string `json:"verticalOffset"`
}
