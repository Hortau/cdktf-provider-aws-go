package macie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/macie/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/macie_member_account_association.html aws_macie_member_account_association}.
type MacieMemberAccountAssociation interface {
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
	MemberAccountId() *string
	SetMemberAccountId(val *string)
	MemberAccountIdInput() *string
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

// The jsii proxy struct for MacieMemberAccountAssociation
type jsiiProxy_MacieMemberAccountAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MacieMemberAccountAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) MemberAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) MemberAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieMemberAccountAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie_member_account_association.html aws_macie_member_account_association} Resource.
func NewMacieMemberAccountAssociation(scope constructs.Construct, id *string, config *MacieMemberAccountAssociationConfig) MacieMemberAccountAssociation {
	_init_.Initialize()

	j := jsiiProxy_MacieMemberAccountAssociation{}

	_jsii_.Create(
		"hashicorp_aws.Macie.MacieMemberAccountAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie_member_account_association.html aws_macie_member_account_association} Resource.
func NewMacieMemberAccountAssociation_Override(m MacieMemberAccountAssociation, scope constructs.Construct, id *string, config *MacieMemberAccountAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie.MacieMemberAccountAssociation",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MacieMemberAccountAssociation) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MacieMemberAccountAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MacieMemberAccountAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MacieMemberAccountAssociation) SetMemberAccountId(val *string) {
	_jsii_.Set(
		j,
		"memberAccountId",
		val,
	)
}

func (j *jsiiProxy_MacieMemberAccountAssociation) SetProvider(val cdktf.TerraformProvider) {
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
func MacieMemberAccountAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Macie.MacieMemberAccountAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MacieMemberAccountAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Macie.MacieMemberAccountAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MacieMemberAccountAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (m *jsiiProxy_MacieMemberAccountAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (m *jsiiProxy_MacieMemberAccountAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MacieMemberAccountAssociationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie_member_account_association.html#member_account_id MacieMemberAccountAssociation#member_account_id}.
	MemberAccountId *string `json:"memberAccountId"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association.html aws_macie_s3_bucket_association}.
type MacieS3BucketAssociation interface {
	cdktf.TerraformResource
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	CdktfStack() cdktf.TerraformStack
	ClassificationType() MacieS3BucketAssociationClassificationTypeOutputReference
	ClassificationTypeInput() *MacieS3BucketAssociationClassificationType
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
	MemberAccountId() *string
	SetMemberAccountId(val *string)
	MemberAccountIdInput() *string
	Node() constructs.Node
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
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
	PutClassificationType(value *MacieS3BucketAssociationClassificationType)
	ResetClassificationType()
	ResetMemberAccountId()
	ResetOverrideLogicalId()
	ResetPrefix()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for MacieS3BucketAssociation
type jsiiProxy_MacieS3BucketAssociation struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MacieS3BucketAssociation) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) ClassificationType() MacieS3BucketAssociationClassificationTypeOutputReference {
	var returns MacieS3BucketAssociationClassificationTypeOutputReference
	_jsii_.Get(
		j,
		"classificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) ClassificationTypeInput() *MacieS3BucketAssociationClassificationType {
	var returns *MacieS3BucketAssociationClassificationType
	_jsii_.Get(
		j,
		"classificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) MemberAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) MemberAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memberAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociation) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association.html aws_macie_s3_bucket_association} Resource.
func NewMacieS3BucketAssociation(scope constructs.Construct, id *string, config *MacieS3BucketAssociationConfig) MacieS3BucketAssociation {
	_init_.Initialize()

	j := jsiiProxy_MacieS3BucketAssociation{}

	_jsii_.Create(
		"hashicorp_aws.Macie.MacieS3BucketAssociation",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association.html aws_macie_s3_bucket_association} Resource.
func NewMacieS3BucketAssociation_Override(m MacieS3BucketAssociation, scope constructs.Construct, id *string, config *MacieS3BucketAssociationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie.MacieS3BucketAssociation",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociation) SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociation) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociation) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociation) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociation) SetMemberAccountId(val *string) {
	_jsii_.Set(
		j,
		"memberAccountId",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociation) SetPrefix(val *string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociation) SetProvider(val cdktf.TerraformProvider) {
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
func MacieS3BucketAssociation_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.Macie.MacieS3BucketAssociation",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MacieS3BucketAssociation_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.Macie.MacieS3BucketAssociation",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MacieS3BucketAssociation) PutClassificationType(value *MacieS3BucketAssociationClassificationType) {
	_jsii_.InvokeVoid(
		m,
		"putClassificationType",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MacieS3BucketAssociation) ResetClassificationType() {
	_jsii_.InvokeVoid(
		m,
		"resetClassificationType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MacieS3BucketAssociation) ResetMemberAccountId() {
	_jsii_.InvokeVoid(
		m,
		"resetMemberAccountId",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MacieS3BucketAssociation) ResetPrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetPrefix",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MacieS3BucketAssociation) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (m *jsiiProxy_MacieS3BucketAssociation) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociation) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MacieS3BucketAssociationClassificationType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association.html#continuous MacieS3BucketAssociation#continuous}.
	Continuous *string `json:"continuous"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association.html#one_time MacieS3BucketAssociation#one_time}.
	OneTime *string `json:"oneTime"`
}

type MacieS3BucketAssociationClassificationTypeOutputReference interface {
	cdktf.ComplexObject
	Continuous() *string
	SetContinuous(val *string)
	ContinuousInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	OneTime() *string
	SetOneTime(val *string)
	OneTimeInput() *string
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
	ResetContinuous()
	ResetOneTime()
}

// The jsii proxy struct for MacieS3BucketAssociationClassificationTypeOutputReference
type jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) Continuous() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuous",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) ContinuousInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"continuousInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) OneTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) OneTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewMacieS3BucketAssociationClassificationTypeOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) MacieS3BucketAssociationClassificationTypeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.Macie.MacieS3BucketAssociationClassificationTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewMacieS3BucketAssociationClassificationTypeOutputReference_Override(m MacieS3BucketAssociationClassificationTypeOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.Macie.MacieS3BucketAssociationClassificationTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		m,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) SetContinuous(val *string) {
	_jsii_.Set(
		j,
		"continuous",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) SetOneTime(val *string) {
	_jsii_.Set(
		j,
		"oneTime",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (m *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) ResetContinuous() {
	_jsii_.InvokeVoid(
		m,
		"resetContinuous",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MacieS3BucketAssociationClassificationTypeOutputReference) ResetOneTime() {
	_jsii_.InvokeVoid(
		m,
		"resetOneTime",
		nil, // no parameters
	)
}

type MacieS3BucketAssociationConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association.html#bucket_name MacieS3BucketAssociation#bucket_name}.
	BucketName *string `json:"bucketName"`
	// classification_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association.html#classification_type MacieS3BucketAssociation#classification_type}
	ClassificationType *MacieS3BucketAssociationClassificationType `json:"classificationType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association.html#member_account_id MacieS3BucketAssociation#member_account_id}.
	MemberAccountId *string `json:"memberAccountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/macie_s3_bucket_association.html#prefix MacieS3BucketAssociation#prefix}.
	Prefix *string `json:"prefix"`
}
