package elasticsearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/elasticsearch/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain aws_elasticsearch_domain}.
type DataAwsElasticsearchDomain interface {
	cdktf.TerraformDataSource
	AccessPolicies() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	Created() cdktf.IResolvable
	Deleted() cdktf.IResolvable
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DomainId() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	ElasticsearchVersion() *string
	Endpoint() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KibanaEndpoint() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Processing() cdktf.IResolvable
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	AddOverride(path *string, value interface{})
	AdvancedOptions(key *string) interface{}
	AdvancedSecurityOptions(index *string) DataAwsElasticsearchDomainAdvancedSecurityOptions
	AutoTuneOptions(index *string) DataAwsElasticsearchDomainAutoTuneOptions
	ClusterConfig(index *string) DataAwsElasticsearchDomainClusterConfig
	CognitoOptions(index *string) DataAwsElasticsearchDomainCognitoOptions
	EbsOptions(index *string) DataAwsElasticsearchDomainEbsOptions
	EncryptionAtRest(index *string) DataAwsElasticsearchDomainEncryptionAtRest
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
	LogPublishingOptions(index *string) DataAwsElasticsearchDomainLogPublishingOptions
	NodeToNodeEncryption(index *string) DataAwsElasticsearchDomainNodeToNodeEncryption
	OverrideLogicalId(newLogicalId *string)
	ResetOverrideLogicalId()
	ResetTags()
	SnapshotOptions(index *string) DataAwsElasticsearchDomainSnapshotOptions
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
	VpcOptions(index *string) DataAwsElasticsearchDomainVpcOptions
}

