package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/cloudfront/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html aws_cloudfront_cache_policy}.
type CloudfrontCachePolicy interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxTtl() *float64
	SetMaxTtl(val *float64)
	MaxTtlInput() *float64
	MinTtl() *float64
	SetMinTtl(val *float64)
	MinTtlInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	ParametersInCacheKeyAndForwardedToOrigin() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference
	ParametersInCacheKeyAndForwardedToOriginInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin
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
	PutParametersInCacheKeyAndForwardedToOrigin(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin)
	ResetComment()
	ResetDefaultTtl()
	ResetEtag()
	ResetMaxTtl()
	ResetMinTtl()
	ResetOverrideLogicalId()
	ResetParametersInCacheKeyAndForwardedToOrigin()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfrontCachePolicy
type jsiiProxy_CloudfrontCachePolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontCachePolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) MinTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) MinTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) ParametersInCacheKeyAndForwardedToOrigin() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference {
	var returns CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference
	_jsii_.Get(
		j,
		"parametersInCacheKeyAndForwardedToOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) ParametersInCacheKeyAndForwardedToOriginInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin
	_jsii_.Get(
		j,
		"parametersInCacheKeyAndForwardedToOriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html aws_cloudfront_cache_policy} Resource.
func NewCloudfrontCachePolicy(scope constructs.Construct, id *string, config *CloudfrontCachePolicyConfig) CloudfrontCachePolicy {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontCachePolicy{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html aws_cloudfront_cache_policy} Resource.
func NewCloudfrontCachePolicy_Override(c CloudfrontCachePolicy, scope constructs.Construct, id *string, config *CloudfrontCachePolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicy",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetDefaultTtl(val *float64) {
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetEtag(val *string) {
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetMaxTtl(val *float64) {
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetMinTtl(val *float64) {
	_jsii_.Set(
		j,
		"minTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicy) SetProvider(val cdktf.TerraformProvider) {
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
func CloudfrontCachePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontCachePolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontCachePolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontCachePolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontCachePolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontCachePolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontCachePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicy) PutParametersInCacheKeyAndForwardedToOrigin(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) {
	_jsii_.InvokeVoid(
		c,
		"putParametersInCacheKeyAndForwardedToOrigin",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicy) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicy) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicy) ResetEtag() {
	_jsii_.InvokeVoid(
		c,
		"resetEtag",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicy) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicy) ResetMinTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMinTtl",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicy) ResetParametersInCacheKeyAndForwardedToOrigin() {
	_jsii_.InvokeVoid(
		c,
		"resetParametersInCacheKeyAndForwardedToOrigin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontCachePolicy) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontCachePolicy) ToString() *string {
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
func (c *jsiiProxy_CloudfrontCachePolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudfrontCachePolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#name CloudfrontCachePolicy#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#comment CloudfrontCachePolicy#comment}.
	Comment *string `json:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#default_ttl CloudfrontCachePolicy#default_ttl}.
	DefaultTtl *float64 `json:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#etag CloudfrontCachePolicy#etag}.
	Etag *string `json:"etag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#max_ttl CloudfrontCachePolicy#max_ttl}.
	MaxTtl *float64 `json:"maxTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#min_ttl CloudfrontCachePolicy#min_ttl}.
	MinTtl *float64 `json:"minTtl"`
	// parameters_in_cache_key_and_forwarded_to_origin block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#parameters_in_cache_key_and_forwarded_to_origin CloudfrontCachePolicy#parameters_in_cache_key_and_forwarded_to_origin}
	ParametersInCacheKeyAndForwardedToOrigin *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin `json:"parametersInCacheKeyAndForwardedToOrigin"`
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin struct {
	// cookies_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#cookies_config CloudfrontCachePolicy#cookies_config}
	CookiesConfig *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig `json:"cookiesConfig"`
	// headers_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#headers_config CloudfrontCachePolicy#headers_config}
	HeadersConfig *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig `json:"headersConfig"`
	// query_strings_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#query_strings_config CloudfrontCachePolicy#query_strings_config}
	QueryStringsConfig *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig `json:"queryStringsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#enable_accept_encoding_brotli CloudfrontCachePolicy#enable_accept_encoding_brotli}.
	EnableAcceptEncodingBrotli interface{} `json:"enableAcceptEncodingBrotli"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#enable_accept_encoding_gzip CloudfrontCachePolicy#enable_accept_encoding_gzip}.
	EnableAcceptEncodingGzip interface{} `json:"enableAcceptEncodingGzip"`
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#cookie_behavior CloudfrontCachePolicy#cookie_behavior}.
	CookieBehavior *string `json:"cookieBehavior"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#cookies CloudfrontCachePolicy#cookies}
	Cookies *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies `json:"cookies"`
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#items CloudfrontCachePolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference
type jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference_Override(c CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference interface {
	cdktf.ComplexObject
	CookieBehavior() *string
	SetCookieBehavior(val *string)
	CookieBehaviorInput() *string
	Cookies() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference
	CookiesInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies
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
	PutCookies(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies)
	ResetCookies()
}

// The jsii proxy struct for CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference
type jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) CookieBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) CookieBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) Cookies() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference {
	var returns CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) CookiesInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference_Override(c CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) SetCookieBehavior(val *string) {
	_jsii_.Set(
		j,
		"cookieBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) PutCookies(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) {
	_jsii_.InvokeVoid(
		c,
		"putCookies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		c,
		"resetCookies",
		nil, // no parameters
	)
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#header_behavior CloudfrontCachePolicy#header_behavior}.
	HeaderBehavior *string `json:"headerBehavior"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#headers CloudfrontCachePolicy#headers}
	Headers *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders `json:"headers"`
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#items CloudfrontCachePolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference
type jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference_Override(c CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference interface {
	cdktf.ComplexObject
	HeaderBehavior() *string
	SetHeaderBehavior(val *string)
	HeaderBehaviorInput() *string
	Headers() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference
	HeadersInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders
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
	PutHeaders(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders)
	ResetHeaderBehavior()
	ResetHeaders()
}

// The jsii proxy struct for CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference
type jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) HeaderBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) HeaderBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) Headers() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference {
	var returns CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) HeadersInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference_Override(c CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) SetHeaderBehavior(val *string) {
	_jsii_.Set(
		j,
		"headerBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) PutHeaders(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) {
	_jsii_.InvokeVoid(
		c,
		"putHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) ResetHeaderBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaderBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaders",
		nil, // no parameters
	)
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference interface {
	cdktf.ComplexObject
	CookiesConfig() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference
	CookiesConfigInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig
	EnableAcceptEncodingBrotli() interface{}
	SetEnableAcceptEncodingBrotli(val interface{})
	EnableAcceptEncodingBrotliInput() interface{}
	EnableAcceptEncodingGzip() interface{}
	SetEnableAcceptEncodingGzip(val interface{})
	EnableAcceptEncodingGzipInput() interface{}
	HeadersConfig() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference
	HeadersConfigInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	QueryStringsConfig() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference
	QueryStringsConfigInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig
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
	PutCookiesConfig(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig)
	PutHeadersConfig(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig)
	PutQueryStringsConfig(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig)
	ResetEnableAcceptEncodingBrotli()
	ResetEnableAcceptEncodingGzip()
}

// The jsii proxy struct for CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference
type jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) CookiesConfig() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference {
	var returns CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference
	_jsii_.Get(
		j,
		"cookiesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) CookiesConfigInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig
	_jsii_.Get(
		j,
		"cookiesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) EnableAcceptEncodingBrotli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) EnableAcceptEncodingBrotliInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) EnableAcceptEncodingGzip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) EnableAcceptEncodingGzipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) HeadersConfig() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference {
	var returns CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference
	_jsii_.Get(
		j,
		"headersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) HeadersConfigInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig
	_jsii_.Get(
		j,
		"headersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) QueryStringsConfig() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference {
	var returns CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference
	_jsii_.Get(
		j,
		"queryStringsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) QueryStringsConfigInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig
	_jsii_.Get(
		j,
		"queryStringsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference_Override(c CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) SetEnableAcceptEncodingBrotli(val interface{}) {
	_jsii_.Set(
		j,
		"enableAcceptEncodingBrotli",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) SetEnableAcceptEncodingGzip(val interface{}) {
	_jsii_.Set(
		j,
		"enableAcceptEncodingGzip",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) PutCookiesConfig(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) {
	_jsii_.InvokeVoid(
		c,
		"putCookiesConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) PutHeadersConfig(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) {
	_jsii_.InvokeVoid(
		c,
		"putHeadersConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) PutQueryStringsConfig(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) {
	_jsii_.InvokeVoid(
		c,
		"putQueryStringsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) ResetEnableAcceptEncodingBrotli() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableAcceptEncodingBrotli",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference) ResetEnableAcceptEncodingGzip() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableAcceptEncodingGzip",
		nil, // no parameters
	)
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#query_string_behavior CloudfrontCachePolicy#query_string_behavior}.
	QueryStringBehavior *string `json:"queryStringBehavior"`
	// query_strings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#query_strings CloudfrontCachePolicy#query_strings}
	QueryStrings *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings `json:"queryStrings"`
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	QueryStringBehavior() *string
	SetQueryStringBehavior(val *string)
	QueryStringBehaviorInput() *string
	QueryStrings() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference
	QueryStringsInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings
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
	PutQueryStrings(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings)
	ResetQueryStrings()
}

// The jsii proxy struct for CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference
type jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) QueryStringBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) QueryStringBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) QueryStrings() CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference {
	var returns CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) QueryStringsInput() *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings {
	var returns *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings
	_jsii_.Get(
		j,
		"queryStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference_Override(c CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) SetQueryStringBehavior(val *string) {
	_jsii_.Set(
		j,
		"queryStringBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) PutQueryStrings(value *CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) {
	_jsii_.InvokeVoid(
		c,
		"putQueryStrings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference) ResetQueryStrings() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStrings",
		nil, // no parameters
	)
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_cache_policy.html#items CloudfrontCachePolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference
type jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference_Override(c CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html aws_cloudfront_distribution}.
type CloudfrontDistribution interface {
	cdktf.TerraformResource
	Aliases() *[]*string
	SetAliases(val *[]*string)
	AliasesInput() *[]*string
	Arn() *string
	CallerReference() *string
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	CustomErrorResponse() *[]*CloudfrontDistributionCustomErrorResponse
	SetCustomErrorResponse(val *[]*CloudfrontDistributionCustomErrorResponse)
	CustomErrorResponseInput() *[]*CloudfrontDistributionCustomErrorResponse
	DefaultCacheBehavior() CloudfrontDistributionDefaultCacheBehaviorOutputReference
	DefaultCacheBehaviorInput() *CloudfrontDistributionDefaultCacheBehavior
	DefaultRootObject() *string
	SetDefaultRootObject(val *string)
	DefaultRootObjectInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DomainName() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Etag() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HostedZoneId() *string
	HttpVersion() *string
	SetHttpVersion(val *string)
	HttpVersionInput() *string
	Id() *string
	InProgressValidationBatches() *float64
	IsIpv6Enabled() interface{}
	SetIsIpv6Enabled(val interface{})
	IsIpv6EnabledInput() interface{}
	LastModifiedTime() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingConfig() CloudfrontDistributionLoggingConfigOutputReference
	LoggingConfigInput() *CloudfrontDistributionLoggingConfig
	Node() constructs.Node
	OrderedCacheBehavior() *[]*CloudfrontDistributionOrderedCacheBehavior
	SetOrderedCacheBehavior(val *[]*CloudfrontDistributionOrderedCacheBehavior)
	OrderedCacheBehaviorInput() *[]*CloudfrontDistributionOrderedCacheBehavior
	Origin() *[]*CloudfrontDistributionOrigin
	SetOrigin(val *[]*CloudfrontDistributionOrigin)
	OriginGroup() *[]*CloudfrontDistributionOriginGroup
	SetOriginGroup(val *[]*CloudfrontDistributionOriginGroup)
	OriginGroupInput() *[]*CloudfrontDistributionOriginGroup
	OriginInput() *[]*CloudfrontDistributionOrigin
	PriceClass() *string
	SetPriceClass(val *string)
	PriceClassInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Restrictions() CloudfrontDistributionRestrictionsOutputReference
	RestrictionsInput() *CloudfrontDistributionRestrictions
	RetainOnDelete() interface{}
	SetRetainOnDelete(val interface{})
	RetainOnDeleteInput() interface{}
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
	ViewerCertificate() CloudfrontDistributionViewerCertificateOutputReference
	ViewerCertificateInput() *CloudfrontDistributionViewerCertificate
	WaitForDeployment() interface{}
	SetWaitForDeployment(val interface{})
	WaitForDeploymentInput() interface{}
	WebAclId() *string
	SetWebAclId(val *string)
	WebAclIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutDefaultCacheBehavior(value *CloudfrontDistributionDefaultCacheBehavior)
	PutLoggingConfig(value *CloudfrontDistributionLoggingConfig)
	PutRestrictions(value *CloudfrontDistributionRestrictions)
	PutViewerCertificate(value *CloudfrontDistributionViewerCertificate)
	ResetAliases()
	ResetComment()
	ResetCustomErrorResponse()
	ResetDefaultRootObject()
	ResetHttpVersion()
	ResetIsIpv6Enabled()
	ResetLoggingConfig()
	ResetOrderedCacheBehavior()
	ResetOriginGroup()
	ResetOverrideLogicalId()
	ResetPriceClass()
	ResetRetainOnDelete()
	ResetTags()
	ResetTagsAll()
	ResetWaitForDeployment()
	ResetWebAclId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
	TrustedKeyGroups(index *string) CloudfrontDistributionTrustedKeyGroups
	TrustedSigners(index *string) CloudfrontDistributionTrustedSigners
}

// The jsii proxy struct for CloudfrontDistribution
type jsiiProxy_CloudfrontDistribution struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontDistribution) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) AliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CustomErrorResponse() *[]*CloudfrontDistributionCustomErrorResponse {
	var returns *[]*CloudfrontDistributionCustomErrorResponse
	_jsii_.Get(
		j,
		"customErrorResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) CustomErrorResponseInput() *[]*CloudfrontDistributionCustomErrorResponse {
	var returns *[]*CloudfrontDistributionCustomErrorResponse
	_jsii_.Get(
		j,
		"customErrorResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DefaultCacheBehavior() CloudfrontDistributionDefaultCacheBehaviorOutputReference {
	var returns CloudfrontDistributionDefaultCacheBehaviorOutputReference
	_jsii_.Get(
		j,
		"defaultCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DefaultCacheBehaviorInput() *CloudfrontDistributionDefaultCacheBehavior {
	var returns *CloudfrontDistributionDefaultCacheBehavior
	_jsii_.Get(
		j,
		"defaultCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DefaultRootObject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DefaultRootObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRootObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) HttpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) InProgressValidationBatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inProgressValidationBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) IsIpv6Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) IsIpv6EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIpv6EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) LoggingConfig() CloudfrontDistributionLoggingConfigOutputReference {
	var returns CloudfrontDistributionLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) LoggingConfigInput() *CloudfrontDistributionLoggingConfig {
	var returns *CloudfrontDistributionLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OrderedCacheBehavior() *[]*CloudfrontDistributionOrderedCacheBehavior {
	var returns *[]*CloudfrontDistributionOrderedCacheBehavior
	_jsii_.Get(
		j,
		"orderedCacheBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OrderedCacheBehaviorInput() *[]*CloudfrontDistributionOrderedCacheBehavior {
	var returns *[]*CloudfrontDistributionOrderedCacheBehavior
	_jsii_.Get(
		j,
		"orderedCacheBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Origin() *[]*CloudfrontDistributionOrigin {
	var returns *[]*CloudfrontDistributionOrigin
	_jsii_.Get(
		j,
		"origin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OriginGroup() *[]*CloudfrontDistributionOriginGroup {
	var returns *[]*CloudfrontDistributionOriginGroup
	_jsii_.Get(
		j,
		"originGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OriginGroupInput() *[]*CloudfrontDistributionOriginGroup {
	var returns *[]*CloudfrontDistributionOriginGroup
	_jsii_.Get(
		j,
		"originGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) OriginInput() *[]*CloudfrontDistributionOrigin {
	var returns *[]*CloudfrontDistributionOrigin
	_jsii_.Get(
		j,
		"originInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) PriceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) PriceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Restrictions() CloudfrontDistributionRestrictionsOutputReference {
	var returns CloudfrontDistributionRestrictionsOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) RestrictionsInput() *CloudfrontDistributionRestrictions {
	var returns *CloudfrontDistributionRestrictions
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) RetainOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) RetainOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retainOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) ViewerCertificate() CloudfrontDistributionViewerCertificateOutputReference {
	var returns CloudfrontDistributionViewerCertificateOutputReference
	_jsii_.Get(
		j,
		"viewerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) ViewerCertificateInput() *CloudfrontDistributionViewerCertificate {
	var returns *CloudfrontDistributionViewerCertificate
	_jsii_.Get(
		j,
		"viewerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) WaitForDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) WaitForDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) WebAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistribution) WebAclIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webAclIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html aws_cloudfront_distribution} Resource.
func NewCloudfrontDistribution(scope constructs.Construct, id *string, config *CloudfrontDistributionConfig) CloudfrontDistribution {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistribution{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html aws_cloudfront_distribution} Resource.
func NewCloudfrontDistribution_Override(c CloudfrontDistribution, scope constructs.Construct, id *string, config *CloudfrontDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistribution",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetAliases(val *[]*string) {
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetCustomErrorResponse(val *[]*CloudfrontDistributionCustomErrorResponse) {
	_jsii_.Set(
		j,
		"customErrorResponse",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetDefaultRootObject(val *string) {
	_jsii_.Set(
		j,
		"defaultRootObject",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetHttpVersion(val *string) {
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetIsIpv6Enabled(val interface{}) {
	_jsii_.Set(
		j,
		"isIpv6Enabled",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetOrderedCacheBehavior(val *[]*CloudfrontDistributionOrderedCacheBehavior) {
	_jsii_.Set(
		j,
		"orderedCacheBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetOrigin(val *[]*CloudfrontDistributionOrigin) {
	_jsii_.Set(
		j,
		"origin",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetOriginGroup(val *[]*CloudfrontDistributionOriginGroup) {
	_jsii_.Set(
		j,
		"originGroup",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetPriceClass(val *string) {
	_jsii_.Set(
		j,
		"priceClass",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetRetainOnDelete(val interface{}) {
	_jsii_.Set(
		j,
		"retainOnDelete",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetWaitForDeployment(val interface{}) {
	_jsii_.Set(
		j,
		"waitForDeployment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistribution) SetWebAclId(val *string) {
	_jsii_.Set(
		j,
		"webAclId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudfrontDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistribution) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistribution) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistribution) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistribution) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistribution) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutDefaultCacheBehavior(value *CloudfrontDistributionDefaultCacheBehavior) {
	_jsii_.InvokeVoid(
		c,
		"putDefaultCacheBehavior",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutLoggingConfig(value *CloudfrontDistributionLoggingConfig) {
	_jsii_.InvokeVoid(
		c,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutRestrictions(value *CloudfrontDistributionRestrictions) {
	_jsii_.InvokeVoid(
		c,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) PutViewerCertificate(value *CloudfrontDistributionViewerCertificate) {
	_jsii_.InvokeVoid(
		c,
		"putViewerCertificate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetAliases() {
	_jsii_.InvokeVoid(
		c,
		"resetAliases",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetCustomErrorResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomErrorResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetDefaultRootObject() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultRootObject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetHttpVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetIsIpv6Enabled() {
	_jsii_.InvokeVoid(
		c,
		"resetIsIpv6Enabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetOrderedCacheBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetOrderedCacheBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetOriginGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginGroup",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetPriceClass() {
	_jsii_.InvokeVoid(
		c,
		"resetPriceClass",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetRetainOnDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetRetainOnDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetTagsAll() {
	_jsii_.InvokeVoid(
		c,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetWaitForDeployment() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitForDeployment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) ResetWebAclId() {
	_jsii_.InvokeVoid(
		c,
		"resetWebAclId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistribution) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontDistribution) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontDistribution) ToString() *string {
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
func (c *jsiiProxy_CloudfrontDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistribution) TrustedKeyGroups(index *string) CloudfrontDistributionTrustedKeyGroups {
	var returns CloudfrontDistributionTrustedKeyGroups

	_jsii_.Invoke(
		c,
		"trustedKeyGroups",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistribution) TrustedSigners(index *string) CloudfrontDistributionTrustedSigners {
	var returns CloudfrontDistributionTrustedSigners

	_jsii_.Invoke(
		c,
		"trustedSigners",
		[]interface{}{index},
		&returns,
	)

	return returns
}

type CloudfrontDistributionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// default_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#default_cache_behavior CloudfrontDistribution#default_cache_behavior}
	DefaultCacheBehavior *CloudfrontDistributionDefaultCacheBehavior `json:"defaultCacheBehavior"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#enabled CloudfrontDistribution#enabled}.
	Enabled interface{} `json:"enabled"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin CloudfrontDistribution#origin}
	Origin *[]*CloudfrontDistributionOrigin `json:"origin"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#restrictions CloudfrontDistribution#restrictions}
	Restrictions *CloudfrontDistributionRestrictions `json:"restrictions"`
	// viewer_certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#viewer_certificate CloudfrontDistribution#viewer_certificate}
	ViewerCertificate *CloudfrontDistributionViewerCertificate `json:"viewerCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#aliases CloudfrontDistribution#aliases}.
	Aliases *[]*string `json:"aliases"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#comment CloudfrontDistribution#comment}.
	Comment *string `json:"comment"`
	// custom_error_response block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#custom_error_response CloudfrontDistribution#custom_error_response}
	CustomErrorResponse *[]*CloudfrontDistributionCustomErrorResponse `json:"customErrorResponse"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#default_root_object CloudfrontDistribution#default_root_object}.
	DefaultRootObject *string `json:"defaultRootObject"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#http_version CloudfrontDistribution#http_version}.
	HttpVersion *string `json:"httpVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#is_ipv6_enabled CloudfrontDistribution#is_ipv6_enabled}.
	IsIpv6Enabled interface{} `json:"isIpv6Enabled"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#logging_config CloudfrontDistribution#logging_config}
	LoggingConfig *CloudfrontDistributionLoggingConfig `json:"loggingConfig"`
	// ordered_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#ordered_cache_behavior CloudfrontDistribution#ordered_cache_behavior}
	OrderedCacheBehavior *[]*CloudfrontDistributionOrderedCacheBehavior `json:"orderedCacheBehavior"`
	// origin_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_group CloudfrontDistribution#origin_group}
	OriginGroup *[]*CloudfrontDistributionOriginGroup `json:"originGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#price_class CloudfrontDistribution#price_class}.
	PriceClass *string `json:"priceClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#retain_on_delete CloudfrontDistribution#retain_on_delete}.
	RetainOnDelete interface{} `json:"retainOnDelete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#tags CloudfrontDistribution#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#tags_all CloudfrontDistribution#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#wait_for_deployment CloudfrontDistribution#wait_for_deployment}.
	WaitForDeployment interface{} `json:"waitForDeployment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#web_acl_id CloudfrontDistribution#web_acl_id}.
	WebAclId *string `json:"webAclId"`
}

type CloudfrontDistributionCustomErrorResponse struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#error_code CloudfrontDistribution#error_code}.
	ErrorCode *float64 `json:"errorCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#error_caching_min_ttl CloudfrontDistribution#error_caching_min_ttl}.
	ErrorCachingMinTtl *float64 `json:"errorCachingMinTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#response_code CloudfrontDistribution#response_code}.
	ResponseCode *float64 `json:"responseCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#response_page_path CloudfrontDistribution#response_page_path}.
	ResponsePagePath *string `json:"responsePagePath"`
}

type CloudfrontDistributionDefaultCacheBehavior struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#allowed_methods CloudfrontDistribution#allowed_methods}.
	AllowedMethods *[]*string `json:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#cached_methods CloudfrontDistribution#cached_methods}.
	CachedMethods *[]*string `json:"cachedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#target_origin_id CloudfrontDistribution#target_origin_id}.
	TargetOriginId *string `json:"targetOriginId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#viewer_protocol_policy CloudfrontDistribution#viewer_protocol_policy}.
	ViewerProtocolPolicy *string `json:"viewerProtocolPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#cache_policy_id CloudfrontDistribution#cache_policy_id}.
	CachePolicyId *string `json:"cachePolicyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#compress CloudfrontDistribution#compress}.
	Compress interface{} `json:"compress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#default_ttl CloudfrontDistribution#default_ttl}.
	DefaultTtl *float64 `json:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#field_level_encryption_id CloudfrontDistribution#field_level_encryption_id}.
	FieldLevelEncryptionId *string `json:"fieldLevelEncryptionId"`
	// forwarded_values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#forwarded_values CloudfrontDistribution#forwarded_values}
	ForwardedValues *CloudfrontDistributionDefaultCacheBehaviorForwardedValues `json:"forwardedValues"`
	// function_association block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#function_association CloudfrontDistribution#function_association}
	FunctionAssociation *[]*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation `json:"functionAssociation"`
	// lambda_function_association block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#lambda_function_association CloudfrontDistribution#lambda_function_association}
	LambdaFunctionAssociation *[]*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation `json:"lambdaFunctionAssociation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#max_ttl CloudfrontDistribution#max_ttl}.
	MaxTtl *float64 `json:"maxTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#min_ttl CloudfrontDistribution#min_ttl}.
	MinTtl *float64 `json:"minTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_request_policy_id CloudfrontDistribution#origin_request_policy_id}.
	OriginRequestPolicyId *string `json:"originRequestPolicyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#realtime_log_config_arn CloudfrontDistribution#realtime_log_config_arn}.
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#response_headers_policy_id CloudfrontDistribution#response_headers_policy_id}.
	ResponseHeadersPolicyId *string `json:"responseHeadersPolicyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#smooth_streaming CloudfrontDistribution#smooth_streaming}.
	SmoothStreaming interface{} `json:"smoothStreaming"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#trusted_key_groups CloudfrontDistribution#trusted_key_groups}.
	TrustedKeyGroups *[]*string `json:"trustedKeyGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#trusted_signers CloudfrontDistribution#trusted_signers}.
	TrustedSigners *[]*string `json:"trustedSigners"`
}

type CloudfrontDistributionDefaultCacheBehaviorForwardedValues struct {
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#cookies CloudfrontDistribution#cookies}
	Cookies *CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies `json:"cookies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#query_string CloudfrontDistribution#query_string}.
	QueryString interface{} `json:"queryString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#headers CloudfrontDistribution#headers}.
	Headers *[]*string `json:"headers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#query_string_cache_keys CloudfrontDistribution#query_string_cache_keys}.
	QueryStringCacheKeys *[]*string `json:"queryStringCacheKeys"`
}

type CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#forward CloudfrontDistribution#forward}.
	Forward *string `json:"forward"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#whitelisted_names CloudfrontDistribution#whitelisted_names}.
	WhitelistedNames *[]*string `json:"whitelistedNames"`
}

type CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference interface {
	cdktf.ComplexObject
	Forward() *string
	SetForward(val *string)
	ForwardInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	WhitelistedNames() *[]*string
	SetWhitelistedNames(val *[]*string)
	WhitelistedNamesInput() *[]*string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetWhitelistedNames()
}

// The jsii proxy struct for CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference
type jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) Forward() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) ForwardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) WhitelistedNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistedNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) WhitelistedNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistedNamesInput",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference_Override(c CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) SetForward(val *string) {
	_jsii_.Set(
		j,
		"forward",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) SetWhitelistedNames(val *[]*string) {
	_jsii_.Set(
		j,
		"whitelistedNames",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference) ResetWhitelistedNames() {
	_jsii_.InvokeVoid(
		c,
		"resetWhitelistedNames",
		nil, // no parameters
	)
}

type CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference interface {
	cdktf.ComplexObject
	Cookies() CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference
	CookiesInput() *CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies
	Headers() *[]*string
	SetHeaders(val *[]*string)
	HeadersInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	QueryString() interface{}
	SetQueryString(val interface{})
	QueryStringCacheKeys() *[]*string
	SetQueryStringCacheKeys(val *[]*string)
	QueryStringCacheKeysInput() *[]*string
	QueryStringInput() interface{}
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
	PutCookies(value *CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies)
	ResetHeaders()
	ResetQueryStringCacheKeys()
}

// The jsii proxy struct for CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference
type jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) Cookies() CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference {
	var returns CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) CookiesInput() *CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies {
	var returns *CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) HeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) QueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) QueryStringCacheKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) QueryStringCacheKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference_Override(c CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) SetHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) SetQueryString(val interface{}) {
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) SetQueryStringCacheKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"queryStringCacheKeys",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) PutCookies(value *CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies) {
	_jsii_.InvokeVoid(
		c,
		"putCookies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference) ResetQueryStringCacheKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringCacheKeys",
		nil, // no parameters
	)
}

type CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#event_type CloudfrontDistribution#event_type}.
	EventType *string `json:"eventType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#function_arn CloudfrontDistribution#function_arn}.
	FunctionArn *string `json:"functionArn"`
}

type CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#event_type CloudfrontDistribution#event_type}.
	EventType *string `json:"eventType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#lambda_arn CloudfrontDistribution#lambda_arn}.
	LambdaArn *string `json:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#include_body CloudfrontDistribution#include_body}.
	IncludeBody interface{} `json:"includeBody"`
}

type CloudfrontDistributionDefaultCacheBehaviorOutputReference interface {
	cdktf.ComplexObject
	AllowedMethods() *[]*string
	SetAllowedMethods(val *[]*string)
	AllowedMethodsInput() *[]*string
	CachedMethods() *[]*string
	SetCachedMethods(val *[]*string)
	CachedMethodsInput() *[]*string
	CachePolicyId() *string
	SetCachePolicyId(val *string)
	CachePolicyIdInput() *string
	Compress() interface{}
	SetCompress(val interface{})
	CompressInput() interface{}
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
	FieldLevelEncryptionId() *string
	SetFieldLevelEncryptionId(val *string)
	FieldLevelEncryptionIdInput() *string
	ForwardedValues() CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference
	ForwardedValuesInput() *CloudfrontDistributionDefaultCacheBehaviorForwardedValues
	FunctionAssociation() *[]*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation
	SetFunctionAssociation(val *[]*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation)
	FunctionAssociationInput() *[]*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	LambdaFunctionAssociation() *[]*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation
	SetLambdaFunctionAssociation(val *[]*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation)
	LambdaFunctionAssociationInput() *[]*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation
	MaxTtl() *float64
	SetMaxTtl(val *float64)
	MaxTtlInput() *float64
	MinTtl() *float64
	SetMinTtl(val *float64)
	MinTtlInput() *float64
	OriginRequestPolicyId() *string
	SetOriginRequestPolicyId(val *string)
	OriginRequestPolicyIdInput() *string
	RealtimeLogConfigArn() *string
	SetRealtimeLogConfigArn(val *string)
	RealtimeLogConfigArnInput() *string
	ResponseHeadersPolicyId() *string
	SetResponseHeadersPolicyId(val *string)
	ResponseHeadersPolicyIdInput() *string
	SmoothStreaming() interface{}
	SetSmoothStreaming(val interface{})
	SmoothStreamingInput() interface{}
	TargetOriginId() *string
	SetTargetOriginId(val *string)
	TargetOriginIdInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TrustedKeyGroups() *[]*string
	SetTrustedKeyGroups(val *[]*string)
	TrustedKeyGroupsInput() *[]*string
	TrustedSigners() *[]*string
	SetTrustedSigners(val *[]*string)
	TrustedSignersInput() *[]*string
	ViewerProtocolPolicy() *string
	SetViewerProtocolPolicy(val *string)
	ViewerProtocolPolicyInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutForwardedValues(value *CloudfrontDistributionDefaultCacheBehaviorForwardedValues)
	ResetCachePolicyId()
	ResetCompress()
	ResetDefaultTtl()
	ResetFieldLevelEncryptionId()
	ResetForwardedValues()
	ResetFunctionAssociation()
	ResetLambdaFunctionAssociation()
	ResetMaxTtl()
	ResetMinTtl()
	ResetOriginRequestPolicyId()
	ResetRealtimeLogConfigArn()
	ResetResponseHeadersPolicyId()
	ResetSmoothStreaming()
	ResetTrustedKeyGroups()
	ResetTrustedSigners()
}

// The jsii proxy struct for CloudfrontDistributionDefaultCacheBehaviorOutputReference
type jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CachedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CachedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cachedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CachePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CachePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) Compress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) CompressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) FieldLevelEncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) FieldLevelEncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldLevelEncryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ForwardedValues() CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference {
	var returns CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference
	_jsii_.Get(
		j,
		"forwardedValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ForwardedValuesInput() *CloudfrontDistributionDefaultCacheBehaviorForwardedValues {
	var returns *CloudfrontDistributionDefaultCacheBehaviorForwardedValues
	_jsii_.Get(
		j,
		"forwardedValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) FunctionAssociation() *[]*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation {
	var returns *[]*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation
	_jsii_.Get(
		j,
		"functionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) FunctionAssociationInput() *[]*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation {
	var returns *[]*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation
	_jsii_.Get(
		j,
		"functionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) LambdaFunctionAssociation() *[]*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation {
	var returns *[]*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation
	_jsii_.Get(
		j,
		"lambdaFunctionAssociation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) LambdaFunctionAssociationInput() *[]*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation {
	var returns *[]*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation
	_jsii_.Get(
		j,
		"lambdaFunctionAssociationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) MinTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) MinTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) OriginRequestPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) OriginRequestPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originRequestPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) RealtimeLogConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) RealtimeLogConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeLogConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResponseHeadersPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResponseHeadersPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeadersPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SmoothStreaming() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SmoothStreamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smoothStreamingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TargetOriginId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TargetOriginIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetOriginIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TrustedKeyGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TrustedKeyGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedKeyGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TrustedSigners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSigners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) TrustedSignersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedSignersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ViewerProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ViewerProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewerProtocolPolicyInput",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionDefaultCacheBehaviorOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionDefaultCacheBehaviorOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionDefaultCacheBehaviorOutputReference_Override(c CloudfrontDistributionDefaultCacheBehaviorOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetAllowedMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetCachedMethods(val *[]*string) {
	_jsii_.Set(
		j,
		"cachedMethods",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetCachePolicyId(val *string) {
	_jsii_.Set(
		j,
		"cachePolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetCompress(val interface{}) {
	_jsii_.Set(
		j,
		"compress",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetDefaultTtl(val *float64) {
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetFieldLevelEncryptionId(val *string) {
	_jsii_.Set(
		j,
		"fieldLevelEncryptionId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetFunctionAssociation(val *[]*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation) {
	_jsii_.Set(
		j,
		"functionAssociation",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetLambdaFunctionAssociation(val *[]*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation) {
	_jsii_.Set(
		j,
		"lambdaFunctionAssociation",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetMaxTtl(val *float64) {
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetMinTtl(val *float64) {
	_jsii_.Set(
		j,
		"minTtl",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetOriginRequestPolicyId(val *string) {
	_jsii_.Set(
		j,
		"originRequestPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetRealtimeLogConfigArn(val *string) {
	_jsii_.Set(
		j,
		"realtimeLogConfigArn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetResponseHeadersPolicyId(val *string) {
	_jsii_.Set(
		j,
		"responseHeadersPolicyId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetSmoothStreaming(val interface{}) {
	_jsii_.Set(
		j,
		"smoothStreaming",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetTargetOriginId(val *string) {
	_jsii_.Set(
		j,
		"targetOriginId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetTrustedKeyGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"trustedKeyGroups",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetTrustedSigners(val *[]*string) {
	_jsii_.Set(
		j,
		"trustedSigners",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) SetViewerProtocolPolicy(val *string) {
	_jsii_.Set(
		j,
		"viewerProtocolPolicy",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) PutForwardedValues(value *CloudfrontDistributionDefaultCacheBehaviorForwardedValues) {
	_jsii_.InvokeVoid(
		c,
		"putForwardedValues",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetCachePolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetCachePolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetCompress() {
	_jsii_.InvokeVoid(
		c,
		"resetCompress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetFieldLevelEncryptionId() {
	_jsii_.InvokeVoid(
		c,
		"resetFieldLevelEncryptionId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetForwardedValues() {
	_jsii_.InvokeVoid(
		c,
		"resetForwardedValues",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetFunctionAssociation() {
	_jsii_.InvokeVoid(
		c,
		"resetFunctionAssociation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetLambdaFunctionAssociation() {
	_jsii_.InvokeVoid(
		c,
		"resetLambdaFunctionAssociation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetMinTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMinTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetOriginRequestPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginRequestPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetRealtimeLogConfigArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRealtimeLogConfigArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetResponseHeadersPolicyId() {
	_jsii_.InvokeVoid(
		c,
		"resetResponseHeadersPolicyId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetSmoothStreaming() {
	_jsii_.InvokeVoid(
		c,
		"resetSmoothStreaming",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetTrustedKeyGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetTrustedKeyGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference) ResetTrustedSigners() {
	_jsii_.InvokeVoid(
		c,
		"resetTrustedSigners",
		nil, // no parameters
	)
}

type CloudfrontDistributionLoggingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#bucket CloudfrontDistribution#bucket}.
	Bucket *string `json:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#include_cookies CloudfrontDistribution#include_cookies}.
	IncludeCookies interface{} `json:"includeCookies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#prefix CloudfrontDistribution#prefix}.
	Prefix *string `json:"prefix"`
}

type CloudfrontDistributionLoggingConfigOutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	IncludeCookies() interface{}
	SetIncludeCookies(val interface{})
	IncludeCookiesInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
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
	ResetIncludeCookies()
	ResetPrefix()
}

// The jsii proxy struct for CloudfrontDistributionLoggingConfigOutputReference
type jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) IncludeCookies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) IncludeCookiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionLoggingConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionLoggingConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionLoggingConfigOutputReference_Override(c CloudfrontDistributionLoggingConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) SetIncludeCookies(val interface{}) {
	_jsii_.Set(
		j,
		"includeCookies",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) SetPrefix(val *string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) ResetIncludeCookies() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeCookies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetPrefix",
		nil, // no parameters
	)
}

type CloudfrontDistributionOrderedCacheBehavior struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#allowed_methods CloudfrontDistribution#allowed_methods}.
	AllowedMethods *[]*string `json:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#cached_methods CloudfrontDistribution#cached_methods}.
	CachedMethods *[]*string `json:"cachedMethods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#path_pattern CloudfrontDistribution#path_pattern}.
	PathPattern *string `json:"pathPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#target_origin_id CloudfrontDistribution#target_origin_id}.
	TargetOriginId *string `json:"targetOriginId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#viewer_protocol_policy CloudfrontDistribution#viewer_protocol_policy}.
	ViewerProtocolPolicy *string `json:"viewerProtocolPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#cache_policy_id CloudfrontDistribution#cache_policy_id}.
	CachePolicyId *string `json:"cachePolicyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#compress CloudfrontDistribution#compress}.
	Compress interface{} `json:"compress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#default_ttl CloudfrontDistribution#default_ttl}.
	DefaultTtl *float64 `json:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#field_level_encryption_id CloudfrontDistribution#field_level_encryption_id}.
	FieldLevelEncryptionId *string `json:"fieldLevelEncryptionId"`
	// forwarded_values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#forwarded_values CloudfrontDistribution#forwarded_values}
	ForwardedValues *CloudfrontDistributionOrderedCacheBehaviorForwardedValues `json:"forwardedValues"`
	// function_association block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#function_association CloudfrontDistribution#function_association}
	FunctionAssociation *[]*CloudfrontDistributionOrderedCacheBehaviorFunctionAssociation `json:"functionAssociation"`
	// lambda_function_association block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#lambda_function_association CloudfrontDistribution#lambda_function_association}
	LambdaFunctionAssociation *[]*CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociation `json:"lambdaFunctionAssociation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#max_ttl CloudfrontDistribution#max_ttl}.
	MaxTtl *float64 `json:"maxTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#min_ttl CloudfrontDistribution#min_ttl}.
	MinTtl *float64 `json:"minTtl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_request_policy_id CloudfrontDistribution#origin_request_policy_id}.
	OriginRequestPolicyId *string `json:"originRequestPolicyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#realtime_log_config_arn CloudfrontDistribution#realtime_log_config_arn}.
	RealtimeLogConfigArn *string `json:"realtimeLogConfigArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#response_headers_policy_id CloudfrontDistribution#response_headers_policy_id}.
	ResponseHeadersPolicyId *string `json:"responseHeadersPolicyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#smooth_streaming CloudfrontDistribution#smooth_streaming}.
	SmoothStreaming interface{} `json:"smoothStreaming"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#trusted_key_groups CloudfrontDistribution#trusted_key_groups}.
	TrustedKeyGroups *[]*string `json:"trustedKeyGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#trusted_signers CloudfrontDistribution#trusted_signers}.
	TrustedSigners *[]*string `json:"trustedSigners"`
}

type CloudfrontDistributionOrderedCacheBehaviorForwardedValues struct {
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#cookies CloudfrontDistribution#cookies}
	Cookies *CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies `json:"cookies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#query_string CloudfrontDistribution#query_string}.
	QueryString interface{} `json:"queryString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#headers CloudfrontDistribution#headers}.
	Headers *[]*string `json:"headers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#query_string_cache_keys CloudfrontDistribution#query_string_cache_keys}.
	QueryStringCacheKeys *[]*string `json:"queryStringCacheKeys"`
}

type CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#forward CloudfrontDistribution#forward}.
	Forward *string `json:"forward"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#whitelisted_names CloudfrontDistribution#whitelisted_names}.
	WhitelistedNames *[]*string `json:"whitelistedNames"`
}

type CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference interface {
	cdktf.ComplexObject
	Forward() *string
	SetForward(val *string)
	ForwardInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	WhitelistedNames() *[]*string
	SetWhitelistedNames(val *[]*string)
	WhitelistedNamesInput() *[]*string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetWhitelistedNames()
}

// The jsii proxy struct for CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference
type jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) Forward() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) ForwardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) WhitelistedNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistedNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) WhitelistedNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"whitelistedNamesInput",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference_Override(c CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) SetForward(val *string) {
	_jsii_.Set(
		j,
		"forward",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) SetWhitelistedNames(val *[]*string) {
	_jsii_.Set(
		j,
		"whitelistedNames",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference) ResetWhitelistedNames() {
	_jsii_.InvokeVoid(
		c,
		"resetWhitelistedNames",
		nil, // no parameters
	)
}

type CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference interface {
	cdktf.ComplexObject
	Cookies() CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference
	CookiesInput() *CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies
	Headers() *[]*string
	SetHeaders(val *[]*string)
	HeadersInput() *[]*string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	QueryString() interface{}
	SetQueryString(val interface{})
	QueryStringCacheKeys() *[]*string
	SetQueryStringCacheKeys(val *[]*string)
	QueryStringCacheKeysInput() *[]*string
	QueryStringInput() interface{}
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
	PutCookies(value *CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies)
	ResetHeaders()
	ResetQueryStringCacheKeys()
}

// The jsii proxy struct for CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference
type jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) Cookies() CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference {
	var returns CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) CookiesInput() *CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies {
	var returns *CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) HeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) QueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) QueryStringCacheKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) QueryStringCacheKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference_Override(c CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) SetHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) SetQueryString(val interface{}) {
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) SetQueryStringCacheKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"queryStringCacheKeys",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) PutCookies(value *CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies) {
	_jsii_.InvokeVoid(
		c,
		"putCookies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference) ResetQueryStringCacheKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringCacheKeys",
		nil, // no parameters
	)
}

type CloudfrontDistributionOrderedCacheBehaviorFunctionAssociation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#event_type CloudfrontDistribution#event_type}.
	EventType *string `json:"eventType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#function_arn CloudfrontDistribution#function_arn}.
	FunctionArn *string `json:"functionArn"`
}

type CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#event_type CloudfrontDistribution#event_type}.
	EventType *string `json:"eventType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#lambda_arn CloudfrontDistribution#lambda_arn}.
	LambdaArn *string `json:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#include_body CloudfrontDistribution#include_body}.
	IncludeBody interface{} `json:"includeBody"`
}

type CloudfrontDistributionOrigin struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#domain_name CloudfrontDistribution#domain_name}.
	DomainName *string `json:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_id CloudfrontDistribution#origin_id}.
	OriginId *string `json:"originId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#connection_attempts CloudfrontDistribution#connection_attempts}.
	ConnectionAttempts *float64 `json:"connectionAttempts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#connection_timeout CloudfrontDistribution#connection_timeout}.
	ConnectionTimeout *float64 `json:"connectionTimeout"`
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#custom_header CloudfrontDistribution#custom_header}
	CustomHeader *[]*CloudfrontDistributionOriginCustomHeader `json:"customHeader"`
	// custom_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#custom_origin_config CloudfrontDistribution#custom_origin_config}
	CustomOriginConfig *CloudfrontDistributionOriginCustomOriginConfig `json:"customOriginConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_path CloudfrontDistribution#origin_path}.
	OriginPath *string `json:"originPath"`
	// origin_shield block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_shield CloudfrontDistribution#origin_shield}
	OriginShield *CloudfrontDistributionOriginOriginShield `json:"originShield"`
	// s3_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#s3_origin_config CloudfrontDistribution#s3_origin_config}
	S3OriginConfig *CloudfrontDistributionOriginS3OriginConfig `json:"s3OriginConfig"`
}

type CloudfrontDistributionOriginCustomHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#name CloudfrontDistribution#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#value CloudfrontDistribution#value}.
	Value *string `json:"value"`
}

type CloudfrontDistributionOriginCustomOriginConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#http_port CloudfrontDistribution#http_port}.
	HttpPort *float64 `json:"httpPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#https_port CloudfrontDistribution#https_port}.
	HttpsPort *float64 `json:"httpsPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_protocol_policy CloudfrontDistribution#origin_protocol_policy}.
	OriginProtocolPolicy *string `json:"originProtocolPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_ssl_protocols CloudfrontDistribution#origin_ssl_protocols}.
	OriginSslProtocols *[]*string `json:"originSslProtocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_keepalive_timeout CloudfrontDistribution#origin_keepalive_timeout}.
	OriginKeepaliveTimeout *float64 `json:"originKeepaliveTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_read_timeout CloudfrontDistribution#origin_read_timeout}.
	OriginReadTimeout *float64 `json:"originReadTimeout"`
}

type CloudfrontDistributionOriginCustomOriginConfigOutputReference interface {
	cdktf.ComplexObject
	HttpPort() *float64
	SetHttpPort(val *float64)
	HttpPortInput() *float64
	HttpsPort() *float64
	SetHttpsPort(val *float64)
	HttpsPortInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OriginKeepaliveTimeout() *float64
	SetOriginKeepaliveTimeout(val *float64)
	OriginKeepaliveTimeoutInput() *float64
	OriginProtocolPolicy() *string
	SetOriginProtocolPolicy(val *string)
	OriginProtocolPolicyInput() *string
	OriginReadTimeout() *float64
	SetOriginReadTimeout(val *float64)
	OriginReadTimeoutInput() *float64
	OriginSslProtocols() *[]*string
	SetOriginSslProtocols(val *[]*string)
	OriginSslProtocolsInput() *[]*string
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
	ResetOriginKeepaliveTimeout()
	ResetOriginReadTimeout()
}

// The jsii proxy struct for CloudfrontDistributionOriginCustomOriginConfigOutputReference
type jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) HttpsPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) HttpsPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) OriginKeepaliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originKeepaliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) OriginKeepaliveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originKeepaliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) OriginProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) OriginProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) OriginReadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originReadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) OriginReadTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originReadTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) OriginSslProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originSslProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) OriginSslProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originSslProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionOriginCustomOriginConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionOriginCustomOriginConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginCustomOriginConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionOriginCustomOriginConfigOutputReference_Override(c CloudfrontDistributionOriginCustomOriginConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginCustomOriginConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) SetHttpPort(val *float64) {
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) SetHttpsPort(val *float64) {
	_jsii_.Set(
		j,
		"httpsPort",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) SetOriginKeepaliveTimeout(val *float64) {
	_jsii_.Set(
		j,
		"originKeepaliveTimeout",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) SetOriginProtocolPolicy(val *string) {
	_jsii_.Set(
		j,
		"originProtocolPolicy",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) SetOriginReadTimeout(val *float64) {
	_jsii_.Set(
		j,
		"originReadTimeout",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) SetOriginSslProtocols(val *[]*string) {
	_jsii_.Set(
		j,
		"originSslProtocols",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) ResetOriginKeepaliveTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginKeepaliveTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference) ResetOriginReadTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetOriginReadTimeout",
		nil, // no parameters
	)
}

type CloudfrontDistributionOriginGroup struct {
	// failover_criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#failover_criteria CloudfrontDistribution#failover_criteria}
	FailoverCriteria *CloudfrontDistributionOriginGroupFailoverCriteria `json:"failoverCriteria"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#member CloudfrontDistribution#member}
	Member *[]*CloudfrontDistributionOriginGroupMember `json:"member"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_id CloudfrontDistribution#origin_id}.
	OriginId *string `json:"originId"`
}

type CloudfrontDistributionOriginGroupFailoverCriteria struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#status_codes CloudfrontDistribution#status_codes}.
	StatusCodes *[]*float64 `json:"statusCodes"`
}

type CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	StatusCodes() *[]*float64
	SetStatusCodes(val *[]*float64)
	StatusCodesInput() *[]*float64
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
}

// The jsii proxy struct for CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference
type jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) StatusCodes() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"statusCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) StatusCodesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"statusCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionOriginGroupFailoverCriteriaOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionOriginGroupFailoverCriteriaOutputReference_Override(c CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) SetStatusCodes(val *[]*float64) {
	_jsii_.Set(
		j,
		"statusCodes",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontDistributionOriginGroupMember struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_id CloudfrontDistribution#origin_id}.
	OriginId *string `json:"originId"`
}

type CloudfrontDistributionOriginOriginShield struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#enabled CloudfrontDistribution#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_shield_region CloudfrontDistribution#origin_shield_region}.
	OriginShieldRegion *string `json:"originShieldRegion"`
}

type CloudfrontDistributionOriginOriginShieldOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OriginShieldRegion() *string
	SetOriginShieldRegion(val *string)
	OriginShieldRegionInput() *string
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
}

// The jsii proxy struct for CloudfrontDistributionOriginOriginShieldOutputReference
type jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) OriginShieldRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originShieldRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) OriginShieldRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originShieldRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionOriginOriginShieldOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionOriginOriginShieldOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginOriginShieldOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionOriginOriginShieldOutputReference_Override(c CloudfrontDistributionOriginOriginShieldOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginOriginShieldOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) SetOriginShieldRegion(val *string) {
	_jsii_.Set(
		j,
		"originShieldRegion",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontDistributionOriginS3OriginConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#origin_access_identity CloudfrontDistribution#origin_access_identity}.
	OriginAccessIdentity *string `json:"originAccessIdentity"`
}

type CloudfrontDistributionOriginS3OriginConfigOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OriginAccessIdentity() *string
	SetOriginAccessIdentity(val *string)
	OriginAccessIdentityInput() *string
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
}

// The jsii proxy struct for CloudfrontDistributionOriginS3OriginConfigOutputReference
type jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) OriginAccessIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) OriginAccessIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originAccessIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionOriginS3OriginConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionOriginS3OriginConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginS3OriginConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionOriginS3OriginConfigOutputReference_Override(c CloudfrontDistributionOriginS3OriginConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginS3OriginConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) SetOriginAccessIdentity(val *string) {
	_jsii_.Set(
		j,
		"originAccessIdentity",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontDistributionRestrictions struct {
	// geo_restriction block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#geo_restriction CloudfrontDistribution#geo_restriction}
	GeoRestriction *CloudfrontDistributionRestrictionsGeoRestriction `json:"geoRestriction"`
}

type CloudfrontDistributionRestrictionsGeoRestriction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#restriction_type CloudfrontDistribution#restriction_type}.
	RestrictionType *string `json:"restrictionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#locations CloudfrontDistribution#locations}.
	Locations *[]*string `json:"locations"`
}

type CloudfrontDistributionRestrictionsGeoRestrictionOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Locations() *[]*string
	SetLocations(val *[]*string)
	LocationsInput() *[]*string
	RestrictionType() *string
	SetRestrictionType(val *string)
	RestrictionTypeInput() *string
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
	ResetLocations()
}

// The jsii proxy struct for CloudfrontDistributionRestrictionsGeoRestrictionOutputReference
type jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) Locations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) LocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) RestrictionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) RestrictionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restrictionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionRestrictionsGeoRestrictionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionRestrictionsGeoRestrictionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionRestrictionsGeoRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionRestrictionsGeoRestrictionOutputReference_Override(c CloudfrontDistributionRestrictionsGeoRestrictionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionRestrictionsGeoRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) SetLocations(val *[]*string) {
	_jsii_.Set(
		j,
		"locations",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) SetRestrictionType(val *string) {
	_jsii_.Set(
		j,
		"restrictionType",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference) ResetLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetLocations",
		nil, // no parameters
	)
}

type CloudfrontDistributionRestrictionsOutputReference interface {
	cdktf.ComplexObject
	GeoRestriction() CloudfrontDistributionRestrictionsGeoRestrictionOutputReference
	GeoRestrictionInput() *CloudfrontDistributionRestrictionsGeoRestriction
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
	PutGeoRestriction(value *CloudfrontDistributionRestrictionsGeoRestriction)
}

// The jsii proxy struct for CloudfrontDistributionRestrictionsOutputReference
type jsiiProxy_CloudfrontDistributionRestrictionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) GeoRestriction() CloudfrontDistributionRestrictionsGeoRestrictionOutputReference {
	var returns CloudfrontDistributionRestrictionsGeoRestrictionOutputReference
	_jsii_.Get(
		j,
		"geoRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) GeoRestrictionInput() *CloudfrontDistributionRestrictionsGeoRestriction {
	var returns *CloudfrontDistributionRestrictionsGeoRestriction
	_jsii_.Get(
		j,
		"geoRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionRestrictionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionRestrictionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionRestrictionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionRestrictionsOutputReference_Override(c CloudfrontDistributionRestrictionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionRestrictionsOutputReference) PutGeoRestriction(value *CloudfrontDistributionRestrictionsGeoRestriction) {
	_jsii_.InvokeVoid(
		c,
		"putGeoRestriction",
		[]interface{}{value},
	)
}

type CloudfrontDistributionTrustedKeyGroups interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Enabled() interface{}
	Items() interface{}
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

// The jsii proxy struct for CloudfrontDistributionTrustedKeyGroups
type jsiiProxy_CloudfrontDistributionTrustedKeyGroups struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) Items() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewCloudfrontDistributionTrustedKeyGroups(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) CloudfrontDistributionTrustedKeyGroups {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionTrustedKeyGroups{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedKeyGroups",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudfrontDistributionTrustedKeyGroups_Override(c CloudfrontDistributionTrustedKeyGroups, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedKeyGroups",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroups) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontDistributionTrustedKeyGroupsItems interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	KeyGroupId() *string
	KeyPairIds() *[]*string
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

// The jsii proxy struct for CloudfrontDistributionTrustedKeyGroupsItems
type jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) KeyGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) KeyPairIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPairIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewCloudfrontDistributionTrustedKeyGroupsItems(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) CloudfrontDistributionTrustedKeyGroupsItems {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedKeyGroupsItems",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudfrontDistributionTrustedKeyGroupsItems_Override(c CloudfrontDistributionTrustedKeyGroupsItems, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedKeyGroupsItems",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontDistributionTrustedSigners interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Enabled() interface{}
	Items() interface{}
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

// The jsii proxy struct for CloudfrontDistributionTrustedSigners
type jsiiProxy_CloudfrontDistributionTrustedSigners struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSigners) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSigners) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSigners) Items() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSigners) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSigners) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewCloudfrontDistributionTrustedSigners(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) CloudfrontDistributionTrustedSigners {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionTrustedSigners{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedSigners",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudfrontDistributionTrustedSigners_Override(c CloudfrontDistributionTrustedSigners, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedSigners",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSigners) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSigners) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSigners) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionTrustedSigners) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionTrustedSigners) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedSigners) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedSigners) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedSigners) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontDistributionTrustedSignersItems interface {
	cdktf.ComplexComputedList
	AwsAccountNumber() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	KeyPairIds() *[]*string
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

// The jsii proxy struct for CloudfrontDistributionTrustedSignersItems
type jsiiProxy_CloudfrontDistributionTrustedSignersItems struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersItems) AwsAccountNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersItems) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersItems) KeyPairIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPairIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersItems) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersItems) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewCloudfrontDistributionTrustedSignersItems(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) CloudfrontDistributionTrustedSignersItems {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionTrustedSignersItems{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedSignersItems",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudfrontDistributionTrustedSignersItems_Override(c CloudfrontDistributionTrustedSignersItems, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedSignersItems",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersItems) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersItems) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersItems) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionTrustedSignersItems) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionTrustedSignersItems) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedSignersItems) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedSignersItems) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionTrustedSignersItems) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontDistributionViewerCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#acm_certificate_arn CloudfrontDistribution#acm_certificate_arn}.
	AcmCertificateArn *string `json:"acmCertificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#cloudfront_default_certificate CloudfrontDistribution#cloudfront_default_certificate}.
	CloudfrontDefaultCertificate interface{} `json:"cloudfrontDefaultCertificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#iam_certificate_id CloudfrontDistribution#iam_certificate_id}.
	IamCertificateId *string `json:"iamCertificateId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#minimum_protocol_version CloudfrontDistribution#minimum_protocol_version}.
	MinimumProtocolVersion *string `json:"minimumProtocolVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html#ssl_support_method CloudfrontDistribution#ssl_support_method}.
	SslSupportMethod *string `json:"sslSupportMethod"`
}

type CloudfrontDistributionViewerCertificateOutputReference interface {
	cdktf.ComplexObject
	AcmCertificateArn() *string
	SetAcmCertificateArn(val *string)
	AcmCertificateArnInput() *string
	CloudfrontDefaultCertificate() interface{}
	SetCloudfrontDefaultCertificate(val interface{})
	CloudfrontDefaultCertificateInput() interface{}
	IamCertificateId() *string
	SetIamCertificateId(val *string)
	IamCertificateIdInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MinimumProtocolVersion() *string
	SetMinimumProtocolVersion(val *string)
	MinimumProtocolVersionInput() *string
	SslSupportMethod() *string
	SetSslSupportMethod(val *string)
	SslSupportMethodInput() *string
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
	ResetAcmCertificateArn()
	ResetCloudfrontDefaultCertificate()
	ResetIamCertificateId()
	ResetMinimumProtocolVersion()
	ResetSslSupportMethod()
}

// The jsii proxy struct for CloudfrontDistributionViewerCertificateOutputReference
type jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) AcmCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) AcmCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acmCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) CloudfrontDefaultCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudfrontDefaultCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) CloudfrontDefaultCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudfrontDefaultCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) IamCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) IamCertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamCertificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) MinimumProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProtocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) MinimumProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumProtocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SslSupportMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSupportMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SslSupportMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSupportMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontDistributionViewerCertificateOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontDistributionViewerCertificateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionViewerCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontDistributionViewerCertificateOutputReference_Override(c CloudfrontDistributionViewerCertificateOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontDistributionViewerCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SetAcmCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"acmCertificateArn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SetCloudfrontDefaultCertificate(val interface{}) {
	_jsii_.Set(
		j,
		"cloudfrontDefaultCertificate",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SetIamCertificateId(val *string) {
	_jsii_.Set(
		j,
		"iamCertificateId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SetMinimumProtocolVersion(val *string) {
	_jsii_.Set(
		j,
		"minimumProtocolVersion",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SetSslSupportMethod(val *string) {
	_jsii_.Set(
		j,
		"sslSupportMethod",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) ResetAcmCertificateArn() {
	_jsii_.InvokeVoid(
		c,
		"resetAcmCertificateArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) ResetCloudfrontDefaultCertificate() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudfrontDefaultCertificate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) ResetIamCertificateId() {
	_jsii_.InvokeVoid(
		c,
		"resetIamCertificateId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) ResetMinimumProtocolVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetMinimumProtocolVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference) ResetSslSupportMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetSslSupportMethod",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_function.html aws_cloudfront_function}.
type CloudfrontFunction interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Code() *string
	SetCode(val *string)
	CodeInput() *string
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
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
	Publish() interface{}
	SetPublish(val interface{})
	PublishInput() interface{}
	RawOverrides() interface{}
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
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
	ResetComment()
	ResetOverrideLogicalId()
	ResetPublish()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfrontFunction
type jsiiProxy_CloudfrontFunction struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Code() *string {
	var returns *string
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) CodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Publish() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) PublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_function.html aws_cloudfront_function} Resource.
func NewCloudfrontFunction(scope constructs.Construct, id *string, config *CloudfrontFunctionConfig) CloudfrontFunction {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontFunction{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_function.html aws_cloudfront_function} Resource.
func NewCloudfrontFunction_Override(c CloudfrontFunction, scope constructs.Construct, id *string, config *CloudfrontFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontFunction",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontFunction) SetCode(val *string) {
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFunction) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFunction) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFunction) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFunction) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFunction) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFunction) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFunction) SetPublish(val interface{}) {
	_jsii_.Set(
		j,
		"publish",
		val,
	)
}

func (j *jsiiProxy_CloudfrontFunction) SetRuntime(val *string) {
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudfrontFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontFunction) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontFunction) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontFunction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontFunction) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontFunction) ResetPublish() {
	_jsii_.InvokeVoid(
		c,
		"resetPublish",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontFunction) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontFunction) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontFunction) ToString() *string {
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
func (c *jsiiProxy_CloudfrontFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudfrontFunctionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_function.html#code CloudfrontFunction#code}.
	Code *string `json:"code"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_function.html#name CloudfrontFunction#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_function.html#runtime CloudfrontFunction#runtime}.
	Runtime *string `json:"runtime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_function.html#comment CloudfrontFunction#comment}.
	Comment *string `json:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_function.html#publish CloudfrontFunction#publish}.
	Publish interface{} `json:"publish"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_key_group.html aws_cloudfront_key_group}.
type CloudfrontKeyGroup interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetComment()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfrontKeyGroup
type jsiiProxy_CloudfrontKeyGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontKeyGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontKeyGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_key_group.html aws_cloudfront_key_group} Resource.
func NewCloudfrontKeyGroup(scope constructs.Construct, id *string, config *CloudfrontKeyGroupConfig) CloudfrontKeyGroup {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontKeyGroup{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontKeyGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_key_group.html aws_cloudfront_key_group} Resource.
func NewCloudfrontKeyGroup_Override(c CloudfrontKeyGroup, scope constructs.Construct, id *string, config *CloudfrontKeyGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontKeyGroup",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontKeyGroup) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontKeyGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontKeyGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontKeyGroup) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontKeyGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontKeyGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfrontKeyGroup) SetProvider(val cdktf.TerraformProvider) {
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
func CloudfrontKeyGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontKeyGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontKeyGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontKeyGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontKeyGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontKeyGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontKeyGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontKeyGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontKeyGroup) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontKeyGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontKeyGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontKeyGroup) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontKeyGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontKeyGroup) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontKeyGroup) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontKeyGroup) ToString() *string {
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
func (c *jsiiProxy_CloudfrontKeyGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudfrontKeyGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_key_group.html#items CloudfrontKeyGroup#items}.
	Items *[]*string `json:"items"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_key_group.html#name CloudfrontKeyGroup#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_key_group.html#comment CloudfrontKeyGroup#comment}.
	Comment *string `json:"comment"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_monitoring_subscription.html aws_cloudfront_monitoring_subscription}.
type CloudfrontMonitoringSubscription interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DistributionId() *string
	SetDistributionId(val *string)
	DistributionIdInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MonitoringSubscription() CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference
	MonitoringSubscriptionInput() *CloudfrontMonitoringSubscriptionMonitoringSubscription
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
	PutMonitoringSubscription(value *CloudfrontMonitoringSubscriptionMonitoringSubscription)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfrontMonitoringSubscription
type jsiiProxy_CloudfrontMonitoringSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) DistributionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) DistributionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) MonitoringSubscription() CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference {
	var returns CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference
	_jsii_.Get(
		j,
		"monitoringSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) MonitoringSubscriptionInput() *CloudfrontMonitoringSubscriptionMonitoringSubscription {
	var returns *CloudfrontMonitoringSubscriptionMonitoringSubscription
	_jsii_.Get(
		j,
		"monitoringSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_monitoring_subscription.html aws_cloudfront_monitoring_subscription} Resource.
func NewCloudfrontMonitoringSubscription(scope constructs.Construct, id *string, config *CloudfrontMonitoringSubscriptionConfig) CloudfrontMonitoringSubscription {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontMonitoringSubscription{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_monitoring_subscription.html aws_cloudfront_monitoring_subscription} Resource.
func NewCloudfrontMonitoringSubscription_Override(c CloudfrontMonitoringSubscription, scope constructs.Construct, id *string, config *CloudfrontMonitoringSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscription",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) SetDistributionId(val *string) {
	_jsii_.Set(
		j,
		"distributionId",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscription) SetProvider(val cdktf.TerraformProvider) {
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
func CloudfrontMonitoringSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontMonitoringSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontMonitoringSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontMonitoringSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscription) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontMonitoringSubscription) PutMonitoringSubscription(value *CloudfrontMonitoringSubscriptionMonitoringSubscription) {
	_jsii_.InvokeVoid(
		c,
		"putMonitoringSubscription",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontMonitoringSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontMonitoringSubscription) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscription) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscription) ToString() *string {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudfrontMonitoringSubscriptionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_monitoring_subscription.html#distribution_id CloudfrontMonitoringSubscription#distribution_id}.
	DistributionId *string `json:"distributionId"`
	// monitoring_subscription block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_monitoring_subscription.html#monitoring_subscription CloudfrontMonitoringSubscription#monitoring_subscription}
	MonitoringSubscription *CloudfrontMonitoringSubscriptionMonitoringSubscription `json:"monitoringSubscription"`
}

type CloudfrontMonitoringSubscriptionMonitoringSubscription struct {
	// realtime_metrics_subscription_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_monitoring_subscription.html#realtime_metrics_subscription_config CloudfrontMonitoringSubscription#realtime_metrics_subscription_config}
	RealtimeMetricsSubscriptionConfig *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig `json:"realtimeMetricsSubscriptionConfig"`
}

type CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RealtimeMetricsSubscriptionConfig() CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference
	RealtimeMetricsSubscriptionConfigInput() *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig
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
	PutRealtimeMetricsSubscriptionConfig(value *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig)
}

// The jsii proxy struct for CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference
type jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) RealtimeMetricsSubscriptionConfig() CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference {
	var returns CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference
	_jsii_.Get(
		j,
		"realtimeMetricsSubscriptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) RealtimeMetricsSubscriptionConfigInput() *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig {
	var returns *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig
	_jsii_.Get(
		j,
		"realtimeMetricsSubscriptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference_Override(c CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference) PutRealtimeMetricsSubscriptionConfig(value *CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig) {
	_jsii_.InvokeVoid(
		c,
		"putRealtimeMetricsSubscriptionConfig",
		[]interface{}{value},
	)
}

type CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_monitoring_subscription.html#realtime_metrics_subscription_status CloudfrontMonitoringSubscription#realtime_metrics_subscription_status}.
	RealtimeMetricsSubscriptionStatus *string `json:"realtimeMetricsSubscriptionStatus"`
}

type CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RealtimeMetricsSubscriptionStatus() *string
	SetRealtimeMetricsSubscriptionStatus(val *string)
	RealtimeMetricsSubscriptionStatusInput() *string
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
}

// The jsii proxy struct for CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference
type jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) RealtimeMetricsSubscriptionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeMetricsSubscriptionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) RealtimeMetricsSubscriptionStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realtimeMetricsSubscriptionStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference_Override(c CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) SetRealtimeMetricsSubscriptionStatus(val *string) {
	_jsii_.Set(
		j,
		"realtimeMetricsSubscriptionStatus",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_access_identity.html aws_cloudfront_origin_access_identity}.
type CloudfrontOriginAccessIdentity interface {
	cdktf.TerraformResource
	CallerReference() *string
	CdktfStack() cdktf.TerraformStack
	CloudfrontAccessIdentityPath() *string
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
	Fqn() *string
	FriendlyUniqueId() *string
	IamArn() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	S3CanonicalUserId() *string
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
	ResetComment()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfrontOriginAccessIdentity
type jsiiProxy_CloudfrontOriginAccessIdentity struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) CloudfrontAccessIdentityPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontAccessIdentityPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) IamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) S3CanonicalUserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3CanonicalUserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_access_identity.html aws_cloudfront_origin_access_identity} Resource.
func NewCloudfrontOriginAccessIdentity(scope constructs.Construct, id *string, config *CloudfrontOriginAccessIdentityConfig) CloudfrontOriginAccessIdentity {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontOriginAccessIdentity{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginAccessIdentity",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_access_identity.html aws_cloudfront_origin_access_identity} Resource.
func NewCloudfrontOriginAccessIdentity_Override(c CloudfrontOriginAccessIdentity, scope constructs.Construct, id *string, config *CloudfrontOriginAccessIdentityConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginAccessIdentity",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginAccessIdentity) SetProvider(val cdktf.TerraformProvider) {
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
func CloudfrontOriginAccessIdentity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontOriginAccessIdentity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontOriginAccessIdentity_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontOriginAccessIdentity",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontOriginAccessIdentity) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontOriginAccessIdentity) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) ToString() *string {
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
func (c *jsiiProxy_CloudfrontOriginAccessIdentity) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudfrontOriginAccessIdentityConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_access_identity.html#comment CloudfrontOriginAccessIdentity#comment}.
	Comment *string `json:"comment"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html aws_cloudfront_origin_request_policy}.
type CloudfrontOriginRequestPolicy interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	CookiesConfig() CloudfrontOriginRequestPolicyCookiesConfigOutputReference
	CookiesConfigInput() *CloudfrontOriginRequestPolicyCookiesConfig
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HeadersConfig() CloudfrontOriginRequestPolicyHeadersConfigOutputReference
	HeadersConfigInput() *CloudfrontOriginRequestPolicyHeadersConfig
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	QueryStringsConfig() CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference
	QueryStringsConfigInput() *CloudfrontOriginRequestPolicyQueryStringsConfig
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
	PutCookiesConfig(value *CloudfrontOriginRequestPolicyCookiesConfig)
	PutHeadersConfig(value *CloudfrontOriginRequestPolicyHeadersConfig)
	PutQueryStringsConfig(value *CloudfrontOriginRequestPolicyQueryStringsConfig)
	ResetComment()
	ResetEtag()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfrontOriginRequestPolicy
type jsiiProxy_CloudfrontOriginRequestPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) CookiesConfig() CloudfrontOriginRequestPolicyCookiesConfigOutputReference {
	var returns CloudfrontOriginRequestPolicyCookiesConfigOutputReference
	_jsii_.Get(
		j,
		"cookiesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) CookiesConfigInput() *CloudfrontOriginRequestPolicyCookiesConfig {
	var returns *CloudfrontOriginRequestPolicyCookiesConfig
	_jsii_.Get(
		j,
		"cookiesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) HeadersConfig() CloudfrontOriginRequestPolicyHeadersConfigOutputReference {
	var returns CloudfrontOriginRequestPolicyHeadersConfigOutputReference
	_jsii_.Get(
		j,
		"headersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) HeadersConfigInput() *CloudfrontOriginRequestPolicyHeadersConfig {
	var returns *CloudfrontOriginRequestPolicyHeadersConfig
	_jsii_.Get(
		j,
		"headersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) QueryStringsConfig() CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference {
	var returns CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference
	_jsii_.Get(
		j,
		"queryStringsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) QueryStringsConfigInput() *CloudfrontOriginRequestPolicyQueryStringsConfig {
	var returns *CloudfrontOriginRequestPolicyQueryStringsConfig
	_jsii_.Get(
		j,
		"queryStringsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html aws_cloudfront_origin_request_policy} Resource.
func NewCloudfrontOriginRequestPolicy(scope constructs.Construct, id *string, config *CloudfrontOriginRequestPolicyConfig) CloudfrontOriginRequestPolicy {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontOriginRequestPolicy{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html aws_cloudfront_origin_request_policy} Resource.
func NewCloudfrontOriginRequestPolicy_Override(c CloudfrontOriginRequestPolicy, scope constructs.Construct, id *string, config *CloudfrontOriginRequestPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicy",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) SetEtag(val *string) {
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicy) SetProvider(val cdktf.TerraformProvider) {
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
func CloudfrontOriginRequestPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontOriginRequestPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicy) PutCookiesConfig(value *CloudfrontOriginRequestPolicyCookiesConfig) {
	_jsii_.InvokeVoid(
		c,
		"putCookiesConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicy) PutHeadersConfig(value *CloudfrontOriginRequestPolicyHeadersConfig) {
	_jsii_.InvokeVoid(
		c,
		"putHeadersConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicy) PutQueryStringsConfig(value *CloudfrontOriginRequestPolicyQueryStringsConfig) {
	_jsii_.InvokeVoid(
		c,
		"putQueryStringsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicy) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicy) ResetEtag() {
	_jsii_.InvokeVoid(
		c,
		"resetEtag",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) ToString() *string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudfrontOriginRequestPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// cookies_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#cookies_config CloudfrontOriginRequestPolicy#cookies_config}
	CookiesConfig *CloudfrontOriginRequestPolicyCookiesConfig `json:"cookiesConfig"`
	// headers_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#headers_config CloudfrontOriginRequestPolicy#headers_config}
	HeadersConfig *CloudfrontOriginRequestPolicyHeadersConfig `json:"headersConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#name CloudfrontOriginRequestPolicy#name}.
	Name *string `json:"name"`
	// query_strings_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#query_strings_config CloudfrontOriginRequestPolicy#query_strings_config}
	QueryStringsConfig *CloudfrontOriginRequestPolicyQueryStringsConfig `json:"queryStringsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#comment CloudfrontOriginRequestPolicy#comment}.
	Comment *string `json:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#etag CloudfrontOriginRequestPolicy#etag}.
	Etag *string `json:"etag"`
}

type CloudfrontOriginRequestPolicyCookiesConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#cookie_behavior CloudfrontOriginRequestPolicy#cookie_behavior}.
	CookieBehavior *string `json:"cookieBehavior"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#cookies CloudfrontOriginRequestPolicy#cookies}
	Cookies *CloudfrontOriginRequestPolicyCookiesConfigCookies `json:"cookies"`
}

type CloudfrontOriginRequestPolicyCookiesConfigCookies struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#items CloudfrontOriginRequestPolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference
type jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference_Override(c CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

type CloudfrontOriginRequestPolicyCookiesConfigOutputReference interface {
	cdktf.ComplexObject
	CookieBehavior() *string
	SetCookieBehavior(val *string)
	CookieBehaviorInput() *string
	Cookies() CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference
	CookiesInput() *CloudfrontOriginRequestPolicyCookiesConfigCookies
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
	PutCookies(value *CloudfrontOriginRequestPolicyCookiesConfigCookies)
	ResetCookies()
}

// The jsii proxy struct for CloudfrontOriginRequestPolicyCookiesConfigOutputReference
type jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) CookieBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) CookieBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) Cookies() CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference {
	var returns CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) CookiesInput() *CloudfrontOriginRequestPolicyCookiesConfigCookies {
	var returns *CloudfrontOriginRequestPolicyCookiesConfigCookies
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontOriginRequestPolicyCookiesConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontOriginRequestPolicyCookiesConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyCookiesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontOriginRequestPolicyCookiesConfigOutputReference_Override(c CloudfrontOriginRequestPolicyCookiesConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyCookiesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) SetCookieBehavior(val *string) {
	_jsii_.Set(
		j,
		"cookieBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) PutCookies(value *CloudfrontOriginRequestPolicyCookiesConfigCookies) {
	_jsii_.InvokeVoid(
		c,
		"putCookies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference) ResetCookies() {
	_jsii_.InvokeVoid(
		c,
		"resetCookies",
		nil, // no parameters
	)
}

type CloudfrontOriginRequestPolicyHeadersConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#header_behavior CloudfrontOriginRequestPolicy#header_behavior}.
	HeaderBehavior *string `json:"headerBehavior"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#headers CloudfrontOriginRequestPolicy#headers}
	Headers *CloudfrontOriginRequestPolicyHeadersConfigHeaders `json:"headers"`
}

type CloudfrontOriginRequestPolicyHeadersConfigHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#items CloudfrontOriginRequestPolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference
type jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference_Override(c CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

type CloudfrontOriginRequestPolicyHeadersConfigOutputReference interface {
	cdktf.ComplexObject
	HeaderBehavior() *string
	SetHeaderBehavior(val *string)
	HeaderBehaviorInput() *string
	Headers() CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference
	HeadersInput() *CloudfrontOriginRequestPolicyHeadersConfigHeaders
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
	PutHeaders(value *CloudfrontOriginRequestPolicyHeadersConfigHeaders)
	ResetHeaderBehavior()
	ResetHeaders()
}

// The jsii proxy struct for CloudfrontOriginRequestPolicyHeadersConfigOutputReference
type jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) HeaderBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) HeaderBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) Headers() CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference {
	var returns CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) HeadersInput() *CloudfrontOriginRequestPolicyHeadersConfigHeaders {
	var returns *CloudfrontOriginRequestPolicyHeadersConfigHeaders
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontOriginRequestPolicyHeadersConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontOriginRequestPolicyHeadersConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontOriginRequestPolicyHeadersConfigOutputReference_Override(c CloudfrontOriginRequestPolicyHeadersConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) SetHeaderBehavior(val *string) {
	_jsii_.Set(
		j,
		"headerBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) PutHeaders(value *CloudfrontOriginRequestPolicyHeadersConfigHeaders) {
	_jsii_.InvokeVoid(
		c,
		"putHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) ResetHeaderBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaderBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaders",
		nil, // no parameters
	)
}

type CloudfrontOriginRequestPolicyQueryStringsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#query_string_behavior CloudfrontOriginRequestPolicy#query_string_behavior}.
	QueryStringBehavior *string `json:"queryStringBehavior"`
	// query_strings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#query_strings CloudfrontOriginRequestPolicy#query_strings}
	QueryStrings *CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings `json:"queryStrings"`
}

type CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	QueryStringBehavior() *string
	SetQueryStringBehavior(val *string)
	QueryStringBehaviorInput() *string
	QueryStrings() CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference
	QueryStringsInput() *CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings
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
	PutQueryStrings(value *CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings)
	ResetQueryStrings()
}

