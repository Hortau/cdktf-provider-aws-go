package codebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/codebuild/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project aws_codebuild_project}.
type CodebuildProject interface {
	cdktf.TerraformResource
	Arn() *string
	Artifacts() CodebuildProjectArtifactsOutputReference
	ArtifactsInput() *CodebuildProjectArtifacts
	BadgeEnabled() interface{}
	SetBadgeEnabled(val interface{})
	BadgeEnabledInput() interface{}
	BadgeUrl() *string
	BuildBatchConfig() CodebuildProjectBuildBatchConfigOutputReference
	BuildBatchConfigInput() *CodebuildProjectBuildBatchConfig
	BuildTimeout() *float64
	SetBuildTimeout(val *float64)
	BuildTimeoutInput() *float64
	Cache() CodebuildProjectCacheOutputReference
	CacheInput() *CodebuildProjectCache
	CdktfStack() cdktf.TerraformStack
	ConcurrentBuildLimit() *float64
	SetConcurrentBuildLimit(val *float64)
	ConcurrentBuildLimitInput() *float64
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	EncryptionKeyInput() *string
	Environment() CodebuildProjectEnvironmentOutputReference
	EnvironmentInput() *CodebuildProjectEnvironment
	FileSystemLocations() interface{}
	SetFileSystemLocations(val interface{})
	FileSystemLocationsInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogsConfig() CodebuildProjectLogsConfigOutputReference
	LogsConfigInput() *CodebuildProjectLogsConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	ProjectVisibility() *string
	SetProjectVisibility(val *string)
	ProjectVisibilityInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	PublicProjectAlias() *string
	QueuedTimeout() *float64
	SetQueuedTimeout(val *float64)
	QueuedTimeoutInput() *float64
	RawOverrides() interface{}
	ResourceAccessRole() *string
	SetResourceAccessRole(val *string)
	ResourceAccessRoleInput() *string
	SecondaryArtifacts() interface{}
	SetSecondaryArtifacts(val interface{})
	SecondaryArtifactsInput() interface{}
	SecondarySources() interface{}
	SetSecondarySources(val interface{})
	SecondarySourcesInput() interface{}
	SecondarySourceVersion() interface{}
	SetSecondarySourceVersion(val interface{})
	SecondarySourceVersionInput() interface{}
	ServiceRole() *string
	SetServiceRole(val *string)
	ServiceRoleInput() *string
	Source() CodebuildProjectSourceOutputReference
	SourceInput() *CodebuildProjectSource
	SourceVersion() *string
	SetSourceVersion(val *string)
	SourceVersionInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VpcConfig() CodebuildProjectVpcConfigOutputReference
	VpcConfigInput() *CodebuildProjectVpcConfig
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
	PutArtifacts(value *CodebuildProjectArtifacts)
	PutBuildBatchConfig(value *CodebuildProjectBuildBatchConfig)
	PutCache(value *CodebuildProjectCache)
	PutEnvironment(value *CodebuildProjectEnvironment)
	PutLogsConfig(value *CodebuildProjectLogsConfig)
	PutSource(value *CodebuildProjectSource)
	PutVpcConfig(value *CodebuildProjectVpcConfig)
	ResetBadgeEnabled()
	ResetBuildBatchConfig()
	ResetBuildTimeout()
	ResetCache()
	ResetConcurrentBuildLimit()
	ResetDescription()
	ResetEncryptionKey()
	ResetFileSystemLocations()
	ResetLogsConfig()
	ResetOverrideLogicalId()
	ResetProjectVisibility()
	ResetQueuedTimeout()
	ResetResourceAccessRole()
	ResetSecondaryArtifacts()
	ResetSecondarySources()
	ResetSecondarySourceVersion()
	ResetSourceVersion()
	ResetTags()
	ResetTagsAll()
	ResetVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CodebuildProject
type jsiiProxy_CodebuildProject struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodebuildProject) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Artifacts() CodebuildProjectArtifactsOutputReference {
	var returns CodebuildProjectArtifactsOutputReference
	_jsii_.Get(
		j,
		"artifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ArtifactsInput() *CodebuildProjectArtifacts {
	var returns *CodebuildProjectArtifacts
	_jsii_.Get(
		j,
		"artifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BadgeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"badgeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BadgeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"badgeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BadgeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badgeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BuildBatchConfig() CodebuildProjectBuildBatchConfigOutputReference {
	var returns CodebuildProjectBuildBatchConfigOutputReference
	_jsii_.Get(
		j,
		"buildBatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BuildBatchConfigInput() *CodebuildProjectBuildBatchConfig {
	var returns *CodebuildProjectBuildBatchConfig
	_jsii_.Get(
		j,
		"buildBatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BuildTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) BuildTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Cache() CodebuildProjectCacheOutputReference {
	var returns CodebuildProjectCacheOutputReference
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) CacheInput() *CodebuildProjectCache {
	var returns *CodebuildProjectCache
	_jsii_.Get(
		j,
		"cacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ConcurrentBuildLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentBuildLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ConcurrentBuildLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentBuildLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) EncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Environment() CodebuildProjectEnvironmentOutputReference {
	var returns CodebuildProjectEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) EnvironmentInput() *CodebuildProjectEnvironment {
	var returns *CodebuildProjectEnvironment
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) FileSystemLocations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) FileSystemLocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) LogsConfig() CodebuildProjectLogsConfigOutputReference {
	var returns CodebuildProjectLogsConfigOutputReference
	_jsii_.Get(
		j,
		"logsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) LogsConfigInput() *CodebuildProjectLogsConfig {
	var returns *CodebuildProjectLogsConfig
	_jsii_.Get(
		j,
		"logsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ProjectVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ProjectVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) PublicProjectAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicProjectAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) QueuedTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queuedTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) QueuedTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queuedTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ResourceAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ResourceAccessRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAccessRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondaryArtifacts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondaryArtifactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondarySources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondarySourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondarySourceVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SecondarySourceVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySourceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Source() CodebuildProjectSourceOutputReference {
	var returns CodebuildProjectSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SourceInput() *CodebuildProjectSource {
	var returns *CodebuildProjectSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SourceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) SourceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) VpcConfig() CodebuildProjectVpcConfigOutputReference {
	var returns CodebuildProjectVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProject) VpcConfigInput() *CodebuildProjectVpcConfig {
	var returns *CodebuildProjectVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project aws_codebuild_project} Resource.
func NewCodebuildProject(scope constructs.Construct, id *string, config *CodebuildProjectConfig) CodebuildProject {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProject{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProject",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project aws_codebuild_project} Resource.
func NewCodebuildProject_Override(c CodebuildProject, scope constructs.Construct, id *string, config *CodebuildProjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProject",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodebuildProject) SetBadgeEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"badgeEnabled",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetBuildTimeout(val *float64) {
	_jsii_.Set(
		j,
		"buildTimeout",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetConcurrentBuildLimit(val *float64) {
	_jsii_.Set(
		j,
		"concurrentBuildLimit",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetFileSystemLocations(val interface{}) {
	_jsii_.Set(
		j,
		"fileSystemLocations",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetProjectVisibility(val *string) {
	_jsii_.Set(
		j,
		"projectVisibility",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetQueuedTimeout(val *float64) {
	_jsii_.Set(
		j,
		"queuedTimeout",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetResourceAccessRole(val *string) {
	_jsii_.Set(
		j,
		"resourceAccessRole",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetSecondaryArtifacts(val interface{}) {
	_jsii_.Set(
		j,
		"secondaryArtifacts",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetSecondarySources(val interface{}) {
	_jsii_.Set(
		j,
		"secondarySources",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetSecondarySourceVersion(val interface{}) {
	_jsii_.Set(
		j,
		"secondarySourceVersion",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetSourceVersion(val *string) {
	_jsii_.Set(
		j,
		"sourceVersion",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CodebuildProject) SetTagsAll(val *map[string]*string) {
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
func CodebuildProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.codebuild.CodebuildProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodebuildProject_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.codebuild.CodebuildProject",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CodebuildProject) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CodebuildProject) PutArtifacts(value *CodebuildProjectArtifacts) {
	_jsii_.InvokeVoid(
		c,
		"putArtifacts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutBuildBatchConfig(value *CodebuildProjectBuildBatchConfig) {
	_jsii_.InvokeVoid(
		c,
		"putBuildBatchConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutCache(value *CodebuildProjectCache) {
	_jsii_.InvokeVoid(
		c,
		"putCache",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutEnvironment(value *CodebuildProjectEnvironment) {
	_jsii_.InvokeVoid(
		c,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutLogsConfig(value *CodebuildProjectLogsConfig) {
	_jsii_.InvokeVoid(
		c,
		"putLogsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutSource(value *CodebuildProjectSource) {
	_jsii_.InvokeVoid(
		c,
		"putSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) PutVpcConfig(value *CodebuildProjectVpcConfig) {
	_jsii_.InvokeVoid(
		c,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProject) ResetBadgeEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetBadgeEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetBuildBatchConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildBatchConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetBuildTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetCache() {
	_jsii_.InvokeVoid(
		c,
		"resetCache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetConcurrentBuildLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetConcurrentBuildLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetFileSystemLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetFileSystemLocations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetLogsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLogsConfig",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CodebuildProject) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetProjectVisibility() {
	_jsii_.InvokeVoid(
		c,
		"resetProjectVisibility",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetQueuedTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetQueuedTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetResourceAccessRole() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceAccessRole",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetSecondaryArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondaryArtifacts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetSecondarySources() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondarySources",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetSecondarySourceVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetSecondarySourceVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetSourceVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProject) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProject) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CodebuildProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CodebuildProject) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CodebuildProjectArtifacts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#artifact_identifier CodebuildProject#artifact_identifier}.
	ArtifactIdentifier *string `json:"artifactIdentifier" yaml:"artifactIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#bucket_owner_access CodebuildProject#bucket_owner_access}.
	BucketOwnerAccess *string `json:"bucketOwnerAccess" yaml:"bucketOwnerAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#encryption_disabled CodebuildProject#encryption_disabled}.
	EncryptionDisabled interface{} `json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#location CodebuildProject#location}.
	Location *string `json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#name CodebuildProject#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#namespace_type CodebuildProject#namespace_type}.
	NamespaceType *string `json:"namespaceType" yaml:"namespaceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#override_artifact_name CodebuildProject#override_artifact_name}.
	OverrideArtifactName interface{} `json:"overrideArtifactName" yaml:"overrideArtifactName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#packaging CodebuildProject#packaging}.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#path CodebuildProject#path}.
	Path *string `json:"path" yaml:"path"`
}

type CodebuildProjectArtifactsOutputReference interface {
	cdktf.ComplexObject
	ArtifactIdentifier() *string
	SetArtifactIdentifier(val *string)
	ArtifactIdentifierInput() *string
	BucketOwnerAccess() *string
	SetBucketOwnerAccess(val *string)
	BucketOwnerAccessInput() *string
	EncryptionDisabled() interface{}
	SetEncryptionDisabled(val interface{})
	EncryptionDisabledInput() interface{}
	InternalValue() *CodebuildProjectArtifacts
	SetInternalValue(val *CodebuildProjectArtifacts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamespaceType() *string
	SetNamespaceType(val *string)
	NamespaceTypeInput() *string
	OverrideArtifactName() interface{}
	SetOverrideArtifactName(val interface{})
	OverrideArtifactNameInput() interface{}
	Packaging() *string
	SetPackaging(val *string)
	PackagingInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	ResetArtifactIdentifier()
	ResetBucketOwnerAccess()
	ResetEncryptionDisabled()
	ResetLocation()
	ResetName()
	ResetNamespaceType()
	ResetOverrideArtifactName()
	ResetPackaging()
	ResetPath()
}

// The jsii proxy struct for CodebuildProjectArtifactsOutputReference
type jsiiProxy_CodebuildProjectArtifactsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) ArtifactIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) ArtifactIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) BucketOwnerAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) BucketOwnerAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) EncryptionDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) EncryptionDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) InternalValue() *CodebuildProjectArtifacts {
	var returns *CodebuildProjectArtifacts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) NamespaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) NamespaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) OverrideArtifactName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideArtifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) OverrideArtifactNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideArtifactNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Packaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) PackagingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewCodebuildProjectArtifactsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectArtifactsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectArtifactsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectArtifactsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectArtifactsOutputReference_Override(c CodebuildProjectArtifactsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectArtifactsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetArtifactIdentifier(val *string) {
	_jsii_.Set(
		j,
		"artifactIdentifier",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetBucketOwnerAccess(val *string) {
	_jsii_.Set(
		j,
		"bucketOwnerAccess",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetEncryptionDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionDisabled",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetInternalValue(val *CodebuildProjectArtifacts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetNamespaceType(val *string) {
	_jsii_.Set(
		j,
		"namespaceType",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetOverrideArtifactName(val interface{}) {
	_jsii_.Set(
		j,
		"overrideArtifactName",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetPackaging(val *string) {
	_jsii_.Set(
		j,
		"packaging",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetArtifactIdentifier() {
	_jsii_.InvokeVoid(
		c,
		"resetArtifactIdentifier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetBucketOwnerAccess() {
	_jsii_.InvokeVoid(
		c,
		"resetBucketOwnerAccess",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetEncryptionDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetNamespaceType() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespaceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetOverrideArtifactName() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideArtifactName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetPackaging() {
	_jsii_.InvokeVoid(
		c,
		"resetPackaging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPath",
		nil, // no parameters
	)
}

type CodebuildProjectBuildBatchConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#service_role CodebuildProject#service_role}.
	ServiceRole *string `json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#combine_artifacts CodebuildProject#combine_artifacts}.
	CombineArtifacts interface{} `json:"combineArtifacts" yaml:"combineArtifacts"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#restrictions CodebuildProject#restrictions}
	Restrictions *CodebuildProjectBuildBatchConfigRestrictions `json:"restrictions" yaml:"restrictions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#timeout_in_mins CodebuildProject#timeout_in_mins}.
	TimeoutInMins *float64 `json:"timeoutInMins" yaml:"timeoutInMins"`
}

type CodebuildProjectBuildBatchConfigOutputReference interface {
	cdktf.ComplexObject
	CombineArtifacts() interface{}
	SetCombineArtifacts(val interface{})
	CombineArtifactsInput() interface{}
	InternalValue() *CodebuildProjectBuildBatchConfig
	SetInternalValue(val *CodebuildProjectBuildBatchConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Restrictions() CodebuildProjectBuildBatchConfigRestrictionsOutputReference
	RestrictionsInput() *CodebuildProjectBuildBatchConfigRestrictions
	ServiceRole() *string
	SetServiceRole(val *string)
	ServiceRoleInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutInMins() *float64
	SetTimeoutInMins(val *float64)
	TimeoutInMinsInput() *float64
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
	PutRestrictions(value *CodebuildProjectBuildBatchConfigRestrictions)
	ResetCombineArtifacts()
	ResetRestrictions()
	ResetTimeoutInMins()
}

// The jsii proxy struct for CodebuildProjectBuildBatchConfigOutputReference
type jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) CombineArtifacts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"combineArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) CombineArtifactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"combineArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) InternalValue() *CodebuildProjectBuildBatchConfig {
	var returns *CodebuildProjectBuildBatchConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) Restrictions() CodebuildProjectBuildBatchConfigRestrictionsOutputReference {
	var returns CodebuildProjectBuildBatchConfigRestrictionsOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) RestrictionsInput() *CodebuildProjectBuildBatchConfigRestrictions {
	var returns *CodebuildProjectBuildBatchConfigRestrictions
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) TimeoutInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) TimeoutInMinsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinsInput",
		&returns,
	)
	return returns
}

func NewCodebuildProjectBuildBatchConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectBuildBatchConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectBuildBatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectBuildBatchConfigOutputReference_Override(c CodebuildProjectBuildBatchConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectBuildBatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) SetCombineArtifacts(val interface{}) {
	_jsii_.Set(
		j,
		"combineArtifacts",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) SetInternalValue(val *CodebuildProjectBuildBatchConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) SetTimeoutInMins(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInMins",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) PutRestrictions(value *CodebuildProjectBuildBatchConfigRestrictions) {
	_jsii_.InvokeVoid(
		c,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) ResetCombineArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetCombineArtifacts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) ResetRestrictions() {
	_jsii_.InvokeVoid(
		c,
		"resetRestrictions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigOutputReference) ResetTimeoutInMins() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutInMins",
		nil, // no parameters
	)
}

type CodebuildProjectBuildBatchConfigRestrictions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#compute_types_allowed CodebuildProject#compute_types_allowed}.
	ComputeTypesAllowed *[]*string `json:"computeTypesAllowed" yaml:"computeTypesAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#maximum_builds_allowed CodebuildProject#maximum_builds_allowed}.
	MaximumBuildsAllowed *float64 `json:"maximumBuildsAllowed" yaml:"maximumBuildsAllowed"`
}

type CodebuildProjectBuildBatchConfigRestrictionsOutputReference interface {
	cdktf.ComplexObject
	ComputeTypesAllowed() *[]*string
	SetComputeTypesAllowed(val *[]*string)
	ComputeTypesAllowedInput() *[]*string
	InternalValue() *CodebuildProjectBuildBatchConfigRestrictions
	SetInternalValue(val *CodebuildProjectBuildBatchConfigRestrictions)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaximumBuildsAllowed() *float64
	SetMaximumBuildsAllowed(val *float64)
	MaximumBuildsAllowedInput() *float64
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
	ResetComputeTypesAllowed()
	ResetMaximumBuildsAllowed()
}

// The jsii proxy struct for CodebuildProjectBuildBatchConfigRestrictionsOutputReference
type jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ComputeTypesAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"computeTypesAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ComputeTypesAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"computeTypesAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) InternalValue() *CodebuildProjectBuildBatchConfigRestrictions {
	var returns *CodebuildProjectBuildBatchConfigRestrictions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) MaximumBuildsAllowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBuildsAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) MaximumBuildsAllowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBuildsAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildProjectBuildBatchConfigRestrictionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectBuildBatchConfigRestrictionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectBuildBatchConfigRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectBuildBatchConfigRestrictionsOutputReference_Override(c CodebuildProjectBuildBatchConfigRestrictionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectBuildBatchConfigRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) SetComputeTypesAllowed(val *[]*string) {
	_jsii_.Set(
		j,
		"computeTypesAllowed",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) SetInternalValue(val *CodebuildProjectBuildBatchConfigRestrictions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) SetMaximumBuildsAllowed(val *float64) {
	_jsii_.Set(
		j,
		"maximumBuildsAllowed",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ResetComputeTypesAllowed() {
	_jsii_.InvokeVoid(
		c,
		"resetComputeTypesAllowed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ResetMaximumBuildsAllowed() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumBuildsAllowed",
		nil, // no parameters
	)
}

type CodebuildProjectCache struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#location CodebuildProject#location}.
	Location *string `json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#modes CodebuildProject#modes}.
	Modes *[]*string `json:"modes" yaml:"modes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
}

type CodebuildProjectCacheOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *CodebuildProjectCache
	SetInternalValue(val *CodebuildProjectCache)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Modes() *[]*string
	SetModes(val *[]*string)
	ModesInput() *[]*string
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
	ResetLocation()
	ResetModes()
	ResetType()
}

// The jsii proxy struct for CodebuildProjectCacheOutputReference
type jsiiProxy_CodebuildProjectCacheOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) InternalValue() *CodebuildProjectCache {
	var returns *CodebuildProjectCache
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) Modes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) ModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"modesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewCodebuildProjectCacheOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectCacheOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectCacheOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectCacheOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectCacheOutputReference_Override(c CodebuildProjectCacheOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectCacheOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) SetInternalValue(val *CodebuildProjectCache) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) SetModes(val *[]*string) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectCacheOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectCacheOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectCacheOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectCacheOutputReference) ResetModes() {
	_jsii_.InvokeVoid(
		c,
		"resetModes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectCacheOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

// AWS CodeBuild.
type CodebuildProjectConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// artifacts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#artifacts CodebuildProject#artifacts}
	Artifacts *CodebuildProjectArtifacts `json:"artifacts" yaml:"artifacts"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#environment CodebuildProject#environment}
	Environment *CodebuildProjectEnvironment `json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#name CodebuildProject#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#service_role CodebuildProject#service_role}.
	ServiceRole *string `json:"serviceRole" yaml:"serviceRole"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#source CodebuildProject#source}
	Source *CodebuildProjectSource `json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#badge_enabled CodebuildProject#badge_enabled}.
	BadgeEnabled interface{} `json:"badgeEnabled" yaml:"badgeEnabled"`
	// build_batch_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#build_batch_config CodebuildProject#build_batch_config}
	BuildBatchConfig *CodebuildProjectBuildBatchConfig `json:"buildBatchConfig" yaml:"buildBatchConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#build_timeout CodebuildProject#build_timeout}.
	BuildTimeout *float64 `json:"buildTimeout" yaml:"buildTimeout"`
	// cache block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#cache CodebuildProject#cache}
	Cache *CodebuildProjectCache `json:"cache" yaml:"cache"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#concurrent_build_limit CodebuildProject#concurrent_build_limit}.
	ConcurrentBuildLimit *float64 `json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#description CodebuildProject#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#encryption_key CodebuildProject#encryption_key}.
	EncryptionKey *string `json:"encryptionKey" yaml:"encryptionKey"`
	// file_system_locations block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#file_system_locations CodebuildProject#file_system_locations}
	FileSystemLocations interface{} `json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// logs_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#logs_config CodebuildProject#logs_config}
	LogsConfig *CodebuildProjectLogsConfig `json:"logsConfig" yaml:"logsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#project_visibility CodebuildProject#project_visibility}.
	ProjectVisibility *string `json:"projectVisibility" yaml:"projectVisibility"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#queued_timeout CodebuildProject#queued_timeout}.
	QueuedTimeout *float64 `json:"queuedTimeout" yaml:"queuedTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#resource_access_role CodebuildProject#resource_access_role}.
	ResourceAccessRole *string `json:"resourceAccessRole" yaml:"resourceAccessRole"`
	// secondary_artifacts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#secondary_artifacts CodebuildProject#secondary_artifacts}
	SecondaryArtifacts interface{} `json:"secondaryArtifacts" yaml:"secondaryArtifacts"`
	// secondary_sources block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#secondary_sources CodebuildProject#secondary_sources}
	SecondarySources interface{} `json:"secondarySources" yaml:"secondarySources"`
	// secondary_source_version block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#secondary_source_version CodebuildProject#secondary_source_version}
	SecondarySourceVersion interface{} `json:"secondarySourceVersion" yaml:"secondarySourceVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#source_version CodebuildProject#source_version}.
	SourceVersion *string `json:"sourceVersion" yaml:"sourceVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#tags CodebuildProject#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#tags_all CodebuildProject#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#vpc_config CodebuildProject#vpc_config}
	VpcConfig *CodebuildProjectVpcConfig `json:"vpcConfig" yaml:"vpcConfig"`
}

type CodebuildProjectEnvironment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#compute_type CodebuildProject#compute_type}.
	ComputeType *string `json:"computeType" yaml:"computeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#image CodebuildProject#image}.
	Image *string `json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#certificate CodebuildProject#certificate}.
	Certificate *string `json:"certificate" yaml:"certificate"`
	// environment_variable block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#environment_variable CodebuildProject#environment_variable}
	EnvironmentVariable interface{} `json:"environmentVariable" yaml:"environmentVariable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#image_pull_credentials_type CodebuildProject#image_pull_credentials_type}.
	ImagePullCredentialsType *string `json:"imagePullCredentialsType" yaml:"imagePullCredentialsType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#privileged_mode CodebuildProject#privileged_mode}.
	PrivilegedMode interface{} `json:"privilegedMode" yaml:"privilegedMode"`
	// registry_credential block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#registry_credential CodebuildProject#registry_credential}
	RegistryCredential *CodebuildProjectEnvironmentRegistryCredential `json:"registryCredential" yaml:"registryCredential"`
}

type CodebuildProjectEnvironmentEnvironmentVariable struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#name CodebuildProject#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#value CodebuildProject#value}.
	Value *string `json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
}

type CodebuildProjectEnvironmentOutputReference interface {
	cdktf.ComplexObject
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	ComputeType() *string
	SetComputeType(val *string)
	ComputeTypeInput() *string
	EnvironmentVariable() interface{}
	SetEnvironmentVariable(val interface{})
	EnvironmentVariableInput() interface{}
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	ImagePullCredentialsType() *string
	SetImagePullCredentialsType(val *string)
	ImagePullCredentialsTypeInput() *string
	InternalValue() *CodebuildProjectEnvironment
	SetInternalValue(val *CodebuildProjectEnvironment)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	PrivilegedMode() interface{}
	SetPrivilegedMode(val interface{})
	PrivilegedModeInput() interface{}
	RegistryCredential() CodebuildProjectEnvironmentRegistryCredentialOutputReference
	RegistryCredentialInput() *CodebuildProjectEnvironmentRegistryCredential
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
	PutRegistryCredential(value *CodebuildProjectEnvironmentRegistryCredential)
	ResetCertificate()
	ResetEnvironmentVariable()
	ResetImagePullCredentialsType()
	ResetPrivilegedMode()
	ResetRegistryCredential()
}

// The jsii proxy struct for CodebuildProjectEnvironmentOutputReference
type jsiiProxy_CodebuildProjectEnvironmentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ComputeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ComputeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) EnvironmentVariable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) EnvironmentVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ImagePullCredentialsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullCredentialsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ImagePullCredentialsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullCredentialsTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) InternalValue() *CodebuildProjectEnvironment {
	var returns *CodebuildProjectEnvironment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) PrivilegedMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) PrivilegedModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) RegistryCredential() CodebuildProjectEnvironmentRegistryCredentialOutputReference {
	var returns CodebuildProjectEnvironmentRegistryCredentialOutputReference
	_jsii_.Get(
		j,
		"registryCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) RegistryCredentialInput() *CodebuildProjectEnvironmentRegistryCredential {
	var returns *CodebuildProjectEnvironmentRegistryCredential
	_jsii_.Get(
		j,
		"registryCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewCodebuildProjectEnvironmentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectEnvironmentOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectEnvironmentOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectEnvironmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectEnvironmentOutputReference_Override(c CodebuildProjectEnvironmentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectEnvironmentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetComputeType(val *string) {
	_jsii_.Set(
		j,
		"computeType",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetEnvironmentVariable(val interface{}) {
	_jsii_.Set(
		j,
		"environmentVariable",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetImage(val *string) {
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetImagePullCredentialsType(val *string) {
	_jsii_.Set(
		j,
		"imagePullCredentialsType",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetInternalValue(val *CodebuildProjectEnvironment) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetPrivilegedMode(val interface{}) {
	_jsii_.Set(
		j,
		"privilegedMode",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) PutRegistryCredential(value *CodebuildProjectEnvironmentRegistryCredential) {
	_jsii_.InvokeVoid(
		c,
		"putRegistryCredential",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		c,
		"resetCertificate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetEnvironmentVariable() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvironmentVariable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetImagePullCredentialsType() {
	_jsii_.InvokeVoid(
		c,
		"resetImagePullCredentialsType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetPrivilegedMode() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivilegedMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectEnvironmentOutputReference) ResetRegistryCredential() {
	_jsii_.InvokeVoid(
		c,
		"resetRegistryCredential",
		nil, // no parameters
	)
}

type CodebuildProjectEnvironmentRegistryCredential struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#credential CodebuildProject#credential}.
	Credential *string `json:"credential" yaml:"credential"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#credential_provider CodebuildProject#credential_provider}.
	CredentialProvider *string `json:"credentialProvider" yaml:"credentialProvider"`
}

type CodebuildProjectEnvironmentRegistryCredentialOutputReference interface {
	cdktf.ComplexObject
	Credential() *string
	SetCredential(val *string)
	CredentialInput() *string
	CredentialProvider() *string
	SetCredentialProvider(val *string)
	CredentialProviderInput() *string
	InternalValue() *CodebuildProjectEnvironmentRegistryCredential
	SetInternalValue(val *CodebuildProjectEnvironmentRegistryCredential)
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

// The jsii proxy struct for CodebuildProjectEnvironmentRegistryCredentialOutputReference
type jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) Credential() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) CredentialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) CredentialProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) CredentialProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) InternalValue() *CodebuildProjectEnvironmentRegistryCredential {
	var returns *CodebuildProjectEnvironmentRegistryCredential
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildProjectEnvironmentRegistryCredentialOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectEnvironmentRegistryCredentialOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectEnvironmentRegistryCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectEnvironmentRegistryCredentialOutputReference_Override(c CodebuildProjectEnvironmentRegistryCredentialOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectEnvironmentRegistryCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) SetCredential(val *string) {
	_jsii_.Set(
		j,
		"credential",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) SetCredentialProvider(val *string) {
	_jsii_.Set(
		j,
		"credentialProvider",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) SetInternalValue(val *CodebuildProjectEnvironmentRegistryCredential) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectEnvironmentRegistryCredentialOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CodebuildProjectFileSystemLocations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#identifier CodebuildProject#identifier}.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#location CodebuildProject#location}.
	Location *string `json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#mount_options CodebuildProject#mount_options}.
	MountOptions *string `json:"mountOptions" yaml:"mountOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#mount_point CodebuildProject#mount_point}.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
}

type CodebuildProjectLogsConfig struct {
	// cloudwatch_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#cloudwatch_logs CodebuildProject#cloudwatch_logs}
	CloudwatchLogs *CodebuildProjectLogsConfigCloudwatchLogs `json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#s3_logs CodebuildProject#s3_logs}
	S3Logs *CodebuildProjectLogsConfigS3Logs `json:"s3Logs" yaml:"s3Logs"`
}

type CodebuildProjectLogsConfigCloudwatchLogs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#group_name CodebuildProject#group_name}.
	GroupName *string `json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#status CodebuildProject#status}.
	Status *string `json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#stream_name CodebuildProject#stream_name}.
	StreamName *string `json:"streamName" yaml:"streamName"`
}

type CodebuildProjectLogsConfigCloudwatchLogsOutputReference interface {
	cdktf.ComplexObject
	GroupName() *string
	SetGroupName(val *string)
	GroupNameInput() *string
	InternalValue() *CodebuildProjectLogsConfigCloudwatchLogs
	SetInternalValue(val *CodebuildProjectLogsConfigCloudwatchLogs)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	StreamName() *string
	SetStreamName(val *string)
	StreamNameInput() *string
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
	ResetGroupName()
	ResetStatus()
	ResetStreamName()
}

// The jsii proxy struct for CodebuildProjectLogsConfigCloudwatchLogsOutputReference
type jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) InternalValue() *CodebuildProjectLogsConfigCloudwatchLogs {
	var returns *CodebuildProjectLogsConfigCloudwatchLogs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) StreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) StreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildProjectLogsConfigCloudwatchLogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectLogsConfigCloudwatchLogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectLogsConfigCloudwatchLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectLogsConfigCloudwatchLogsOutputReference_Override(c CodebuildProjectLogsConfigCloudwatchLogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectLogsConfigCloudwatchLogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) SetGroupName(val *string) {
	_jsii_.Set(
		j,
		"groupName",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) SetInternalValue(val *CodebuildProjectLogsConfigCloudwatchLogs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) SetStreamName(val *string) {
	_jsii_.Set(
		j,
		"streamName",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) ResetGroupName() {
	_jsii_.InvokeVoid(
		c,
		"resetGroupName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectLogsConfigCloudwatchLogsOutputReference) ResetStreamName() {
	_jsii_.InvokeVoid(
		c,
		"resetStreamName",
		nil, // no parameters
	)
}

type CodebuildProjectLogsConfigOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLogs() CodebuildProjectLogsConfigCloudwatchLogsOutputReference
	CloudwatchLogsInput() *CodebuildProjectLogsConfigCloudwatchLogs
	InternalValue() *CodebuildProjectLogsConfig
	SetInternalValue(val *CodebuildProjectLogsConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	S3Logs() CodebuildProjectLogsConfigS3LogsOutputReference
	S3LogsInput() *CodebuildProjectLogsConfigS3Logs
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
	PutCloudwatchLogs(value *CodebuildProjectLogsConfigCloudwatchLogs)
	PutS3Logs(value *CodebuildProjectLogsConfigS3Logs)
	ResetCloudwatchLogs()
	ResetS3Logs()
}

// The jsii proxy struct for CodebuildProjectLogsConfigOutputReference
type jsiiProxy_CodebuildProjectLogsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) CloudwatchLogs() CodebuildProjectLogsConfigCloudwatchLogsOutputReference {
	var returns CodebuildProjectLogsConfigCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) CloudwatchLogsInput() *CodebuildProjectLogsConfigCloudwatchLogs {
	var returns *CodebuildProjectLogsConfigCloudwatchLogs
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) InternalValue() *CodebuildProjectLogsConfig {
	var returns *CodebuildProjectLogsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) S3Logs() CodebuildProjectLogsConfigS3LogsOutputReference {
	var returns CodebuildProjectLogsConfigS3LogsOutputReference
	_jsii_.Get(
		j,
		"s3Logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) S3LogsInput() *CodebuildProjectLogsConfigS3Logs {
	var returns *CodebuildProjectLogsConfigS3Logs
	_jsii_.Get(
		j,
		"s3LogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildProjectLogsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectLogsConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectLogsConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectLogsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectLogsConfigOutputReference_Override(c CodebuildProjectLogsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectLogsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) SetInternalValue(val *CodebuildProjectLogsConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) PutCloudwatchLogs(value *CodebuildProjectLogsConfigCloudwatchLogs) {
	_jsii_.InvokeVoid(
		c,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) PutS3Logs(value *CodebuildProjectLogsConfigS3Logs) {
	_jsii_.InvokeVoid(
		c,
		"putS3Logs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectLogsConfigOutputReference) ResetS3Logs() {
	_jsii_.InvokeVoid(
		c,
		"resetS3Logs",
		nil, // no parameters
	)
}

type CodebuildProjectLogsConfigS3Logs struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#bucket_owner_access CodebuildProject#bucket_owner_access}.
	BucketOwnerAccess *string `json:"bucketOwnerAccess" yaml:"bucketOwnerAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#encryption_disabled CodebuildProject#encryption_disabled}.
	EncryptionDisabled interface{} `json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#location CodebuildProject#location}.
	Location *string `json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#status CodebuildProject#status}.
	Status *string `json:"status" yaml:"status"`
}

type CodebuildProjectLogsConfigS3LogsOutputReference interface {
	cdktf.ComplexObject
	BucketOwnerAccess() *string
	SetBucketOwnerAccess(val *string)
	BucketOwnerAccessInput() *string
	EncryptionDisabled() interface{}
	SetEncryptionDisabled(val interface{})
	EncryptionDisabledInput() interface{}
	InternalValue() *CodebuildProjectLogsConfigS3Logs
	SetInternalValue(val *CodebuildProjectLogsConfigS3Logs)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
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
	ResetBucketOwnerAccess()
	ResetEncryptionDisabled()
	ResetLocation()
	ResetStatus()
}

// The jsii proxy struct for CodebuildProjectLogsConfigS3LogsOutputReference
type jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) BucketOwnerAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) BucketOwnerAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) EncryptionDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) EncryptionDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) InternalValue() *CodebuildProjectLogsConfigS3Logs {
	var returns *CodebuildProjectLogsConfigS3Logs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildProjectLogsConfigS3LogsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectLogsConfigS3LogsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectLogsConfigS3LogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectLogsConfigS3LogsOutputReference_Override(c CodebuildProjectLogsConfigS3LogsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectLogsConfigS3LogsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) SetBucketOwnerAccess(val *string) {
	_jsii_.Set(
		j,
		"bucketOwnerAccess",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) SetEncryptionDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionDisabled",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) SetInternalValue(val *CodebuildProjectLogsConfigS3Logs) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) ResetBucketOwnerAccess() {
	_jsii_.InvokeVoid(
		c,
		"resetBucketOwnerAccess",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) ResetEncryptionDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectLogsConfigS3LogsOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetStatus",
		nil, // no parameters
	)
}

type CodebuildProjectSecondaryArtifacts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#artifact_identifier CodebuildProject#artifact_identifier}.
	ArtifactIdentifier *string `json:"artifactIdentifier" yaml:"artifactIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#bucket_owner_access CodebuildProject#bucket_owner_access}.
	BucketOwnerAccess *string `json:"bucketOwnerAccess" yaml:"bucketOwnerAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#encryption_disabled CodebuildProject#encryption_disabled}.
	EncryptionDisabled interface{} `json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#location CodebuildProject#location}.
	Location *string `json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#name CodebuildProject#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#namespace_type CodebuildProject#namespace_type}.
	NamespaceType *string `json:"namespaceType" yaml:"namespaceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#override_artifact_name CodebuildProject#override_artifact_name}.
	OverrideArtifactName interface{} `json:"overrideArtifactName" yaml:"overrideArtifactName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#packaging CodebuildProject#packaging}.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#path CodebuildProject#path}.
	Path *string `json:"path" yaml:"path"`
}

type CodebuildProjectSecondarySourceVersion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#source_identifier CodebuildProject#source_identifier}.
	SourceIdentifier *string `json:"sourceIdentifier" yaml:"sourceIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#source_version CodebuildProject#source_version}.
	SourceVersion *string `json:"sourceVersion" yaml:"sourceVersion"`
}

type CodebuildProjectSecondarySources struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#source_identifier CodebuildProject#source_identifier}.
	SourceIdentifier *string `json:"sourceIdentifier" yaml:"sourceIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#auth CodebuildProject#auth}
	Auth *CodebuildProjectSecondarySourcesAuth `json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#buildspec CodebuildProject#buildspec}.
	Buildspec *string `json:"buildspec" yaml:"buildspec"`
	// build_status_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#build_status_config CodebuildProject#build_status_config}
	BuildStatusConfig *CodebuildProjectSecondarySourcesBuildStatusConfig `json:"buildStatusConfig" yaml:"buildStatusConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#git_clone_depth CodebuildProject#git_clone_depth}.
	GitCloneDepth *float64 `json:"gitCloneDepth" yaml:"gitCloneDepth"`
	// git_submodules_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#git_submodules_config CodebuildProject#git_submodules_config}
	GitSubmodulesConfig *CodebuildProjectSecondarySourcesGitSubmodulesConfig `json:"gitSubmodulesConfig" yaml:"gitSubmodulesConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#insecure_ssl CodebuildProject#insecure_ssl}.
	InsecureSsl interface{} `json:"insecureSsl" yaml:"insecureSsl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#location CodebuildProject#location}.
	Location *string `json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#report_build_status CodebuildProject#report_build_status}.
	ReportBuildStatus interface{} `json:"reportBuildStatus" yaml:"reportBuildStatus"`
}

type CodebuildProjectSecondarySourcesAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#resource CodebuildProject#resource}.
	Resource *string `json:"resource" yaml:"resource"`
}

type CodebuildProjectSecondarySourcesAuthOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *CodebuildProjectSecondarySourcesAuth
	SetInternalValue(val *CodebuildProjectSecondarySourcesAuth)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Resource() *string
	SetResource(val *string)
	ResourceInput() *string
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
	ResetResource()
}

// The jsii proxy struct for CodebuildProjectSecondarySourcesAuthOutputReference
type jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) InternalValue() *CodebuildProjectSecondarySourcesAuth {
	var returns *CodebuildProjectSecondarySourcesAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewCodebuildProjectSecondarySourcesAuthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectSecondarySourcesAuthOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSecondarySourcesAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectSecondarySourcesAuthOutputReference_Override(c CodebuildProjectSecondarySourcesAuthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSecondarySourcesAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) SetInternalValue(val *CodebuildProjectSecondarySourcesAuth) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) SetResource(val *string) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesAuthOutputReference) ResetResource() {
	_jsii_.InvokeVoid(
		c,
		"resetResource",
		nil, // no parameters
	)
}

type CodebuildProjectSecondarySourcesBuildStatusConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#context CodebuildProject#context}.
	Context *string `json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#target_url CodebuildProject#target_url}.
	TargetUrl *string `json:"targetUrl" yaml:"targetUrl"`
}

type CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference interface {
	cdktf.ComplexObject
	Context() *string
	SetContext(val *string)
	ContextInput() *string
	InternalValue() *CodebuildProjectSecondarySourcesBuildStatusConfig
	SetInternalValue(val *CodebuildProjectSecondarySourcesBuildStatusConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TargetUrl() *string
	SetTargetUrl(val *string)
	TargetUrlInput() *string
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
	ResetContext()
	ResetTargetUrl()
}

// The jsii proxy struct for CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference
type jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) InternalValue() *CodebuildProjectSecondarySourcesBuildStatusConfig {
	var returns *CodebuildProjectSecondarySourcesBuildStatusConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) TargetUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) TargetUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildProjectSecondarySourcesBuildStatusConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectSecondarySourcesBuildStatusConfigOutputReference_Override(c CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) SetInternalValue(val *CodebuildProjectSecondarySourcesBuildStatusConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) SetTargetUrl(val *string) {
	_jsii_.Set(
		j,
		"targetUrl",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) ResetContext() {
	_jsii_.InvokeVoid(
		c,
		"resetContext",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference) ResetTargetUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetUrl",
		nil, // no parameters
	)
}

type CodebuildProjectSecondarySourcesGitSubmodulesConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#fetch_submodules CodebuildProject#fetch_submodules}.
	FetchSubmodules interface{} `json:"fetchSubmodules" yaml:"fetchSubmodules"`
}

type CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference interface {
	cdktf.ComplexObject
	FetchSubmodules() interface{}
	SetFetchSubmodules(val interface{})
	FetchSubmodulesInput() interface{}
	InternalValue() *CodebuildProjectSecondarySourcesGitSubmodulesConfig
	SetInternalValue(val *CodebuildProjectSecondarySourcesGitSubmodulesConfig)
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

// The jsii proxy struct for CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference
type jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) FetchSubmodules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchSubmodules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) FetchSubmodulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchSubmodulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) InternalValue() *CodebuildProjectSecondarySourcesGitSubmodulesConfig {
	var returns *CodebuildProjectSecondarySourcesGitSubmodulesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference_Override(c CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) SetFetchSubmodules(val interface{}) {
	_jsii_.Set(
		j,
		"fetchSubmodules",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) SetInternalValue(val *CodebuildProjectSecondarySourcesGitSubmodulesConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CodebuildProjectSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#auth CodebuildProject#auth}
	Auth *CodebuildProjectSourceAuth `json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#buildspec CodebuildProject#buildspec}.
	Buildspec *string `json:"buildspec" yaml:"buildspec"`
	// build_status_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#build_status_config CodebuildProject#build_status_config}
	BuildStatusConfig *CodebuildProjectSourceBuildStatusConfig `json:"buildStatusConfig" yaml:"buildStatusConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#git_clone_depth CodebuildProject#git_clone_depth}.
	GitCloneDepth *float64 `json:"gitCloneDepth" yaml:"gitCloneDepth"`
	// git_submodules_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#git_submodules_config CodebuildProject#git_submodules_config}
	GitSubmodulesConfig *CodebuildProjectSourceGitSubmodulesConfig `json:"gitSubmodulesConfig" yaml:"gitSubmodulesConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#insecure_ssl CodebuildProject#insecure_ssl}.
	InsecureSsl interface{} `json:"insecureSsl" yaml:"insecureSsl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#location CodebuildProject#location}.
	Location *string `json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#report_build_status CodebuildProject#report_build_status}.
	ReportBuildStatus interface{} `json:"reportBuildStatus" yaml:"reportBuildStatus"`
}

type CodebuildProjectSourceAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#type CodebuildProject#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#resource CodebuildProject#resource}.
	Resource *string `json:"resource" yaml:"resource"`
}

type CodebuildProjectSourceAuthOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *CodebuildProjectSourceAuth
	SetInternalValue(val *CodebuildProjectSourceAuth)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Resource() *string
	SetResource(val *string)
	ResourceInput() *string
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
	ResetResource()
}

// The jsii proxy struct for CodebuildProjectSourceAuthOutputReference
type jsiiProxy_CodebuildProjectSourceAuthOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) InternalValue() *CodebuildProjectSourceAuth {
	var returns *CodebuildProjectSourceAuth
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) ResourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewCodebuildProjectSourceAuthOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectSourceAuthOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectSourceAuthOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSourceAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectSourceAuthOutputReference_Override(c CodebuildProjectSourceAuthOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSourceAuthOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) SetInternalValue(val *CodebuildProjectSourceAuth) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) SetResource(val *string) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceAuthOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectSourceAuthOutputReference) ResetResource() {
	_jsii_.InvokeVoid(
		c,
		"resetResource",
		nil, // no parameters
	)
}

type CodebuildProjectSourceBuildStatusConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#context CodebuildProject#context}.
	Context *string `json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#target_url CodebuildProject#target_url}.
	TargetUrl *string `json:"targetUrl" yaml:"targetUrl"`
}

type CodebuildProjectSourceBuildStatusConfigOutputReference interface {
	cdktf.ComplexObject
	Context() *string
	SetContext(val *string)
	ContextInput() *string
	InternalValue() *CodebuildProjectSourceBuildStatusConfig
	SetInternalValue(val *CodebuildProjectSourceBuildStatusConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TargetUrl() *string
	SetTargetUrl(val *string)
	TargetUrlInput() *string
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
	ResetContext()
	ResetTargetUrl()
}

// The jsii proxy struct for CodebuildProjectSourceBuildStatusConfigOutputReference
type jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) InternalValue() *CodebuildProjectSourceBuildStatusConfig {
	var returns *CodebuildProjectSourceBuildStatusConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) TargetUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) TargetUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildProjectSourceBuildStatusConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectSourceBuildStatusConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSourceBuildStatusConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectSourceBuildStatusConfigOutputReference_Override(c CodebuildProjectSourceBuildStatusConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSourceBuildStatusConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) SetInternalValue(val *CodebuildProjectSourceBuildStatusConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) SetTargetUrl(val *string) {
	_jsii_.Set(
		j,
		"targetUrl",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) ResetContext() {
	_jsii_.InvokeVoid(
		c,
		"resetContext",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSourceBuildStatusConfigOutputReference) ResetTargetUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetUrl",
		nil, // no parameters
	)
}

type CodebuildProjectSourceGitSubmodulesConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#fetch_submodules CodebuildProject#fetch_submodules}.
	FetchSubmodules interface{} `json:"fetchSubmodules" yaml:"fetchSubmodules"`
}

type CodebuildProjectSourceGitSubmodulesConfigOutputReference interface {
	cdktf.ComplexObject
	FetchSubmodules() interface{}
	SetFetchSubmodules(val interface{})
	FetchSubmodulesInput() interface{}
	InternalValue() *CodebuildProjectSourceGitSubmodulesConfig
	SetInternalValue(val *CodebuildProjectSourceGitSubmodulesConfig)
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

// The jsii proxy struct for CodebuildProjectSourceGitSubmodulesConfigOutputReference
type jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) FetchSubmodules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchSubmodules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) FetchSubmodulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchSubmodulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) InternalValue() *CodebuildProjectSourceGitSubmodulesConfig {
	var returns *CodebuildProjectSourceGitSubmodulesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildProjectSourceGitSubmodulesConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectSourceGitSubmodulesConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSourceGitSubmodulesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectSourceGitSubmodulesConfigOutputReference_Override(c CodebuildProjectSourceGitSubmodulesConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSourceGitSubmodulesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) SetFetchSubmodules(val interface{}) {
	_jsii_.Set(
		j,
		"fetchSubmodules",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) SetInternalValue(val *CodebuildProjectSourceGitSubmodulesConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceGitSubmodulesConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CodebuildProjectSourceOutputReference interface {
	cdktf.ComplexObject
	Auth() CodebuildProjectSourceAuthOutputReference
	AuthInput() *CodebuildProjectSourceAuth
	Buildspec() *string
	SetBuildspec(val *string)
	BuildspecInput() *string
	BuildStatusConfig() CodebuildProjectSourceBuildStatusConfigOutputReference
	BuildStatusConfigInput() *CodebuildProjectSourceBuildStatusConfig
	GitCloneDepth() *float64
	SetGitCloneDepth(val *float64)
	GitCloneDepthInput() *float64
	GitSubmodulesConfig() CodebuildProjectSourceGitSubmodulesConfigOutputReference
	GitSubmodulesConfigInput() *CodebuildProjectSourceGitSubmodulesConfig
	InsecureSsl() interface{}
	SetInsecureSsl(val interface{})
	InsecureSslInput() interface{}
	InternalValue() *CodebuildProjectSource
	SetInternalValue(val *CodebuildProjectSource)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ReportBuildStatus() interface{}
	SetReportBuildStatus(val interface{})
	ReportBuildStatusInput() interface{}
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
	PutAuth(value *CodebuildProjectSourceAuth)
	PutBuildStatusConfig(value *CodebuildProjectSourceBuildStatusConfig)
	PutGitSubmodulesConfig(value *CodebuildProjectSourceGitSubmodulesConfig)
	ResetAuth()
	ResetBuildspec()
	ResetBuildStatusConfig()
	ResetGitCloneDepth()
	ResetGitSubmodulesConfig()
	ResetInsecureSsl()
	ResetLocation()
	ResetReportBuildStatus()
}

// The jsii proxy struct for CodebuildProjectSourceOutputReference
type jsiiProxy_CodebuildProjectSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) Auth() CodebuildProjectSourceAuthOutputReference {
	var returns CodebuildProjectSourceAuthOutputReference
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) AuthInput() *CodebuildProjectSourceAuth {
	var returns *CodebuildProjectSourceAuth
	_jsii_.Get(
		j,
		"authInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) Buildspec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) BuildspecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) BuildStatusConfig() CodebuildProjectSourceBuildStatusConfigOutputReference {
	var returns CodebuildProjectSourceBuildStatusConfigOutputReference
	_jsii_.Get(
		j,
		"buildStatusConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) BuildStatusConfigInput() *CodebuildProjectSourceBuildStatusConfig {
	var returns *CodebuildProjectSourceBuildStatusConfig
	_jsii_.Get(
		j,
		"buildStatusConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) GitCloneDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) GitCloneDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) GitSubmodulesConfig() CodebuildProjectSourceGitSubmodulesConfigOutputReference {
	var returns CodebuildProjectSourceGitSubmodulesConfigOutputReference
	_jsii_.Get(
		j,
		"gitSubmodulesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) GitSubmodulesConfigInput() *CodebuildProjectSourceGitSubmodulesConfig {
	var returns *CodebuildProjectSourceGitSubmodulesConfig
	_jsii_.Get(
		j,
		"gitSubmodulesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) InsecureSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) InsecureSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) InternalValue() *CodebuildProjectSource {
	var returns *CodebuildProjectSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) ReportBuildStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) ReportBuildStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewCodebuildProjectSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectSourceOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectSourceOutputReference_Override(c CodebuildProjectSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetBuildspec(val *string) {
	_jsii_.Set(
		j,
		"buildspec",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetGitCloneDepth(val *float64) {
	_jsii_.Set(
		j,
		"gitCloneDepth",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetInsecureSsl(val interface{}) {
	_jsii_.Set(
		j,
		"insecureSsl",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetInternalValue(val *CodebuildProjectSource) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetReportBuildStatus(val interface{}) {
	_jsii_.Set(
		j,
		"reportBuildStatus",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSourceOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) PutAuth(value *CodebuildProjectSourceAuth) {
	_jsii_.InvokeVoid(
		c,
		"putAuth",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) PutBuildStatusConfig(value *CodebuildProjectSourceBuildStatusConfig) {
	_jsii_.InvokeVoid(
		c,
		"putBuildStatusConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) PutGitSubmodulesConfig(value *CodebuildProjectSourceGitSubmodulesConfig) {
	_jsii_.InvokeVoid(
		c,
		"putGitSubmodulesConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) ResetAuth() {
	_jsii_.InvokeVoid(
		c,
		"resetAuth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) ResetBuildspec() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildspec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) ResetBuildStatusConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildStatusConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) ResetGitCloneDepth() {
	_jsii_.InvokeVoid(
		c,
		"resetGitCloneDepth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) ResetGitSubmodulesConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGitSubmodulesConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) ResetInsecureSsl() {
	_jsii_.InvokeVoid(
		c,
		"resetInsecureSsl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSourceOutputReference) ResetReportBuildStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetReportBuildStatus",
		nil, // no parameters
	)
}

type CodebuildProjectVpcConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#security_group_ids CodebuildProject#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#subnets CodebuildProject#subnets}.
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_project#vpc_id CodebuildProject#vpc_id}.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

type CodebuildProjectVpcConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *CodebuildProjectVpcConfig
	SetInternalValue(val *CodebuildProjectVpcConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	SubnetsInput() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcId() *string
	SetVpcId(val *string)
	VpcIdInput() *string
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

// The jsii proxy struct for CodebuildProjectVpcConfigOutputReference
type jsiiProxy_CodebuildProjectVpcConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) InternalValue() *CodebuildProjectVpcConfig {
	var returns *CodebuildProjectVpcConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func NewCodebuildProjectVpcConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildProjectVpcConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildProjectVpcConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildProjectVpcConfigOutputReference_Override(c CodebuildProjectVpcConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildProjectVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SetInternalValue(val *CodebuildProjectVpcConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SetSubnets(val *[]*string) {
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectVpcConfigOutputReference) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildProjectVpcConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group aws_codebuild_report_group}.
type CodebuildReportGroup interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	Created() *string
	DeleteReports() interface{}
	SetDeleteReports(val interface{})
	DeleteReportsInput() interface{}
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ExportConfig() CodebuildReportGroupExportConfigOutputReference
	ExportConfigInput() *CodebuildReportGroupExportConfig
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
	PutExportConfig(value *CodebuildReportGroupExportConfig)
	ResetDeleteReports()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CodebuildReportGroup
type jsiiProxy_CodebuildReportGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodebuildReportGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) DeleteReports() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) DeleteReportsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) ExportConfig() CodebuildReportGroupExportConfigOutputReference {
	var returns CodebuildReportGroupExportConfigOutputReference
	_jsii_.Get(
		j,
		"exportConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) ExportConfigInput() *CodebuildReportGroupExportConfig {
	var returns *CodebuildReportGroupExportConfig
	_jsii_.Get(
		j,
		"exportConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroup) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group aws_codebuild_report_group} Resource.
func NewCodebuildReportGroup(scope constructs.Construct, id *string, config *CodebuildReportGroupConfig) CodebuildReportGroup {
	_init_.Initialize()

	j := jsiiProxy_CodebuildReportGroup{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildReportGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group aws_codebuild_report_group} Resource.
func NewCodebuildReportGroup_Override(c CodebuildReportGroup, scope constructs.Construct, id *string, config *CodebuildReportGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildReportGroup",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodebuildReportGroup) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroup) SetDeleteReports(val interface{}) {
	_jsii_.Set(
		j,
		"deleteReports",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroup) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroup) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroup) SetType(val *string) {
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
func CodebuildReportGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.codebuild.CodebuildReportGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodebuildReportGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.codebuild.CodebuildReportGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CodebuildReportGroup) PutExportConfig(value *CodebuildReportGroupExportConfig) {
	_jsii_.InvokeVoid(
		c,
		"putExportConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildReportGroup) ResetDeleteReports() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteReports",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildReportGroup) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildReportGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildReportGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CodebuildReportGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CodebuildReportGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS CodeBuild.
type CodebuildReportGroupConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// export_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#export_config CodebuildReportGroup#export_config}
	ExportConfig *CodebuildReportGroupExportConfig `json:"exportConfig" yaml:"exportConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#name CodebuildReportGroup#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#type CodebuildReportGroup#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#delete_reports CodebuildReportGroup#delete_reports}.
	DeleteReports interface{} `json:"deleteReports" yaml:"deleteReports"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#tags CodebuildReportGroup#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#tags_all CodebuildReportGroup#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

type CodebuildReportGroupExportConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#type CodebuildReportGroup#type}.
	Type *string `json:"type" yaml:"type"`
	// s3_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#s3_destination CodebuildReportGroup#s3_destination}
	S3Destination *CodebuildReportGroupExportConfigS3Destination `json:"s3Destination" yaml:"s3Destination"`
}

type CodebuildReportGroupExportConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *CodebuildReportGroupExportConfig
	SetInternalValue(val *CodebuildReportGroupExportConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	S3Destination() CodebuildReportGroupExportConfigS3DestinationOutputReference
	S3DestinationInput() *CodebuildReportGroupExportConfigS3Destination
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
	PutS3Destination(value *CodebuildReportGroupExportConfigS3Destination)
	ResetS3Destination()
}

// The jsii proxy struct for CodebuildReportGroupExportConfigOutputReference
type jsiiProxy_CodebuildReportGroupExportConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) InternalValue() *CodebuildReportGroupExportConfig {
	var returns *CodebuildReportGroupExportConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) S3Destination() CodebuildReportGroupExportConfigS3DestinationOutputReference {
	var returns CodebuildReportGroupExportConfigS3DestinationOutputReference
	_jsii_.Get(
		j,
		"s3Destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) S3DestinationInput() *CodebuildReportGroupExportConfigS3Destination {
	var returns *CodebuildReportGroupExportConfigS3Destination
	_jsii_.Get(
		j,
		"s3DestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewCodebuildReportGroupExportConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildReportGroupExportConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildReportGroupExportConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildReportGroupExportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildReportGroupExportConfigOutputReference_Override(c CodebuildReportGroupExportConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildReportGroupExportConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) SetInternalValue(val *CodebuildReportGroupExportConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) PutS3Destination(value *CodebuildReportGroupExportConfigS3Destination) {
	_jsii_.InvokeVoid(
		c,
		"putS3Destination",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildReportGroupExportConfigOutputReference) ResetS3Destination() {
	_jsii_.InvokeVoid(
		c,
		"resetS3Destination",
		nil, // no parameters
	)
}

type CodebuildReportGroupExportConfigS3Destination struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#bucket CodebuildReportGroup#bucket}.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#encryption_key CodebuildReportGroup#encryption_key}.
	EncryptionKey *string `json:"encryptionKey" yaml:"encryptionKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#encryption_disabled CodebuildReportGroup#encryption_disabled}.
	EncryptionDisabled interface{} `json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#packaging CodebuildReportGroup#packaging}.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_report_group#path CodebuildReportGroup#path}.
	Path *string `json:"path" yaml:"path"`
}

type CodebuildReportGroupExportConfigS3DestinationOutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	EncryptionDisabled() interface{}
	SetEncryptionDisabled(val interface{})
	EncryptionDisabledInput() interface{}
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	EncryptionKeyInput() *string
	InternalValue() *CodebuildReportGroupExportConfigS3Destination
	SetInternalValue(val *CodebuildReportGroupExportConfigS3Destination)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Packaging() *string
	SetPackaging(val *string)
	PackagingInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	ResetEncryptionDisabled()
	ResetPackaging()
	ResetPath()
}

// The jsii proxy struct for CodebuildReportGroupExportConfigS3DestinationOutputReference
type jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) EncryptionDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) EncryptionDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) EncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) InternalValue() *CodebuildReportGroupExportConfigS3Destination {
	var returns *CodebuildReportGroupExportConfigS3Destination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) Packaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) PackagingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCodebuildReportGroupExportConfigS3DestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) CodebuildReportGroupExportConfigS3DestinationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildReportGroupExportConfigS3DestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCodebuildReportGroupExportConfigS3DestinationOutputReference_Override(c CodebuildReportGroupExportConfigS3DestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildReportGroupExportConfigS3DestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) SetEncryptionDisabled(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionDisabled",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) SetEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) SetInternalValue(val *CodebuildReportGroupExportConfigS3Destination) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) SetPackaging(val *string) {
	_jsii_.Set(
		j,
		"packaging",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) ResetEncryptionDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) ResetPackaging() {
	_jsii_.InvokeVoid(
		c,
		"resetPackaging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildReportGroupExportConfigS3DestinationOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPath",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/codebuild_resource_policy aws_codebuild_resource_policy}.
type CodebuildResourcePolicy interface {
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
	Node() constructs.Node
	Policy() *string
	SetPolicy(val *string)
	PolicyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceArn() *string
	SetResourceArn(val *string)
	ResourceArnInput() *string
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

// The jsii proxy struct for CodebuildResourcePolicy
type jsiiProxy_CodebuildResourcePolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodebuildResourcePolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) ResourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildResourcePolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_resource_policy aws_codebuild_resource_policy} Resource.
func NewCodebuildResourcePolicy(scope constructs.Construct, id *string, config *CodebuildResourcePolicyConfig) CodebuildResourcePolicy {
	_init_.Initialize()

	j := jsiiProxy_CodebuildResourcePolicy{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildResourcePolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_resource_policy aws_codebuild_resource_policy} Resource.
func NewCodebuildResourcePolicy_Override(c CodebuildResourcePolicy, scope constructs.Construct, id *string, config *CodebuildResourcePolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildResourcePolicy",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodebuildResourcePolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodebuildResourcePolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodebuildResourcePolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodebuildResourcePolicy) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_CodebuildResourcePolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodebuildResourcePolicy) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CodebuildResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.codebuild.CodebuildResourcePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodebuildResourcePolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.codebuild.CodebuildResourcePolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildResourcePolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CodebuildResourcePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CodebuildResourcePolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS CodeBuild.
type CodebuildResourcePolicyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_resource_policy#policy CodebuildResourcePolicy#policy}.
	Policy *string `json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_resource_policy#resource_arn CodebuildResourcePolicy#resource_arn}.
	ResourceArn *string `json:"resourceArn" yaml:"resourceArn"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/codebuild_source_credential aws_codebuild_source_credential}.
type CodebuildSourceCredential interface {
	cdktf.TerraformResource
	Arn() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
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
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServerType() *string
	SetServerType(val *string)
	ServerTypeInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
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
	ResetUserName()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CodebuildSourceCredential
type jsiiProxy_CodebuildSourceCredential struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodebuildSourceCredential) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) ServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) ServerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildSourceCredential) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_source_credential aws_codebuild_source_credential} Resource.
func NewCodebuildSourceCredential(scope constructs.Construct, id *string, config *CodebuildSourceCredentialConfig) CodebuildSourceCredential {
	_init_.Initialize()

	j := jsiiProxy_CodebuildSourceCredential{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildSourceCredential",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_source_credential aws_codebuild_source_credential} Resource.
func NewCodebuildSourceCredential_Override(c CodebuildSourceCredential, scope constructs.Construct, id *string, config *CodebuildSourceCredentialConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildSourceCredential",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodebuildSourceCredential) SetAuthType(val *string) {
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_CodebuildSourceCredential) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodebuildSourceCredential) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodebuildSourceCredential) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodebuildSourceCredential) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CodebuildSourceCredential) SetServerType(val *string) {
	_jsii_.Set(
		j,
		"serverType",
		val,
	)
}

func (j *jsiiProxy_CodebuildSourceCredential) SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_CodebuildSourceCredential) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CodebuildSourceCredential_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.codebuild.CodebuildSourceCredential",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodebuildSourceCredential_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.codebuild.CodebuildSourceCredential",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildSourceCredential) ResetUserName() {
	_jsii_.InvokeVoid(
		c,
		"resetUserName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildSourceCredential) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CodebuildSourceCredential) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CodebuildSourceCredential) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS CodeBuild.
type CodebuildSourceCredentialConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_source_credential#auth_type CodebuildSourceCredential#auth_type}.
	AuthType *string `json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_source_credential#server_type CodebuildSourceCredential#server_type}.
	ServerType *string `json:"serverType" yaml:"serverType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_source_credential#token CodebuildSourceCredential#token}.
	Token *string `json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_source_credential#user_name CodebuildSourceCredential#user_name}.
	UserName *string `json:"userName" yaml:"userName"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook aws_codebuild_webhook}.
type CodebuildWebhook interface {
	cdktf.TerraformResource
	BranchFilter() *string
	SetBranchFilter(val *string)
	BranchFilterInput() *string
	BuildType() *string
	SetBuildType(val *string)
	BuildTypeInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FilterGroup() interface{}
	SetFilterGroup(val interface{})
	FilterGroupInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PayloadUrl() *string
	ProjectName() *string
	SetProjectName(val *string)
	ProjectNameInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Secret() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Url() *string
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
	ResetBranchFilter()
	ResetBuildType()
	ResetFilterGroup()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CodebuildWebhook
type jsiiProxy_CodebuildWebhook struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CodebuildWebhook) BranchFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) BranchFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) BuildType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) BuildTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) FilterGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) FilterGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) PayloadUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) ProjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) Secret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildWebhook) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook aws_codebuild_webhook} Resource.
func NewCodebuildWebhook(scope constructs.Construct, id *string, config *CodebuildWebhookConfig) CodebuildWebhook {
	_init_.Initialize()

	j := jsiiProxy_CodebuildWebhook{}

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildWebhook",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook aws_codebuild_webhook} Resource.
func NewCodebuildWebhook_Override(c CodebuildWebhook, scope constructs.Construct, id *string, config *CodebuildWebhookConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.codebuild.CodebuildWebhook",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CodebuildWebhook) SetBranchFilter(val *string) {
	_jsii_.Set(
		j,
		"branchFilter",
		val,
	)
}

func (j *jsiiProxy_CodebuildWebhook) SetBuildType(val *string) {
	_jsii_.Set(
		j,
		"buildType",
		val,
	)
}

func (j *jsiiProxy_CodebuildWebhook) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CodebuildWebhook) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CodebuildWebhook) SetFilterGroup(val interface{}) {
	_jsii_.Set(
		j,
		"filterGroup",
		val,
	)
}

func (j *jsiiProxy_CodebuildWebhook) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CodebuildWebhook) SetProjectName(val *string) {
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_CodebuildWebhook) SetProvider(val cdktf.TerraformProvider) {
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
func CodebuildWebhook_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.codebuild.CodebuildWebhook",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CodebuildWebhook_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.codebuild.CodebuildWebhook",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CodebuildWebhook) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CodebuildWebhook) ResetBranchFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetBranchFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildWebhook) ResetBuildType() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildWebhook) ResetFilterGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetFilterGroup",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CodebuildWebhook) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildWebhook) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CodebuildWebhook) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CodebuildWebhook) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (c *jsiiProxy_CodebuildWebhook) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS CodeBuild.
type CodebuildWebhookConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook#project_name CodebuildWebhook#project_name}.
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook#branch_filter CodebuildWebhook#branch_filter}.
	BranchFilter *string `json:"branchFilter" yaml:"branchFilter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook#build_type CodebuildWebhook#build_type}.
	BuildType *string `json:"buildType" yaml:"buildType"`
	// filter_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook#filter_group CodebuildWebhook#filter_group}
	FilterGroup interface{} `json:"filterGroup" yaml:"filterGroup"`
}

type CodebuildWebhookFilterGroup struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook#filter CodebuildWebhook#filter}
	Filter interface{} `json:"filter" yaml:"filter"`
}

type CodebuildWebhookFilterGroupFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook#pattern CodebuildWebhook#pattern}.
	Pattern *string `json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook#type CodebuildWebhook#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/codebuild_webhook#exclude_matched_pattern CodebuildWebhook#exclude_matched_pattern}.
	ExcludeMatchedPattern interface{} `json:"excludeMatchedPattern" yaml:"excludeMatchedPattern"`
}
