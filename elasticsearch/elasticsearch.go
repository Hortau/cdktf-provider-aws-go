package elasticsearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/elasticsearch/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain.html aws_elasticsearch_domain}.
type DataAwsElasticsearchDomain interface {
	cdktf.TerraformDataSource
	AccessPolicies() *string
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	Created() interface{}
	Deleted() interface{}
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
	Processing() interface{}
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
	AdvancedOptions(key *string) *string
	AdvancedSecurityOptions(index *string) DataAwsElasticsearchDomainAdvancedSecurityOptions
	ClusterConfig(index *string) DataAwsElasticsearchDomainClusterConfig
	CognitoOptions(index *string) DataAwsElasticsearchDomainCognitoOptions
	EbsOptions(index *string) DataAwsElasticsearchDomainEbsOptions
	EncryptionAtRest(index *string) DataAwsElasticsearchDomainEncryptionAtRest
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_DataAwsElasticsearchDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Created() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) Deleted() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomain) Processing() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomain) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) TagsInput() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain.html aws_elasticsearch_domain} Data Source.
func NewDataAwsElasticsearchDomain(scope constructs.Construct, id *string, config *DataAwsElasticsearchDomainConfig) DataAwsElasticsearchDomain {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomain{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain.html aws_elasticsearch_domain} Data Source.
func NewDataAwsElasticsearchDomain_Override(d DataAwsElasticsearchDomain, scope constructs.Construct, id *string, config *DataAwsElasticsearchDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomain",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetCount(val interface{}) {
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

func (j *jsiiProxy_DataAwsElasticsearchDomain) SetTags(val interface{}) {
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
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomain",
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
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomain",
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

func (d *jsiiProxy_DataAwsElasticsearchDomain) AdvancedOptions(key *string) *string {
	var returns *string

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
	Enabled() interface{}
	InternalUserDatabaseEnabled() interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) InternalUserDatabaseEnabled() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainAdvancedSecurityOptions(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainAdvancedSecurityOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainAdvancedSecurityOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainAdvancedSecurityOptions_Override(d DataAwsElasticsearchDomainAdvancedSecurityOptions, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainAdvancedSecurityOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptions) GetBooleanAttribute(terraformAttribute *string) interface{} {
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

type DataAwsElasticsearchDomainClusterConfig interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	DedicatedMasterCount() *float64
	DedicatedMasterEnabled() interface{}
	DedicatedMasterType() *string
	InstanceCount() *float64
	InstanceType() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	WarmCount() *float64
	WarmEnabled() interface{}
	WarmType() *string
	ZoneAwarenessConfig() interface{}
	ZoneAwarenessEnabled() interface{}
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) DedicatedMasterEnabled() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) WarmEnabled() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) ZoneAwarenessConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) ZoneAwarenessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneAwarenessEnabled",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainClusterConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainClusterConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainClusterConfig{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainClusterConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainClusterConfig_Override(d DataAwsElasticsearchDomainClusterConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainClusterConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig_Override(d DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainClusterConfigZoneAwarenessConfig) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	Enabled() interface{}
	IdentityPoolId() *string
	RoleArn() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UserPoolId() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) Enabled() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

