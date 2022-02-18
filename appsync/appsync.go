package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/appsync/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_cache aws_appsync_api_cache}.
type AppsyncApiCache interface {
	cdktf.TerraformResource
	ApiCachingBehavior() *string
	SetApiCachingBehavior(val *string)
	ApiCachingBehaviorInput() *string
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	AtRestEncryptionEnabled() interface{}
	SetAtRestEncryptionEnabled(val interface{})
	AtRestEncryptionEnabledInput() interface{}
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
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TransitEncryptionEnabled() interface{}
	SetTransitEncryptionEnabled(val interface{})
	TransitEncryptionEnabledInput() interface{}
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
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
	ResetAtRestEncryptionEnabled()
	ResetOverrideLogicalId()
	ResetTransitEncryptionEnabled()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppsyncApiCache
type jsiiProxy_AppsyncApiCache struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncApiCache) ApiCachingBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiCachingBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) ApiCachingBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiCachingBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) AtRestEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atRestEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) AtRestEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atRestEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) TransitEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) TransitEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiCache) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_cache aws_appsync_api_cache} Resource.
func NewAppsyncApiCache(scope constructs.Construct, id *string, config *AppsyncApiCacheConfig) AppsyncApiCache {
	_init_.Initialize()

	j := jsiiProxy_AppsyncApiCache{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncApiCache",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_cache aws_appsync_api_cache} Resource.
func NewAppsyncApiCache_Override(a AppsyncApiCache, scope constructs.Construct, id *string, config *AppsyncApiCacheConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncApiCache",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetApiCachingBehavior(val *string) {
	_jsii_.Set(
		j,
		"apiCachingBehavior",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetAtRestEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"atRestEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetTransitEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"transitEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetTtl(val *float64) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiCache) SetType(val *string) {
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
func AppsyncApiCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appsync.AppsyncApiCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncApiCache_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appsync.AppsyncApiCache",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppsyncApiCache) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncApiCache) ResetAtRestEncryptionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAtRestEncryptionEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppsyncApiCache) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiCache) ResetTransitEncryptionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitEncryptionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiCache) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiCache) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppsyncApiCache) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppsyncApiCache) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS AppSync.
type AppsyncApiCacheConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_cache#api_caching_behavior AppsyncApiCache#api_caching_behavior}.
	ApiCachingBehavior *string `json:"apiCachingBehavior" yaml:"apiCachingBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_cache#api_id AppsyncApiCache#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_cache#ttl AppsyncApiCache#ttl}.
	Ttl *float64 `json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_cache#type AppsyncApiCache#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_cache#at_rest_encryption_enabled AppsyncApiCache#at_rest_encryption_enabled}.
	AtRestEncryptionEnabled interface{} `json:"atRestEncryptionEnabled" yaml:"atRestEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_cache#transit_encryption_enabled AppsyncApiCache#transit_encryption_enabled}.
	TransitEncryptionEnabled interface{} `json:"transitEncryptionEnabled" yaml:"transitEncryptionEnabled"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_key aws_appsync_api_key}.
type AppsyncApiKey interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Expires() *string
	SetExpires(val *string)
	ExpiresInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Key() *string
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
	ResetExpires()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppsyncApiKey
type jsiiProxy_AppsyncApiKey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncApiKey) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) Expires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) ExpiresInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncApiKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_key aws_appsync_api_key} Resource.
func NewAppsyncApiKey(scope constructs.Construct, id *string, config *AppsyncApiKeyConfig) AppsyncApiKey {
	_init_.Initialize()

	j := jsiiProxy_AppsyncApiKey{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncApiKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_key aws_appsync_api_key} Resource.
func NewAppsyncApiKey_Override(a AppsyncApiKey, scope constructs.Construct, id *string, config *AppsyncApiKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncApiKey",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncApiKey) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiKey) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiKey) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiKey) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiKey) SetExpires(val *string) {
	_jsii_.Set(
		j,
		"expires",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiKey) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncApiKey) SetProvider(val cdktf.TerraformProvider) {
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
func AppsyncApiKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appsync.AppsyncApiKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncApiKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appsync.AppsyncApiKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppsyncApiKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncApiKey) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiKey) ResetExpires() {
	_jsii_.InvokeVoid(
		a,
		"resetExpires",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppsyncApiKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncApiKey) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncApiKey) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppsyncApiKey) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppsyncApiKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS AppSync.
type AppsyncApiKeyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_key#api_id AppsyncApiKey#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_key#description AppsyncApiKey#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_api_key#expires AppsyncApiKey#expires}.
	Expires *string `json:"expires" yaml:"expires"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource aws_appsync_datasource}.
type AppsyncDatasource interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DynamodbConfig() AppsyncDatasourceDynamodbConfigOutputReference
	DynamodbConfigInput() *AppsyncDatasourceDynamodbConfig
	ElasticsearchConfig() AppsyncDatasourceElasticsearchConfigOutputReference
	ElasticsearchConfigInput() *AppsyncDatasourceElasticsearchConfig
	Fqn() *string
	FriendlyUniqueId() *string
	HttpConfig() AppsyncDatasourceHttpConfigOutputReference
	HttpConfigInput() *AppsyncDatasourceHttpConfig
	Id() *string
	LambdaConfig() AppsyncDatasourceLambdaConfigOutputReference
	LambdaConfigInput() *AppsyncDatasourceLambdaConfig
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RelationalDatabaseConfig() AppsyncDatasourceRelationalDatabaseConfigOutputReference
	RelationalDatabaseConfigInput() *AppsyncDatasourceRelationalDatabaseConfig
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	ServiceRoleArnInput() *string
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
	PutDynamodbConfig(value *AppsyncDatasourceDynamodbConfig)
	PutElasticsearchConfig(value *AppsyncDatasourceElasticsearchConfig)
	PutHttpConfig(value *AppsyncDatasourceHttpConfig)
	PutLambdaConfig(value *AppsyncDatasourceLambdaConfig)
	PutRelationalDatabaseConfig(value *AppsyncDatasourceRelationalDatabaseConfig)
	ResetDescription()
	ResetDynamodbConfig()
	ResetElasticsearchConfig()
	ResetHttpConfig()
	ResetLambdaConfig()
	ResetOverrideLogicalId()
	ResetRelationalDatabaseConfig()
	ResetServiceRoleArn()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppsyncDatasource
type jsiiProxy_AppsyncDatasource struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncDatasource) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) DynamodbConfig() AppsyncDatasourceDynamodbConfigOutputReference {
	var returns AppsyncDatasourceDynamodbConfigOutputReference
	_jsii_.Get(
		j,
		"dynamodbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) DynamodbConfigInput() *AppsyncDatasourceDynamodbConfig {
	var returns *AppsyncDatasourceDynamodbConfig
	_jsii_.Get(
		j,
		"dynamodbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ElasticsearchConfig() AppsyncDatasourceElasticsearchConfigOutputReference {
	var returns AppsyncDatasourceElasticsearchConfigOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ElasticsearchConfigInput() *AppsyncDatasourceElasticsearchConfig {
	var returns *AppsyncDatasourceElasticsearchConfig
	_jsii_.Get(
		j,
		"elasticsearchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) HttpConfig() AppsyncDatasourceHttpConfigOutputReference {
	var returns AppsyncDatasourceHttpConfigOutputReference
	_jsii_.Get(
		j,
		"httpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) HttpConfigInput() *AppsyncDatasourceHttpConfig {
	var returns *AppsyncDatasourceHttpConfig
	_jsii_.Get(
		j,
		"httpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) LambdaConfig() AppsyncDatasourceLambdaConfigOutputReference {
	var returns AppsyncDatasourceLambdaConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) LambdaConfigInput() *AppsyncDatasourceLambdaConfig {
	var returns *AppsyncDatasourceLambdaConfig
	_jsii_.Get(
		j,
		"lambdaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) RelationalDatabaseConfig() AppsyncDatasourceRelationalDatabaseConfigOutputReference {
	var returns AppsyncDatasourceRelationalDatabaseConfigOutputReference
	_jsii_.Get(
		j,
		"relationalDatabaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) RelationalDatabaseConfigInput() *AppsyncDatasourceRelationalDatabaseConfig {
	var returns *AppsyncDatasourceRelationalDatabaseConfig
	_jsii_.Get(
		j,
		"relationalDatabaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasource) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource aws_appsync_datasource} Resource.
func NewAppsyncDatasource(scope constructs.Construct, id *string, config *AppsyncDatasourceConfig) AppsyncDatasource {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasource{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource aws_appsync_datasource} Resource.
func NewAppsyncDatasource_Override(a AppsyncDatasource, scope constructs.Construct, id *string, config *AppsyncDatasourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasource",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasource) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasource) SetType(val *string) {
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
func AppsyncDatasource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appsync.AppsyncDatasource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncDatasource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appsync.AppsyncDatasource",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppsyncDatasource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutDynamodbConfig(value *AppsyncDatasourceDynamodbConfig) {
	_jsii_.InvokeVoid(
		a,
		"putDynamodbConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutElasticsearchConfig(value *AppsyncDatasourceElasticsearchConfig) {
	_jsii_.InvokeVoid(
		a,
		"putElasticsearchConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutHttpConfig(value *AppsyncDatasourceHttpConfig) {
	_jsii_.InvokeVoid(
		a,
		"putHttpConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutLambdaConfig(value *AppsyncDatasourceLambdaConfig) {
	_jsii_.InvokeVoid(
		a,
		"putLambdaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) PutRelationalDatabaseConfig(value *AppsyncDatasourceRelationalDatabaseConfig) {
	_jsii_.InvokeVoid(
		a,
		"putRelationalDatabaseConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetDynamodbConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodbConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetElasticsearchConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetHttpConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetLambdaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConfig",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppsyncDatasource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetRelationalDatabaseConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRelationalDatabaseConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) ResetServiceRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppsyncDatasource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppsyncDatasource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS AppSync.
type AppsyncDatasourceConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#api_id AppsyncDatasource#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#name AppsyncDatasource#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#type AppsyncDatasource#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#description AppsyncDatasource#description}.
	Description *string `json:"description" yaml:"description"`
	// dynamodb_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#dynamodb_config AppsyncDatasource#dynamodb_config}
	DynamodbConfig *AppsyncDatasourceDynamodbConfig `json:"dynamodbConfig" yaml:"dynamodbConfig"`
	// elasticsearch_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#elasticsearch_config AppsyncDatasource#elasticsearch_config}
	ElasticsearchConfig *AppsyncDatasourceElasticsearchConfig `json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// http_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#http_config AppsyncDatasource#http_config}
	HttpConfig *AppsyncDatasourceHttpConfig `json:"httpConfig" yaml:"httpConfig"`
	// lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#lambda_config AppsyncDatasource#lambda_config}
	LambdaConfig *AppsyncDatasourceLambdaConfig `json:"lambdaConfig" yaml:"lambdaConfig"`
	// relational_database_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#relational_database_config AppsyncDatasource#relational_database_config}
	RelationalDatabaseConfig *AppsyncDatasourceRelationalDatabaseConfig `json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#service_role_arn AppsyncDatasource#service_role_arn}.
	ServiceRoleArn *string `json:"serviceRoleArn" yaml:"serviceRoleArn"`
}

type AppsyncDatasourceDynamodbConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#table_name AppsyncDatasource#table_name}.
	TableName *string `json:"tableName" yaml:"tableName"`
	// delta_sync_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#delta_sync_config AppsyncDatasource#delta_sync_config}
	DeltaSyncConfig *AppsyncDatasourceDynamodbConfigDeltaSyncConfig `json:"deltaSyncConfig" yaml:"deltaSyncConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#region AppsyncDatasource#region}.
	Region *string `json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#use_caller_credentials AppsyncDatasource#use_caller_credentials}.
	UseCallerCredentials interface{} `json:"useCallerCredentials" yaml:"useCallerCredentials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#versioned AppsyncDatasource#versioned}.
	Versioned interface{} `json:"versioned" yaml:"versioned"`
}

