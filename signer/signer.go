package signer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/signer/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/signer_signing_job aws_signer_signing_job}.
type DataAwsSignerSigningJob interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	CompletedAt() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedAt() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	JobId() *string
	SetJobId(val *string)
	JobIdInput() *string
	JobInvoker() *string
	JobOwner() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PlatformDisplayName() *string
	PlatformId() *string
	ProfileName() *string
	ProfileVersion() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequestedBy() *string
	SignatureExpiresAt() *string
	Status() *string
	StatusReason() *string
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
	RevocationRecord(index *string) DataAwsSignerSigningJobRevocationRecord
	SignedObject(index *string) DataAwsSignerSigningJobSignedObject
	Source(index *string) DataAwsSignerSigningJobSource
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsSignerSigningJob
type jsiiProxy_DataAwsSignerSigningJob struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsSignerSigningJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) CompletedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) JobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) JobInvoker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobInvoker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) JobOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) PlatformDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) ProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) ProfileVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) RequestedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) SignatureExpiresAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureExpiresAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/signer_signing_job aws_signer_signing_job} Data Source.
func NewDataAwsSignerSigningJob(scope constructs.Construct, id *string, config *DataAwsSignerSigningJobConfig) DataAwsSignerSigningJob {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSignerSigningJob{}

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/signer_signing_job aws_signer_signing_job} Data Source.
func NewDataAwsSignerSigningJob_Override(d DataAwsSignerSigningJob, scope constructs.Construct, id *string, config *DataAwsSignerSigningJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJob",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJob) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJob) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJob) SetJobId(val *string) {
	_jsii_.Set(
		j,
		"jobId",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJob) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJob) SetProvider(val cdktf.TerraformProvider) {
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
func DataAwsSignerSigningJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.signer.DataAwsSignerSigningJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSignerSigningJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.signer.DataAwsSignerSigningJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningJob) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSignerSigningJob) RevocationRecord(index *string) DataAwsSignerSigningJobRevocationRecord {
	var returns DataAwsSignerSigningJobRevocationRecord

	_jsii_.Invoke(
		d,
		"revocationRecord",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSignerSigningJob) SignedObject(index *string) DataAwsSignerSigningJobSignedObject {
	var returns DataAwsSignerSigningJobSignedObject

	_jsii_.Invoke(
		d,
		"signedObject",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSignerSigningJob) Source(index *string) DataAwsSignerSigningJobSource {
	var returns DataAwsSignerSigningJobSource

	_jsii_.Invoke(
		d,
		"source",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSignerSigningJob) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) ToString() *string {
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
func (d *jsiiProxy_DataAwsSignerSigningJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Signer.
type DataAwsSignerSigningJobConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/signer_signing_job#job_id DataAwsSignerSigningJob#job_id}.
	JobId *string `json:"jobId" yaml:"jobId"`
}

type DataAwsSignerSigningJobRevocationRecord interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Reason() *string
	RevokedAt() *string
	RevokedBy() *string
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

// The jsii proxy struct for DataAwsSignerSigningJobRevocationRecord
type jsiiProxy_DataAwsSignerSigningJobRevocationRecord struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) Reason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) RevokedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) RevokedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsSignerSigningJobRevocationRecord(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsSignerSigningJobRevocationRecord {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSignerSigningJobRevocationRecord{}

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobRevocationRecord",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsSignerSigningJobRevocationRecord_Override(d DataAwsSignerSigningJobRevocationRecord, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobRevocationRecord",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobRevocationRecord) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsSignerSigningJobSignedObject interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	S3() cdktf.IResolvable
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

// The jsii proxy struct for DataAwsSignerSigningJobSignedObject
type jsiiProxy_DataAwsSignerSigningJobSignedObject struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObject) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObject) S3() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObject) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObject) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObject) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsSignerSigningJobSignedObject(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsSignerSigningJobSignedObject {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSignerSigningJobSignedObject{}

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobSignedObject",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsSignerSigningJobSignedObject_Override(d DataAwsSignerSigningJobSignedObject, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobSignedObject",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObject) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObject) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObject) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObject) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObject) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsSignerSigningJobSignedObjectS3 interface {
	cdktf.ComplexComputedList
	Bucket() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Key() *string
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

// The jsii proxy struct for DataAwsSignerSigningJobSignedObjectS3
type jsiiProxy_DataAwsSignerSigningJobSignedObjectS3 struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsSignerSigningJobSignedObjectS3(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsSignerSigningJobSignedObjectS3 {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSignerSigningJobSignedObjectS3{}

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobSignedObjectS3",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsSignerSigningJobSignedObjectS3_Override(d DataAwsSignerSigningJobSignedObjectS3, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobSignedObjectS3",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSignedObjectS3) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsSignerSigningJobSource interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	S3() cdktf.IResolvable
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

// The jsii proxy struct for DataAwsSignerSigningJobSource
type jsiiProxy_DataAwsSignerSigningJobSource struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsSignerSigningJobSource) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSource) S3() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSource) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSource) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSource) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsSignerSigningJobSource(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsSignerSigningJobSource {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSignerSigningJobSource{}

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobSource",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsSignerSigningJobSource_Override(d DataAwsSignerSigningJobSource, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobSource",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSource) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSource) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSource) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSource) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningJobSource) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSource) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSource) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSource) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSource) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSource) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSource) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSource) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSource) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSource) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsSignerSigningJobSourceS3 interface {
	cdktf.ComplexComputedList
	Bucket() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Key() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
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

// The jsii proxy struct for DataAwsSignerSigningJobSourceS3
type jsiiProxy_DataAwsSignerSigningJobSourceS3 struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsSignerSigningJobSourceS3(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsSignerSigningJobSourceS3 {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSignerSigningJobSourceS3{}

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobSourceS3",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsSignerSigningJobSourceS3_Override(d DataAwsSignerSigningJobSourceS3, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningJobSourceS3",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningJobSourceS3) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningJobSourceS3) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/signer_signing_profile aws_signer_signing_profile}.
type DataAwsSignerSigningProfile interface {
	cdktf.TerraformDataSource
	Arn() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	Node() constructs.Node
	PlatformDisplayName() *string
	PlatformId() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Status() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Version() *string
	VersionArn() *string
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
	ResetTags()
	RevocationRecord(index *string) DataAwsSignerSigningProfileRevocationRecord
	SignatureValidityPeriod(index *string) DataAwsSignerSigningProfileSignatureValidityPeriod
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsSignerSigningProfile
type jsiiProxy_DataAwsSignerSigningProfile struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) PlatformDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) VersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionArn",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/signer_signing_profile aws_signer_signing_profile} Data Source.
func NewDataAwsSignerSigningProfile(scope constructs.Construct, id *string, config *DataAwsSignerSigningProfileConfig) DataAwsSignerSigningProfile {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSignerSigningProfile{}

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/signer_signing_profile aws_signer_signing_profile} Data Source.
func NewDataAwsSignerSigningProfile_Override(d DataAwsSignerSigningProfile, scope constructs.Construct, id *string, config *DataAwsSignerSigningProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningProfile",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfile) SetTags(val *map[string]*string) {
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
func DataAwsSignerSigningProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.signer.DataAwsSignerSigningProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsSignerSigningProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.signer.DataAwsSignerSigningProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSignerSigningProfile) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsSignerSigningProfile) RevocationRecord(index *string) DataAwsSignerSigningProfileRevocationRecord {
	var returns DataAwsSignerSigningProfileRevocationRecord

	_jsii_.Invoke(
		d,
		"revocationRecord",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSignerSigningProfile) SignatureValidityPeriod(index *string) DataAwsSignerSigningProfileSignatureValidityPeriod {
	var returns DataAwsSignerSigningProfileSignatureValidityPeriod

	_jsii_.Invoke(
		d,
		"signatureValidityPeriod",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsSignerSigningProfile) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) ToString() *string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Signer.
type DataAwsSignerSigningProfileConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/signer_signing_profile#name DataAwsSignerSigningProfile#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/signer_signing_profile#tags DataAwsSignerSigningProfile#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type DataAwsSignerSigningProfileRevocationRecord interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	RevocationEffectiveFrom() *string
	RevokedAt() *string
	RevokedBy() *string
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

// The jsii proxy struct for DataAwsSignerSigningProfileRevocationRecord
type jsiiProxy_DataAwsSignerSigningProfileRevocationRecord struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) RevocationEffectiveFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revocationEffectiveFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) RevokedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) RevokedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsSignerSigningProfileRevocationRecord(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsSignerSigningProfileRevocationRecord {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSignerSigningProfileRevocationRecord{}

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningProfileRevocationRecord",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsSignerSigningProfileRevocationRecord_Override(d DataAwsSignerSigningProfileRevocationRecord, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningProfileRevocationRecord",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileRevocationRecord) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsSignerSigningProfileSignatureValidityPeriod interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	Value() *float64
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

// The jsii proxy struct for DataAwsSignerSigningProfileSignatureValidityPeriod
type jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsSignerSigningProfileSignatureValidityPeriod(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsSignerSigningProfileSignatureValidityPeriod {
	_init_.Initialize()

	j := jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod{}

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningProfileSignatureValidityPeriod",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsSignerSigningProfileSignatureValidityPeriod_Override(d DataAwsSignerSigningProfileSignatureValidityPeriod, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.DataAwsSignerSigningProfileSignatureValidityPeriod",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsSignerSigningProfileSignatureValidityPeriod) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job aws_signer_signing_job}.
type SignerSigningJob interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	CompletedAt() *string
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	CreatedAt() *string
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	Destination() SignerSigningJobDestinationOutputReference
	DestinationInput() *SignerSigningJobDestination
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IgnoreSigningJobFailure() interface{}
	SetIgnoreSigningJobFailure(val interface{})
	IgnoreSigningJobFailureInput() interface{}
	JobId() *string
	JobInvoker() *string
	JobOwner() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	PlatformDisplayName() *string
	PlatformId() *string
	ProfileName() *string
	SetProfileName(val *string)
	ProfileNameInput() *string
	ProfileVersion() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	RequestedBy() *string
	SignatureExpiresAt() *string
	Source() SignerSigningJobSourceOutputReference
	SourceInput() *SignerSigningJobSource
	Status() *string
	StatusReason() *string
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
	PutDestination(value *SignerSigningJobDestination)
	PutSource(value *SignerSigningJobSource)
	ResetIgnoreSigningJobFailure()
	ResetOverrideLogicalId()
	RevocationRecord(index *string) SignerSigningJobRevocationRecord
	SignedObject(index *string) SignerSigningJobSignedObject
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SignerSigningJob
type jsiiProxy_SignerSigningJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SignerSigningJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) CompletedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"completedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) Destination() SignerSigningJobDestinationOutputReference {
	var returns SignerSigningJobDestinationOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) DestinationInput() *SignerSigningJobDestination {
	var returns *SignerSigningJobDestination
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) IgnoreSigningJobFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreSigningJobFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) IgnoreSigningJobFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreSigningJobFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) JobInvoker() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobInvoker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) JobOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) PlatformDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) ProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) ProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) ProfileVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) RequestedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) SignatureExpiresAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signatureExpiresAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) Source() SignerSigningJobSourceOutputReference {
	var returns SignerSigningJobSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) SourceInput() *SignerSigningJobSource {
	var returns *SignerSigningJobSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job aws_signer_signing_job} Resource.
