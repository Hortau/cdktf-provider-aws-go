package resourcegroups

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hortau/cdktf-provider-aws-go/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/hortau/cdktf-provider-aws-go/resourcegroups/internal"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html aws_resourcegroupstaggingapi_resources}.
type DataAwsResourcegroupstaggingapiResources interface {
	cdktf.TerraformDataSource
	CdktfStack() cdktf.TerraformStack
	ConstructNodeMetadata() *map[string]interface{}
	Count() interface{}
	SetCount(val interface{})
	DependsOn() *[]*string
	SetDependsOn(val *[]*string)
	ExcludeCompliantResources() interface{}
	SetExcludeCompliantResources(val interface{})
	ExcludeCompliantResourcesInput() interface{}
	Fqn() *string
	FriendlyUniqueId() *string
	Id() *string
	IncludeComplianceDetails() interface{}
	SetIncludeComplianceDetails(val interface{})
	IncludeComplianceDetailsInput() interface{}
	Lifecycle() *cdktf.TerraformResourceLifecycle
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Node() constructs.Node
	Provider() cdktf.TerraformProvider
	SetProvider(val cdktf.TerraformProvider)
	RawOverrides() interface{}
	ResourceArnList() *[]*string
	SetResourceArnList(val *[]*string)
	ResourceArnListInput() *[]*string
	ResourceTypeFilters() *[]*string
	SetResourceTypeFilters(val *[]*string)
	ResourceTypeFiltersInput() *[]*string
	TagFilter() *[]*DataAwsResourcegroupstaggingapiResourcesTagFilter
	SetTagFilter(val *[]*DataAwsResourcegroupstaggingapiResourcesTagFilter)
	TagFilterInput() *[]*DataAwsResourcegroupstaggingapiResourcesTagFilter
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
	ResetExcludeCompliantResources()
	ResetIncludeComplianceDetails()
	ResetOverrideLogicalId()
	ResetResourceArnList()
	ResetResourceTypeFilters()
	ResetTagFilter()
	ResourceTagMappingList(index *string) DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsResourcegroupstaggingapiResources
type jsiiProxy_DataAwsResourcegroupstaggingapiResources struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ExcludeCompliantResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCompliantResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ExcludeCompliantResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCompliantResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) IncludeComplianceDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeComplianceDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) IncludeComplianceDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeComplianceDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceArnList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceArnListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceTypeFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceTypeFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypeFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TagFilter() *[]*DataAwsResourcegroupstaggingapiResourcesTagFilter {
	var returns *[]*DataAwsResourcegroupstaggingapiResourcesTagFilter
	_jsii_.Get(
		j,
		"tagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TagFilterInput() *[]*DataAwsResourcegroupstaggingapiResourcesTagFilter {
	var returns *[]*DataAwsResourcegroupstaggingapiResourcesTagFilter
	_jsii_.Get(
		j,
		"tagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html aws_resourcegroupstaggingapi_resources} Data Source.
func NewDataAwsResourcegroupstaggingapiResources(scope constructs.Construct, id *string, config *DataAwsResourcegroupstaggingapiResourcesConfig) DataAwsResourcegroupstaggingapiResources {
	_init_.Initialize()

	j := jsiiProxy_DataAwsResourcegroupstaggingapiResources{}

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.DataAwsResourcegroupstaggingapiResources",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html aws_resourcegroupstaggingapi_resources} Data Source.
func NewDataAwsResourcegroupstaggingapiResources_Override(d DataAwsResourcegroupstaggingapiResources, scope constructs.Construct, id *string, config *DataAwsResourcegroupstaggingapiResourcesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.DataAwsResourcegroupstaggingapiResources",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SetExcludeCompliantResources(val interface{}) {
	_jsii_.Set(
		j,
		"excludeCompliantResources",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SetIncludeComplianceDetails(val interface{}) {
	_jsii_.Set(
		j,
		"includeComplianceDetails",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SetResourceArnList(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceArnList",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SetResourceTypeFilters(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceTypeFilters",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SetTagFilter(val *[]*DataAwsResourcegroupstaggingapiResourcesTagFilter) {
	_jsii_.Set(
		j,
		"tagFilter",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DataAwsResourcegroupstaggingapiResources_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ResourceGroups.DataAwsResourcegroupstaggingapiResources",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsResourcegroupstaggingapiResources_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ResourceGroups.DataAwsResourcegroupstaggingapiResources",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetExcludeCompliantResources() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeCompliantResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetIncludeComplianceDetails() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeComplianceDetails",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetResourceArnList() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceArnList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetResourceTypeFilters() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceTypeFilters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResetTagFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetTagFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ResourceTagMappingList(index *string) DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList {
	var returns DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList

	_jsii_.Invoke(
		d,
		"resourceTagMappingList",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) SynthesizeAttributes() *map[string]interface{} {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ToMetadata() interface{} {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ToString() *string {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResources) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataAwsResourcegroupstaggingapiResourcesConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html#exclude_compliant_resources DataAwsResourcegroupstaggingapiResources#exclude_compliant_resources}.
	ExcludeCompliantResources interface{} `json:"excludeCompliantResources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html#include_compliance_details DataAwsResourcegroupstaggingapiResources#include_compliance_details}.
	IncludeComplianceDetails interface{} `json:"includeComplianceDetails"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html#resource_arn_list DataAwsResourcegroupstaggingapiResources#resource_arn_list}.
	ResourceArnList *[]*string `json:"resourceArnList"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html#resource_type_filters DataAwsResourcegroupstaggingapiResources#resource_type_filters}.
	ResourceTypeFilters *[]*string `json:"resourceTypeFilters"`
	// tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html#tag_filter DataAwsResourcegroupstaggingapiResources#tag_filter}
	TagFilter *[]*DataAwsResourcegroupstaggingapiResourcesTagFilter `json:"tagFilter"`
}

type DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ComplianceDetails() interface{}
	ResourceArn() *string
	Tags() interface{}
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

// The jsii proxy struct for DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList
type jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) ComplianceDetails() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsResourcegroupstaggingapiResourcesResourceTagMappingList(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList {
	_init_.Initialize()

	j := jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList{}

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsResourcegroupstaggingapiResourcesResourceTagMappingList_Override(d DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingList) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails interface {
	cdktf.ComplexComputedList
	ComplexComputedListIndex() *string
	SetComplexComputedListIndex(val *string)
	ComplianceStatus() interface{}
	KeysWithNoncompliantValues() *[]*string
	NonCompliantKeys() *[]*string
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

// The jsii proxy struct for DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails
type jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails struct {
	internal.Type__cdktfComplexComputedList
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) ComplexComputedListIndex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complexComputedListIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) ComplianceStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) KeysWithNoncompliantValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"keysWithNoncompliantValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) NonCompliantKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nonCompliantKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

// Experimental.
func NewDataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails(terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails {
	_init_.Initialize()

	j := jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails{}

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails_Override(d DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails, terraformResource cdktf.ITerraformResource, terraformAttribute *string, complexComputedListIndex *string) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails",
		[]interface{}{terraformResource, terraformAttribute, complexComputedListIndex},
		d,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) SetComplexComputedListIndex(val *string) {
	_jsii_.Set(
		j,
		"complexComputedListIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

// Experimental.
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) GetBooleanAttribute(terraformAttribute *string) interface{} {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) GetListAttribute(terraformAttribute *string) *[]*string {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) GetNumberAttribute(terraformAttribute *string) *float64 {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) GetStringAttribute(terraformAttribute *string) *string {
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
func (d *jsiiProxy_DataAwsResourcegroupstaggingapiResourcesResourceTagMappingListComplianceDetails) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

type DataAwsResourcegroupstaggingapiResourcesTagFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html#key DataAwsResourcegroupstaggingapiResources#key}.
	Key *string `json:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/resourcegroupstaggingapi_resources.html#values DataAwsResourcegroupstaggingapiResources#values}.
	Values *[]*string `json:"values"`
}

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html aws_resourcegroups_group}.
type ResourcegroupsGroup interface {
	cdktf.TerraformResource
	Arn() *string
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
	ResourceQuery() ResourcegroupsGroupResourceQueryOutputReference
	ResourceQueryInput() *ResourcegroupsGroupResourceQuery
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
	PutResourceQuery(value *ResourcegroupsGroupResourceQuery)
	ResetDescription()
	ResetOverrideLogicalId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	ToMetadata() interface{}
	ToString() *string
	ToTerraform() interface{}
}

// The jsii proxy struct for ResourcegroupsGroup
type jsiiProxy_ResourcegroupsGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ResourcegroupsGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) ResourceQuery() ResourcegroupsGroupResourceQueryOutputReference {
	var returns ResourcegroupsGroupResourceQueryOutputReference
	_jsii_.Get(
		j,
		"resourceQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) ResourceQueryInput() *ResourcegroupsGroupResourceQuery {
	var returns *ResourcegroupsGroupResourceQuery
	_jsii_.Get(
		j,
		"resourceQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) Tags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) TagsAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) TagsAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html aws_resourcegroups_group} Resource.
func NewResourcegroupsGroup(scope constructs.Construct, id *string, config *ResourcegroupsGroupConfig) ResourcegroupsGroup {
	_init_.Initialize()

	j := jsiiProxy_ResourcegroupsGroup{}

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.ResourcegroupsGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html aws_resourcegroups_group} Resource.
func NewResourcegroupsGroup_Override(r ResourcegroupsGroup, scope constructs.Construct, id *string, config *ResourcegroupsGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.ResourcegroupsGroup",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ResourcegroupsGroup) SetCount(val interface{}) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroup) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroup) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroup) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroup) SetTags(val interface{}) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroup) SetTagsAll(val interface{}) {
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
func ResourcegroupsGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"hashicorp_aws.ResourceGroups.ResourcegroupsGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ResourcegroupsGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"hashicorp_aws.ResourceGroups.ResourcegroupsGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ResourcegroupsGroup) PutResourceQuery(value *ResourcegroupsGroupResourceQuery) {
	_jsii_.InvokeVoid(
		r,
		"putResourceQuery",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourcegroupsGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		r,
		"resetDescription",
		nil, // no parameters
	)
}

// Resets a previously passed logical Id to use the auto-generated logical id again.
// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcegroupsGroup) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcegroupsGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		r,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourcegroupsGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_ResourcegroupsGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds this resource to the terraform JSON output.
// Experimental.
func (r *jsiiProxy_ResourcegroupsGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ResourcegroupsGroupConfig struct {
	// Experimental.
	Count interface{} `json:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `json:"dependsOn"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `json:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `json:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html#name ResourcegroupsGroup#name}.
	Name *string `json:"name"`
	// resource_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html#resource_query ResourcegroupsGroup#resource_query}
	ResourceQuery *ResourcegroupsGroupResourceQuery `json:"resourceQuery"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html#description ResourcegroupsGroup#description}.
	Description *string `json:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html#tags ResourcegroupsGroup#tags}.
	Tags interface{} `json:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html#tags_all ResourcegroupsGroup#tags_all}.
	TagsAll interface{} `json:"tagsAll"`
}

type ResourcegroupsGroupResourceQuery struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html#query ResourcegroupsGroup#query}.
	Query *string `json:"query"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/resourcegroups_group.html#type ResourcegroupsGroup#type}.
	Type *string `json:"type"`
}

type ResourcegroupsGroupResourceQueryOutputReference interface {
	cdktf.ComplexObject
	IsSingleItem() *bool
	SetIsSingleItem(val *bool)
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	TerraformResource() cdktf.ITerraformResource
	SetTerraformResource(val cdktf.ITerraformResource)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	GetBooleanAttribute(terraformAttribute *string) interface{}
	GetListAttribute(terraformAttribute *string) *[]*string
	GetNumberAttribute(terraformAttribute *string) *float64
	GetStringAttribute(terraformAttribute *string) *string
	InterpolationAsList() cdktf.IResolvable
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetType()
}

// The jsii proxy struct for ResourcegroupsGroupResourceQueryOutputReference
type jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) IsSingleItem() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSingleItem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) TerraformResource() cdktf.ITerraformResource {
	var returns cdktf.ITerraformResource
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func NewResourcegroupsGroupResourceQueryOutputReference(terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) ResourcegroupsGroupResourceQueryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference{}

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.ResourcegroupsGroupResourceQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		&j,
	)

	return &j
}

func NewResourcegroupsGroupResourceQueryOutputReference_Override(r ResourcegroupsGroupResourceQueryOutputReference, terraformResource cdktf.ITerraformResource, terraformAttribute *string, isSingleItem *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"hashicorp_aws.ResourceGroups.ResourcegroupsGroupResourceQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute, isSingleItem},
		r,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) SetIsSingleItem(val *bool) {
	_jsii_.Set(
		j,
		"isSingleItem",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) SetQuery(val *string) {
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) SetTerraformResource(val cdktf.ITerraformResource) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (r *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourcegroupsGroupResourceQueryOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		r,
		"resetType",
		nil, // no parameters
	)
}
