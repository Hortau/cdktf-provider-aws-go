package acm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/acm/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate aws_acm_certificate}.
type AcmCertificate interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
	CertificateAuthorityArnInput() *string
	CertificateBody() *string
	SetCertificateBody(val *string)
	CertificateBodyInput() *string
	CertificateChain() *string
	SetCertificateChain(val *string)
	CertificateChainInput() *string
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
	Options() AcmCertificateOptionsOutputReference
	OptionsInput() *AcmCertificateOptions
	PrivateKey() *string
	SetPrivateKey(val *string)
	PrivateKeyInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	SubjectAlternativeNames() *[]*string
	SetSubjectAlternativeNames(val *[]*string)
	SubjectAlternativeNamesInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	ValidationEmails() *[]*string
	ValidationMethod() *string
	SetValidationMethod(val *string)
	ValidationMethodInput() *string
	AddOverride(path *string, value interface{})
	DomainValidationOptions(index *string) AcmCertificateDomainValidationOptions
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
	PutOptions(value *AcmCertificateOptions)
	ResetCertificateAuthorityArn()
	ResetCertificateBody()
	ResetCertificateChain()
	ResetDomainName()
	ResetOptions()
	ResetOverrideLogicalId()
	ResetPrivateKey()
	ResetSubjectAlternativeNames()
	ResetTags()
	ResetTagsAll()
	ResetValidationMethod()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AcmCertificate