func NewSignerSigningJob(scope constructs.Construct, id *string, config *SignerSigningJobConfig) SignerSigningJob {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningJob{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job aws_signer_signing_job} Resource.
func NewSignerSigningJob_Override(s SignerSigningJob, scope constructs.Construct, id *string, config *SignerSigningJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJob",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SignerSigningJob) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJob) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJob) SetIgnoreSigningJobFailure(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreSigningJobFailure",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJob) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJob) SetProfileName(val *string) {
	_jsii_.Set(
		j,
		"profileName",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJob) SetProvider(val cdktf.TerraformProvider) {
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
func SignerSigningJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.signer.SignerSigningJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SignerSigningJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.signer.SignerSigningJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_SignerSigningJob) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SignerSigningJob) PutDestination(value *SignerSigningJobDestination) {
	_jsii_.InvokeVoid(
		s,
		"putDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignerSigningJob) PutSource(value *SignerSigningJobSource) {
	_jsii_.InvokeVoid(
		s,
		"putSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignerSigningJob) ResetIgnoreSigningJobFailure() {
	_jsii_.InvokeVoid(
		s,
		"resetIgnoreSigningJobFailure",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SignerSigningJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningJob) RevocationRecord(index *string) SignerSigningJobRevocationRecord {
	var returns SignerSigningJobRevocationRecord

	_jsii_.Invoke(
		s,
		"revocationRecord",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignerSigningJob) SignedObject(index *string) SignerSigningJobSignedObject {
	var returns SignerSigningJobSignedObject

	_jsii_.Invoke(
		s,
		"signedObject",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignerSigningJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SignerSigningJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_SignerSigningJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Signer.
type SignerSigningJobConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#destination SignerSigningJob#destination}
	Destination *SignerSigningJobDestination `json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#profile_name SignerSigningJob#profile_name}.
	ProfileName *string `json:"profileName" yaml:"profileName"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#source SignerSigningJob#source}
	Source *SignerSigningJobSource `json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#ignore_signing_job_failure SignerSigningJob#ignore_signing_job_failure}.
	IgnoreSigningJobFailure interface{} `json:"ignoreSigningJobFailure" yaml:"ignoreSigningJobFailure"`
}

type SignerSigningJobDestination struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#s3 SignerSigningJob#s3}
	S3 *SignerSigningJobDestinationS3 `json:"s3" yaml:"s3"`
}

type SignerSigningJobDestinationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SignerSigningJobDestination
	SetInternalValue(val *SignerSigningJobDestination)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	S3() SignerSigningJobDestinationS3OutputReference
	S3Input() *SignerSigningJobDestinationS3
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
	PutS3(value *SignerSigningJobDestinationS3)
}

// The jsii proxy struct for SignerSigningJobDestinationOutputReference
type jsiiProxy_SignerSigningJobDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) InternalValue() *SignerSigningJobDestination {
	var returns *SignerSigningJobDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) S3() SignerSigningJobDestinationS3OutputReference {
	var returns SignerSigningJobDestinationS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) S3Input() *SignerSigningJobDestinationS3 {
	var returns *SignerSigningJobDestinationS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSignerSigningJobDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SignerSigningJobDestinationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningJobDestinationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSignerSigningJobDestinationOutputReference_Override(s SignerSigningJobDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) SetInternalValue(val *SignerSigningJobDestination) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignerSigningJobDestinationOutputReference) PutS3(value *SignerSigningJobDestinationS3) {
	_jsii_.InvokeVoid(
		s,
		"putS3",
		[]interface{}{value},
	)
}

type SignerSigningJobDestinationS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#bucket SignerSigningJob#bucket}.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#prefix SignerSigningJob#prefix}.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

type SignerSigningJobDestinationS3OutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	InternalValue() *SignerSigningJobDestinationS3
	SetInternalValue(val *SignerSigningJobDestinationS3)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
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
	ResetPrefix()
}

// The jsii proxy struct for SignerSigningJobDestinationS3OutputReference
type jsiiProxy_SignerSigningJobDestinationS3OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) InternalValue() *SignerSigningJobDestinationS3 {
	var returns *SignerSigningJobDestinationS3
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSignerSigningJobDestinationS3OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SignerSigningJobDestinationS3OutputReference {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningJobDestinationS3OutputReference{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobDestinationS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSignerSigningJobDestinationS3OutputReference_Override(s SignerSigningJobDestinationS3OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobDestinationS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) SetInternalValue(val *SignerSigningJobDestinationS3) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) SetPrefix(val *string) {
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobDestinationS3OutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignerSigningJobDestinationS3OutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetPrefix",
		nil, // no parameters
	)
}

type SignerSigningJobRevocationRecord interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Reason() *string
	RevokedAt() *string
	RevokedBy() *string
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

// The jsii proxy struct for SignerSigningJobRevocationRecord
type jsiiProxy_SignerSigningJobRevocationRecord struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) Reason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) RevokedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) RevokedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewSignerSigningJobRevocationRecord(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) SignerSigningJobRevocationRecord {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningJobRevocationRecord{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobRevocationRecord",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewSignerSigningJobRevocationRecord_Override(s SignerSigningJobRevocationRecord, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobRevocationRecord",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobRevocationRecord) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobRevocationRecord) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SignerSigningJobSignedObject interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	S3() cdktf.IResolvable
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

// The jsii proxy struct for SignerSigningJobSignedObject
type jsiiProxy_SignerSigningJobSignedObject struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_SignerSigningJobSignedObject) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSignedObject) S3() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSignedObject) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSignedObject) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSignedObject) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewSignerSigningJobSignedObject(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) SignerSigningJobSignedObject {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningJobSignedObject{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobSignedObject",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewSignerSigningJobSignedObject_Override(s SignerSigningJobSignedObject, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobSignedObject",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SignerSigningJobSignedObject) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSignedObject) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSignedObject) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSignedObject) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObject) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SignerSigningJobSignedObjectS3 interface {
	cdktf.ComplexComputedList
	Bucket() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Key() *string
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

// The jsii proxy struct for SignerSigningJobSignedObjectS3
type jsiiProxy_SignerSigningJobSignedObjectS3 struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewSignerSigningJobSignedObjectS3(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) SignerSigningJobSignedObjectS3 {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningJobSignedObjectS3{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobSignedObjectS3",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewSignerSigningJobSignedObjectS3_Override(s SignerSigningJobSignedObjectS3, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobSignedObjectS3",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSignedObjectS3) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSignedObjectS3) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SignerSigningJobSource struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#s3 SignerSigningJob#s3}
	S3 *SignerSigningJobSourceS3 `json:"s3" yaml:"s3"`
}

type SignerSigningJobSourceOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SignerSigningJobSource
	SetInternalValue(val *SignerSigningJobSource)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	S3() SignerSigningJobSourceS3OutputReference
	S3Input() *SignerSigningJobSourceS3
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
	PutS3(value *SignerSigningJobSourceS3)
}

// The jsii proxy struct for SignerSigningJobSourceOutputReference
type jsiiProxy_SignerSigningJobSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) InternalValue() *SignerSigningJobSource {
	var returns *SignerSigningJobSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) S3() SignerSigningJobSourceS3OutputReference {
	var returns SignerSigningJobSourceS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) S3Input() *SignerSigningJobSourceS3 {
	var returns *SignerSigningJobSourceS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewSignerSigningJobSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SignerSigningJobSourceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningJobSourceOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSignerSigningJobSourceOutputReference_Override(s SignerSigningJobSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) SetInternalValue(val *SignerSigningJobSource) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignerSigningJobSourceOutputReference) PutS3(value *SignerSigningJobSourceS3) {
	_jsii_.InvokeVoid(
		s,
		"putS3",
		[]interface{}{value},
	)
}

type SignerSigningJobSourceS3 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#bucket SignerSigningJob#bucket}.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#key SignerSigningJob#key}.
	Key *string `json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_job#version SignerSigningJob#version}.
	Version *string `json:"version" yaml:"version"`
}

type SignerSigningJobSourceS3OutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	InternalValue() *SignerSigningJobSourceS3
	SetInternalValue(val *SignerSigningJobSourceS3)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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

// The jsii proxy struct for SignerSigningJobSourceS3OutputReference
type jsiiProxy_SignerSigningJobSourceS3OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) InternalValue() *SignerSigningJobSourceS3 {
	var returns *SignerSigningJobSourceS3
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func NewSignerSigningJobSourceS3OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SignerSigningJobSourceS3OutputReference {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningJobSourceS3OutputReference{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobSourceS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSignerSigningJobSourceS3OutputReference_Override(s SignerSigningJobSourceS3OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningJobSourceS3OutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) SetInternalValue(val *SignerSigningJobSourceS3) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) SetKey(val *string) {
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SignerSigningJobSourceS3OutputReference) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningJobSourceS3OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile aws_signer_signing_profile}.
type SignerSigningProfile interface {
	cdktf.TerraformResource
	Arn() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	Node() constructs.Node
	PlatformDisplayName() *string
	PlatformId() *string
	SetPlatformId(val *string)
	PlatformIdInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SignatureValidityPeriod() SignerSigningProfileSignatureValidityPeriodOutputReference
	SignatureValidityPeriodInput() *SignerSigningProfileSignatureValidityPeriod
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
	Version() *string
	VersionArn() *string
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
	PutSignatureValidityPeriod(value *SignerSigningProfileSignatureValidityPeriod)
	ResetName()
	ResetNamePrefix()
	ResetOverrideLogicalId()
	ResetSignatureValidityPeriod()
	ResetTags()
	ResetTagsAll()
	RevocationRecord(index *string) SignerSigningProfileRevocationRecord
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SignerSigningProfile
type jsiiProxy_SignerSigningProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SignerSigningProfile) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) PlatformDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) PlatformId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) PlatformIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) SignatureValidityPeriod() SignerSigningProfileSignatureValidityPeriodOutputReference {
	var returns SignerSigningProfileSignatureValidityPeriodOutputReference
	_jsii_.Get(
		j,
		"signatureValidityPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) SignatureValidityPeriodInput() *SignerSigningProfileSignatureValidityPeriod {
	var returns *SignerSigningProfileSignatureValidityPeriod
	_jsii_.Get(
		j,
		"signatureValidityPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfile) VersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionArn",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile aws_signer_signing_profile} Resource.
func NewSignerSigningProfile(scope constructs.Construct, id *string, config *SignerSigningProfileConfig) SignerSigningProfile {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningProfile{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile aws_signer_signing_profile} Resource.
func NewSignerSigningProfile_Override(s SignerSigningProfile, scope constructs.Construct, id *string, config *SignerSigningProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningProfile",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SignerSigningProfile) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfile) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfile) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfile) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfile) SetNamePrefix(val *string) {
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfile) SetPlatformId(val *string) {
	_jsii_.Set(
		j,
		"platformId",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfile) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfile) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfile) SetTagsAll(val *map[string]*string) {
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
func SignerSigningProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.signer.SignerSigningProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SignerSigningProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.signer.SignerSigningProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_SignerSigningProfile) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SignerSigningProfile) PutSignatureValidityPeriod(value *SignerSigningProfileSignatureValidityPeriod) {
	_jsii_.InvokeVoid(
		s,
		"putSignatureValidityPeriod",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SignerSigningProfile) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningProfile) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetNamePrefix",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SignerSigningProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningProfile) ResetSignatureValidityPeriod() {
	_jsii_.InvokeVoid(
		s,
		"resetSignatureValidityPeriod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningProfile) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningProfile) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningProfile) RevocationRecord(index *string) SignerSigningProfileRevocationRecord {
	var returns SignerSigningProfileRevocationRecord

	_jsii_.Invoke(
		s,
		"revocationRecord",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SignerSigningProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SignerSigningProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_SignerSigningProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Signer.
type SignerSigningProfileConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile#platform_id SignerSigningProfile#platform_id}.
	PlatformId *string `json:"platformId" yaml:"platformId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile#name SignerSigningProfile#name}.
	Name *string `json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile#name_prefix SignerSigningProfile#name_prefix}.
	NamePrefix *string `json:"namePrefix" yaml:"namePrefix"`
	// signature_validity_period block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile#signature_validity_period SignerSigningProfile#signature_validity_period}
	SignatureValidityPeriod *SignerSigningProfileSignatureValidityPeriod `json:"signatureValidityPeriod" yaml:"signatureValidityPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile#tags SignerSigningProfile#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile#tags_all SignerSigningProfile#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile_permission aws_signer_signing_profile_permission}.
type SignerSigningProfilePermission interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
	ProfileName() *string
	SetProfileName(val *string)
	ProfileNameInput() *string
	ProfileVersion() *string
	SetProfileVersion(val *string)
	ProfileVersionInput() *string
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	StatementId() *string
	SetStatementId(val *string)
	StatementIdInput() *string
	StatementIdPrefix() *string
	SetStatementIdPrefix(val *string)
	StatementIdPrefixInput() *string
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
	ResetProfileVersion()
	ResetStatementId()
	ResetStatementIdPrefix()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for SignerSigningProfilePermission
type jsiiProxy_SignerSigningProfilePermission struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SignerSigningProfilePermission) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) ProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) ProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) ProfileVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) ProfileVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) StatementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) StatementIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) StatementIdPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) StatementIdPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfilePermission) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile_permission aws_signer_signing_profile_permission} Resource.
func NewSignerSigningProfilePermission(scope constructs.Construct, id *string, config *SignerSigningProfilePermissionConfig) SignerSigningProfilePermission {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningProfilePermission{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningProfilePermission",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile_permission aws_signer_signing_profile_permission} Resource.
func NewSignerSigningProfilePermission_Override(s SignerSigningProfilePermission, scope constructs.Construct, id *string, config *SignerSigningProfilePermissionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningProfilePermission",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetProfileName(val *string) {
	_jsii_.Set(
		j,
		"profileName",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetProfileVersion(val *string) {
	_jsii_.Set(
		j,
		"profileVersion",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetStatementId(val *string) {
	_jsii_.Set(
		j,
		"statementId",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfilePermission) SetStatementIdPrefix(val *string) {
	_jsii_.Set(
		j,
		"statementIdPrefix",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SignerSigningProfilePermission_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.signer.SignerSigningProfilePermission",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SignerSigningProfilePermission_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.signer.SignerSigningProfilePermission",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningProfilePermission) ResetProfileVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetProfileVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningProfilePermission) ResetStatementId() {
	_jsii_.InvokeVoid(
		s,
		"resetStatementId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningProfilePermission) ResetStatementIdPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetStatementIdPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SignerSigningProfilePermission) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SignerSigningProfilePermission) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (s *jsiiProxy_SignerSigningProfilePermission) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS Signer.
type SignerSigningProfilePermissionConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile_permission#action SignerSigningProfilePermission#action}.
	Action *string `json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile_permission#principal SignerSigningProfilePermission#principal}.
	Principal *string `json:"principal" yaml:"principal"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile_permission#profile_name SignerSigningProfilePermission#profile_name}.
	ProfileName *string `json:"profileName" yaml:"profileName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile_permission#profile_version SignerSigningProfilePermission#profile_version}.
	ProfileVersion *string `json:"profileVersion" yaml:"profileVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile_permission#statement_id SignerSigningProfilePermission#statement_id}.
	StatementId *string `json:"statementId" yaml:"statementId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile_permission#statement_id_prefix SignerSigningProfilePermission#statement_id_prefix}.
	StatementIdPrefix *string `json:"statementIdPrefix" yaml:"statementIdPrefix"`
}

type SignerSigningProfileRevocationRecord interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	RevocationEffectiveFrom() *string
	RevokedAt() *string
	RevokedBy() *string
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

// The jsii proxy struct for SignerSigningProfileRevocationRecord
type jsiiProxy_SignerSigningProfileRevocationRecord struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) RevocationEffectiveFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revocationEffectiveFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) RevokedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) RevokedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewSignerSigningProfileRevocationRecord(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) SignerSigningProfileRevocationRecord {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningProfileRevocationRecord{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningProfileRevocationRecord",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewSignerSigningProfileRevocationRecord_Override(s SignerSigningProfileRevocationRecord, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningProfileRevocationRecord",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfileRevocationRecord) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileRevocationRecord) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type SignerSigningProfileSignatureValidityPeriod struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile#type SignerSigningProfile#type}.
	Type *string `json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/signer_signing_profile#value SignerSigningProfile#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type SignerSigningProfileSignatureValidityPeriodOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *SignerSigningProfileSignatureValidityPeriod
	SetInternalValue(val *SignerSigningProfileSignatureValidityPeriod)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Value() *float64
	SetValue(val *float64)
	ValueInput() *float64
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

// The jsii proxy struct for SignerSigningProfileSignatureValidityPeriodOutputReference
type jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) InternalValue() *SignerSigningProfileSignatureValidityPeriod {
	var returns *SignerSigningProfileSignatureValidityPeriod
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewSignerSigningProfileSignatureValidityPeriodOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) SignerSigningProfileSignatureValidityPeriodOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningProfileSignatureValidityPeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewSignerSigningProfileSignatureValidityPeriodOutputReference_Override(s SignerSigningProfileSignatureValidityPeriodOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.signer.SignerSigningProfileSignatureValidityPeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		s,
	)
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) SetInternalValue(val *SignerSigningProfileSignatureValidityPeriod) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (s *jsiiProxy_SignerSigningProfileSignatureValidityPeriodOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}