type AppsyncDatasourceDynamodbConfigDeltaSyncConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#delta_sync_table_name AppsyncDatasource#delta_sync_table_name}.
	DeltaSyncTableName *string `json:"deltaSyncTableName" yaml:"deltaSyncTableName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#base_table_ttl AppsyncDatasource#base_table_ttl}.
	BaseTableTtl *float64 `json:"baseTableTtl" yaml:"baseTableTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#delta_sync_table_ttl AppsyncDatasource#delta_sync_table_ttl}.
	DeltaSyncTableTtl *float64 `json:"deltaSyncTableTtl" yaml:"deltaSyncTableTtl"`
}

type AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference interface {
	cdktf.ComplexObject
	BaseTableTtl() *float64
	SetBaseTableTtl(val *float64)
	BaseTableTtlInput() *float64
	DeltaSyncTableName() *string
	SetDeltaSyncTableName(val *string)
	DeltaSyncTableNameInput() *string
	DeltaSyncTableTtl() *float64
	SetDeltaSyncTableTtl(val *float64)
	DeltaSyncTableTtlInput() *float64
	InternalValue() *AppsyncDatasourceDynamodbConfigDeltaSyncConfig
	SetInternalValue(val *AppsyncDatasourceDynamodbConfigDeltaSyncConfig)
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
	ResetBaseTableTtl()
	ResetDeltaSyncTableTtl()
}

// The jsii proxy struct for AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference
type jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) BaseTableTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseTableTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) BaseTableTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseTableTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) DeltaSyncTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSyncTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) DeltaSyncTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSyncTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) DeltaSyncTableTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSyncTableTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) DeltaSyncTableTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSyncTableTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) InternalValue() *AppsyncDatasourceDynamodbConfigDeltaSyncConfig {
	var returns *AppsyncDatasourceDynamodbConfigDeltaSyncConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference_Override(a AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) SetBaseTableTtl(val *float64) {
	_jsii_.Set(
		j,
		"baseTableTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) SetDeltaSyncTableName(val *string) {
	_jsii_.Set(
		j,
		"deltaSyncTableName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) SetDeltaSyncTableTtl(val *float64) {
	_jsii_.Set(
		j,
		"deltaSyncTableTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) SetInternalValue(val *AppsyncDatasourceDynamodbConfigDeltaSyncConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) ResetBaseTableTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseTableTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference) ResetDeltaSyncTableTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetDeltaSyncTableTtl",
		nil, // no parameters
	)
}

type AppsyncDatasourceDynamodbConfigOutputReference interface {
	cdktf.ComplexObject
	DeltaSyncConfig() AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference
	DeltaSyncConfigInput() *AppsyncDatasourceDynamodbConfigDeltaSyncConfig
	InternalValue() *AppsyncDatasourceDynamodbConfig
	SetInternalValue(val *AppsyncDatasourceDynamodbConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseCallerCredentials() interface{}
	SetUseCallerCredentials(val interface{})
	UseCallerCredentialsInput() interface{}
	Versioned() interface{}
	SetVersioned(val interface{})
	VersionedInput() interface{}
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
	PutDeltaSyncConfig(value *AppsyncDatasourceDynamodbConfigDeltaSyncConfig)
	ResetDeltaSyncConfig()
	ResetRegion()
	ResetUseCallerCredentials()
	ResetVersioned()
}

// The jsii proxy struct for AppsyncDatasourceDynamodbConfigOutputReference
type jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) DeltaSyncConfig() AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference {
	var returns AppsyncDatasourceDynamodbConfigDeltaSyncConfigOutputReference
	_jsii_.Get(
		j,
		"deltaSyncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) DeltaSyncConfigInput() *AppsyncDatasourceDynamodbConfigDeltaSyncConfig {
	var returns *AppsyncDatasourceDynamodbConfigDeltaSyncConfig
	_jsii_.Get(
		j,
		"deltaSyncConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) InternalValue() *AppsyncDatasourceDynamodbConfig {
	var returns *AppsyncDatasourceDynamodbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) UseCallerCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCallerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) UseCallerCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCallerCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) Versioned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) VersionedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionedInput",
		&returns,
	)
	return returns
}

func NewAppsyncDatasourceDynamodbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncDatasourceDynamodbConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceDynamodbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceDynamodbConfigOutputReference_Override(a AppsyncDatasourceDynamodbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceDynamodbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) SetInternalValue(val *AppsyncDatasourceDynamodbConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) SetUseCallerCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"useCallerCredentials",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) SetVersioned(val interface{}) {
	_jsii_.Set(
		j,
		"versioned",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) PutDeltaSyncConfig(value *AppsyncDatasourceDynamodbConfigDeltaSyncConfig) {
	_jsii_.InvokeVoid(
		a,
		"putDeltaSyncConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) ResetDeltaSyncConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDeltaSyncConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) ResetUseCallerCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetUseCallerCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceDynamodbConfigOutputReference) ResetVersioned() {
	_jsii_.InvokeVoid(
		a,
		"resetVersioned",
		nil, // no parameters
	)
}

type AppsyncDatasourceElasticsearchConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#endpoint AppsyncDatasource#endpoint}.
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#region AppsyncDatasource#region}.
	Region *string `json:"region" yaml:"region"`
}

type AppsyncDatasourceElasticsearchConfigOutputReference interface {
	cdktf.ComplexObject
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	InternalValue() *AppsyncDatasourceElasticsearchConfig
	SetInternalValue(val *AppsyncDatasourceElasticsearchConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	ResetRegion()
}

// The jsii proxy struct for AppsyncDatasourceElasticsearchConfigOutputReference
type jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) InternalValue() *AppsyncDatasourceElasticsearchConfig {
	var returns *AppsyncDatasourceElasticsearchConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncDatasourceElasticsearchConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncDatasourceElasticsearchConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceElasticsearchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceElasticsearchConfigOutputReference_Override(a AppsyncDatasourceElasticsearchConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceElasticsearchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) SetInternalValue(val *AppsyncDatasourceElasticsearchConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceElasticsearchConfigOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

type AppsyncDatasourceHttpConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#endpoint AppsyncDatasource#endpoint}.
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#authorization_config AppsyncDatasource#authorization_config}
	AuthorizationConfig *AppsyncDatasourceHttpConfigAuthorizationConfig `json:"authorizationConfig" yaml:"authorizationConfig"`
}

type AppsyncDatasourceHttpConfigAuthorizationConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#authorization_type AppsyncDatasource#authorization_type}.
	AuthorizationType *string `json:"authorizationType" yaml:"authorizationType"`
	// aws_iam_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#aws_iam_config AppsyncDatasource#aws_iam_config}
	AwsIamConfig *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig `json:"awsIamConfig" yaml:"awsIamConfig"`
}

type AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#signing_region AppsyncDatasource#signing_region}.
	SigningRegion *string `json:"signingRegion" yaml:"signingRegion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#signing_service_name AppsyncDatasource#signing_service_name}.
	SigningServiceName *string `json:"signingServiceName" yaml:"signingServiceName"`
}

type AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig
	SetInternalValue(val *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SigningRegion() *string
	SetSigningRegion(val *string)
	SigningRegionInput() *string
	SigningServiceName() *string
	SetSigningServiceName(val *string)
	SigningServiceNameInput() *string
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
	ResetSigningRegion()
	ResetSigningServiceName()
}

// The jsii proxy struct for AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference
type jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) InternalValue() *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig {
	var returns *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SigningRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SigningRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SigningServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SigningServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference_Override(a AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SetInternalValue(val *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SetSigningRegion(val *string) {
	_jsii_.Set(
		j,
		"signingRegion",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SetSigningServiceName(val *string) {
	_jsii_.Set(
		j,
		"signingServiceName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) ResetSigningRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetSigningRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference) ResetSigningServiceName() {
	_jsii_.InvokeVoid(
		a,
		"resetSigningServiceName",
		nil, // no parameters
	)
}

type AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference interface {
	cdktf.ComplexObject
	AuthorizationType() *string
	SetAuthorizationType(val *string)
	AuthorizationTypeInput() *string
	AwsIamConfig() AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference
	AwsIamConfigInput() *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig
	InternalValue() *AppsyncDatasourceHttpConfigAuthorizationConfig
	SetInternalValue(val *AppsyncDatasourceHttpConfigAuthorizationConfig)
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
	PutAwsIamConfig(value *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig)
	ResetAuthorizationType()
	ResetAwsIamConfig()
}

// The jsii proxy struct for AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference
type jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) AuthorizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) AwsIamConfig() AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference {
	var returns AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfigOutputReference
	_jsii_.Get(
		j,
		"awsIamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) AwsIamConfigInput() *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig {
	var returns *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig
	_jsii_.Get(
		j,
		"awsIamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) InternalValue() *AppsyncDatasourceHttpConfigAuthorizationConfig {
	var returns *AppsyncDatasourceHttpConfigAuthorizationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncDatasourceHttpConfigAuthorizationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceHttpConfigAuthorizationConfigOutputReference_Override(a AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) SetAuthorizationType(val *string) {
	_jsii_.Set(
		j,
		"authorizationType",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) SetInternalValue(val *AppsyncDatasourceHttpConfigAuthorizationConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) PutAwsIamConfig(value *AppsyncDatasourceHttpConfigAuthorizationConfigAwsIamConfig) {
	_jsii_.InvokeVoid(
		a,
		"putAwsIamConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) ResetAuthorizationType() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference) ResetAwsIamConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsIamConfig",
		nil, // no parameters
	)
}

type AppsyncDatasourceHttpConfigOutputReference interface {
	cdktf.ComplexObject
	AuthorizationConfig() AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference
	AuthorizationConfigInput() *AppsyncDatasourceHttpConfigAuthorizationConfig
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	InternalValue() *AppsyncDatasourceHttpConfig
	SetInternalValue(val *AppsyncDatasourceHttpConfig)
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
	PutAuthorizationConfig(value *AppsyncDatasourceHttpConfigAuthorizationConfig)
	ResetAuthorizationConfig()
}

// The jsii proxy struct for AppsyncDatasourceHttpConfigOutputReference
type jsiiProxy_AppsyncDatasourceHttpConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) AuthorizationConfig() AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference {
	var returns AppsyncDatasourceHttpConfigAuthorizationConfigOutputReference
	_jsii_.Get(
		j,
		"authorizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) AuthorizationConfigInput() *AppsyncDatasourceHttpConfigAuthorizationConfig {
	var returns *AppsyncDatasourceHttpConfigAuthorizationConfig
	_jsii_.Get(
		j,
		"authorizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) InternalValue() *AppsyncDatasourceHttpConfig {
	var returns *AppsyncDatasourceHttpConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncDatasourceHttpConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncDatasourceHttpConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasourceHttpConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceHttpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceHttpConfigOutputReference_Override(a AppsyncDatasourceHttpConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceHttpConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) SetInternalValue(val *AppsyncDatasourceHttpConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) PutAuthorizationConfig(value *AppsyncDatasourceHttpConfigAuthorizationConfig) {
	_jsii_.InvokeVoid(
		a,
		"putAuthorizationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasourceHttpConfigOutputReference) ResetAuthorizationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationConfig",
		nil, // no parameters
	)
}

type AppsyncDatasourceLambdaConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#function_arn AppsyncDatasource#function_arn}.
	FunctionArn *string `json:"functionArn" yaml:"functionArn"`
}

type AppsyncDatasourceLambdaConfigOutputReference interface {
	cdktf.ComplexObject
	FunctionArn() *string
	SetFunctionArn(val *string)
	FunctionArnInput() *string
	InternalValue() *AppsyncDatasourceLambdaConfig
	SetInternalValue(val *AppsyncDatasourceLambdaConfig)
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

// The jsii proxy struct for AppsyncDatasourceLambdaConfigOutputReference
type jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) FunctionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) InternalValue() *AppsyncDatasourceLambdaConfig {
	var returns *AppsyncDatasourceLambdaConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncDatasourceLambdaConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncDatasourceLambdaConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceLambdaConfigOutputReference_Override(a AppsyncDatasourceLambdaConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceLambdaConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) SetFunctionArn(val *string) {
	_jsii_.Set(
		j,
		"functionArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) SetInternalValue(val *AppsyncDatasourceLambdaConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceLambdaConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type AppsyncDatasourceRelationalDatabaseConfig struct {
	// http_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#http_endpoint_config AppsyncDatasource#http_endpoint_config}
	HttpEndpointConfig *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig `json:"httpEndpointConfig" yaml:"httpEndpointConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#source_type AppsyncDatasource#source_type}.
	SourceType *string `json:"sourceType" yaml:"sourceType"`
}

type AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#aws_secret_store_arn AppsyncDatasource#aws_secret_store_arn}.
	AwsSecretStoreArn *string `json:"awsSecretStoreArn" yaml:"awsSecretStoreArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#db_cluster_identifier AppsyncDatasource#db_cluster_identifier}.
	DbClusterIdentifier *string `json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#database_name AppsyncDatasource#database_name}.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#region AppsyncDatasource#region}.
	Region *string `json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_datasource#schema AppsyncDatasource#schema}.
	Schema *string `json:"schema" yaml:"schema"`
}

type AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference interface {
	cdktf.ComplexObject
	AwsSecretStoreArn() *string
	SetAwsSecretStoreArn(val *string)
	AwsSecretStoreArnInput() *string
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	DbClusterIdentifierInput() *string
	InternalValue() *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig
	SetInternalValue(val *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
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
	ResetDatabaseName()
	ResetRegion()
	ResetSchema()
}

// The jsii proxy struct for AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference
type jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) AwsSecretStoreArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretStoreArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) AwsSecretStoreArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsSecretStoreArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) DbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) InternalValue() *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig {
	var returns *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference_Override(a AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SetAwsSecretStoreArn(val *string) {
	_jsii_.Set(
		j,
		"awsSecretStoreArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SetDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SetInternalValue(val *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SetSchema(val *string) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference) ResetSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetSchema",
		nil, // no parameters
	)
}

type AppsyncDatasourceRelationalDatabaseConfigOutputReference interface {
	cdktf.ComplexObject
	HttpEndpointConfig() AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference
	HttpEndpointConfigInput() *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig
	InternalValue() *AppsyncDatasourceRelationalDatabaseConfig
	SetInternalValue(val *AppsyncDatasourceRelationalDatabaseConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SourceType() *string
	SetSourceType(val *string)
	SourceTypeInput() *string
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
	PutHttpEndpointConfig(value *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig)
	ResetHttpEndpointConfig()
	ResetSourceType()
}

// The jsii proxy struct for AppsyncDatasourceRelationalDatabaseConfigOutputReference
type jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) HttpEndpointConfig() AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference {
	var returns AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"httpEndpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) HttpEndpointConfigInput() *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig {
	var returns *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig
	_jsii_.Get(
		j,
		"httpEndpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) InternalValue() *AppsyncDatasourceRelationalDatabaseConfig {
	var returns *AppsyncDatasourceRelationalDatabaseConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) SourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncDatasourceRelationalDatabaseConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncDatasourceRelationalDatabaseConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceRelationalDatabaseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncDatasourceRelationalDatabaseConfigOutputReference_Override(a AppsyncDatasourceRelationalDatabaseConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDatasourceRelationalDatabaseConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) SetInternalValue(val *AppsyncDatasourceRelationalDatabaseConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) SetSourceType(val *string) {
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) PutHttpEndpointConfig(value *AppsyncDatasourceRelationalDatabaseConfigHttpEndpointConfig) {
	_jsii_.InvokeVoid(
		a,
		"putHttpEndpointConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) ResetHttpEndpointConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpEndpointConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDatasourceRelationalDatabaseConfigOutputReference) ResetSourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceType",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name aws_appsync_domain_name}.
type AppsyncDomainName interface {
	cdktf.TerraformResource
	AppsyncDomainName() *string
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HostedZoneId() *string
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
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppsyncDomainName
type jsiiProxy_AppsyncDomainName struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncDomainName) AppsyncDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appsyncDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainName) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name aws_appsync_domain_name} Resource.
func NewAppsyncDomainName(scope constructs.Construct, id *string, config *AppsyncDomainNameConfig) AppsyncDomainName {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDomainName{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDomainName",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name aws_appsync_domain_name} Resource.
func NewAppsyncDomainName_Override(a AppsyncDomainName, scope constructs.Construct, id *string, config *AppsyncDomainNameConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDomainName",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncDomainName) SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainName) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainName) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainName) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainName) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainName) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainName) SetProvider(val cdktf.TerraformProvider) {
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
func AppsyncDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appsync.AppsyncDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncDomainName_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appsync.AppsyncDomainName",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppsyncDomainName) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncDomainName) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppsyncDomainName) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDomainName) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainName) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppsyncDomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppsyncDomainName) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name_api_association aws_appsync_domain_name_api_association}.
type AppsyncDomainNameApiAssociation interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
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

// The jsii proxy struct for AppsyncDomainNameApiAssociation
type jsiiProxy_AppsyncDomainNameApiAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name_api_association aws_appsync_domain_name_api_association} Resource.
func NewAppsyncDomainNameApiAssociation(scope constructs.Construct, id *string, config *AppsyncDomainNameApiAssociationConfig) AppsyncDomainNameApiAssociation {
	_init_.Initialize()

	j := jsiiProxy_AppsyncDomainNameApiAssociation{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDomainNameApiAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name_api_association aws_appsync_domain_name_api_association} Resource.
func NewAppsyncDomainNameApiAssociation_Override(a AppsyncDomainNameApiAssociation, scope constructs.Construct, id *string, config *AppsyncDomainNameApiAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncDomainNameApiAssociation",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncDomainNameApiAssociation) SetProvider(val cdktf.TerraformProvider) {
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
func AppsyncDomainNameApiAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appsync.AppsyncDomainNameApiAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncDomainNameApiAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appsync.AppsyncDomainNameApiAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncDomainNameApiAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppsyncDomainNameApiAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS AppSync.
type AppsyncDomainNameApiAssociationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name_api_association#api_id AppsyncDomainNameApiAssociation#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name_api_association#domain_name AppsyncDomainNameApiAssociation#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
}

// AWS AppSync.
type AppsyncDomainNameConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name#certificate_arn AppsyncDomainName#certificate_arn}.
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name#domain_name AppsyncDomainName#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_domain_name#description AppsyncDomainName#description}.
	Description *string `json:"description" yaml:"description"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appsync_function aws_appsync_function}.
type AppsyncFunction interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	FunctionId() *string
	FunctionVersion() *string
	SetFunctionVersion(val *string)
	FunctionVersionInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxBatchSize() *float64
	SetMaxBatchSize(val *float64)
	MaxBatchSizeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequestMappingTemplate() *string
	SetRequestMappingTemplate(val *string)
	RequestMappingTemplateInput() *string
	ResponseMappingTemplate() *string
	SetResponseMappingTemplate(val *string)
	ResponseMappingTemplateInput() *string
	SyncConfig() AppsyncFunctionSyncConfigOutputReference
	SyncConfigInput() *AppsyncFunctionSyncConfig
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
	PutSyncConfig(value *AppsyncFunctionSyncConfig)
	ResetDescription()
	ResetFunctionVersion()
	ResetMaxBatchSize()
	ResetOverrideLogicalId()
	ResetSyncConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppsyncFunction
type jsiiProxy_AppsyncFunction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncFunction) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FunctionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FunctionVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) MaxBatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) MaxBatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) RequestMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) RequestMappingTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) ResponseMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) ResponseMappingTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) SyncConfig() AppsyncFunctionSyncConfigOutputReference {
	var returns AppsyncFunctionSyncConfigOutputReference
	_jsii_.Get(
		j,
		"syncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) SyncConfigInput() *AppsyncFunctionSyncConfig {
	var returns *AppsyncFunctionSyncConfig
	_jsii_.Get(
		j,
		"syncConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_function aws_appsync_function} Resource.
func NewAppsyncFunction(scope constructs.Construct, id *string, config *AppsyncFunctionConfig) AppsyncFunction {
	_init_.Initialize()

	j := jsiiProxy_AppsyncFunction{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_function aws_appsync_function} Resource.
func NewAppsyncFunction_Override(a AppsyncFunction, scope constructs.Construct, id *string, config *AppsyncFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncFunction",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetDataSource(val *string) {
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetFunctionVersion(val *string) {
	_jsii_.Set(
		j,
		"functionVersion",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetMaxBatchSize(val *float64) {
	_jsii_.Set(
		j,
		"maxBatchSize",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetRequestMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"requestMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunction) SetResponseMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"responseMappingTemplate",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AppsyncFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appsync.AppsyncFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appsync.AppsyncFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppsyncFunction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncFunction) PutSyncConfig(value *AppsyncFunctionSyncConfig) {
	_jsii_.InvokeVoid(
		a,
		"putSyncConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunction) ResetFunctionVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetFunctionVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunction) ResetMaxBatchSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxBatchSize",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppsyncFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunction) ResetSyncConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppsyncFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppsyncFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS AppSync.
type AppsyncFunctionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#api_id AppsyncFunction#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#data_source AppsyncFunction#data_source}.
	DataSource *string `json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#name AppsyncFunction#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#request_mapping_template AppsyncFunction#request_mapping_template}.
	RequestMappingTemplate *string `json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#response_mapping_template AppsyncFunction#response_mapping_template}.
	ResponseMappingTemplate *string `json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#description AppsyncFunction#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#function_version AppsyncFunction#function_version}.
	FunctionVersion *string `json:"functionVersion" yaml:"functionVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#max_batch_size AppsyncFunction#max_batch_size}.
	MaxBatchSize *float64 `json:"maxBatchSize" yaml:"maxBatchSize"`
	// sync_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#sync_config AppsyncFunction#sync_config}
	SyncConfig *AppsyncFunctionSyncConfig `json:"syncConfig" yaml:"syncConfig"`
}

type AppsyncFunctionSyncConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#conflict_detection AppsyncFunction#conflict_detection}.
	ConflictDetection *string `json:"conflictDetection" yaml:"conflictDetection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#conflict_handler AppsyncFunction#conflict_handler}.
	ConflictHandler *string `json:"conflictHandler" yaml:"conflictHandler"`
	// lambda_conflict_handler_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#lambda_conflict_handler_config AppsyncFunction#lambda_conflict_handler_config}
	LambdaConflictHandlerConfig *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig `json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

type AppsyncFunctionSyncConfigLambdaConflictHandlerConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_function#lambda_conflict_handler_arn AppsyncFunction#lambda_conflict_handler_arn}.
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn" yaml:"lambdaConflictHandlerArn"`
}

type AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig
	SetInternalValue(val *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LambdaConflictHandlerArn() *string
	SetLambdaConflictHandlerArn(val *string)
	LambdaConflictHandlerArnInput() *string
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
	ResetLambdaConflictHandlerArn()
}

// The jsii proxy struct for AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference
type jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) InternalValue() *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig {
	var returns *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) LambdaConflictHandlerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaConflictHandlerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) LambdaConflictHandlerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaConflictHandlerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference_Override(a AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) SetInternalValue(val *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) SetLambdaConflictHandlerArn(val *string) {
	_jsii_.Set(
		j,
		"lambdaConflictHandlerArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference) ResetLambdaConflictHandlerArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConflictHandlerArn",
		nil, // no parameters
	)
}

type AppsyncFunctionSyncConfigOutputReference interface {
	cdktf.ComplexObject
	ConflictDetection() *string
	SetConflictDetection(val *string)
	ConflictDetectionInput() *string
	ConflictHandler() *string
	SetConflictHandler(val *string)
	ConflictHandlerInput() *string
	InternalValue() *AppsyncFunctionSyncConfig
	SetInternalValue(val *AppsyncFunctionSyncConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LambdaConflictHandlerConfig() AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference
	LambdaConflictHandlerConfigInput() *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig
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
	PutLambdaConflictHandlerConfig(value *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig)
	ResetConflictDetection()
	ResetConflictHandler()
	ResetLambdaConflictHandlerConfig()
}

// The jsii proxy struct for AppsyncFunctionSyncConfigOutputReference
type jsiiProxy_AppsyncFunctionSyncConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) ConflictDetection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) ConflictDetectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) ConflictHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) ConflictHandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) InternalValue() *AppsyncFunctionSyncConfig {
	var returns *AppsyncFunctionSyncConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) LambdaConflictHandlerConfig() AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference {
	var returns AppsyncFunctionSyncConfigLambdaConflictHandlerConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaConflictHandlerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) LambdaConflictHandlerConfigInput() *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig {
	var returns *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig
	_jsii_.Get(
		j,
		"lambdaConflictHandlerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncFunctionSyncConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncFunctionSyncConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncFunctionSyncConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncFunctionSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncFunctionSyncConfigOutputReference_Override(a AppsyncFunctionSyncConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncFunctionSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) SetConflictDetection(val *string) {
	_jsii_.Set(
		j,
		"conflictDetection",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) SetConflictHandler(val *string) {
	_jsii_.Set(
		j,
		"conflictHandler",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) SetInternalValue(val *AppsyncFunctionSyncConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) PutLambdaConflictHandlerConfig(value *AppsyncFunctionSyncConfigLambdaConflictHandlerConfig) {
	_jsii_.InvokeVoid(
		a,
		"putLambdaConflictHandlerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) ResetConflictDetection() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictDetection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) ResetConflictHandler() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictHandler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncFunctionSyncConfigOutputReference) ResetLambdaConflictHandlerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConflictHandlerConfig",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api aws_appsync_graphql_api}.
type AppsyncGraphqlApi interface {
	cdktf.TerraformResource
	AdditionalAuthenticationProvider() interface{}
	SetAdditionalAuthenticationProvider(val interface{})
	AdditionalAuthenticationProviderInput() interface{}
	Arn() *string
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LambdaAuthorizerConfig() AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference
	LambdaAuthorizerConfigInput() *AppsyncGraphqlApiLambdaAuthorizerConfig
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogConfig() AppsyncGraphqlApiLogConfigOutputReference
	LogConfigInput() *AppsyncGraphqlApiLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	OpenidConnectConfig() AppsyncGraphqlApiOpenidConnectConfigOutputReference
	OpenidConnectConfigInput() *AppsyncGraphqlApiOpenidConnectConfig
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserPoolConfig() AppsyncGraphqlApiUserPoolConfigOutputReference
	UserPoolConfigInput() *AppsyncGraphqlApiUserPoolConfig
	XrayEnabled() interface{}
	SetXrayEnabled(val interface{})
	XrayEnabledInput() interface{}
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
	PutLambdaAuthorizerConfig(value *AppsyncGraphqlApiLambdaAuthorizerConfig)
	PutLogConfig(value *AppsyncGraphqlApiLogConfig)
	PutOpenidConnectConfig(value *AppsyncGraphqlApiOpenidConnectConfig)
	PutUserPoolConfig(value *AppsyncGraphqlApiUserPoolConfig)
	ResetAdditionalAuthenticationProvider()
	ResetLambdaAuthorizerConfig()
	ResetLogConfig()
	ResetOpenidConnectConfig()
	ResetOverrideLogicalId()
	ResetSchema()
	ResetTags()
	ResetTagsAll()
	ResetUserPoolConfig()
	ResetXrayEnabled()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
	Uris(key *string) interface{}
}

// The jsii proxy struct for AppsyncGraphqlApi
type jsiiProxy_AppsyncGraphqlApi struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncGraphqlApi) AdditionalAuthenticationProvider() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAuthenticationProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) AdditionalAuthenticationProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAuthenticationProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) LambdaAuthorizerConfig() AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference {
	var returns AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) LambdaAuthorizerConfigInput() *AppsyncGraphqlApiLambdaAuthorizerConfig {
	var returns *AppsyncGraphqlApiLambdaAuthorizerConfig
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) LogConfig() AppsyncGraphqlApiLogConfigOutputReference {
	var returns AppsyncGraphqlApiLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) LogConfigInput() *AppsyncGraphqlApiLogConfig {
	var returns *AppsyncGraphqlApiLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) OpenidConnectConfig() AppsyncGraphqlApiOpenidConnectConfigOutputReference {
	var returns AppsyncGraphqlApiOpenidConnectConfigOutputReference
	_jsii_.Get(
		j,
		"openidConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) OpenidConnectConfigInput() *AppsyncGraphqlApiOpenidConnectConfig {
	var returns *AppsyncGraphqlApiOpenidConnectConfig
	_jsii_.Get(
		j,
		"openidConnectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) UserPoolConfig() AppsyncGraphqlApiUserPoolConfigOutputReference {
	var returns AppsyncGraphqlApiUserPoolConfigOutputReference
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) UserPoolConfigInput() *AppsyncGraphqlApiUserPoolConfig {
	var returns *AppsyncGraphqlApiUserPoolConfig
	_jsii_.Get(
		j,
		"userPoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) XrayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApi) XrayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabledInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api aws_appsync_graphql_api} Resource.
func NewAppsyncGraphqlApi(scope constructs.Construct, id *string, config *AppsyncGraphqlApiConfig) AppsyncGraphqlApi {
	_init_.Initialize()

	j := jsiiProxy_AppsyncGraphqlApi{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api aws_appsync_graphql_api} Resource.
func NewAppsyncGraphqlApi_Override(a AppsyncGraphqlApi, scope constructs.Construct, id *string, config *AppsyncGraphqlApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApi",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetAdditionalAuthenticationProvider(val interface{}) {
	_jsii_.Set(
		j,
		"additionalAuthenticationProvider",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetAuthenticationType(val *string) {
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetSchema(val *string) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApi) SetXrayEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"xrayEnabled",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AppsyncGraphqlApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appsync.AppsyncGraphqlApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncGraphqlApi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appsync.AppsyncGraphqlApi",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) PutLambdaAuthorizerConfig(value *AppsyncGraphqlApiLambdaAuthorizerConfig) {
	_jsii_.InvokeVoid(
		a,
		"putLambdaAuthorizerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) PutLogConfig(value *AppsyncGraphqlApiLogConfig) {
	_jsii_.InvokeVoid(
		a,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) PutOpenidConnectConfig(value *AppsyncGraphqlApiOpenidConnectConfig) {
	_jsii_.InvokeVoid(
		a,
		"putOpenidConnectConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) PutUserPoolConfig(value *AppsyncGraphqlApiUserPoolConfig) {
	_jsii_.InvokeVoid(
		a,
		"putUserPoolConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetAdditionalAuthenticationProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalAuthenticationProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetLambdaAuthorizerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaAuthorizerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetLogConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetOpenidConnectConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenidConnectConfig",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetUserPoolConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPoolConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) ResetXrayEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetXrayEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppsyncGraphqlApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApi) Uris(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"uris",
		[]interface{}{key},
		&returns,
	)

	return returns
}

type AppsyncGraphqlApiAdditionalAuthenticationProvider struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#authentication_type AppsyncGraphqlApi#authentication_type}.
	AuthenticationType *string `json:"authenticationType" yaml:"authenticationType"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#lambda_authorizer_config AppsyncGraphqlApi#lambda_authorizer_config}
	LambdaAuthorizerConfig *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig `json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#openid_connect_config AppsyncGraphqlApi#openid_connect_config}
	OpenidConnectConfig *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig `json:"openidConnectConfig" yaml:"openidConnectConfig"`
	// user_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#user_pool_config AppsyncGraphqlApi#user_pool_config}
	UserPoolConfig *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig `json:"userPoolConfig" yaml:"userPoolConfig"`
}

type AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#authorizer_uri AppsyncGraphqlApi#authorizer_uri}.
	AuthorizerUri *string `json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#authorizer_result_ttl_in_seconds AppsyncGraphqlApi#authorizer_result_ttl_in_seconds}.
	AuthorizerResultTtlInSeconds *float64 `json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#identity_validation_expression AppsyncGraphqlApi#identity_validation_expression}.
	IdentityValidationExpression *string `json:"identityValidationExpression" yaml:"identityValidationExpression"`
}

type AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference interface {
	cdktf.ComplexObject
	AuthorizerResultTtlInSeconds() *float64
	SetAuthorizerResultTtlInSeconds(val *float64)
	AuthorizerResultTtlInSecondsInput() *float64
	AuthorizerUri() *string
	SetAuthorizerUri(val *string)
	AuthorizerUriInput() *string
	IdentityValidationExpression() *string
	SetIdentityValidationExpression(val *string)
	IdentityValidationExpressionInput() *string
	InternalValue() *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig
	SetInternalValue(val *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig)
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
	ResetAuthorizerResultTtlInSeconds()
	ResetIdentityValidationExpression()
}

// The jsii proxy struct for AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference
type jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) AuthorizerResultTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) AuthorizerResultTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) AuthorizerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) AuthorizerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) IdentityValidationExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) IdentityValidationExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) InternalValue() *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig {
	var returns *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference_Override(a AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) SetAuthorizerResultTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"authorizerResultTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) SetAuthorizerUri(val *string) {
	_jsii_.Set(
		j,
		"authorizerUri",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) SetIdentityValidationExpression(val *string) {
	_jsii_.Set(
		j,
		"identityValidationExpression",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) SetInternalValue(val *AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) ResetAuthorizerResultTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerResultTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderLambdaAuthorizerConfigOutputReference) ResetIdentityValidationExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityValidationExpression",
		nil, // no parameters
	)
}

type AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#issuer AppsyncGraphqlApi#issuer}.
	Issuer *string `json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#auth_ttl AppsyncGraphqlApi#auth_ttl}.
	AuthTtl *float64 `json:"authTtl" yaml:"authTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#client_id AppsyncGraphqlApi#client_id}.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#iat_ttl AppsyncGraphqlApi#iat_ttl}.
	IatTtl *float64 `json:"iatTtl" yaml:"iatTtl"`
}

type AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference interface {
	cdktf.ComplexObject
	AuthTtl() *float64
	SetAuthTtl(val *float64)
	AuthTtlInput() *float64
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	IatTtl() *float64
	SetIatTtl(val *float64)
	IatTtlInput() *float64
	InternalValue() *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig
	SetInternalValue(val *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
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
	ResetAuthTtl()
	ResetClientId()
	ResetIatTtl()
}

// The jsii proxy struct for AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference
type jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) AuthTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) AuthTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) IatTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iatTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) IatTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iatTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) InternalValue() *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig {
	var returns *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference_Override(a AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) SetAuthTtl(val *float64) {
	_jsii_.Set(
		j,
		"authTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) SetIatTtl(val *float64) {
	_jsii_.Set(
		j,
		"iatTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) SetInternalValue(val *AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) SetIssuer(val *string) {
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) ResetAuthTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderOpenidConnectConfigOutputReference) ResetIatTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetIatTtl",
		nil, // no parameters
	)
}

type AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#user_pool_id AppsyncGraphqlApi#user_pool_id}.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#app_id_client_regex AppsyncGraphqlApi#app_id_client_regex}.
	AppIdClientRegex *string `json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#aws_region AppsyncGraphqlApi#aws_region}.
	AwsRegion *string `json:"awsRegion" yaml:"awsRegion"`
}

type AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference interface {
	cdktf.ComplexObject
	AppIdClientRegex() *string
	SetAppIdClientRegex(val *string)
	AppIdClientRegexInput() *string
	AwsRegion() *string
	SetAwsRegion(val *string)
	AwsRegionInput() *string
	InternalValue() *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig
	SetInternalValue(val *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
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
	ResetAppIdClientRegex()
	ResetAwsRegion()
}

// The jsii proxy struct for AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference
type jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) AppIdClientRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdClientRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) AppIdClientRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdClientRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) InternalValue() *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig {
	var returns *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

func NewAppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference_Override(a AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) SetAppIdClientRegex(val *string) {
	_jsii_.Set(
		j,
		"appIdClientRegex",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) SetAwsRegion(val *string) {
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) SetInternalValue(val *AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) ResetAppIdClientRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetAppIdClientRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiAdditionalAuthenticationProviderUserPoolConfigOutputReference) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsRegion",
		nil, // no parameters
	)
}

// AWS AppSync.
type AppsyncGraphqlApiConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#authentication_type AppsyncGraphqlApi#authentication_type}.
	AuthenticationType *string `json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#name AppsyncGraphqlApi#name}.
	Name *string `json:"name" yaml:"name"`
	// additional_authentication_provider block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#additional_authentication_provider AppsyncGraphqlApi#additional_authentication_provider}
	AdditionalAuthenticationProvider interface{} `json:"additionalAuthenticationProvider" yaml:"additionalAuthenticationProvider"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#lambda_authorizer_config AppsyncGraphqlApi#lambda_authorizer_config}
	LambdaAuthorizerConfig *AppsyncGraphqlApiLambdaAuthorizerConfig `json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#log_config AppsyncGraphqlApi#log_config}
	LogConfig *AppsyncGraphqlApiLogConfig `json:"logConfig" yaml:"logConfig"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#openid_connect_config AppsyncGraphqlApi#openid_connect_config}
	OpenidConnectConfig *AppsyncGraphqlApiOpenidConnectConfig `json:"openidConnectConfig" yaml:"openidConnectConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#schema AppsyncGraphqlApi#schema}.
	Schema *string `json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#tags AppsyncGraphqlApi#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#tags_all AppsyncGraphqlApi#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// user_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#user_pool_config AppsyncGraphqlApi#user_pool_config}
	UserPoolConfig *AppsyncGraphqlApiUserPoolConfig `json:"userPoolConfig" yaml:"userPoolConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#xray_enabled AppsyncGraphqlApi#xray_enabled}.
	XrayEnabled interface{} `json:"xrayEnabled" yaml:"xrayEnabled"`
}

type AppsyncGraphqlApiLambdaAuthorizerConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#authorizer_uri AppsyncGraphqlApi#authorizer_uri}.
	AuthorizerUri *string `json:"authorizerUri" yaml:"authorizerUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#authorizer_result_ttl_in_seconds AppsyncGraphqlApi#authorizer_result_ttl_in_seconds}.
	AuthorizerResultTtlInSeconds *float64 `json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#identity_validation_expression AppsyncGraphqlApi#identity_validation_expression}.
	IdentityValidationExpression *string `json:"identityValidationExpression" yaml:"identityValidationExpression"`
}

type AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference interface {
	cdktf.ComplexObject
	AuthorizerResultTtlInSeconds() *float64
	SetAuthorizerResultTtlInSeconds(val *float64)
	AuthorizerResultTtlInSecondsInput() *float64
	AuthorizerUri() *string
	SetAuthorizerUri(val *string)
	AuthorizerUriInput() *string
	IdentityValidationExpression() *string
	SetIdentityValidationExpression(val *string)
	IdentityValidationExpressionInput() *string
	InternalValue() *AppsyncGraphqlApiLambdaAuthorizerConfig
	SetInternalValue(val *AppsyncGraphqlApiLambdaAuthorizerConfig)
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
	ResetAuthorizerResultTtlInSeconds()
	ResetIdentityValidationExpression()
}

// The jsii proxy struct for AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference
type jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) AuthorizerResultTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) AuthorizerResultTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) AuthorizerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) AuthorizerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) IdentityValidationExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) IdentityValidationExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) InternalValue() *AppsyncGraphqlApiLambdaAuthorizerConfig {
	var returns *AppsyncGraphqlApiLambdaAuthorizerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncGraphqlApiLambdaAuthorizerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncGraphqlApiLambdaAuthorizerConfigOutputReference_Override(a AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) SetAuthorizerResultTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"authorizerResultTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) SetAuthorizerUri(val *string) {
	_jsii_.Set(
		j,
		"authorizerUri",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) SetIdentityValidationExpression(val *string) {
	_jsii_.Set(
		j,
		"identityValidationExpression",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) SetInternalValue(val *AppsyncGraphqlApiLambdaAuthorizerConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) ResetAuthorizerResultTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerResultTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiLambdaAuthorizerConfigOutputReference) ResetIdentityValidationExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityValidationExpression",
		nil, // no parameters
	)
}

type AppsyncGraphqlApiLogConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#cloudwatch_logs_role_arn AppsyncGraphqlApi#cloudwatch_logs_role_arn}.
	CloudwatchLogsRoleArn *string `json:"cloudwatchLogsRoleArn" yaml:"cloudwatchLogsRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#field_log_level AppsyncGraphqlApi#field_log_level}.
	FieldLogLevel *string `json:"fieldLogLevel" yaml:"fieldLogLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#exclude_verbose_content AppsyncGraphqlApi#exclude_verbose_content}.
	ExcludeVerboseContent interface{} `json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
}

type AppsyncGraphqlApiLogConfigOutputReference interface {
	cdktf.ComplexObject
	CloudwatchLogsRoleArn() *string
	SetCloudwatchLogsRoleArn(val *string)
	CloudwatchLogsRoleArnInput() *string
	ExcludeVerboseContent() interface{}
	SetExcludeVerboseContent(val interface{})
	ExcludeVerboseContentInput() interface{}
	FieldLogLevel() *string
	SetFieldLogLevel(val *string)
	FieldLogLevelInput() *string
	InternalValue() *AppsyncGraphqlApiLogConfig
	SetInternalValue(val *AppsyncGraphqlApiLogConfig)
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
	ResetExcludeVerboseContent()
}

// The jsii proxy struct for AppsyncGraphqlApiLogConfigOutputReference
type jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) CloudwatchLogsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) CloudwatchLogsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) ExcludeVerboseContent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeVerboseContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) ExcludeVerboseContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeVerboseContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) FieldLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) FieldLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) InternalValue() *AppsyncGraphqlApiLogConfig {
	var returns *AppsyncGraphqlApiLogConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncGraphqlApiLogConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncGraphqlApiLogConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiLogConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncGraphqlApiLogConfigOutputReference_Override(a AppsyncGraphqlApiLogConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiLogConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) SetCloudwatchLogsRoleArn(val *string) {
	_jsii_.Set(
		j,
		"cloudwatchLogsRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) SetExcludeVerboseContent(val interface{}) {
	_jsii_.Set(
		j,
		"excludeVerboseContent",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) SetFieldLogLevel(val *string) {
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) SetInternalValue(val *AppsyncGraphqlApiLogConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiLogConfigOutputReference) ResetExcludeVerboseContent() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludeVerboseContent",
		nil, // no parameters
	)
}

type AppsyncGraphqlApiOpenidConnectConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#issuer AppsyncGraphqlApi#issuer}.
	Issuer *string `json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#auth_ttl AppsyncGraphqlApi#auth_ttl}.
	AuthTtl *float64 `json:"authTtl" yaml:"authTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#client_id AppsyncGraphqlApi#client_id}.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#iat_ttl AppsyncGraphqlApi#iat_ttl}.
	IatTtl *float64 `json:"iatTtl" yaml:"iatTtl"`
}

type AppsyncGraphqlApiOpenidConnectConfigOutputReference interface {
	cdktf.ComplexObject
	AuthTtl() *float64
	SetAuthTtl(val *float64)
	AuthTtlInput() *float64
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	IatTtl() *float64
	SetIatTtl(val *float64)
	IatTtlInput() *float64
	InternalValue() *AppsyncGraphqlApiOpenidConnectConfig
	SetInternalValue(val *AppsyncGraphqlApiOpenidConnectConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
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
	ResetAuthTtl()
	ResetClientId()
	ResetIatTtl()
}

// The jsii proxy struct for AppsyncGraphqlApiOpenidConnectConfigOutputReference
type jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) AuthTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) AuthTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) IatTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iatTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) IatTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iatTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) InternalValue() *AppsyncGraphqlApiOpenidConnectConfig {
	var returns *AppsyncGraphqlApiOpenidConnectConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncGraphqlApiOpenidConnectConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncGraphqlApiOpenidConnectConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiOpenidConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncGraphqlApiOpenidConnectConfigOutputReference_Override(a AppsyncGraphqlApiOpenidConnectConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiOpenidConnectConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) SetAuthTtl(val *float64) {
	_jsii_.Set(
		j,
		"authTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) SetIatTtl(val *float64) {
	_jsii_.Set(
		j,
		"iatTtl",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) SetInternalValue(val *AppsyncGraphqlApiOpenidConnectConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) SetIssuer(val *string) {
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) ResetAuthTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiOpenidConnectConfigOutputReference) ResetIatTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetIatTtl",
		nil, // no parameters
	)
}

type AppsyncGraphqlApiUserPoolConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#default_action AppsyncGraphqlApi#default_action}.
	DefaultAction *string `json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#user_pool_id AppsyncGraphqlApi#user_pool_id}.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#app_id_client_regex AppsyncGraphqlApi#app_id_client_regex}.
	AppIdClientRegex *string `json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_graphql_api#aws_region AppsyncGraphqlApi#aws_region}.
	AwsRegion *string `json:"awsRegion" yaml:"awsRegion"`
}

type AppsyncGraphqlApiUserPoolConfigOutputReference interface {
	cdktf.ComplexObject
	AppIdClientRegex() *string
	SetAppIdClientRegex(val *string)
	AppIdClientRegexInput() *string
	AwsRegion() *string
	SetAwsRegion(val *string)
	AwsRegionInput() *string
	DefaultAction() *string
	SetDefaultAction(val *string)
	DefaultActionInput() *string
	InternalValue() *AppsyncGraphqlApiUserPoolConfig
	SetInternalValue(val *AppsyncGraphqlApiUserPoolConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
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
	ResetAppIdClientRegex()
	ResetAwsRegion()
}

// The jsii proxy struct for AppsyncGraphqlApiUserPoolConfigOutputReference
type jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) AppIdClientRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdClientRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) AppIdClientRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appIdClientRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) AwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) AwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) DefaultAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) DefaultActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) InternalValue() *AppsyncGraphqlApiUserPoolConfig {
	var returns *AppsyncGraphqlApiUserPoolConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

func NewAppsyncGraphqlApiUserPoolConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncGraphqlApiUserPoolConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiUserPoolConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncGraphqlApiUserPoolConfigOutputReference_Override(a AppsyncGraphqlApiUserPoolConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncGraphqlApiUserPoolConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) SetAppIdClientRegex(val *string) {
	_jsii_.Set(
		j,
		"appIdClientRegex",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) SetAwsRegion(val *string) {
	_jsii_.Set(
		j,
		"awsRegion",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) SetDefaultAction(val *string) {
	_jsii_.Set(
		j,
		"defaultAction",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) SetInternalValue(val *AppsyncGraphqlApiUserPoolConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) ResetAppIdClientRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetAppIdClientRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncGraphqlApiUserPoolConfigOutputReference) ResetAwsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsRegion",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver aws_appsync_resolver}.
type AppsyncResolver interface {
	cdktf.TerraformResource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	Arn() *string
	CachingConfig() AppsyncResolverCachingConfigOutputReference
	CachingConfigInput() *AppsyncResolverCachingConfig
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DataSource() *string
	SetDataSource(val *string)
	DataSourceInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Field() *string
	SetField(val *string)
	FieldInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxBatchSize() *float64
	SetMaxBatchSize(val *float64)
	MaxBatchSizeInput() *float64
	Node() constructs.Node
	PipelineConfig() AppsyncResolverPipelineConfigOutputReference
	PipelineConfigInput() *AppsyncResolverPipelineConfig
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequestTemplate() *string
	SetRequestTemplate(val *string)
	RequestTemplateInput() *string
	ResponseTemplate() *string
	SetResponseTemplate(val *string)
	ResponseTemplateInput() *string
	SyncConfig() AppsyncResolverSyncConfigOutputReference
	SyncConfigInput() *AppsyncResolverSyncConfig
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
	PutCachingConfig(value *AppsyncResolverCachingConfig)
	PutPipelineConfig(value *AppsyncResolverPipelineConfig)
	PutSyncConfig(value *AppsyncResolverSyncConfig)
	ResetCachingConfig()
	ResetDataSource()
	ResetKind()
	ResetMaxBatchSize()
	ResetOverrideLogicalId()
	ResetPipelineConfig()
	ResetRequestTemplate()
	ResetResponseTemplate()
	ResetSyncConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppsyncResolver
type jsiiProxy_AppsyncResolver struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppsyncResolver) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) CachingConfig() AppsyncResolverCachingConfigOutputReference {
	var returns AppsyncResolverCachingConfigOutputReference
	_jsii_.Get(
		j,
		"cachingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) CachingConfigInput() *AppsyncResolverCachingConfig {
	var returns *AppsyncResolverCachingConfig
	_jsii_.Get(
		j,
		"cachingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) MaxBatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) MaxBatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBatchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) PipelineConfig() AppsyncResolverPipelineConfigOutputReference {
	var returns AppsyncResolverPipelineConfigOutputReference
	_jsii_.Get(
		j,
		"pipelineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) PipelineConfigInput() *AppsyncResolverPipelineConfig {
	var returns *AppsyncResolverPipelineConfig
	_jsii_.Get(
		j,
		"pipelineConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) RequestTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) RequestTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) ResponseTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) ResponseTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) SyncConfig() AppsyncResolverSyncConfigOutputReference {
	var returns AppsyncResolverSyncConfigOutputReference
	_jsii_.Get(
		j,
		"syncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) SyncConfigInput() *AppsyncResolverSyncConfig {
	var returns *AppsyncResolverSyncConfig
	_jsii_.Get(
		j,
		"syncConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolver) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver aws_appsync_resolver} Resource.
func NewAppsyncResolver(scope constructs.Construct, id *string, config *AppsyncResolverConfig) AppsyncResolver {
	_init_.Initialize()

	j := jsiiProxy_AppsyncResolver{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolver",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver aws_appsync_resolver} Resource.
func NewAppsyncResolver_Override(a AppsyncResolver, scope constructs.Construct, id *string, config *AppsyncResolverConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolver",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetDataSource(val *string) {
	_jsii_.Set(
		j,
		"dataSource",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetField(val *string) {
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetKind(val *string) {
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetMaxBatchSize(val *float64) {
	_jsii_.Set(
		j,
		"maxBatchSize",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetRequestTemplate(val *string) {
	_jsii_.Set(
		j,
		"requestTemplate",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetResponseTemplate(val *string) {
	_jsii_.Set(
		j,
		"responseTemplate",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolver) SetType(val *string) {
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
func AppsyncResolver_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appsync.AppsyncResolver",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppsyncResolver_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appsync.AppsyncResolver",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AppsyncResolver) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppsyncResolver) PutCachingConfig(value *AppsyncResolverCachingConfig) {
	_jsii_.InvokeVoid(
		a,
		"putCachingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncResolver) PutPipelineConfig(value *AppsyncResolverPipelineConfig) {
	_jsii_.InvokeVoid(
		a,
		"putPipelineConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncResolver) PutSyncConfig(value *AppsyncResolverSyncConfig) {
	_jsii_.InvokeVoid(
		a,
		"putSyncConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetCachingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCachingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetDataSource() {
	_jsii_.InvokeVoid(
		a,
		"resetDataSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetKind() {
	_jsii_.InvokeVoid(
		a,
		"resetKind",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetMaxBatchSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxBatchSize",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppsyncResolver) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetPipelineConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPipelineConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetRequestTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetResponseTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) ResetSyncConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolver) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolver) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_AppsyncResolver) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (a *jsiiProxy_AppsyncResolver) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppsyncResolverCachingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#caching_keys AppsyncResolver#caching_keys}.
	CachingKeys *[]*string `json:"cachingKeys" yaml:"cachingKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#ttl AppsyncResolver#ttl}.
	Ttl *float64 `json:"ttl" yaml:"ttl"`
}

