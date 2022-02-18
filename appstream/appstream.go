package appstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/appstream/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_directory_config aws_appstream_directory_config}.
type AppstreamDirectoryConfig interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DirectoryName() *string
	SetDirectoryName(val *string)
	DirectoryNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	OrganizationalUnitDistinguishedNames() *[]*string
	SetOrganizationalUnitDistinguishedNames(val *[]*string)
	OrganizationalUnitDistinguishedNamesInput() *[]*string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ServiceAccountCredentials() AppstreamDirectoryConfigServiceAccountCredentialsOutputReference
	ServiceAccountCredentialsInput() *AppstreamDirectoryConfigServiceAccountCredentials
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
	PutServiceAccountCredentials(value *AppstreamDirectoryConfigServiceAccountCredentials)
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamDirectoryConfig
type jsiiProxy_AppstreamDirectoryConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamDirectoryConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) DirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) OrganizationalUnitDistinguishedNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) OrganizationalUnitDistinguishedNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) ServiceAccountCredentials() AppstreamDirectoryConfigServiceAccountCredentialsOutputReference {
	var returns AppstreamDirectoryConfigServiceAccountCredentialsOutputReference
	_jsii_.Get(
		j,
		"serviceAccountCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) ServiceAccountCredentialsInput() *AppstreamDirectoryConfigServiceAccountCredentials {
	var returns *AppstreamDirectoryConfigServiceAccountCredentials
	_jsii_.Get(
		j,
		"serviceAccountCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_directory_config aws_appstream_directory_config} Resource.
func NewAppstreamDirectoryConfig(scope constructs.Construct, id *string, config *AppstreamDirectoryConfigConfig) AppstreamDirectoryConfig {
	_init_.Initialize()

	j := jsiiProxy_AppstreamDirectoryConfig{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamDirectoryConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_directory_config aws_appstream_directory_config} Resource.
func NewAppstreamDirectoryConfig_Override(a AppstreamDirectoryConfig, scope constructs.Construct, id *string, config *AppstreamDirectoryConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamDirectoryConfig",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfig) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfig) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfig) SetDirectoryName(val *string) {
	_jsii_.Set(
		j,
		"directoryName",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfig) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfig) SetOrganizationalUnitDistinguishedNames(val *[]*string) {
	_jsii_.Set(
		j,
		"organizationalUnitDistinguishedNames",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfig) SetProvider(val cdktf.TerraformProvider) {
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
func AppstreamDirectoryConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appstream.AppstreamDirectoryConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamDirectoryConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appstream.AppstreamDirectoryConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamDirectoryConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamDirectoryConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamDirectoryConfig) PutServiceAccountCredentials(value *AppstreamDirectoryConfigServiceAccountCredentials) {
	_jsii_.InvokeVoid(
		a,
		"putServiceAccountCredentials",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamDirectoryConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamDirectoryConfig) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) ToString() *string {
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
func (a *jsiiProxy_AppstreamDirectoryConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppStream.
type AppstreamDirectoryConfigConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_directory_config#directory_name AppstreamDirectoryConfig#directory_name}.
	DirectoryName *string `json:"directoryName" yaml:"directoryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_directory_config#organizational_unit_distinguished_names AppstreamDirectoryConfig#organizational_unit_distinguished_names}.
	OrganizationalUnitDistinguishedNames *[]*string `json:"organizationalUnitDistinguishedNames" yaml:"organizationalUnitDistinguishedNames"`
	// service_account_credentials block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_directory_config#service_account_credentials AppstreamDirectoryConfig#service_account_credentials}
	ServiceAccountCredentials *AppstreamDirectoryConfigServiceAccountCredentials `json:"serviceAccountCredentials" yaml:"serviceAccountCredentials"`
}

type AppstreamDirectoryConfigServiceAccountCredentials struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_directory_config#account_name AppstreamDirectoryConfig#account_name}.
	AccountName *string `json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_directory_config#account_password AppstreamDirectoryConfig#account_password}.
	AccountPassword *string `json:"accountPassword" yaml:"accountPassword"`
}

type AppstreamDirectoryConfigServiceAccountCredentialsOutputReference interface {
	cdktf.ComplexObject
	AccountName() *string
	SetAccountName(val *string)
	AccountNameInput() *string
	AccountPassword() *string
	SetAccountPassword(val *string)
	AccountPasswordInput() *string
	InternalValue() *AppstreamDirectoryConfigServiceAccountCredentials
	SetInternalValue(val *AppstreamDirectoryConfigServiceAccountCredentials)
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

// The jsii proxy struct for AppstreamDirectoryConfigServiceAccountCredentialsOutputReference
type jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) AccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) AccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) AccountPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) AccountPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) InternalValue() *AppstreamDirectoryConfigServiceAccountCredentials {
	var returns *AppstreamDirectoryConfigServiceAccountCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamDirectoryConfigServiceAccountCredentialsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppstreamDirectoryConfigServiceAccountCredentialsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamDirectoryConfigServiceAccountCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamDirectoryConfigServiceAccountCredentialsOutputReference_Override(a AppstreamDirectoryConfigServiceAccountCredentialsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamDirectoryConfigServiceAccountCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) SetAccountName(val *string) {
	_jsii_.Set(
		j,
		"accountName",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) SetAccountPassword(val *string) {
	_jsii_.Set(
		j,
		"accountPassword",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) SetInternalValue(val *AppstreamDirectoryConfigServiceAccountCredentials) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamDirectoryConfigServiceAccountCredentialsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet aws_appstream_fleet}.
type AppstreamFleet interface {
	cdktf.TerraformResource
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ComputeCapacity() AppstreamFleetComputeCapacityOutputReference
	ComputeCapacityInput() *AppstreamFleetComputeCapacity
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisconnectTimeoutInSeconds() *float64
	SetDisconnectTimeoutInSeconds(val *float64)
	DisconnectTimeoutInSecondsInput() *float64
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DomainJoinInfo() AppstreamFleetDomainJoinInfoOutputReference
	DomainJoinInfoInput() *AppstreamFleetDomainJoinInfo
	EnableDefaultInternetAccess() interface{}
	SetEnableDefaultInternetAccess(val interface{})
	EnableDefaultInternetAccessInput() interface{}
	FleetType() *string
	SetFleetType(val *string)
	FleetTypeInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	Id() *string
	IdleDisconnectTimeoutInSeconds() *float64
	SetIdleDisconnectTimeoutInSeconds(val *float64)
	IdleDisconnectTimeoutInSecondsInput() *float64
	ImageArn() *string
	SetImageArn(val *string)
	ImageArnInput() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxUserDurationInSeconds() *float64
	SetMaxUserDurationInSeconds(val *float64)
	MaxUserDurationInSecondsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	State() *string
	StreamView() *string
	SetStreamView(val *string)
	StreamViewInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VpcConfig() AppstreamFleetVpcConfigOutputReference
	VpcConfigInput() *AppstreamFleetVpcConfig
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
	PutComputeCapacity(value *AppstreamFleetComputeCapacity)
	PutDomainJoinInfo(value *AppstreamFleetDomainJoinInfo)
	PutVpcConfig(value *AppstreamFleetVpcConfig)
	ResetDescription()
	ResetDisconnectTimeoutInSeconds()
	ResetDisplayName()
	ResetDomainJoinInfo()
	ResetEnableDefaultInternetAccess()
	ResetFleetType()
	ResetIamRoleArn()
	ResetIdleDisconnectTimeoutInSeconds()
	ResetImageArn()
	ResetImageName()
	ResetMaxUserDurationInSeconds()
	ResetOverrideLogicalId()
	ResetStreamView()
	ResetTags()
	ResetTagsAll()
	ResetVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamFleet
type jsiiProxy_AppstreamFleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamFleet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ComputeCapacity() AppstreamFleetComputeCapacityOutputReference {
	var returns AppstreamFleetComputeCapacityOutputReference
	_jsii_.Get(
		j,
		"computeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ComputeCapacityInput() *AppstreamFleetComputeCapacity {
	var returns *AppstreamFleetComputeCapacity
	_jsii_.Get(
		j,
		"computeCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisconnectTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisconnectTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"disconnectTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DomainJoinInfo() AppstreamFleetDomainJoinInfoOutputReference {
	var returns AppstreamFleetDomainJoinInfoOutputReference
	_jsii_.Get(
		j,
		"domainJoinInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) DomainJoinInfoInput() *AppstreamFleetDomainJoinInfo {
	var returns *AppstreamFleetDomainJoinInfo
	_jsii_.Get(
		j,
		"domainJoinInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) EnableDefaultInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) EnableDefaultInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) FleetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IdleDisconnectTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) IdleDisconnectTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleDisconnectTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxUserDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) MaxUserDurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUserDurationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) StreamView() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamView",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) StreamViewInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamViewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) VpcConfig() AppstreamFleetVpcConfigOutputReference {
	var returns AppstreamFleetVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleet) VpcConfigInput() *AppstreamFleetVpcConfig {
	var returns *AppstreamFleetVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet aws_appstream_fleet} Resource.
func NewAppstreamFleet(scope constructs.Construct, id *string, config *AppstreamFleetConfig) AppstreamFleet {
	_init_.Initialize()

	j := jsiiProxy_AppstreamFleet{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet aws_appstream_fleet} Resource.
func NewAppstreamFleet_Override(a AppstreamFleet, scope constructs.Construct, id *string, config *AppstreamFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleet",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetDisconnectTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"disconnectTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetEnableDefaultInternetAccess(val interface{}) {
	_jsii_.Set(
		j,
		"enableDefaultInternetAccess",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetFleetType(val *string) {
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetIdleDisconnectTimeoutInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"idleDisconnectTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetImageArn(val *string) {
	_jsii_.Set(
		j,
		"imageArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetMaxUserDurationInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxUserDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetStreamView(val *string) {
	_jsii_.Set(
		j,
		"streamView",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleet) SetTagsAll(val *map[string]*string) {
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
func AppstreamFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appstream.AppstreamFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appstream.AppstreamFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleet) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamFleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamFleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamFleet) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamFleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamFleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamFleet) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamFleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamFleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleet) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutComputeCapacity(value *AppstreamFleetComputeCapacity) {
	_jsii_.InvokeVoid(
		a,
		"putComputeCapacity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutDomainJoinInfo(value *AppstreamFleetDomainJoinInfo) {
	_jsii_.InvokeVoid(
		a,
		"putDomainJoinInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) PutVpcConfig(value *AppstreamFleetVpcConfig) {
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDisconnectTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetDisconnectTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetDomainJoinInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainJoinInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetEnableDefaultInternetAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDefaultInternetAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetFleetType() {
	_jsii_.InvokeVoid(
		a,
		"resetFleetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetIamRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetIdleDisconnectTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleDisconnectTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetImageArn() {
	_jsii_.InvokeVoid(
		a,
		"resetImageArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetImageName() {
	_jsii_.InvokeVoid(
		a,
		"resetImageName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetMaxUserDurationInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxUserDurationInSeconds",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetStreamView() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamView",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleet) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamFleet) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppstreamFleet) ToString() *string {
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
func (a *jsiiProxy_AppstreamFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppstreamFleetComputeCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#desired_instances AppstreamFleet#desired_instances}.
	DesiredInstances *float64 `json:"desiredInstances" yaml:"desiredInstances"`
}

type AppstreamFleetComputeCapacityOutputReference interface {
	cdktf.ComplexObject
	Available() *float64
	DesiredInstances() *float64
	SetDesiredInstances(val *float64)
	DesiredInstancesInput() *float64
	InternalValue() *AppstreamFleetComputeCapacity
	SetInternalValue(val *AppstreamFleetComputeCapacity)
	InUse() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Running() *float64
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

// The jsii proxy struct for AppstreamFleetComputeCapacityOutputReference
type jsiiProxy_AppstreamFleetComputeCapacityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) Available() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"available",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) DesiredInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) DesiredInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) InternalValue() *AppstreamFleetComputeCapacity {
	var returns *AppstreamFleetComputeCapacity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) InUse() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inUse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) Running() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"running",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamFleetComputeCapacityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppstreamFleetComputeCapacityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamFleetComputeCapacityOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleetComputeCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamFleetComputeCapacityOutputReference_Override(a AppstreamFleetComputeCapacityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleetComputeCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) SetDesiredInstances(val *float64) {
	_jsii_.Set(
		j,
		"desiredInstances",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) SetInternalValue(val *AppstreamFleetComputeCapacity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleetComputeCapacityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// AppStream.
type AppstreamFleetConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// compute_capacity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#compute_capacity AppstreamFleet#compute_capacity}
	ComputeCapacity *AppstreamFleetComputeCapacity `json:"computeCapacity" yaml:"computeCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#instance_type AppstreamFleet#instance_type}.
	InstanceType *string `json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#name AppstreamFleet#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#description AppstreamFleet#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#disconnect_timeout_in_seconds AppstreamFleet#disconnect_timeout_in_seconds}.
	DisconnectTimeoutInSeconds *float64 `json:"disconnectTimeoutInSeconds" yaml:"disconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#display_name AppstreamFleet#display_name}.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// domain_join_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#domain_join_info AppstreamFleet#domain_join_info}
	DomainJoinInfo *AppstreamFleetDomainJoinInfo `json:"domainJoinInfo" yaml:"domainJoinInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#enable_default_internet_access AppstreamFleet#enable_default_internet_access}.
	EnableDefaultInternetAccess interface{} `json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#fleet_type AppstreamFleet#fleet_type}.
	FleetType *string `json:"fleetType" yaml:"fleetType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#iam_role_arn AppstreamFleet#iam_role_arn}.
	IamRoleArn *string `json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#idle_disconnect_timeout_in_seconds AppstreamFleet#idle_disconnect_timeout_in_seconds}.
	IdleDisconnectTimeoutInSeconds *float64 `json:"idleDisconnectTimeoutInSeconds" yaml:"idleDisconnectTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#image_arn AppstreamFleet#image_arn}.
	ImageArn *string `json:"imageArn" yaml:"imageArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#image_name AppstreamFleet#image_name}.
	ImageName *string `json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#max_user_duration_in_seconds AppstreamFleet#max_user_duration_in_seconds}.
	MaxUserDurationInSeconds *float64 `json:"maxUserDurationInSeconds" yaml:"maxUserDurationInSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#stream_view AppstreamFleet#stream_view}.
	StreamView *string `json:"streamView" yaml:"streamView"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#tags AppstreamFleet#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#tags_all AppstreamFleet#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#vpc_config AppstreamFleet#vpc_config}
	VpcConfig *AppstreamFleetVpcConfig `json:"vpcConfig" yaml:"vpcConfig"`
}

type AppstreamFleetDomainJoinInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#directory_name AppstreamFleet#directory_name}.
	DirectoryName *string `json:"directoryName" yaml:"directoryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#organizational_unit_distinguished_name AppstreamFleet#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

type AppstreamFleetDomainJoinInfoOutputReference interface {
	cdktf.ComplexObject
	DirectoryName() *string
	SetDirectoryName(val *string)
	DirectoryNameInput() *string
	InternalValue() *AppstreamFleetDomainJoinInfo
	SetInternalValue(val *AppstreamFleetDomainJoinInfo)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OrganizationalUnitDistinguishedName() *string
	SetOrganizationalUnitDistinguishedName(val *string)
	OrganizationalUnitDistinguishedNameInput() *string
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
	ResetDirectoryName()
	ResetOrganizationalUnitDistinguishedName()
}

// The jsii proxy struct for AppstreamFleetDomainJoinInfoOutputReference
type jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) DirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) InternalValue() *AppstreamFleetDomainJoinInfo {
	var returns *AppstreamFleetDomainJoinInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) OrganizationalUnitDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) OrganizationalUnitDistinguishedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamFleetDomainJoinInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppstreamFleetDomainJoinInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleetDomainJoinInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamFleetDomainJoinInfoOutputReference_Override(a AppstreamFleetDomainJoinInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleetDomainJoinInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetDirectoryName(val *string) {
	_jsii_.Set(
		j,
		"directoryName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetInternalValue(val *AppstreamFleetDomainJoinInfo) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetOrganizationalUnitDistinguishedName(val *string) {
	_jsii_.Set(
		j,
		"organizationalUnitDistinguishedName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) ResetDirectoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleetDomainJoinInfoOutputReference) ResetOrganizationalUnitDistinguishedName() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganizationalUnitDistinguishedName",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet_stack_association aws_appstream_fleet_stack_association}.
type AppstreamFleetStackAssociation interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	FleetName() *string
	SetFleetName(val *string)
	FleetNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StackName() *string
	SetStackName(val *string)
	StackNameInput() *string
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

// The jsii proxy struct for AppstreamFleetStackAssociation
type jsiiProxy_AppstreamFleetStackAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) FleetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) FleetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) StackNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet_stack_association aws_appstream_fleet_stack_association} Resource.
func NewAppstreamFleetStackAssociation(scope constructs.Construct, id *string, config *AppstreamFleetStackAssociationConfig) AppstreamFleetStackAssociation {
	_init_.Initialize()

	j := jsiiProxy_AppstreamFleetStackAssociation{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleetStackAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet_stack_association aws_appstream_fleet_stack_association} Resource.
func NewAppstreamFleetStackAssociation_Override(a AppstreamFleetStackAssociation, scope constructs.Construct, id *string, config *AppstreamFleetStackAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleetStackAssociation",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) SetFleetName(val *string) {
	_jsii_.Set(
		j,
		"fleetName",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetStackAssociation) SetStackName(val *string) {
	_jsii_.Set(
		j,
		"stackName",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AppstreamFleetStackAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appstream.AppstreamFleetStackAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamFleetStackAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appstream.AppstreamFleetStackAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetStackAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetStackAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamFleetStackAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleetStackAssociation) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) ToString() *string {
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
func (a *jsiiProxy_AppstreamFleetStackAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppStream.
type AppstreamFleetStackAssociationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet_stack_association#fleet_name AppstreamFleetStackAssociation#fleet_name}.
	FleetName *string `json:"fleetName" yaml:"fleetName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet_stack_association#stack_name AppstreamFleetStackAssociation#stack_name}.
	StackName *string `json:"stackName" yaml:"stackName"`
}

type AppstreamFleetVpcConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#security_group_ids AppstreamFleet#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_fleet#subnet_ids AppstreamFleet#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

type AppstreamFleetVpcConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *AppstreamFleetVpcConfig
	SetInternalValue(val *AppstreamFleetVpcConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	ResetSecurityGroupIds()
	ResetSubnetIds()
}

// The jsii proxy struct for AppstreamFleetVpcConfigOutputReference
type jsiiProxy_AppstreamFleetVpcConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) InternalValue() *AppstreamFleetVpcConfig {
	var returns *AppstreamFleetVpcConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamFleetVpcConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppstreamFleetVpcConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamFleetVpcConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleetVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamFleetVpcConfigOutputReference_Override(a AppstreamFleetVpcConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamFleetVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetInternalValue(val *AppstreamFleetVpcConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamFleetVpcConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamFleetVpcConfigOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetIds",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder aws_appstream_image_builder}.
type AppstreamImageBuilder interface {
	cdktf.TerraformResource
	AccessEndpoint() interface{}
	SetAccessEndpoint(val interface{})
	AccessEndpointInput() interface{}
	AppstreamAgentVersion() *string
	SetAppstreamAgentVersion(val *string)
	AppstreamAgentVersionInput() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	DomainJoinInfo() AppstreamImageBuilderDomainJoinInfoOutputReference
	DomainJoinInfoInput() *AppstreamImageBuilderDomainJoinInfo
	EnableDefaultInternetAccess() interface{}
	SetEnableDefaultInternetAccess(val interface{})
	EnableDefaultInternetAccessInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	IamRoleArn() *string
	SetIamRoleArn(val *string)
	IamRoleArnInput() *string
	Id() *string
	ImageArn() *string
	SetImageArn(val *string)
	ImageArnInput() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	State() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	VpcConfig() AppstreamImageBuilderVpcConfigOutputReference
	VpcConfigInput() *AppstreamImageBuilderVpcConfig
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
	PutDomainJoinInfo(value *AppstreamImageBuilderDomainJoinInfo)
	PutVpcConfig(value *AppstreamImageBuilderVpcConfig)
	ResetAccessEndpoint()
	ResetAppstreamAgentVersion()
	ResetDescription()
	ResetDisplayName()
	ResetDomainJoinInfo()
	ResetEnableDefaultInternetAccess()
	ResetIamRoleArn()
	ResetImageArn()
	ResetImageName()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	ResetVpcConfig()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamImageBuilder
type jsiiProxy_AppstreamImageBuilder struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamImageBuilder) AccessEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) AccessEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) AppstreamAgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appstreamAgentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) AppstreamAgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appstreamAgentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DomainJoinInfo() AppstreamImageBuilderDomainJoinInfoOutputReference {
	var returns AppstreamImageBuilderDomainJoinInfoOutputReference
	_jsii_.Get(
		j,
		"domainJoinInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) DomainJoinInfoInput() *AppstreamImageBuilderDomainJoinInfo {
	var returns *AppstreamImageBuilderDomainJoinInfo
	_jsii_.Get(
		j,
		"domainJoinInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) EnableDefaultInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) EnableDefaultInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDefaultInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) IamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) IamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) VpcConfig() AppstreamImageBuilderVpcConfigOutputReference {
	var returns AppstreamImageBuilderVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilder) VpcConfigInput() *AppstreamImageBuilderVpcConfig {
	var returns *AppstreamImageBuilderVpcConfig
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder aws_appstream_image_builder} Resource.
func NewAppstreamImageBuilder(scope constructs.Construct, id *string, config *AppstreamImageBuilderConfig) AppstreamImageBuilder {
	_init_.Initialize()

	j := jsiiProxy_AppstreamImageBuilder{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamImageBuilder",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder aws_appstream_image_builder} Resource.
func NewAppstreamImageBuilder_Override(a AppstreamImageBuilder, scope constructs.Construct, id *string, config *AppstreamImageBuilderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamImageBuilder",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetAccessEndpoint(val interface{}) {
	_jsii_.Set(
		j,
		"accessEndpoint",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetAppstreamAgentVersion(val *string) {
	_jsii_.Set(
		j,
		"appstreamAgentVersion",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetEnableDefaultInternetAccess(val interface{}) {
	_jsii_.Set(
		j,
		"enableDefaultInternetAccess",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"iamRoleArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetImageArn(val *string) {
	_jsii_.Set(
		j,
		"imageArn",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetImageName(val *string) {
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilder) SetTagsAll(val *map[string]*string) {
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
func AppstreamImageBuilder_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appstream.AppstreamImageBuilder",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamImageBuilder_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appstream.AppstreamImageBuilder",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamImageBuilder) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamImageBuilder) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamImageBuilder) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamImageBuilder) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamImageBuilder) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamImageBuilder) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamImageBuilder) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamImageBuilder) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamImageBuilder) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamImageBuilder) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) PutDomainJoinInfo(value *AppstreamImageBuilderDomainJoinInfo) {
	_jsii_.InvokeVoid(
		a,
		"putDomainJoinInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) PutVpcConfig(value *AppstreamImageBuilderVpcConfig) {
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetAccessEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetAppstreamAgentVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetAppstreamAgentVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetDomainJoinInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainJoinInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetEnableDefaultInternetAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDefaultInternetAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetIamRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetImageArn() {
	_jsii_.InvokeVoid(
		a,
		"resetImageArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetImageName() {
	_jsii_.InvokeVoid(
		a,
		"resetImageName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamImageBuilder) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilder) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamImageBuilder) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppstreamImageBuilder) ToString() *string {
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
func (a *jsiiProxy_AppstreamImageBuilder) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppstreamImageBuilderAccessEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#endpoint_type AppstreamImageBuilder#endpoint_type}.
	EndpointType *string `json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#vpce_id AppstreamImageBuilder#vpce_id}.
	VpceId *string `json:"vpceId" yaml:"vpceId"`
}

// AppStream.
type AppstreamImageBuilderConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#instance_type AppstreamImageBuilder#instance_type}.
	InstanceType *string `json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#name AppstreamImageBuilder#name}.
	Name *string `json:"name" yaml:"name"`
	// access_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#access_endpoint AppstreamImageBuilder#access_endpoint}
	AccessEndpoint interface{} `json:"accessEndpoint" yaml:"accessEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#appstream_agent_version AppstreamImageBuilder#appstream_agent_version}.
	AppstreamAgentVersion *string `json:"appstreamAgentVersion" yaml:"appstreamAgentVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#description AppstreamImageBuilder#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#display_name AppstreamImageBuilder#display_name}.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// domain_join_info block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#domain_join_info AppstreamImageBuilder#domain_join_info}
	DomainJoinInfo *AppstreamImageBuilderDomainJoinInfo `json:"domainJoinInfo" yaml:"domainJoinInfo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#enable_default_internet_access AppstreamImageBuilder#enable_default_internet_access}.
	EnableDefaultInternetAccess interface{} `json:"enableDefaultInternetAccess" yaml:"enableDefaultInternetAccess"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#iam_role_arn AppstreamImageBuilder#iam_role_arn}.
	IamRoleArn *string `json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#image_arn AppstreamImageBuilder#image_arn}.
	ImageArn *string `json:"imageArn" yaml:"imageArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#image_name AppstreamImageBuilder#image_name}.
	ImageName *string `json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#tags AppstreamImageBuilder#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#tags_all AppstreamImageBuilder#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#vpc_config AppstreamImageBuilder#vpc_config}
	VpcConfig *AppstreamImageBuilderVpcConfig `json:"vpcConfig" yaml:"vpcConfig"`
}

type AppstreamImageBuilderDomainJoinInfo struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#directory_name AppstreamImageBuilder#directory_name}.
	DirectoryName *string `json:"directoryName" yaml:"directoryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#organizational_unit_distinguished_name AppstreamImageBuilder#organizational_unit_distinguished_name}.
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

type AppstreamImageBuilderDomainJoinInfoOutputReference interface {
	cdktf.ComplexObject
	DirectoryName() *string
	SetDirectoryName(val *string)
	DirectoryNameInput() *string
	InternalValue() *AppstreamImageBuilderDomainJoinInfo
	SetInternalValue(val *AppstreamImageBuilderDomainJoinInfo)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OrganizationalUnitDistinguishedName() *string
	SetOrganizationalUnitDistinguishedName(val *string)
	OrganizationalUnitDistinguishedNameInput() *string
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
	ResetDirectoryName()
	ResetOrganizationalUnitDistinguishedName()
}

// The jsii proxy struct for AppstreamImageBuilderDomainJoinInfoOutputReference
type jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) DirectoryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) InternalValue() *AppstreamImageBuilderDomainJoinInfo {
	var returns *AppstreamImageBuilderDomainJoinInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) OrganizationalUnitDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) OrganizationalUnitDistinguishedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamImageBuilderDomainJoinInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppstreamImageBuilderDomainJoinInfoOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamImageBuilderDomainJoinInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamImageBuilderDomainJoinInfoOutputReference_Override(a AppstreamImageBuilderDomainJoinInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamImageBuilderDomainJoinInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetDirectoryName(val *string) {
	_jsii_.Set(
		j,
		"directoryName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetInternalValue(val *AppstreamImageBuilderDomainJoinInfo) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetOrganizationalUnitDistinguishedName(val *string) {
	_jsii_.Set(
		j,
		"organizationalUnitDistinguishedName",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) ResetDirectoryName() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilderDomainJoinInfoOutputReference) ResetOrganizationalUnitDistinguishedName() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganizationalUnitDistinguishedName",
		nil, // no parameters
	)
}

type AppstreamImageBuilderVpcConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#security_group_ids AppstreamImageBuilder#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_image_builder#subnet_ids AppstreamImageBuilder#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

type AppstreamImageBuilderVpcConfigOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *AppstreamImageBuilderVpcConfig
	SetInternalValue(val *AppstreamImageBuilderVpcConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
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
	ResetSecurityGroupIds()
	ResetSubnetIds()
}

// The jsii proxy struct for AppstreamImageBuilderVpcConfigOutputReference
type jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) InternalValue() *AppstreamImageBuilderVpcConfig {
	var returns *AppstreamImageBuilderVpcConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamImageBuilderVpcConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppstreamImageBuilderVpcConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamImageBuilderVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamImageBuilderVpcConfigOutputReference_Override(a AppstreamImageBuilderVpcConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamImageBuilderVpcConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetInternalValue(val *AppstreamImageBuilderVpcConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamImageBuilderVpcConfigOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetIds",
		nil, // no parameters
	)
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack aws_appstream_stack}.
type AppstreamStack interface {
	cdktf.TerraformResource
	AccessEndpoints() interface{}
	SetAccessEndpoints(val interface{})
	AccessEndpointsInput() interface{}
	ApplicationSettings() AppstreamStackApplicationSettingsOutputReference
	ApplicationSettingsInput() *AppstreamStackApplicationSettings
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EmbedHostDomains() *[]*string
	SetEmbedHostDomains(val *[]*string)
	EmbedHostDomainsInput() *[]*string
	FeedbackUrl() *string
	SetFeedbackUrl(val *string)
	FeedbackUrlInput() *string
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
	RedirectUrl() *string
	SetRedirectUrl(val *string)
	RedirectUrlInput() *string
	StorageConnectors() interface{}
	SetStorageConnectors(val interface{})
	StorageConnectorsInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	UserSettings() interface{}
	SetUserSettings(val interface{})
	UserSettingsInput() interface{}
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
	PutApplicationSettings(value *AppstreamStackApplicationSettings)
	ResetAccessEndpoints()
	ResetApplicationSettings()
	ResetDescription()
	ResetDisplayName()
	ResetEmbedHostDomains()
	ResetFeedbackUrl()
	ResetOverrideLogicalId()
	ResetRedirectUrl()
	ResetStorageConnectors()
	ResetTags()
	ResetTagsAll()
	ResetUserSettings()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamStack
type jsiiProxy_AppstreamStack struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamStack) AccessEndpoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) AccessEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ApplicationSettings() AppstreamStackApplicationSettingsOutputReference {
	var returns AppstreamStackApplicationSettingsOutputReference
	_jsii_.Get(
		j,
		"applicationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ApplicationSettingsInput() *AppstreamStackApplicationSettings {
	var returns *AppstreamStackApplicationSettings
	_jsii_.Get(
		j,
		"applicationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) EmbedHostDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"embedHostDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) EmbedHostDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"embedHostDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) FeedbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedbackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) FeedbackUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedbackUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) RedirectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) RedirectUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) StorageConnectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConnectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) StorageConnectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConnectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) UserSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStack) UserSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSettingsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack aws_appstream_stack} Resource.
func NewAppstreamStack(scope constructs.Construct, id *string, config *AppstreamStackConfig) AppstreamStack {
	_init_.Initialize()

	j := jsiiProxy_AppstreamStack{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamStack",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack aws_appstream_stack} Resource.
func NewAppstreamStack_Override(a AppstreamStack, scope constructs.Construct, id *string, config *AppstreamStackConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamStack",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamStack) SetAccessEndpoints(val interface{}) {
	_jsii_.Set(
		j,
		"accessEndpoints",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetEmbedHostDomains(val *[]*string) {
	_jsii_.Set(
		j,
		"embedHostDomains",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetFeedbackUrl(val *string) {
	_jsii_.Set(
		j,
		"feedbackUrl",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetRedirectUrl(val *string) {
	_jsii_.Set(
		j,
		"redirectUrl",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetStorageConnectors(val interface{}) {
	_jsii_.Set(
		j,
		"storageConnectors",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetTagsAll(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AppstreamStack) SetUserSettings(val interface{}) {
	_jsii_.Set(
		j,
		"userSettings",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AppstreamStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appstream.AppstreamStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamStack_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appstream.AppstreamStack",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamStack) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamStack) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamStack) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamStack) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamStack) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamStack) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamStack) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamStack) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamStack) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamStack) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamStack) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamStack) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamStack) PutApplicationSettings(value *AppstreamStackApplicationSettings) {
	_jsii_.InvokeVoid(
		a,
		"putApplicationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppstreamStack) ResetAccessEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetApplicationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetEmbedHostDomains() {
	_jsii_.InvokeVoid(
		a,
		"resetEmbedHostDomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetFeedbackUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetFeedbackUrl",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamStack) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetRedirectUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetStorageConnectors() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageConnectors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) ResetUserSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetUserSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStack) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamStack) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppstreamStack) ToString() *string {
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
func (a *jsiiProxy_AppstreamStack) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type AppstreamStackAccessEndpoints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#endpoint_type AppstreamStack#endpoint_type}.
	EndpointType *string `json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#vpce_id AppstreamStack#vpce_id}.
	VpceId *string `json:"vpceId" yaml:"vpceId"`
}

type AppstreamStackApplicationSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#enabled AppstreamStack#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#settings_group AppstreamStack#settings_group}.
	SettingsGroup *string `json:"settingsGroup" yaml:"settingsGroup"`
}

type AppstreamStackApplicationSettingsOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *AppstreamStackApplicationSettings
	SetInternalValue(val *AppstreamStackApplicationSettings)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	SettingsGroup() *string
	SetSettingsGroup(val *string)
	SettingsGroupInput() *string
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
	ResetEnabled()
	ResetSettingsGroup()
}

// The jsii proxy struct for AppstreamStackApplicationSettingsOutputReference
type jsiiProxy_AppstreamStackApplicationSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) InternalValue() *AppstreamStackApplicationSettings {
	var returns *AppstreamStackApplicationSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SettingsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SettingsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"settingsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewAppstreamStackApplicationSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) AppstreamStackApplicationSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_AppstreamStackApplicationSettingsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamStackApplicationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewAppstreamStackApplicationSettingsOutputReference_Override(a AppstreamStackApplicationSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamStackApplicationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		a,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetInternalValue(val *AppstreamStackApplicationSettings) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetSettingsGroup(val *string) {
	_jsii_.Set(
		j,
		"settingsGroup",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamStackApplicationSettingsOutputReference) ResetSettingsGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetSettingsGroup",
		nil, // no parameters
	)
}

// AppStream.
type AppstreamStackConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#name AppstreamStack#name}.
	Name *string `json:"name" yaml:"name"`
	// access_endpoints block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#access_endpoints AppstreamStack#access_endpoints}
	AccessEndpoints interface{} `json:"accessEndpoints" yaml:"accessEndpoints"`
	// application_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#application_settings AppstreamStack#application_settings}
	ApplicationSettings *AppstreamStackApplicationSettings `json:"applicationSettings" yaml:"applicationSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#description AppstreamStack#description}.
	Description *string `json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#display_name AppstreamStack#display_name}.
	DisplayName *string `json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#embed_host_domains AppstreamStack#embed_host_domains}.
	EmbedHostDomains *[]*string `json:"embedHostDomains" yaml:"embedHostDomains"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#feedback_url AppstreamStack#feedback_url}.
	FeedbackUrl *string `json:"feedbackUrl" yaml:"feedbackUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#redirect_url AppstreamStack#redirect_url}.
	RedirectUrl *string `json:"redirectUrl" yaml:"redirectUrl"`
	// storage_connectors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#storage_connectors AppstreamStack#storage_connectors}
	StorageConnectors interface{} `json:"storageConnectors" yaml:"storageConnectors"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#tags AppstreamStack#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#tags_all AppstreamStack#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// user_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#user_settings AppstreamStack#user_settings}
	UserSettings interface{} `json:"userSettings" yaml:"userSettings"`
}

type AppstreamStackStorageConnectors struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#connector_type AppstreamStack#connector_type}.
	ConnectorType *string `json:"connectorType" yaml:"connectorType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#domains AppstreamStack#domains}.
	Domains *[]*string `json:"domains" yaml:"domains"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#resource_identifier AppstreamStack#resource_identifier}.
	ResourceIdentifier *string `json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

type AppstreamStackUserSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#action AppstreamStack#action}.
	Action *string `json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_stack#permission AppstreamStack#permission}.
	Permission *string `json:"permission" yaml:"permission"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_user aws_appstream_user}.
type AppstreamUser interface {
	cdktf.TerraformResource
	Arn() *string
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedTime() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	LastName() *string
	SetLastName(val *string)
	LastNameInput() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SendEmailNotification() interface{}
	SetSendEmailNotification(val interface{})
	SendEmailNotificationInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
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
	ResetEnabled()
	ResetFirstName()
	ResetLastName()
	ResetOverrideLogicalId()
	ResetSendEmailNotification()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamUser
type jsiiProxy_AppstreamUser struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamUser) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) SendEmailNotification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) SendEmailNotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUser) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_user aws_appstream_user} Resource.
func NewAppstreamUser(scope constructs.Construct, id *string, config *AppstreamUserConfig) AppstreamUser {
	_init_.Initialize()

	j := jsiiProxy_AppstreamUser{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_user aws_appstream_user} Resource.
func NewAppstreamUser_Override(a AppstreamUser, scope constructs.Construct, id *string, config *AppstreamUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamUser",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamUser) SetAuthenticationType(val *string) {
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_AppstreamUser) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamUser) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamUser) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AppstreamUser) SetFirstName(val *string) {
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_AppstreamUser) SetLastName(val *string) {
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_AppstreamUser) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamUser) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamUser) SetSendEmailNotification(val interface{}) {
	_jsii_.Set(
		j,
		"sendEmailNotification",
		val,
	)
}

func (j *jsiiProxy_AppstreamUser) SetUserName(val *string) {
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
func AppstreamUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appstream.AppstreamUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appstream.AppstreamUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamUser) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamUser) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamUser) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamUser) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppstreamUser) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamUser) ResetFirstName() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamUser) ResetLastName() {
	_jsii_.InvokeVoid(
		a,
		"resetLastName",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamUser) ResetSendEmailNotification() {
	_jsii_.InvokeVoid(
		a,
		"resetSendEmailNotification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamUser) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamUser) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppstreamUser) ToString() *string {
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
func (a *jsiiProxy_AppstreamUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppStream.
type AppstreamUserConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user#authentication_type AppstreamUser#authentication_type}.
	AuthenticationType *string `json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user#user_name AppstreamUser#user_name}.
	UserName *string `json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user#enabled AppstreamUser#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user#first_name AppstreamUser#first_name}.
	FirstName *string `json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user#last_name AppstreamUser#last_name}.
	LastName *string `json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user#send_email_notification AppstreamUser#send_email_notification}.
	SendEmailNotification interface{} `json:"sendEmailNotification" yaml:"sendEmailNotification"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/appstream_user_stack_association aws_appstream_user_stack_association}.
type AppstreamUserStackAssociation interface {
	cdktf.TerraformResource
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
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SendEmailNotification() interface{}
	SetSendEmailNotification(val interface{})
	SendEmailNotificationInput() interface{}
	StackName() *string
	SetStackName(val *string)
	StackNameInput() *string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
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
	ResetSendEmailNotification()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for AppstreamUserStackAssociation
type jsiiProxy_AppstreamUserStackAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppstreamUserStackAssociation) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SendEmailNotification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SendEmailNotificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendEmailNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) StackNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppstreamUserStackAssociation) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_user_stack_association aws_appstream_user_stack_association} Resource.
func NewAppstreamUserStackAssociation(scope constructs.Construct, id *string, config *AppstreamUserStackAssociationConfig) AppstreamUserStackAssociation {
	_init_.Initialize()

	j := jsiiProxy_AppstreamUserStackAssociation{}

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamUserStackAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/appstream_user_stack_association aws_appstream_user_stack_association} Resource.
func NewAppstreamUserStackAssociation_Override(a AppstreamUserStackAssociation, scope constructs.Construct, id *string, config *AppstreamUserStackAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.appstream.AppstreamUserStackAssociation",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SetAuthenticationType(val *string) {
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SetSendEmailNotification(val interface{}) {
	_jsii_.Set(
		j,
		"sendEmailNotification",
		val,
	)
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SetStackName(val *string) {
	_jsii_.Set(
		j,
		"stackName",
		val,
	)
}

func (j *jsiiProxy_AppstreamUserStackAssociation) SetUserName(val *string) {
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
func AppstreamUserStackAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.appstream.AppstreamUserStackAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppstreamUserStackAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.appstream.AppstreamUserStackAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (a *jsiiProxy_AppstreamUserStackAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (a *jsiiProxy_AppstreamUserStackAssociation) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) GetStringAttribute(terraformAttribute *string) *string {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (a *jsiiProxy_AppstreamUserStackAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamUserStackAssociation) ResetSendEmailNotification() {
	_jsii_.InvokeVoid(
		a,
		"resetSendEmailNotification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppstreamUserStackAssociation) SynthesizeAttributes() *map[string]interface{} {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) ToMetadata() interface{} {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) ToString() *string {
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
func (a *jsiiProxy_AppstreamUserStackAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AppStream.
type AppstreamUserStackAssociationConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user_stack_association#authentication_type AppstreamUserStackAssociation#authentication_type}.
	AuthenticationType *string `json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user_stack_association#stack_name AppstreamUserStackAssociation#stack_name}.
	StackName *string `json:"stackName" yaml:"stackName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user_stack_association#user_name AppstreamUserStackAssociation#user_name}.
	UserName *string `json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appstream_user_stack_association#send_email_notification AppstreamUserStackAssociation#send_email_notification}.
	SendEmailNotification interface{} `json:"sendEmailNotification" yaml:"sendEmailNotification"`
}