// The jsii proxy struct for DataAwsElasticsearchDomain
type jsiiProxy_DataAwsElasticsearchDomain struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Created() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Deleted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deleted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) ElasticsearchVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) KibanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kibanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Processing() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"processing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain aws_elasticsearch_domain} Data Source.
func NewDataAwsElasticsearchDomain(scope constructs.Construct, id *string, config *DataAwsElasticsearchDomainConfig) DataAwsElasticsearchDomain {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomain{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain aws_elasticsearch_domain} Data Source.
func NewDataAwsElasticsearchDomain_Override(d DataAwsElasticsearchDomain, scope constructs.Construct, id *string, config *DataAwsElasticsearchDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomain",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetTags(val *map[string]*string) {
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
func DataAwsElasticsearchDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsElasticsearchDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) AdvancedOptions(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"advancedOptions",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) AdvancedSecurityOptions(index *string) DataAwsElasticsearchDomainAdvancedSecurityOptions {
	var returns DataAwsElasticsearchDomainAdvancedSecurityOptions

	_jsii_.Invoke(
		d,
		"advancedSecurityOptions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) AutoTuneOptions(index *string) DataAwsElasticsearchDomainAutoTuneOptions {
	var returns DataAwsElasticsearchDomainAutoTuneOptions

	_jsii_.Invoke(
		d,
		"autoTuneOptions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) ClusterConfig(index *string) DataAwsElasticsearchDomainClusterConfig {
	var returns DataAwsElasticsearchDomainClusterConfig

	_jsii_.Invoke(
		d,
		"clusterConfig",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) CognitoOptions(index *string) DataAwsElasticsearchDomainCognitoOptions {
	var returns DataAwsElasticsearchDomainCognitoOptions

	_jsii_.Invoke(
		d,
		"cognitoOptions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) EbsOptions(index *string) DataAwsElasticsearchDomainEbsOptions {
	var returns DataAwsElasticsearchDomainEbsOptions

	_jsii_.Invoke(
		d,
		"ebsOptions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) EncryptionAtRest(index *string) DataAwsElasticsearchDomainEncryptionAtRest {
	var returns DataAwsElasticsearchDomainEncryptionAtRest

	_jsii_.Invoke(
		d,
		"encryptionAtRest",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) LogPublishingOptions(index *string) DataAwsElasticsearchDomainLogPublishingOptions {
	var returns DataAwsElasticsearchDomainLogPublishingOptions

	_jsii_.Invoke(
		d,
		"logPublishingOptions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) NodeToNodeEncryption(index *string) DataAwsElasticsearchDomainNodeToNodeEncryption {
	var returns DataAwsElasticsearchDomainNodeToNodeEncryption

	_jsii_.Invoke(
		d,
		"nodeToNodeEncryption",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) SnapshotOptions(index *string) DataAwsElasticsearchDomainSnapshotOptions {
	var returns DataAwsElasticsearchDomainSnapshotOptions

	_jsii_.Invoke(
		d,
		"snapshotOptions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) ToString() *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsElasticsearchDomain) VpcOptions(index *string) DataAwsElasticsearchDomainVpcOptions {
	var returns DataAwsElasticsearchDomainVpcOptions

	_jsii_.Invoke(
		d,
		"vpcOptions",
		[]interface{}{index},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainAdvancedSecurityOptions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Enabled() cdktf.IResolvable
	InternalUserDatabaseEnabled() cdktf.IResolvable
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

// The jsii proxy struct for DataAwsElasticsearchDomainAdvancedSecurityOptions
type jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) InternalUserDatabaseEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainAdvancedSecurityOptions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainAdvancedSecurityOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainAdvancedSecurityOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainAdvancedSecurityOptions_Override(d DataAwsElasticsearchDomainAdvancedSecurityOptions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainAdvancedSecurityOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainAutoTuneOptions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DesiredState() *string
	MaintenanceSchedule() cdktf.IResolvable
	RollbackOnDisable() *string
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

// The jsii proxy struct for DataAwsElasticsearchDomainAutoTuneOptions
type jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) MaintenanceSchedule() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"maintenanceSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) RollbackOnDisable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rollbackOnDisable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainAutoTuneOptions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainAutoTuneOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainAutoTuneOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainAutoTuneOptions_Override(d DataAwsElasticsearchDomainAutoTuneOptions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainAutoTuneOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	CronExpressionForRecurrence() *string
	Duration() cdktf.IResolvable
	StartAt() *string
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

// The jsii proxy struct for DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule
type jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) CronExpressionForRecurrence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cronExpressionForRecurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) Duration() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) StartAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule_Override(d DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceSchedule) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
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

// The jsii proxy struct for DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration
type jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration_Override(d DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainClusterConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DedicatedMasterCount() *float64
	DedicatedMasterEnabled() cdktf.IResolvable
	DedicatedMasterType() *string
	InstanceCount() *float64
	InstanceType() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WarmCount() *float64
	WarmEnabled() cdktf.IResolvable
	WarmType() *string
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	ZoneAwarenessConfig() cdktf.IResolvable
	ZoneAwarenessEnabled() cdktf.IResolvable
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

// The jsii proxy struct for DataAwsElasticsearchDomainClusterConfig
type jsiiProxy_DataAwsElasticsearchDomainClusterConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) DedicatedMasterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) DedicatedMasterEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"dedicatedMasterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) DedicatedMasterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) WarmCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) WarmEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"warmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) WarmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) ZoneAwarenessConfig() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"zoneAwarenessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) ZoneAwarenessEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"zoneAwarenessEnabled",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainClusterConfig(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainClusterConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainClusterConfig{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainClusterConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainClusterConfig_Override(d DataAwsElasticsearchDomainClusterConfig, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainClusterConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig interface {
	cdktf.ComplexComputedList
	AvailabilityZoneCount() *float64
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
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

// The jsii proxy struct for DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig
type jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) AvailabilityZoneCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityZoneCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig_Override(d DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainCognitoOptions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Enabled() cdktf.IResolvable
	IdentityPoolId() *string
	RoleArn() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserPoolId() *string
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

// The jsii proxy struct for DataAwsElasticsearchDomainCognitoOptions
type jsiiProxy_DataAwsElasticsearchDomainCognitoOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainCognitoOptions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainCognitoOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainCognitoOptions{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainCognitoOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainCognitoOptions_Override(d DataAwsElasticsearchDomainCognitoOptions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainCognitoOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// AWS ElasticSearch Service.
type DataAwsElasticsearchDomainConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain#domain_name DataAwsElasticsearchDomain#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain#tags DataAwsElasticsearchDomain#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

type DataAwsElasticsearchDomainEbsOptions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	EbsEnabled() cdktf.IResolvable
	Iops() *float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeSize() *float64
	VolumeType() *string
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

// The jsii proxy struct for DataAwsElasticsearchDomainEbsOptions
type jsiiProxy_DataAwsElasticsearchDomainEbsOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) EbsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ebsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) VolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainEbsOptions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainEbsOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainEbsOptions{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainEbsOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainEbsOptions_Override(d DataAwsElasticsearchDomainEbsOptions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainEbsOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainEncryptionAtRest interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Enabled() cdktf.IResolvable
	KmsKeyId() *string
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

// The jsii proxy struct for DataAwsElasticsearchDomainEncryptionAtRest
type jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainEncryptionAtRest(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainEncryptionAtRest {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainEncryptionAtRest",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainEncryptionAtRest_Override(d DataAwsElasticsearchDomainEncryptionAtRest, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainEncryptionAtRest",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainLogPublishingOptions interface {
	cdktf.ComplexComputedList
	CloudwatchLogGroupArn() *string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Enabled() cdktf.IResolvable
	LogType() *string
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

// The jsii proxy struct for DataAwsElasticsearchDomainLogPublishingOptions
type jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) CloudwatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainLogPublishingOptions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainLogPublishingOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainLogPublishingOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainLogPublishingOptions_Override(d DataAwsElasticsearchDomainLogPublishingOptions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainLogPublishingOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainNodeToNodeEncryption interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	Enabled() cdktf.IResolvable
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

// The jsii proxy struct for DataAwsElasticsearchDomainNodeToNodeEncryption
type jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainNodeToNodeEncryption(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainNodeToNodeEncryption {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainNodeToNodeEncryption",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainNodeToNodeEncryption_Override(d DataAwsElasticsearchDomainNodeToNodeEncryption, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainNodeToNodeEncryption",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainSnapshotOptions interface {
	cdktf.ComplexComputedList
	AutomatedSnapshotStartHour() *float64
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
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

// The jsii proxy struct for DataAwsElasticsearchDomainSnapshotOptions
type jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) AutomatedSnapshotStartHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotStartHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainSnapshotOptions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainSnapshotOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainSnapshotOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainSnapshotOptions_Override(d DataAwsElasticsearchDomainSnapshotOptions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainSnapshotOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsElasticsearchDomainVpcOptions interface {
	cdktf.ComplexComputedList
	AvailabilityZones() *[]*string
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	SecurityGroupIds() *[]*string
	SubnetIds() *[]*string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcId() *string
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

// The jsii proxy struct for DataAwsElasticsearchDomainVpcOptions
type jsiiProxy_DataAwsElasticsearchDomainVpcOptions struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainVpcOptions(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) DataAwsElasticsearchDomainVpcOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainVpcOptions{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainVpcOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainVpcOptions_Override(d DataAwsElasticsearchDomainVpcOptions, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexComputedListIndex *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.DataAwsElasticsearchDomainVpcOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain aws_elasticsearch_domain}.
type ElasticsearchDomain interface {
	cdktf.TerraformResource
	AccessPolicies() *string
	SetAccessPolicies(val *string)
	AccessPoliciesInput() *string
	AdvancedOptions() *map[string]*string
	SetAdvancedOptions(val *map[string]*string)
	AdvancedOptionsInput() *map[string]*string
	AdvancedSecurityOptions() ElasticsearchDomainAdvancedSecurityOptionsOutputReference
	AdvancedSecurityOptionsInput() *ElasticsearchDomainAdvancedSecurityOptions
	Arn() *string
	AutoTuneOptions() ElasticsearchDomainAutoTuneOptionsOutputReference
	AutoTuneOptionsInput() *ElasticsearchDomainAutoTuneOptions
	CdktfStack() cdktf.TerraformStack
	ClusterConfig() ElasticsearchDomainClusterConfigOutputReference
	ClusterConfigInput() *ElasticsearchDomainClusterConfig
	CognitoOptions() ElasticsearchDomainCognitoOptionsOutputReference
	CognitoOptionsInput() *ElasticsearchDomainCognitoOptions
	ConstructNodeMetadata() *map[string]interface{}
	Count() *float64
	SetCount(val *float64)
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	DomainEndpointOptions() ElasticsearchDomainDomainEndpointOptionsOutputReference
	DomainEndpointOptionsInput() *ElasticsearchDomainDomainEndpointOptions
	DomainId() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EbsOptions() ElasticsearchDomainEbsOptionsOutputReference
	EbsOptionsInput() *ElasticsearchDomainEbsOptions
	ElasticsearchVersion() *string
	SetElasticsearchVersion(val *string)
	ElasticsearchVersionInput() *string
	EncryptAtRest() ElasticsearchDomainEncryptAtRestOutputReference
	EncryptAtRestInput() *ElasticsearchDomainEncryptAtRest
	Endpoint() *string
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	KibanaEndpoint() *string
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogPublishingOptions() interface{}
	SetLogPublishingOptions(val interface{})
	LogPublishingOptionsInput() interface{}
	Node() constructs.Node
	NodeToNodeEncryption() ElasticsearchDomainNodeToNodeEncryptionOutputReference
	NodeToNodeEncryptionInput() *ElasticsearchDomainNodeToNodeEncryption
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SnapshotOptions() ElasticsearchDomainSnapshotOptionsOutputReference
	SnapshotOptionsInput() *ElasticsearchDomainSnapshotOptions
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() ElasticsearchDomainTimeoutsOutputReference
	TimeoutsInput() *ElasticsearchDomainTimeouts
	VpcOptions() ElasticsearchDomainVpcOptionsOutputReference
	VpcOptionsInput() *ElasticsearchDomainVpcOptions
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
	PutAdvancedSecurityOptions(value *ElasticsearchDomainAdvancedSecurityOptions)
	PutAutoTuneOptions(value *ElasticsearchDomainAutoTuneOptions)
	PutClusterConfig(value *ElasticsearchDomainClusterConfig)
	PutCognitoOptions(value *ElasticsearchDomainCognitoOptions)
	PutDomainEndpointOptions(value *ElasticsearchDomainDomainEndpointOptions)
	PutEbsOptions(value *ElasticsearchDomainEbsOptions)
	PutEncryptAtRest(value *ElasticsearchDomainEncryptAtRest)
	PutNodeToNodeEncryption(value *ElasticsearchDomainNodeToNodeEncryption)
	PutSnapshotOptions(value *ElasticsearchDomainSnapshotOptions)
	PutTimeouts(value *ElasticsearchDomainTimeouts)
	PutVpcOptions(value *ElasticsearchDomainVpcOptions)
	ResetAccessPolicies()
	ResetAdvancedOptions()
	ResetAdvancedSecurityOptions()
	ResetAutoTuneOptions()
	ResetClusterConfig()
	ResetCognitoOptions()
	ResetDomainEndpointOptions()
	ResetEbsOptions()
	ResetElasticsearchVersion()
	ResetEncryptAtRest()
	ResetLogPublishingOptions()
	ResetNodeToNodeEncryption()
	ResetOverrideLogicalId()
	ResetSnapshotOptions()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVpcOptions()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ElasticsearchDomain
type jsiiProxy_ElasticsearchDomain struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElasticsearchDomain) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AccessPoliciesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AdvancedOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AdvancedOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AdvancedSecurityOptions() ElasticsearchDomainAdvancedSecurityOptionsOutputReference {
	var returns ElasticsearchDomainAdvancedSecurityOptionsOutputReference
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AdvancedSecurityOptionsInput() *ElasticsearchDomainAdvancedSecurityOptions {
	var returns *ElasticsearchDomainAdvancedSecurityOptions
	_jsii_.Get(
		j,
		"advancedSecurityOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AutoTuneOptions() ElasticsearchDomainAutoTuneOptionsOutputReference {
	var returns ElasticsearchDomainAutoTuneOptionsOutputReference
	_jsii_.Get(
		j,
		"autoTuneOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AutoTuneOptionsInput() *ElasticsearchDomainAutoTuneOptions {
	var returns *ElasticsearchDomainAutoTuneOptions
	_jsii_.Get(
		j,
		"autoTuneOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ClusterConfig() ElasticsearchDomainClusterConfigOutputReference {
	var returns ElasticsearchDomainClusterConfigOutputReference
	_jsii_.Get(
		j,
		"clusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ClusterConfigInput() *ElasticsearchDomainClusterConfig {
	var returns *ElasticsearchDomainClusterConfig
	_jsii_.Get(
		j,
		"clusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) CognitoOptions() ElasticsearchDomainCognitoOptionsOutputReference {
	var returns ElasticsearchDomainCognitoOptionsOutputReference
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) CognitoOptionsInput() *ElasticsearchDomainCognitoOptions {
	var returns *ElasticsearchDomainCognitoOptions
	_jsii_.Get(
		j,
		"cognitoOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainEndpointOptions() ElasticsearchDomainDomainEndpointOptionsOutputReference {
	var returns ElasticsearchDomainDomainEndpointOptionsOutputReference
	_jsii_.Get(
		j,
		"domainEndpointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainEndpointOptionsInput() *ElasticsearchDomainDomainEndpointOptions {
	var returns *ElasticsearchDomainDomainEndpointOptions
	_jsii_.Get(
		j,
		"domainEndpointOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) EbsOptions() ElasticsearchDomainEbsOptionsOutputReference {
	var returns ElasticsearchDomainEbsOptionsOutputReference
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) EbsOptionsInput() *ElasticsearchDomainEbsOptions {
	var returns *ElasticsearchDomainEbsOptions
	_jsii_.Get(
		j,
		"ebsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ElasticsearchVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) ElasticsearchVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticsearchVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) EncryptAtRest() ElasticsearchDomainEncryptAtRestOutputReference {
	var returns ElasticsearchDomainEncryptAtRestOutputReference
	_jsii_.Get(
		j,
		"encryptAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) EncryptAtRestInput() *ElasticsearchDomainEncryptAtRest {
	var returns *ElasticsearchDomainEncryptAtRest
	_jsii_.Get(
		j,
		"encryptAtRestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) KibanaEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kibanaEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) LogPublishingOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) LogPublishingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPublishingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) NodeToNodeEncryption() ElasticsearchDomainNodeToNodeEncryptionOutputReference {
	var returns ElasticsearchDomainNodeToNodeEncryptionOutputReference
	_jsii_.Get(
		j,
		"nodeToNodeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) NodeToNodeEncryptionInput() *ElasticsearchDomainNodeToNodeEncryption {
	var returns *ElasticsearchDomainNodeToNodeEncryption
	_jsii_.Get(
		j,
		"nodeToNodeEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) SnapshotOptions() ElasticsearchDomainSnapshotOptionsOutputReference {
	var returns ElasticsearchDomainSnapshotOptionsOutputReference
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) SnapshotOptionsInput() *ElasticsearchDomainSnapshotOptions {
	var returns *ElasticsearchDomainSnapshotOptions
	_jsii_.Get(
		j,
		"snapshotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) Timeouts() ElasticsearchDomainTimeoutsOutputReference {
	var returns ElasticsearchDomainTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TimeoutsInput() *ElasticsearchDomainTimeouts {
	var returns *ElasticsearchDomainTimeouts
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) VpcOptions() ElasticsearchDomainVpcOptionsOutputReference {
	var returns ElasticsearchDomainVpcOptionsOutputReference
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) VpcOptionsInput() *ElasticsearchDomainVpcOptions {
	var returns *ElasticsearchDomainVpcOptions
	_jsii_.Get(
		j,
		"vpcOptionsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain aws_elasticsearch_domain} Resource.
func NewElasticsearchDomain(scope constructs.Construct, id *string, config *ElasticsearchDomainConfig) ElasticsearchDomain {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomain{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain aws_elasticsearch_domain} Resource.
func NewElasticsearchDomain_Override(e ElasticsearchDomain, scope constructs.Construct, id *string, config *ElasticsearchDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomain",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetAccessPolicies(val *string) {
	_jsii_.Set(
		j,
		"accessPolicies",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetAdvancedOptions(val *map[string]*string) {
	_jsii_.Set(
		j,
		"advancedOptions",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetElasticsearchVersion(val *string) {
	_jsii_.Set(
		j,
		"elasticsearchVersion",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetLogPublishingOptions(val interface{}) {
	_jsii_.Set(
		j,
		"logPublishingOptions",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetTagsAll(val *map[string]*string) {
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
func ElasticsearchDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.elasticsearch.ElasticsearchDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticsearchDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.elasticsearch.ElasticsearchDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutAdvancedSecurityOptions(value *ElasticsearchDomainAdvancedSecurityOptions) {
	_jsii_.InvokeVoid(
		e,
		"putAdvancedSecurityOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutAutoTuneOptions(value *ElasticsearchDomainAutoTuneOptions) {
	_jsii_.InvokeVoid(
		e,
		"putAutoTuneOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutClusterConfig(value *ElasticsearchDomainClusterConfig) {
	_jsii_.InvokeVoid(
		e,
		"putClusterConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutCognitoOptions(value *ElasticsearchDomainCognitoOptions) {
	_jsii_.InvokeVoid(
		e,
		"putCognitoOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutDomainEndpointOptions(value *ElasticsearchDomainDomainEndpointOptions) {
	_jsii_.InvokeVoid(
		e,
		"putDomainEndpointOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutEbsOptions(value *ElasticsearchDomainEbsOptions) {
	_jsii_.InvokeVoid(
		e,
		"putEbsOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutEncryptAtRest(value *ElasticsearchDomainEncryptAtRest) {
	_jsii_.InvokeVoid(
		e,
		"putEncryptAtRest",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutNodeToNodeEncryption(value *ElasticsearchDomainNodeToNodeEncryption) {
	_jsii_.InvokeVoid(
		e,
		"putNodeToNodeEncryption",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutSnapshotOptions(value *ElasticsearchDomainSnapshotOptions) {
	_jsii_.InvokeVoid(
		e,
		"putSnapshotOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutTimeouts(value *ElasticsearchDomainTimeouts) {
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) PutVpcOptions(value *ElasticsearchDomainVpcOptions) {
	_jsii_.InvokeVoid(
		e,
		"putVpcOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetAccessPolicies() {
	_jsii_.InvokeVoid(
		e,
		"resetAccessPolicies",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetAdvancedOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetAdvancedSecurityOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedSecurityOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetAutoTuneOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoTuneOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetClusterConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetClusterConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetCognitoOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetCognitoOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetDomainEndpointOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetDomainEndpointOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetEbsOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetEbsOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetElasticsearchVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetElasticsearchVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetEncryptAtRest() {
	_jsii_.InvokeVoid(
		e,
		"resetEncryptAtRest",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetLogPublishingOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetLogPublishingOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetNodeToNodeEncryption() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeToNodeEncryption",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetSnapshotOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) ResetVpcOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetVpcOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_ElasticsearchDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_ElasticsearchDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ElasticsearchDomainAdvancedSecurityOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#internal_user_database_enabled ElasticsearchDomain#internal_user_database_enabled}.
	InternalUserDatabaseEnabled interface{} `json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// master_user_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#master_user_options ElasticsearchDomain#master_user_options}
	MasterUserOptions *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions `json:"masterUserOptions" yaml:"masterUserOptions"`
}

type ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#master_user_arn ElasticsearchDomain#master_user_arn}.
	MasterUserArn *string `json:"masterUserArn" yaml:"masterUserArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#master_user_name ElasticsearchDomain#master_user_name}.
	MasterUserName *string `json:"masterUserName" yaml:"masterUserName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#master_user_password ElasticsearchDomain#master_user_password}.
	MasterUserPassword *string `json:"masterUserPassword" yaml:"masterUserPassword"`
}

type ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions
	SetInternalValue(val *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MasterUserArn() *string
	SetMasterUserArn(val *string)
	MasterUserArnInput() *string
	MasterUserName() *string
	SetMasterUserName(val *string)
	MasterUserNameInput() *string
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	MasterUserPasswordInput() *string
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
	ResetMasterUserArn()
	ResetMasterUserName()
	ResetMasterUserPassword()
}

// The jsii proxy struct for ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
type jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) InternalValue() *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions {
	var returns *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference_Override(e ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) SetInternalValue(val *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) SetMasterUserArn(val *string) {
	_jsii_.Set(
		j,
		"masterUserArn",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) SetMasterUserName(val *string) {
	_jsii_.Set(
		j,
		"masterUserName",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ResetMasterUserArn() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterUserArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ResetMasterUserName() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterUserName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ResetMasterUserPassword() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterUserPassword",
		nil, // no parameters
	)
}

type ElasticsearchDomainAdvancedSecurityOptionsOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalUserDatabaseEnabled() interface{}
	SetInternalUserDatabaseEnabled(val interface{})
	InternalUserDatabaseEnabledInput() interface{}
	InternalValue() *ElasticsearchDomainAdvancedSecurityOptions
	SetInternalValue(val *ElasticsearchDomainAdvancedSecurityOptions)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MasterUserOptions() ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
	MasterUserOptionsInput() *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions
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
	PutMasterUserOptions(value *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions)
	ResetInternalUserDatabaseEnabled()
	ResetMasterUserOptions()
}

// The jsii proxy struct for ElasticsearchDomainAdvancedSecurityOptionsOutputReference
type jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) InternalUserDatabaseEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) InternalUserDatabaseEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalUserDatabaseEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) InternalValue() *ElasticsearchDomainAdvancedSecurityOptions {
	var returns *ElasticsearchDomainAdvancedSecurityOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) MasterUserOptions() ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference {
	var returns ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
	_jsii_.Get(
		j,
		"masterUserOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) MasterUserOptionsInput() *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions {
	var returns *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions
	_jsii_.Get(
		j,
		"masterUserOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainAdvancedSecurityOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainAdvancedSecurityOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainAdvancedSecurityOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainAdvancedSecurityOptionsOutputReference_Override(e ElasticsearchDomainAdvancedSecurityOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainAdvancedSecurityOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) SetInternalUserDatabaseEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"internalUserDatabaseEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) SetInternalValue(val *ElasticsearchDomainAdvancedSecurityOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) PutMasterUserOptions(value *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions) {
	_jsii_.InvokeVoid(
		e,
		"putMasterUserOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) ResetInternalUserDatabaseEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetInternalUserDatabaseEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) ResetMasterUserOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterUserOptions",
		nil, // no parameters
	)
}

type ElasticsearchDomainAutoTuneOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#desired_state ElasticsearchDomain#desired_state}.
	DesiredState *string `json:"desiredState" yaml:"desiredState"`
	// maintenance_schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#maintenance_schedule ElasticsearchDomain#maintenance_schedule}
	MaintenanceSchedule interface{} `json:"maintenanceSchedule" yaml:"maintenanceSchedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#rollback_on_disable ElasticsearchDomain#rollback_on_disable}.
	RollbackOnDisable *string `json:"rollbackOnDisable" yaml:"rollbackOnDisable"`
}

type ElasticsearchDomainAutoTuneOptionsMaintenanceSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#cron_expression_for_recurrence ElasticsearchDomain#cron_expression_for_recurrence}.
	CronExpressionForRecurrence *string `json:"cronExpressionForRecurrence" yaml:"cronExpressionForRecurrence"`
	// duration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#duration ElasticsearchDomain#duration}
	Duration *ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration `json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#start_at ElasticsearchDomain#start_at}.
	StartAt *string `json:"startAt" yaml:"startAt"`
}

type ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#unit ElasticsearchDomain#unit}.
	Unit *string `json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#value ElasticsearchDomain#value}.
	Value *float64 `json:"value" yaml:"value"`
}

type ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration
	SetInternalValue(val *ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unit() *string
	SetUnit(val *string)
	UnitInput() *string
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

// The jsii proxy struct for ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference
type jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) InternalValue() *ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration {
	var returns *ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) ValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference_Override(e ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) SetInternalValue(val *ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) SetUnit(val *string) {
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) SetValue(val *float64) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type ElasticsearchDomainAutoTuneOptionsOutputReference interface {
	cdktf.ComplexObject
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	InternalValue() *ElasticsearchDomainAutoTuneOptions
	SetInternalValue(val *ElasticsearchDomainAutoTuneOptions)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MaintenanceSchedule() interface{}
	SetMaintenanceSchedule(val interface{})
	MaintenanceScheduleInput() interface{}
	RollbackOnDisable() *string
	SetRollbackOnDisable(val *string)
	RollbackOnDisableInput() *string
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
	ResetMaintenanceSchedule()
	ResetRollbackOnDisable()
}

// The jsii proxy struct for ElasticsearchDomainAutoTuneOptionsOutputReference
type jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) InternalValue() *ElasticsearchDomainAutoTuneOptions {
	var returns *ElasticsearchDomainAutoTuneOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) MaintenanceSchedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) MaintenanceScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) RollbackOnDisable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rollbackOnDisable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) RollbackOnDisableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rollbackOnDisableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainAutoTuneOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainAutoTuneOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainAutoTuneOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainAutoTuneOptionsOutputReference_Override(e ElasticsearchDomainAutoTuneOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainAutoTuneOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) SetDesiredState(val *string) {
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) SetInternalValue(val *ElasticsearchDomainAutoTuneOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) SetMaintenanceSchedule(val interface{}) {
	_jsii_.Set(
		j,
		"maintenanceSchedule",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) SetRollbackOnDisable(val *string) {
	_jsii_.Set(
		j,
		"rollbackOnDisable",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) ResetMaintenanceSchedule() {
	_jsii_.InvokeVoid(
		e,
		"resetMaintenanceSchedule",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainAutoTuneOptionsOutputReference) ResetRollbackOnDisable() {
	_jsii_.InvokeVoid(
		e,
		"resetRollbackOnDisable",
		nil, // no parameters
	)
}

type ElasticsearchDomainClusterConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#dedicated_master_count ElasticsearchDomain#dedicated_master_count}.
	DedicatedMasterCount *float64 `json:"dedicatedMasterCount" yaml:"dedicatedMasterCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#dedicated_master_enabled ElasticsearchDomain#dedicated_master_enabled}.
	DedicatedMasterEnabled interface{} `json:"dedicatedMasterEnabled" yaml:"dedicatedMasterEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#dedicated_master_type ElasticsearchDomain#dedicated_master_type}.
	DedicatedMasterType *string `json:"dedicatedMasterType" yaml:"dedicatedMasterType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#instance_count ElasticsearchDomain#instance_count}.
	InstanceCount *float64 `json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#instance_type ElasticsearchDomain#instance_type}.
	InstanceType *string `json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#warm_count ElasticsearchDomain#warm_count}.
	WarmCount *float64 `json:"warmCount" yaml:"warmCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#warm_enabled ElasticsearchDomain#warm_enabled}.
	WarmEnabled interface{} `json:"warmEnabled" yaml:"warmEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#warm_type ElasticsearchDomain#warm_type}.
	WarmType *string `json:"warmType" yaml:"warmType"`
	// zone_awareness_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#zone_awareness_config ElasticsearchDomain#zone_awareness_config}
	ZoneAwarenessConfig *ElasticsearchDomainClusterConfigZoneAwarenessConfig `json:"zoneAwarenessConfig" yaml:"zoneAwarenessConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#zone_awareness_enabled ElasticsearchDomain#zone_awareness_enabled}.
	ZoneAwarenessEnabled interface{} `json:"zoneAwarenessEnabled" yaml:"zoneAwarenessEnabled"`
}

type ElasticsearchDomainClusterConfigOutputReference interface {
	cdktf.ComplexObject
	DedicatedMasterCount() *float64
	SetDedicatedMasterCount(val *float64)
	DedicatedMasterCountInput() *float64
	DedicatedMasterEnabled() interface{}
	SetDedicatedMasterEnabled(val interface{})
	DedicatedMasterEnabledInput() interface{}
	DedicatedMasterType() *string
	SetDedicatedMasterType(val *string)
	DedicatedMasterTypeInput() *string
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() *ElasticsearchDomainClusterConfig
	SetInternalValue(val *ElasticsearchDomainClusterConfig)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WarmCount() *float64
	SetWarmCount(val *float64)
	WarmCountInput() *float64
	WarmEnabled() interface{}
	SetWarmEnabled(val interface{})
	WarmEnabledInput() interface{}
	WarmType() *string
	SetWarmType(val *string)
	WarmTypeInput() *string
	ZoneAwarenessConfig() ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference
	ZoneAwarenessConfigInput() *ElasticsearchDomainClusterConfigZoneAwarenessConfig
	ZoneAwarenessEnabled() interface{}
	SetZoneAwarenessEnabled(val interface{})
	ZoneAwarenessEnabledInput() interface{}
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
	PutZoneAwarenessConfig(value *ElasticsearchDomainClusterConfigZoneAwarenessConfig)
	ResetDedicatedMasterCount()
	ResetDedicatedMasterEnabled()
	ResetDedicatedMasterType()
	ResetInstanceCount()
	ResetInstanceType()
	ResetWarmCount()
	ResetWarmEnabled()
	ResetWarmType()
	ResetZoneAwarenessConfig()
	ResetZoneAwarenessEnabled()
}

// The jsii proxy struct for ElasticsearchDomainClusterConfigOutputReference
type jsiiProxy_ElasticsearchDomainClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dedicatedMasterCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedMasterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) DedicatedMasterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedMasterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InternalValue() *ElasticsearchDomainClusterConfig {
	var returns *ElasticsearchDomainClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) WarmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ZoneAwarenessConfig() ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference {
	var returns ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference
	_jsii_.Get(
		j,
		"zoneAwarenessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ZoneAwarenessConfigInput() *ElasticsearchDomainClusterConfigZoneAwarenessConfig {
	var returns *ElasticsearchDomainClusterConfigZoneAwarenessConfig
	_jsii_.Get(
		j,
		"zoneAwarenessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ZoneAwarenessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ZoneAwarenessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabledInput",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainClusterConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainClusterConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainClusterConfigOutputReference_Override(e ElasticsearchDomainClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetDedicatedMasterCount(val *float64) {
	_jsii_.Set(
		j,
		"dedicatedMasterCount",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetDedicatedMasterEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dedicatedMasterEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetDedicatedMasterType(val *string) {
	_jsii_.Set(
		j,
		"dedicatedMasterType",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetInternalValue(val *ElasticsearchDomainClusterConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetWarmCount(val *float64) {
	_jsii_.Set(
		j,
		"warmCount",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetWarmEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"warmEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetWarmType(val *string) {
	_jsii_.Set(
		j,
		"warmType",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetZoneAwarenessEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"zoneAwarenessEnabled",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) PutZoneAwarenessConfig(value *ElasticsearchDomainClusterConfigZoneAwarenessConfig) {
	_jsii_.InvokeVoid(
		e,
		"putZoneAwarenessConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetDedicatedMasterCount() {
	_jsii_.InvokeVoid(
		e,
		"resetDedicatedMasterCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetDedicatedMasterEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetDedicatedMasterEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetDedicatedMasterType() {
	_jsii_.InvokeVoid(
		e,
		"resetDedicatedMasterType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetInstanceCount() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetWarmCount() {
	_jsii_.InvokeVoid(
		e,
		"resetWarmCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetWarmEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetWarmEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetWarmType() {
	_jsii_.InvokeVoid(
		e,
		"resetWarmType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetZoneAwarenessConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetZoneAwarenessConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) ResetZoneAwarenessEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetZoneAwarenessEnabled",
		nil, // no parameters
	)
}

type ElasticsearchDomainClusterConfigZoneAwarenessConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#availability_zone_count ElasticsearchDomain#availability_zone_count}.
	AvailabilityZoneCount *float64 `json:"availabilityZoneCount" yaml:"availabilityZoneCount"`
}

type ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference interface {
	cdktf.ComplexObject
	AvailabilityZoneCount() *float64
	SetAvailabilityZoneCount(val *float64)
	AvailabilityZoneCountInput() *float64
	InternalValue() *ElasticsearchDomainClusterConfigZoneAwarenessConfig
	SetInternalValue(val *ElasticsearchDomainClusterConfigZoneAwarenessConfig)
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
	ResetAvailabilityZoneCount()
}

// The jsii proxy struct for ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference
type jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) AvailabilityZoneCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityZoneCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) AvailabilityZoneCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availabilityZoneCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) InternalValue() *ElasticsearchDomainClusterConfigZoneAwarenessConfig {
	var returns *ElasticsearchDomainClusterConfigZoneAwarenessConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference_Override(e ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) SetAvailabilityZoneCount(val *float64) {
	_jsii_.Set(
		j,
		"availabilityZoneCount",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) SetInternalValue(val *ElasticsearchDomainClusterConfigZoneAwarenessConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) ResetAvailabilityZoneCount() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZoneCount",
		nil, // no parameters
	)
}

type ElasticsearchDomainCognitoOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#identity_pool_id ElasticsearchDomain#identity_pool_id}.
	IdentityPoolId *string `json:"identityPoolId" yaml:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#role_arn ElasticsearchDomain#role_arn}.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#user_pool_id ElasticsearchDomain#user_pool_id}.
	UserPoolId *string `json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

type ElasticsearchDomainCognitoOptionsOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	IdentityPoolId() *string
	SetIdentityPoolId(val *string)
	IdentityPoolIdInput() *string
	InternalValue() *ElasticsearchDomainCognitoOptions
	SetInternalValue(val *ElasticsearchDomainCognitoOptions)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	ResetEnabled()
}

// The jsii proxy struct for ElasticsearchDomainCognitoOptionsOutputReference
type jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) IdentityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) InternalValue() *ElasticsearchDomainCognitoOptions {
	var returns *ElasticsearchDomainCognitoOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainCognitoOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainCognitoOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainCognitoOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainCognitoOptionsOutputReference_Override(e ElasticsearchDomainCognitoOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainCognitoOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) SetIdentityPoolId(val *string) {
	_jsii_.Set(
		j,
		"identityPoolId",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) SetInternalValue(val *ElasticsearchDomainCognitoOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) SetUserPoolId(val *string) {
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetEnabled",
		nil, // no parameters
	)
}

// AWS ElasticSearch Service.
type ElasticsearchDomainConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#domain_name ElasticsearchDomain#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#access_policies ElasticsearchDomain#access_policies}.
	AccessPolicies *string `json:"accessPolicies" yaml:"accessPolicies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#advanced_options ElasticsearchDomain#advanced_options}.
	AdvancedOptions *map[string]*string `json:"advancedOptions" yaml:"advancedOptions"`
	// advanced_security_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#advanced_security_options ElasticsearchDomain#advanced_security_options}
	AdvancedSecurityOptions *ElasticsearchDomainAdvancedSecurityOptions `json:"advancedSecurityOptions" yaml:"advancedSecurityOptions"`
	// auto_tune_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#auto_tune_options ElasticsearchDomain#auto_tune_options}
	AutoTuneOptions *ElasticsearchDomainAutoTuneOptions `json:"autoTuneOptions" yaml:"autoTuneOptions"`
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#cluster_config ElasticsearchDomain#cluster_config}
	ClusterConfig *ElasticsearchDomainClusterConfig `json:"clusterConfig" yaml:"clusterConfig"`
	// cognito_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#cognito_options ElasticsearchDomain#cognito_options}
	CognitoOptions *ElasticsearchDomainCognitoOptions `json:"cognitoOptions" yaml:"cognitoOptions"`
	// domain_endpoint_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#domain_endpoint_options ElasticsearchDomain#domain_endpoint_options}
	DomainEndpointOptions *ElasticsearchDomainDomainEndpointOptions `json:"domainEndpointOptions" yaml:"domainEndpointOptions"`
	// ebs_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#ebs_options ElasticsearchDomain#ebs_options}
	EbsOptions *ElasticsearchDomainEbsOptions `json:"ebsOptions" yaml:"ebsOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#elasticsearch_version ElasticsearchDomain#elasticsearch_version}.
	ElasticsearchVersion *string `json:"elasticsearchVersion" yaml:"elasticsearchVersion"`
	// encrypt_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#encrypt_at_rest ElasticsearchDomain#encrypt_at_rest}
	EncryptAtRest *ElasticsearchDomainEncryptAtRest `json:"encryptAtRest" yaml:"encryptAtRest"`
	// log_publishing_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#log_publishing_options ElasticsearchDomain#log_publishing_options}
	LogPublishingOptions interface{} `json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// node_to_node_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#node_to_node_encryption ElasticsearchDomain#node_to_node_encryption}
	NodeToNodeEncryption *ElasticsearchDomainNodeToNodeEncryption `json:"nodeToNodeEncryption" yaml:"nodeToNodeEncryption"`
	// snapshot_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#snapshot_options ElasticsearchDomain#snapshot_options}
	SnapshotOptions *ElasticsearchDomainSnapshotOptions `json:"snapshotOptions" yaml:"snapshotOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#tags ElasticsearchDomain#tags}.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#tags_all ElasticsearchDomain#tags_all}.
	TagsAll *map[string]*string `json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#timeouts ElasticsearchDomain#timeouts}
	Timeouts *ElasticsearchDomainTimeouts `json:"timeouts" yaml:"timeouts"`
	// vpc_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#vpc_options ElasticsearchDomain#vpc_options}
	VpcOptions *ElasticsearchDomainVpcOptions `json:"vpcOptions" yaml:"vpcOptions"`
}

type ElasticsearchDomainDomainEndpointOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#custom_endpoint ElasticsearchDomain#custom_endpoint}.
	CustomEndpoint *string `json:"customEndpoint" yaml:"customEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#custom_endpoint_certificate_arn ElasticsearchDomain#custom_endpoint_certificate_arn}.
	CustomEndpointCertificateArn *string `json:"customEndpointCertificateArn" yaml:"customEndpointCertificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#custom_endpoint_enabled ElasticsearchDomain#custom_endpoint_enabled}.
	CustomEndpointEnabled interface{} `json:"customEndpointEnabled" yaml:"customEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#enforce_https ElasticsearchDomain#enforce_https}.
	EnforceHttps interface{} `json:"enforceHttps" yaml:"enforceHttps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#tls_security_policy ElasticsearchDomain#tls_security_policy}.
	TlsSecurityPolicy *string `json:"tlsSecurityPolicy" yaml:"tlsSecurityPolicy"`
}

type ElasticsearchDomainDomainEndpointOptionsOutputReference interface {
	cdktf.ComplexObject
	CustomEndpoint() *string
	SetCustomEndpoint(val *string)
	CustomEndpointCertificateArn() *string
	SetCustomEndpointCertificateArn(val *string)
	CustomEndpointCertificateArnInput() *string
	CustomEndpointEnabled() interface{}
	SetCustomEndpointEnabled(val interface{})
	CustomEndpointEnabledInput() interface{}
	CustomEndpointInput() *string
	EnforceHttps() interface{}
	SetEnforceHttps(val interface{})
	EnforceHttpsInput() interface{}
	InternalValue() *ElasticsearchDomainDomainEndpointOptions
	SetInternalValue(val *ElasticsearchDomainDomainEndpointOptions)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsSecurityPolicy() *string
	SetTlsSecurityPolicy(val *string)
	TlsSecurityPolicyInput() *string
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
	ResetCustomEndpoint()
	ResetCustomEndpointCertificateArn()
	ResetCustomEndpointEnabled()
	ResetEnforceHttps()
	ResetTlsSecurityPolicy()
}

// The jsii proxy struct for ElasticsearchDomainDomainEndpointOptionsOutputReference
type jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpointCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpointCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) EnforceHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) EnforceHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) InternalValue() *ElasticsearchDomainDomainEndpointOptions {
	var returns *ElasticsearchDomainDomainEndpointOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) TlsSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) TlsSecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSecurityPolicyInput",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainDomainEndpointOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainDomainEndpointOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainDomainEndpointOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainDomainEndpointOptionsOutputReference_Override(e ElasticsearchDomainDomainEndpointOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainDomainEndpointOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetCustomEndpoint(val *string) {
	_jsii_.Set(
		j,
		"customEndpoint",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetCustomEndpointCertificateArn(val *string) {
	_jsii_.Set(
		j,
		"customEndpointCertificateArn",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetCustomEndpointEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"customEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetEnforceHttps(val interface{}) {
	_jsii_.Set(
		j,
		"enforceHttps",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetInternalValue(val *ElasticsearchDomainDomainEndpointOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetTlsSecurityPolicy(val *string) {
	_jsii_.Set(
		j,
		"tlsSecurityPolicy",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetCustomEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetCustomEndpointCertificateArn() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomEndpointCertificateArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetCustomEndpointEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomEndpointEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetEnforceHttps() {
	_jsii_.InvokeVoid(
		e,
		"resetEnforceHttps",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) ResetTlsSecurityPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetTlsSecurityPolicy",
		nil, // no parameters
	)
}

type ElasticsearchDomainEbsOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#ebs_enabled ElasticsearchDomain#ebs_enabled}.
	EbsEnabled interface{} `json:"ebsEnabled" yaml:"ebsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#iops ElasticsearchDomain#iops}.
	Iops *float64 `json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#volume_size ElasticsearchDomain#volume_size}.
	VolumeSize *float64 `json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#volume_type ElasticsearchDomain#volume_type}.
	VolumeType *string `json:"volumeType" yaml:"volumeType"`
}

type ElasticsearchDomainEbsOptionsOutputReference interface {
	cdktf.ComplexObject
	EbsEnabled() interface{}
	SetEbsEnabled(val interface{})
	EbsEnabledInput() interface{}
	InternalValue() *ElasticsearchDomainEbsOptions
	SetInternalValue(val *ElasticsearchDomainEbsOptions)
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeSize() *float64
	SetVolumeSize(val *float64)
	VolumeSizeInput() *float64
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
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
	ResetIops()
	ResetVolumeSize()
	ResetVolumeType()
}

// The jsii proxy struct for ElasticsearchDomainEbsOptionsOutputReference
type jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) EbsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) EbsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) InternalValue() *ElasticsearchDomainEbsOptions {
	var returns *ElasticsearchDomainEbsOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) VolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) VolumeSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainEbsOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainEbsOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainEbsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainEbsOptionsOutputReference_Override(e ElasticsearchDomainEbsOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainEbsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) SetEbsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"ebsEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) SetInternalValue(val *ElasticsearchDomainEbsOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) SetIops(val *float64) {
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) SetVolumeSize(val *float64) {
	_jsii_.Set(
		j,
		"volumeSize",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) SetVolumeType(val *string) {
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) ResetIops() {
	_jsii_.InvokeVoid(
		e,
		"resetIops",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) ResetVolumeSize() {
	_jsii_.InvokeVoid(
		e,
		"resetVolumeSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) ResetVolumeType() {
	_jsii_.InvokeVoid(
		e,
		"resetVolumeType",
		nil, // no parameters
	)
}

type ElasticsearchDomainEncryptAtRest struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#kms_key_id ElasticsearchDomain#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
}

type ElasticsearchDomainEncryptAtRestOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *ElasticsearchDomainEncryptAtRest
	SetInternalValue(val *ElasticsearchDomainEncryptAtRest)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
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
	ResetKmsKeyId()
}

// The jsii proxy struct for ElasticsearchDomainEncryptAtRestOutputReference
type jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) InternalValue() *ElasticsearchDomainEncryptAtRest {
	var returns *ElasticsearchDomainEncryptAtRest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainEncryptAtRestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainEncryptAtRestOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainEncryptAtRestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainEncryptAtRestOutputReference_Override(e ElasticsearchDomainEncryptAtRestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainEncryptAtRestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) SetInternalValue(val *ElasticsearchDomainEncryptAtRest) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		e,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

type ElasticsearchDomainLogPublishingOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#cloudwatch_log_group_arn ElasticsearchDomain#cloudwatch_log_group_arn}.
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#log_type ElasticsearchDomain#log_type}.
	LogType *string `json:"logType" yaml:"logType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

type ElasticsearchDomainNodeToNodeEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

type ElasticsearchDomainNodeToNodeEncryptionOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	InternalValue() *ElasticsearchDomainNodeToNodeEncryption
	SetInternalValue(val *ElasticsearchDomainNodeToNodeEncryption)
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

// The jsii proxy struct for ElasticsearchDomainNodeToNodeEncryptionOutputReference
type jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) InternalValue() *ElasticsearchDomainNodeToNodeEncryption {
	var returns *ElasticsearchDomainNodeToNodeEncryption
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainNodeToNodeEncryptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainNodeToNodeEncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainNodeToNodeEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainNodeToNodeEncryptionOutputReference_Override(e ElasticsearchDomainNodeToNodeEncryptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainNodeToNodeEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) SetInternalValue(val *ElasticsearchDomainNodeToNodeEncryption) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy aws_elasticsearch_domain_policy}.
type ElasticsearchDomainPolicy interface {
	cdktf.TerraformResource
	AccessPolicies() *string
	SetAccessPolicies(val *string)
	AccessPoliciesInput() *string
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

// The jsii proxy struct for ElasticsearchDomainPolicy
type jsiiProxy_ElasticsearchDomainPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) AccessPoliciesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy aws_elasticsearch_domain_policy} Resource.
func NewElasticsearchDomainPolicy(scope constructs.Construct, id *string, config *ElasticsearchDomainPolicyConfig) ElasticsearchDomainPolicy {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainPolicy{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy aws_elasticsearch_domain_policy} Resource.
func NewElasticsearchDomainPolicy_Override(e ElasticsearchDomainPolicy, scope constructs.Construct, id *string, config *ElasticsearchDomainPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainPolicy",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) SetAccessPolicies(val *string) {
	_jsii_.Set(
		j,
		"accessPolicies",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainPolicy) SetProvider(val cdktf.TerraformProvider) {
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
func ElasticsearchDomainPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticsearchDomainPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_ElasticsearchDomainPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_ElasticsearchDomainPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS ElasticSearch Service.
type ElasticsearchDomainPolicyConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy#access_policies ElasticsearchDomainPolicy#access_policies}.
	AccessPolicies *string `json:"accessPolicies" yaml:"accessPolicies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy#domain_name ElasticsearchDomainPolicy#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options aws_elasticsearch_domain_saml_options}.
type ElasticsearchDomainSamlOptions interface {
	cdktf.TerraformResource
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
	SamlOptions() ElasticsearchDomainSamlOptionsSamlOptionsOutputReference
	SamlOptionsInput() *ElasticsearchDomainSamlOptionsSamlOptions
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
	PutSamlOptions(value *ElasticsearchDomainSamlOptionsSamlOptions)
	ResetOverrideLogicalId()
	ResetSamlOptions()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ElasticsearchDomainSamlOptions
type jsiiProxy_ElasticsearchDomainSamlOptions struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) SamlOptions() ElasticsearchDomainSamlOptionsSamlOptionsOutputReference {
	var returns ElasticsearchDomainSamlOptionsSamlOptionsOutputReference
	_jsii_.Get(
		j,
		"samlOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) SamlOptionsInput() *ElasticsearchDomainSamlOptionsSamlOptions {
	var returns *ElasticsearchDomainSamlOptionsSamlOptions
	_jsii_.Get(
		j,
		"samlOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options aws_elasticsearch_domain_saml_options} Resource.
func NewElasticsearchDomainSamlOptions(scope constructs.Construct, id *string, config *ElasticsearchDomainSamlOptionsConfig) ElasticsearchDomainSamlOptions {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainSamlOptions{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSamlOptions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options aws_elasticsearch_domain_saml_options} Resource.
func NewElasticsearchDomainSamlOptions_Override(e ElasticsearchDomainSamlOptions, scope constructs.Construct, id *string, config *ElasticsearchDomainSamlOptionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSamlOptions",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) SetProvider(val cdktf.TerraformProvider) {
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
func ElasticsearchDomainSamlOptions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSamlOptions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticsearchDomainSamlOptions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSamlOptions",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptions) PutSamlOptions(value *ElasticsearchDomainSamlOptionsSamlOptions) {
	_jsii_.InvokeVoid(
		e,
		"putSamlOptions",
		[]interface{}{value},
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptions) ResetSamlOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetSamlOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptions) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// AWS ElasticSearch Service.
type ElasticsearchDomainSamlOptionsConfig struct {
	// Experimental.
	Count *float64 `json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#domain_name ElasticsearchDomainSamlOptions#domain_name}.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// saml_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#saml_options ElasticsearchDomainSamlOptions#saml_options}
	SamlOptions *ElasticsearchDomainSamlOptionsSamlOptions `json:"samlOptions" yaml:"samlOptions"`
}

type ElasticsearchDomainSamlOptionsSamlOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#enabled ElasticsearchDomainSamlOptions#enabled}.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// idp block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#idp ElasticsearchDomainSamlOptions#idp}
	Idp *ElasticsearchDomainSamlOptionsSamlOptionsIdp `json:"idp" yaml:"idp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#master_backend_role ElasticsearchDomainSamlOptions#master_backend_role}.
	MasterBackendRole *string `json:"masterBackendRole" yaml:"masterBackendRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#master_user_name ElasticsearchDomainSamlOptions#master_user_name}.
	MasterUserName *string `json:"masterUserName" yaml:"masterUserName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#roles_key ElasticsearchDomainSamlOptions#roles_key}.
	RolesKey *string `json:"rolesKey" yaml:"rolesKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#session_timeout_minutes ElasticsearchDomainSamlOptions#session_timeout_minutes}.
	SessionTimeoutMinutes *float64 `json:"sessionTimeoutMinutes" yaml:"sessionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#subject_key ElasticsearchDomainSamlOptions#subject_key}.
	SubjectKey *string `json:"subjectKey" yaml:"subjectKey"`
}

type ElasticsearchDomainSamlOptionsSamlOptionsIdp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#entity_id ElasticsearchDomainSamlOptions#entity_id}.
	EntityId *string `json:"entityId" yaml:"entityId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options#metadata_content ElasticsearchDomainSamlOptions#metadata_content}.
	MetadataContent *string `json:"metadataContent" yaml:"metadataContent"`
}

type ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference interface {
	cdktf.ComplexObject
	EntityId() *string
	SetEntityId(val *string)
	EntityIdInput() *string
	InternalValue() *ElasticsearchDomainSamlOptionsSamlOptionsIdp
	SetInternalValue(val *ElasticsearchDomainSamlOptionsSamlOptionsIdp)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MetadataContent() *string
	SetMetadataContent(val *string)
	MetadataContentInput() *string
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

// The jsii proxy struct for ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference
type jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) EntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) EntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) InternalValue() *ElasticsearchDomainSamlOptionsSamlOptionsIdp {
	var returns *ElasticsearchDomainSamlOptionsSamlOptionsIdp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) MetadataContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) MetadataContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference_Override(e ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetEntityId(val *string) {
	_jsii_.Set(
		j,
		"entityId",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetInternalValue(val *ElasticsearchDomainSamlOptionsSamlOptionsIdp) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetMetadataContent(val *string) {
	_jsii_.Set(
		j,
		"metadataContent",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type ElasticsearchDomainSamlOptionsSamlOptionsOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Idp() ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference
	IdpInput() *ElasticsearchDomainSamlOptionsSamlOptionsIdp
	InternalValue() *ElasticsearchDomainSamlOptionsSamlOptions
	SetInternalValue(val *ElasticsearchDomainSamlOptionsSamlOptions)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MasterBackendRole() *string
	SetMasterBackendRole(val *string)
	MasterBackendRoleInput() *string
	MasterUserName() *string
	SetMasterUserName(val *string)
	MasterUserNameInput() *string
	RolesKey() *string
	SetRolesKey(val *string)
	RolesKeyInput() *string
	SessionTimeoutMinutes() *float64
	SetSessionTimeoutMinutes(val *float64)
	SessionTimeoutMinutesInput() *float64
	SubjectKey() *string
	SetSubjectKey(val *string)
	SubjectKeyInput() *string
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
	PutIdp(value *ElasticsearchDomainSamlOptionsSamlOptionsIdp)
	ResetEnabled()
	ResetIdp()
	ResetMasterBackendRole()
	ResetMasterUserName()
	ResetRolesKey()
	ResetSessionTimeoutMinutes()
	ResetSubjectKey()
}

// The jsii proxy struct for ElasticsearchDomainSamlOptionsSamlOptionsOutputReference
type jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) Idp() ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference {
	var returns ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference
	_jsii_.Get(
		j,
		"idp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) IdpInput() *ElasticsearchDomainSamlOptionsSamlOptionsIdp {
	var returns *ElasticsearchDomainSamlOptionsSamlOptionsIdp
	_jsii_.Get(
		j,
		"idpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) InternalValue() *ElasticsearchDomainSamlOptionsSamlOptions {
	var returns *ElasticsearchDomainSamlOptionsSamlOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) MasterBackendRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterBackendRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) MasterBackendRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterBackendRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) MasterUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) MasterUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) RolesKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) RolesKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SessionTimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SessionTimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SubjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SubjectKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subjectKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainSamlOptionsSamlOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainSamlOptionsSamlOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSamlOptionsSamlOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainSamlOptionsSamlOptionsOutputReference_Override(e ElasticsearchDomainSamlOptionsSamlOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSamlOptionsSamlOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetInternalValue(val *ElasticsearchDomainSamlOptionsSamlOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetMasterBackendRole(val *string) {
	_jsii_.Set(
		j,
		"masterBackendRole",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetMasterUserName(val *string) {
	_jsii_.Set(
		j,
		"masterUserName",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetRolesKey(val *string) {
	_jsii_.Set(
		j,
		"rolesKey",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetSessionTimeoutMinutes(val *float64) {
	_jsii_.Set(
		j,
		"sessionTimeoutMinutes",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetSubjectKey(val *string) {
	_jsii_.Set(
		j,
		"subjectKey",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) PutIdp(value *ElasticsearchDomainSamlOptionsSamlOptionsIdp) {
	_jsii_.InvokeVoid(
		e,
		"putIdp",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetIdp() {
	_jsii_.InvokeVoid(
		e,
		"resetIdp",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetMasterBackendRole() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterBackendRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetMasterUserName() {
	_jsii_.InvokeVoid(
		e,
		"resetMasterUserName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetRolesKey() {
	_jsii_.InvokeVoid(
		e,
		"resetRolesKey",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetSessionTimeoutMinutes() {
	_jsii_.InvokeVoid(
		e,
		"resetSessionTimeoutMinutes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) ResetSubjectKey() {
	_jsii_.InvokeVoid(
		e,
		"resetSubjectKey",
		nil, // no parameters
	)
}

type ElasticsearchDomainSnapshotOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#automated_snapshot_start_hour ElasticsearchDomain#automated_snapshot_start_hour}.
	AutomatedSnapshotStartHour *float64 `json:"automatedSnapshotStartHour" yaml:"automatedSnapshotStartHour"`
}

type ElasticsearchDomainSnapshotOptionsOutputReference interface {
	cdktf.ComplexObject
	AutomatedSnapshotStartHour() *float64
	SetAutomatedSnapshotStartHour(val *float64)
	AutomatedSnapshotStartHourInput() *float64
	InternalValue() *ElasticsearchDomainSnapshotOptions
	SetInternalValue(val *ElasticsearchDomainSnapshotOptions)
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

// The jsii proxy struct for ElasticsearchDomainSnapshotOptionsOutputReference
type jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) AutomatedSnapshotStartHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotStartHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) AutomatedSnapshotStartHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotStartHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) InternalValue() *ElasticsearchDomainSnapshotOptions {
	var returns *ElasticsearchDomainSnapshotOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainSnapshotOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainSnapshotOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSnapshotOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainSnapshotOptionsOutputReference_Override(e ElasticsearchDomainSnapshotOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainSnapshotOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) SetAutomatedSnapshotStartHour(val *float64) {
	_jsii_.Set(
		j,
		"automatedSnapshotStartHour",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) SetInternalValue(val *ElasticsearchDomainSnapshotOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type ElasticsearchDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#update ElasticsearchDomain#update}.
	Update *string `json:"update" yaml:"update"`
}

type ElasticsearchDomainTimeoutsOutputReference interface {
	cdktf.ComplexObject
	InternalValue() *ElasticsearchDomainTimeouts
	SetInternalValue(val *ElasticsearchDomainTimeouts)
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetUpdate()
}

// The jsii proxy struct for ElasticsearchDomainTimeoutsOutputReference
type jsiiProxy_ElasticsearchDomainTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) InternalValue() *ElasticsearchDomainTimeouts {
	var returns *ElasticsearchDomainTimeouts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainTimeoutsOutputReference_Override(e ElasticsearchDomainTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) SetInternalValue(val *ElasticsearchDomainTimeouts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		e,
		"resetUpdate",
		nil, // no parameters
	)
}

type ElasticsearchDomainVpcOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#security_group_ids ElasticsearchDomain#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain#subnet_ids ElasticsearchDomain#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

type ElasticsearchDomainVpcOptionsOutputReference interface {
	cdktf.ComplexObject
	AvailabilityZones() *[]*string
	InternalValue() *ElasticsearchDomainVpcOptions
	SetInternalValue(val *ElasticsearchDomainVpcOptions)
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
	VpcId() *string
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

// The jsii proxy struct for ElasticsearchDomainVpcOptionsOutputReference
type jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) InternalValue() *ElasticsearchDomainVpcOptions {
	var returns *ElasticsearchDomainVpcOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainVpcOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainVpcOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainVpcOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainVpcOptionsOutputReference_Override(e ElasticsearchDomainVpcOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.elasticsearch.ElasticsearchDomainVpcOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SetInternalValue(val *ElasticsearchDomainVpcOptions) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetIds",
		nil, // no parameters
	)
}