type jsiiProxy_AcmCertificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AcmCertificate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) CertificateAuthorityArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) CertificateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) CertificateBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) CertificateChainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) Options() AcmCertificateOptionsOutputReference {
	var returns AcmCertificateOptionsOutputReference
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) OptionsInput() *AcmCertificateOptions {
	var returns *AcmCertificateOptions
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) SubjectAlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) SubjectAlternativeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) ValidationEmails() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validationEmails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) ValidationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificate) ValidationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationMethodInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate aws_acm_certificate} Resource.
func NewAcmCertificate(scope constructs.Construct, id *string, config *AcmCertificateConfig) AcmCertificate {
	_init_.Initialize()

	j := jsiiProxy_AcmCertificate{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate aws_acm_certificate} Resource.
func NewAcmCertificate_Override(a AcmCertificate, scope constructs.Construct, id *string, config *AcmCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificate",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AcmCertificate) SetCertificateAuthorityArn(val *string) {
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetCertificateBody(val *string) {
	_jsii_.Set(
		j,
		"certificateBody",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetCertificateChain(val *string) {
	_jsii_.Set(
		j,
		"certificateChain",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetPrivateKey(val *string) {
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetSubjectAlternativeNames(val *[]*string) {
	_jsii_.Set(
		j,
		"subjectAlternativeNames",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AcmCertificate) SetValidationMethod(val *string) {
	_jsii_.Set(
		j,
		"validationMethod",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AcmCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.acm.AcmCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AcmCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.acm.AcmCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AcmCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AcmCertificate) DomainValidationOptions(index *string) AcmCertificateDomainValidationOptions {
	var returns AcmCertificateDomainValidationOptions

	_jsii_.Invoke(
		a,
		"domainValidationOptions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (a *jsiiProxy_AcmCertificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmCertificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmCertificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmCertificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmCertificate) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmCertificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AcmCertificate) PutOptions(value *AcmCertificateOptions) {
	_jsii_.InvokeVoid(
		a,
		"putOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmCertificate) ResetCertificateAuthorityArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateAuthorityArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) ResetCertificateBody() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) ResetCertificateChain() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateChain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) ResetDomainName() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) ResetOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetOptions",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AcmCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) ResetSubjectAlternativeNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectAlternativeNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) ResetValidationMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetValidationMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificate) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AcmCertificate) ToMetadata() interface{} {
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
func (a *jsiiProxy_AcmCertificate) ToString() *string {
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
func (a *jsiiProxy_AcmCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon Certificate Manager.
type AcmCertificateConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#certificate_authority_arn AcmCertificate#certificate_authority_arn}.
	CertificateAuthorityArn *string `json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#certificate_body AcmCertificate#certificate_body}.
	CertificateBody *string `json:"certificateBody" yaml:"certificateBody"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#certificate_chain AcmCertificate#certificate_chain}.
	CertificateChain *string `json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#domain_name AcmCertificate#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#options AcmCertificate#options}
	Options *AcmCertificateOptions `json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#private_key AcmCertificate#private_key}.
	PrivateKey *string `json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#subject_alternative_names AcmCertificate#subject_alternative_names}.
	SubjectAlternativeNames *[]*string `json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#tags AcmCertificate#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#tags_all AcmCertificate#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#validation_method AcmCertificate#validation_method}.
	ValidationMethod *string `json:"validationMethod" yaml:"validationMethod"`
}

type AcmCertificateDomainValidationOptions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DomainName() *string
	ResourceRecordName() *string
	ResourceRecordType() *string
	ResourceRecordValue() *string
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

// The jsii proxy struct for AcmCertificateDomainValidationOptions
type jsiiProxy_AcmCertificateDomainValidationOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) ResourceRecordName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceRecordName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) ResourceRecordType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceRecordType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) ResourceRecordValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceRecordValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewAcmCertificateDomainValidationOptions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) AcmCertificateDomainValidationOptions {
	_init_.Initialize()

	j := jsiiProxy_AcmCertificateDomainValidationOptions{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificateDomainValidationOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAcmCertificateDomainValidationOptions_Override(a AcmCertificateDomainValidationOptions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificateDomainValidationOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateDomainValidationOptions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmCertificateDomainValidationOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type AcmCertificateOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate#certificate_transparency_logging_preference AcmCertificate#certificate_transparency_logging_preference}.
	CertificateTransparencyLoggingPreference *string `json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
}

type AcmCertificateOptionsOutputReference interface {
	cdktf.ComplexObject
	CertificateTransparencyLoggingPreference() *string
	SetCertificateTransparencyLoggingPreference(val *string)
	CertificateTransparencyLoggingPreferenceInput() *string
	InternalValue() *AcmCertificateOptions
	SetInternalValue(val *AcmCertificateOptions)
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
	ResetCertificateTransparencyLoggingPreference()
}

// The jsii proxy struct for AcmCertificateOptionsOutputReference
type jsiiProxy_AcmCertificateOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) CertificateTransparencyLoggingPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateTransparencyLoggingPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) CertificateTransparencyLoggingPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateTransparencyLoggingPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) InternalValue() *AcmCertificateOptions {
	var returns *AcmCertificateOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAcmCertificateOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AcmCertificateOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AcmCertificateOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificateOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAcmCertificateOptionsOutputReference_Override(a AcmCertificateOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificateOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) SetCertificateTransparencyLoggingPreference(val *string) {
	_jsii_.Set(
		j,
		"certificateTransparencyLoggingPreference",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) SetInternalValue(val *AcmCertificateOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AcmCertificateOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmCertificateOptionsOutputReference) ResetCertificateTransparencyLoggingPreference() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateTransparencyLoggingPreference",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate_validation aws_acm_certificate_validation}.
type AcmCertificateValidation interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
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
	Timeouts() AcmCertificateValidationTimeoutsOutputReference
	TimeoutsInput() *AcmCertificateValidationTimeouts
	ValidationRecordFqdns() *[]*string
	SetValidationRecordFqdns(val *[]*string)
	ValidationRecordFqdnsInput() *[]*string
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
	PutTimeouts(value *AcmCertificateValidationTimeouts)
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetValidationRecordFqdns()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AcmCertificateValidation
type jsiiProxy_AcmCertificateValidation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AcmCertificateValidation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) Timeouts() AcmCertificateValidationTimeoutsOutputReference {
	var returns AcmCertificateValidationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) TimeoutsInput() *AcmCertificateValidationTimeouts {
	var returns *AcmCertificateValidationTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) ValidationRecordFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validationRecordFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidation) ValidationRecordFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validationRecordFqdnsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate_validation aws_acm_certificate_validation} Resource.
func NewAcmCertificateValidation(scope constructs.Construct, id *string, config *AcmCertificateValidationConfig) AcmCertificateValidation {
	_init_.Initialize()

	j := jsiiProxy_AcmCertificateValidation{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificateValidation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate_validation aws_acm_certificate_validation} Resource.
func NewAcmCertificateValidation_Override(a AcmCertificateValidation, scope constructs.Construct, id *string, config *AcmCertificateValidationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificateValidation",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AcmCertificateValidation) SetCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateValidation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateValidation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateValidation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateValidation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateValidation) SetValidationRecordFqdns(val *[]*string) {
	_jsii_.Set(
		j,
		"validationRecordFqdns",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AcmCertificateValidation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.acm.AcmCertificateValidation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AcmCertificateValidation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.acm.AcmCertificateValidation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AcmCertificateValidation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AcmCertificateValidation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmCertificateValidation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmCertificateValidation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmCertificateValidation) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmCertificateValidation) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmCertificateValidation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmCertificateValidation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmCertificateValidation) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmCertificateValidation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmCertificateValidation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmCertificateValidation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AcmCertificateValidation) PutTimeouts(value *AcmCertificateValidationTimeouts) {
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AcmCertificateValidation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificateValidation) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificateValidation) ResetValidationRecordFqdns() {
	_jsii_.InvokeVoid(
		a,
		"resetValidationRecordFqdns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmCertificateValidation) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AcmCertificateValidation) ToMetadata() interface{} {
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
func (a *jsiiProxy_AcmCertificateValidation) ToString() *string {
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
func (a *jsiiProxy_AcmCertificateValidation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon Certificate Manager.
type AcmCertificateValidationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate_validation#certificate_arn AcmCertificateValidation#certificate_arn}.
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate_validation#timeouts AcmCertificateValidation#timeouts}
	Timeouts *AcmCertificateValidationTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate_validation#validation_record_fqdns AcmCertificateValidation#validation_record_fqdns}.
	ValidationRecordFqdns *[]*string `json:"validationRecordFqdns" yaml:"validationRecordFqdns"`
}

type AcmCertificateValidationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acm_certificate_validation#create AcmCertificateValidation#create}.
	Create *string `json:"create" yaml:"create"`
}

type AcmCertificateValidationTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	InternalValue() *AcmCertificateValidationTimeouts
	SetInternalValue(val *AcmCertificateValidationTimeouts)
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
}

// The jsii proxy struct for AcmCertificateValidationTimeoutsOutputReference
type jsiiProxy_AcmCertificateValidationTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) InternalValue() *AcmCertificateValidationTimeouts {
	var returns *AcmCertificateValidationTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAcmCertificateValidationTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AcmCertificateValidationTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AcmCertificateValidationTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificateValidationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAcmCertificateValidationTimeoutsOutputReference_Override(a AcmCertificateValidationTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmCertificateValidationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) SetInternalValue(val *AcmCertificateValidationTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmCertificateValidationTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		a,
		"resetCreate",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate aws_acmpca_certificate}.
type AcmpcaCertificate interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
	CertificateAuthorityArnInput() *string
	CertificateChain() *string
	CertificateSigningRequest() *string
	SetCertificateSigningRequest(val *string)
	CertificateSigningRequestInput() *string
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
	SigningAlgorithm() *string
	SetSigningAlgorithm(val *string)
	SigningAlgorithmInput() *string
	TemplateArn() *string
	SetTemplateArn(val *string)
	TemplateArnInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Validity() AcmpcaCertificateValidityOutputReference
	ValidityInput() *AcmpcaCertificateValidity
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
	PutValidity(value *AcmpcaCertificateValidity)
	ResetOverrideLogicalId()
	ResetTemplateArn()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AcmpcaCertificate
type jsiiProxy_AcmpcaCertificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AcmpcaCertificate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) CertificateAuthorityArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) CertificateSigningRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSigningRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) CertificateSigningRequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSigningRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) SigningAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) SigningAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) TemplateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) TemplateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) Validity() AcmpcaCertificateValidityOutputReference {
	var returns AcmpcaCertificateValidityOutputReference
	_jsii_.Get(
		j,
		"validity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificate) ValidityInput() *AcmpcaCertificateValidity {
	var returns *AcmpcaCertificateValidity
	_jsii_.Get(
		j,
		"validityInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate aws_acmpca_certificate} Resource.
func NewAcmpcaCertificate(scope constructs.Construct, id *string, config *AcmpcaCertificateConfig) AcmpcaCertificate {
	_init_.Initialize()

	j := jsiiProxy_AcmpcaCertificate{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate aws_acmpca_certificate} Resource.
func NewAcmpcaCertificate_Override(a AcmpcaCertificate, scope constructs.Construct, id *string, config *AcmpcaCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificate",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificate) SetCertificateAuthorityArn(val *string) {
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificate) SetCertificateSigningRequest(val *string) {
	_jsii_.Set(
		j,
		"certificateSigningRequest",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificate) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificate) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificate) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificate) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificate) SetSigningAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"signingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificate) SetTemplateArn(val *string) {
	_jsii_.Set(
		j,
		"templateArn",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AcmpcaCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.acm.AcmpcaCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AcmpcaCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.acm.AcmpcaCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmpcaCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmpcaCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmpcaCertificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificate) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmpcaCertificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmpcaCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AcmpcaCertificate) PutValidity(value *AcmpcaCertificateValidity) {
	_jsii_.InvokeVoid(
		a,
		"putValidity",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AcmpcaCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificate) ResetTemplateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTemplateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificate) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificate) ToMetadata() interface{} {
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
func (a *jsiiProxy_AcmpcaCertificate) ToString() *string {
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
func (a *jsiiProxy_AcmpcaCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority aws_acmpca_certificate_authority}.
type AcmpcaCertificateAuthority interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	CertificateAuthorityConfiguration() AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference
	CertificateAuthorityConfigurationInput() *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration
	CertificateChain() *string
	CertificateSigningRequest() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	NotAfter() *string
	NotBefore() *string
	PermanentDeletionTimeInDays() *float64
	SetPermanentDeletionTimeInDays(val *float64)
	PermanentDeletionTimeInDaysInput() *float64
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RevocationConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationOutputReference
	RevocationConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfiguration
	Serial() *string
	Status() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() AcmpcaCertificateAuthorityTimeoutsOutputReference
	TimeoutsInput() *AcmpcaCertificateAuthorityTimeouts
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
	PutCertificateAuthorityConfiguration(value *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration)
	PutRevocationConfiguration(value *AcmpcaCertificateAuthorityRevocationConfiguration)
	PutTimeouts(value *AcmpcaCertificateAuthorityTimeouts)
	ResetEnabled()
	ResetOverrideLogicalId()
	ResetPermanentDeletionTimeInDays()
	ResetRevocationConfiguration()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AcmpcaCertificateAuthority
type jsiiProxy_AcmpcaCertificateAuthority struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CertificateAuthorityConfiguration() AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference {
	var returns AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference
	_jsii_.Get(
		j,
		"certificateAuthorityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CertificateAuthorityConfigurationInput() *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration {
	var returns *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration
	_jsii_.Get(
		j,
		"certificateAuthorityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) CertificateSigningRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSigningRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) NotAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) NotBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) PermanentDeletionTimeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"permanentDeletionTimeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) PermanentDeletionTimeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"permanentDeletionTimeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) RevocationConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationOutputReference {
	var returns AcmpcaCertificateAuthorityRevocationConfigurationOutputReference
	_jsii_.Get(
		j,
		"revocationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) RevocationConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfiguration {
	var returns *AcmpcaCertificateAuthorityRevocationConfiguration
	_jsii_.Get(
		j,
		"revocationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Serial() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Timeouts() AcmpcaCertificateAuthorityTimeoutsOutputReference {
	var returns AcmpcaCertificateAuthorityTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TimeoutsInput() *AcmpcaCertificateAuthorityTimeouts {
	var returns *AcmpcaCertificateAuthorityTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority aws_acmpca_certificate_authority} Resource.
func NewAcmpcaCertificateAuthority(scope constructs.Construct, id *string, config *AcmpcaCertificateAuthorityConfig) AcmpcaCertificateAuthority {
	_init_.Initialize()

	j := jsiiProxy_AcmpcaCertificateAuthority{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthority",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority aws_acmpca_certificate_authority} Resource.
func NewAcmpcaCertificateAuthority_Override(a AcmpcaCertificateAuthority, scope constructs.Construct, id *string, config *AcmpcaCertificateAuthorityConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthority",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) SetPermanentDeletionTimeInDays(val *float64) {
	_jsii_.Set(
		j,
		"permanentDeletionTimeInDays",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthority) SetType(val *string) {
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
func AcmpcaCertificateAuthority_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.acm.AcmpcaCertificateAuthority",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AcmpcaCertificateAuthority_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.acm.AcmpcaCertificateAuthority",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthority) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthority) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) PutCertificateAuthorityConfiguration(value *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putCertificateAuthorityConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) PutRevocationConfiguration(value *AcmpcaCertificateAuthorityRevocationConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putRevocationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) PutTimeouts(value *AcmpcaCertificateAuthorityTimeouts) {
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetPermanentDeletionTimeInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetPermanentDeletionTimeInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetRevocationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRevocationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthority) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) ToMetadata() interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) ToString() *string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthority) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority_certificate aws_acmpca_certificate_authority_certificate}.
type AcmpcaCertificateAuthorityCertificate interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
	CertificateAuthorityArnInput() *string
	CertificateChain() *string
	SetCertificateChain(val *string)
	CertificateChainInput() *string
	CertificateInput() *string
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
	ResetCertificateChain()
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AcmpcaCertificateAuthorityCertificate
type jsiiProxy_AcmpcaCertificateAuthorityCertificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) CertificateAuthorityArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) CertificateChainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority_certificate aws_acmpca_certificate_authority_certificate} Resource.
func NewAcmpcaCertificateAuthorityCertificate(scope constructs.Construct, id *string, config *AcmpcaCertificateAuthorityCertificateConfig) AcmpcaCertificateAuthorityCertificate {
	_init_.Initialize()

	j := jsiiProxy_AcmpcaCertificateAuthorityCertificate{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority_certificate aws_acmpca_certificate_authority_certificate} Resource.
func NewAcmpcaCertificateAuthorityCertificate_Override(a AcmpcaCertificateAuthorityCertificate, scope constructs.Construct, id *string, config *AcmpcaCertificateAuthorityCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityCertificate",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) SetCertificate(val *string) {
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) SetCertificateAuthorityArn(val *string) {
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) SetCertificateChain(val *string) {
	_jsii_.Set(
		j,
		"certificateChain",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificate) SetProvider(val cdktf.TerraformProvider) {
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
func AcmpcaCertificateAuthorityCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AcmpcaCertificateAuthorityCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) ResetCertificateChain() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateChain",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) ToMetadata() interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) ToString() *string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AcmpcaCertificateAuthorityCertificateAuthorityConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#key_algorithm AcmpcaCertificateAuthority#key_algorithm}.
	KeyAlgorithm *string `json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#signing_algorithm AcmpcaCertificateAuthority#signing_algorithm}.
	SigningAlgorithm *string `json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#subject AcmpcaCertificateAuthority#subject}
	Subject *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject `json:"subject" yaml:"subject"`
}

type AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration
	SetInternalValue(val *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KeyAlgorithm() *string
	SetKeyAlgorithm(val *string)
	KeyAlgorithmInput() *string
	SigningAlgorithm() *string
	SetSigningAlgorithm(val *string)
	SigningAlgorithmInput() *string
	Subject() AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference
	SubjectInput() *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject
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
	PutSubject(value *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject)
}

// The jsii proxy struct for AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference
type jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) InternalValue() *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration {
	var returns *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) KeyAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) KeyAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) SigningAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) SigningAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) Subject() AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference {
	var returns AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) SubjectInput() *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject {
	var returns *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject
	_jsii_.Get(
		j,
		"subjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference_Override(a AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) SetInternalValue(val *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) SetKeyAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"keyAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) SetSigningAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"signingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationOutputReference) PutSubject(value *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject) {
	_jsii_.InvokeVoid(
		a,
		"putSubject",
		[]interface{}{value},
	)
}

type AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#common_name AcmpcaCertificateAuthority#common_name}.
	CommonName *string `json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#country AcmpcaCertificateAuthority#country}.
	Country *string `json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#distinguished_name_qualifier AcmpcaCertificateAuthority#distinguished_name_qualifier}.
	DistinguishedNameQualifier *string `json:"distinguishedNameQualifier" yaml:"distinguishedNameQualifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#generation_qualifier AcmpcaCertificateAuthority#generation_qualifier}.
	GenerationQualifier *string `json:"generationQualifier" yaml:"generationQualifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#given_name AcmpcaCertificateAuthority#given_name}.
	GivenName *string `json:"givenName" yaml:"givenName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#initials AcmpcaCertificateAuthority#initials}.
	Initials *string `json:"initials" yaml:"initials"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#locality AcmpcaCertificateAuthority#locality}.
	Locality *string `json:"locality" yaml:"locality"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#organization AcmpcaCertificateAuthority#organization}.
	Organization *string `json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#organizational_unit AcmpcaCertificateAuthority#organizational_unit}.
	OrganizationalUnit *string `json:"organizationalUnit" yaml:"organizationalUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#pseudonym AcmpcaCertificateAuthority#pseudonym}.
	Pseudonym *string `json:"pseudonym" yaml:"pseudonym"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#state AcmpcaCertificateAuthority#state}.
	State *string `json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#surname AcmpcaCertificateAuthority#surname}.
	Surname *string `json:"surname" yaml:"surname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#title AcmpcaCertificateAuthority#title}.
	Title *string `json:"title" yaml:"title"`
}

type AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference interface {
	cdktf.ComplexObject
	CommonName() *string
	SetCommonName(val *string)
	CommonNameInput() *string
	Country() *string
	SetCountry(val *string)
	CountryInput() *string
	DistinguishedNameQualifier() *string
	SetDistinguishedNameQualifier(val *string)
	DistinguishedNameQualifierInput() *string
	GenerationQualifier() *string
	SetGenerationQualifier(val *string)
	GenerationQualifierInput() *string
	GivenName() *string
	SetGivenName(val *string)
	GivenNameInput() *string
	Initials() *string
	SetInitials(val *string)
	InitialsInput() *string
	InternalValue() *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject
	SetInternalValue(val *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Locality() *string
	SetLocality(val *string)
	LocalityInput() *string
	Organization() *string
	SetOrganization(val *string)
	OrganizationalUnit() *string
	SetOrganizationalUnit(val *string)
	OrganizationalUnitInput() *string
	OrganizationInput() *string
	Pseudonym() *string
	SetPseudonym(val *string)
	PseudonymInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	Surname() *string
	SetSurname(val *string)
	SurnameInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
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
	ResetCommonName()
	ResetCountry()
	ResetDistinguishedNameQualifier()
	ResetGenerationQualifier()
	ResetGivenName()
	ResetInitials()
	ResetLocality()
	ResetOrganization()
	ResetOrganizationalUnit()
	ResetPseudonym()
	ResetState()
	ResetSurname()
	ResetTitle()
}

// The jsii proxy struct for AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference
type jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) DistinguishedNameQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distinguishedNameQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) DistinguishedNameQualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distinguishedNameQualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GenerationQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generationQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GenerationQualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generationQualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GivenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GivenNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) Initials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) InitialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) InternalValue() *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject {
	var returns *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) OrganizationalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) OrganizationalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) Pseudonym() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pseudonym",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) PseudonymInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pseudonymInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) Surname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SurnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"surnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func NewAcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference_Override(a AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetCommonName(val *string) {
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetCountry(val *string) {
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetDistinguishedNameQualifier(val *string) {
	_jsii_.Set(
		j,
		"distinguishedNameQualifier",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetGenerationQualifier(val *string) {
	_jsii_.Set(
		j,
		"generationQualifier",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetGivenName(val *string) {
	_jsii_.Set(
		j,
		"givenName",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetInitials(val *string) {
	_jsii_.Set(
		j,
		"initials",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetInternalValue(val *AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubject) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetLocality(val *string) {
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetOrganization(val *string) {
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetOrganizationalUnit(val *string) {
	_jsii_.Set(
		j,
		"organizationalUnit",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetPseudonym(val *string) {
	_jsii_.Set(
		j,
		"pseudonym",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetSurname(val *string) {
	_jsii_.Set(
		j,
		"surname",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) SetTitle(val *string) {
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		a,
		"resetCommonName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetCountry() {
	_jsii_.InvokeVoid(
		a,
		"resetCountry",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetDistinguishedNameQualifier() {
	_jsii_.InvokeVoid(
		a,
		"resetDistinguishedNameQualifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetGenerationQualifier() {
	_jsii_.InvokeVoid(
		a,
		"resetGenerationQualifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetGivenName() {
	_jsii_.InvokeVoid(
		a,
		"resetGivenName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetInitials() {
	_jsii_.InvokeVoid(
		a,
		"resetInitials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetLocality() {
	_jsii_.InvokeVoid(
		a,
		"resetLocality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetOrganization() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetOrganizationalUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganizationalUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetPseudonym() {
	_jsii_.InvokeVoid(
		a,
		"resetPseudonym",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		a,
		"resetState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetSurname() {
	_jsii_.InvokeVoid(
		a,
		"resetSurname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCertificateAuthorityConfigurationSubjectOutputReference) ResetTitle() {
	_jsii_.InvokeVoid(
		a,
		"resetTitle",
		nil, // no parameters
	)
}

// Amazon Certificate Manager.
type AcmpcaCertificateAuthorityCertificateConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority_certificate#certificate AcmpcaCertificateAuthorityCertificate#certificate}.
	Certificate *string `json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority_certificate#certificate_authority_arn AcmpcaCertificateAuthorityCertificate#certificate_authority_arn}.
	CertificateAuthorityArn *string `json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority_certificate#certificate_chain AcmpcaCertificateAuthorityCertificate#certificate_chain}.
	CertificateChain *string `json:"certificateChain" yaml:"certificateChain"`
}

// Amazon Certificate Manager.
type AcmpcaCertificateAuthorityConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// certificate_authority_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#certificate_authority_configuration AcmpcaCertificateAuthority#certificate_authority_configuration}
	CertificateAuthorityConfiguration *AcmpcaCertificateAuthorityCertificateAuthorityConfiguration `json:"certificateAuthorityConfiguration" yaml:"certificateAuthorityConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#enabled AcmpcaCertificateAuthority#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#permanent_deletion_time_in_days AcmpcaCertificateAuthority#permanent_deletion_time_in_days}.
	PermanentDeletionTimeInDays *float64 `json:"permanentDeletionTimeInDays" yaml:"permanentDeletionTimeInDays"`
	// revocation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#revocation_configuration AcmpcaCertificateAuthority#revocation_configuration}
	RevocationConfiguration *AcmpcaCertificateAuthorityRevocationConfiguration `json:"revocationConfiguration" yaml:"revocationConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#tags AcmpcaCertificateAuthority#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#tags_all AcmpcaCertificateAuthority#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#timeouts AcmpcaCertificateAuthority#timeouts}
	Timeouts *AcmpcaCertificateAuthorityTimeouts `json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#type AcmpcaCertificateAuthority#type}.
	Type *string `json:"type" yaml:"type"`
}

type AcmpcaCertificateAuthorityRevocationConfiguration struct {
	// crl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#crl_configuration AcmpcaCertificateAuthority#crl_configuration}
	CrlConfiguration *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration `json:"crlConfiguration" yaml:"crlConfiguration"`
}

type AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#expiration_in_days AcmpcaCertificateAuthority#expiration_in_days}.
	ExpirationInDays *float64 `json:"expirationInDays" yaml:"expirationInDays"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#custom_cname AcmpcaCertificateAuthority#custom_cname}.
	CustomCname *string `json:"customCname" yaml:"customCname"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#enabled AcmpcaCertificateAuthority#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#s3_bucket_name AcmpcaCertificateAuthority#s3_bucket_name}.
	S3BucketName *string `json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#s3_object_acl AcmpcaCertificateAuthority#s3_object_acl}.
	S3ObjectAcl *string `json:"s3ObjectAcl" yaml:"s3ObjectAcl"`
}

type AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference interface {
	cdktf.ComplexObject
	CustomCname() *string
	SetCustomCname(val *string)
	CustomCnameInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExpirationInDays() *float64
	SetExpirationInDays(val *float64)
	ExpirationInDaysInput() *float64
	InternalValue() *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration
	SetInternalValue(val *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3BucketNameInput() *string
	S3ObjectAcl() *string
	SetS3ObjectAcl(val *string)
	S3ObjectAclInput() *string
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
	ResetCustomCname()
	ResetEnabled()
	ResetS3BucketName()
	ResetS3ObjectAcl()
}

// The jsii proxy struct for AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference
type jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) CustomCname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) CustomCnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ExpirationInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ExpirationInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) InternalValue() *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration {
	var returns *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) S3ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) S3ObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference_Override(a AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) SetCustomCname(val *string) {
	_jsii_.Set(
		j,
		"customCname",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) SetExpirationInDays(val *float64) {
	_jsii_.Set(
		j,
		"expirationInDays",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) SetInternalValue(val *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) SetS3BucketName(val *string) {
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) SetS3ObjectAcl(val *string) {
	_jsii_.Set(
		j,
		"s3ObjectAcl",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ResetCustomCname() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomCname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ResetS3BucketName() {
	_jsii_.InvokeVoid(
		a,
		"resetS3BucketName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference) ResetS3ObjectAcl() {
	_jsii_.InvokeVoid(
		a,
		"resetS3ObjectAcl",
		nil, // no parameters
	)
}

type AcmpcaCertificateAuthorityRevocationConfigurationOutputReference interface {
	cdktf.ComplexObject
	CrlConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference
	CrlConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration
	InternalValue() *AcmpcaCertificateAuthorityRevocationConfiguration
	SetInternalValue(val *AcmpcaCertificateAuthorityRevocationConfiguration)
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
	PutCrlConfiguration(value *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration)
	ResetCrlConfiguration()
}

// The jsii proxy struct for AcmpcaCertificateAuthorityRevocationConfigurationOutputReference
type jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) CrlConfiguration() AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference {
	var returns AcmpcaCertificateAuthorityRevocationConfigurationCrlConfigurationOutputReference
	_jsii_.Get(
		j,
		"crlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) CrlConfigurationInput() *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration {
	var returns *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration
	_jsii_.Get(
		j,
		"crlConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) InternalValue() *AcmpcaCertificateAuthorityRevocationConfiguration {
	var returns *AcmpcaCertificateAuthorityRevocationConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAcmpcaCertificateAuthorityRevocationConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AcmpcaCertificateAuthorityRevocationConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityRevocationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateAuthorityRevocationConfigurationOutputReference_Override(a AcmpcaCertificateAuthorityRevocationConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityRevocationConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) SetInternalValue(val *AcmpcaCertificateAuthorityRevocationConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) PutCrlConfiguration(value *AcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration) {
	_jsii_.InvokeVoid(
		a,
		"putCrlConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityRevocationConfigurationOutputReference) ResetCrlConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCrlConfiguration",
		nil, // no parameters
	)
}

type AcmpcaCertificateAuthorityTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate_authority#create AcmpcaCertificateAuthority#create}.
	Create *string `json:"create" yaml:"create"`
}

type AcmpcaCertificateAuthorityTimeoutsOutputReference interface {
	cdktf.ComplexObject
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	InternalValue() *AcmpcaCertificateAuthorityTimeouts
	SetInternalValue(val *AcmpcaCertificateAuthorityTimeouts)
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
}

// The jsii proxy struct for AcmpcaCertificateAuthorityTimeoutsOutputReference
type jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) InternalValue() *AcmpcaCertificateAuthorityTimeouts {
	var returns *AcmpcaCertificateAuthorityTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAcmpcaCertificateAuthorityTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AcmpcaCertificateAuthorityTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateAuthorityTimeoutsOutputReference_Override(a AcmpcaCertificateAuthorityTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateAuthorityTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) SetInternalValue(val *AcmpcaCertificateAuthorityTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		a,
		"resetCreate",
		nil, // no parameters
	)
}

// Amazon Certificate Manager.
type AcmpcaCertificateConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate#certificate_authority_arn AcmpcaCertificate#certificate_authority_arn}.
	CertificateAuthorityArn *string `json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate#certificate_signing_request AcmpcaCertificate#certificate_signing_request}.
	CertificateSigningRequest *string `json:"certificateSigningRequest" yaml:"certificateSigningRequest"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate#signing_algorithm AcmpcaCertificate#signing_algorithm}.
	SigningAlgorithm *string `json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// validity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate#validity AcmpcaCertificate#validity}
	Validity *AcmpcaCertificateValidity `json:"validity" yaml:"validity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate#template_arn AcmpcaCertificate#template_arn}.
	TemplateArn *string `json:"templateArn" yaml:"templateArn"`
}

type AcmpcaCertificateValidity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate#type AcmpcaCertificate#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/acmpca_certificate#value AcmpcaCertificate#value}.
	Value *string `json:"value" yaml:"value"`
}

type AcmpcaCertificateValidityOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *AcmpcaCertificateValidity
	SetInternalValue(val *AcmpcaCertificateValidity)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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

// The jsii proxy struct for AcmpcaCertificateValidityOutputReference
type jsiiProxy_AcmpcaCertificateValidityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) InternalValue() *AcmpcaCertificateValidity {
	var returns *AcmpcaCertificateValidity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewAcmpcaCertificateValidityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AcmpcaCertificateValidityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AcmpcaCertificateValidityOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateValidityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateValidityOutputReference_Override(a AcmpcaCertificateValidityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.AcmpcaCertificateValidityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) SetInternalValue(val *AcmpcaCertificateValidity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateValidityOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AcmpcaCertificateValidityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/acm_certificate aws_acm_certificate}.
type DataAwsAcmCertificate interface {
	cdktf.TerraformDataSource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KeyTypes() *[]*string
	SetKeyTypes(val *[]*string)
	KeyTypesInput() *[]*string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MostRecent() interface{}
	SetMostRecent(val interface{})
	MostRecentInput() interface{}
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	Statuses() *[]*string
	SetStatuses(val *[]*string)
	StatusesInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Types() *[]*string
	SetTypes(val *[]*string)
	TypesInput() *[]*string
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
	ResetKeyTypes()
	ResetMostRecent()
	ResetOverrideLogicalId()
	ResetStatuses()
	ResetTags()
	ResetTypes()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsAcmCertificate
type jsiiProxy_DataAwsAcmCertificate struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsAcmCertificate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) KeyTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) KeyTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keyTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) MostRecent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mostRecent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) MostRecentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mostRecentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Statuses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statuses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) StatusesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmCertificate) TypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"typesInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/acm_certificate aws_acm_certificate} Data Source.
func NewDataAwsAcmCertificate(scope constructs.Construct, id *string, config *DataAwsAcmCertificateConfig) DataAwsAcmCertificate {
	_init_.Initialize()

	j := jsiiProxy_DataAwsAcmCertificate{}

	_jsii_.Create(
		"hashicorp_aws.acm.DataAwsAcmCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/acm_certificate aws_acm_certificate} Data Source.
func NewDataAwsAcmCertificate_Override(d DataAwsAcmCertificate, scope constructs.Construct, id *string, config *DataAwsAcmCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.DataAwsAcmCertificate",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetDomain(val *string) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetKeyTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"keyTypes",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetMostRecent(val interface{}) {
	_jsii_.Set(
		j,
		"mostRecent",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetStatuses(val *[]*string) {
	_jsii_.Set(
		j,
		"statuses",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmCertificate) SetTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"types",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsAcmCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.acm.DataAwsAcmCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsAcmCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.acm.DataAwsAcmCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmCertificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsAcmCertificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsAcmCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsAcmCertificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmCertificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmCertificate) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsAcmCertificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsAcmCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsAcmCertificate) ResetKeyTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmCertificate) ResetMostRecent() {
	_jsii_.InvokeVoid(
		d,
		"resetMostRecent",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsAcmCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmCertificate) ResetStatuses() {
	_jsii_.InvokeVoid(
		d,
		"resetStatuses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmCertificate) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmCertificate) ResetTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmCertificate) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsAcmCertificate) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsAcmCertificate) ToString() *string {
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
func (d *jsiiProxy_DataAwsAcmCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon Certificate Manager.
type DataAwsAcmCertificateConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acm_certificate#domain DataAwsAcmCertificate#domain}.
	Domain *string `json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acm_certificate#key_types DataAwsAcmCertificate#key_types}.
	KeyTypes *[]*string `json:"keyTypes" yaml:"keyTypes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acm_certificate#most_recent DataAwsAcmCertificate#most_recent}.
	MostRecent interface{} `json:"mostRecent" yaml:"mostRecent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acm_certificate#statuses DataAwsAcmCertificate#statuses}.
	Statuses *[]*string `json:"statuses" yaml:"statuses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acm_certificate#tags DataAwsAcmCertificate#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acm_certificate#types DataAwsAcmCertificate#types}.
	Types *[]*string `json:"types" yaml:"types"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate aws_acmpca_certificate}.
type DataAwsAcmpcaCertificate interface {
	cdktf.TerraformDataSource
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	CertificateAuthorityArn() *string
	SetCertificateAuthorityArn(val *string)
	CertificateAuthorityArnInput() *string
	CertificateChain() *string
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

// The jsii proxy struct for DataAwsAcmpcaCertificate
type jsiiProxy_DataAwsAcmpcaCertificate struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) CertificateAuthorityArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate aws_acmpca_certificate} Data Source.
func NewDataAwsAcmpcaCertificate(scope constructs.Construct, id *string, config *DataAwsAcmpcaCertificateConfig) DataAwsAcmpcaCertificate {
	_init_.Initialize()

	j := jsiiProxy_DataAwsAcmpcaCertificate{}

	_jsii_.Create(
		"hashicorp_aws.acm.DataAwsAcmpcaCertificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate aws_acmpca_certificate} Data Source.
func NewDataAwsAcmpcaCertificate_Override(d DataAwsAcmpcaCertificate, scope constructs.Construct, id *string, config *DataAwsAcmpcaCertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.DataAwsAcmpcaCertificate",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) SetCertificateAuthorityArn(val *string) {
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificate) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsAcmpcaCertificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.acm.DataAwsAcmpcaCertificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsAcmpcaCertificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.acm.DataAwsAcmpcaCertificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificate) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificate) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificate) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificate) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificate) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmpcaCertificate) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificate) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificate) ToString() *string {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate_authority aws_acmpca_certificate_authority}.
type DataAwsAcmpcaCertificateAuthority interface {
	cdktf.TerraformDataSource
	Arn() *string
	SetArn(val *string)
	ArnInput() *string
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	CertificateChain() *string
	CertificateSigningRequest() *string
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
	NotAfter() *string
	NotBefore() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RevocationConfiguration() interface{}
	SetRevocationConfiguration(val interface{})
	RevocationConfigurationInput() interface{}
	Serial() *string
	Status() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Type() *string
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
	ResetRevocationConfiguration()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsAcmpcaCertificateAuthority
type jsiiProxy_DataAwsAcmpcaCertificateAuthority struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) CertificateSigningRequest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSigningRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) NotAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) NotBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) RevocationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revocationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) RevocationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revocationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Serial() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate_authority aws_acmpca_certificate_authority} Data Source.
func NewDataAwsAcmpcaCertificateAuthority(scope constructs.Construct, id *string, config *DataAwsAcmpcaCertificateAuthorityConfig) DataAwsAcmpcaCertificateAuthority {
	_init_.Initialize()

	j := jsiiProxy_DataAwsAcmpcaCertificateAuthority{}

	_jsii_.Create(
		"hashicorp_aws.acm.DataAwsAcmpcaCertificateAuthority",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate_authority aws_acmpca_certificate_authority} Data Source.
func NewDataAwsAcmpcaCertificateAuthority_Override(d DataAwsAcmpcaCertificateAuthority, scope constructs.Construct, id *string, config *DataAwsAcmpcaCertificateAuthorityConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.acm.DataAwsAcmpcaCertificateAuthority",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) SetArn(val *string) {
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) SetRevocationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"revocationConfiguration",
		val,
	)
}

