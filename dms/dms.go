package dms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/dms/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/dms_certificate.html aws_dms_certificate}.
type DmsCertificate interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	CertificateId() *string
	SetCertificateId(val *string)
	CertificateIdInput() *string
	CertificatePem() *string
	SetCertificatePem(val *string)
	CertificatePemInput() *string
	CertificateWallet() *string
	SetCertificateWallet(val *string)
	CertificateWalletInput() *string
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
	ResetCertificatePem()
	ResetCertificateWallet()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DmsCertificate
type jsiiProxy_DmsCertificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) CertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) CertificateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) CertificatePem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) CertificatePemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) CertificateWallet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateWallet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) CertificateWalletInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateWalletInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_certificate.html aws_dms_certificate} Resource.
func NewDmsCertificate(scope constructs.Construct, id *string, config *DmsCertificateConfig) DmsCertificate {
	_init_.Initialize()

	j := jsiiProxy_DmsCertificate{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_certificate.html aws_dms_certificate} Resource.
func NewDmsCertificate_Override(d DmsCertificate, scope constructs.Construct, id *string, config *DmsCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsCertificate",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsCertificate) SetCertificateId(val *string) {
	_jsii_.Set(
		j,
		"certificateId",
		val,
	)
}

func (j *jsiiProxy_DmsCertificate) SetCertificatePem(val *string) {
	_jsii_.Set(
		j,
		"certificatePem",
		val,
	)
}

func (j *jsiiProxy_DmsCertificate) SetCertificateWallet(val *string) {
	_jsii_.Set(
		j,
		"certificateWallet",
		val,
	)
}

func (j *jsiiProxy_DmsCertificate) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsCertificate) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsCertificate) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsCertificate) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsCertificate) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsCertificate) SetTagsAll(val interface{}) {
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
func DmsCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DMS.DmsCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DMS.DmsCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DmsCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DmsCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsCertificate) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsCertificate) ResetCertificatePem() {
	_jsii_.InvokeVoid(
		d,
		"resetCertificatePem",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsCertificate) ResetCertificateWallet() {
	_jsii_.InvokeVoid(
		d,
		"resetCertificateWallet",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DmsCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsCertificate) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsCertificate) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsCertificate) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DmsCertificate) ToMetadata() interface{} {
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
func (d *jsiiProxy_DmsCertificate) ToString() *string {
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
func (d *jsiiProxy_DmsCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DmsCertificateConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_certificate.html#certificate_id DmsCertificate#certificate_id}.
	CertificateId *string `json:"certificateId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_certificate.html#certificate_pem DmsCertificate#certificate_pem}.
	CertificatePem *string `json:"certificatePem"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_certificate.html#certificate_wallet DmsCertificate#certificate_wallet}.
	CertificateWallet *string `json:"certificateWallet"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_certificate.html#tags DmsCertificate#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_certificate.html#tags_all DmsCertificate#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html aws_dms_endpoint}.
type DmsEndpoint interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ElasticsearchSettings() DmsEndpointElasticsearchSettingsOutputReference
	ElasticsearchSettingsInput() *DmsEndpointElasticsearchSettings
	EndpointArn() *string
	EndpointId() *string
	SetEndpointId(val *string)
	EndpointIdInput() *string
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	EngineName() *string
	SetEngineName(val *string)
	EngineNameInput() *string
	ExtraConnectionAttributes() *string
	SetExtraConnectionAttributes(val *string)
	ExtraConnectionAttributesInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KafkaSettings() DmsEndpointKafkaSettingsOutputReference
	KafkaSettingsInput() *DmsEndpointKafkaSettings
	KinesisSettings() DmsEndpointKinesisSettingsOutputReference
	KinesisSettingsInput() *DmsEndpointKinesisSettings
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MongodbSettings() DmsEndpointMongodbSettingsOutputReference
	MongodbSettingsInput() *DmsEndpointMongodbSettings
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	S3Settings() DmsEndpointS3SettingsOutputReference
	S3SettingsInput() *DmsEndpointS3Settings
	ServerName() *string
	SetServerName(val *string)
	ServerNameInput() *string
	ServiceAccessRole() *string
	SetServiceAccessRole(val *string)
	ServiceAccessRoleInput() *string
	SslMode() *string
	SetSslMode(val *string)
	SslModeInput() *string
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
	PutElasticsearchSettings(value *DmsEndpointElasticsearchSettings)
	PutKafkaSettings(value *DmsEndpointKafkaSettings)
	PutKinesisSettings(value *DmsEndpointKinesisSettings)
	PutMongodbSettings(value *DmsEndpointMongodbSettings)
	PutS3Settings(value *DmsEndpointS3Settings)
	ResetCertificateArn()
	ResetDatabaseName()
	ResetElasticsearchSettings()
	ResetExtraConnectionAttributes()
	ResetKafkaSettings()
	ResetKinesisSettings()
	ResetKmsKeyArn()
	ResetMongodbSettings()
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPort()
	ResetS3Settings()
	ResetServerName()
	ResetServiceAccessRole()
	ResetSslMode()
	ResetTags()
	ResetTagsAll()
	ResetUsername()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DmsEndpoint
type jsiiProxy_DmsEndpoint struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsEndpoint) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ElasticsearchSettings() DmsEndpointElasticsearchSettingsOutputReference {
	var returns DmsEndpointElasticsearchSettingsOutputReference
	_jsii_.Get(
		j,
		"elasticsearchSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ElasticsearchSettingsInput() *DmsEndpointElasticsearchSettings {
	var returns *DmsEndpointElasticsearchSettings
	_jsii_.Get(
		j,
		"elasticsearchSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EngineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EngineNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ExtraConnectionAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ExtraConnectionAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KafkaSettings() DmsEndpointKafkaSettingsOutputReference {
	var returns DmsEndpointKafkaSettingsOutputReference
	_jsii_.Get(
		j,
		"kafkaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KafkaSettingsInput() *DmsEndpointKafkaSettings {
	var returns *DmsEndpointKafkaSettings
	_jsii_.Get(
		j,
		"kafkaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KinesisSettings() DmsEndpointKinesisSettingsOutputReference {
	var returns DmsEndpointKinesisSettingsOutputReference
	_jsii_.Get(
		j,
		"kinesisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KinesisSettingsInput() *DmsEndpointKinesisSettings {
	var returns *DmsEndpointKinesisSettings
	_jsii_.Get(
		j,
		"kinesisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MongodbSettings() DmsEndpointMongodbSettingsOutputReference {
	var returns DmsEndpointMongodbSettingsOutputReference
	_jsii_.Get(
		j,
		"mongodbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MongodbSettingsInput() *DmsEndpointMongodbSettings {
	var returns *DmsEndpointMongodbSettings
	_jsii_.Get(
		j,
		"mongodbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) S3Settings() DmsEndpointS3SettingsOutputReference {
	var returns DmsEndpointS3SettingsOutputReference
	_jsii_.Get(
		j,
		"s3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) S3SettingsInput() *DmsEndpointS3Settings {
	var returns *DmsEndpointS3Settings
	_jsii_.Get(
		j,
		"s3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServiceAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServiceAccessRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html aws_dms_endpoint} Resource.
func NewDmsEndpoint(scope constructs.Construct, id *string, config *DmsEndpointConfig) DmsEndpoint {
	_init_.Initialize()

	j := jsiiProxy_DmsEndpoint{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html aws_dms_endpoint} Resource.
func NewDmsEndpoint_Override(d DmsEndpoint, scope constructs.Construct, id *string, config *DmsEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpoint",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetEndpointId(val *string) {
	_jsii_.Set(
		j,
		"endpointId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetEndpointType(val *string) {
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetEngineName(val *string) {
	_jsii_.Set(
		j,
		"engineName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetExtraConnectionAttributes(val *string) {
	_jsii_.Set(
		j,
		"extraConnectionAttributes",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetServerName(val *string) {
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetServiceAccessRole(val *string) {
	_jsii_.Set(
		j,
		"serviceAccessRole",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetSslMode(val *string) {
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint) SetUsername(val *string) {
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
func DmsEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DMS.DmsEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DMS.DmsEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DmsEndpoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DmsEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsEndpoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutElasticsearchSettings(value *DmsEndpointElasticsearchSettings) {
	_jsii_.InvokeVoid(
		d,
		"putElasticsearchSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutKafkaSettings(value *DmsEndpointKafkaSettings) {
	_jsii_.InvokeVoid(
		d,
		"putKafkaSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutKinesisSettings(value *DmsEndpointKinesisSettings) {
	_jsii_.InvokeVoid(
		d,
		"putKinesisSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutMongodbSettings(value *DmsEndpointMongodbSettings) {
	_jsii_.InvokeVoid(
		d,
		"putMongodbSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutS3Settings(value *DmsEndpointS3Settings) {
	_jsii_.InvokeVoid(
		d,
		"putS3Settings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetElasticsearchSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetElasticsearchSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetExtraConnectionAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetExtraConnectionAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetKafkaSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetKafkaSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetKinesisSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetKinesisSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetMongodbSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMongodbSettings",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DmsEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetS3Settings() {
	_jsii_.InvokeVoid(
		d,
		"resetS3Settings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetServerName() {
	_jsii_.InvokeVoid(
		d,
		"resetServerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetServiceAccessRole() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccessRole",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetSslMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSslMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DmsEndpoint) ToMetadata() interface{} {
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
func (d *jsiiProxy_DmsEndpoint) ToString() *string {
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
func (d *jsiiProxy_DmsEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DmsEndpointConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#endpoint_id DmsEndpoint#endpoint_id}.
	EndpointId *string `json:"endpointId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#endpoint_type DmsEndpoint#endpoint_type}.
	EndpointType *string `json:"endpointType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#engine_name DmsEndpoint#engine_name}.
	EngineName *string `json:"engineName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#certificate_arn DmsEndpoint#certificate_arn}.
	CertificateArn *string `json:"certificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#database_name DmsEndpoint#database_name}.
	DatabaseName *string `json:"databaseName"`
	// elasticsearch_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#elasticsearch_settings DmsEndpoint#elasticsearch_settings}
	ElasticsearchSettings *DmsEndpointElasticsearchSettings `json:"elasticsearchSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#extra_connection_attributes DmsEndpoint#extra_connection_attributes}.
	ExtraConnectionAttributes *string `json:"extraConnectionAttributes"`
	// kafka_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#kafka_settings DmsEndpoint#kafka_settings}
	KafkaSettings *DmsEndpointKafkaSettings `json:"kafkaSettings"`
	// kinesis_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#kinesis_settings DmsEndpoint#kinesis_settings}
	KinesisSettings *DmsEndpointKinesisSettings `json:"kinesisSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#kms_key_arn DmsEndpoint#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn"`
	// mongodb_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#mongodb_settings DmsEndpoint#mongodb_settings}
	MongodbSettings *DmsEndpointMongodbSettings `json:"mongodbSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#password DmsEndpoint#password}.
	Password *string `json:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#port DmsEndpoint#port}.
	Port *float64 `json:"port"`
	// s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#s3_settings DmsEndpoint#s3_settings}
	S3Settings *DmsEndpointS3Settings `json:"s3Settings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#server_name DmsEndpoint#server_name}.
	ServerName *string `json:"serverName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#service_access_role DmsEndpoint#service_access_role}.
	ServiceAccessRole *string `json:"serviceAccessRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#ssl_mode DmsEndpoint#ssl_mode}.
	SslMode *string `json:"sslMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#tags DmsEndpoint#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#tags_all DmsEndpoint#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#username DmsEndpoint#username}.
	Username *string `json:"username"`
}

type DmsEndpointElasticsearchSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#endpoint_uri DmsEndpoint#endpoint_uri}.
	EndpointUri *string `json:"endpointUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#service_access_role_arn DmsEndpoint#service_access_role_arn}.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#error_retry_duration DmsEndpoint#error_retry_duration}.
	ErrorRetryDuration *float64 `json:"errorRetryDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#full_load_error_percentage DmsEndpoint#full_load_error_percentage}.
	FullLoadErrorPercentage *float64 `json:"fullLoadErrorPercentage"`
}

type DmsEndpointElasticsearchSettingsOutputReference interface {
	cdktf.ComplexObject
	EndpointUri() *string
	SetEndpointUri(val *string)
	EndpointUriInput() *string
	ErrorRetryDuration() *float64
	SetErrorRetryDuration(val *float64)
	ErrorRetryDurationInput() *float64
	FullLoadErrorPercentage() *float64
	SetFullLoadErrorPercentage(val *float64)
	FullLoadErrorPercentageInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ServiceAccessRoleArn() *string
	SetServiceAccessRoleArn(val *string)
	ServiceAccessRoleArnInput() *string
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
	ResetErrorRetryDuration()
	ResetFullLoadErrorPercentage()
}

// The jsii proxy struct for DmsEndpointElasticsearchSettingsOutputReference
type jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) EndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) EndpointUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) ErrorRetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorRetryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) ErrorRetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorRetryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) FullLoadErrorPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullLoadErrorPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) FullLoadErrorPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullLoadErrorPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDmsEndpointElasticsearchSettingsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) DmsEndpointElasticsearchSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointElasticsearchSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDmsEndpointElasticsearchSettingsOutputReference_Override(d DmsEndpointElasticsearchSettingsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointElasticsearchSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) SetEndpointUri(val *string) {
	_jsii_.Set(
		j,
		"endpointUri",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) SetErrorRetryDuration(val *float64) {
	_jsii_.Set(
		j,
		"errorRetryDuration",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) SetFullLoadErrorPercentage(val *float64) {
	_jsii_.Set(
		j,
		"fullLoadErrorPercentage",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) SetServiceAccessRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) ResetErrorRetryDuration() {
	_jsii_.InvokeVoid(
		d,
		"resetErrorRetryDuration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointElasticsearchSettingsOutputReference) ResetFullLoadErrorPercentage() {
	_jsii_.InvokeVoid(
		d,
		"resetFullLoadErrorPercentage",
		nil, // no parameters
	)
}

type DmsEndpointKafkaSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#broker DmsEndpoint#broker}.
	Broker *string `json:"broker"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_control_details DmsEndpoint#include_control_details}.
	IncludeControlDetails interface{} `json:"includeControlDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_null_and_empty DmsEndpoint#include_null_and_empty}.
	IncludeNullAndEmpty interface{} `json:"includeNullAndEmpty"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_partition_value DmsEndpoint#include_partition_value}.
	IncludePartitionValue interface{} `json:"includePartitionValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_table_alter_operations DmsEndpoint#include_table_alter_operations}.
	IncludeTableAlterOperations interface{} `json:"includeTableAlterOperations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_transaction_details DmsEndpoint#include_transaction_details}.
	IncludeTransactionDetails interface{} `json:"includeTransactionDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#message_format DmsEndpoint#message_format}.
	MessageFormat *string `json:"messageFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#message_max_bytes DmsEndpoint#message_max_bytes}.
	MessageMaxBytes *float64 `json:"messageMaxBytes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#no_hex_prefix DmsEndpoint#no_hex_prefix}.
	NoHexPrefix interface{} `json:"noHexPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#partition_include_schema_table DmsEndpoint#partition_include_schema_table}.
	PartitionIncludeSchemaTable interface{} `json:"partitionIncludeSchemaTable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#sasl_password DmsEndpoint#sasl_password}.
	SaslPassword *string `json:"saslPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#sasl_username DmsEndpoint#sasl_username}.
	SaslUsername *string `json:"saslUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#security_protocol DmsEndpoint#security_protocol}.
	SecurityProtocol *string `json:"securityProtocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#ssl_ca_certificate_arn DmsEndpoint#ssl_ca_certificate_arn}.
	SslCaCertificateArn *string `json:"sslCaCertificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#ssl_client_certificate_arn DmsEndpoint#ssl_client_certificate_arn}.
	SslClientCertificateArn *string `json:"sslClientCertificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#ssl_client_key_arn DmsEndpoint#ssl_client_key_arn}.
	SslClientKeyArn *string `json:"sslClientKeyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#ssl_client_key_password DmsEndpoint#ssl_client_key_password}.
	SslClientKeyPassword *string `json:"sslClientKeyPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#topic DmsEndpoint#topic}.
	Topic *string `json:"topic"`
}

type DmsEndpointKafkaSettingsOutputReference interface {
	cdktf.ComplexObject
	Broker() *string
	SetBroker(val *string)
	BrokerInput() *string
	IncludeControlDetails() interface{}
	SetIncludeControlDetails(val interface{})
	IncludeControlDetailsInput() interface{}
	IncludeNullAndEmpty() interface{}
	SetIncludeNullAndEmpty(val interface{})
	IncludeNullAndEmptyInput() interface{}
	IncludePartitionValue() interface{}
	SetIncludePartitionValue(val interface{})
	IncludePartitionValueInput() interface{}
	IncludeTableAlterOperations() interface{}
	SetIncludeTableAlterOperations(val interface{})
	IncludeTableAlterOperationsInput() interface{}
	IncludeTransactionDetails() interface{}
	SetIncludeTransactionDetails(val interface{})
	IncludeTransactionDetailsInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MessageFormat() *string
	SetMessageFormat(val *string)
	MessageFormatInput() *string
	MessageMaxBytes() *float64
	SetMessageMaxBytes(val *float64)
	MessageMaxBytesInput() *float64
	NoHexPrefix() interface{}
	SetNoHexPrefix(val interface{})
	NoHexPrefixInput() interface{}
	PartitionIncludeSchemaTable() interface{}
	SetPartitionIncludeSchemaTable(val interface{})
	PartitionIncludeSchemaTableInput() interface{}
	SaslPassword() *string
	SetSaslPassword(val *string)
	SaslPasswordInput() *string
	SaslUsername() *string
	SetSaslUsername(val *string)
	SaslUsernameInput() *string
	SecurityProtocol() *string
	SetSecurityProtocol(val *string)
	SecurityProtocolInput() *string
	SslCaCertificateArn() *string
	SetSslCaCertificateArn(val *string)
	SslCaCertificateArnInput() *string
	SslClientCertificateArn() *string
	SetSslClientCertificateArn(val *string)
	SslClientCertificateArnInput() *string
	SslClientKeyArn() *string
	SetSslClientKeyArn(val *string)
	SslClientKeyArnInput() *string
	SslClientKeyPassword() *string
	SetSslClientKeyPassword(val *string)
	SslClientKeyPasswordInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIncludeControlDetails()
	ResetIncludeNullAndEmpty()
	ResetIncludePartitionValue()
	ResetIncludeTableAlterOperations()
	ResetIncludeTransactionDetails()
	ResetMessageFormat()
	ResetMessageMaxBytes()
	ResetNoHexPrefix()
	ResetPartitionIncludeSchemaTable()
	ResetSaslPassword()
	ResetSaslUsername()
	ResetSecurityProtocol()
	ResetSslCaCertificateArn()
	ResetSslClientCertificateArn()
	ResetSslClientKeyArn()
	ResetSslClientKeyPassword()
	ResetTopic()
}

// The jsii proxy struct for DmsEndpointKafkaSettingsOutputReference
type jsiiProxy_DmsEndpointKafkaSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) Broker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"broker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) BrokerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeControlDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeControlDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeNullAndEmpty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeNullAndEmptyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmptyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludePartitionValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludePartitionValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeTableAlterOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeTableAlterOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeTransactionDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IncludeTransactionDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) MessageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) MessageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) MessageMaxBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) MessageMaxBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageMaxBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) NoHexPrefix() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHexPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) NoHexPrefixInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noHexPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) PartitionIncludeSchemaTable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) PartitionIncludeSchemaTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SaslUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslCaCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslCaCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientKeyPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SslClientKeyPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslClientKeyPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}

func NewDmsEndpointKafkaSettingsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) DmsEndpointKafkaSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DmsEndpointKafkaSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointKafkaSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDmsEndpointKafkaSettingsOutputReference_Override(d DmsEndpointKafkaSettingsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointKafkaSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetBroker(val *string) {
	_jsii_.Set(
		j,
		"broker",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetIncludeControlDetails(val interface{}) {
	_jsii_.Set(
		j,
		"includeControlDetails",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetIncludeNullAndEmpty(val interface{}) {
	_jsii_.Set(
		j,
		"includeNullAndEmpty",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetIncludePartitionValue(val interface{}) {
	_jsii_.Set(
		j,
		"includePartitionValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetIncludeTableAlterOperations(val interface{}) {
	_jsii_.Set(
		j,
		"includeTableAlterOperations",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetIncludeTransactionDetails(val interface{}) {
	_jsii_.Set(
		j,
		"includeTransactionDetails",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetMessageFormat(val *string) {
	_jsii_.Set(
		j,
		"messageFormat",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetMessageMaxBytes(val *float64) {
	_jsii_.Set(
		j,
		"messageMaxBytes",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetNoHexPrefix(val interface{}) {
	_jsii_.Set(
		j,
		"noHexPrefix",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetPartitionIncludeSchemaTable(val interface{}) {
	_jsii_.Set(
		j,
		"partitionIncludeSchemaTable",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetSaslPassword(val *string) {
	_jsii_.Set(
		j,
		"saslPassword",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetSaslUsername(val *string) {
	_jsii_.Set(
		j,
		"saslUsername",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetSecurityProtocol(val *string) {
	_jsii_.Set(
		j,
		"securityProtocol",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetSslCaCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"sslCaCertificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetSslClientCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"sslClientCertificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetSslClientKeyArn(val *string) {
	_jsii_.Set(
		j,
		"sslClientKeyArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetSslClientKeyPassword(val *string) {
	_jsii_.Set(
		j,
		"sslClientKeyPassword",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) SetTopic(val *string) {
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludeControlDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeControlDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludeNullAndEmpty() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeNullAndEmpty",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludePartitionValue() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludePartitionValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludeTableAlterOperations() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeTableAlterOperations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetIncludeTransactionDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeTransactionDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetMessageFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetMessageMaxBytes() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageMaxBytes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetNoHexPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetNoHexPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetPartitionIncludeSchemaTable() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionIncludeSchemaTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSaslPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetSaslPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSaslUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetSaslUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSecurityProtocol() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityProtocol",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSslCaCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSslCaCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSslClientCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSslClientCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSslClientKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSslClientKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetSslClientKeyPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetSslClientKeyPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKafkaSettingsOutputReference) ResetTopic() {
	_jsii_.InvokeVoid(
		d,
		"resetTopic",
		nil, // no parameters
	)
}

type DmsEndpointKinesisSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_control_details DmsEndpoint#include_control_details}.
	IncludeControlDetails interface{} `json:"includeControlDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_null_and_empty DmsEndpoint#include_null_and_empty}.
	IncludeNullAndEmpty interface{} `json:"includeNullAndEmpty"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_partition_value DmsEndpoint#include_partition_value}.
	IncludePartitionValue interface{} `json:"includePartitionValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_table_alter_operations DmsEndpoint#include_table_alter_operations}.
	IncludeTableAlterOperations interface{} `json:"includeTableAlterOperations"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#include_transaction_details DmsEndpoint#include_transaction_details}.
	IncludeTransactionDetails interface{} `json:"includeTransactionDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#message_format DmsEndpoint#message_format}.
	MessageFormat *string `json:"messageFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#partition_include_schema_table DmsEndpoint#partition_include_schema_table}.
	PartitionIncludeSchemaTable interface{} `json:"partitionIncludeSchemaTable"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#service_access_role_arn DmsEndpoint#service_access_role_arn}.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#stream_arn DmsEndpoint#stream_arn}.
	StreamArn *string `json:"streamArn"`
}

type DmsEndpointKinesisSettingsOutputReference interface {
	cdktf.ComplexObject
	IncludeControlDetails() interface{}
	SetIncludeControlDetails(val interface{})
	IncludeControlDetailsInput() interface{}
	IncludeNullAndEmpty() interface{}
	SetIncludeNullAndEmpty(val interface{})
	IncludeNullAndEmptyInput() interface{}
	IncludePartitionValue() interface{}
	SetIncludePartitionValue(val interface{})
	IncludePartitionValueInput() interface{}
	IncludeTableAlterOperations() interface{}
	SetIncludeTableAlterOperations(val interface{})
	IncludeTableAlterOperationsInput() interface{}
	IncludeTransactionDetails() interface{}
	SetIncludeTransactionDetails(val interface{})
	IncludeTransactionDetailsInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MessageFormat() *string
	SetMessageFormat(val *string)
	MessageFormatInput() *string
	PartitionIncludeSchemaTable() interface{}
	SetPartitionIncludeSchemaTable(val interface{})
	PartitionIncludeSchemaTableInput() interface{}
	ServiceAccessRoleArn() *string
	SetServiceAccessRoleArn(val *string)
	ServiceAccessRoleArnInput() *string
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
	ResetIncludeControlDetails()
	ResetIncludeNullAndEmpty()
	ResetIncludePartitionValue()
	ResetIncludeTableAlterOperations()
	ResetIncludeTransactionDetails()
	ResetMessageFormat()
	ResetPartitionIncludeSchemaTable()
	ResetServiceAccessRoleArn()
	ResetStreamArn()
}

// The jsii proxy struct for DmsEndpointKinesisSettingsOutputReference
type jsiiProxy_DmsEndpointKinesisSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludeControlDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludeControlDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeControlDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludeNullAndEmpty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludeNullAndEmptyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNullAndEmptyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludePartitionValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludePartitionValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includePartitionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludeTableAlterOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludeTableAlterOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTableAlterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludeTransactionDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IncludeTransactionDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTransactionDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) MessageFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) MessageFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) PartitionIncludeSchemaTable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) PartitionIncludeSchemaTableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIncludeSchemaTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) StreamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDmsEndpointKinesisSettingsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) DmsEndpointKinesisSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DmsEndpointKinesisSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointKinesisSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDmsEndpointKinesisSettingsOutputReference_Override(d DmsEndpointKinesisSettingsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointKinesisSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetIncludeControlDetails(val interface{}) {
	_jsii_.Set(
		j,
		"includeControlDetails",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetIncludeNullAndEmpty(val interface{}) {
	_jsii_.Set(
		j,
		"includeNullAndEmpty",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetIncludePartitionValue(val interface{}) {
	_jsii_.Set(
		j,
		"includePartitionValue",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetIncludeTableAlterOperations(val interface{}) {
	_jsii_.Set(
		j,
		"includeTableAlterOperations",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetIncludeTransactionDetails(val interface{}) {
	_jsii_.Set(
		j,
		"includeTransactionDetails",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetMessageFormat(val *string) {
	_jsii_.Set(
		j,
		"messageFormat",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetPartitionIncludeSchemaTable(val interface{}) {
	_jsii_.Set(
		j,
		"partitionIncludeSchemaTable",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetServiceAccessRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetStreamArn(val *string) {
	_jsii_.Set(
		j,
		"streamArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ResetIncludeControlDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeControlDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ResetIncludeNullAndEmpty() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeNullAndEmpty",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ResetIncludePartitionValue() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludePartitionValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ResetIncludeTableAlterOperations() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeTableAlterOperations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ResetIncludeTransactionDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeTransactionDetails",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ResetMessageFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetMessageFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ResetPartitionIncludeSchemaTable() {
	_jsii_.InvokeVoid(
		d,
		"resetPartitionIncludeSchemaTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointKinesisSettingsOutputReference) ResetStreamArn() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamArn",
		nil, // no parameters
	)
}

type DmsEndpointMongodbSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#auth_mechanism DmsEndpoint#auth_mechanism}.
	AuthMechanism *string `json:"authMechanism"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#auth_source DmsEndpoint#auth_source}.
	AuthSource *string `json:"authSource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#auth_type DmsEndpoint#auth_type}.
	AuthType *string `json:"authType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#docs_to_investigate DmsEndpoint#docs_to_investigate}.
	DocsToInvestigate *string `json:"docsToInvestigate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#extract_doc_id DmsEndpoint#extract_doc_id}.
	ExtractDocId *string `json:"extractDocId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#nesting_level DmsEndpoint#nesting_level}.
	NestingLevel *string `json:"nestingLevel"`
}

type DmsEndpointMongodbSettingsOutputReference interface {
	cdktf.ComplexObject
	AuthMechanism() *string
	SetAuthMechanism(val *string)
	AuthMechanismInput() *string
	AuthSource() *string
	SetAuthSource(val *string)
	AuthSourceInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	DocsToInvestigate() *string
	SetDocsToInvestigate(val *string)
	DocsToInvestigateInput() *string
	ExtractDocId() *string
	SetExtractDocId(val *string)
	ExtractDocIdInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	NestingLevel() *string
	SetNestingLevel(val *string)
	NestingLevelInput() *string
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
	ResetAuthMechanism()
	ResetAuthSource()
	ResetAuthType()
	ResetDocsToInvestigate()
	ResetExtractDocId()
	ResetNestingLevel()
}

// The jsii proxy struct for DmsEndpointMongodbSettingsOutputReference
type jsiiProxy_DmsEndpointMongodbSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthMechanism() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMechanism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthMechanismInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMechanismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) DocsToInvestigate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsToInvestigate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) DocsToInvestigateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"docsToInvestigateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ExtractDocId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractDocId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ExtractDocIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extractDocIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) NestingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nestingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) NestingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nestingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDmsEndpointMongodbSettingsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) DmsEndpointMongodbSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DmsEndpointMongodbSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointMongodbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDmsEndpointMongodbSettingsOutputReference_Override(d DmsEndpointMongodbSettingsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointMongodbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) SetAuthMechanism(val *string) {
	_jsii_.Set(
		j,
		"authMechanism",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) SetAuthSource(val *string) {
	_jsii_.Set(
		j,
		"authSource",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) SetAuthType(val *string) {
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) SetDocsToInvestigate(val *string) {
	_jsii_.Set(
		j,
		"docsToInvestigate",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) SetExtractDocId(val *string) {
	_jsii_.Set(
		j,
		"extractDocId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) SetNestingLevel(val *string) {
	_jsii_.Set(
		j,
		"nestingLevel",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetAuthMechanism() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthMechanism",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetAuthSource() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthSource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetAuthType() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetDocsToInvestigate() {
	_jsii_.InvokeVoid(
		d,
		"resetDocsToInvestigate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetExtractDocId() {
	_jsii_.InvokeVoid(
		d,
		"resetExtractDocId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointMongodbSettingsOutputReference) ResetNestingLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetNestingLevel",
		nil, // no parameters
	)
}

type DmsEndpointS3Settings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#bucket_folder DmsEndpoint#bucket_folder}.
	BucketFolder *string `json:"bucketFolder"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#bucket_name DmsEndpoint#bucket_name}.
	BucketName *string `json:"bucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#compression_type DmsEndpoint#compression_type}.
	CompressionType *string `json:"compressionType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#csv_delimiter DmsEndpoint#csv_delimiter}.
	CsvDelimiter *string `json:"csvDelimiter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#csv_row_delimiter DmsEndpoint#csv_row_delimiter}.
	CsvRowDelimiter *string `json:"csvRowDelimiter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#data_format DmsEndpoint#data_format}.
	DataFormat *string `json:"dataFormat"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#date_partition_enabled DmsEndpoint#date_partition_enabled}.
	DatePartitionEnabled interface{} `json:"datePartitionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#encryption_mode DmsEndpoint#encryption_mode}.
	EncryptionMode *string `json:"encryptionMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#external_table_definition DmsEndpoint#external_table_definition}.
	ExternalTableDefinition *string `json:"externalTableDefinition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#parquet_timestamp_in_millisecond DmsEndpoint#parquet_timestamp_in_millisecond}.
	ParquetTimestampInMillisecond interface{} `json:"parquetTimestampInMillisecond"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#parquet_version DmsEndpoint#parquet_version}.
	ParquetVersion *string `json:"parquetVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#server_side_encryption_kms_key_id DmsEndpoint#server_side_encryption_kms_key_id}.
	ServerSideEncryptionKmsKeyId *string `json:"serverSideEncryptionKmsKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_endpoint.html#service_access_role_arn DmsEndpoint#service_access_role_arn}.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn"`
}

type DmsEndpointS3SettingsOutputReference interface {
	cdktf.ComplexObject
	BucketFolder() *string
	SetBucketFolder(val *string)
	BucketFolderInput() *string
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	CompressionType() *string
	SetCompressionType(val *string)
	CompressionTypeInput() *string
	CsvDelimiter() *string
	SetCsvDelimiter(val *string)
	CsvDelimiterInput() *string
	CsvRowDelimiter() *string
	SetCsvRowDelimiter(val *string)
	CsvRowDelimiterInput() *string
	DataFormat() *string
	SetDataFormat(val *string)
	DataFormatInput() *string
	DatePartitionEnabled() interface{}
	SetDatePartitionEnabled(val interface{})
	DatePartitionEnabledInput() interface{}
	EncryptionMode() *string
	SetEncryptionMode(val *string)
	EncryptionModeInput() *string
	ExternalTableDefinition() *string
	SetExternalTableDefinition(val *string)
	ExternalTableDefinitionInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	ParquetTimestampInMillisecond() interface{}
	SetParquetTimestampInMillisecond(val interface{})
	ParquetTimestampInMillisecondInput() interface{}
	ParquetVersion() *string
	SetParquetVersion(val *string)
	ParquetVersionInput() *string
	ServerSideEncryptionKmsKeyId() *string
	SetServerSideEncryptionKmsKeyId(val *string)
	ServerSideEncryptionKmsKeyIdInput() *string
	ServiceAccessRoleArn() *string
	SetServiceAccessRoleArn(val *string)
	ServiceAccessRoleArnInput() *string
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
	ResetBucketFolder()
	ResetBucketName()
	ResetCompressionType()
	ResetCsvDelimiter()
	ResetCsvRowDelimiter()
	ResetDataFormat()
	ResetDatePartitionEnabled()
	ResetEncryptionMode()
	ResetExternalTableDefinition()
	ResetParquetTimestampInMillisecond()
	ResetParquetVersion()
	ResetServerSideEncryptionKmsKeyId()
	ResetServiceAccessRoleArn()
}

// The jsii proxy struct for DmsEndpointS3SettingsOutputReference
type jsiiProxy_DmsEndpointS3SettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) BucketFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) BucketFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvRowDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) CsvRowDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DatePartitionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datePartitionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) DatePartitionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datePartitionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) EncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ExternalTableDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ExternalTableDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ParquetTimestampInMillisecond() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ParquetTimestampInMillisecondInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ParquetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ParquetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ServerSideEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ServerSideEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDmsEndpointS3SettingsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) DmsEndpointS3SettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DmsEndpointS3SettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointS3SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDmsEndpointS3SettingsOutputReference_Override(d DmsEndpointS3SettingsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEndpointS3SettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetBucketFolder(val *string) {
	_jsii_.Set(
		j,
		"bucketFolder",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetCompressionType(val *string) {
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetCsvDelimiter(val *string) {
	_jsii_.Set(
		j,
		"csvDelimiter",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetCsvRowDelimiter(val *string) {
	_jsii_.Set(
		j,
		"csvRowDelimiter",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetDataFormat(val *string) {
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetDatePartitionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"datePartitionEnabled",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetEncryptionMode(val *string) {
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetExternalTableDefinition(val *string) {
	_jsii_.Set(
		j,
		"externalTableDefinition",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetParquetTimestampInMillisecond(val interface{}) {
	_jsii_.Set(
		j,
		"parquetTimestampInMillisecond",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetParquetVersion(val *string) {
	_jsii_.Set(
		j,
		"parquetVersion",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetServerSideEncryptionKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"serverSideEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetServiceAccessRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEndpointS3SettingsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetBucketFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		d,
		"resetBucketName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCsvDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetCsvRowDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvRowDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetDataFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetDatePartitionEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDatePartitionEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetEncryptionMode() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetExternalTableDefinition() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalTableDefinition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetParquetTimestampInMillisecond() {
	_jsii_.InvokeVoid(
		d,
		"resetParquetTimestampInMillisecond",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetParquetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetParquetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetServerSideEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetServerSideEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpointS3SettingsOutputReference) ResetServiceAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccessRoleArn",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html aws_dms_event_subscription}.
type DmsEventSubscription interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EventCategories() *[]*string
	SetEventCategories(val *[]*string)
	EventCategoriesInput() *[]*string
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
	SnsTopicArn() *string
	SetSnsTopicArn(val *string)
	SnsTopicArnInput() *string
	SourceIds() *[]*string
	SetSourceIds(val *[]*string)
	SourceIdsInput() *[]*string
	SourceType() *string
	SetSourceType(val *string)
	SourceTypeInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() DmsEventSubscriptionTimeoutsOutputReference
	TimeoutsInput() *DmsEventSubscriptionTimeouts
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *DmsEventSubscriptionTimeouts)
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetSourceIds()
	ResetSourceType()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DmsEventSubscription
type jsiiProxy_DmsEventSubscription struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsEventSubscription) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) EventCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) EventCategoriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) SnsTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) SourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) SourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) SourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) Timeouts() DmsEventSubscriptionTimeoutsOutputReference {
	var returns DmsEventSubscriptionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscription) TimeoutsInput() *DmsEventSubscriptionTimeouts {
	var returns *DmsEventSubscriptionTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html aws_dms_event_subscription} Resource.
func NewDmsEventSubscription(scope constructs.Construct, id *string, config *DmsEventSubscriptionConfig) DmsEventSubscription {
	_init_.Initialize()

	j := jsiiProxy_DmsEventSubscription{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEventSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html aws_dms_event_subscription} Resource.
func NewDmsEventSubscription_Override(d DmsEventSubscription, scope constructs.Construct, id *string, config *DmsEventSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEventSubscription",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetEventCategories(val *[]*string) {
	_jsii_.Set(
		j,
		"eventCategories",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetSnsTopicArn(val *string) {
	_jsii_.Set(
		j,
		"snsTopicArn",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetSourceIds(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceIds",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetSourceType(val *string) {
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscription) SetTagsAll(val interface{}) {
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
func DmsEventSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DMS.DmsEventSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsEventSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DMS.DmsEventSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DmsEventSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DmsEventSubscription) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsEventSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsEventSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsEventSubscription) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsEventSubscription) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsEventSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsEventSubscription) PutTimeouts(value *DmsEventSubscriptionTimeouts) {
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEventSubscription) ResetEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DmsEventSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEventSubscription) ResetSourceIds() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEventSubscription) ResetSourceType() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEventSubscription) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEventSubscription) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEventSubscription) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEventSubscription) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DmsEventSubscription) ToMetadata() interface{} {
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
func (d *jsiiProxy_DmsEventSubscription) ToString() *string {
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
func (d *jsiiProxy_DmsEventSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DmsEventSubscriptionConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#event_categories DmsEventSubscription#event_categories}.
	EventCategories *[]*string `json:"eventCategories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#name DmsEventSubscription#name}.
	Name *string `json:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#sns_topic_arn DmsEventSubscription#sns_topic_arn}.
	SnsTopicArn *string `json:"snsTopicArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#enabled DmsEventSubscription#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#source_ids DmsEventSubscription#source_ids}.
	SourceIds *[]*string `json:"sourceIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#source_type DmsEventSubscription#source_type}.
	SourceType *string `json:"sourceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#tags DmsEventSubscription#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#tags_all DmsEventSubscription#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#timeouts DmsEventSubscription#timeouts}
	Timeouts *DmsEventSubscriptionTimeouts `json:"timeouts"`
}

type DmsEventSubscriptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#create DmsEventSubscription#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#delete DmsEventSubscription#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_event_subscription.html#update DmsEventSubscription#update}.
	Update *string `json:"update"`
}

type DmsEventSubscriptionTimeoutsOutputReference interface {
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

// The jsii proxy struct for DmsEventSubscriptionTimeoutsOutputReference
type jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewDmsEventSubscriptionTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) DmsEventSubscriptionTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEventSubscriptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDmsEventSubscriptionTimeoutsOutputReference_Override(d DmsEventSubscriptionTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsEventSubscriptionTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEventSubscriptionTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html aws_dms_replication_instance}.
type DmsReplicationInstance interface {
	cdktf.TerraformResource
	AllocatedStorage() *float64
	SetAllocatedStorage(val *float64)
	AllocatedStorageInput() *float64
	AllowMajorVersionUpgrade() interface{}
	SetAllowMajorVersionUpgrade(val interface{})
	AllowMajorVersionUpgradeInput() interface{}
	ApplyImmediately() interface{}
	SetApplyImmediately(val interface{})
	ApplyImmediatelyInput() interface{}
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AutoMinorVersionUpgradeInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MultiAz() interface{}
	SetMultiAz(val interface{})
	MultiAzInput() interface{}
	Node() constructs.Node
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PreferredMaintenanceWindowInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	PubliclyAccessibleInput() interface{}
	RawOverrides() interface{}
	ReplicationInstanceArn() *string
	ReplicationInstanceClass() *string
	SetReplicationInstanceClass(val *string)
	ReplicationInstanceClassInput() *string
	ReplicationInstanceId() *string
	SetReplicationInstanceId(val *string)
	ReplicationInstanceIdInput() *string
	ReplicationInstancePrivateIps() *[]*string
	ReplicationInstancePublicIps() *[]*string
	ReplicationSubnetGroupId() *string
	SetReplicationSubnetGroupId(val *string)
	ReplicationSubnetGroupIdInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() DmsReplicationInstanceTimeoutsOutputReference
	TimeoutsInput() *DmsReplicationInstanceTimeouts
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *DmsReplicationInstanceTimeouts)
	ResetAllocatedStorage()
	ResetAllowMajorVersionUpgrade()
	ResetApplyImmediately()
	ResetAutoMinorVersionUpgrade()
	ResetAvailabilityZone()
	ResetEngineVersion()
	ResetKmsKeyArn()
	ResetMultiAz()
	ResetOverrideLogicalId()
	ResetPreferredMaintenanceWindow()
	ResetPubliclyAccessible()
	ResetReplicationSubnetGroupId()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVpcSecurityGroupIds()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DmsReplicationInstance
type jsiiProxy_DmsReplicationInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsReplicationInstance) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) AllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) AllowMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationInstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationInstancePrivateIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationInstancePrivateIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationInstancePublicIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicationInstancePublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationSubnetGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) ReplicationSubnetGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) Timeouts() DmsReplicationInstanceTimeoutsOutputReference {
	var returns DmsReplicationInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) TimeoutsInput() *DmsReplicationInstanceTimeouts {
	var returns *DmsReplicationInstanceTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstance) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html aws_dms_replication_instance} Resource.
func NewDmsReplicationInstance(scope constructs.Construct, id *string, config *DmsReplicationInstanceConfig) DmsReplicationInstance {
	_init_.Initialize()

	j := jsiiProxy_DmsReplicationInstance{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsReplicationInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html aws_dms_replication_instance} Resource.
func NewDmsReplicationInstance_Override(d DmsReplicationInstance, scope constructs.Construct, id *string, config *DmsReplicationInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsReplicationInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetAllocatedStorage(val *float64) {
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetAllowMajorVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetApplyImmediately(val interface{}) {
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetAutoMinorVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetMultiAz(val interface{}) {
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetPubliclyAccessible(val interface{}) {
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetReplicationInstanceClass(val *string) {
	_jsii_.Set(
		j,
		"replicationInstanceClass",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetReplicationInstanceId(val *string) {
	_jsii_.Set(
		j,
		"replicationInstanceId",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetReplicationSubnetGroupId(val *string) {
	_jsii_.Set(
		j,
		"replicationSubnetGroupId",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstance) SetVpcSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DmsReplicationInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DMS.DmsReplicationInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsReplicationInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DMS.DmsReplicationInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DmsReplicationInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DmsReplicationInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsReplicationInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsReplicationInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsReplicationInstance) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsReplicationInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsReplicationInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsReplicationInstance) PutTimeouts(value *DmsReplicationInstanceTimeouts) {
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		d,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetAllowMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		d,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetMultiAz() {
	_jsii_.InvokeVoid(
		d,
		"resetMultiAz",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DmsReplicationInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		d,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetReplicationSubnetGroupId() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicationSubnetGroupId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstance) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DmsReplicationInstance) ToMetadata() interface{} {
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
func (d *jsiiProxy_DmsReplicationInstance) ToString() *string {
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
func (d *jsiiProxy_DmsReplicationInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DmsReplicationInstanceConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#replication_instance_class DmsReplicationInstance#replication_instance_class}.
	ReplicationInstanceClass *string `json:"replicationInstanceClass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#replication_instance_id DmsReplicationInstance#replication_instance_id}.
	ReplicationInstanceId *string `json:"replicationInstanceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#allocated_storage DmsReplicationInstance#allocated_storage}.
	AllocatedStorage *float64 `json:"allocatedStorage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#allow_major_version_upgrade DmsReplicationInstance#allow_major_version_upgrade}.
	AllowMajorVersionUpgrade interface{} `json:"allowMajorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#apply_immediately DmsReplicationInstance#apply_immediately}.
	ApplyImmediately interface{} `json:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#auto_minor_version_upgrade DmsReplicationInstance#auto_minor_version_upgrade}.
	AutoMinorVersionUpgrade interface{} `json:"autoMinorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#availability_zone DmsReplicationInstance#availability_zone}.
	AvailabilityZone *string `json:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#engine_version DmsReplicationInstance#engine_version}.
	EngineVersion *string `json:"engineVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#kms_key_arn DmsReplicationInstance#kms_key_arn}.
	KmsKeyArn *string `json:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#multi_az DmsReplicationInstance#multi_az}.
	MultiAz interface{} `json:"multiAz"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#preferred_maintenance_window DmsReplicationInstance#preferred_maintenance_window}.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#publicly_accessible DmsReplicationInstance#publicly_accessible}.
	PubliclyAccessible interface{} `json:"publiclyAccessible"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#replication_subnet_group_id DmsReplicationInstance#replication_subnet_group_id}.
	ReplicationSubnetGroupId *string `json:"replicationSubnetGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#tags DmsReplicationInstance#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#tags_all DmsReplicationInstance#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#timeouts DmsReplicationInstance#timeouts}
	Timeouts *DmsReplicationInstanceTimeouts `json:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#vpc_security_group_ids DmsReplicationInstance#vpc_security_group_ids}.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds"`
}

type DmsReplicationInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#create DmsReplicationInstance#create}.
	Create *string `json:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#delete DmsReplicationInstance#delete}.
	Delete *string `json:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_instance.html#update DmsReplicationInstance#update}.
	Update *string `json:"update"`
}

type DmsReplicationInstanceTimeoutsOutputReference interface {
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

// The jsii proxy struct for DmsReplicationInstanceTimeoutsOutputReference
type jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewDmsReplicationInstanceTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) DmsReplicationInstanceTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsReplicationInstanceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewDmsReplicationInstanceTimeoutsOutputReference_Override(d DmsReplicationInstanceTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsReplicationInstanceTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		d,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationInstanceTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_subnet_group.html aws_dms_replication_subnet_group}.
type DmsReplicationSubnetGroup interface {
	cdktf.TerraformResource
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
	ReplicationSubnetGroupArn() *string
	ReplicationSubnetGroupDescription() *string
	SetReplicationSubnetGroupDescription(val *string)
	ReplicationSubnetGroupDescriptionInput() *string
	ReplicationSubnetGroupId() *string
	SetReplicationSubnetGroupId(val *string)
	ReplicationSubnetGroupIdInput() *string
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
	VpcId() *string
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DmsReplicationSubnetGroup
type jsiiProxy_DmsReplicationSubnetGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) ReplicationSubnetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) ReplicationSubnetGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) ReplicationSubnetGroupDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) ReplicationSubnetGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) ReplicationSubnetGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSubnetGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_subnet_group.html aws_dms_replication_subnet_group} Resource.
func NewDmsReplicationSubnetGroup(scope constructs.Construct, id *string, config *DmsReplicationSubnetGroupConfig) DmsReplicationSubnetGroup {
	_init_.Initialize()

	j := jsiiProxy_DmsReplicationSubnetGroup{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsReplicationSubnetGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_subnet_group.html aws_dms_replication_subnet_group} Resource.
func NewDmsReplicationSubnetGroup_Override(d DmsReplicationSubnetGroup, scope constructs.Construct, id *string, config *DmsReplicationSubnetGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsReplicationSubnetGroup",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SetReplicationSubnetGroupDescription(val *string) {
	_jsii_.Set(
		j,
		"replicationSubnetGroupDescription",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SetReplicationSubnetGroupId(val *string) {
	_jsii_.Set(
		j,
		"replicationSubnetGroupId",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationSubnetGroup) SetTagsAll(val interface{}) {
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
func DmsReplicationSubnetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DMS.DmsReplicationSubnetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsReplicationSubnetGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DMS.DmsReplicationSubnetGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DmsReplicationSubnetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DmsReplicationSubnetGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsReplicationSubnetGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsReplicationSubnetGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsReplicationSubnetGroup) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsReplicationSubnetGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsReplicationSubnetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DmsReplicationSubnetGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationSubnetGroup) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationSubnetGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationSubnetGroup) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DmsReplicationSubnetGroup) ToMetadata() interface{} {
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
func (d *jsiiProxy_DmsReplicationSubnetGroup) ToString() *string {
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
func (d *jsiiProxy_DmsReplicationSubnetGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DmsReplicationSubnetGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_subnet_group.html#replication_subnet_group_description DmsReplicationSubnetGroup#replication_subnet_group_description}.
	ReplicationSubnetGroupDescription *string `json:"replicationSubnetGroupDescription"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_subnet_group.html#replication_subnet_group_id DmsReplicationSubnetGroup#replication_subnet_group_id}.
	ReplicationSubnetGroupId *string `json:"replicationSubnetGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_subnet_group.html#subnet_ids DmsReplicationSubnetGroup#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_subnet_group.html#tags DmsReplicationSubnetGroup#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_subnet_group.html#tags_all DmsReplicationSubnetGroup#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html aws_dms_replication_task}.
type DmsReplicationTask interface {
	cdktf.TerraformResource
	CdcStartPosition() *string
	SetCdcStartPosition(val *string)
	CdcStartPositionInput() *string
	CdcStartTime() *string
	SetCdcStartTime(val *string)
	CdcStartTimeInput() *string
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
	MigrationType() *string
	SetMigrationType(val *string)
	MigrationTypeInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ReplicationInstanceArn() *string
	SetReplicationInstanceArn(val *string)
	ReplicationInstanceArnInput() *string
	ReplicationTaskArn() *string
	ReplicationTaskId() *string
	SetReplicationTaskId(val *string)
	ReplicationTaskIdInput() *string
	ReplicationTaskSettings() *string
	SetReplicationTaskSettings(val *string)
	ReplicationTaskSettingsInput() *string
	SourceEndpointArn() *string
	SetSourceEndpointArn(val *string)
	SourceEndpointArnInput() *string
	TableMappings() *string
	SetTableMappings(val *string)
	TableMappingsInput() *string
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TargetEndpointArn() *string
	SetTargetEndpointArn(val *string)
	TargetEndpointArnInput() *string
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
	ResetCdcStartPosition()
	ResetCdcStartTime()
	ResetOverrideLogicalId()
	ResetReplicationTaskSettings()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DmsReplicationTask
type jsiiProxy_DmsReplicationTask struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsReplicationTask) CdcStartPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) CdcStartPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) CdcStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) CdcStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) MigrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) MigrationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) ReplicationInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) ReplicationInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) ReplicationTaskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) ReplicationTaskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) ReplicationTaskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) ReplicationTaskSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) ReplicationTaskSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) SourceEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) SourceEndpointArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndpointArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TableMappings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TableMappingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TargetEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TargetEndpointArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEndpointArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsReplicationTask) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html aws_dms_replication_task} Resource.
func NewDmsReplicationTask(scope constructs.Construct, id *string, config *DmsReplicationTaskConfig) DmsReplicationTask {
	_init_.Initialize()

	j := jsiiProxy_DmsReplicationTask{}

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsReplicationTask",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html aws_dms_replication_task} Resource.
func NewDmsReplicationTask_Override(d DmsReplicationTask, scope constructs.Construct, id *string, config *DmsReplicationTaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.DMS.DmsReplicationTask",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetCdcStartPosition(val *string) {
	_jsii_.Set(
		j,
		"cdcStartPosition",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetCdcStartTime(val *string) {
	_jsii_.Set(
		j,
		"cdcStartTime",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetMigrationType(val *string) {
	_jsii_.Set(
		j,
		"migrationType",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetReplicationInstanceArn(val *string) {
	_jsii_.Set(
		j,
		"replicationInstanceArn",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetReplicationTaskId(val *string) {
	_jsii_.Set(
		j,
		"replicationTaskId",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetReplicationTaskSettings(val *string) {
	_jsii_.Set(
		j,
		"replicationTaskSettings",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetSourceEndpointArn(val *string) {
	_jsii_.Set(
		j,
		"sourceEndpointArn",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetTableMappings(val *string) {
	_jsii_.Set(
		j,
		"tableMappings",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetTagsAll(val interface{}) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DmsReplicationTask) SetTargetEndpointArn(val *string) {
	_jsii_.Set(
		j,
		"targetEndpointArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DmsReplicationTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.DMS.DmsReplicationTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsReplicationTask_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.DMS.DmsReplicationTask",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DmsReplicationTask) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DmsReplicationTask) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsReplicationTask) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DmsReplicationTask) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DmsReplicationTask) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DmsReplicationTask) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DmsReplicationTask) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsReplicationTask) ResetCdcStartPosition() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcStartPosition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationTask) ResetCdcStartTime() {
	_jsii_.InvokeVoid(
		d,
		"resetCdcStartTime",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DmsReplicationTask) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationTask) ResetReplicationTaskSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicationTaskSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationTask) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationTask) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsReplicationTask) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DmsReplicationTask) ToMetadata() interface{} {
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
func (d *jsiiProxy_DmsReplicationTask) ToString() *string {
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
func (d *jsiiProxy_DmsReplicationTask) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DmsReplicationTaskConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#migration_type DmsReplicationTask#migration_type}.
	MigrationType *string `json:"migrationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#replication_instance_arn DmsReplicationTask#replication_instance_arn}.
	ReplicationInstanceArn *string `json:"replicationInstanceArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#replication_task_id DmsReplicationTask#replication_task_id}.
	ReplicationTaskId *string `json:"replicationTaskId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#source_endpoint_arn DmsReplicationTask#source_endpoint_arn}.
	SourceEndpointArn *string `json:"sourceEndpointArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#table_mappings DmsReplicationTask#table_mappings}.
	TableMappings *string `json:"tableMappings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#target_endpoint_arn DmsReplicationTask#target_endpoint_arn}.
	TargetEndpointArn *string `json:"targetEndpointArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#cdc_start_position DmsReplicationTask#cdc_start_position}.
	CdcStartPosition *string `json:"cdcStartPosition"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#cdc_start_time DmsReplicationTask#cdc_start_time}.
	CdcStartTime *string `json:"cdcStartTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#replication_task_settings DmsReplicationTask#replication_task_settings}.
	ReplicationTaskSettings *string `json:"replicationTaskSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#tags DmsReplicationTask#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dms_replication_task.html#tags_all DmsReplicationTask#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}
