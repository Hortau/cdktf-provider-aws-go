package appmesh

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRoute",
		reflect.TypeOf((*AppmeshGatewayRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDate", GoGetter: "LastUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshNameInput", GoGetter: "MeshNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwnerInput", GoGetter: "MeshOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putSpec", GoMethod: "PutSpec"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetMeshOwner", GoMethod: "ResetMeshOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "resourceOwner", GoGetter: "ResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "specInput", GoGetter: "SpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGatewayName", GoGetter: "VirtualGatewayName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualGatewayNameInput", GoGetter: "VirtualGatewayNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRoute{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteConfig",
		reflect.TypeOf((*AppmeshGatewayRouteConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpec",
		reflect.TypeOf((*AppmeshGatewayRouteSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRoute",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRouteAction",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRouteAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRouteActionOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRouteActionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putTarget", GoMethod: "PutTarget"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberProperty{JsiiProperty: "targetInput", GoGetter: "TargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecGrpcRouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRouteActionTarget",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRouteActionTarget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRouteActionTargetOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRouteActionTargetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putVirtualService", GoMethod: "PutVirtualService"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualService", GoGetter: "VirtualService"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceInput", GoGetter: "VirtualServiceInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecGrpcRouteActionTargetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRouteActionTargetVirtualService",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRouteActionTargetVirtualService)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRouteActionTargetVirtualServiceOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRouteActionTargetVirtualServiceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceName", GoGetter: "VirtualServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceNameInput", GoGetter: "VirtualServiceNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecGrpcRouteActionTargetVirtualServiceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRouteMatch",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRouteMatchOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRouteMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNameInput", GoGetter: "ServiceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecGrpcRouteMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecGrpcRouteOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecGrpcRouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAction", GoMethod: "PutAction"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecGrpcRouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2Route",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2Route)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2RouteAction",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2RouteAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2RouteActionOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2RouteActionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putTarget", GoMethod: "PutTarget"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberProperty{JsiiProperty: "targetInput", GoGetter: "TargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2RouteActionTarget",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2RouteActionTarget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2RouteActionTargetOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2RouteActionTargetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putVirtualService", GoMethod: "PutVirtualService"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualService", GoGetter: "VirtualService"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceInput", GoGetter: "VirtualServiceInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteActionTargetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2RouteActionTargetVirtualService",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2RouteActionTargetVirtualService)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2RouteActionTargetVirtualServiceOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2RouteActionTargetVirtualServiceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceName", GoGetter: "VirtualServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceNameInput", GoGetter: "VirtualServiceNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteActionTargetVirtualServiceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2RouteMatch",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2RouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttp2RouteOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttp2RouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAction", GoMethod: "PutAction"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttp2RouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRoute",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRouteAction",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRouteAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRouteActionOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRouteActionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putTarget", GoMethod: "PutTarget"},
			_jsii_.MemberProperty{JsiiProperty: "target", GoGetter: "Target"},
			_jsii_.MemberProperty{JsiiProperty: "targetInput", GoGetter: "TargetInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttpRouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRouteActionTarget",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRouteActionTarget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRouteActionTargetOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRouteActionTargetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putVirtualService", GoMethod: "PutVirtualService"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualService", GoGetter: "VirtualService"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceInput", GoGetter: "VirtualServiceInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttpRouteActionTargetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRouteActionTargetVirtualService",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRouteActionTargetVirtualService)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRouteActionTargetVirtualServiceOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRouteActionTargetVirtualServiceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceName", GoGetter: "VirtualServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceNameInput", GoGetter: "VirtualServiceNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttpRouteActionTargetVirtualServiceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRouteMatch",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRouteMatchOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRouteMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttpRouteMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecHttpRouteOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecHttpRouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAction", GoMethod: "PutAction"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecHttpRouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshGatewayRouteSpecOutputReference",
		reflect.TypeOf((*AppmeshGatewayRouteSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grpcRoute", GoGetter: "GrpcRoute"},
			_jsii_.MemberProperty{JsiiProperty: "grpcRouteInput", GoGetter: "GrpcRouteInput"},
			_jsii_.MemberProperty{JsiiProperty: "http2Route", GoGetter: "Http2Route"},
			_jsii_.MemberProperty{JsiiProperty: "http2RouteInput", GoGetter: "Http2RouteInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpRoute", GoGetter: "HttpRoute"},
			_jsii_.MemberProperty{JsiiProperty: "httpRouteInput", GoGetter: "HttpRouteInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putGrpcRoute", GoMethod: "PutGrpcRoute"},
			_jsii_.MemberMethod{JsiiMethod: "putHttp2Route", GoMethod: "PutHttp2Route"},
			_jsii_.MemberMethod{JsiiMethod: "putHttpRoute", GoMethod: "PutHttpRoute"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrpcRoute", GoMethod: "ResetGrpcRoute"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp2Route", GoMethod: "ResetHttp2Route"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpRoute", GoMethod: "ResetHttpRoute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshGatewayRouteSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshMesh",
		reflect.TypeOf((*AppmeshMesh)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDate", GoGetter: "LastUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putSpec", GoMethod: "PutSpec"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpec", GoMethod: "ResetSpec"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "resourceOwner", GoGetter: "ResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "specInput", GoGetter: "SpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshMesh{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshMeshConfig",
		reflect.TypeOf((*AppmeshMeshConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshMeshSpec",
		reflect.TypeOf((*AppmeshMeshSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshMeshSpecEgressFilter",
		reflect.TypeOf((*AppmeshMeshSpecEgressFilter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshMeshSpecEgressFilterOutputReference",
		reflect.TypeOf((*AppmeshMeshSpecEgressFilterOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshMeshSpecEgressFilterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshMeshSpecOutputReference",
		reflect.TypeOf((*AppmeshMeshSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "egressFilter", GoGetter: "EgressFilter"},
			_jsii_.MemberProperty{JsiiProperty: "egressFilterInput", GoGetter: "EgressFilterInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putEgressFilter", GoMethod: "PutEgressFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetEgressFilter", GoMethod: "ResetEgressFilter"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshMeshSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRoute",
		reflect.TypeOf((*AppmeshRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDate", GoGetter: "LastUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshNameInput", GoGetter: "MeshNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwnerInput", GoGetter: "MeshOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putSpec", GoMethod: "PutSpec"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetMeshOwner", GoMethod: "ResetMeshOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "resourceOwner", GoGetter: "ResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "specInput", GoGetter: "SpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterName", GoGetter: "VirtualRouterName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterNameInput", GoGetter: "VirtualRouterNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRoute{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteConfig",
		reflect.TypeOf((*AppmeshRouteConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpec",
		reflect.TypeOf((*AppmeshRouteSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRoute",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteAction",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteActionOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteActionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "weightedTarget", GoGetter: "WeightedTarget"},
			_jsii_.MemberProperty{JsiiProperty: "weightedTargetInput", GoGetter: "WeightedTargetInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteActionWeightedTarget",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteActionWeightedTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteMatch",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteMatchMetadata",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteMatchMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteMatchMetadataMatch",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteMatchMetadataMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exact", GoGetter: "Exact"},
			_jsii_.MemberProperty{JsiiProperty: "exactInput", GoGetter: "ExactInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putRange", GoMethod: "PutRange"},
			_jsii_.MemberProperty{JsiiProperty: "range", GoGetter: "Range"},
			_jsii_.MemberProperty{JsiiProperty: "rangeInput", GoGetter: "RangeInput"},
			_jsii_.MemberProperty{JsiiProperty: "regex", GoGetter: "Regex"},
			_jsii_.MemberProperty{JsiiProperty: "regexInput", GoGetter: "RegexInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExact", GoMethod: "ResetExact"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetRange", GoMethod: "ResetRange"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegex", GoMethod: "ResetRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuffix", GoMethod: "ResetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "suffix", GoGetter: "Suffix"},
			_jsii_.MemberProperty{JsiiProperty: "suffixInput", GoGetter: "SuffixInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteMatchMetadataMatchRange",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteMatchMetadataMatchRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteMatchMetadataMatchRangeOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteMatchMetadataMatchRangeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "end", GoGetter: "End"},
			_jsii_.MemberProperty{JsiiProperty: "endInput", GoGetter: "EndInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "start", GoGetter: "Start"},
			_jsii_.MemberProperty{JsiiProperty: "startInput", GoGetter: "StartInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteMatchOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "methodName", GoGetter: "MethodName"},
			_jsii_.MemberProperty{JsiiProperty: "methodNameInput", GoGetter: "MethodNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetMethodName", GoMethod: "ResetMethodName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceName", GoMethod: "ResetServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNameInput", GoGetter: "ServiceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAction", GoMethod: "PutAction"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryPolicy", GoMethod: "PutRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeout", GoMethod: "PutTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetMatch", GoMethod: "ResetMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryPolicy", GoMethod: "ResetRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicy", GoGetter: "RetryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicyInput", GoGetter: "RetryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteRetryPolicy",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteRetryPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grpcRetryEvents", GoGetter: "GrpcRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "grpcRetryEventsInput", GoGetter: "GrpcRetryEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryEvents", GoGetter: "HttpRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryEventsInput", GoGetter: "HttpRetryEventsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "perRetryTimeout", GoGetter: "PerRetryTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "perRetryTimeoutInput", GoGetter: "PerRetryTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPerRetryTimeout", GoMethod: "PutPerRetryTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrpcRetryEvents", GoMethod: "ResetGrpcRetryEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpRetryEvents", GoMethod: "ResetHttpRetryEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetTcpRetryEvents", GoMethod: "ResetTcpRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "tcpRetryEvents", GoGetter: "TcpRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "tcpRetryEventsInput", GoGetter: "TcpRetryEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeout",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeoutOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeoutOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteRetryPolicyPerRetryTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteTimeout",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteTimeoutIdle",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteTimeoutIdle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteTimeoutIdleOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteTimeoutIdleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteTimeoutIdleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteTimeoutOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteTimeoutOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idle", GoGetter: "Idle"},
			_jsii_.MemberProperty{JsiiProperty: "idleInput", GoGetter: "IdleInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "perRequest", GoGetter: "PerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "perRequestInput", GoGetter: "PerRequestInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIdle", GoMethod: "PutIdle"},
			_jsii_.MemberMethod{JsiiMethod: "putPerRequest", GoMethod: "PutPerRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdle", GoMethod: "ResetIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerRequest", GoMethod: "ResetPerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteTimeoutPerRequest",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteTimeoutPerRequest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecGrpcRouteTimeoutPerRequestOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecGrpcRouteTimeoutPerRequestOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecGrpcRouteTimeoutPerRequestOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2Route",
		reflect.TypeOf((*AppmeshRouteSpecHttp2Route)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteAction",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteActionOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteActionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "weightedTarget", GoGetter: "WeightedTarget"},
			_jsii_.MemberProperty{JsiiProperty: "weightedTargetInput", GoGetter: "WeightedTargetInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteActionWeightedTarget",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteActionWeightedTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteMatch",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteMatchHeader",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteMatchHeader)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteMatchHeaderMatch",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteMatchHeaderMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteMatchHeaderMatchOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteMatchHeaderMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exact", GoGetter: "Exact"},
			_jsii_.MemberProperty{JsiiProperty: "exactInput", GoGetter: "ExactInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putRange", GoMethod: "PutRange"},
			_jsii_.MemberProperty{JsiiProperty: "range", GoGetter: "Range"},
			_jsii_.MemberProperty{JsiiProperty: "rangeInput", GoGetter: "RangeInput"},
			_jsii_.MemberProperty{JsiiProperty: "regex", GoGetter: "Regex"},
			_jsii_.MemberProperty{JsiiProperty: "regexInput", GoGetter: "RegexInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExact", GoMethod: "ResetExact"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetRange", GoMethod: "ResetRange"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegex", GoMethod: "ResetRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuffix", GoMethod: "ResetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "suffix", GoGetter: "Suffix"},
			_jsii_.MemberProperty{JsiiProperty: "suffixInput", GoGetter: "SuffixInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteMatchHeaderMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteMatchHeaderMatchRange",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteMatchHeaderMatchRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteMatchHeaderMatchRangeOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteMatchHeaderMatchRangeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "end", GoGetter: "End"},
			_jsii_.MemberProperty{JsiiProperty: "endInput", GoGetter: "EndInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "start", GoGetter: "Start"},
			_jsii_.MemberProperty{JsiiProperty: "startInput", GoGetter: "StartInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteMatchHeaderMatchRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteMatchOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "header", GoGetter: "Header"},
			_jsii_.MemberProperty{JsiiProperty: "headerInput", GoGetter: "HeaderInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "methodInput", GoGetter: "MethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeader", GoMethod: "ResetHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetMethod", GoMethod: "ResetMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheme", GoMethod: "ResetScheme"},
			_jsii_.MemberProperty{JsiiProperty: "scheme", GoGetter: "Scheme"},
			_jsii_.MemberProperty{JsiiProperty: "schemeInput", GoGetter: "SchemeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAction", GoMethod: "PutAction"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryPolicy", GoMethod: "PutRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeout", GoMethod: "PutTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryPolicy", GoMethod: "ResetRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicy", GoGetter: "RetryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicyInput", GoGetter: "RetryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteRetryPolicy",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteRetryPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteRetryPolicyOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteRetryPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryEvents", GoGetter: "HttpRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryEventsInput", GoGetter: "HttpRetryEventsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "perRetryTimeout", GoGetter: "PerRetryTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "perRetryTimeoutInput", GoGetter: "PerRetryTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPerRetryTimeout", GoMethod: "PutPerRetryTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpRetryEvents", GoMethod: "ResetHttpRetryEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetTcpRetryEvents", GoMethod: "ResetTcpRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "tcpRetryEvents", GoGetter: "TcpRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "tcpRetryEventsInput", GoGetter: "TcpRetryEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteRetryPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteRetryPolicyPerRetryTimeout",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteRetryPolicyPerRetryTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteRetryPolicyPerRetryTimeoutOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteRetryPolicyPerRetryTimeoutOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteRetryPolicyPerRetryTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteTimeout",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteTimeoutIdle",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteTimeoutIdle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteTimeoutIdleOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteTimeoutIdleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteTimeoutIdleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteTimeoutOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteTimeoutOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idle", GoGetter: "Idle"},
			_jsii_.MemberProperty{JsiiProperty: "idleInput", GoGetter: "IdleInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "perRequest", GoGetter: "PerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "perRequestInput", GoGetter: "PerRequestInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIdle", GoMethod: "PutIdle"},
			_jsii_.MemberMethod{JsiiMethod: "putPerRequest", GoMethod: "PutPerRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdle", GoMethod: "ResetIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerRequest", GoMethod: "ResetPerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteTimeoutPerRequest",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteTimeoutPerRequest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttp2RouteTimeoutPerRequestOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttp2RouteTimeoutPerRequestOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttp2RouteTimeoutPerRequestOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRoute",
		reflect.TypeOf((*AppmeshRouteSpecHttpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteAction",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteActionOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteActionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "weightedTarget", GoGetter: "WeightedTarget"},
			_jsii_.MemberProperty{JsiiProperty: "weightedTargetInput", GoGetter: "WeightedTargetInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteActionWeightedTarget",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteActionWeightedTarget)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteMatch",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteMatch)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteMatchHeader",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteMatchHeader)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteMatchHeaderMatch",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteMatchHeaderMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteMatchHeaderMatchOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteMatchHeaderMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exact", GoGetter: "Exact"},
			_jsii_.MemberProperty{JsiiProperty: "exactInput", GoGetter: "ExactInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "putRange", GoMethod: "PutRange"},
			_jsii_.MemberProperty{JsiiProperty: "range", GoGetter: "Range"},
			_jsii_.MemberProperty{JsiiProperty: "rangeInput", GoGetter: "RangeInput"},
			_jsii_.MemberProperty{JsiiProperty: "regex", GoGetter: "Regex"},
			_jsii_.MemberProperty{JsiiProperty: "regexInput", GoGetter: "RegexInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExact", GoMethod: "ResetExact"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetRange", GoMethod: "ResetRange"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegex", GoMethod: "ResetRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuffix", GoMethod: "ResetSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "suffix", GoGetter: "Suffix"},
			_jsii_.MemberProperty{JsiiProperty: "suffixInput", GoGetter: "SuffixInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteMatchHeaderMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteMatchHeaderMatchRange",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteMatchHeaderMatchRange)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteMatchHeaderMatchRangeOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteMatchHeaderMatchRangeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "end", GoGetter: "End"},
			_jsii_.MemberProperty{JsiiProperty: "endInput", GoGetter: "EndInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "start", GoGetter: "Start"},
			_jsii_.MemberProperty{JsiiProperty: "startInput", GoGetter: "StartInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteMatchHeaderMatchRangeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteMatchOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "header", GoGetter: "Header"},
			_jsii_.MemberProperty{JsiiProperty: "headerInput", GoGetter: "HeaderInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "methodInput", GoGetter: "MethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeader", GoMethod: "ResetHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetMethod", GoMethod: "ResetMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheme", GoMethod: "ResetScheme"},
			_jsii_.MemberProperty{JsiiProperty: "scheme", GoGetter: "Scheme"},
			_jsii_.MemberProperty{JsiiProperty: "schemeInput", GoGetter: "SchemeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAction", GoMethod: "PutAction"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberMethod{JsiiMethod: "putRetryPolicy", GoMethod: "PutRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeout", GoMethod: "PutTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetryPolicy", GoMethod: "ResetRetryPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicy", GoGetter: "RetryPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "retryPolicyInput", GoGetter: "RetryPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteRetryPolicy",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteRetryPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteRetryPolicyOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteRetryPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryEvents", GoGetter: "HttpRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "httpRetryEventsInput", GoGetter: "HttpRetryEventsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "perRetryTimeout", GoGetter: "PerRetryTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "perRetryTimeoutInput", GoGetter: "PerRetryTimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPerRetryTimeout", GoMethod: "PutPerRetryTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpRetryEvents", GoMethod: "ResetHttpRetryEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetTcpRetryEvents", GoMethod: "ResetTcpRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "tcpRetryEvents", GoGetter: "TcpRetryEvents"},
			_jsii_.MemberProperty{JsiiProperty: "tcpRetryEventsInput", GoGetter: "TcpRetryEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeout",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeoutOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeoutOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteRetryPolicyPerRetryTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteTimeout",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteTimeoutIdle",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteTimeoutIdle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteTimeoutIdleOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteTimeoutIdleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteTimeoutIdleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteTimeoutOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteTimeoutOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idle", GoGetter: "Idle"},
			_jsii_.MemberProperty{JsiiProperty: "idleInput", GoGetter: "IdleInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "perRequest", GoGetter: "PerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "perRequestInput", GoGetter: "PerRequestInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIdle", GoMethod: "PutIdle"},
			_jsii_.MemberMethod{JsiiMethod: "putPerRequest", GoMethod: "PutPerRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdle", GoMethod: "ResetIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerRequest", GoMethod: "ResetPerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteTimeoutPerRequest",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteTimeoutPerRequest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecHttpRouteTimeoutPerRequestOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecHttpRouteTimeoutPerRequestOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecHttpRouteTimeoutPerRequestOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grpcRoute", GoGetter: "GrpcRoute"},
			_jsii_.MemberProperty{JsiiProperty: "grpcRouteInput", GoGetter: "GrpcRouteInput"},
			_jsii_.MemberProperty{JsiiProperty: "http2Route", GoGetter: "Http2Route"},
			_jsii_.MemberProperty{JsiiProperty: "http2RouteInput", GoGetter: "Http2RouteInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpRoute", GoGetter: "HttpRoute"},
			_jsii_.MemberProperty{JsiiProperty: "httpRouteInput", GoGetter: "HttpRouteInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberMethod{JsiiMethod: "putGrpcRoute", GoMethod: "PutGrpcRoute"},
			_jsii_.MemberMethod{JsiiMethod: "putHttp2Route", GoMethod: "PutHttp2Route"},
			_jsii_.MemberMethod{JsiiMethod: "putHttpRoute", GoMethod: "PutHttpRoute"},
			_jsii_.MemberMethod{JsiiMethod: "putTcpRoute", GoMethod: "PutTcpRoute"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrpcRoute", GoMethod: "ResetGrpcRoute"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp2Route", GoMethod: "ResetHttp2Route"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpRoute", GoMethod: "ResetHttpRoute"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriority", GoMethod: "ResetPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resetTcpRoute", GoMethod: "ResetTcpRoute"},
			_jsii_.MemberProperty{JsiiProperty: "tcpRoute", GoGetter: "TcpRoute"},
			_jsii_.MemberProperty{JsiiProperty: "tcpRouteInput", GoGetter: "TcpRouteInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecTcpRoute",
		reflect.TypeOf((*AppmeshRouteSpecTcpRoute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecTcpRouteAction",
		reflect.TypeOf((*AppmeshRouteSpecTcpRouteAction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecTcpRouteActionOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecTcpRouteActionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "weightedTarget", GoGetter: "WeightedTarget"},
			_jsii_.MemberProperty{JsiiProperty: "weightedTargetInput", GoGetter: "WeightedTargetInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecTcpRouteActionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecTcpRouteActionWeightedTarget",
		reflect.TypeOf((*AppmeshRouteSpecTcpRouteActionWeightedTarget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecTcpRouteOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecTcpRouteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "actionInput", GoGetter: "ActionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAction", GoMethod: "PutAction"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeout", GoMethod: "PutTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecTcpRouteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecTcpRouteTimeout",
		reflect.TypeOf((*AppmeshRouteSpecTcpRouteTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecTcpRouteTimeoutIdle",
		reflect.TypeOf((*AppmeshRouteSpecTcpRouteTimeoutIdle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecTcpRouteTimeoutIdleOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecTcpRouteTimeoutIdleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecTcpRouteTimeoutIdleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshRouteSpecTcpRouteTimeoutOutputReference",
		reflect.TypeOf((*AppmeshRouteSpecTcpRouteTimeoutOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idle", GoGetter: "Idle"},
			_jsii_.MemberProperty{JsiiProperty: "idleInput", GoGetter: "IdleInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putIdle", GoMethod: "PutIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdle", GoMethod: "ResetIdle"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshRouteSpecTcpRouteTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGateway",
		reflect.TypeOf((*AppmeshVirtualGateway)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDate", GoGetter: "LastUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshNameInput", GoGetter: "MeshNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwnerInput", GoGetter: "MeshOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putSpec", GoMethod: "PutSpec"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetMeshOwner", GoMethod: "ResetMeshOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "resourceOwner", GoGetter: "ResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "specInput", GoGetter: "SpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGateway{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewayConfig",
		reflect.TypeOf((*AppmeshVirtualGatewayConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpec",
		reflect.TypeOf((*AppmeshVirtualGatewaySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaults",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaults)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicy",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putTls", GoMethod: "PutTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetTls", GoMethod: "ResetTls"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tls", GoGetter: "Tls"},
			_jsii_.MemberProperty{JsiiProperty: "tlsInput", GoGetter: "TlsInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTls",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificate",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateFile",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateSds",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsCertificateSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "enforce", GoGetter: "Enforce"},
			_jsii_.MemberProperty{JsiiProperty: "enforceInput", GoGetter: "EnforceInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "portsInput", GoGetter: "PortsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "putValidation", GoMethod: "PutValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnforce", GoMethod: "ResetEnforce"},
			_jsii_.MemberMethod{JsiiMethod: "resetPorts", GoMethod: "ResetPorts"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "validation", GoGetter: "Validation"},
			_jsii_.MemberProperty{JsiiProperty: "validationInput", GoGetter: "ValidationInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidation",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putSubjectAlternativeNames", GoMethod: "PutSubjectAlternativeNames"},
			_jsii_.MemberMethod{JsiiMethod: "putTrust", GoMethod: "PutTrust"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubjectAlternativeNames", GoMethod: "ResetSubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNames", GoGetter: "SubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNamesInput", GoGetter: "SubjectAlternativeNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "trust", GoGetter: "Trust"},
			_jsii_.MemberProperty{JsiiProperty: "trustInput", GoGetter: "TrustInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exact", GoGetter: "Exact"},
			_jsii_.MemberProperty{JsiiProperty: "exactInput", GoGetter: "ExactInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrust)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcm)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcmOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcmOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateAuthorityArns", GoGetter: "CertificateAuthorityArns"},
			_jsii_.MemberProperty{JsiiProperty: "certificateAuthorityArnsInput", GoGetter: "CertificateAuthorityArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustAcmOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFile",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acm", GoGetter: "Acm"},
			_jsii_.MemberProperty{JsiiProperty: "acmInput", GoGetter: "AcmInput"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAcm", GoMethod: "PutAcm"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcm", GoMethod: "ResetAcm"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsClientPolicyTlsValidationTrustSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecBackendDefaultsOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecBackendDefaultsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientPolicy", GoGetter: "ClientPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clientPolicyInput", GoGetter: "ClientPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putClientPolicy", GoMethod: "PutClientPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientPolicy", GoMethod: "ResetClientPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecBackendDefaultsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListener",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListener)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerConnectionPool",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerConnectionPoolGrpc",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerConnectionPoolGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerConnectionPoolGrpcOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerConnectionPoolGrpcOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxRequests", GoGetter: "MaxRequests"},
			_jsii_.MemberProperty{JsiiProperty: "maxRequestsInput", GoGetter: "MaxRequestsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerConnectionPoolGrpcOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerConnectionPoolHttp",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerConnectionPoolHttp2",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerConnectionPoolHttp2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerConnectionPoolHttp2OutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerConnectionPoolHttp2OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxRequests", GoGetter: "MaxRequests"},
			_jsii_.MemberProperty{JsiiProperty: "maxRequestsInput", GoGetter: "MaxRequestsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerConnectionPoolHttp2OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerConnectionPoolHttpOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerConnectionPoolHttpOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnections", GoGetter: "MaxConnections"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnectionsInput", GoGetter: "MaxConnectionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxPendingRequests", GoGetter: "MaxPendingRequests"},
			_jsii_.MemberProperty{JsiiProperty: "maxPendingRequestsInput", GoGetter: "MaxPendingRequestsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxPendingRequests", GoMethod: "ResetMaxPendingRequests"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerConnectionPoolHttpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerConnectionPoolOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerConnectionPoolOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grpc", GoGetter: "Grpc"},
			_jsii_.MemberProperty{JsiiProperty: "grpcInput", GoGetter: "GrpcInput"},
			_jsii_.MemberProperty{JsiiProperty: "http", GoGetter: "Http"},
			_jsii_.MemberProperty{JsiiProperty: "http2", GoGetter: "Http2"},
			_jsii_.MemberProperty{JsiiProperty: "http2Input", GoGetter: "Http2Input"},
			_jsii_.MemberProperty{JsiiProperty: "httpInput", GoGetter: "HttpInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putGrpc", GoMethod: "PutGrpc"},
			_jsii_.MemberMethod{JsiiMethod: "putHttp", GoMethod: "PutHttp"},
			_jsii_.MemberMethod{JsiiMethod: "putHttp2", GoMethod: "PutHttp2"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrpc", GoMethod: "ResetGrpc"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp", GoMethod: "ResetHttp"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp2", GoMethod: "ResetHttp2"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerConnectionPoolOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerHealthCheck",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerHealthCheck)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerHealthCheckOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerHealthCheckOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "healthyThreshold", GoGetter: "HealthyThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "healthyThresholdInput", GoGetter: "HealthyThresholdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "intervalMillis", GoGetter: "IntervalMillis"},
			_jsii_.MemberProperty{JsiiProperty: "intervalMillisInput", GoGetter: "IntervalMillisInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetPort", GoMethod: "ResetPort"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutMillis", GoGetter: "TimeoutMillis"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutMillisInput", GoGetter: "TimeoutMillisInput"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyThreshold", GoGetter: "UnhealthyThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyThresholdInput", GoGetter: "UnhealthyThresholdInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerHealthCheckOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connectionPool", GoGetter: "ConnectionPool"},
			_jsii_.MemberProperty{JsiiProperty: "connectionPoolInput", GoGetter: "ConnectionPoolInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheck", GoGetter: "HealthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckInput", GoGetter: "HealthCheckInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "portMapping", GoGetter: "PortMapping"},
			_jsii_.MemberProperty{JsiiProperty: "portMappingInput", GoGetter: "PortMappingInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConnectionPool", GoMethod: "PutConnectionPool"},
			_jsii_.MemberMethod{JsiiMethod: "putHealthCheck", GoMethod: "PutHealthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "putPortMapping", GoMethod: "PutPortMapping"},
			_jsii_.MemberMethod{JsiiMethod: "putTls", GoMethod: "PutTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionPool", GoMethod: "ResetConnectionPool"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheck", GoMethod: "ResetHealthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetTls", GoMethod: "ResetTls"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tls", GoGetter: "Tls"},
			_jsii_.MemberProperty{JsiiProperty: "tlsInput", GoGetter: "TlsInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerPortMapping",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerPortMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerPortMappingOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerPortMappingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerPortMappingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTls",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsCertificate",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsCertificateAcm",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsCertificateAcm)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsCertificateAcmOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsCertificateAcmOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateArn", GoGetter: "CertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "certificateArnInput", GoGetter: "CertificateArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateAcmOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsCertificateFile",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsCertificateFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsCertificateFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsCertificateFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acm", GoGetter: "Acm"},
			_jsii_.MemberProperty{JsiiProperty: "acmInput", GoGetter: "AcmInput"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAcm", GoMethod: "PutAcm"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcm", GoMethod: "ResetAcm"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsCertificateSds",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsCertificateSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsCertificateSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsCertificateSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsCertificateSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "modeInput", GoGetter: "ModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "putValidation", GoMethod: "PutValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidation", GoMethod: "ResetValidation"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "validation", GoGetter: "Validation"},
			_jsii_.MemberProperty{JsiiProperty: "validationInput", GoGetter: "ValidationInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidation",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putSubjectAlternativeNames", GoMethod: "PutSubjectAlternativeNames"},
			_jsii_.MemberMethod{JsiiMethod: "putTrust", GoMethod: "PutTrust"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubjectAlternativeNames", GoMethod: "ResetSubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNames", GoGetter: "SubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNamesInput", GoGetter: "SubjectAlternativeNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "trust", GoGetter: "Trust"},
			_jsii_.MemberProperty{JsiiProperty: "trustInput", GoGetter: "TrustInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsValidationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNames",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesMatch",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesMatchOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exact", GoGetter: "Exact"},
			_jsii_.MemberProperty{JsiiProperty: "exactInput", GoGetter: "ExactInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsValidationSubjectAlternativeNamesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationTrust",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationTrust)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationTrustFile",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationTrustFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationTrustFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationTrustFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsValidationTrustFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationTrustOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationTrustOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsValidationTrustOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationTrustSds",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationTrustSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecListenerTlsValidationTrustSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecListenerTlsValidationTrustSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecListenerTlsValidationTrustSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecLogging",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecLogging)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecLoggingAccessLog",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecLoggingAccessLog)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecLoggingAccessLogFile",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecLoggingAccessLogFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecLoggingAccessLogFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecLoggingAccessLogFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecLoggingAccessLogFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecLoggingAccessLogOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecLoggingAccessLogOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecLoggingAccessLogOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecLoggingOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecLoggingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessLog", GoGetter: "AccessLog"},
			_jsii_.MemberProperty{JsiiProperty: "accessLogInput", GoGetter: "AccessLogInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAccessLog", GoMethod: "PutAccessLog"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessLog", GoMethod: "ResetAccessLog"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecLoggingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualGatewaySpecOutputReference",
		reflect.TypeOf((*AppmeshVirtualGatewaySpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendDefaults", GoGetter: "BackendDefaults"},
			_jsii_.MemberProperty{JsiiProperty: "backendDefaultsInput", GoGetter: "BackendDefaultsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "listener", GoGetter: "Listener"},
			_jsii_.MemberProperty{JsiiProperty: "listenerInput", GoGetter: "ListenerInput"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "loggingInput", GoGetter: "LoggingInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBackendDefaults", GoMethod: "PutBackendDefaults"},
			_jsii_.MemberMethod{JsiiMethod: "putListener", GoMethod: "PutListener"},
			_jsii_.MemberMethod{JsiiMethod: "putLogging", GoMethod: "PutLogging"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackendDefaults", GoMethod: "ResetBackendDefaults"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogging", GoMethod: "ResetLogging"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualGatewaySpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNode",
		reflect.TypeOf((*AppmeshVirtualNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDate", GoGetter: "LastUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshNameInput", GoGetter: "MeshNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwnerInput", GoGetter: "MeshOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putSpec", GoMethod: "PutSpec"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetMeshOwner", GoMethod: "ResetMeshOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "resourceOwner", GoGetter: "ResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "specInput", GoGetter: "SpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNode{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeConfig",
		reflect.TypeOf((*AppmeshVirtualNodeConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpec",
		reflect.TypeOf((*AppmeshVirtualNodeSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackend",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaults",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaults)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicy",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putTls", GoMethod: "PutTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetTls", GoMethod: "ResetTls"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tls", GoGetter: "Tls"},
			_jsii_.MemberProperty{JsiiProperty: "tlsInput", GoGetter: "TlsInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSds",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsCertificateSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "enforce", GoGetter: "Enforce"},
			_jsii_.MemberProperty{JsiiProperty: "enforceInput", GoGetter: "EnforceInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "portsInput", GoGetter: "PortsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "putValidation", GoMethod: "PutValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnforce", GoMethod: "ResetEnforce"},
			_jsii_.MemberMethod{JsiiMethod: "resetPorts", GoMethod: "ResetPorts"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "validation", GoGetter: "Validation"},
			_jsii_.MemberProperty{JsiiProperty: "validationInput", GoGetter: "ValidationInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidation",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putSubjectAlternativeNames", GoMethod: "PutSubjectAlternativeNames"},
			_jsii_.MemberMethod{JsiiMethod: "putTrust", GoMethod: "PutTrust"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubjectAlternativeNames", GoMethod: "ResetSubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNames", GoGetter: "SubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNamesInput", GoGetter: "SubjectAlternativeNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "trust", GoGetter: "Trust"},
			_jsii_.MemberProperty{JsiiProperty: "trustInput", GoGetter: "TrustInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exact", GoGetter: "Exact"},
			_jsii_.MemberProperty{JsiiProperty: "exactInput", GoGetter: "ExactInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrust",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrust)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustAcm",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustAcm)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustAcmOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustAcmOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateAuthorityArns", GoGetter: "CertificateAuthorityArns"},
			_jsii_.MemberProperty{JsiiProperty: "certificateAuthorityArnsInput", GoGetter: "CertificateAuthorityArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustAcmOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustFile",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acm", GoGetter: "Acm"},
			_jsii_.MemberProperty{JsiiProperty: "acmInput", GoGetter: "AcmInput"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAcm", GoMethod: "PutAcm"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcm", GoMethod: "ResetAcm"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSds",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsClientPolicyTlsValidationTrustSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendDefaultsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendDefaultsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientPolicy", GoGetter: "ClientPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clientPolicyInput", GoGetter: "ClientPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putClientPolicy", GoMethod: "PutClientPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientPolicy", GoMethod: "ResetClientPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendDefaultsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualService",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualService)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicy",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putTls", GoMethod: "PutTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetTls", GoMethod: "ResetTls"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tls", GoGetter: "Tls"},
			_jsii_.MemberProperty{JsiiProperty: "tlsInput", GoGetter: "TlsInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTls",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificate",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateFile",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateSds",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsCertificateSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "enforce", GoGetter: "Enforce"},
			_jsii_.MemberProperty{JsiiProperty: "enforceInput", GoGetter: "EnforceInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "ports", GoGetter: "Ports"},
			_jsii_.MemberProperty{JsiiProperty: "portsInput", GoGetter: "PortsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "putValidation", GoMethod: "PutValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnforce", GoMethod: "ResetEnforce"},
			_jsii_.MemberMethod{JsiiMethod: "resetPorts", GoMethod: "ResetPorts"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "validation", GoGetter: "Validation"},
			_jsii_.MemberProperty{JsiiProperty: "validationInput", GoGetter: "ValidationInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidation",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putSubjectAlternativeNames", GoMethod: "PutSubjectAlternativeNames"},
			_jsii_.MemberMethod{JsiiMethod: "putTrust", GoMethod: "PutTrust"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubjectAlternativeNames", GoMethod: "ResetSubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNames", GoGetter: "SubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNamesInput", GoGetter: "SubjectAlternativeNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "trust", GoGetter: "Trust"},
			_jsii_.MemberProperty{JsiiProperty: "trustInput", GoGetter: "TrustInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNames",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesMatch",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exact", GoGetter: "Exact"},
			_jsii_.MemberProperty{JsiiProperty: "exactInput", GoGetter: "ExactInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrust",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrust)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustAcm",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustAcm)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateAuthorityArns", GoGetter: "CertificateAuthorityArns"},
			_jsii_.MemberProperty{JsiiProperty: "certificateAuthorityArnsInput", GoGetter: "CertificateAuthorityArnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustFile",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acm", GoGetter: "Acm"},
			_jsii_.MemberProperty{JsiiProperty: "acmInput", GoGetter: "AcmInput"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAcm", GoMethod: "PutAcm"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcm", GoMethod: "ResetAcm"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustSds",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecBackendVirtualServiceOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecBackendVirtualServiceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientPolicy", GoGetter: "ClientPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clientPolicyInput", GoGetter: "ClientPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putClientPolicy", GoMethod: "PutClientPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientPolicy", GoMethod: "ResetClientPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceName", GoGetter: "VirtualServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualServiceNameInput", GoGetter: "VirtualServiceNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecBackendVirtualServiceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListener",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListener)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPool",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPool)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPoolGrpc",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPoolGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPoolGrpcOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPoolGrpcOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxRequests", GoGetter: "MaxRequests"},
			_jsii_.MemberProperty{JsiiProperty: "maxRequestsInput", GoGetter: "MaxRequestsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerConnectionPoolGrpcOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPoolHttp",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPoolHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPoolHttp2",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPoolHttp2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPoolHttp2OutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPoolHttp2OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxRequests", GoGetter: "MaxRequests"},
			_jsii_.MemberProperty{JsiiProperty: "maxRequestsInput", GoGetter: "MaxRequestsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerConnectionPoolHttp2OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPoolHttpOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPoolHttpOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnections", GoGetter: "MaxConnections"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnectionsInput", GoGetter: "MaxConnectionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxPendingRequests", GoGetter: "MaxPendingRequests"},
			_jsii_.MemberProperty{JsiiProperty: "maxPendingRequestsInput", GoGetter: "MaxPendingRequestsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxPendingRequests", GoMethod: "ResetMaxPendingRequests"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerConnectionPoolHttpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPoolOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPoolOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grpc", GoGetter: "Grpc"},
			_jsii_.MemberProperty{JsiiProperty: "grpcInput", GoGetter: "GrpcInput"},
			_jsii_.MemberProperty{JsiiProperty: "http", GoGetter: "Http"},
			_jsii_.MemberProperty{JsiiProperty: "http2", GoGetter: "Http2"},
			_jsii_.MemberProperty{JsiiProperty: "http2Input", GoGetter: "Http2Input"},
			_jsii_.MemberProperty{JsiiProperty: "httpInput", GoGetter: "HttpInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putGrpc", GoMethod: "PutGrpc"},
			_jsii_.MemberMethod{JsiiMethod: "putHttp", GoMethod: "PutHttp"},
			_jsii_.MemberMethod{JsiiMethod: "putHttp2", GoMethod: "PutHttp2"},
			_jsii_.MemberMethod{JsiiMethod: "putTcp", GoMethod: "PutTcp"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrpc", GoMethod: "ResetGrpc"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp", GoMethod: "ResetHttp"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp2", GoMethod: "ResetHttp2"},
			_jsii_.MemberMethod{JsiiMethod: "resetTcp", GoMethod: "ResetTcp"},
			_jsii_.MemberProperty{JsiiProperty: "tcp", GoGetter: "Tcp"},
			_jsii_.MemberProperty{JsiiProperty: "tcpInput", GoGetter: "TcpInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerConnectionPoolOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPoolTcp",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPoolTcp)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerConnectionPoolTcpOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerConnectionPoolTcpOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnections", GoGetter: "MaxConnections"},
			_jsii_.MemberProperty{JsiiProperty: "maxConnectionsInput", GoGetter: "MaxConnectionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerConnectionPoolTcpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerHealthCheck",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerHealthCheck)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerHealthCheckOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerHealthCheckOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "healthyThreshold", GoGetter: "HealthyThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "healthyThresholdInput", GoGetter: "HealthyThresholdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "intervalMillis", GoGetter: "IntervalMillis"},
			_jsii_.MemberProperty{JsiiProperty: "intervalMillisInput", GoGetter: "IntervalMillisInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetPort", GoMethod: "ResetPort"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutMillis", GoGetter: "TimeoutMillis"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutMillisInput", GoGetter: "TimeoutMillisInput"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyThreshold", GoGetter: "UnhealthyThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "unhealthyThresholdInput", GoGetter: "UnhealthyThresholdInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerHealthCheckOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerOutlierDetection",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerOutlierDetection)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDuration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDurationOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionBaseEjectionDurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerOutlierDetectionInterval",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerOutlierDetectionInterval)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerOutlierDetectionIntervalOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerOutlierDetectionIntervalOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionIntervalOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "baseEjectionDuration", GoGetter: "BaseEjectionDuration"},
			_jsii_.MemberProperty{JsiiProperty: "baseEjectionDurationInput", GoGetter: "BaseEjectionDurationInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInput", GoGetter: "IntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxEjectionPercent", GoGetter: "MaxEjectionPercent"},
			_jsii_.MemberProperty{JsiiProperty: "maxEjectionPercentInput", GoGetter: "MaxEjectionPercentInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxServerErrors", GoGetter: "MaxServerErrors"},
			_jsii_.MemberProperty{JsiiProperty: "maxServerErrorsInput", GoGetter: "MaxServerErrorsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBaseEjectionDuration", GoMethod: "PutBaseEjectionDuration"},
			_jsii_.MemberMethod{JsiiMethod: "putInterval", GoMethod: "PutInterval"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerOutlierDetectionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connectionPool", GoGetter: "ConnectionPool"},
			_jsii_.MemberProperty{JsiiProperty: "connectionPoolInput", GoGetter: "ConnectionPoolInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheck", GoGetter: "HealthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckInput", GoGetter: "HealthCheckInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "outlierDetection", GoGetter: "OutlierDetection"},
			_jsii_.MemberProperty{JsiiProperty: "outlierDetectionInput", GoGetter: "OutlierDetectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "portMapping", GoGetter: "PortMapping"},
			_jsii_.MemberProperty{JsiiProperty: "portMappingInput", GoGetter: "PortMappingInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConnectionPool", GoMethod: "PutConnectionPool"},
			_jsii_.MemberMethod{JsiiMethod: "putHealthCheck", GoMethod: "PutHealthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "putOutlierDetection", GoMethod: "PutOutlierDetection"},
			_jsii_.MemberMethod{JsiiMethod: "putPortMapping", GoMethod: "PutPortMapping"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeout", GoMethod: "PutTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "putTls", GoMethod: "PutTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionPool", GoMethod: "ResetConnectionPool"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheck", GoMethod: "ResetHealthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutlierDetection", GoMethod: "ResetOutlierDetection"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetTls", GoMethod: "ResetTls"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "tls", GoGetter: "Tls"},
			_jsii_.MemberProperty{JsiiProperty: "tlsInput", GoGetter: "TlsInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerPortMapping",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerPortMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerPortMappingOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerPortMappingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerPortMappingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeout",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeout)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutGrpc",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutGrpcIdle",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutGrpcIdle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutGrpcIdleOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutGrpcIdleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutGrpcIdleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutGrpcOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutGrpcOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idle", GoGetter: "Idle"},
			_jsii_.MemberProperty{JsiiProperty: "idleInput", GoGetter: "IdleInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "perRequest", GoGetter: "PerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "perRequestInput", GoGetter: "PerRequestInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIdle", GoMethod: "PutIdle"},
			_jsii_.MemberMethod{JsiiMethod: "putPerRequest", GoMethod: "PutPerRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdle", GoMethod: "ResetIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerRequest", GoMethod: "ResetPerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutGrpcOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutGrpcPerRequest",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutGrpcPerRequest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutGrpcPerRequestOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutGrpcPerRequestOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutGrpcPerRequestOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttp",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttp2",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttp2)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttp2Idle",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttp2Idle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttp2IdleOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttp2IdleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutHttp2IdleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttp2OutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttp2OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idle", GoGetter: "Idle"},
			_jsii_.MemberProperty{JsiiProperty: "idleInput", GoGetter: "IdleInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "perRequest", GoGetter: "PerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "perRequestInput", GoGetter: "PerRequestInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIdle", GoMethod: "PutIdle"},
			_jsii_.MemberMethod{JsiiMethod: "putPerRequest", GoMethod: "PutPerRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdle", GoMethod: "ResetIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerRequest", GoMethod: "ResetPerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutHttp2OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttp2PerRequest",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttp2PerRequest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttp2PerRequestOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttp2PerRequestOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutHttp2PerRequestOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttpIdle",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttpIdle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttpIdleOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttpIdleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutHttpIdleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttpOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttpOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idle", GoGetter: "Idle"},
			_jsii_.MemberProperty{JsiiProperty: "idleInput", GoGetter: "IdleInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "perRequest", GoGetter: "PerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "perRequestInput", GoGetter: "PerRequestInput"},
			_jsii_.MemberMethod{JsiiMethod: "putIdle", GoMethod: "PutIdle"},
			_jsii_.MemberMethod{JsiiMethod: "putPerRequest", GoMethod: "PutPerRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdle", GoMethod: "ResetIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetPerRequest", GoMethod: "ResetPerRequest"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutHttpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttpPerRequest",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttpPerRequest)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutHttpPerRequestOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutHttpPerRequestOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutHttpPerRequestOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grpc", GoGetter: "Grpc"},
			_jsii_.MemberProperty{JsiiProperty: "grpcInput", GoGetter: "GrpcInput"},
			_jsii_.MemberProperty{JsiiProperty: "http", GoGetter: "Http"},
			_jsii_.MemberProperty{JsiiProperty: "http2", GoGetter: "Http2"},
			_jsii_.MemberProperty{JsiiProperty: "http2Input", GoGetter: "Http2Input"},
			_jsii_.MemberProperty{JsiiProperty: "httpInput", GoGetter: "HttpInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putGrpc", GoMethod: "PutGrpc"},
			_jsii_.MemberMethod{JsiiMethod: "putHttp", GoMethod: "PutHttp"},
			_jsii_.MemberMethod{JsiiMethod: "putHttp2", GoMethod: "PutHttp2"},
			_jsii_.MemberMethod{JsiiMethod: "putTcp", GoMethod: "PutTcp"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrpc", GoMethod: "ResetGrpc"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp", GoMethod: "ResetHttp"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttp2", GoMethod: "ResetHttp2"},
			_jsii_.MemberMethod{JsiiMethod: "resetTcp", GoMethod: "ResetTcp"},
			_jsii_.MemberProperty{JsiiProperty: "tcp", GoGetter: "Tcp"},
			_jsii_.MemberProperty{JsiiProperty: "tcpInput", GoGetter: "TcpInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutTcp",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutTcp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutTcpIdle",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutTcpIdle)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutTcpIdleOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutTcpIdleOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutTcpIdleOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTimeoutTcpOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTimeoutTcpOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idle", GoGetter: "Idle"},
			_jsii_.MemberProperty{JsiiProperty: "idleInput", GoGetter: "IdleInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putIdle", GoMethod: "PutIdle"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdle", GoMethod: "ResetIdle"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTimeoutTcpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTls",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsCertificate",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsCertificateAcm",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsCertificateAcm)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsCertificateAcmOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsCertificateAcmOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateArn", GoGetter: "CertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "certificateArnInput", GoGetter: "CertificateArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsCertificateAcmOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsCertificateFile",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsCertificateFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsCertificateFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsCertificateFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyInput", GoGetter: "PrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsCertificateFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsCertificateOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsCertificateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acm", GoGetter: "Acm"},
			_jsii_.MemberProperty{JsiiProperty: "acmInput", GoGetter: "AcmInput"},
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAcm", GoMethod: "PutAcm"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcm", GoMethod: "ResetAcm"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsCertificateSds",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsCertificateSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsCertificateSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsCertificateSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsCertificateSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "modeInput", GoGetter: "ModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "putValidation", GoMethod: "PutValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidation", GoMethod: "ResetValidation"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "validation", GoGetter: "Validation"},
			_jsii_.MemberProperty{JsiiProperty: "validationInput", GoGetter: "ValidationInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidation",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putSubjectAlternativeNames", GoMethod: "PutSubjectAlternativeNames"},
			_jsii_.MemberMethod{JsiiMethod: "putTrust", GoMethod: "PutTrust"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubjectAlternativeNames", GoMethod: "ResetSubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNames", GoGetter: "SubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNamesInput", GoGetter: "SubjectAlternativeNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "trust", GoGetter: "Trust"},
			_jsii_.MemberProperty{JsiiProperty: "trustInput", GoGetter: "TrustInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsValidationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNames",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatch",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatchOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatchOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exact", GoGetter: "Exact"},
			_jsii_.MemberProperty{JsiiProperty: "exactInput", GoGetter: "ExactInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsValidationSubjectAlternativeNamesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationTrust",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationTrust)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationTrustFile",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationTrustFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationTrustFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationTrustFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certificateChain", GoGetter: "CertificateChain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateChainInput", GoGetter: "CertificateChainInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsValidationTrustFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationTrustOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationTrustOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "putSds", GoMethod: "PutSds"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSds", GoMethod: "ResetSds"},
			_jsii_.MemberProperty{JsiiProperty: "sds", GoGetter: "Sds"},
			_jsii_.MemberProperty{JsiiProperty: "sdsInput", GoGetter: "SdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsValidationTrustOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationTrustSds",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationTrustSds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecListenerTlsValidationTrustSdsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecListenerTlsValidationTrustSdsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "secretName", GoGetter: "SecretName"},
			_jsii_.MemberProperty{JsiiProperty: "secretNameInput", GoGetter: "SecretNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecListenerTlsValidationTrustSdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecLogging",
		reflect.TypeOf((*AppmeshVirtualNodeSpecLogging)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecLoggingAccessLog",
		reflect.TypeOf((*AppmeshVirtualNodeSpecLoggingAccessLog)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecLoggingAccessLogFile",
		reflect.TypeOf((*AppmeshVirtualNodeSpecLoggingAccessLogFile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecLoggingAccessLogFileOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecLoggingAccessLogFileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecLoggingAccessLogFileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecLoggingAccessLogOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecLoggingAccessLogOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "file", GoGetter: "File"},
			_jsii_.MemberProperty{JsiiProperty: "fileInput", GoGetter: "FileInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putFile", GoMethod: "PutFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetFile", GoMethod: "ResetFile"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecLoggingAccessLogOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecLoggingOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecLoggingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessLog", GoGetter: "AccessLog"},
			_jsii_.MemberProperty{JsiiProperty: "accessLogInput", GoGetter: "AccessLogInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAccessLog", GoMethod: "PutAccessLog"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessLog", GoMethod: "ResetAccessLog"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecLoggingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendDefaults", GoGetter: "BackendDefaults"},
			_jsii_.MemberProperty{JsiiProperty: "backendDefaultsInput", GoGetter: "BackendDefaultsInput"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "listener", GoGetter: "Listener"},
			_jsii_.MemberProperty{JsiiProperty: "listenerInput", GoGetter: "ListenerInput"},
			_jsii_.MemberProperty{JsiiProperty: "logging", GoGetter: "Logging"},
			_jsii_.MemberProperty{JsiiProperty: "loggingInput", GoGetter: "LoggingInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBackendDefaults", GoMethod: "PutBackendDefaults"},
			_jsii_.MemberMethod{JsiiMethod: "putListener", GoMethod: "PutListener"},
			_jsii_.MemberMethod{JsiiMethod: "putLogging", GoMethod: "PutLogging"},
			_jsii_.MemberMethod{JsiiMethod: "putServiceDiscovery", GoMethod: "PutServiceDiscovery"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackend", GoMethod: "ResetBackend"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackendDefaults", GoMethod: "ResetBackendDefaults"},
			_jsii_.MemberMethod{JsiiMethod: "resetListener", GoMethod: "ResetListener"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogging", GoMethod: "ResetLogging"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceDiscovery", GoMethod: "ResetServiceDiscovery"},
			_jsii_.MemberProperty{JsiiProperty: "serviceDiscovery", GoGetter: "ServiceDiscovery"},
			_jsii_.MemberProperty{JsiiProperty: "serviceDiscoveryInput", GoGetter: "ServiceDiscoveryInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecServiceDiscovery",
		reflect.TypeOf((*AppmeshVirtualNodeSpecServiceDiscovery)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMap",
		reflect.TypeOf((*AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMap)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMapOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMapOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
			_jsii_.MemberProperty{JsiiProperty: "attributesInput", GoGetter: "AttributesInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceName", GoGetter: "NamespaceName"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceNameInput", GoGetter: "NamespaceNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttributes", GoMethod: "ResetAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNameInput", GoGetter: "ServiceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecServiceDiscoveryAwsCloudMapOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecServiceDiscoveryDns",
		reflect.TypeOf((*AppmeshVirtualNodeSpecServiceDiscoveryDns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecServiceDiscoveryDnsOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecServiceDiscoveryDnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "hostnameInput", GoGetter: "HostnameInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecServiceDiscoveryDnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualNodeSpecServiceDiscoveryOutputReference",
		reflect.TypeOf((*AppmeshVirtualNodeSpecServiceDiscoveryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsCloudMap", GoGetter: "AwsCloudMap"},
			_jsii_.MemberProperty{JsiiProperty: "awsCloudMapInput", GoGetter: "AwsCloudMapInput"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberProperty{JsiiProperty: "dnsInput", GoGetter: "DnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putAwsCloudMap", GoMethod: "PutAwsCloudMap"},
			_jsii_.MemberMethod{JsiiMethod: "putDns", GoMethod: "PutDns"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsCloudMap", GoMethod: "ResetAwsCloudMap"},
			_jsii_.MemberMethod{JsiiMethod: "resetDns", GoMethod: "ResetDns"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualNodeSpecServiceDiscoveryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualRouter",
		reflect.TypeOf((*AppmeshVirtualRouter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDate", GoGetter: "LastUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshNameInput", GoGetter: "MeshNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwnerInput", GoGetter: "MeshOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putSpec", GoMethod: "PutSpec"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetMeshOwner", GoMethod: "ResetMeshOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "resourceOwner", GoGetter: "ResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "specInput", GoGetter: "SpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualRouter{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualRouterConfig",
		reflect.TypeOf((*AppmeshVirtualRouterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualRouterSpec",
		reflect.TypeOf((*AppmeshVirtualRouterSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualRouterSpecListener",
		reflect.TypeOf((*AppmeshVirtualRouterSpecListener)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualRouterSpecListenerOutputReference",
		reflect.TypeOf((*AppmeshVirtualRouterSpecListenerOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "portMapping", GoGetter: "PortMapping"},
			_jsii_.MemberProperty{JsiiProperty: "portMappingInput", GoGetter: "PortMappingInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPortMapping", GoMethod: "PutPortMapping"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualRouterSpecListenerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualRouterSpecListenerPortMapping",
		reflect.TypeOf((*AppmeshVirtualRouterSpecListenerPortMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualRouterSpecListenerPortMappingOutputReference",
		reflect.TypeOf((*AppmeshVirtualRouterSpecListenerPortMappingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualRouterSpecListenerPortMappingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualRouterSpecOutputReference",
		reflect.TypeOf((*AppmeshVirtualRouterSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "listener", GoGetter: "Listener"},
			_jsii_.MemberProperty{JsiiProperty: "listenerInput", GoGetter: "ListenerInput"},
			_jsii_.MemberMethod{JsiiMethod: "putListener", GoMethod: "PutListener"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualRouterSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualService",
		reflect.TypeOf((*AppmeshVirtualService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDate", GoGetter: "LastUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshNameInput", GoGetter: "MeshNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwnerInput", GoGetter: "MeshOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putSpec", GoMethod: "PutSpec"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetMeshOwner", GoMethod: "ResetMeshOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "resourceOwner", GoGetter: "ResourceOwner"},
			_jsii_.MemberProperty{JsiiProperty: "spec", GoGetter: "Spec"},
			_jsii_.MemberProperty{JsiiProperty: "specInput", GoGetter: "SpecInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAll", GoGetter: "TagsAll"},
			_jsii_.MemberProperty{JsiiProperty: "tagsAllInput", GoGetter: "TagsAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualService{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualServiceConfig",
		reflect.TypeOf((*AppmeshVirtualServiceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualServiceSpec",
		reflect.TypeOf((*AppmeshVirtualServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualServiceSpecOutputReference",
		reflect.TypeOf((*AppmeshVirtualServiceSpecOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "providerInput", GoGetter: "ProviderInput"},
			_jsii_.MemberMethod{JsiiMethod: "putProvider", GoMethod: "PutProvider"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvider", GoMethod: "ResetProvider"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualServiceSpecOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualServiceSpecProvider",
		reflect.TypeOf((*AppmeshVirtualServiceSpecProvider)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualServiceSpecProviderOutputReference",
		reflect.TypeOf((*AppmeshVirtualServiceSpecProviderOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putVirtualNode", GoMethod: "PutVirtualNode"},
			_jsii_.MemberMethod{JsiiMethod: "putVirtualRouter", GoMethod: "PutVirtualRouter"},
			_jsii_.MemberMethod{JsiiMethod: "resetVirtualNode", GoMethod: "ResetVirtualNode"},
			_jsii_.MemberMethod{JsiiMethod: "resetVirtualRouter", GoMethod: "ResetVirtualRouter"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNode", GoGetter: "VirtualNode"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNodeInput", GoGetter: "VirtualNodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouter", GoGetter: "VirtualRouter"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterInput", GoGetter: "VirtualRouterInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualServiceSpecProviderOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualServiceSpecProviderVirtualNode",
		reflect.TypeOf((*AppmeshVirtualServiceSpecProviderVirtualNode)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualServiceSpecProviderVirtualNodeOutputReference",
		reflect.TypeOf((*AppmeshVirtualServiceSpecProviderVirtualNodeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNodeName", GoGetter: "VirtualNodeName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNodeNameInput", GoGetter: "VirtualNodeNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualServiceSpecProviderVirtualNodeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.AppmeshVirtualServiceSpecProviderVirtualRouter",
		reflect.TypeOf((*AppmeshVirtualServiceSpecProviderVirtualRouter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.AppmeshVirtualServiceSpecProviderVirtualRouterOutputReference",
		reflect.TypeOf((*AppmeshVirtualServiceSpecProviderVirtualRouterOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterName", GoGetter: "VirtualRouterName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterNameInput", GoGetter: "VirtualRouterNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppmeshVirtualServiceSpecProviderVirtualRouterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.DataAwsAppmeshMesh",
		reflect.TypeOf((*DataAwsAppmeshMesh)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDate", GoGetter: "LastUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwnerInput", GoGetter: "MeshOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetMeshOwner", GoMethod: "ResetMeshOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "resourceOwner", GoGetter: "ResourceOwner"},
			_jsii_.MemberMethod{JsiiMethod: "spec", GoMethod: "Spec"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsAppmeshMesh{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.DataAwsAppmeshMeshConfig",
		reflect.TypeOf((*DataAwsAppmeshMeshConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.DataAwsAppmeshMeshSpec",
		reflect.TypeOf((*DataAwsAppmeshMeshSpec)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "egressFilter", GoGetter: "EgressFilter"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsAppmeshMeshSpec{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.DataAwsAppmeshMeshSpecEgressFilter",
		reflect.TypeOf((*DataAwsAppmeshMeshSpecEgressFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsAppmeshMeshSpecEgressFilter{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.DataAwsAppmeshVirtualService",
		reflect.TypeOf((*DataAwsAppmeshVirtualService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdDate", GoGetter: "CreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdatedDate", GoGetter: "LastUpdatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "meshName", GoGetter: "MeshName"},
			_jsii_.MemberProperty{JsiiProperty: "meshNameInput", GoGetter: "MeshNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwner", GoGetter: "MeshOwner"},
			_jsii_.MemberProperty{JsiiProperty: "meshOwnerInput", GoGetter: "MeshOwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetMeshOwner", GoMethod: "ResetMeshOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "resourceOwner", GoGetter: "ResourceOwner"},
			_jsii_.MemberMethod{JsiiMethod: "spec", GoMethod: "Spec"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsAppmeshVirtualService{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.AppMesh.DataAwsAppmeshVirtualServiceConfig",
		reflect.TypeOf((*DataAwsAppmeshVirtualServiceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.DataAwsAppmeshVirtualServiceSpec",
		reflect.TypeOf((*DataAwsAppmeshVirtualServiceSpec)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsAppmeshVirtualServiceSpec{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.DataAwsAppmeshVirtualServiceSpecProvider",
		reflect.TypeOf((*DataAwsAppmeshVirtualServiceSpecProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNode", GoGetter: "VirtualNode"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouter", GoGetter: "VirtualRouter"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsAppmeshVirtualServiceSpecProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.DataAwsAppmeshVirtualServiceSpecProviderVirtualNode",
		reflect.TypeOf((*DataAwsAppmeshVirtualServiceSpecProviderVirtualNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNodeName", GoGetter: "VirtualNodeName"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsAppmeshVirtualServiceSpecProviderVirtualNode{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.AppMesh.DataAwsAppmeshVirtualServiceSpecProviderVirtualRouter",
		reflect.TypeOf((*DataAwsAppmeshVirtualServiceSpecProviderVirtualRouter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouterName", GoGetter: "VirtualRouterName"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsAppmeshVirtualServiceSpecProviderVirtualRouter{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
}