// Experimental.
func NewDataAwsElasticsearchDomainCognitoOptions(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainCognitoOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainCognitoOptions{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainCognitoOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainCognitoOptions_Override(d DataAwsElasticsearchDomainCognitoOptions, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainCognitoOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainCognitoOptions) GetBooleanAttribute(terraformAttribute *string) interface{} {
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

type DataAwsElasticsearchDomainConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain.html#domain_name DataAwsElasticsearchDomain#domain_name}.
	DomainName *string `json:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/elasticsearch_domain.html#tags DataAwsElasticsearchDomain#tags}.
	Tags interface{} `json:"tags"`
}

type DataAwsElasticsearchDomainEbsOptions interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	EbsEnabled() interface{}
	Iops() *float64
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	VolumeSize() *float64
	VolumeType() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) EbsEnabled() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

// Experimental.
func NewDataAwsElasticsearchDomainEbsOptions(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainEbsOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainEbsOptions{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainEbsOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainEbsOptions_Override(d DataAwsElasticsearchDomainEbsOptions, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainEbsOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainEbsOptions) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	Enabled() interface{}
	KmsKeyId() *string
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

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) Enabled() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainEncryptionAtRest(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainEncryptionAtRest {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainEncryptionAtRest",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainEncryptionAtRest_Override(d DataAwsElasticsearchDomainEncryptionAtRest, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainEncryptionAtRest",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRest) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	Enabled() interface{}
	LogType() *string
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

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) Enabled() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainLogPublishingOptions(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainLogPublishingOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainLogPublishingOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainLogPublishingOptions_Override(d DataAwsElasticsearchDomainLogPublishingOptions, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainLogPublishingOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainLogPublishingOptions) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	Enabled() interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) Enabled() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainNodeToNodeEncryption(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainNodeToNodeEncryption {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainNodeToNodeEncryption",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainNodeToNodeEncryption_Override(d DataAwsElasticsearchDomainNodeToNodeEncryption, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainNodeToNodeEncryption",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainNodeToNodeEncryption) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsElasticsearchDomainSnapshotOptions(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainSnapshotOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainSnapshotOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainSnapshotOptions_Override(d DataAwsElasticsearchDomainSnapshotOptions, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainSnapshotOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptions) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	VpcId() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

// Experimental.
func NewDataAwsElasticsearchDomainVpcOptions(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsElasticsearchDomainVpcOptions {
	_init_.Initialize()

	j := jsiiProxy_DataAwsElasticsearchDomainVpcOptions{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainVpcOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsElasticsearchDomainVpcOptions_Override(d DataAwsElasticsearchDomainVpcOptions, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.DataAwsElasticsearchDomainVpcOptions",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
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

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptions) GetBooleanAttribute(terraformAttribute *string) interface{} {
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

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html aws_elasticsearch_domain}.
type ElasticsearchDomain interface {
	cdktf.TerraformResource
	AccessPolicies() *string
	SetAccessPolicies(val *string)
	AccessPoliciesInput() *string
	AdvancedOptions() interface{}
	SetAdvancedOptions(val interface{})
	AdvancedOptionsInput() interface{}
	AdvancedSecurityOptions() ElasticsearchDomainAdvancedSecurityOptionsOutputReference
	AdvancedSecurityOptionsInput() *ElasticsearchDomainAdvancedSecurityOptions
	Arn() *string
	CdktfStack() cdktf.TerraformStack
	ClusterConfig() ElasticsearchDomainClusterConfigOutputReference
	ClusterConfigInput() *ElasticsearchDomainClusterConfig
	CognitoOptions() ElasticsearchDomainCognitoOptionsOutputReference
	CognitoOptionsInput() *ElasticsearchDomainCognitoOptions
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	LogPublishingOptions() *[]*ElasticsearchDomainLogPublishingOptions
	SetLogPublishingOptions(val *[]*ElasticsearchDomainLogPublishingOptions)
	LogPublishingOptionsInput() *[]*ElasticsearchDomainLogPublishingOptions
	Node() constructs.Node
	NodeToNodeEncryption() ElasticsearchDomainNodeToNodeEncryptionOutputReference
	NodeToNodeEncryptionInput() *ElasticsearchDomainNodeToNodeEncryption
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	SnapshotOptions() ElasticsearchDomainSnapshotOptionsOutputReference
	SnapshotOptionsInput() *ElasticsearchDomainSnapshotOptions
	Tags() interface{}
	SetTags(val interface{})
	TagsAll() interface{}
	SetTagsAll(val interface{})
	TagsAllInput() interface{}
	TagsInput() interface{}
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformMetaArguments() *map[string]interface{}
	TerraformResourceType() *string
	Timeouts() ElasticsearchDomainTimeoutsOutputReference
	TimeoutsInput() *ElasticsearchDomainTimeouts
	VpcOptions() ElasticsearchDomainVpcOptionsOutputReference
	VpcOptionsInput() *ElasticsearchDomainVpcOptions
	AddOverride(path *string, value interface{})
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	OverrideLogicalId(newLogicalId *string)
	PutAdvancedSecurityOptions(value *ElasticsearchDomainAdvancedSecurityOptions)
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

func (j *jsiiProxy_ElasticsearchDomain) AdvancedOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) AdvancedOptionsInput() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ElasticsearchDomain) Count() interface{} {
	var returns interface{}
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

func (j *jsiiProxy_ElasticsearchDomain) LogPublishingOptions() *[]*ElasticsearchDomainLogPublishingOptions {
	var returns *[]*ElasticsearchDomainLogPublishingOptions
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) LogPublishingOptionsInput() *[]*ElasticsearchDomainLogPublishingOptions {
	var returns *[]*ElasticsearchDomainLogPublishingOptions
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

func (j *jsiiProxy_ElasticsearchDomain) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticsearchDomain) TagsInput() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html aws_elasticsearch_domain} Resource.
func NewElasticsearchDomain(scope constructs.Construct, id *string, config *ElasticsearchDomainConfig) ElasticsearchDomain {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomain{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html aws_elasticsearch_domain} Resource.
func NewElasticsearchDomain_Override(e ElasticsearchDomain, scope constructs.Construct, id *string, config *ElasticsearchDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomain",
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

func (j *jsiiProxy_ElasticsearchDomain) SetAdvancedOptions(val interface{}) {
	_jsii_.Set(
		j,
		"advancedOptions",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetCount(val interface{}) {
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

func (j *jsiiProxy_ElasticsearchDomain) SetLogPublishingOptions(val *[]*ElasticsearchDomainLogPublishingOptions) {
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

func (j *jsiiProxy_ElasticsearchDomain) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ElasticsearchDomain) SetTagsAll(val interface{}) {
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
		"hashicorp_aws.ElasticSearch.ElasticsearchDomain",
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
		"hashicorp_aws.ElasticSearch.ElasticsearchDomain",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#internal_user_database_enabled ElasticsearchDomain#internal_user_database_enabled}.
	InternalUserDatabaseEnabled interface{} `json:"internalUserDatabaseEnabled"`
	// master_user_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#master_user_options ElasticsearchDomain#master_user_options}
	MasterUserOptions *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions `json:"masterUserOptions"`
}

type ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#master_user_arn ElasticsearchDomain#master_user_arn}.
	MasterUserArn *string `json:"masterUserArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#master_user_name ElasticsearchDomain#master_user_name}.
	MasterUserName *string `json:"masterUserName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#master_user_password ElasticsearchDomain#master_user_password}.
	MasterUserPassword *string `json:"masterUserPassword"`
}

type ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference interface {
	cdktf.ComplexObject
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference_Override(e ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
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

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MasterUserOptions() ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
	MasterUserOptionsInput() *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions
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

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainAdvancedSecurityOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainAdvancedSecurityOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainAdvancedSecurityOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainAdvancedSecurityOptionsOutputReference_Override(e ElasticsearchDomainAdvancedSecurityOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainAdvancedSecurityOptionsOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainAdvancedSecurityOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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

type ElasticsearchDomainClusterConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#dedicated_master_count ElasticsearchDomain#dedicated_master_count}.
	DedicatedMasterCount *float64 `json:"dedicatedMasterCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#dedicated_master_enabled ElasticsearchDomain#dedicated_master_enabled}.
	DedicatedMasterEnabled interface{} `json:"dedicatedMasterEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#dedicated_master_type ElasticsearchDomain#dedicated_master_type}.
	DedicatedMasterType *string `json:"dedicatedMasterType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#instance_count ElasticsearchDomain#instance_count}.
	InstanceCount *float64 `json:"instanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#instance_type ElasticsearchDomain#instance_type}.
	InstanceType *string `json:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#warm_count ElasticsearchDomain#warm_count}.
	WarmCount *float64 `json:"warmCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#warm_enabled ElasticsearchDomain#warm_enabled}.
	WarmEnabled interface{} `json:"warmEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#warm_type ElasticsearchDomain#warm_type}.
	WarmType *string `json:"warmType"`
	// zone_awareness_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#zone_awareness_config ElasticsearchDomain#zone_awareness_config}
	ZoneAwarenessConfig *ElasticsearchDomainClusterConfigZoneAwarenessConfig `json:"zoneAwarenessConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#zone_awareness_enabled ElasticsearchDomain#zone_awareness_enabled}.
	ZoneAwarenessEnabled interface{} `json:"zoneAwarenessEnabled"`
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
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
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
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func NewElasticsearchDomainClusterConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainClusterConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainClusterConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainClusterConfigOutputReference_Override(e ElasticsearchDomainClusterConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainClusterConfigOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (e *jsiiProxy_ElasticsearchDomainClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#availability_zone_count ElasticsearchDomain#availability_zone_count}.
	AvailabilityZoneCount *float64 `json:"availabilityZoneCount"`
}

type ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference interface {
	cdktf.ComplexObject
	AvailabilityZoneCount() *float64
	SetAvailabilityZoneCount(val *float64)
	AvailabilityZoneCountInput() *float64
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

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference_Override(e ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainClusterConfigZoneAwarenessConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#identity_pool_id ElasticsearchDomain#identity_pool_id}.
	IdentityPoolId *string `json:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#role_arn ElasticsearchDomain#role_arn}.
	RoleArn *string `json:"roleArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#user_pool_id ElasticsearchDomain#user_pool_id}.
	UserPoolId *string `json:"userPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled"`
}

type ElasticsearchDomainCognitoOptionsOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	IdentityPoolId() *string
	SetIdentityPoolId(val *string)
	IdentityPoolIdInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func NewElasticsearchDomainCognitoOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainCognitoOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainCognitoOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainCognitoOptionsOutputReference_Override(e ElasticsearchDomainCognitoOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainCognitoOptionsOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (e *jsiiProxy_ElasticsearchDomainCognitoOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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

type ElasticsearchDomainConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#domain_name ElasticsearchDomain#domain_name}.
	DomainName *string `json:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#access_policies ElasticsearchDomain#access_policies}.
	AccessPolicies *string `json:"accessPolicies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#advanced_options ElasticsearchDomain#advanced_options}.
	AdvancedOptions interface{} `json:"advancedOptions"`
	// advanced_security_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#advanced_security_options ElasticsearchDomain#advanced_security_options}
	AdvancedSecurityOptions *ElasticsearchDomainAdvancedSecurityOptions `json:"advancedSecurityOptions"`
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#cluster_config ElasticsearchDomain#cluster_config}
	ClusterConfig *ElasticsearchDomainClusterConfig `json:"clusterConfig"`
	// cognito_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#cognito_options ElasticsearchDomain#cognito_options}
	CognitoOptions *ElasticsearchDomainCognitoOptions `json:"cognitoOptions"`
	// domain_endpoint_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#domain_endpoint_options ElasticsearchDomain#domain_endpoint_options}
	DomainEndpointOptions *ElasticsearchDomainDomainEndpointOptions `json:"domainEndpointOptions"`
	// ebs_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#ebs_options ElasticsearchDomain#ebs_options}
	EbsOptions *ElasticsearchDomainEbsOptions `json:"ebsOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#elasticsearch_version ElasticsearchDomain#elasticsearch_version}.
	ElasticsearchVersion *string `json:"elasticsearchVersion"`
	// encrypt_at_rest block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#encrypt_at_rest ElasticsearchDomain#encrypt_at_rest}
	EncryptAtRest *ElasticsearchDomainEncryptAtRest `json:"encryptAtRest"`
	// log_publishing_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#log_publishing_options ElasticsearchDomain#log_publishing_options}
	LogPublishingOptions *[]*ElasticsearchDomainLogPublishingOptions `json:"logPublishingOptions"`
	// node_to_node_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#node_to_node_encryption ElasticsearchDomain#node_to_node_encryption}
	NodeToNodeEncryption *ElasticsearchDomainNodeToNodeEncryption `json:"nodeToNodeEncryption"`
	// snapshot_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#snapshot_options ElasticsearchDomain#snapshot_options}
	SnapshotOptions *ElasticsearchDomainSnapshotOptions `json:"snapshotOptions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#tags ElasticsearchDomain#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#tags_all ElasticsearchDomain#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#timeouts ElasticsearchDomain#timeouts}
	Timeouts *ElasticsearchDomainTimeouts `json:"timeouts"`
	// vpc_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#vpc_options ElasticsearchDomain#vpc_options}
	VpcOptions *ElasticsearchDomainVpcOptions `json:"vpcOptions"`
}

type ElasticsearchDomainDomainEndpointOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#custom_endpoint ElasticsearchDomain#custom_endpoint}.
	CustomEndpoint *string `json:"customEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#custom_endpoint_certificate_arn ElasticsearchDomain#custom_endpoint_certificate_arn}.
	CustomEndpointCertificateArn *string `json:"customEndpointCertificateArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#custom_endpoint_enabled ElasticsearchDomain#custom_endpoint_enabled}.
	CustomEndpointEnabled interface{} `json:"customEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#enforce_https ElasticsearchDomain#enforce_https}.
	EnforceHttps interface{} `json:"enforceHttps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#tls_security_policy ElasticsearchDomain#tls_security_policy}.
	TlsSecurityPolicy *string `json:"tlsSecurityPolicy"`
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
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	TlsSecurityPolicy() *string
	SetTlsSecurityPolicy(val *string)
	TlsSecurityPolicyInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func NewElasticsearchDomainDomainEndpointOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainDomainEndpointOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainDomainEndpointOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainDomainEndpointOptionsOutputReference_Override(e ElasticsearchDomainDomainEndpointOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainDomainEndpointOptionsOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (e *jsiiProxy_ElasticsearchDomainDomainEndpointOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#ebs_enabled ElasticsearchDomain#ebs_enabled}.
	EbsEnabled interface{} `json:"ebsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#iops ElasticsearchDomain#iops}.
	Iops *float64 `json:"iops"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#volume_size ElasticsearchDomain#volume_size}.
	VolumeSize *float64 `json:"volumeSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#volume_type ElasticsearchDomain#volume_type}.
	VolumeType *string `json:"volumeType"`
}

type ElasticsearchDomainEbsOptionsOutputReference interface {
	cdktf.ComplexObject
	EbsEnabled() interface{}
	SetEbsEnabled(val interface{})
	EbsEnabledInput() interface{}
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	VolumeSize() *float64
	SetVolumeSize(val *float64)
	VolumeSizeInput() *float64
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func NewElasticsearchDomainEbsOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainEbsOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainEbsOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainEbsOptionsOutputReference_Override(e ElasticsearchDomainEbsOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainEbsOptionsOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (e *jsiiProxy_ElasticsearchDomainEbsOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#kms_key_id ElasticsearchDomain#kms_key_id}.
	KmsKeyId *string `json:"kmsKeyId"`
}

type ElasticsearchDomainEncryptAtRestOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
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

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainEncryptAtRestOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainEncryptAtRestOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainEncryptAtRestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainEncryptAtRestOutputReference_Override(e ElasticsearchDomainEncryptAtRestOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainEncryptAtRestOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainEncryptAtRestOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#cloudwatch_log_group_arn ElasticsearchDomain#cloudwatch_log_group_arn}.
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#log_type ElasticsearchDomain#log_type}.
	LogType *string `json:"logType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled"`
}

type ElasticsearchDomainNodeToNodeEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `json:"enabled"`
}

type ElasticsearchDomainNodeToNodeEncryptionOutputReference interface {
	cdktf.ComplexObject
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainNodeToNodeEncryptionOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainNodeToNodeEncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainNodeToNodeEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainNodeToNodeEncryptionOutputReference_Override(e ElasticsearchDomainNodeToNodeEncryptionOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainNodeToNodeEncryptionOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainNodeToNodeEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy.html aws_elasticsearch_domain_policy}.
type ElasticsearchDomainPolicy interface {
	cdktf.TerraformResource
	AccessPolicies() *string
	SetAccessPolicies(val *string)
	AccessPoliciesInput() *string
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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

func (j *jsiiProxy_ElasticsearchDomainPolicy) Count() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy.html aws_elasticsearch_domain_policy} Resource.
func NewElasticsearchDomainPolicy(scope constructs.Construct, id *string, config *ElasticsearchDomainPolicyConfig) ElasticsearchDomainPolicy {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainPolicy{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy.html aws_elasticsearch_domain_policy} Resource.
func NewElasticsearchDomainPolicy_Override(e ElasticsearchDomainPolicy, scope constructs.Construct, id *string, config *ElasticsearchDomainPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainPolicy",
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

func (j *jsiiProxy_ElasticsearchDomainPolicy) SetCount(val interface{}) {
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
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainPolicy",
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
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainPolicy",
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

type ElasticsearchDomainPolicyConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy.html#access_policies ElasticsearchDomainPolicy#access_policies}.
	AccessPolicies *string `json:"accessPolicies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy.html#domain_name ElasticsearchDomainPolicy#domain_name}.
	DomainName *string `json:"domainName"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html aws_elasticsearch_domain_saml_options}.
type ElasticsearchDomainSamlOptions interface {
	cdktf.TerraformResource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
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
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) Count() interface{} {
	var returns interface{}
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

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html aws_elasticsearch_domain_saml_options} Resource.
func NewElasticsearchDomainSamlOptions(scope constructs.Construct, id *string, config *ElasticsearchDomainSamlOptionsConfig) ElasticsearchDomainSamlOptions {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainSamlOptions{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSamlOptions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html aws_elasticsearch_domain_saml_options} Resource.
func NewElasticsearchDomainSamlOptions_Override(e ElasticsearchDomainSamlOptions, scope constructs.Construct, id *string, config *ElasticsearchDomainSamlOptionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSamlOptions",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticsearchDomainSamlOptions) SetCount(val interface{}) {
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
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSamlOptions",
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
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSamlOptions",
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

type ElasticsearchDomainSamlOptionsConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#domain_name ElasticsearchDomainSamlOptions#domain_name}.
	DomainName *string `json:"domainName"`
	// saml_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#saml_options ElasticsearchDomainSamlOptions#saml_options}
	SamlOptions *ElasticsearchDomainSamlOptionsSamlOptions `json:"samlOptions"`
}

type ElasticsearchDomainSamlOptionsSamlOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#enabled ElasticsearchDomainSamlOptions#enabled}.
	Enabled interface{} `json:"enabled"`
	// idp block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#idp ElasticsearchDomainSamlOptions#idp}
	Idp *ElasticsearchDomainSamlOptionsSamlOptionsIdp `json:"idp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#master_backend_role ElasticsearchDomainSamlOptions#master_backend_role}.
	MasterBackendRole *string `json:"masterBackendRole"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#master_user_name ElasticsearchDomainSamlOptions#master_user_name}.
	MasterUserName *string `json:"masterUserName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#roles_key ElasticsearchDomainSamlOptions#roles_key}.
	RolesKey *string `json:"rolesKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#session_timeout_minutes ElasticsearchDomainSamlOptions#session_timeout_minutes}.
	SessionTimeoutMinutes *float64 `json:"sessionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#subject_key ElasticsearchDomainSamlOptions#subject_key}.
	SubjectKey *string `json:"subjectKey"`
}

type ElasticsearchDomainSamlOptionsSamlOptionsIdp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#entity_id ElasticsearchDomainSamlOptions#entity_id}.
	EntityId *string `json:"entityId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_saml_options.html#metadata_content ElasticsearchDomainSamlOptions#metadata_content}.
	MetadataContent *string `json:"metadataContent"`
}

type ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference interface {
	cdktf.ComplexObject
	EntityId() *string
	SetEntityId(val *string)
	EntityIdInput() *string
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	MetadataContent() *string
	SetMetadataContent(val *string)
	MetadataContentInput() *string
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

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference_Override(e ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
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

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainSamlOptionsSamlOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainSamlOptionsSamlOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSamlOptionsSamlOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainSamlOptionsSamlOptionsOutputReference_Override(e ElasticsearchDomainSamlOptionsSamlOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSamlOptionsSamlOptionsOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSamlOptionsSamlOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#automated_snapshot_start_hour ElasticsearchDomain#automated_snapshot_start_hour}.
	AutomatedSnapshotStartHour *float64 `json:"automatedSnapshotStartHour"`
}

type ElasticsearchDomainSnapshotOptionsOutputReference interface {
	cdktf.ComplexObject
	AutomatedSnapshotStartHour() *float64
	SetAutomatedSnapshotStartHour(val *float64)
	AutomatedSnapshotStartHourInput() *float64
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

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainSnapshotOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainSnapshotOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSnapshotOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainSnapshotOptionsOutputReference_Override(e ElasticsearchDomainSnapshotOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainSnapshotOptionsOutputReference",
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

func (j *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainSnapshotOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#update ElasticsearchDomain#update}.
	Update *string `json:"update"`
}

type ElasticsearchDomainTimeoutsOutputReference interface {
	cdktf.ComplexObject
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
	ResetUpdate()
}

// The jsii proxy struct for ElasticsearchDomainTimeoutsOutputReference
type jsiiProxy_ElasticsearchDomainTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
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

func NewElasticsearchDomainTimeoutsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainTimeoutsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainTimeoutsOutputReference_Override(e ElasticsearchDomainTimeoutsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
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

func (j *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
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
func (e *jsiiProxy_ElasticsearchDomainTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#security_group_ids ElasticsearchDomain#security_group_ids}.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html#subnet_ids ElasticsearchDomain#subnet_ids}.
	SubnetIds *[]*string `json:"subnetIds"`
}

type ElasticsearchDomainVpcOptionsOutputReference interface {
	cdktf.ComplexObject
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
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetSecurityGroupIds()
	ResetSubnetIds()
}

// The jsii proxy struct for ElasticsearchDomainVpcOptionsOutputReference
type jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
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

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewElasticsearchDomainVpcOptionsOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ElasticsearchDomainVpcOptionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainVpcOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewElasticsearchDomainVpcOptionsOutputReference_Override(e ElasticsearchDomainVpcOptionsOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ElasticSearch.ElasticsearchDomainVpcOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		e,
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

func (j *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (e *jsiiProxy_ElasticsearchDomainVpcOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
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
