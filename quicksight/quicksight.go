package quicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/quicksight/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html aws_quicksight_data_source}.
type QuicksightDataSource interface {
	cdktf.TerraformResource
	Arn() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	Credentials() QuicksightDataSourceCredentialsOutputReference
	CredentialsInput() *QuicksightDataSourceCredentials
	DataSourceId() *string
	SetDataSourceId(val *string)
	DataSourceIdInput() *string
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
	Parameters() QuicksightDataSourceParametersOutputReference
	ParametersInput() *QuicksightDataSourceParameters
	Permission() *[]*QuicksightDataSourcePermission
	SetPermission(val *[]*QuicksightDataSourcePermission)
	PermissionInput() *[]*QuicksightDataSourcePermission
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SslProperties() QuicksightDataSourceSslPropertiesOutputReference
	SslPropertiesInput() *QuicksightDataSourceSslProperties
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
	VpcConnectionProperties() QuicksightDataSourceVpcConnectionPropertiesOutputReference
	VpcConnectionPropertiesInput() *QuicksightDataSourceVpcConnectionProperties
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutCredentials(value *QuicksightDataSourceCredentials)
	PutParameters(value *QuicksightDataSourceParameters)
	PutSslProperties(value *QuicksightDataSourceSslProperties)
	PutVpcConnectionProperties(value *QuicksightDataSourceVpcConnectionProperties)
	ResetAwsAccountId()
	ResetCredentials()
	ResetOverrideLogicalId()
	ResetPermission()
	ResetSslProperties()
	ResetTags()
	ResetTagsAll()
	ResetVpcConnectionProperties()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for QuicksightDataSource
type jsiiProxy_QuicksightDataSource struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QuicksightDataSource) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Credentials() QuicksightDataSourceCredentialsOutputReference {
	var returns QuicksightDataSourceCredentialsOutputReference
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) CredentialsInput() *QuicksightDataSourceCredentials {
	var returns *QuicksightDataSourceCredentials
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) DataSourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Parameters() QuicksightDataSourceParametersOutputReference {
	var returns QuicksightDataSourceParametersOutputReference
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) ParametersInput() *QuicksightDataSourceParameters {
	var returns *QuicksightDataSourceParameters
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Permission() *[]*QuicksightDataSourcePermission {
	var returns *[]*QuicksightDataSourcePermission
	_jsii_.Get(
		j,
		"permission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) PermissionInput() *[]*QuicksightDataSourcePermission {
	var returns *[]*QuicksightDataSourcePermission
	_jsii_.Get(
		j,
		"permissionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) SslProperties() QuicksightDataSourceSslPropertiesOutputReference {
	var returns QuicksightDataSourceSslPropertiesOutputReference
	_jsii_.Get(
		j,
		"sslProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) SslPropertiesInput() *QuicksightDataSourceSslProperties {
	var returns *QuicksightDataSourceSslProperties
	_jsii_.Get(
		j,
		"sslPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) VpcConnectionProperties() QuicksightDataSourceVpcConnectionPropertiesOutputReference {
	var returns QuicksightDataSourceVpcConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"vpcConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSource) VpcConnectionPropertiesInput() *QuicksightDataSourceVpcConnectionProperties {
	var returns *QuicksightDataSourceVpcConnectionProperties
	_jsii_.Get(
		j,
		"vpcConnectionPropertiesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html aws_quicksight_data_source} Resource.
func NewQuicksightDataSource(scope constructs.Construct, id *string, config *QuicksightDataSourceConfig) QuicksightDataSource {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSource{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSource",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html aws_quicksight_data_source} Resource.
func NewQuicksightDataSource_Override(q QuicksightDataSource, scope constructs.Construct, id *string, config *QuicksightDataSourceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSource",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetDataSourceId(val *string) {
	_jsii_.Set(
		j,
		"dataSourceId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetPermission(val *[]*QuicksightDataSourcePermission) {
	_jsii_.Set(
		j,
		"permission",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSource) SetType(val *string) {
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
func QuicksightDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.QuickSight.QuicksightDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QuicksightDataSource_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.QuickSight.QuicksightDataSource",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSource) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSource) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSource) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSource) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (q *jsiiProxy_QuicksightDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QuicksightDataSource) PutCredentials(value *QuicksightDataSourceCredentials) {
	_jsii_.InvokeVoid(
		q,
		"putCredentials",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSource) PutParameters(value *QuicksightDataSourceParameters) {
	_jsii_.InvokeVoid(
		q,
		"putParameters",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSource) PutSslProperties(value *QuicksightDataSourceSslProperties) {
	_jsii_.InvokeVoid(
		q,
		"putSslProperties",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSource) PutVpcConnectionProperties(value *QuicksightDataSourceVpcConnectionProperties) {
	_jsii_.InvokeVoid(
		q,
		"putVpcConnectionProperties",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSource) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		q,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSource) ResetCredentials() {
	_jsii_.InvokeVoid(
		q,
		"resetCredentials",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (q *jsiiProxy_QuicksightDataSource) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSource) ResetPermission() {
	_jsii_.InvokeVoid(
		q,
		"resetPermission",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSource) ResetSslProperties() {
	_jsii_.InvokeVoid(
		q,
		"resetSslProperties",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSource) ResetTags() {
	_jsii_.InvokeVoid(
		q,
		"resetTags",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSource) ResetTagsAll() {
	_jsii_.InvokeVoid(
		q,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSource) ResetVpcConnectionProperties() {
	_jsii_.InvokeVoid(
		q,
		"resetVpcConnectionProperties",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSource) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSource) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (q *jsiiProxy_QuicksightDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (q *jsiiProxy_QuicksightDataSource) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type QuicksightDataSourceConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#data_source_id QuicksightDataSource#data_source_id}.
	DataSourceId *string `json:"dataSourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#name QuicksightDataSource#name}.
	Name *string `json:"name"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#parameters QuicksightDataSource#parameters}
	Parameters *QuicksightDataSourceParameters `json:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#type QuicksightDataSource#type}.
	Type *string `json:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#aws_account_id QuicksightDataSource#aws_account_id}.
	AwsAccountId *string `json:"awsAccountId"`
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#credentials QuicksightDataSource#credentials}
	Credentials *QuicksightDataSourceCredentials `json:"credentials"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#permission QuicksightDataSource#permission}
	Permission *[]*QuicksightDataSourcePermission `json:"permission"`
	// ssl_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#ssl_properties QuicksightDataSource#ssl_properties}
	SslProperties *QuicksightDataSourceSslProperties `json:"sslProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#tags QuicksightDataSource#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#tags_all QuicksightDataSource#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// vpc_connection_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#vpc_connection_properties QuicksightDataSource#vpc_connection_properties}
	VpcConnectionProperties *QuicksightDataSourceVpcConnectionProperties `json:"vpcConnectionProperties"`
}

type QuicksightDataSourceCredentials struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#copy_source_arn QuicksightDataSource#copy_source_arn}.
	CopySourceArn *string `json:"copySourceArn"`
	// credential_pair block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#credential_pair QuicksightDataSource#credential_pair}
	CredentialPair *QuicksightDataSourceCredentialsCredentialPair `json:"credentialPair"`
}

type QuicksightDataSourceCredentialsCredentialPair struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#password QuicksightDataSource#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#username QuicksightDataSource#username}.
	Username *string `json:"username"`
}

type QuicksightDataSourceCredentialsCredentialPairOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
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
}

// The jsii proxy struct for QuicksightDataSourceCredentialsCredentialPairOutputReference
type jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceCredentialsCredentialPairOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceCredentialsCredentialPairOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceCredentialsCredentialPairOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceCredentialsCredentialPairOutputReference_Override(q QuicksightDataSourceCredentialsCredentialPairOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceCredentialsCredentialPairOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsCredentialPairOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceCredentialsOutputReference interface {
	cdktf.ComplexObject
	CopySourceArn() *string
	SetCopySourceArn(val *string)
	CopySourceArnInput() *string
	CredentialPair() QuicksightDataSourceCredentialsCredentialPairOutputReference
	CredentialPairInput() *QuicksightDataSourceCredentialsCredentialPair
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
	PutCredentialPair(value *QuicksightDataSourceCredentialsCredentialPair)
	ResetCopySourceArn()
	ResetCredentialPair()
}

// The jsii proxy struct for QuicksightDataSourceCredentialsOutputReference
type jsiiProxy_QuicksightDataSourceCredentialsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) CopySourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copySourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) CopySourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copySourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) CredentialPair() QuicksightDataSourceCredentialsCredentialPairOutputReference {
	var returns QuicksightDataSourceCredentialsCredentialPairOutputReference
	_jsii_.Get(
		j,
		"credentialPair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) CredentialPairInput() *QuicksightDataSourceCredentialsCredentialPair {
	var returns *QuicksightDataSourceCredentialsCredentialPair
	_jsii_.Get(
		j,
		"credentialPairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceCredentialsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceCredentialsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceCredentialsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceCredentialsOutputReference_Override(q QuicksightDataSourceCredentialsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) SetCopySourceArn(val *string) {
	_jsii_.Set(
		j,
		"copySourceArn",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) PutCredentialPair(value *QuicksightDataSourceCredentialsCredentialPair) {
	_jsii_.InvokeVoid(
		q,
		"putCredentialPair",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) ResetCopySourceArn() {
	_jsii_.InvokeVoid(
		q,
		"resetCopySourceArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceCredentialsOutputReference) ResetCredentialPair() {
	_jsii_.InvokeVoid(
		q,
		"resetCredentialPair",
		nil, // no parameters
	)
}

type QuicksightDataSourceParameters struct {
	// amazon_elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#amazon_elasticsearch QuicksightDataSource#amazon_elasticsearch}
	AmazonElasticsearch *QuicksightDataSourceParametersAmazonElasticsearch `json:"amazonElasticsearch"`
	// athena block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#athena QuicksightDataSource#athena}
	Athena *QuicksightDataSourceParametersAthena `json:"athena"`
	// aurora block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#aurora QuicksightDataSource#aurora}
	Aurora *QuicksightDataSourceParametersAurora `json:"aurora"`
	// aurora_postgresql block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#aurora_postgresql QuicksightDataSource#aurora_postgresql}
	AuroraPostgresql *QuicksightDataSourceParametersAuroraPostgresql `json:"auroraPostgresql"`
	// aws_iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#aws_iot_analytics QuicksightDataSource#aws_iot_analytics}
	AwsIotAnalytics *QuicksightDataSourceParametersAwsIotAnalytics `json:"awsIotAnalytics"`
	// jira block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#jira QuicksightDataSource#jira}
	Jira *QuicksightDataSourceParametersJira `json:"jira"`
	// maria_db block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#maria_db QuicksightDataSource#maria_db}
	MariaDb *QuicksightDataSourceParametersMariaDb `json:"mariaDb"`
	// mysql block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#mysql QuicksightDataSource#mysql}
	Mysql *QuicksightDataSourceParametersMysql `json:"mysql"`
	// oracle block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#oracle QuicksightDataSource#oracle}
	Oracle *QuicksightDataSourceParametersOracle `json:"oracle"`
	// postgresql block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#postgresql QuicksightDataSource#postgresql}
	Postgresql *QuicksightDataSourceParametersPostgresql `json:"postgresql"`
	// presto block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#presto QuicksightDataSource#presto}
	Presto *QuicksightDataSourceParametersPresto `json:"presto"`
	// rds block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#rds QuicksightDataSource#rds}
	Rds *QuicksightDataSourceParametersRds `json:"rds"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#redshift QuicksightDataSource#redshift}
	Redshift *QuicksightDataSourceParametersRedshift `json:"redshift"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#s3 QuicksightDataSource#s3}
	S3 *QuicksightDataSourceParametersS3 `json:"s3"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#service_now QuicksightDataSource#service_now}
	ServiceNow *QuicksightDataSourceParametersServiceNow `json:"serviceNow"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#snowflake QuicksightDataSource#snowflake}
	Snowflake *QuicksightDataSourceParametersSnowflake `json:"snowflake"`
	// spark block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#spark QuicksightDataSource#spark}
	Spark *QuicksightDataSourceParametersSpark `json:"spark"`
	// sql_server block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#sql_server QuicksightDataSource#sql_server}
	SqlServer *QuicksightDataSourceParametersSqlServer `json:"sqlServer"`
	// teradata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#teradata QuicksightDataSource#teradata}
	Teradata *QuicksightDataSourceParametersTeradata `json:"teradata"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#twitter QuicksightDataSource#twitter}
	Twitter *QuicksightDataSourceParametersTwitter `json:"twitter"`
}

type QuicksightDataSourceParametersAmazonElasticsearch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#domain QuicksightDataSource#domain}.
	Domain *string `json:"domain"`
}

type QuicksightDataSourceParametersAmazonElasticsearchOutputReference interface {
	cdktf.ComplexObject
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
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
}

// The jsii proxy struct for QuicksightDataSourceParametersAmazonElasticsearchOutputReference
type jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersAmazonElasticsearchOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersAmazonElasticsearchOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAmazonElasticsearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersAmazonElasticsearchOutputReference_Override(q QuicksightDataSourceParametersAmazonElasticsearchOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAmazonElasticsearchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAmazonElasticsearchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersAthena struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#work_group QuicksightDataSource#work_group}.
	WorkGroup *string `json:"workGroup"`
}

type QuicksightDataSourceParametersAthenaOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	WorkGroup() *string
	SetWorkGroup(val *string)
	WorkGroupInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetWorkGroup()
}

// The jsii proxy struct for QuicksightDataSourceParametersAthenaOutputReference
type jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) WorkGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) WorkGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workGroupInput",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersAthenaOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersAthenaOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAthenaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersAthenaOutputReference_Override(q QuicksightDataSourceParametersAthenaOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAthenaOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) SetWorkGroup(val *string) {
	_jsii_.Set(
		j,
		"workGroup",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersAthenaOutputReference) ResetWorkGroup() {
	_jsii_.InvokeVoid(
		q,
		"resetWorkGroup",
		nil, // no parameters
	)
}

type QuicksightDataSourceParametersAurora struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersAuroraOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersAuroraOutputReference
type jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersAuroraOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersAuroraOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAuroraOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersAuroraOutputReference_Override(q QuicksightDataSourceParametersAuroraOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAuroraOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersAuroraPostgresql struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersAuroraPostgresqlOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersAuroraPostgresqlOutputReference
type jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersAuroraPostgresqlOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersAuroraPostgresqlOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAuroraPostgresqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersAuroraPostgresqlOutputReference_Override(q QuicksightDataSourceParametersAuroraPostgresqlOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAuroraPostgresqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAuroraPostgresqlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersAwsIotAnalytics struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#data_set_name QuicksightDataSource#data_set_name}.
	DataSetName *string `json:"dataSetName"`
}

type QuicksightDataSourceParametersAwsIotAnalyticsOutputReference interface {
	cdktf.ComplexObject
	DataSetName() *string
	SetDataSetName(val *string)
	DataSetNameInput() *string
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
}

// The jsii proxy struct for QuicksightDataSourceParametersAwsIotAnalyticsOutputReference
type jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) DataSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) DataSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersAwsIotAnalyticsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersAwsIotAnalyticsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAwsIotAnalyticsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersAwsIotAnalyticsOutputReference_Override(q QuicksightDataSourceParametersAwsIotAnalyticsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersAwsIotAnalyticsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) SetDataSetName(val *string) {
	_jsii_.Set(
		j,
		"dataSetName",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersAwsIotAnalyticsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersJira struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#site_base_url QuicksightDataSource#site_base_url}.
	SiteBaseUrl *string `json:"siteBaseUrl"`
}

type QuicksightDataSourceParametersJiraOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SiteBaseUrl() *string
	SetSiteBaseUrl(val *string)
	SiteBaseUrlInput() *string
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

// The jsii proxy struct for QuicksightDataSourceParametersJiraOutputReference
type jsiiProxy_QuicksightDataSourceParametersJiraOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) SiteBaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteBaseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) SiteBaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteBaseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersJiraOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersJiraOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersJiraOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersJiraOutputReference_Override(q QuicksightDataSourceParametersJiraOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersJiraOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) SetSiteBaseUrl(val *string) {
	_jsii_.Set(
		j,
		"siteBaseUrl",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersJiraOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersMariaDb struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersMariaDbOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersMariaDbOutputReference
type jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersMariaDbOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersMariaDbOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersMariaDbOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersMariaDbOutputReference_Override(q QuicksightDataSourceParametersMariaDbOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersMariaDbOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMariaDbOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersMysql struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersMysqlOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersMysqlOutputReference
type jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersMysqlOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersMysqlOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersMysqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersMysqlOutputReference_Override(q QuicksightDataSourceParametersMysqlOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersMysqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersMysqlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersOracle struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersOracleOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersOracleOutputReference
type jsiiProxy_QuicksightDataSourceParametersOracleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersOracleOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersOracleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersOracleOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersOracleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersOracleOutputReference_Override(q QuicksightDataSourceParametersOracleOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersOracleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOracleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersOutputReference interface {
	cdktf.ComplexObject
	AmazonElasticsearch() QuicksightDataSourceParametersAmazonElasticsearchOutputReference
	AmazonElasticsearchInput() *QuicksightDataSourceParametersAmazonElasticsearch
	Athena() QuicksightDataSourceParametersAthenaOutputReference
	AthenaInput() *QuicksightDataSourceParametersAthena
	Aurora() QuicksightDataSourceParametersAuroraOutputReference
	AuroraInput() *QuicksightDataSourceParametersAurora
	AuroraPostgresql() QuicksightDataSourceParametersAuroraPostgresqlOutputReference
	AuroraPostgresqlInput() *QuicksightDataSourceParametersAuroraPostgresql
	AwsIotAnalytics() QuicksightDataSourceParametersAwsIotAnalyticsOutputReference
	AwsIotAnalyticsInput() *QuicksightDataSourceParametersAwsIotAnalytics
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Jira() QuicksightDataSourceParametersJiraOutputReference
	JiraInput() *QuicksightDataSourceParametersJira
	MariaDb() QuicksightDataSourceParametersMariaDbOutputReference
	MariaDbInput() *QuicksightDataSourceParametersMariaDb
	Mysql() QuicksightDataSourceParametersMysqlOutputReference
	MysqlInput() *QuicksightDataSourceParametersMysql
	Oracle() QuicksightDataSourceParametersOracleOutputReference
	OracleInput() *QuicksightDataSourceParametersOracle
	Postgresql() QuicksightDataSourceParametersPostgresqlOutputReference
	PostgresqlInput() *QuicksightDataSourceParametersPostgresql
	Presto() QuicksightDataSourceParametersPrestoOutputReference
	PrestoInput() *QuicksightDataSourceParametersPresto
	Rds() QuicksightDataSourceParametersRdsOutputReference
	RdsInput() *QuicksightDataSourceParametersRds
	Redshift() QuicksightDataSourceParametersRedshiftOutputReference
	RedshiftInput() *QuicksightDataSourceParametersRedshift
	S3() QuicksightDataSourceParametersS3OutputReference
	S3Input() *QuicksightDataSourceParametersS3
	ServiceNow() QuicksightDataSourceParametersServiceNowOutputReference
	ServiceNowInput() *QuicksightDataSourceParametersServiceNow
	Snowflake() QuicksightDataSourceParametersSnowflakeOutputReference
	SnowflakeInput() *QuicksightDataSourceParametersSnowflake
	Spark() QuicksightDataSourceParametersSparkOutputReference
	SparkInput() *QuicksightDataSourceParametersSpark
	SqlServer() QuicksightDataSourceParametersSqlServerOutputReference
	SqlServerInput() *QuicksightDataSourceParametersSqlServer
	Teradata() QuicksightDataSourceParametersTeradataOutputReference
	TeradataInput() *QuicksightDataSourceParametersTeradata
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Twitter() QuicksightDataSourceParametersTwitterOutputReference
	TwitterInput() *QuicksightDataSourceParametersTwitter
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutAmazonElasticsearch(value *QuicksightDataSourceParametersAmazonElasticsearch)
	PutAthena(value *QuicksightDataSourceParametersAthena)
	PutAurora(value *QuicksightDataSourceParametersAurora)
	PutAuroraPostgresql(value *QuicksightDataSourceParametersAuroraPostgresql)
	PutAwsIotAnalytics(value *QuicksightDataSourceParametersAwsIotAnalytics)
	PutJira(value *QuicksightDataSourceParametersJira)
	PutMariaDb(value *QuicksightDataSourceParametersMariaDb)
	PutMysql(value *QuicksightDataSourceParametersMysql)
	PutOracle(value *QuicksightDataSourceParametersOracle)
	PutPostgresql(value *QuicksightDataSourceParametersPostgresql)
	PutPresto(value *QuicksightDataSourceParametersPresto)
	PutRds(value *QuicksightDataSourceParametersRds)
	PutRedshift(value *QuicksightDataSourceParametersRedshift)
	PutS3(value *QuicksightDataSourceParametersS3)
	PutServiceNow(value *QuicksightDataSourceParametersServiceNow)
	PutSnowflake(value *QuicksightDataSourceParametersSnowflake)
	PutSpark(value *QuicksightDataSourceParametersSpark)
	PutSqlServer(value *QuicksightDataSourceParametersSqlServer)
	PutTeradata(value *QuicksightDataSourceParametersTeradata)
	PutTwitter(value *QuicksightDataSourceParametersTwitter)
	ResetAmazonElasticsearch()
	ResetAthena()
	ResetAurora()
	ResetAuroraPostgresql()
	ResetAwsIotAnalytics()
	ResetJira()
	ResetMariaDb()
	ResetMysql()
	ResetOracle()
	ResetPostgresql()
	ResetPresto()
	ResetRds()
	ResetRedshift()
	ResetS3()
	ResetServiceNow()
	ResetSnowflake()
	ResetSpark()
	ResetSqlServer()
	ResetTeradata()
	ResetTwitter()
}

// The jsii proxy struct for QuicksightDataSourceParametersOutputReference
type jsiiProxy_QuicksightDataSourceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AmazonElasticsearch() QuicksightDataSourceParametersAmazonElasticsearchOutputReference {
	var returns QuicksightDataSourceParametersAmazonElasticsearchOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AmazonElasticsearchInput() *QuicksightDataSourceParametersAmazonElasticsearch {
	var returns *QuicksightDataSourceParametersAmazonElasticsearch
	_jsii_.Get(
		j,
		"amazonElasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Athena() QuicksightDataSourceParametersAthenaOutputReference {
	var returns QuicksightDataSourceParametersAthenaOutputReference
	_jsii_.Get(
		j,
		"athena",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AthenaInput() *QuicksightDataSourceParametersAthena {
	var returns *QuicksightDataSourceParametersAthena
	_jsii_.Get(
		j,
		"athenaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Aurora() QuicksightDataSourceParametersAuroraOutputReference {
	var returns QuicksightDataSourceParametersAuroraOutputReference
	_jsii_.Get(
		j,
		"aurora",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AuroraInput() *QuicksightDataSourceParametersAurora {
	var returns *QuicksightDataSourceParametersAurora
	_jsii_.Get(
		j,
		"auroraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AuroraPostgresql() QuicksightDataSourceParametersAuroraPostgresqlOutputReference {
	var returns QuicksightDataSourceParametersAuroraPostgresqlOutputReference
	_jsii_.Get(
		j,
		"auroraPostgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AuroraPostgresqlInput() *QuicksightDataSourceParametersAuroraPostgresql {
	var returns *QuicksightDataSourceParametersAuroraPostgresql
	_jsii_.Get(
		j,
		"auroraPostgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AwsIotAnalytics() QuicksightDataSourceParametersAwsIotAnalyticsOutputReference {
	var returns QuicksightDataSourceParametersAwsIotAnalyticsOutputReference
	_jsii_.Get(
		j,
		"awsIotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) AwsIotAnalyticsInput() *QuicksightDataSourceParametersAwsIotAnalytics {
	var returns *QuicksightDataSourceParametersAwsIotAnalytics
	_jsii_.Get(
		j,
		"awsIotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Jira() QuicksightDataSourceParametersJiraOutputReference {
	var returns QuicksightDataSourceParametersJiraOutputReference
	_jsii_.Get(
		j,
		"jira",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) JiraInput() *QuicksightDataSourceParametersJira {
	var returns *QuicksightDataSourceParametersJira
	_jsii_.Get(
		j,
		"jiraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) MariaDb() QuicksightDataSourceParametersMariaDbOutputReference {
	var returns QuicksightDataSourceParametersMariaDbOutputReference
	_jsii_.Get(
		j,
		"mariaDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) MariaDbInput() *QuicksightDataSourceParametersMariaDb {
	var returns *QuicksightDataSourceParametersMariaDb
	_jsii_.Get(
		j,
		"mariaDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Mysql() QuicksightDataSourceParametersMysqlOutputReference {
	var returns QuicksightDataSourceParametersMysqlOutputReference
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) MysqlInput() *QuicksightDataSourceParametersMysql {
	var returns *QuicksightDataSourceParametersMysql
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Oracle() QuicksightDataSourceParametersOracleOutputReference {
	var returns QuicksightDataSourceParametersOracleOutputReference
	_jsii_.Get(
		j,
		"oracle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) OracleInput() *QuicksightDataSourceParametersOracle {
	var returns *QuicksightDataSourceParametersOracle
	_jsii_.Get(
		j,
		"oracleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Postgresql() QuicksightDataSourceParametersPostgresqlOutputReference {
	var returns QuicksightDataSourceParametersPostgresqlOutputReference
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) PostgresqlInput() *QuicksightDataSourceParametersPostgresql {
	var returns *QuicksightDataSourceParametersPostgresql
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Presto() QuicksightDataSourceParametersPrestoOutputReference {
	var returns QuicksightDataSourceParametersPrestoOutputReference
	_jsii_.Get(
		j,
		"presto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) PrestoInput() *QuicksightDataSourceParametersPresto {
	var returns *QuicksightDataSourceParametersPresto
	_jsii_.Get(
		j,
		"prestoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Rds() QuicksightDataSourceParametersRdsOutputReference {
	var returns QuicksightDataSourceParametersRdsOutputReference
	_jsii_.Get(
		j,
		"rds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) RdsInput() *QuicksightDataSourceParametersRds {
	var returns *QuicksightDataSourceParametersRds
	_jsii_.Get(
		j,
		"rdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Redshift() QuicksightDataSourceParametersRedshiftOutputReference {
	var returns QuicksightDataSourceParametersRedshiftOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) RedshiftInput() *QuicksightDataSourceParametersRedshift {
	var returns *QuicksightDataSourceParametersRedshift
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) S3() QuicksightDataSourceParametersS3OutputReference {
	var returns QuicksightDataSourceParametersS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) S3Input() *QuicksightDataSourceParametersS3 {
	var returns *QuicksightDataSourceParametersS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) ServiceNow() QuicksightDataSourceParametersServiceNowOutputReference {
	var returns QuicksightDataSourceParametersServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) ServiceNowInput() *QuicksightDataSourceParametersServiceNow {
	var returns *QuicksightDataSourceParametersServiceNow
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Snowflake() QuicksightDataSourceParametersSnowflakeOutputReference {
	var returns QuicksightDataSourceParametersSnowflakeOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SnowflakeInput() *QuicksightDataSourceParametersSnowflake {
	var returns *QuicksightDataSourceParametersSnowflake
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Spark() QuicksightDataSourceParametersSparkOutputReference {
	var returns QuicksightDataSourceParametersSparkOutputReference
	_jsii_.Get(
		j,
		"spark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SparkInput() *QuicksightDataSourceParametersSpark {
	var returns *QuicksightDataSourceParametersSpark
	_jsii_.Get(
		j,
		"sparkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SqlServer() QuicksightDataSourceParametersSqlServerOutputReference {
	var returns QuicksightDataSourceParametersSqlServerOutputReference
	_jsii_.Get(
		j,
		"sqlServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SqlServerInput() *QuicksightDataSourceParametersSqlServer {
	var returns *QuicksightDataSourceParametersSqlServer
	_jsii_.Get(
		j,
		"sqlServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Teradata() QuicksightDataSourceParametersTeradataOutputReference {
	var returns QuicksightDataSourceParametersTeradataOutputReference
	_jsii_.Get(
		j,
		"teradata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) TeradataInput() *QuicksightDataSourceParametersTeradata {
	var returns *QuicksightDataSourceParametersTeradata
	_jsii_.Get(
		j,
		"teradataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) Twitter() QuicksightDataSourceParametersTwitterOutputReference {
	var returns QuicksightDataSourceParametersTwitterOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) TwitterInput() *QuicksightDataSourceParametersTwitter {
	var returns *QuicksightDataSourceParametersTwitter
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersOutputReference_Override(q QuicksightDataSourceParametersOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAmazonElasticsearch(value *QuicksightDataSourceParametersAmazonElasticsearch) {
	_jsii_.InvokeVoid(
		q,
		"putAmazonElasticsearch",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAthena(value *QuicksightDataSourceParametersAthena) {
	_jsii_.InvokeVoid(
		q,
		"putAthena",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAurora(value *QuicksightDataSourceParametersAurora) {
	_jsii_.InvokeVoid(
		q,
		"putAurora",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAuroraPostgresql(value *QuicksightDataSourceParametersAuroraPostgresql) {
	_jsii_.InvokeVoid(
		q,
		"putAuroraPostgresql",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutAwsIotAnalytics(value *QuicksightDataSourceParametersAwsIotAnalytics) {
	_jsii_.InvokeVoid(
		q,
		"putAwsIotAnalytics",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutJira(value *QuicksightDataSourceParametersJira) {
	_jsii_.InvokeVoid(
		q,
		"putJira",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutMariaDb(value *QuicksightDataSourceParametersMariaDb) {
	_jsii_.InvokeVoid(
		q,
		"putMariaDb",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutMysql(value *QuicksightDataSourceParametersMysql) {
	_jsii_.InvokeVoid(
		q,
		"putMysql",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutOracle(value *QuicksightDataSourceParametersOracle) {
	_jsii_.InvokeVoid(
		q,
		"putOracle",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutPostgresql(value *QuicksightDataSourceParametersPostgresql) {
	_jsii_.InvokeVoid(
		q,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutPresto(value *QuicksightDataSourceParametersPresto) {
	_jsii_.InvokeVoid(
		q,
		"putPresto",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutRds(value *QuicksightDataSourceParametersRds) {
	_jsii_.InvokeVoid(
		q,
		"putRds",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutRedshift(value *QuicksightDataSourceParametersRedshift) {
	_jsii_.InvokeVoid(
		q,
		"putRedshift",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutS3(value *QuicksightDataSourceParametersS3) {
	_jsii_.InvokeVoid(
		q,
		"putS3",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutServiceNow(value *QuicksightDataSourceParametersServiceNow) {
	_jsii_.InvokeVoid(
		q,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutSnowflake(value *QuicksightDataSourceParametersSnowflake) {
	_jsii_.InvokeVoid(
		q,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutSpark(value *QuicksightDataSourceParametersSpark) {
	_jsii_.InvokeVoid(
		q,
		"putSpark",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutSqlServer(value *QuicksightDataSourceParametersSqlServer) {
	_jsii_.InvokeVoid(
		q,
		"putSqlServer",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutTeradata(value *QuicksightDataSourceParametersTeradata) {
	_jsii_.InvokeVoid(
		q,
		"putTeradata",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) PutTwitter(value *QuicksightDataSourceParametersTwitter) {
	_jsii_.InvokeVoid(
		q,
		"putTwitter",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAmazonElasticsearch() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonElasticsearch",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAthena() {
	_jsii_.InvokeVoid(
		q,
		"resetAthena",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAurora() {
	_jsii_.InvokeVoid(
		q,
		"resetAurora",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAuroraPostgresql() {
	_jsii_.InvokeVoid(
		q,
		"resetAuroraPostgresql",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetAwsIotAnalytics() {
	_jsii_.InvokeVoid(
		q,
		"resetAwsIotAnalytics",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetJira() {
	_jsii_.InvokeVoid(
		q,
		"resetJira",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetMariaDb() {
	_jsii_.InvokeVoid(
		q,
		"resetMariaDb",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetMysql() {
	_jsii_.InvokeVoid(
		q,
		"resetMysql",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetOracle() {
	_jsii_.InvokeVoid(
		q,
		"resetOracle",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetPostgresql() {
	_jsii_.InvokeVoid(
		q,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetPresto() {
	_jsii_.InvokeVoid(
		q,
		"resetPresto",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetRds() {
	_jsii_.InvokeVoid(
		q,
		"resetRds",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		q,
		"resetRedshift",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		q,
		"resetS3",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		q,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		q,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetSpark() {
	_jsii_.InvokeVoid(
		q,
		"resetSpark",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetSqlServer() {
	_jsii_.InvokeVoid(
		q,
		"resetSqlServer",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetTeradata() {
	_jsii_.InvokeVoid(
		q,
		"resetTeradata",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		q,
		"resetTwitter",
		nil, // no parameters
	)
}

type QuicksightDataSourceParametersPostgresql struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersPostgresqlOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersPostgresqlOutputReference
type jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersPostgresqlOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersPostgresqlOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersPostgresqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersPostgresqlOutputReference_Override(q QuicksightDataSourceParametersPostgresqlOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersPostgresqlOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPostgresqlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersPresto struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#catalog QuicksightDataSource#catalog}.
	Catalog *string `json:"catalog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersPrestoOutputReference interface {
	cdktf.ComplexObject
	Catalog() *string
	SetCatalog(val *string)
	CatalogInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersPrestoOutputReference
type jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) Catalog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) CatalogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersPrestoOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersPrestoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersPrestoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersPrestoOutputReference_Override(q QuicksightDataSourceParametersPrestoOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersPrestoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) SetCatalog(val *string) {
	_jsii_.Set(
		j,
		"catalog",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersPrestoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersRds struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#instance_id QuicksightDataSource#instance_id}.
	InstanceId *string `json:"instanceId"`
}

type QuicksightDataSourceParametersRdsOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
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
}

// The jsii proxy struct for QuicksightDataSourceParametersRdsOutputReference
type jsiiProxy_QuicksightDataSourceParametersRdsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersRdsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersRdsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersRdsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersRdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersRdsOutputReference_Override(q QuicksightDataSourceParametersRdsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersRdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) SetInstanceId(val *string) {
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRdsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersRedshift struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#cluster_id QuicksightDataSource#cluster_id}.
	ClusterId *string `json:"clusterId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersRedshiftOutputReference interface {
	cdktf.ComplexObject
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	ResetClusterId()
	ResetHost()
	ResetPort()
}

// The jsii proxy struct for QuicksightDataSourceParametersRedshiftOutputReference
type jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersRedshiftOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersRedshiftOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersRedshiftOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersRedshiftOutputReference_Override(q QuicksightDataSourceParametersRedshiftOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersRedshiftOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		q,
		"resetClusterId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		q,
		"resetHost",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceParametersRedshiftOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		q,
		"resetPort",
		nil, // no parameters
	)
}

type QuicksightDataSourceParametersS3 struct {
	// manifest_file_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#manifest_file_location QuicksightDataSource#manifest_file_location}
	ManifestFileLocation *QuicksightDataSourceParametersS3ManifestFileLocation `json:"manifestFileLocation"`
}

type QuicksightDataSourceParametersS3ManifestFileLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#bucket QuicksightDataSource#bucket}.
	Bucket *string `json:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#key QuicksightDataSource#key}.
	Key *string `json:"key"`
}

type QuicksightDataSourceParametersS3ManifestFileLocationOutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
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

// The jsii proxy struct for QuicksightDataSourceParametersS3ManifestFileLocationOutputReference
type jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersS3ManifestFileLocationOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersS3ManifestFileLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersS3ManifestFileLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersS3ManifestFileLocationOutputReference_Override(q QuicksightDataSourceParametersS3ManifestFileLocationOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersS3ManifestFileLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3ManifestFileLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersS3OutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ManifestFileLocation() QuicksightDataSourceParametersS3ManifestFileLocationOutputReference
	ManifestFileLocationInput() *QuicksightDataSourceParametersS3ManifestFileLocation
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
	PutManifestFileLocation(value *QuicksightDataSourceParametersS3ManifestFileLocation)
}

// The jsii proxy struct for QuicksightDataSourceParametersS3OutputReference
type jsiiProxy_QuicksightDataSourceParametersS3OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) ManifestFileLocation() QuicksightDataSourceParametersS3ManifestFileLocationOutputReference {
	var returns QuicksightDataSourceParametersS3ManifestFileLocationOutputReference
	_jsii_.Get(
		j,
		"manifestFileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) ManifestFileLocationInput() *QuicksightDataSourceParametersS3ManifestFileLocation {
	var returns *QuicksightDataSourceParametersS3ManifestFileLocation
	_jsii_.Get(
		j,
		"manifestFileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersS3OutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersS3OutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersS3OutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersS3OutputReference_Override(q QuicksightDataSourceParametersS3OutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceParametersS3OutputReference) PutManifestFileLocation(value *QuicksightDataSourceParametersS3ManifestFileLocation) {
	_jsii_.InvokeVoid(
		q,
		"putManifestFileLocation",
		[]interface{}{value},
	)
}

type QuicksightDataSourceParametersServiceNow struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#site_base_url QuicksightDataSource#site_base_url}.
	SiteBaseUrl *string `json:"siteBaseUrl"`
}

type QuicksightDataSourceParametersServiceNowOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SiteBaseUrl() *string
	SetSiteBaseUrl(val *string)
	SiteBaseUrlInput() *string
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

// The jsii proxy struct for QuicksightDataSourceParametersServiceNowOutputReference
type jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) SiteBaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteBaseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) SiteBaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteBaseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersServiceNowOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersServiceNowOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersServiceNowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersServiceNowOutputReference_Override(q QuicksightDataSourceParametersServiceNowOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersServiceNowOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) SetSiteBaseUrl(val *string) {
	_jsii_.Set(
		j,
		"siteBaseUrl",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersServiceNowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersSnowflake struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#warehouse QuicksightDataSource#warehouse}.
	Warehouse *string `json:"warehouse"`
}

type QuicksightDataSourceParametersSnowflakeOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Warehouse() *string
	SetWarehouse(val *string)
	WarehouseInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for QuicksightDataSourceParametersSnowflakeOutputReference
type jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) Warehouse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) WarehouseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warehouseInput",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersSnowflakeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersSnowflakeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersSnowflakeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersSnowflakeOutputReference_Override(q QuicksightDataSourceParametersSnowflakeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersSnowflakeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) SetWarehouse(val *string) {
	_jsii_.Set(
		j,
		"warehouse",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSnowflakeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersSpark struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersSparkOutputReference interface {
	cdktf.ComplexObject
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersSparkOutputReference
type jsiiProxy_QuicksightDataSourceParametersSparkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersSparkOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersSparkOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersSparkOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersSparkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersSparkOutputReference_Override(q QuicksightDataSourceParametersSparkOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersSparkOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSparkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersSqlServer struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersSqlServerOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersSqlServerOutputReference
type jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersSqlServerOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersSqlServerOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersSqlServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersSqlServerOutputReference_Override(q QuicksightDataSourceParametersSqlServerOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersSqlServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersSqlServerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersTeradata struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#database QuicksightDataSource#database}.
	Database *string `json:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#host QuicksightDataSource#host}.
	Host *string `json:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#port QuicksightDataSource#port}.
	Port *float64 `json:"port"`
}

type QuicksightDataSourceParametersTeradataOutputReference interface {
	cdktf.ComplexObject
	Database() *string
	SetDatabase(val *string)
	DatabaseInput() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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

// The jsii proxy struct for QuicksightDataSourceParametersTeradataOutputReference
type jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersTeradataOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersTeradataOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersTeradataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersTeradataOutputReference_Override(q QuicksightDataSourceParametersTeradataOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersTeradataOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) SetDatabase(val *string) {
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTeradataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceParametersTwitter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#max_rows QuicksightDataSource#max_rows}.
	MaxRows *float64 `json:"maxRows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#query QuicksightDataSource#query}.
	Query *string `json:"query"`
}

type QuicksightDataSourceParametersTwitterOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaxRows() *float64
	SetMaxRows(val *float64)
	MaxRowsInput() *float64
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
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

// The jsii proxy struct for QuicksightDataSourceParametersTwitterOutputReference
type jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) MaxRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) MaxRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceParametersTwitterOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceParametersTwitterOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersTwitterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceParametersTwitterOutputReference_Override(q QuicksightDataSourceParametersTwitterOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceParametersTwitterOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) SetMaxRows(val *float64) {
	_jsii_.Set(
		j,
		"maxRows",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) SetQuery(val *string) {
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceParametersTwitterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourcePermission struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#actions QuicksightDataSource#actions}.
	Actions *[]*string `json:"actions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#principal QuicksightDataSource#principal}.
	Principal *string `json:"principal"`
}

type QuicksightDataSourceSslProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#disable_ssl QuicksightDataSource#disable_ssl}.
	DisableSsl interface{} `json:"disableSsl"`
}

type QuicksightDataSourceSslPropertiesOutputReference interface {
	cdktf.ComplexObject
	DisableSsl() interface{}
	SetDisableSsl(val interface{})
	DisableSslInput() interface{}
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
}

// The jsii proxy struct for QuicksightDataSourceSslPropertiesOutputReference
type jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) DisableSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) DisableSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceSslPropertiesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceSslPropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceSslPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceSslPropertiesOutputReference_Override(q QuicksightDataSourceSslPropertiesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceSslPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) SetDisableSsl(val interface{}) {
	_jsii_.Set(
		j,
		"disableSsl",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceSslPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type QuicksightDataSourceVpcConnectionProperties struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_data_source.html#vpc_connection_arn QuicksightDataSource#vpc_connection_arn}.
	VpcConnectionArn *string `json:"vpcConnectionArn"`
}

type QuicksightDataSourceVpcConnectionPropertiesOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	VpcConnectionArn() *string
	SetVpcConnectionArn(val *string)
	VpcConnectionArnInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
}

// The jsii proxy struct for QuicksightDataSourceVpcConnectionPropertiesOutputReference
type jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) VpcConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) VpcConnectionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectionArnInput",
		&returns,
	)
	return returns
}

func NewQuicksightDataSourceVpcConnectionPropertiesOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) QuicksightDataSourceVpcConnectionPropertiesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceVpcConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceVpcConnectionPropertiesOutputReference_Override(q QuicksightDataSourceVpcConnectionPropertiesOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightDataSourceVpcConnectionPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) SetVpcConnectionArn(val *string) {
	_jsii_.Set(
		j,
		"vpcConnectionArn",
		val,
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightDataSourceVpcConnectionPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group.html aws_quicksight_group}.
type QuicksightGroup interface {
	cdktf.TerraformResource
	Arn() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	GroupName() *string
	SetGroupName(val *string)
	GroupNameInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
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
	ResetAwsAccountId()
	ResetDescription()
	ResetNamespace()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for QuicksightGroup
type jsiiProxy_QuicksightGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QuicksightGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) GroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group.html aws_quicksight_group} Resource.
func NewQuicksightGroup(scope constructs.Construct, id *string, config *QuicksightGroupConfig) QuicksightGroup {
	_init_.Initialize()

	j := jsiiProxy_QuicksightGroup{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group.html aws_quicksight_group} Resource.
func NewQuicksightGroup_Override(q QuicksightGroup, scope constructs.Construct, id *string, config *QuicksightGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightGroup",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QuicksightGroup) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroup) SetGroupName(val *string) {
	_jsii_.Set(
		j,
		"groupName",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroup) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroup) SetProvider(val cdktf.TerraformProvider) {
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
func QuicksightGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.QuickSight.QuicksightGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QuicksightGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.QuickSight.QuicksightGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (q *jsiiProxy_QuicksightGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QuicksightGroup) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		q,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		q,
		"resetDescription",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightGroup) ResetNamespace() {
	_jsii_.InvokeVoid(
		q,
		"resetNamespace",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (q *jsiiProxy_QuicksightGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (q *jsiiProxy_QuicksightGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (q *jsiiProxy_QuicksightGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type QuicksightGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group.html#group_name QuicksightGroup#group_name}.
	GroupName *string `json:"groupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group.html#aws_account_id QuicksightGroup#aws_account_id}.
	AwsAccountId *string `json:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group.html#description QuicksightGroup#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group.html#namespace QuicksightGroup#namespace}.
	Namespace *string `json:"namespace"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group_membership.html aws_quicksight_group_membership}.
type QuicksightGroupMembership interface {
	cdktf.TerraformResource
	Arn() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	GroupName() *string
	SetGroupName(val *string)
	GroupNameInput() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MemberName() *string
	SetMemberName(val *string)
	MemberNameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
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
	ResetAwsAccountId()
	ResetNamespace()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for QuicksightGroupMembership
type jsiiProxy_QuicksightGroupMembership struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QuicksightGroupMembership) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) GroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) MemberName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) MemberNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightGroupMembership) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group_membership.html aws_quicksight_group_membership} Resource.
func NewQuicksightGroupMembership(scope constructs.Construct, id *string, config *QuicksightGroupMembershipConfig) QuicksightGroupMembership {
	_init_.Initialize()

	j := jsiiProxy_QuicksightGroupMembership{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightGroupMembership",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group_membership.html aws_quicksight_group_membership} Resource.
func NewQuicksightGroupMembership_Override(q QuicksightGroupMembership, scope constructs.Construct, id *string, config *QuicksightGroupMembershipConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightGroupMembership",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QuicksightGroupMembership) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroupMembership) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroupMembership) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroupMembership) SetGroupName(val *string) {
	_jsii_.Set(
		j,
		"groupName",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroupMembership) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroupMembership) SetMemberName(val *string) {
	_jsii_.Set(
		j,
		"memberName",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroupMembership) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_QuicksightGroupMembership) SetProvider(val cdktf.TerraformProvider) {
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
func QuicksightGroupMembership_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.QuickSight.QuicksightGroupMembership",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QuicksightGroupMembership_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.QuickSight.QuicksightGroupMembership",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QuicksightGroupMembership) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		q,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightGroupMembership) ResetNamespace() {
	_jsii_.InvokeVoid(
		q,
		"resetNamespace",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightGroupMembership) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (q *jsiiProxy_QuicksightGroupMembership) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (q *jsiiProxy_QuicksightGroupMembership) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type QuicksightGroupMembershipConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group_membership.html#group_name QuicksightGroupMembership#group_name}.
	GroupName *string `json:"groupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group_membership.html#member_name QuicksightGroupMembership#member_name}.
	MemberName *string `json:"memberName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group_membership.html#aws_account_id QuicksightGroupMembership#aws_account_id}.
	AwsAccountId *string `json:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_group_membership.html#namespace QuicksightGroupMembership#namespace}.
	Namespace *string `json:"namespace"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html aws_quicksight_user}.
type QuicksightUser interface {
	cdktf.TerraformResource
	Arn() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	IamArn() *string
	SetIamArn(val *string)
	IamArnInput() *string
	Id() *string
	IdentityType() *string
	SetIdentityType(val *string)
	IdentityTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SessionName() *string
	SetSessionName(val *string)
	SessionNameInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
	UserRole() *string
	SetUserRole(val *string)
	UserRoleInput() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetAwsAccountId()
	ResetIamArn()
	ResetNamespace()
	ResetOverrideLogicalId()
	ResetSessionName()
	ResetUserName()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for QuicksightUser
type jsiiProxy_QuicksightUser struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_QuicksightUser) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) IamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) IamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) SessionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) SessionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) UserRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightUser) UserRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html aws_quicksight_user} Resource.
func NewQuicksightUser(scope constructs.Construct, id *string, config *QuicksightUserConfig) QuicksightUser {
	_init_.Initialize()

	j := jsiiProxy_QuicksightUser{}

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html aws_quicksight_user} Resource.
func NewQuicksightUser_Override(q QuicksightUser, scope constructs.Construct, id *string, config *QuicksightUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.QuickSight.QuicksightUser",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QuicksightUser) SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetEmail(val *string) {
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetIamArn(val *string) {
	_jsii_.Set(
		j,
		"iamArn",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetIdentityType(val *string) {
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetSessionName(val *string) {
	_jsii_.Set(
		j,
		"sessionName",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetUserName(val *string) {
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (j *jsiiProxy_QuicksightUser) SetUserRole(val *string) {
	_jsii_.Set(
		j,
		"userRole",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func QuicksightUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.QuickSight.QuicksightUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QuicksightUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.QuickSight.QuicksightUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightUser) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (q *jsiiProxy_QuicksightUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightUser) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightUser) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightUser) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (q *jsiiProxy_QuicksightUser) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QuicksightUser) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		q,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightUser) ResetIamArn() {
	_jsii_.InvokeVoid(
		q,
		"resetIamArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightUser) ResetNamespace() {
	_jsii_.InvokeVoid(
		q,
		"resetNamespace",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (q *jsiiProxy_QuicksightUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightUser) ResetSessionName() {
	_jsii_.InvokeVoid(
		q,
		"resetSessionName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightUser) ResetUserName() {
	_jsii_.InvokeVoid(
		q,
		"resetUserName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (q *jsiiProxy_QuicksightUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (q *jsiiProxy_QuicksightUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (q *jsiiProxy_QuicksightUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type QuicksightUserConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html#email QuicksightUser#email}.
	Email *string `json:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html#identity_type QuicksightUser#identity_type}.
	IdentityType *string `json:"identityType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html#user_role QuicksightUser#user_role}.
	UserRole *string `json:"userRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html#aws_account_id QuicksightUser#aws_account_id}.
	AwsAccountId *string `json:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html#iam_arn QuicksightUser#iam_arn}.
	IamArn *string `json:"iamArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html#namespace QuicksightUser#namespace}.
	Namespace *string `json:"namespace"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html#session_name QuicksightUser#session_name}.
	SessionName *string `json:"sessionName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_user.html#user_name QuicksightUser#user_name}.
	UserName *string `json:"userName"`
}