func (j *jsiiProxy_DataAwsAcmpcaCertificateAuthority) SetTags(val *map[string]*string) {
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
func DataAwsAcmpcaCertificateAuthority_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.acm.DataAwsAcmpcaCertificateAuthority",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsAcmpcaCertificateAuthority_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.acm.DataAwsAcmpcaCertificateAuthority",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) ResetRevocationConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetRevocationConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) ToString() *string {
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
func (d *jsiiProxy_DataAwsAcmpcaCertificateAuthority) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Amazon Certificate Manager.
type DataAwsAcmpcaCertificateAuthorityConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate_authority#arn DataAwsAcmpcaCertificateAuthority#arn}.
	Arn *string `json:"arn" yaml:"arn"`
	// revocation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate_authority#revocation_configuration DataAwsAcmpcaCertificateAuthority#revocation_configuration}
	RevocationConfiguration interface{} `json:"revocationConfiguration" yaml:"revocationConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate_authority#tags DataAwsAcmpcaCertificateAuthority#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type DataAwsAcmpcaCertificateAuthorityRevocationConfiguration struct {
	// crl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate_authority#crl_configuration DataAwsAcmpcaCertificateAuthority#crl_configuration}
	CrlConfiguration interface{} `json:"crlConfiguration" yaml:"crlConfiguration"`
}

type DataAwsAcmpcaCertificateAuthorityRevocationConfigurationCrlConfiguration struct {
}

// Amazon Certificate Manager.
type DataAwsAcmpcaCertificateConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate#arn DataAwsAcmpcaCertificate#arn}.
	Arn *string `json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/acmpca_certificate#certificate_authority_arn DataAwsAcmpcaCertificate#certificate_authority_arn}.
	CertificateAuthorityArn *string `json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
}
