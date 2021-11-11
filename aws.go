// hashicorp_aws
package aws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws aws}.
type AwsProvider interface {
	cdktf.TerraformProvider
	AccessKey() *string
	SetAccessKey(val *string)
	AccessKeyInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AllowedAccountIds() *[]*string
	SetAllowedAccountIds(val *[]*string)
	AllowedAccountIdsInput() *[]*string
	AssumeRole() *AwsProviderAssumeRole
	SetAssumeRole(val *AwsProviderAssumeRole)
	AssumeRoleInput() *AwsProviderAssumeRole
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	DefaultTags() *AwsProviderDefaultTags
	SetDefaultTags(val *AwsProviderDefaultTags)
	DefaultTagsInput() *AwsProviderDefaultTags
	Endpoints() *[]*AwsProviderEndpoints
	SetEndpoints(val *[]*AwsProviderEndpoints)
	EndpointsInput() *[]*AwsProviderEndpoints
	ForbiddenAccountIds() *[]*string
	SetForbiddenAccountIds(val *[]*string)
	ForbiddenAccountIdsInput() *[]*string
	Fqn() *string
	FriendlyUniqueId() *string
	HttpProxy() *string
	SetHttpProxy(val *string)
	HttpProxyInput() *string
	IgnoreTags() *AwsProviderIgnoreTags
	SetIgnoreTags(val *AwsProviderIgnoreTags)
	IgnoreTagsInput() *AwsProviderIgnoreTags
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
	MetaAttributes() *map[string]interface{}
	Node() constructs.Node
	Profile() *string
	SetProfile(val *string)
	ProfileInput() *string
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	S3ForcePathStyle() interface{}
	SetS3ForcePathStyle(val interface{})
	S3ForcePathStyleInput() interface{}
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	SharedCredentialsFile() *string
	SetSharedCredentialsFile(val *string)
	SharedCredentialsFileInput() *string
	SkipCredentialsValidation() interface{}
	SetSkipCredentialsValidation(val interface{})
	SkipCredentialsValidationInput() interface{}
	SkipGetEc2Platforms() interface{}
	SetSkipGetEc2Platforms(val interface{})
	SkipGetEc2PlatformsInput() interface{}
	SkipMetadataApiCheck() interface{}
	SetSkipMetadataApiCheck(val interface{})
	SkipMetadataApiCheckInput() interface{}
	SkipRegionValidation() interface{}
	SetSkipRegionValidation(val interface{})
	SkipRegionValidationInput() interface{}
	SkipRequestingAccountId() interface{}
	SetSkipRequestingAccountId(val interface{})
	SkipRequestingAccountIdInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformProviderSource() *string
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	AddOverride(path *string, value interface{})
	OverrideLogicalId(newLogicalId *string)
	ResetAccessKey()
	ResetAlias()
	ResetAllowedAccountIds()
	ResetAssumeRole()
	ResetDefaultTags()
	ResetEndpoints()
	ResetForbiddenAccountIds()
	ResetHttpProxy()
	ResetIgnoreTags()
	ResetInsecure()
	ResetMaxRetries()
	ResetOverrideLogicalId()
	ResetProfile()
	ResetS3ForcePathStyle()
	ResetSecretKey()
	ResetSharedCredentialsFile()
	ResetSkipCredentialsValidation()
	ResetSkipGetEc2Platforms()
	ResetSkipMetadataApiCheck()
	ResetSkipRegionValidation()
	ResetSkipRequestingAccountId()
	ResetToken()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AwsProvider
type jsiiProxy_AwsProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_AwsProvider) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AllowedAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AllowedAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRole() *AwsProviderAssumeRole {
	var returns *AwsProviderAssumeRole
	_jsii_.Get(
		j,
		"assumeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRoleInput() *AwsProviderAssumeRole {
	var returns *AwsProviderAssumeRole
	_jsii_.Get(
		j,
		"assumeRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) DefaultTags() *AwsProviderDefaultTags {
	var returns *AwsProviderDefaultTags
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) DefaultTagsInput() *AwsProviderDefaultTags {
	var returns *AwsProviderDefaultTags
	_jsii_.Get(
		j,
		"defaultTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Endpoints() *[]*AwsProviderEndpoints {
	var returns *[]*AwsProviderEndpoints
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) EndpointsInput() *[]*AwsProviderEndpoints {
	var returns *[]*AwsProviderEndpoints
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ForbiddenAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ForbiddenAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) HttpProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) HttpProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) IgnoreTags() *AwsProviderIgnoreTags {
	var returns *AwsProviderIgnoreTags
	_jsii_.Get(
		j,
		"ignoreTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) IgnoreTagsInput() *AwsProviderIgnoreTags {
	var returns *AwsProviderIgnoreTags
	_jsii_.Get(
		j,
		"ignoreTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3ForcePathStyle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ForcePathStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3ForcePathStyleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ForcePathStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedCredentialsFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedCredentialsFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedCredentialsFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedCredentialsFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipCredentialsValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCredentialsValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipCredentialsValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCredentialsValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipGetEc2Platforms() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGetEc2Platforms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipGetEc2PlatformsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipGetEc2PlatformsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipMetadataApiCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMetadataApiCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipMetadataApiCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipMetadataApiCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRegionValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRegionValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRegionValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRegionValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRequestingAccountId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRequestingAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRequestingAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRequestingAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws aws} Resource.
func NewAwsProvider(scope constructs.Construct, id *string, config *AwsProviderConfig) AwsProvider {
	_init_.Initialize()

	j := jsiiProxy_AwsProvider{}

	_jsii_.Create(
		"hashicorp_aws.AwsProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws aws} Resource.
func NewAwsProvider_Override(a AwsProvider, scope constructs.Construct, id *string, config *AwsProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AwsProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsProvider) SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetAllowedAccountIds(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedAccountIds",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetAssumeRole(val *AwsProviderAssumeRole) {
	_jsii_.Set(
		j,
		"assumeRole",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetDefaultTags(val *AwsProviderDefaultTags) {
	_jsii_.Set(
		j,
		"defaultTags",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetEndpoints(val *[]*AwsProviderEndpoints) {
	_jsii_.Set(
		j,
		"endpoints",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetForbiddenAccountIds(val *[]*string) {
	_jsii_.Set(
		j,
		"forbiddenAccountIds",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetHttpProxy(val *string) {
	_jsii_.Set(
		j,
		"httpProxy",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetIgnoreTags(val *AwsProviderIgnoreTags) {
	_jsii_.Set(
		j,
		"ignoreTags",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetInsecure(val interface{}) {
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetProfile(val *string) {
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetS3ForcePathStyle(val interface{}) {
	_jsii_.Set(
		j,
		"s3ForcePathStyle",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetSharedCredentialsFile(val *string) {
	_jsii_.Set(
		j,
		"sharedCredentialsFile",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetSkipCredentialsValidation(val interface{}) {
	_jsii_.Set(
		j,
		"skipCredentialsValidation",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetSkipGetEc2Platforms(val interface{}) {
	_jsii_.Set(
		j,
		"skipGetEc2Platforms",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetSkipMetadataApiCheck(val interface{}) {
	_jsii_.Set(
		j,
		"skipMetadataApiCheck",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetSkipRegionValidation(val interface{}) {
	_jsii_.Set(
		j,
		"skipRegionValidation",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetSkipRequestingAccountId(val interface{}) {
	_jsii_.Set(
		j,
		"skipRequestingAccountId",
		val,
	)
}

func (j *jsiiProxy_AwsProvider) SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AwsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.AwsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.AwsProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AwsProvider) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (a *jsiiProxy_AwsProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsProvider) ResetAccessKey() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAllowedAccountIds() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedAccountIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAssumeRole() {
	_jsii_.InvokeVoid(
		a,
		"resetAssumeRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetDefaultTags() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetForbiddenAccountIds() {
	_jsii_.InvokeVoid(
		a,
		"resetForbiddenAccountIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetHttpProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetIgnoreTags() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxRetries",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AwsProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetS3ForcePathStyle() {
	_jsii_.InvokeVoid(
		a,
		"resetS3ForcePathStyle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSecretKey() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSharedCredentialsFile() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedCredentialsFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipCredentialsValidation() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipCredentialsValidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipGetEc2Platforms() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipGetEc2Platforms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipMetadataApiCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipMetadataApiCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipRegionValidation() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipRegionValidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipRequestingAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipRequestingAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetToken() {
	_jsii_.InvokeVoid(
		a,
		"resetToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AwsProvider) ToMetadata() interface{} {
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
func (a *jsiiProxy_AwsProvider) ToString() *string {
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
func (a *jsiiProxy_AwsProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AwsProviderAssumeRole struct {
	// Seconds to restrict the assume role session duration.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#duration_seconds AwsProvider#duration_seconds}
	DurationSeconds *float64 `json:"durationSeconds"`
	// Unique identifier that might be required for assuming a role in another account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#external_id AwsProvider#external_id}
	ExternalId *string `json:"externalId"`
	// IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#policy AwsProvider#policy}
	Policy *string `json:"policy"`
	// Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#policy_arns AwsProvider#policy_arns}
	PolicyArns *[]*string `json:"policyArns"`
	// Amazon Resource Name of an IAM Role to assume prior to making API calls.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#role_arn AwsProvider#role_arn}
	RoleArn *string `json:"roleArn"`
	// Identifier for the assumed role session.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#session_name AwsProvider#session_name}
	SessionName *string `json:"sessionName"`
	// Assume role session tags.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#tags AwsProvider#tags}
	Tags interface{} `json:"tags"`
	// Assume role session tag keys to pass to any subsequent sessions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#transitive_tag_keys AwsProvider#transitive_tag_keys}
	TransitiveTagKeys *[]*string `json:"transitiveTagKeys"`
}

type AwsProviderAssumeRoleOutputReference interface {
	cdktf.ComplexObject
	DurationSeconds() *float64
	SetDurationSeconds(val *float64)
	DurationSecondsInput() *float64
	ExternalId() *string
	SetExternalId(val *string)
	ExternalIdInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Policy() *string
	SetPolicy(val *string)
	PolicyArns() *[]*string
	SetPolicyArns(val *[]*string)
	PolicyArnsInput() *[]*string
	PolicyInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	SessionName() *string
	SetSessionName(val *string)
	SessionNameInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsInput() interface{}
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TransitiveTagKeys() *[]*string
	SetTransitiveTagKeys(val *[]*string)
	TransitiveTagKeysInput() *[]*string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetDurationSeconds()
	ResetExternalId()
	ResetPolicy()
	ResetPolicyArns()
	ResetRoleArn()
	ResetSessionName()
	ResetTags()
	ResetTransitiveTagKeys()
}

// The jsii proxy struct for AwsProviderAssumeRoleOutputReference
type jsiiProxy_AwsProviderAssumeRoleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) DurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) DurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) ExternalIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) PolicyArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) PolicyArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SessionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SessionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) TransitiveTagKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transitiveTagKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) TransitiveTagKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"transitiveTagKeysInput",
		&returns,
	)
	return returns
}

func NewAwsProviderAssumeRoleOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AwsProviderAssumeRoleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AwsProviderAssumeRoleOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AwsProviderAssumeRoleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAwsProviderAssumeRoleOutputReference_Override(a AwsProviderAssumeRoleOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AwsProviderAssumeRoleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"durationSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetExternalId(val *string) {
	_jsii_.Set(
		j,
		"externalId",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetPolicyArns(val *[]*string) {
	_jsii_.Set(
		j,
		"policyArns",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetSessionName(val *string) {
	_jsii_.Set(
		j,
		"sessionName",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsProviderAssumeRoleOutputReference) SetTransitiveTagKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"transitiveTagKeys",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) ResetDurationSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetDurationSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) ResetExternalId() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) ResetPolicyArns() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) ResetSessionName() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProviderAssumeRoleOutputReference) ResetTransitiveTagKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitiveTagKeys",
		nil, // no parameters
	)
}

type AwsProviderConfig struct {
	// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#region AwsProvider#region}
	Region *string `json:"region"`
	// The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#access_key AwsProvider#access_key}
	AccessKey *string `json:"accessKey"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#alias AwsProvider#alias}
	Alias *string `json:"alias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#allowed_account_ids AwsProvider#allowed_account_ids}.
	AllowedAccountIds *[]*string `json:"allowedAccountIds"`
	// assume_role block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#assume_role AwsProvider#assume_role}
	AssumeRole *AwsProviderAssumeRole `json:"assumeRole"`
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#default_tags AwsProvider#default_tags}
	DefaultTags *AwsProviderDefaultTags `json:"defaultTags"`
	// endpoints block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#endpoints AwsProvider#endpoints}
	Endpoints *[]*AwsProviderEndpoints `json:"endpoints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#forbidden_account_ids AwsProvider#forbidden_account_ids}.
	ForbiddenAccountIds *[]*string `json:"forbiddenAccountIds"`
	// The address of an HTTP proxy to use when accessing the AWS API.
	//
	// Can also be configured using the `HTTP_PROXY` or `HTTPS_PROXY` environment variables.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#http_proxy AwsProvider#http_proxy}
	HttpProxy *string `json:"httpProxy"`
	// ignore_tags block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ignore_tags AwsProvider#ignore_tags}
	IgnoreTags *AwsProviderIgnoreTags `json:"ignoreTags"`
	// Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is `false`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#insecure AwsProvider#insecure}
	Insecure interface{} `json:"insecure"`
	// The maximum number of times an AWS API request is being executed.
	//
	// If the API request still fails, an error is
	// thrown.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#max_retries AwsProvider#max_retries}
	MaxRetries *float64 `json:"maxRetries"`
	// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#profile AwsProvider#profile}
	Profile *string `json:"profile"`
	// Set this to true to force the request to use path-style addressing, i.e., http://s3.amazonaws.com/BUCKET/KEY. By default, the S3 client will use virtual hosted bucket addressing when possible (http://BUCKET.s3.amazonaws.com/KEY). Specific to the Amazon S3 service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#s3_force_path_style AwsProvider#s3_force_path_style}
	S3ForcePathStyle interface{} `json:"s3ForcePathStyle"`
	// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#secret_key AwsProvider#secret_key}
	SecretKey *string `json:"secretKey"`
	// The path to the shared credentials file. If not set this defaults to ~/.aws/credentials.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#shared_credentials_file AwsProvider#shared_credentials_file}
	SharedCredentialsFile *string `json:"sharedCredentialsFile"`
	// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS available/implemented.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_credentials_validation AwsProvider#skip_credentials_validation}
	SkipCredentialsValidation interface{} `json:"skipCredentialsValidation"`
	// Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_get_ec2_platforms AwsProvider#skip_get_ec2_platforms}
	SkipGetEc2Platforms interface{} `json:"skipGetEc2Platforms"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_metadata_api_check AwsProvider#skip_metadata_api_check}.
	SkipMetadataApiCheck interface{} `json:"skipMetadataApiCheck"`
	// Skip static validation of region name.
	//
	// Used by users of alternative AWS-like APIs or users w/ access to regions that are not public (yet).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_region_validation AwsProvider#skip_region_validation}
	SkipRegionValidation interface{} `json:"skipRegionValidation"`
	// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#skip_requesting_account_id AwsProvider#skip_requesting_account_id}
	SkipRequestingAccountId interface{} `json:"skipRequestingAccountId"`
	// session token. A session token is only required if you are using temporary security credentials.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#token AwsProvider#token}
	Token *string `json:"token"`
}

type AwsProviderDefaultTags struct {
	// Resource tags to default across all resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#tags AwsProvider#tags}
	Tags interface{} `json:"tags"`
}

type AwsProviderDefaultTagsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Tags() interface{}
	SetTags(val interface{})
	TagsInput() interface{}
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
	ResetTags()
}

// The jsii proxy struct for AwsProviderDefaultTagsOutputReference
type jsiiProxy_AwsProviderDefaultTagsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AwsProviderDefaultTagsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderDefaultTagsOutputReference) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderDefaultTagsOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderDefaultTagsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderDefaultTagsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAwsProviderDefaultTagsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AwsProviderDefaultTagsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AwsProviderDefaultTagsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AwsProviderDefaultTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAwsProviderDefaultTagsOutputReference_Override(a AwsProviderDefaultTagsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AwsProviderDefaultTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AwsProviderDefaultTagsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AwsProviderDefaultTagsOutputReference) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsProviderDefaultTagsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsProviderDefaultTagsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AwsProviderDefaultTagsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AwsProviderDefaultTagsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AwsProviderDefaultTagsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AwsProviderDefaultTagsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AwsProviderDefaultTagsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AwsProviderDefaultTagsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProviderDefaultTagsOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

type AwsProviderEndpoints struct {
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#accessanalyzer AwsProvider#accessanalyzer}
	Accessanalyzer *string `json:"accessanalyzer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#acm AwsProvider#acm}
	Acm *string `json:"acm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#acmpca AwsProvider#acmpca}
	Acmpca *string `json:"acmpca"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#alexaforbusiness AwsProvider#alexaforbusiness}
	Alexaforbusiness *string `json:"alexaforbusiness"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#amplify AwsProvider#amplify}
	Amplify *string `json:"amplify"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#amplifybackend AwsProvider#amplifybackend}
	Amplifybackend *string `json:"amplifybackend"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#apigateway AwsProvider#apigateway}
	Apigateway *string `json:"apigateway"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#apigatewayv2 AwsProvider#apigatewayv2}
	Apigatewayv2 *string `json:"apigatewayv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#appautoscaling AwsProvider#appautoscaling}
	Appautoscaling *string `json:"appautoscaling"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#appconfig AwsProvider#appconfig}
	Appconfig *string `json:"appconfig"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#appflow AwsProvider#appflow}
	Appflow *string `json:"appflow"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#appintegrations AwsProvider#appintegrations}
	Appintegrations *string `json:"appintegrations"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#appintegrationsservice AwsProvider#appintegrationsservice}
	Appintegrationsservice *string `json:"appintegrationsservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#applicationautoscaling AwsProvider#applicationautoscaling}
	Applicationautoscaling *string `json:"applicationautoscaling"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#applicationcostprofiler AwsProvider#applicationcostprofiler}
	Applicationcostprofiler *string `json:"applicationcostprofiler"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#applicationdiscovery AwsProvider#applicationdiscovery}
	Applicationdiscovery *string `json:"applicationdiscovery"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#applicationdiscoveryservice AwsProvider#applicationdiscoveryservice}
	Applicationdiscoveryservice *string `json:"applicationdiscoveryservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#applicationinsights AwsProvider#applicationinsights}
	Applicationinsights *string `json:"applicationinsights"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#appmesh AwsProvider#appmesh}
	Appmesh *string `json:"appmesh"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#appregistry AwsProvider#appregistry}
	Appregistry *string `json:"appregistry"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#apprunner AwsProvider#apprunner}
	Apprunner *string `json:"apprunner"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#appstream AwsProvider#appstream}
	Appstream *string `json:"appstream"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#appsync AwsProvider#appsync}
	Appsync *string `json:"appsync"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#athena AwsProvider#athena}
	Athena *string `json:"athena"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#auditmanager AwsProvider#auditmanager}
	Auditmanager *string `json:"auditmanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#augmentedairuntime AwsProvider#augmentedairuntime}
	Augmentedairuntime *string `json:"augmentedairuntime"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#autoscaling AwsProvider#autoscaling}
	Autoscaling *string `json:"autoscaling"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#autoscalingplans AwsProvider#autoscalingplans}
	Autoscalingplans *string `json:"autoscalingplans"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#backup AwsProvider#backup}
	Backup *string `json:"backup"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#batch AwsProvider#batch}
	Batch *string `json:"batch"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#braket AwsProvider#braket}
	Braket *string `json:"braket"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#budgets AwsProvider#budgets}
	Budgets *string `json:"budgets"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#chime AwsProvider#chime}
	Chime *string `json:"chime"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloud9 AwsProvider#cloud9}
	Cloud9 *string `json:"cloud9"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudcontrol AwsProvider#cloudcontrol}
	Cloudcontrol *string `json:"cloudcontrol"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudcontrolapi AwsProvider#cloudcontrolapi}
	Cloudcontrolapi *string `json:"cloudcontrolapi"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#clouddirectory AwsProvider#clouddirectory}
	Clouddirectory *string `json:"clouddirectory"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudformation AwsProvider#cloudformation}
	Cloudformation *string `json:"cloudformation"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudfront AwsProvider#cloudfront}
	Cloudfront *string `json:"cloudfront"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudhsm AwsProvider#cloudhsm}
	Cloudhsm *string `json:"cloudhsm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudhsmv2 AwsProvider#cloudhsmv2}
	Cloudhsmv2 *string `json:"cloudhsmv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudsearch AwsProvider#cloudsearch}
	Cloudsearch *string `json:"cloudsearch"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudsearchdomain AwsProvider#cloudsearchdomain}
	Cloudsearchdomain *string `json:"cloudsearchdomain"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudtrail AwsProvider#cloudtrail}
	Cloudtrail *string `json:"cloudtrail"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudwatch AwsProvider#cloudwatch}
	Cloudwatch *string `json:"cloudwatch"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudwatchevents AwsProvider#cloudwatchevents}
	Cloudwatchevents *string `json:"cloudwatchevents"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cloudwatchlogs AwsProvider#cloudwatchlogs}
	Cloudwatchlogs *string `json:"cloudwatchlogs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codeartifact AwsProvider#codeartifact}
	Codeartifact *string `json:"codeartifact"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codebuild AwsProvider#codebuild}
	Codebuild *string `json:"codebuild"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codecommit AwsProvider#codecommit}
	Codecommit *string `json:"codecommit"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codedeploy AwsProvider#codedeploy}
	Codedeploy *string `json:"codedeploy"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codeguruprofiler AwsProvider#codeguruprofiler}
	Codeguruprofiler *string `json:"codeguruprofiler"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codegurureviewer AwsProvider#codegurureviewer}
	Codegurureviewer *string `json:"codegurureviewer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codepipeline AwsProvider#codepipeline}
	Codepipeline *string `json:"codepipeline"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codestar AwsProvider#codestar}
	Codestar *string `json:"codestar"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codestarconnections AwsProvider#codestarconnections}
	Codestarconnections *string `json:"codestarconnections"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#codestarnotifications AwsProvider#codestarnotifications}
	Codestarnotifications *string `json:"codestarnotifications"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cognitoidentity AwsProvider#cognitoidentity}
	Cognitoidentity *string `json:"cognitoidentity"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cognitoidentityprovider AwsProvider#cognitoidentityprovider}
	Cognitoidentityprovider *string `json:"cognitoidentityprovider"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cognitoidp AwsProvider#cognitoidp}
	Cognitoidp *string `json:"cognitoidp"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cognitosync AwsProvider#cognitosync}
	Cognitosync *string `json:"cognitosync"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#comprehend AwsProvider#comprehend}
	Comprehend *string `json:"comprehend"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#comprehendmedical AwsProvider#comprehendmedical}
	Comprehendmedical *string `json:"comprehendmedical"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#config AwsProvider#config}
	Config *string `json:"config"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#configservice AwsProvider#configservice}
	Configservice *string `json:"configservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#connect AwsProvider#connect}
	Connect *string `json:"connect"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#connectcontactlens AwsProvider#connectcontactlens}
	Connectcontactlens *string `json:"connectcontactlens"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#connectparticipant AwsProvider#connectparticipant}
	Connectparticipant *string `json:"connectparticipant"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#costandusagereportservice AwsProvider#costandusagereportservice}
	Costandusagereportservice *string `json:"costandusagereportservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#costexplorer AwsProvider#costexplorer}
	Costexplorer *string `json:"costexplorer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#cur AwsProvider#cur}
	Cur *string `json:"cur"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#databasemigration AwsProvider#databasemigration}
	Databasemigration *string `json:"databasemigration"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#databasemigrationservice AwsProvider#databasemigrationservice}
	Databasemigrationservice *string `json:"databasemigrationservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#dataexchange AwsProvider#dataexchange}
	Dataexchange *string `json:"dataexchange"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#datapipeline AwsProvider#datapipeline}
	Datapipeline *string `json:"datapipeline"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#datasync AwsProvider#datasync}
	Datasync *string `json:"datasync"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#dax AwsProvider#dax}
	Dax *string `json:"dax"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#detective AwsProvider#detective}
	Detective *string `json:"detective"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#devicefarm AwsProvider#devicefarm}
	Devicefarm *string `json:"devicefarm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#devopsguru AwsProvider#devopsguru}
	Devopsguru *string `json:"devopsguru"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#directconnect AwsProvider#directconnect}
	Directconnect *string `json:"directconnect"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#dlm AwsProvider#dlm}
	Dlm *string `json:"dlm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#dms AwsProvider#dms}
	Dms *string `json:"dms"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#docdb AwsProvider#docdb}
	Docdb *string `json:"docdb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ds AwsProvider#ds}
	Ds *string `json:"ds"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#dynamodb AwsProvider#dynamodb}
	Dynamodb *string `json:"dynamodb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#dynamodbstreams AwsProvider#dynamodbstreams}
	Dynamodbstreams *string `json:"dynamodbstreams"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ec2 AwsProvider#ec2}
	Ec2 *string `json:"ec2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ec2instanceconnect AwsProvider#ec2instanceconnect}
	Ec2Instanceconnect *string `json:"ec2Instanceconnect"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ecr AwsProvider#ecr}
	Ecr *string `json:"ecr"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ecrpublic AwsProvider#ecrpublic}
	Ecrpublic *string `json:"ecrpublic"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ecs AwsProvider#ecs}
	Ecs *string `json:"ecs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#efs AwsProvider#efs}
	Efs *string `json:"efs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#eks AwsProvider#eks}
	Eks *string `json:"eks"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#elasticache AwsProvider#elasticache}
	Elasticache *string `json:"elasticache"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#elasticbeanstalk AwsProvider#elasticbeanstalk}
	Elasticbeanstalk *string `json:"elasticbeanstalk"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#elasticinference AwsProvider#elasticinference}
	Elasticinference *string `json:"elasticinference"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#elasticsearch AwsProvider#elasticsearch}
	Elasticsearch *string `json:"elasticsearch"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#elasticsearchservice AwsProvider#elasticsearchservice}
	Elasticsearchservice *string `json:"elasticsearchservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#elastictranscoder AwsProvider#elastictranscoder}
	Elastictranscoder *string `json:"elastictranscoder"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#elb AwsProvider#elb}
	Elb *string `json:"elb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#elbv2 AwsProvider#elbv2}
	Elbv2 *string `json:"elbv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#emr AwsProvider#emr}
	Emr *string `json:"emr"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#emrcontainers AwsProvider#emrcontainers}
	Emrcontainers *string `json:"emrcontainers"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#es AwsProvider#es}
	Es *string `json:"es"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#finspace AwsProvider#finspace}
	Finspace *string `json:"finspace"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#finspacedata AwsProvider#finspacedata}
	Finspacedata *string `json:"finspacedata"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#firehose AwsProvider#firehose}
	Firehose *string `json:"firehose"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#fis AwsProvider#fis}
	Fis *string `json:"fis"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#fms AwsProvider#fms}
	Fms *string `json:"fms"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#forecast AwsProvider#forecast}
	Forecast *string `json:"forecast"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#forecastquery AwsProvider#forecastquery}
	Forecastquery *string `json:"forecastquery"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#forecastqueryservice AwsProvider#forecastqueryservice}
	Forecastqueryservice *string `json:"forecastqueryservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#forecastservice AwsProvider#forecastservice}
	Forecastservice *string `json:"forecastservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#frauddetector AwsProvider#frauddetector}
	Frauddetector *string `json:"frauddetector"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#fsx AwsProvider#fsx}
	Fsx *string `json:"fsx"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#gamelift AwsProvider#gamelift}
	Gamelift *string `json:"gamelift"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#glacier AwsProvider#glacier}
	Glacier *string `json:"glacier"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#globalaccelerator AwsProvider#globalaccelerator}
	Globalaccelerator *string `json:"globalaccelerator"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#glue AwsProvider#glue}
	Glue *string `json:"glue"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#gluedatabrew AwsProvider#gluedatabrew}
	Gluedatabrew *string `json:"gluedatabrew"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#greengrass AwsProvider#greengrass}
	Greengrass *string `json:"greengrass"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#greengrassv2 AwsProvider#greengrassv2}
	Greengrassv2 *string `json:"greengrassv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#groundstation AwsProvider#groundstation}
	Groundstation *string `json:"groundstation"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#guardduty AwsProvider#guardduty}
	Guardduty *string `json:"guardduty"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#health AwsProvider#health}
	Health *string `json:"health"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#healthlake AwsProvider#healthlake}
	Healthlake *string `json:"healthlake"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#honeycode AwsProvider#honeycode}
	Honeycode *string `json:"honeycode"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iam AwsProvider#iam}
	Iam *string `json:"iam"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#identitystore AwsProvider#identitystore}
	Identitystore *string `json:"identitystore"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#imagebuilder AwsProvider#imagebuilder}
	Imagebuilder *string `json:"imagebuilder"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#inspector AwsProvider#inspector}
	Inspector *string `json:"inspector"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iot AwsProvider#iot}
	Iot *string `json:"iot"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iot1clickdevices AwsProvider#iot1clickdevices}
	Iot1Clickdevices *string `json:"iot1Clickdevices"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iot1clickdevicesservice AwsProvider#iot1clickdevicesservice}
	Iot1Clickdevicesservice *string `json:"iot1Clickdevicesservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iot1clickprojects AwsProvider#iot1clickprojects}
	Iot1Clickprojects *string `json:"iot1Clickprojects"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotanalytics AwsProvider#iotanalytics}
	Iotanalytics *string `json:"iotanalytics"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotdataplane AwsProvider#iotdataplane}
	Iotdataplane *string `json:"iotdataplane"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotdeviceadvisor AwsProvider#iotdeviceadvisor}
	Iotdeviceadvisor *string `json:"iotdeviceadvisor"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotevents AwsProvider#iotevents}
	Iotevents *string `json:"iotevents"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ioteventsdata AwsProvider#ioteventsdata}
	Ioteventsdata *string `json:"ioteventsdata"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotfleethub AwsProvider#iotfleethub}
	Iotfleethub *string `json:"iotfleethub"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotjobsdataplane AwsProvider#iotjobsdataplane}
	Iotjobsdataplane *string `json:"iotjobsdataplane"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotsecuretunneling AwsProvider#iotsecuretunneling}
	Iotsecuretunneling *string `json:"iotsecuretunneling"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotsitewise AwsProvider#iotsitewise}
	Iotsitewise *string `json:"iotsitewise"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotthingsgraph AwsProvider#iotthingsgraph}
	Iotthingsgraph *string `json:"iotthingsgraph"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#iotwireless AwsProvider#iotwireless}
	Iotwireless *string `json:"iotwireless"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kafka AwsProvider#kafka}
	Kafka *string `json:"kafka"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kendra AwsProvider#kendra}
	Kendra *string `json:"kendra"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kinesis AwsProvider#kinesis}
	Kinesis *string `json:"kinesis"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kinesisanalytics AwsProvider#kinesisanalytics}
	Kinesisanalytics *string `json:"kinesisanalytics"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kinesisanalyticsv2 AwsProvider#kinesisanalyticsv2}
	Kinesisanalyticsv2 *string `json:"kinesisanalyticsv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kinesisvideo AwsProvider#kinesisvideo}
	Kinesisvideo *string `json:"kinesisvideo"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kinesisvideoarchivedmedia AwsProvider#kinesisvideoarchivedmedia}
	Kinesisvideoarchivedmedia *string `json:"kinesisvideoarchivedmedia"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kinesisvideomedia AwsProvider#kinesisvideomedia}
	Kinesisvideomedia *string `json:"kinesisvideomedia"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kinesisvideosignalingchannels AwsProvider#kinesisvideosignalingchannels}
	Kinesisvideosignalingchannels *string `json:"kinesisvideosignalingchannels"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#kms AwsProvider#kms}
	Kms *string `json:"kms"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lakeformation AwsProvider#lakeformation}
	Lakeformation *string `json:"lakeformation"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lambda AwsProvider#lambda}
	Lambda *string `json:"lambda"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lexmodelbuilding AwsProvider#lexmodelbuilding}
	Lexmodelbuilding *string `json:"lexmodelbuilding"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lexmodelbuildingservice AwsProvider#lexmodelbuildingservice}
	Lexmodelbuildingservice *string `json:"lexmodelbuildingservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lexmodels AwsProvider#lexmodels}
	Lexmodels *string `json:"lexmodels"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lexmodelsv2 AwsProvider#lexmodelsv2}
	Lexmodelsv2 *string `json:"lexmodelsv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lexruntime AwsProvider#lexruntime}
	Lexruntime *string `json:"lexruntime"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lexruntimeservice AwsProvider#lexruntimeservice}
	Lexruntimeservice *string `json:"lexruntimeservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lexruntimev2 AwsProvider#lexruntimev2}
	Lexruntimev2 *string `json:"lexruntimev2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#licensemanager AwsProvider#licensemanager}
	Licensemanager *string `json:"licensemanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lightsail AwsProvider#lightsail}
	Lightsail *string `json:"lightsail"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#location AwsProvider#location}
	Location *string `json:"location"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lookoutequipment AwsProvider#lookoutequipment}
	Lookoutequipment *string `json:"lookoutequipment"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lookoutforvision AwsProvider#lookoutforvision}
	Lookoutforvision *string `json:"lookoutforvision"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#lookoutmetrics AwsProvider#lookoutmetrics}
	Lookoutmetrics *string `json:"lookoutmetrics"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#machinelearning AwsProvider#machinelearning}
	Machinelearning *string `json:"machinelearning"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#macie AwsProvider#macie}
	Macie *string `json:"macie"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#macie2 AwsProvider#macie2}
	Macie2 *string `json:"macie2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#managedblockchain AwsProvider#managedblockchain}
	Managedblockchain *string `json:"managedblockchain"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#marketplacecatalog AwsProvider#marketplacecatalog}
	Marketplacecatalog *string `json:"marketplacecatalog"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#marketplacecommerceanalytics AwsProvider#marketplacecommerceanalytics}
	Marketplacecommerceanalytics *string `json:"marketplacecommerceanalytics"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#marketplaceentitlement AwsProvider#marketplaceentitlement}
	Marketplaceentitlement *string `json:"marketplaceentitlement"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#marketplaceentitlementservice AwsProvider#marketplaceentitlementservice}
	Marketplaceentitlementservice *string `json:"marketplaceentitlementservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#marketplacemetering AwsProvider#marketplacemetering}
	Marketplacemetering *string `json:"marketplacemetering"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mediaconnect AwsProvider#mediaconnect}
	Mediaconnect *string `json:"mediaconnect"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mediaconvert AwsProvider#mediaconvert}
	Mediaconvert *string `json:"mediaconvert"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#medialive AwsProvider#medialive}
	Medialive *string `json:"medialive"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mediapackage AwsProvider#mediapackage}
	Mediapackage *string `json:"mediapackage"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mediapackagevod AwsProvider#mediapackagevod}
	Mediapackagevod *string `json:"mediapackagevod"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mediastore AwsProvider#mediastore}
	Mediastore *string `json:"mediastore"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mediastoredata AwsProvider#mediastoredata}
	Mediastoredata *string `json:"mediastoredata"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mediatailor AwsProvider#mediatailor}
	Mediatailor *string `json:"mediatailor"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#memorydb AwsProvider#memorydb}
	Memorydb *string `json:"memorydb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mgn AwsProvider#mgn}
	Mgn *string `json:"mgn"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#migrationhub AwsProvider#migrationhub}
	Migrationhub *string `json:"migrationhub"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#migrationhubconfig AwsProvider#migrationhubconfig}
	Migrationhubconfig *string `json:"migrationhubconfig"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mobile AwsProvider#mobile}
	Mobile *string `json:"mobile"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mobileanalytics AwsProvider#mobileanalytics}
	Mobileanalytics *string `json:"mobileanalytics"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mq AwsProvider#mq}
	Mq *string `json:"mq"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mturk AwsProvider#mturk}
	Mturk *string `json:"mturk"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#mwaa AwsProvider#mwaa}
	Mwaa *string `json:"mwaa"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#neptune AwsProvider#neptune}
	Neptune *string `json:"neptune"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#networkfirewall AwsProvider#networkfirewall}
	Networkfirewall *string `json:"networkfirewall"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#networkmanager AwsProvider#networkmanager}
	Networkmanager *string `json:"networkmanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#nimblestudio AwsProvider#nimblestudio}
	Nimblestudio *string `json:"nimblestudio"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#opsworks AwsProvider#opsworks}
	Opsworks *string `json:"opsworks"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#opsworkscm AwsProvider#opsworkscm}
	Opsworkscm *string `json:"opsworkscm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#organizations AwsProvider#organizations}
	Organizations *string `json:"organizations"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#outposts AwsProvider#outposts}
	Outposts *string `json:"outposts"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#personalize AwsProvider#personalize}
	Personalize *string `json:"personalize"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#personalizeevents AwsProvider#personalizeevents}
	Personalizeevents *string `json:"personalizeevents"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#personalizeruntime AwsProvider#personalizeruntime}
	Personalizeruntime *string `json:"personalizeruntime"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#pi AwsProvider#pi}
	Pi *string `json:"pi"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#pinpoint AwsProvider#pinpoint}
	Pinpoint *string `json:"pinpoint"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#pinpointemail AwsProvider#pinpointemail}
	Pinpointemail *string `json:"pinpointemail"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#pinpointsmsvoice AwsProvider#pinpointsmsvoice}
	Pinpointsmsvoice *string `json:"pinpointsmsvoice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#polly AwsProvider#polly}
	Polly *string `json:"polly"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#pricing AwsProvider#pricing}
	Pricing *string `json:"pricing"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#prometheus AwsProvider#prometheus}
	Prometheus *string `json:"prometheus"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#prometheusservice AwsProvider#prometheusservice}
	Prometheusservice *string `json:"prometheusservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#proton AwsProvider#proton}
	Proton *string `json:"proton"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#qldb AwsProvider#qldb}
	Qldb *string `json:"qldb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#qldbsession AwsProvider#qldbsession}
	Qldbsession *string `json:"qldbsession"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#quicksight AwsProvider#quicksight}
	Quicksight *string `json:"quicksight"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ram AwsProvider#ram}
	Ram *string `json:"ram"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#rds AwsProvider#rds}
	Rds *string `json:"rds"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#rdsdata AwsProvider#rdsdata}
	Rdsdata *string `json:"rdsdata"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#rdsdataservice AwsProvider#rdsdataservice}
	Rdsdataservice *string `json:"rdsdataservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#redshift AwsProvider#redshift}
	Redshift *string `json:"redshift"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#redshiftdata AwsProvider#redshiftdata}
	Redshiftdata *string `json:"redshiftdata"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#rekognition AwsProvider#rekognition}
	Rekognition *string `json:"rekognition"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#resourcegroups AwsProvider#resourcegroups}
	Resourcegroups *string `json:"resourcegroups"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#resourcegroupstagging AwsProvider#resourcegroupstagging}
	Resourcegroupstagging *string `json:"resourcegroupstagging"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#resourcegroupstaggingapi AwsProvider#resourcegroupstaggingapi}
	Resourcegroupstaggingapi *string `json:"resourcegroupstaggingapi"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#robomaker AwsProvider#robomaker}
	Robomaker *string `json:"robomaker"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#route53 AwsProvider#route53}
	Route53 *string `json:"route53"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#route53domains AwsProvider#route53domains}
	Route53Domains *string `json:"route53Domains"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#route53recoverycontrolconfig AwsProvider#route53recoverycontrolconfig}
	Route53Recoverycontrolconfig *string `json:"route53Recoverycontrolconfig"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#route53recoveryreadiness AwsProvider#route53recoveryreadiness}
	Route53Recoveryreadiness *string `json:"route53Recoveryreadiness"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#route53resolver AwsProvider#route53resolver}
	Route53Resolver *string `json:"route53Resolver"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#s3 AwsProvider#s3}
	S3 *string `json:"s3"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#s3control AwsProvider#s3control}
	S3Control *string `json:"s3Control"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#s3outposts AwsProvider#s3outposts}
	S3Outposts *string `json:"s3Outposts"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sagemaker AwsProvider#sagemaker}
	Sagemaker *string `json:"sagemaker"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sagemakeredgemanager AwsProvider#sagemakeredgemanager}
	Sagemakeredgemanager *string `json:"sagemakeredgemanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sagemakerfeaturestoreruntime AwsProvider#sagemakerfeaturestoreruntime}
	Sagemakerfeaturestoreruntime *string `json:"sagemakerfeaturestoreruntime"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sagemakerruntime AwsProvider#sagemakerruntime}
	Sagemakerruntime *string `json:"sagemakerruntime"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#savingsplans AwsProvider#savingsplans}
	Savingsplans *string `json:"savingsplans"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#schemas AwsProvider#schemas}
	Schemas *string `json:"schemas"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sdb AwsProvider#sdb}
	Sdb *string `json:"sdb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#secretsmanager AwsProvider#secretsmanager}
	Secretsmanager *string `json:"secretsmanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#securityhub AwsProvider#securityhub}
	Securityhub *string `json:"securityhub"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#serverlessapplicationrepository AwsProvider#serverlessapplicationrepository}
	Serverlessapplicationrepository *string `json:"serverlessapplicationrepository"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#serverlessapprepo AwsProvider#serverlessapprepo}
	Serverlessapprepo *string `json:"serverlessapprepo"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#serverlessrepo AwsProvider#serverlessrepo}
	Serverlessrepo *string `json:"serverlessrepo"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#servicecatalog AwsProvider#servicecatalog}
	Servicecatalog *string `json:"servicecatalog"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#servicediscovery AwsProvider#servicediscovery}
	Servicediscovery *string `json:"servicediscovery"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#servicequotas AwsProvider#servicequotas}
	Servicequotas *string `json:"servicequotas"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ses AwsProvider#ses}
	Ses *string `json:"ses"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sesv2 AwsProvider#sesv2}
	Sesv2 *string `json:"sesv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sfn AwsProvider#sfn}
	Sfn *string `json:"sfn"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#shield AwsProvider#shield}
	Shield *string `json:"shield"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#signer AwsProvider#signer}
	Signer *string `json:"signer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#simpledb AwsProvider#simpledb}
	Simpledb *string `json:"simpledb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sms AwsProvider#sms}
	Sms *string `json:"sms"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#snowball AwsProvider#snowball}
	Snowball *string `json:"snowball"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sns AwsProvider#sns}
	Sns *string `json:"sns"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sqs AwsProvider#sqs}
	Sqs *string `json:"sqs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ssm AwsProvider#ssm}
	Ssm *string `json:"ssm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ssmcontacts AwsProvider#ssmcontacts}
	Ssmcontacts *string `json:"ssmcontacts"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ssmincidents AwsProvider#ssmincidents}
	Ssmincidents *string `json:"ssmincidents"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sso AwsProvider#sso}
	Sso *string `json:"sso"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ssoadmin AwsProvider#ssoadmin}
	Ssoadmin *string `json:"ssoadmin"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#ssooidc AwsProvider#ssooidc}
	Ssooidc *string `json:"ssooidc"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#stepfunctions AwsProvider#stepfunctions}
	Stepfunctions *string `json:"stepfunctions"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#storagegateway AwsProvider#storagegateway}
	Storagegateway *string `json:"storagegateway"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#sts AwsProvider#sts}
	Sts *string `json:"sts"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#support AwsProvider#support}
	Support *string `json:"support"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#swf AwsProvider#swf}
	Swf *string `json:"swf"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#synthetics AwsProvider#synthetics}
	Synthetics *string `json:"synthetics"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#textract AwsProvider#textract}
	Textract *string `json:"textract"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#timestreamquery AwsProvider#timestreamquery}
	Timestreamquery *string `json:"timestreamquery"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#timestreamwrite AwsProvider#timestreamwrite}
	Timestreamwrite *string `json:"timestreamwrite"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#transcribe AwsProvider#transcribe}
	Transcribe *string `json:"transcribe"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#transcribeservice AwsProvider#transcribeservice}
	Transcribeservice *string `json:"transcribeservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#transcribestreaming AwsProvider#transcribestreaming}
	Transcribestreaming *string `json:"transcribestreaming"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#transcribestreamingservice AwsProvider#transcribestreamingservice}
	Transcribestreamingservice *string `json:"transcribestreamingservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#transfer AwsProvider#transfer}
	Transfer *string `json:"transfer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#translate AwsProvider#translate}
	Translate *string `json:"translate"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#waf AwsProvider#waf}
	Waf *string `json:"waf"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#wafregional AwsProvider#wafregional}
	Wafregional *string `json:"wafregional"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#wafv2 AwsProvider#wafv2}
	Wafv2 *string `json:"wafv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#wellarchitected AwsProvider#wellarchitected}
	Wellarchitected *string `json:"wellarchitected"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#workdocs AwsProvider#workdocs}
	Workdocs *string `json:"workdocs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#worklink AwsProvider#worklink}
	Worklink *string `json:"worklink"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#workmail AwsProvider#workmail}
	Workmail *string `json:"workmail"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#workmailmessageflow AwsProvider#workmailmessageflow}
	Workmailmessageflow *string `json:"workmailmessageflow"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#workspaces AwsProvider#workspaces}
	Workspaces *string `json:"workspaces"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#xray AwsProvider#xray}
	Xray *string `json:"xray"`
}

type AwsProviderIgnoreTags struct {
	// Resource tag key prefixes to ignore across all resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#key_prefixes AwsProvider#key_prefixes}
	KeyPrefixes *[]*string `json:"keyPrefixes"`
	// Resource tag keys to ignore across all resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#keys AwsProvider#keys}
	Keys *[]*string `json:"keys"`
}

type AwsProviderIgnoreTagsOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KeyPrefixes() *[]*string
	SetKeyPrefixes(val *[]*string)
	KeyPrefixesInput() *[]*string
	Keys() *[]*string
	SetKeys(val *[]*string)
	KeysInput() *[]*string
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
	ResetKeyPrefixes()
	ResetKeys()
}

// The jsii proxy struct for AwsProviderIgnoreTagsOutputReference
type jsiiProxy_AwsProviderIgnoreTagsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) KeyPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) KeyPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) Keys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) KeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAwsProviderIgnoreTagsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) AwsProviderIgnoreTagsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AwsProviderIgnoreTagsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.AwsProviderIgnoreTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAwsProviderIgnoreTagsOutputReference_Override(a AwsProviderIgnoreTagsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.AwsProviderIgnoreTagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) SetKeyPrefixes(val *[]*string) {
	_jsii_.Set(
		j,
		"keyPrefixes",
		val,
	)
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) SetKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"keys",
		val,
	)
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsProviderIgnoreTagsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AwsProviderIgnoreTagsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AwsProviderIgnoreTagsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AwsProviderIgnoreTagsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AwsProviderIgnoreTagsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AwsProviderIgnoreTagsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AwsProviderIgnoreTagsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProviderIgnoreTagsOutputReference) ResetKeyPrefixes() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyPrefixes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProviderIgnoreTagsOutputReference) ResetKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetKeys",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html aws_cloudcontrolapi_resource}.
type CloudcontrolapiResource interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Properties() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	Schema() *string
	SetSchema(val *string)
	SchemaInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() CloudcontrolapiResourceTimeoutsOutputReference
	TimeoutsInput() *CloudcontrolapiResourceTimeouts
	TypeName() *string
	SetTypeName(val *string)
	TypeNameInput() *string
	TypeVersionId() *string
	SetTypeVersionId(val *string)
	TypeVersionIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *CloudcontrolapiResourceTimeouts)
	ResetOverrideLogicalId()
	ResetRoleArn()
	ResetSchema()
	ResetTimeouts()
	ResetTypeVersionId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudcontrolapiResource
type jsiiProxy_CloudcontrolapiResource struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudcontrolapiResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) Properties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) Timeouts() CloudcontrolapiResourceTimeoutsOutputReference {
	var returns CloudcontrolapiResourceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) TimeoutsInput() *CloudcontrolapiResourceTimeouts {
	var returns *CloudcontrolapiResourceTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) TypeVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResource) TypeVersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeVersionIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html aws_cloudcontrolapi_resource} Resource.
func NewCloudcontrolapiResource(scope constructs.Construct, id *string, config *CloudcontrolapiResourceConfig) CloudcontrolapiResource {
	_init_.Initialize()

	j := jsiiProxy_CloudcontrolapiResource{}

	_jsii_.Create(
		"hashicorp_aws.CloudcontrolapiResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html aws_cloudcontrolapi_resource} Resource.
func NewCloudcontrolapiResource_Override(c CloudcontrolapiResource, scope constructs.Construct, id *string, config *CloudcontrolapiResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudcontrolapiResource",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudcontrolapiResource) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResource) SetDesiredState(val *string) {
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResource) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResource) SetSchema(val *string) {
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResource) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResource) SetTypeVersionId(val *string) {
	_jsii_.Set(
		j,
		"typeVersionId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudcontrolapiResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.CloudcontrolapiResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudcontrolapiResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.CloudcontrolapiResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (c *jsiiProxy_CloudcontrolapiResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (c *jsiiProxy_CloudcontrolapiResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudcontrolapiResource) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudcontrolapiResource) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudcontrolapiResource) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudcontrolapiResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (c *jsiiProxy_CloudcontrolapiResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudcontrolapiResource) PutTimeouts(value *CloudcontrolapiResourceTimeouts) {
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (c *jsiiProxy_CloudcontrolapiResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudcontrolapiResource) ResetRoleArn() {
	_jsii_.InvokeVoid(
		c,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudcontrolapiResource) ResetSchema() {
	_jsii_.InvokeVoid(
		c,
		"resetSchema",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudcontrolapiResource) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudcontrolapiResource) ResetTypeVersionId() {
	_jsii_.InvokeVoid(
		c,
		"resetTypeVersionId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudcontrolapiResource) SynthesizeAttributes() *map[string]interface{} {
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
func (c *jsiiProxy_CloudcontrolapiResource) ToMetadata() interface{} {
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
func (c *jsiiProxy_CloudcontrolapiResource) ToString() *string {
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
func (c *jsiiProxy_CloudcontrolapiResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudcontrolapiResourceConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html#desired_state CloudcontrolapiResource#desired_state}.
	DesiredState *string `json:"desiredState"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html#type_name CloudcontrolapiResource#type_name}.
	TypeName *string `json:"typeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html#role_arn CloudcontrolapiResource#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html#schema CloudcontrolapiResource#schema}.
	Schema *string `json:"schema"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html#timeouts CloudcontrolapiResource#timeouts}
	Timeouts *CloudcontrolapiResourceTimeouts `json:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html#type_version_id CloudcontrolapiResource#type_version_id}.
	TypeVersionId *string `json:"typeVersionId"`
}

type CloudcontrolapiResourceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html#create CloudcontrolapiResource#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html#delete CloudcontrolapiResource#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cloudcontrolapi_resource.html#update CloudcontrolapiResource#update}.
	Update *string `json:"update"`
}

type CloudcontrolapiResourceTimeoutsOutputReference interface {
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

// The jsii proxy struct for CloudcontrolapiResourceTimeoutsOutputReference
type jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewCloudcontrolapiResourceTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) CloudcontrolapiResourceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.CloudcontrolapiResourceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewCloudcontrolapiResourceTimeoutsOutputReference_Override(c CloudcontrolapiResourceTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.CloudcontrolapiResourceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		c,
	)
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (c *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (c *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (c *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (c *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (c *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (c *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		c,
		"resetCreate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudcontrolapiResourceTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/cloudcontrolapi_resource.html aws_cloudcontrolapi_resource}.
type DataAwsCloudcontrolapiResource interface {
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
	Identifier() *string
	SetIdentifier(val *string)
	IdentifierInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Properties() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	TypeName() *string
	SetTypeName(val *string)
	TypeNameInput() *string
	TypeVersionId() *string
	SetTypeVersionId(val *string)
	TypeVersionIdInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetRoleArn()
	ResetTypeVersionId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsCloudcontrolapiResource
type jsiiProxy_DataAwsCloudcontrolapiResource struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) Properties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) TypeVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) TypeVersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeVersionIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudcontrolapi_resource.html aws_cloudcontrolapi_resource} Data Source.
func NewDataAwsCloudcontrolapiResource(scope constructs.Construct, id *string, config *DataAwsCloudcontrolapiResourceConfig) DataAwsCloudcontrolapiResource {
	_init_.Initialize()

	j := jsiiProxy_DataAwsCloudcontrolapiResource{}

	_jsii_.Create(
		"hashicorp_aws.DataAwsCloudcontrolapiResource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/cloudcontrolapi_resource.html aws_cloudcontrolapi_resource} Data Source.
func NewDataAwsCloudcontrolapiResource_Override(d DataAwsCloudcontrolapiResource, scope constructs.Construct, id *string, config *DataAwsCloudcontrolapiResourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DataAwsCloudcontrolapiResource",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) SetIdentifier(val *string) {
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudcontrolapiResource) SetTypeVersionId(val *string) {
	_jsii_.Set(
		j,
		"typeVersionId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsCloudcontrolapiResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DataAwsCloudcontrolapiResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsCloudcontrolapiResource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DataAwsCloudcontrolapiResource",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudcontrolapiResource) ResetRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudcontrolapiResource) ResetTypeVersionId() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeVersionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudcontrolapiResource) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) ToString() *string {
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
func (d *jsiiProxy_DataAwsCloudcontrolapiResource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsCloudcontrolapiResourceConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudcontrolapi_resource.html#identifier DataAwsCloudcontrolapiResource#identifier}.
	Identifier *string `json:"identifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudcontrolapi_resource.html#type_name DataAwsCloudcontrolapiResource#type_name}.
	TypeName *string `json:"typeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudcontrolapi_resource.html#role_arn DataAwsCloudcontrolapiResource#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/cloudcontrolapi_resource.html#type_version_id DataAwsCloudcontrolapiResource#type_version_id}.
	TypeVersionId *string `json:"typeVersionId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/default_tags.html aws_default_tags}.
type DataAwsDefaultTags interface {
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

// The jsii proxy struct for DataAwsDefaultTags
type jsiiProxy_DataAwsDefaultTags struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsDefaultTags) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDefaultTags) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/default_tags.html aws_default_tags} Data Source.
func NewDataAwsDefaultTags(scope constructs.Construct, id *string, config *DataAwsDefaultTagsConfig) DataAwsDefaultTags {
	_init_.Initialize()

	j := jsiiProxy_DataAwsDefaultTags{}

	_jsii_.Create(
		"hashicorp_aws.DataAwsDefaultTags",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/default_tags.html aws_default_tags} Data Source.
func NewDataAwsDefaultTags_Override(d DataAwsDefaultTags, scope constructs.Construct, id *string, config *DataAwsDefaultTagsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DataAwsDefaultTags",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsDefaultTags) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsDefaultTags) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsDefaultTags) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsDefaultTags) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsDefaultTags) SetTags(val interface{}) {
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
func DataAwsDefaultTags_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DataAwsDefaultTags",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsDefaultTags_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DataAwsDefaultTags",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsDefaultTags) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsDefaultTags) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsDefaultTags) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsDefaultTags) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsDefaultTags) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsDefaultTags) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsDefaultTags) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsDefaultTags) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsDefaultTags) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsDefaultTags) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsDefaultTags) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsDefaultTags) ToString() *string {
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
func (d *jsiiProxy_DataAwsDefaultTags) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsDefaultTagsConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/default_tags.html#tags DataAwsDefaultTags#tags}.
	Tags interface{} `json:"tags"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/identitystore_group.html aws_identitystore_group}.
type DataAwsIdentitystoreGroup interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DisplayName() *string
	Filter() *[]*DataAwsIdentitystoreGroupFilter
	SetFilter(val *[]*DataAwsIdentitystoreGroupFilter)
	FilterInput() *[]*DataAwsIdentitystoreGroupFilter
	Fqn() *string
	FriendlyUniqueId() *string
	GroupId() *string
	SetGroupId(val *string)
	GroupIdInput() *string
	Id() *string
	IdentityStoreId() *string
	SetIdentityStoreId(val *string)
	IdentityStoreIdInput() *string
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
	ResetGroupId()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsIdentitystoreGroup
type jsiiProxy_DataAwsIdentitystoreGroup struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) Filter() *[]*DataAwsIdentitystoreGroupFilter {
	var returns *[]*DataAwsIdentitystoreGroupFilter
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) FilterInput() *[]*DataAwsIdentitystoreGroupFilter {
	var returns *[]*DataAwsIdentitystoreGroupFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) GroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) GroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) IdentityStoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityStoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) IdentityStoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityStoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/identitystore_group.html aws_identitystore_group} Data Source.
func NewDataAwsIdentitystoreGroup(scope constructs.Construct, id *string, config *DataAwsIdentitystoreGroupConfig) DataAwsIdentitystoreGroup {
	_init_.Initialize()

	j := jsiiProxy_DataAwsIdentitystoreGroup{}

	_jsii_.Create(
		"hashicorp_aws.DataAwsIdentitystoreGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/identitystore_group.html aws_identitystore_group} Data Source.
func NewDataAwsIdentitystoreGroup_Override(d DataAwsIdentitystoreGroup, scope constructs.Construct, id *string, config *DataAwsIdentitystoreGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DataAwsIdentitystoreGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) SetFilter(val *[]*DataAwsIdentitystoreGroupFilter) {
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) SetGroupId(val *string) {
	_jsii_.Set(
		j,
		"groupId",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) SetIdentityStoreId(val *string) {
	_jsii_.Set(
		j,
		"identityStoreId",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreGroup) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsIdentitystoreGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DataAwsIdentitystoreGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsIdentitystoreGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DataAwsIdentitystoreGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsIdentitystoreGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsIdentitystoreGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsIdentitystoreGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsIdentitystoreGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsIdentitystoreGroup) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsIdentitystoreGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsIdentitystoreGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsIdentitystoreGroup) ResetGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetGroupId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsIdentitystoreGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIdentitystoreGroup) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsIdentitystoreGroup) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsIdentitystoreGroup) ToString() *string {
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
func (d *jsiiProxy_DataAwsIdentitystoreGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsIdentitystoreGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_group.html#filter DataAwsIdentitystoreGroup#filter}
	Filter *[]*DataAwsIdentitystoreGroupFilter `json:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_group.html#identity_store_id DataAwsIdentitystoreGroup#identity_store_id}.
	IdentityStoreId *string `json:"identityStoreId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_group.html#group_id DataAwsIdentitystoreGroup#group_id}.
	GroupId *string `json:"groupId"`
}

type DataAwsIdentitystoreGroupFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_group.html#attribute_path DataAwsIdentitystoreGroup#attribute_path}.
	AttributePath *string `json:"attributePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_group.html#attribute_value DataAwsIdentitystoreGroup#attribute_value}.
	AttributeValue *string `json:"attributeValue"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user.html aws_identitystore_user}.
type DataAwsIdentitystoreUser interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Filter() *[]*DataAwsIdentitystoreUserFilter
	SetFilter(val *[]*DataAwsIdentitystoreUserFilter)
	FilterInput() *[]*DataAwsIdentitystoreUserFilter
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IdentityStoreId() *string
	SetIdentityStoreId(val *string)
	IdentityStoreIdInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
	UserName() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetUserId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsIdentitystoreUser
type jsiiProxy_DataAwsIdentitystoreUser struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) Filter() *[]*DataAwsIdentitystoreUserFilter {
	var returns *[]*DataAwsIdentitystoreUserFilter
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) FilterInput() *[]*DataAwsIdentitystoreUserFilter {
	var returns *[]*DataAwsIdentitystoreUserFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) IdentityStoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityStoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) IdentityStoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityStoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user.html aws_identitystore_user} Data Source.
func NewDataAwsIdentitystoreUser(scope constructs.Construct, id *string, config *DataAwsIdentitystoreUserConfig) DataAwsIdentitystoreUser {
	_init_.Initialize()

	j := jsiiProxy_DataAwsIdentitystoreUser{}

	_jsii_.Create(
		"hashicorp_aws.DataAwsIdentitystoreUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user.html aws_identitystore_user} Data Source.
func NewDataAwsIdentitystoreUser_Override(d DataAwsIdentitystoreUser, scope constructs.Construct, id *string, config *DataAwsIdentitystoreUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DataAwsIdentitystoreUser",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) SetFilter(val *[]*DataAwsIdentitystoreUserFilter) {
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) SetIdentityStoreId(val *string) {
	_jsii_.Set(
		j,
		"identityStoreId",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsIdentitystoreUser) SetUserId(val *string) {
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsIdentitystoreUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DataAwsIdentitystoreUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsIdentitystoreUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DataAwsIdentitystoreUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsIdentitystoreUser) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsIdentitystoreUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsIdentitystoreUser) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsIdentitystoreUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsIdentitystoreUser) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsIdentitystoreUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsIdentitystoreUser) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsIdentitystoreUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIdentitystoreUser) ResetUserId() {
	_jsii_.InvokeVoid(
		d,
		"resetUserId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsIdentitystoreUser) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsIdentitystoreUser) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsIdentitystoreUser) ToString() *string {
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
func (d *jsiiProxy_DataAwsIdentitystoreUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsIdentitystoreUserConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user.html#filter DataAwsIdentitystoreUser#filter}
	Filter *[]*DataAwsIdentitystoreUserFilter `json:"filter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user.html#identity_store_id DataAwsIdentitystoreUser#identity_store_id}.
	IdentityStoreId *string `json:"identityStoreId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user.html#user_id DataAwsIdentitystoreUser#user_id}.
	UserId *string `json:"userId"`
}

type DataAwsIdentitystoreUserFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user.html#attribute_path DataAwsIdentitystoreUser#attribute_path}.
	AttributePath *string `json:"attributePath"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/identitystore_user.html#attribute_value DataAwsIdentitystoreUser#attribute_value}.
	AttributeValue *string `json:"attributeValue"`
}