type AppsyncResolverCachingConfigOutputReference interface {
	cdktf.ComplexObject
	CachingKeys() *[]*string
	SetCachingKeys(val *[]*string)
	CachingKeysInput() *[]*string
	InternalValue() *AppsyncResolverCachingConfig
	SetInternalValue(val *AppsyncResolverCachingConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
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
	ResetCachingKeys()
	ResetTtl()
}

// The jsii proxy struct for AppsyncResolverCachingConfigOutputReference
type jsiiProxy_AppsyncResolverCachingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) CachingKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachingKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) CachingKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachingKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) InternalValue() *AppsyncResolverCachingConfig {
	var returns *AppsyncResolverCachingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func NewAppsyncResolverCachingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncResolverCachingConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncResolverCachingConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolverCachingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncResolverCachingConfigOutputReference_Override(a AppsyncResolverCachingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolverCachingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) SetCachingKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"cachingKeys",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) SetInternalValue(val *AppsyncResolverCachingConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverCachingConfigOutputReference) SetTtl(val *float64) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) ResetCachingKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetCachingKeys",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolverCachingConfigOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTtl",
		nil, // no parameters
	)
}

// AWS AppSync.
type AppsyncResolverConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#api_id AppsyncResolver#api_id}.
	ApiId *string `json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#field AppsyncResolver#field}.
	Field *string `json:"field" yaml:"field"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#type AppsyncResolver#type}.
	Type *string `json:"type" yaml:"type"`
	// caching_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#caching_config AppsyncResolver#caching_config}
	CachingConfig *AppsyncResolverCachingConfig `json:"cachingConfig" yaml:"cachingConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#data_source AppsyncResolver#data_source}.
	DataSource *string `json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#kind AppsyncResolver#kind}.
	Kind *string `json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#max_batch_size AppsyncResolver#max_batch_size}.
	MaxBatchSize *float64 `json:"maxBatchSize" yaml:"maxBatchSize"`
	// pipeline_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#pipeline_config AppsyncResolver#pipeline_config}
	PipelineConfig *AppsyncResolverPipelineConfig `json:"pipelineConfig" yaml:"pipelineConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#request_template AppsyncResolver#request_template}.
	RequestTemplate *string `json:"requestTemplate" yaml:"requestTemplate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#response_template AppsyncResolver#response_template}.
	ResponseTemplate *string `json:"responseTemplate" yaml:"responseTemplate"`
	// sync_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#sync_config AppsyncResolver#sync_config}
	SyncConfig *AppsyncResolverSyncConfig `json:"syncConfig" yaml:"syncConfig"`
}

type AppsyncResolverPipelineConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#functions AppsyncResolver#functions}.
	Functions *[]*string `json:"functions" yaml:"functions"`
}

type AppsyncResolverPipelineConfigOutputReference interface {
	cdktf.ComplexObject
	Functions() *[]*string
	SetFunctions(val *[]*string)
	FunctionsInput() *[]*string
	InternalValue() *AppsyncResolverPipelineConfig
	SetInternalValue(val *AppsyncResolverPipelineConfig)
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
	ResetFunctions()
}

// The jsii proxy struct for AppsyncResolverPipelineConfigOutputReference
type jsiiProxy_AppsyncResolverPipelineConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) Functions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) FunctionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) InternalValue() *AppsyncResolverPipelineConfig {
	var returns *AppsyncResolverPipelineConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncResolverPipelineConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncResolverPipelineConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncResolverPipelineConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolverPipelineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncResolverPipelineConfigOutputReference_Override(a AppsyncResolverPipelineConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolverPipelineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) SetFunctions(val *[]*string) {
	_jsii_.Set(
		j,
		"functions",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) SetInternalValue(val *AppsyncResolverPipelineConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolverPipelineConfigOutputReference) ResetFunctions() {
	_jsii_.InvokeVoid(
		a,
		"resetFunctions",
		nil, // no parameters
	)
}

type AppsyncResolverSyncConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#conflict_detection AppsyncResolver#conflict_detection}.
	ConflictDetection *string `json:"conflictDetection" yaml:"conflictDetection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#conflict_handler AppsyncResolver#conflict_handler}.
	ConflictHandler *string `json:"conflictHandler" yaml:"conflictHandler"`
	// lambda_conflict_handler_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#lambda_conflict_handler_config AppsyncResolver#lambda_conflict_handler_config}
	LambdaConflictHandlerConfig *AppsyncResolverSyncConfigLambdaConflictHandlerConfig `json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

type AppsyncResolverSyncConfigLambdaConflictHandlerConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appsync_resolver#lambda_conflict_handler_arn AppsyncResolver#lambda_conflict_handler_arn}.
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn" yaml:"lambdaConflictHandlerArn"`
}

type AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *AppsyncResolverSyncConfigLambdaConflictHandlerConfig
	SetInternalValue(val *AppsyncResolverSyncConfigLambdaConflictHandlerConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LambdaConflictHandlerArn() *string
	SetLambdaConflictHandlerArn(val *string)
	LambdaConflictHandlerArnInput() *string
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
	ResetLambdaConflictHandlerArn()
}

// The jsii proxy struct for AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference
type jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) InternalValue() *AppsyncResolverSyncConfigLambdaConflictHandlerConfig {
	var returns *AppsyncResolverSyncConfigLambdaConflictHandlerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) LambdaConflictHandlerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaConflictHandlerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) LambdaConflictHandlerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaConflictHandlerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference_Override(a AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) SetInternalValue(val *AppsyncResolverSyncConfigLambdaConflictHandlerConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) SetLambdaConflictHandlerArn(val *string) {
	_jsii_.Set(
		j,
		"lambdaConflictHandlerArn",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference) ResetLambdaConflictHandlerArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConflictHandlerArn",
		nil, // no parameters
	)
}

type AppsyncResolverSyncConfigOutputReference interface {
	cdktf.ComplexObject
	ConflictDetection() *string
	SetConflictDetection(val *string)
	ConflictDetectionInput() *string
	ConflictHandler() *string
	SetConflictHandler(val *string)
	ConflictHandlerInput() *string
	InternalValue() *AppsyncResolverSyncConfig
	SetInternalValue(val *AppsyncResolverSyncConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LambdaConflictHandlerConfig() AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference
	LambdaConflictHandlerConfigInput() *AppsyncResolverSyncConfigLambdaConflictHandlerConfig
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
	PutLambdaConflictHandlerConfig(value *AppsyncResolverSyncConfigLambdaConflictHandlerConfig)
	ResetConflictDetection()
	ResetConflictHandler()
	ResetLambdaConflictHandlerConfig()
}

// The jsii proxy struct for AppsyncResolverSyncConfigOutputReference
type jsiiProxy_AppsyncResolverSyncConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ConflictDetection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ConflictDetectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ConflictHandler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ConflictHandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conflictHandlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) InternalValue() *AppsyncResolverSyncConfig {
	var returns *AppsyncResolverSyncConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) LambdaConflictHandlerConfig() AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference {
	var returns AppsyncResolverSyncConfigLambdaConflictHandlerConfigOutputReference
	_jsii_.Get(
		j,
		"lambdaConflictHandlerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) LambdaConflictHandlerConfigInput() *AppsyncResolverSyncConfigLambdaConflictHandlerConfig {
	var returns *AppsyncResolverSyncConfigLambdaConflictHandlerConfig
	_jsii_.Get(
		j,
		"lambdaConflictHandlerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppsyncResolverSyncConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppsyncResolverSyncConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppsyncResolverSyncConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolverSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppsyncResolverSyncConfigOutputReference_Override(a AppsyncResolverSyncConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appsync.AppsyncResolverSyncConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) SetConflictDetection(val *string) {
	_jsii_.Set(
		j,
		"conflictDetection",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) SetConflictHandler(val *string) {
	_jsii_.Set(
		j,
		"conflictHandler",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) SetInternalValue(val *AppsyncResolverSyncConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppsyncResolverSyncConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) PutLambdaConflictHandlerConfig(value *AppsyncResolverSyncConfigLambdaConflictHandlerConfig) {
	_jsii_.InvokeVoid(
		a,
		"putLambdaConflictHandlerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ResetConflictDetection() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictDetection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ResetConflictHandler() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictHandler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppsyncResolverSyncConfigOutputReference) ResetLambdaConflictHandlerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaConflictHandlerConfig",
		nil, // no parameters
	)
}