// The jsii proxy struct for CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference
type jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) QueryStringBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) QueryStringBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) QueryStrings() CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference {
	var returns CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) QueryStringsInput() *CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings {
	var returns *CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings
	_jsii_.Get(
		j,
		"queryStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontOriginRequestPolicyQueryStringsConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontOriginRequestPolicyQueryStringsConfigOutputReference_Override(c CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) SetQueryStringBehavior(val *string) {
	_jsii_.Set(
		j,
		"queryStringBehavior",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) PutQueryStrings(value *CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) {
	_jsii_.InvokeVoid(
		c,
		"putQueryStrings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference) ResetQueryStrings() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStrings",
		nil, // no parameters
	)
}

type CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_request_policy.html#items CloudfrontOriginRequestPolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference
type jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference_Override(c CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_public_key.html aws_cloudfront_public_key}.
type CloudfrontPublicKey interface {
	cdktf.TerraformResource
	CallerReference() *string
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EncodedKey() *string
	SetEncodedKey(val *string)
	EncodedKeyInput() *string
	Etag() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
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
	ResetComment()
	ResetName()
	ResetNamePrefix()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfrontPublicKey
type jsiiProxy_CloudfrontPublicKey struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontPublicKey) CallerReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) EncodedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) EncodedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontPublicKey) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_public_key.html aws_cloudfront_public_key} Resource.
func NewCloudfrontPublicKey(scope constructs.Construct, id *string, config *CloudfrontPublicKeyConfig) CloudfrontPublicKey {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontPublicKey{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontPublicKey",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_public_key.html aws_cloudfront_public_key} Resource.
func NewCloudfrontPublicKey_Override(c CloudfrontPublicKey, scope constructs.Construct, id *string, config *CloudfrontPublicKeyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontPublicKey",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontPublicKey) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontPublicKey) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontPublicKey) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontPublicKey) SetEncodedKey(val *string) {
	_jsii_.Set(
		j,
		"encodedKey",
		val,
	)
}

func (j *jsiiProxy_CloudfrontPublicKey) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontPublicKey) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfrontPublicKey) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_CloudfrontPublicKey) SetProvider(val cdktf.TerraformProvider) {
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
func CloudfrontPublicKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontPublicKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontPublicKey_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontPublicKey",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontPublicKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontPublicKey) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontPublicKey) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontPublicKey) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontPublicKey) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontPublicKey) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontPublicKey) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontPublicKey) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontPublicKey) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontPublicKey) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		c,
		"resetNamePrefix",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontPublicKey) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontPublicKey) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontPublicKey) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontPublicKey) ToString() *string {
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
func (c *jsiiProxy_CloudfrontPublicKey) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudfrontPublicKeyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_public_key.html#encoded_key CloudfrontPublicKey#encoded_key}.
	EncodedKey *string `json:"encodedKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_public_key.html#comment CloudfrontPublicKey#comment}.
	Comment *string `json:"comment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_public_key.html#name CloudfrontPublicKey#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_public_key.html#name_prefix CloudfrontPublicKey#name_prefix}.
	NamePrefix *string `json:"namePrefix"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html aws_cloudfront_realtime_log_config}.
type CloudfrontRealtimeLogConfig interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Endpoint() CloudfrontRealtimeLogConfigEndpointOutputReference
	EndpointInput() *CloudfrontRealtimeLogConfigEndpoint
	Fields() *[]*string
	SetFields(val *[]*string)
	FieldsInput() *[]*string
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
	SamplingRate() *float64
	SetSamplingRate(val *float64)
	SamplingRateInput() *float64
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
	PutEndpoint(value *CloudfrontRealtimeLogConfigEndpoint)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfrontRealtimeLogConfig
type jsiiProxy_CloudfrontRealtimeLogConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Endpoint() CloudfrontRealtimeLogConfigEndpointOutputReference {
	var returns CloudfrontRealtimeLogConfigEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) EndpointInput() *CloudfrontRealtimeLogConfigEndpoint {
	var returns *CloudfrontRealtimeLogConfigEndpoint
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Fields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) FieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) SamplingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) SamplingRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html aws_cloudfront_realtime_log_config} Resource.
func NewCloudfrontRealtimeLogConfig(scope constructs.Construct, id *string, config *CloudfrontRealtimeLogConfigConfig) CloudfrontRealtimeLogConfig {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontRealtimeLogConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html aws_cloudfront_realtime_log_config} Resource.
func NewCloudfrontRealtimeLogConfig_Override(c CloudfrontRealtimeLogConfig, scope constructs.Construct, id *string, config *CloudfrontRealtimeLogConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfig",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) SetFields(val *[]*string) {
	_jsii_.Set(
		j,
		"fields",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfig) SetSamplingRate(val *float64) {
	_jsii_.Set(
		j,
		"samplingRate",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudfrontRealtimeLogConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontRealtimeLogConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontRealtimeLogConfig) PutEndpoint(value *CloudfrontRealtimeLogConfigEndpoint) {
	_jsii_.InvokeVoid(
		c,
		"putEndpoint",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontRealtimeLogConfig) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) ToString() *string {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudfrontRealtimeLogConfigConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html#endpoint CloudfrontRealtimeLogConfig#endpoint}
	Endpoint *CloudfrontRealtimeLogConfigEndpoint `json:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html#fields CloudfrontRealtimeLogConfig#fields}.
	Fields *[]*string `json:"fields"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html#name CloudfrontRealtimeLogConfig#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html#sampling_rate CloudfrontRealtimeLogConfig#sampling_rate}.
	SamplingRate *float64 `json:"samplingRate"`
}

type CloudfrontRealtimeLogConfigEndpoint struct {
	// kinesis_stream_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html#kinesis_stream_config CloudfrontRealtimeLogConfig#kinesis_stream_config}
	KinesisStreamConfig *CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig `json:"kinesisStreamConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html#stream_type CloudfrontRealtimeLogConfig#stream_type}.
	StreamType *string `json:"streamType"`
}

type CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html#role_arn CloudfrontRealtimeLogConfig#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_realtime_log_config.html#stream_arn CloudfrontRealtimeLogConfig#stream_arn}.
	StreamArn *string `json:"streamArn"`
}

type CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	StreamArn() *string
	SetStreamArn(val *string)
	StreamArnInput() *string
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
}

// The jsii proxy struct for CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference
type jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) StreamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference_Override(c CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) SetStreamArn(val *string) {
	_jsii_.Set(
		j,
		"streamArn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontRealtimeLogConfigEndpointOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KinesisStreamConfig() CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference
	KinesisStreamConfigInput() *CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig
	StreamType() *string
	SetStreamType(val *string)
	StreamTypeInput() *string
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
	PutKinesisStreamConfig(value *CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig)
}

// The jsii proxy struct for CloudfrontRealtimeLogConfigEndpointOutputReference
type jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) KinesisStreamConfig() CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference {
	var returns CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) KinesisStreamConfigInput() *CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig {
	var returns *CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig
	_jsii_.Get(
		j,
		"kinesisStreamConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) StreamType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) StreamTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontRealtimeLogConfigEndpointOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontRealtimeLogConfigEndpointOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfigEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontRealtimeLogConfigEndpointOutputReference_Override(c CloudfrontRealtimeLogConfigEndpointOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfigEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) SetStreamType(val *string) {
	_jsii_.Set(
		j,
		"streamType",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference) PutKinesisStreamConfig(value *CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig) {
	_jsii_.InvokeVoid(
		c,
		"putKinesisStreamConfig",
		[]interface{}{value},
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html aws_cloudfront_response_headers_policy}.
type CloudfrontResponseHeadersPolicy interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	SetComment(val *string)
	CommentInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	CorsConfig() CloudfrontResponseHeadersPolicyCorsConfigOutputReference
	CorsConfigInput() *CloudfrontResponseHeadersPolicyCorsConfig
	Count() interface{}
	SetCount(val interface{})
	CustomHeadersConfig() CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference
	CustomHeadersConfigInput() *CloudfrontResponseHeadersPolicyCustomHeadersConfig
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
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
	SecurityHeadersConfig() CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference
	SecurityHeadersConfigInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfig
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
	PutCorsConfig(value *CloudfrontResponseHeadersPolicyCorsConfig)
	PutCustomHeadersConfig(value *CloudfrontResponseHeadersPolicyCustomHeadersConfig)
	PutSecurityHeadersConfig(value *CloudfrontResponseHeadersPolicySecurityHeadersConfig)
	ResetComment()
	ResetCorsConfig()
	ResetCustomHeadersConfig()
	ResetEtag()
	ResetOverrideLogicalId()
	ResetSecurityHeadersConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicy
type jsiiProxy_CloudfrontResponseHeadersPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CommentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CorsConfig() CloudfrontResponseHeadersPolicyCorsConfigOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigOutputReference
	_jsii_.Get(
		j,
		"corsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CorsConfigInput() *CloudfrontResponseHeadersPolicyCorsConfig {
	var returns *CloudfrontResponseHeadersPolicyCorsConfig
	_jsii_.Get(
		j,
		"corsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CustomHeadersConfig() CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference {
	var returns CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference
	_jsii_.Get(
		j,
		"customHeadersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) CustomHeadersConfigInput() *CloudfrontResponseHeadersPolicyCustomHeadersConfig {
	var returns *CloudfrontResponseHeadersPolicyCustomHeadersConfig
	_jsii_.Get(
		j,
		"customHeadersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SecurityHeadersConfig() CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference
	_jsii_.Get(
		j,
		"securityHeadersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SecurityHeadersConfigInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfig {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfig
	_jsii_.Get(
		j,
		"securityHeadersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html aws_cloudfront_response_headers_policy} Resource.
func NewCloudfrontResponseHeadersPolicy(scope constructs.Construct, id *string, config *CloudfrontResponseHeadersPolicyConfig) CloudfrontResponseHeadersPolicy {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicy{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html aws_cloudfront_response_headers_policy} Resource.
func NewCloudfrontResponseHeadersPolicy_Override(c CloudfrontResponseHeadersPolicy, scope constructs.Construct, id *string, config *CloudfrontResponseHeadersPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicy",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SetComment(val *string) {
	_jsii_.Set(
		j,
		"comment",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SetEtag(val *string) {
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicy) SetProvider(val cdktf.TerraformProvider) {
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
func CloudfrontResponseHeadersPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudfrontResponseHeadersPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) PutCorsConfig(value *CloudfrontResponseHeadersPolicyCorsConfig) {
	_jsii_.InvokeVoid(
		c,
		"putCorsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) PutCustomHeadersConfig(value *CloudfrontResponseHeadersPolicyCustomHeadersConfig) {
	_jsii_.InvokeVoid(
		c,
		"putCustomHeadersConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) PutSecurityHeadersConfig(value *CloudfrontResponseHeadersPolicySecurityHeadersConfig) {
	_jsii_.InvokeVoid(
		c,
		"putSecurityHeadersConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetComment() {
	_jsii_.InvokeVoid(
		c,
		"resetComment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetCorsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCorsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetCustomHeadersConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomHeadersConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetEtag() {
	_jsii_.InvokeVoid(
		c,
		"resetEtag",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ResetSecurityHeadersConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityHeadersConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ToString() *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudfrontResponseHeadersPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#name CloudfrontResponseHeadersPolicy#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#comment CloudfrontResponseHeadersPolicy#comment}.
	Comment *string `json:"comment"`
	// cors_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#cors_config CloudfrontResponseHeadersPolicy#cors_config}
	CorsConfig *CloudfrontResponseHeadersPolicyCorsConfig `json:"corsConfig"`
	// custom_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#custom_headers_config CloudfrontResponseHeadersPolicy#custom_headers_config}
	CustomHeadersConfig *CloudfrontResponseHeadersPolicyCustomHeadersConfig `json:"customHeadersConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#etag CloudfrontResponseHeadersPolicy#etag}.
	Etag *string `json:"etag"`
	// security_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#security_headers_config CloudfrontResponseHeadersPolicy#security_headers_config}
	SecurityHeadersConfig *CloudfrontResponseHeadersPolicySecurityHeadersConfig `json:"securityHeadersConfig"`
}

type CloudfrontResponseHeadersPolicyCorsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#access_control_allow_credentials CloudfrontResponseHeadersPolicy#access_control_allow_credentials}.
	AccessControlAllowCredentials interface{} `json:"accessControlAllowCredentials"`
	// access_control_allow_headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#access_control_allow_headers CloudfrontResponseHeadersPolicy#access_control_allow_headers}
	AccessControlAllowHeaders *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders `json:"accessControlAllowHeaders"`
	// access_control_allow_methods block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#access_control_allow_methods CloudfrontResponseHeadersPolicy#access_control_allow_methods}
	AccessControlAllowMethods *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods `json:"accessControlAllowMethods"`
	// access_control_allow_origins block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#access_control_allow_origins CloudfrontResponseHeadersPolicy#access_control_allow_origins}
	AccessControlAllowOrigins *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins `json:"accessControlAllowOrigins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#origin_override CloudfrontResponseHeadersPolicy#origin_override}.
	OriginOverride interface{} `json:"originOverride"`
	// access_control_expose_headers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#access_control_expose_headers CloudfrontResponseHeadersPolicy#access_control_expose_headers}
	AccessControlExposeHeaders *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders `json:"accessControlExposeHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#access_control_max_age_sec CloudfrontResponseHeadersPolicy#access_control_max_age_sec}.
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec"`
}

type CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#items CloudfrontResponseHeadersPolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference_Override(c CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

type CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#items CloudfrontResponseHeadersPolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference_Override(c CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

type CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#items CloudfrontResponseHeadersPolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference_Override(c CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

type CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#items CloudfrontResponseHeadersPolicy#items}.
	Items *[]*string `json:"items"`
}

type CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*string
	SetItems(val *[]*string)
	ItemsInput() *[]*string
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) ItemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference_Override(c CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) SetItems(val *[]*string) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

type CloudfrontResponseHeadersPolicyCorsConfigOutputReference interface {
	cdktf.ComplexObject
	AccessControlAllowCredentials() interface{}
	SetAccessControlAllowCredentials(val interface{})
	AccessControlAllowCredentialsInput() interface{}
	AccessControlAllowHeaders() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference
	AccessControlAllowHeadersInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders
	AccessControlAllowMethods() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference
	AccessControlAllowMethodsInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods
	AccessControlAllowOrigins() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference
	AccessControlAllowOriginsInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins
	AccessControlExposeHeaders() CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference
	AccessControlExposeHeadersInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders
	AccessControlMaxAgeSec() *float64
	SetAccessControlMaxAgeSec(val *float64)
	AccessControlMaxAgeSecInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OriginOverride() interface{}
	SetOriginOverride(val interface{})
	OriginOverrideInput() interface{}
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
	PutAccessControlAllowHeaders(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders)
	PutAccessControlAllowMethods(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods)
	PutAccessControlAllowOrigins(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins)
	PutAccessControlExposeHeaders(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders)
	ResetAccessControlExposeHeaders()
	ResetAccessControlMaxAgeSec()
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicyCorsConfigOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowHeaders() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowHeadersInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders {
	var returns *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders
	_jsii_.Get(
		j,
		"accessControlAllowHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowMethods() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowMethodsInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods {
	var returns *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods
	_jsii_.Get(
		j,
		"accessControlAllowMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowOrigins() CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference
	_jsii_.Get(
		j,
		"accessControlAllowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlAllowOriginsInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins {
	var returns *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins
	_jsii_.Get(
		j,
		"accessControlAllowOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlExposeHeaders() CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference {
	var returns CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference
	_jsii_.Get(
		j,
		"accessControlExposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlExposeHeadersInput() *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders {
	var returns *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders
	_jsii_.Get(
		j,
		"accessControlExposeHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) AccessControlMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) OriginOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) OriginOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicyCorsConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicyCorsConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicyCorsConfigOutputReference_Override(c CloudfrontResponseHeadersPolicyCorsConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) SetAccessControlAllowCredentials(val interface{}) {
	_jsii_.Set(
		j,
		"accessControlAllowCredentials",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) SetAccessControlMaxAgeSec(val *float64) {
	_jsii_.Set(
		j,
		"accessControlMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) SetOriginOverride(val interface{}) {
	_jsii_.Set(
		j,
		"originOverride",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) PutAccessControlAllowHeaders(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) {
	_jsii_.InvokeVoid(
		c,
		"putAccessControlAllowHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) PutAccessControlAllowMethods(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) {
	_jsii_.InvokeVoid(
		c,
		"putAccessControlAllowMethods",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) PutAccessControlAllowOrigins(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) {
	_jsii_.InvokeVoid(
		c,
		"putAccessControlAllowOrigins",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) PutAccessControlExposeHeaders(value *CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) {
	_jsii_.InvokeVoid(
		c,
		"putAccessControlExposeHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) ResetAccessControlExposeHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessControlExposeHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference) ResetAccessControlMaxAgeSec() {
	_jsii_.InvokeVoid(
		c,
		"resetAccessControlMaxAgeSec",
		nil, // no parameters
	)
}

type CloudfrontResponseHeadersPolicyCustomHeadersConfig struct {
	// items block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#items CloudfrontResponseHeadersPolicy#items}
	Items *[]*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems `json:"items"`
}

type CloudfrontResponseHeadersPolicyCustomHeadersConfigItems struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#header CloudfrontResponseHeadersPolicy#header}.
	Header *string `json:"header"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `json:"override"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#value CloudfrontResponseHeadersPolicy#value}.
	Value *string `json:"value"`
}

type CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Items() *[]*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems
	SetItems(val *[]*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems)
	ItemsInput() *[]*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems
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
	ResetItems()
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) Items() *[]*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems {
	var returns *[]*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) ItemsInput() *[]*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems {
	var returns *[]*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference_Override(c CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) SetItems(val *[]*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems) {
	_jsii_.Set(
		j,
		"items",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfig struct {
	// content_security_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#content_security_policy CloudfrontResponseHeadersPolicy#content_security_policy}
	ContentSecurityPolicy *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy `json:"contentSecurityPolicy"`
	// content_type_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#content_type_options CloudfrontResponseHeadersPolicy#content_type_options}
	ContentTypeOptions *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions `json:"contentTypeOptions"`
	// frame_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#frame_options CloudfrontResponseHeadersPolicy#frame_options}
	FrameOptions *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions `json:"frameOptions"`
	// referrer_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#referrer_policy CloudfrontResponseHeadersPolicy#referrer_policy}
	ReferrerPolicy *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy `json:"referrerPolicy"`
	// strict_transport_security block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#strict_transport_security CloudfrontResponseHeadersPolicy#strict_transport_security}
	StrictTransportSecurity *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity `json:"strictTransportSecurity"`
	// xss_protection block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#xss_protection CloudfrontResponseHeadersPolicy#xss_protection}
	XssProtection *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection `json:"xssProtection"`
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#content_security_policy CloudfrontResponseHeadersPolicy#content_security_policy}.
	ContentSecurityPolicy *string `json:"contentSecurityPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `json:"override"`
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference interface {
	cdktf.ComplexObject
	ContentSecurityPolicy() *string
	SetContentSecurityPolicy(val *string)
	ContentSecurityPolicyInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Override() interface{}
	SetOverride(val interface{})
	OverrideInput() interface{}
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
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) ContentSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) ContentSecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) SetContentSecurityPolicy(val *string) {
	_jsii_.Set(
		j,
		"contentSecurityPolicy",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) SetOverride(val interface{}) {
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `json:"override"`
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Override() interface{}
	SetOverride(val interface{})
	OverrideInput() interface{}
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
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) SetOverride(val interface{}) {
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#frame_option CloudfrontResponseHeadersPolicy#frame_option}.
	FrameOption *string `json:"frameOption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `json:"override"`
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference interface {
	cdktf.ComplexObject
	FrameOption() *string
	SetFrameOption(val *string)
	FrameOptionInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Override() interface{}
	SetOverride(val interface{})
	OverrideInput() interface{}
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
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) FrameOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) FrameOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) SetFrameOption(val *string) {
	_jsii_.Set(
		j,
		"frameOption",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) SetOverride(val interface{}) {
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference interface {
	cdktf.ComplexObject
	ContentSecurityPolicy() CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference
	ContentSecurityPolicyInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy
	ContentTypeOptions() CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference
	ContentTypeOptionsInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions
	FrameOptions() CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference
	FrameOptionsInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ReferrerPolicy() CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference
	ReferrerPolicyInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy
	StrictTransportSecurity() CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference
	StrictTransportSecurityInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	XssProtection() CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference
	XssProtectionInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutContentSecurityPolicy(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy)
	PutContentTypeOptions(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions)
	PutFrameOptions(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions)
	PutReferrerPolicy(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy)
	PutStrictTransportSecurity(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity)
	PutXssProtection(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection)
	ResetContentSecurityPolicy()
	ResetContentTypeOptions()
	ResetFrameOptions()
	ResetReferrerPolicy()
	ResetStrictTransportSecurity()
	ResetXssProtection()
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ContentSecurityPolicy() CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference
	_jsii_.Get(
		j,
		"contentSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ContentSecurityPolicyInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy
	_jsii_.Get(
		j,
		"contentSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ContentTypeOptions() CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"contentTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ContentTypeOptionsInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions
	_jsii_.Get(
		j,
		"contentTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) FrameOptions() CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference
	_jsii_.Get(
		j,
		"frameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) FrameOptionsInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions
	_jsii_.Get(
		j,
		"frameOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ReferrerPolicy() CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference
	_jsii_.Get(
		j,
		"referrerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ReferrerPolicyInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy
	_jsii_.Get(
		j,
		"referrerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) StrictTransportSecurity() CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference
	_jsii_.Get(
		j,
		"strictTransportSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) StrictTransportSecurityInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity
	_jsii_.Get(
		j,
		"strictTransportSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) XssProtection() CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference {
	var returns CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference
	_jsii_.Get(
		j,
		"xssProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) XssProtectionInput() *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection
	_jsii_.Get(
		j,
		"xssProtectionInput",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutContentSecurityPolicy(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putContentSecurityPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutContentTypeOptions(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) {
	_jsii_.InvokeVoid(
		c,
		"putContentTypeOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutFrameOptions(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) {
	_jsii_.InvokeVoid(
		c,
		"putFrameOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutReferrerPolicy(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) {
	_jsii_.InvokeVoid(
		c,
		"putReferrerPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutStrictTransportSecurity(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) {
	_jsii_.InvokeVoid(
		c,
		"putStrictTransportSecurity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) PutXssProtection(value *CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) {
	_jsii_.InvokeVoid(
		c,
		"putXssProtection",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetContentSecurityPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetContentSecurityPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetContentTypeOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetContentTypeOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetFrameOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetFrameOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetReferrerPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetReferrerPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetStrictTransportSecurity() {
	_jsii_.InvokeVoid(
		c,
		"resetStrictTransportSecurity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference) ResetXssProtection() {
	_jsii_.InvokeVoid(
		c,
		"resetXssProtection",
		nil, // no parameters
	)
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `json:"override"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#referrer_policy CloudfrontResponseHeadersPolicy#referrer_policy}.
	ReferrerPolicy *string `json:"referrerPolicy"`
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Override() interface{}
	SetOverride(val interface{})
	OverrideInput() interface{}
	ReferrerPolicy() *string
	SetReferrerPolicy(val *string)
	ReferrerPolicyInput() *string
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
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) ReferrerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referrerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) ReferrerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referrerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) SetOverride(val interface{}) {
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) SetReferrerPolicy(val *string) {
	_jsii_.Set(
		j,
		"referrerPolicy",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#access_control_max_age_sec CloudfrontResponseHeadersPolicy#access_control_max_age_sec}.
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `json:"override"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#include_subdomains CloudfrontResponseHeadersPolicy#include_subdomains}.
	IncludeSubdomains interface{} `json:"includeSubdomains"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#preload CloudfrontResponseHeadersPolicy#preload}.
	Preload interface{} `json:"preload"`
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference interface {
	cdktf.ComplexObject
	AccessControlMaxAgeSec() *float64
	SetAccessControlMaxAgeSec(val *float64)
	AccessControlMaxAgeSecInput() *float64
	IncludeSubdomains() interface{}
	SetIncludeSubdomains(val interface{})
	IncludeSubdomainsInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Override() interface{}
	SetOverride(val interface{})
	OverrideInput() interface{}
	Preload() interface{}
	SetPreload(val interface{})
	PreloadInput() interface{}
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
	ResetIncludeSubdomains()
	ResetPreload()
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) AccessControlMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) IncludeSubdomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) IncludeSubdomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) Preload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) PreloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetAccessControlMaxAgeSec(val *float64) {
	_jsii_.Set(
		j,
		"accessControlMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetIncludeSubdomains(val interface{}) {
	_jsii_.Set(
		j,
		"includeSubdomains",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetOverride(val interface{}) {
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetPreload(val interface{}) {
	_jsii_.Set(
		j,
		"preload",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) ResetIncludeSubdomains() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeSubdomains",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) ResetPreload() {
	_jsii_.InvokeVoid(
		c,
		"resetPreload",
		nil, // no parameters
	)
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `json:"override"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#protection CloudfrontResponseHeadersPolicy#protection}.
	Protection interface{} `json:"protection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#mode_block CloudfrontResponseHeadersPolicy#mode_block}.
	ModeBlock interface{} `json:"modeBlock"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudfront_response_headers_policy.html#report_uri CloudfrontResponseHeadersPolicy#report_uri}.
	ReportUri *string `json:"reportUri"`
}

type CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ModeBlock() interface{}
	SetModeBlock(val interface{})
	ModeBlockInput() interface{}
	Override() interface{}
	SetOverride(val interface{})
	OverrideInput() interface{}
	Protection() interface{}
	SetProtection(val interface{})
	ProtectionInput() interface{}
	ReportUri() *string
	SetReportUri(val *string)
	ReportUriInput() *string
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
	ResetModeBlock()
	ResetReportUri()
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ModeBlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modeBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ModeBlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modeBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) OverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) Protection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ReportUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ReportUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) SetModeBlock(val interface{}) {
	_jsii_.Set(
		j,
		"modeBlock",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) SetOverride(val interface{}) {
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) SetProtection(val interface{}) {
	_jsii_.Set(
		j,
		"protection",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) SetReportUri(val *string) {
	_jsii_.Set(
		j,
		"reportUri",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ResetModeBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetModeBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference) ResetReportUri() {
	_jsii_.InvokeVoid(
		c,
		"resetReportUri",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_cache_policy.html aws_cloudfront_cache_policy}.
type DataAwsCloudfrontCachePolicy interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DefaultTtl() *float64
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxTtl() *float64
	MinTtl() *float64
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
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ParametersInCacheKeyAndForwardedToOrigin(index *string) DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin
	ResetId()
	ResetName()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsCloudfrontCachePolicy
type jsiiProxy_DataAwsCloudfrontCachePolicy struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) MinTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_cache_policy.html aws_cloudfront_cache_policy} Data Source.
func NewDataAwsCloudfrontCachePolicy(scope constructs.Construct, id *string, config *DataAwsCloudfrontCachePolicyConfig) DataAwsCloudfrontCachePolicy {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontCachePolicy{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_cache_policy.html aws_cloudfront_cache_policy} Data Source.
func NewDataAwsCloudfrontCachePolicy_Override(d DataAwsCloudfrontCachePolicy, scope constructs.Construct, id *string, config *DataAwsCloudfrontCachePolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicy",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicy) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsCloudfrontCachePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCloudfrontCachePolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) ParametersInCacheKeyAndForwardedToOrigin(index *string) DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin {
	var returns DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin

	_jsii_.Invoke(
		d,
		"parametersInCacheKeyAndForwardedToOrigin",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) ToString() *string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCloudfrontCachePolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_cache_policy.html#id DataAwsCloudfrontCachePolicy#id}.
	Id *string `json:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_cache_policy.html#name DataAwsCloudfrontCachePolicy#name}.
	Name *string `json:"name"`
}

type DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	CookiesConfig() interface{}
	EnableAcceptEncodingBrotli() interface{}
	EnableAcceptEncodingGzip() interface{}
	HeadersConfig() interface{}
	QueryStringsConfig() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin
type jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) CookiesConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) EnableAcceptEncodingBrotli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) EnableAcceptEncodingGzip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) HeadersConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) QueryStringsConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin_Override(d DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	CookieBehavior() *string
	Cookies() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig
type jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) CookieBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) Cookies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig_Override(d DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies
type jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies_Override(d DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	HeaderBehavior() *string
	Headers() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig
type jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) HeaderBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) Headers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig_Override(d DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders
type jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders_Override(d DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	QueryStringBehavior() *string
	QueryStrings() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig
type jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) QueryStringBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) QueryStrings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig_Override(d DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings
type jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings_Override(d DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_distribution.html aws_cloudfront_distribution}.
type DataAwsCloudfrontDistribution interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DomainName() *string
	Enabled() interface{}
	Etag() *string
	Fqn() *string
	FriendlyUniqueId() *string
	HostedZoneId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InProgressValidationBatches() *float64
	LastModifiedTime() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	Tags() interface{}
	SetTags(val interface{})
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
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsCloudfrontDistribution
type jsiiProxy_DataAwsCloudfrontDistribution struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) InProgressValidationBatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inProgressValidationBatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_distribution.html aws_cloudfront_distribution} Data Source.
func NewDataAwsCloudfrontDistribution(scope constructs.Construct, id *string, config *DataAwsCloudfrontDistributionConfig) DataAwsCloudfrontDistribution {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontDistribution{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontDistribution",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_distribution.html aws_cloudfront_distribution} Data Source.
func NewDataAwsCloudfrontDistribution_Override(d DataAwsCloudfrontDistribution, scope constructs.Construct, id *string, config *DataAwsCloudfrontDistributionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontDistribution",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontDistribution) SetTags(val interface{}) {
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
func DataAwsCloudfrontDistribution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontDistribution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCloudfrontDistribution_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontDistribution",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontDistribution) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontDistribution) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontDistribution) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontDistribution) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontDistribution) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontDistribution) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontDistribution) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontDistribution) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontDistribution) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontDistribution) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontDistribution) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontDistribution) ToString() *string {
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
func (d *jsiiProxy_DataAwsCloudfrontDistribution) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCloudfrontDistributionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_distribution.html#id DataAwsCloudfrontDistribution#id}.
	Id *string `json:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_distribution.html#tags DataAwsCloudfrontDistribution#tags}.
	Tags interface{} `json:"tags"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_function.html aws_cloudfront_function}.
type DataAwsCloudfrontFunction interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Code() *string
	Comment() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastModifiedTime() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Runtime() *string
	Stage() *string
	SetStage(val *string)
	StageInput() *string
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

// The jsii proxy struct for DataAwsCloudfrontFunction
type jsiiProxy_DataAwsCloudfrontFunction struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Code() *string {
	var returns *string
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) StageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_function.html aws_cloudfront_function} Data Source.
func NewDataAwsCloudfrontFunction(scope constructs.Construct, id *string, config *DataAwsCloudfrontFunctionConfig) DataAwsCloudfrontFunction {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontFunction{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_function.html aws_cloudfront_function} Data Source.
func NewDataAwsCloudfrontFunction_Override(d DataAwsCloudfrontFunction, scope constructs.Construct, id *string, config *DataAwsCloudfrontFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontFunction",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontFunction) SetStage(val *string) {
	_jsii_.Set(
		j,
		"stage",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsCloudfrontFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCloudfrontFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontFunction) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontFunction) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontFunction) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontFunction) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontFunction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontFunction) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontFunction) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontFunction) ToString() *string {
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
func (d *jsiiProxy_DataAwsCloudfrontFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCloudfrontFunctionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_function.html#name DataAwsCloudfrontFunction#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_function.html#stage DataAwsCloudfrontFunction#stage}.
	Stage *string `json:"stage"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_log_delivery_canonical_user_id.html aws_cloudfront_log_delivery_canonical_user_id}.
type DataAwsCloudfrontLogDeliveryCanonicalUserId interface {
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	ResetRegion()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsCloudfrontLogDeliveryCanonicalUserId
type jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_log_delivery_canonical_user_id.html aws_cloudfront_log_delivery_canonical_user_id} Data Source.
func NewDataAwsCloudfrontLogDeliveryCanonicalUserId(scope constructs.Construct, id *string, config *DataAwsCloudfrontLogDeliveryCanonicalUserIdConfig) DataAwsCloudfrontLogDeliveryCanonicalUserId {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontLogDeliveryCanonicalUserId",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_log_delivery_canonical_user_id.html aws_cloudfront_log_delivery_canonical_user_id} Data Source.
func NewDataAwsCloudfrontLogDeliveryCanonicalUserId_Override(d DataAwsCloudfrontLogDeliveryCanonicalUserId, scope constructs.Construct, id *string, config *DataAwsCloudfrontLogDeliveryCanonicalUserIdConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontLogDeliveryCanonicalUserId",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsCloudfrontLogDeliveryCanonicalUserId_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontLogDeliveryCanonicalUserId",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCloudfrontLogDeliveryCanonicalUserId_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontLogDeliveryCanonicalUserId",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) ToString() *string {
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
func (d *jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCloudfrontLogDeliveryCanonicalUserIdConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_log_delivery_canonical_user_id.html#region DataAwsCloudfrontLogDeliveryCanonicalUserId#region}.
	Region *string `json:"region"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_origin_request_policy.html aws_cloudfront_origin_request_policy}.
type DataAwsCloudfrontOriginRequestPolicy interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
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
	AddOverride(path *string, value interface{})
	CookiesConfig(index *string) DataAwsCloudfrontOriginRequestPolicyCookiesConfig
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	HeadersConfig(index *string) DataAwsCloudfrontOriginRequestPolicyHeadersConfig
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	QueryStringsConfig(index *string) DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig
	ResetId()
	ResetName()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsCloudfrontOriginRequestPolicy
type jsiiProxy_DataAwsCloudfrontOriginRequestPolicy struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_origin_request_policy.html aws_cloudfront_origin_request_policy} Data Source.
func NewDataAwsCloudfrontOriginRequestPolicy(scope constructs.Construct, id *string, config *DataAwsCloudfrontOriginRequestPolicyConfig) DataAwsCloudfrontOriginRequestPolicy {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicy{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_origin_request_policy.html aws_cloudfront_origin_request_policy} Data Source.
func NewDataAwsCloudfrontOriginRequestPolicy_Override(d DataAwsCloudfrontOriginRequestPolicy, scope constructs.Construct, id *string, config *DataAwsCloudfrontOriginRequestPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicy",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsCloudfrontOriginRequestPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCloudfrontOriginRequestPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) CookiesConfig(index *string) DataAwsCloudfrontOriginRequestPolicyCookiesConfig {
	var returns DataAwsCloudfrontOriginRequestPolicyCookiesConfig

	_jsii_.Invoke(
		d,
		"cookiesConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) HeadersConfig(index *string) DataAwsCloudfrontOriginRequestPolicyHeadersConfig {
	var returns DataAwsCloudfrontOriginRequestPolicyHeadersConfig

	_jsii_.Invoke(
		d,
		"headersConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) QueryStringsConfig(index *string) DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig {
	var returns DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig

	_jsii_.Invoke(
		d,
		"queryStringsConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) ToString() *string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCloudfrontOriginRequestPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_origin_request_policy.html#id DataAwsCloudfrontOriginRequestPolicy#id}.
	Id *string `json:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_origin_request_policy.html#name DataAwsCloudfrontOriginRequestPolicy#name}.
	Name *string `json:"name"`
}

type DataAwsCloudfrontOriginRequestPolicyCookiesConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	CookieBehavior() *string
	Cookies() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontOriginRequestPolicyCookiesConfig
type jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) CookieBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) Cookies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyCookiesConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontOriginRequestPolicyCookiesConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyCookiesConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyCookiesConfig_Override(d DataAwsCloudfrontOriginRequestPolicyCookiesConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyCookiesConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies
type jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies_Override(d DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontOriginRequestPolicyHeadersConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	HeaderBehavior() *string
	Headers() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontOriginRequestPolicyHeadersConfig
type jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) HeaderBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) Headers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyHeadersConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontOriginRequestPolicyHeadersConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyHeadersConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyHeadersConfig_Override(d DataAwsCloudfrontOriginRequestPolicyHeadersConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyHeadersConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders
type jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders_Override(d DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	QueryStringBehavior() *string
	QueryStrings() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig
type jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) QueryStringBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) QueryStrings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyQueryStringsConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyQueryStringsConfig_Override(d DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings
type jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings_Override(d DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_response_headers_policy.html aws_cloudfront_response_headers_policy}.
type DataAwsCloudfrontResponseHeadersPolicy interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	Comment() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Etag() *string
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
	AddOverride(path *string, value interface{})
	CorsConfig(index *string) DataAwsCloudfrontResponseHeadersPolicyCorsConfig
	CustomHeadersConfig(index *string) DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetName()
	ResetOverrideLogicalId()
	SecurityHeadersConfig(index *string) DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicy
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_response_headers_policy.html aws_cloudfront_response_headers_policy} Data Source.
func NewDataAwsCloudfrontResponseHeadersPolicy(scope constructs.Construct, id *string, config *DataAwsCloudfrontResponseHeadersPolicyConfig) DataAwsCloudfrontResponseHeadersPolicy {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_response_headers_policy.html aws_cloudfront_response_headers_policy} Data Source.
func NewDataAwsCloudfrontResponseHeadersPolicy_Override(d DataAwsCloudfrontResponseHeadersPolicy, scope constructs.Construct, id *string, config *DataAwsCloudfrontResponseHeadersPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicy",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsCloudfrontResponseHeadersPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCloudfrontResponseHeadersPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) CorsConfig(index *string) DataAwsCloudfrontResponseHeadersPolicyCorsConfig {
	var returns DataAwsCloudfrontResponseHeadersPolicyCorsConfig

	_jsii_.Invoke(
		d,
		"corsConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) CustomHeadersConfig(index *string) DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig {
	var returns DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig

	_jsii_.Invoke(
		d,
		"customHeadersConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) SecurityHeadersConfig(index *string) DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig {
	var returns DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig

	_jsii_.Invoke(
		d,
		"securityHeadersConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) ToString() *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudfront_response_headers_policy.html#name DataAwsCloudfrontResponseHeadersPolicy#name}.
	Name *string `json:"name"`
}

type DataAwsCloudfrontResponseHeadersPolicyCorsConfig interface {
	cdktf.ComplexComputedList
	AccessControlAllowCredentials() interface{}
	AccessControlAllowHeaders() interface{}
	AccessControlAllowMethods() interface{}
	AccessControlAllowOrigins() interface{}
	AccessControlExposeHeaders() interface{}
	AccessControlMaxAgeSec() *float64
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	OriginOverride() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicyCorsConfig
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) AccessControlAllowCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) AccessControlAllowHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) AccessControlAllowMethods() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) AccessControlAllowOrigins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlAllowOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) AccessControlExposeHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlExposeHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) OriginOverride() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"originOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicyCorsConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfig_Override(d DataAwsCloudfrontResponseHeadersPolicyCorsConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders_Override(d DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods_Override(d DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins_Override(d DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Items() *[]*string
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) Items() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders_Override(d DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Header() *string
	Override() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Value() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) Header() *string {
	var returns *string
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig_Override(d DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ContentSecurityPolicy() interface{}
	ContentTypeOptions() interface{}
	FrameOptions() interface{}
	ReferrerPolicy() interface{}
	StrictTransportSecurity() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	XssProtection() interface{}
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) ContentSecurityPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) ContentTypeOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) FrameOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"frameOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) ReferrerPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"referrerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) StrictTransportSecurity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictTransportSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) XssProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xssProtection",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig_Override(d DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ContentSecurityPolicy() *string
	Override() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) ContentSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy_Override(d DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Override() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions_Override(d DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	FrameOption() *string
	Override() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) FrameOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frameOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions_Override(d DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Override() interface{}
	ReferrerPolicy() *string
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) ReferrerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referrerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy_Override(d DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity interface {
	cdktf.ComplexComputedList
	AccessControlMaxAgeSec() *float64
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	IncludeSubdomains() interface{}
	Override() interface{}
	Preload() interface{}
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) IncludeSubdomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) Preload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity_Override(d DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ModeBlock() interface{}
	Override() interface{}
	Protection() interface{}
	ReportUri() *string
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

// The jsii proxy struct for DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection
type jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) ModeBlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modeBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) Override() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) Protection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) ReportUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection{}

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection_Override(d DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}
