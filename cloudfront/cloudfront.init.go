package cloudfront

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicy",
		reflect.TypeOf((*CloudfrontCachePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtl", GoGetter: "DefaultTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtlInput", GoGetter: "DefaultTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "etagInput", GoGetter: "EtagInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "minTtl", GoGetter: "MinTtl"},
			_jsii_.MemberProperty{JsiiProperty: "minTtlInput", GoGetter: "MinTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInCacheKeyAndForwardedToOrigin", GoGetter: "ParametersInCacheKeyAndForwardedToOrigin"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInCacheKeyAndForwardedToOriginInput", GoGetter: "ParametersInCacheKeyAndForwardedToOriginInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putParametersInCacheKeyAndForwardedToOrigin", GoMethod: "PutParametersInCacheKeyAndForwardedToOrigin"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTtl", GoMethod: "ResetDefaultTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetEtag", GoMethod: "ResetEtag"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinTtl", GoMethod: "ResetMinTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetParametersInCacheKeyAndForwardedToOrigin", GoMethod: "ResetParametersInCacheKeyAndForwardedToOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontCachePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyConfig",
		reflect.TypeOf((*CloudfrontCachePolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cookieBehavior", GoGetter: "CookieBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "cookieBehaviorInput", GoGetter: "CookieBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesInput", GoGetter: "CookiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putCookies", GoMethod: "PutCookies"},
			_jsii_.MemberMethod{JsiiMethod: "resetCookies", GoMethod: "ResetCookies"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headerBehavior", GoGetter: "HeaderBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "headerBehaviorInput", GoGetter: "HeaderBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaders", GoMethod: "PutHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderBehavior", GoMethod: "ResetHeaderBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cookiesConfig", GoGetter: "CookiesConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesConfigInput", GoGetter: "CookiesConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableAcceptEncodingBrotli", GoGetter: "EnableAcceptEncodingBrotli"},
			_jsii_.MemberProperty{JsiiProperty: "enableAcceptEncodingBrotliInput", GoGetter: "EnableAcceptEncodingBrotliInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableAcceptEncodingGzip", GoGetter: "EnableAcceptEncodingGzip"},
			_jsii_.MemberProperty{JsiiProperty: "enableAcceptEncodingGzipInput", GoGetter: "EnableAcceptEncodingGzipInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headersConfig", GoGetter: "HeadersConfig"},
			_jsii_.MemberProperty{JsiiProperty: "headersConfigInput", GoGetter: "HeadersConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putCookiesConfig", GoMethod: "PutCookiesConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putHeadersConfig", GoMethod: "PutHeadersConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryStringsConfig", GoMethod: "PutQueryStringsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringsConfig", GoGetter: "QueryStringsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringsConfigInput", GoGetter: "QueryStringsConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableAcceptEncodingBrotli", GoMethod: "ResetEnableAcceptEncodingBrotli"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableAcceptEncodingGzip", GoMethod: "ResetEnableAcceptEncodingGzip"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryStrings", GoMethod: "PutQueryStrings"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringBehavior", GoGetter: "QueryStringBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringBehaviorInput", GoGetter: "QueryStringBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryStrings", GoGetter: "QueryStrings"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringsInput", GoGetter: "QueryStringsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStrings", GoMethod: "ResetQueryStrings"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference",
		reflect.TypeOf((*CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistribution",
		reflect.TypeOf((*CloudfrontDistribution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "aliases", GoGetter: "Aliases"},
			_jsii_.MemberProperty{JsiiProperty: "aliasesInput", GoGetter: "AliasesInput"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "callerReference", GoGetter: "CallerReference"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customErrorResponse", GoGetter: "CustomErrorResponse"},
			_jsii_.MemberProperty{JsiiProperty: "customErrorResponseInput", GoGetter: "CustomErrorResponseInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCacheBehavior", GoGetter: "DefaultCacheBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCacheBehaviorInput", GoGetter: "DefaultCacheBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRootObject", GoGetter: "DefaultRootObject"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRootObjectInput", GoGetter: "DefaultRootObjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hostedZoneId", GoGetter: "HostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "httpVersion", GoGetter: "HttpVersion"},
			_jsii_.MemberProperty{JsiiProperty: "httpVersionInput", GoGetter: "HttpVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inProgressValidationBatches", GoGetter: "InProgressValidationBatches"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isIpv6Enabled", GoGetter: "IsIpv6Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "isIpv6EnabledInput", GoGetter: "IsIpv6EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "lastModifiedTime", GoGetter: "LastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orderedCacheBehavior", GoGetter: "OrderedCacheBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "orderedCacheBehaviorInput", GoGetter: "OrderedCacheBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originGroup", GoGetter: "OriginGroup"},
			_jsii_.MemberProperty{JsiiProperty: "originGroupInput", GoGetter: "OriginGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "priceClass", GoGetter: "PriceClass"},
			_jsii_.MemberProperty{JsiiProperty: "priceClassInput", GoGetter: "PriceClassInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putDefaultCacheBehavior", GoMethod: "PutDefaultCacheBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putRestrictions", GoMethod: "PutRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "putViewerCertificate", GoMethod: "PutViewerCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAliases", GoMethod: "ResetAliases"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomErrorResponse", GoMethod: "ResetCustomErrorResponse"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRootObject", GoMethod: "ResetDefaultRootObject"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpVersion", GoMethod: "ResetHttpVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsIpv6Enabled", GoMethod: "ResetIsIpv6Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrderedCacheBehavior", GoMethod: "ResetOrderedCacheBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginGroup", GoMethod: "ResetOriginGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriceClass", GoMethod: "ResetPriceClass"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetainOnDelete", GoMethod: "ResetRetainOnDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsAll", GoMethod: "ResetTagsAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetWaitForDeployment", GoMethod: "ResetWaitForDeployment"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebAclId", GoMethod: "ResetWebAclId"},
			_jsii_.MemberProperty{JsiiProperty: "restrictions", GoGetter: "Restrictions"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionsInput", GoGetter: "RestrictionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "retainOnDelete", GoGetter: "RetainOnDelete"},
			_jsii_.MemberProperty{JsiiProperty: "retainOnDeleteInput", GoGetter: "RetainOnDeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
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
			_jsii_.MemberMethod{JsiiMethod: "trustedKeyGroups", GoMethod: "TrustedKeyGroups"},
			_jsii_.MemberMethod{JsiiMethod: "trustedSigners", GoMethod: "TrustedSigners"},
			_jsii_.MemberProperty{JsiiProperty: "viewerCertificate", GoGetter: "ViewerCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "viewerCertificateInput", GoGetter: "ViewerCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "waitForDeployment", GoGetter: "WaitForDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "waitForDeploymentInput", GoGetter: "WaitForDeploymentInput"},
			_jsii_.MemberProperty{JsiiProperty: "webAclId", GoGetter: "WebAclId"},
			_jsii_.MemberProperty{JsiiProperty: "webAclIdInput", GoGetter: "WebAclIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistribution{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionConfig",
		reflect.TypeOf((*CloudfrontDistributionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionCustomErrorResponse",
		reflect.TypeOf((*CloudfrontDistributionCustomErrorResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehavior",
		reflect.TypeOf((*CloudfrontDistributionDefaultCacheBehavior)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorForwardedValues",
		reflect.TypeOf((*CloudfrontDistributionDefaultCacheBehaviorForwardedValues)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies",
		reflect.TypeOf((*CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookies)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "forward", GoGetter: "Forward"},
			_jsii_.MemberProperty{JsiiProperty: "forwardInput", GoGetter: "ForwardInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhitelistedNames", GoMethod: "ResetWhitelistedNames"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNames", GoGetter: "WhitelistedNames"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNamesInput", GoGetter: "WhitelistedNamesInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesCookiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesInput", GoGetter: "CookiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putCookies", GoMethod: "PutCookies"},
			_jsii_.MemberProperty{JsiiProperty: "queryString", GoGetter: "QueryString"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeys", GoGetter: "QueryStringCacheKeys"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeysInput", GoGetter: "QueryStringCacheKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringInput", GoGetter: "QueryStringInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStringCacheKeys", GoMethod: "ResetQueryStringCacheKeys"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorForwardedValuesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation",
		reflect.TypeOf((*CloudfrontDistributionDefaultCacheBehaviorFunctionAssociation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation",
		reflect.TypeOf((*CloudfrontDistributionDefaultCacheBehaviorLambdaFunctionAssociation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionDefaultCacheBehaviorOutputReference",
		reflect.TypeOf((*CloudfrontDistributionDefaultCacheBehaviorOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowedMethods", GoGetter: "AllowedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "allowedMethodsInput", GoGetter: "AllowedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethods", GoGetter: "CachedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "cachedMethodsInput", GoGetter: "CachedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyId", GoGetter: "CachePolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "cachePolicyIdInput", GoGetter: "CachePolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "compress", GoGetter: "Compress"},
			_jsii_.MemberProperty{JsiiProperty: "compressInput", GoGetter: "CompressInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtl", GoGetter: "DefaultTtl"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtlInput", GoGetter: "DefaultTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionId", GoGetter: "FieldLevelEncryptionId"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLevelEncryptionIdInput", GoGetter: "FieldLevelEncryptionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValues", GoGetter: "ForwardedValues"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedValuesInput", GoGetter: "ForwardedValuesInput"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociation", GoGetter: "FunctionAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "functionAssociationInput", GoGetter: "FunctionAssociationInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociation", GoGetter: "LambdaFunctionAssociation"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionAssociationInput", GoGetter: "LambdaFunctionAssociationInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtlInput", GoGetter: "MaxTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "minTtl", GoGetter: "MinTtl"},
			_jsii_.MemberProperty{JsiiProperty: "minTtlInput", GoGetter: "MinTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyId", GoGetter: "OriginRequestPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "originRequestPolicyIdInput", GoGetter: "OriginRequestPolicyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putForwardedValues", GoMethod: "PutForwardedValues"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArn", GoGetter: "RealtimeLogConfigArn"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeLogConfigArnInput", GoGetter: "RealtimeLogConfigArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCachePolicyId", GoMethod: "ResetCachePolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompress", GoMethod: "ResetCompress"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTtl", GoMethod: "ResetDefaultTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetFieldLevelEncryptionId", GoMethod: "ResetFieldLevelEncryptionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetForwardedValues", GoMethod: "ResetForwardedValues"},
			_jsii_.MemberMethod{JsiiMethod: "resetFunctionAssociation", GoMethod: "ResetFunctionAssociation"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaFunctionAssociation", GoMethod: "ResetLambdaFunctionAssociation"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxTtl", GoMethod: "ResetMaxTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinTtl", GoMethod: "ResetMinTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginRequestPolicyId", GoMethod: "ResetOriginRequestPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealtimeLogConfigArn", GoMethod: "ResetRealtimeLogConfigArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetResponseHeadersPolicyId", GoMethod: "ResetResponseHeadersPolicyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSmoothStreaming", GoMethod: "ResetSmoothStreaming"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedKeyGroups", GoMethod: "ResetTrustedKeyGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedSigners", GoMethod: "ResetTrustedSigners"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyId", GoGetter: "ResponseHeadersPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeadersPolicyIdInput", GoGetter: "ResponseHeadersPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreaming", GoGetter: "SmoothStreaming"},
			_jsii_.MemberProperty{JsiiProperty: "smoothStreamingInput", GoGetter: "SmoothStreamingInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginId", GoGetter: "TargetOriginId"},
			_jsii_.MemberProperty{JsiiProperty: "targetOriginIdInput", GoGetter: "TargetOriginIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroups", GoGetter: "TrustedKeyGroups"},
			_jsii_.MemberProperty{JsiiProperty: "trustedKeyGroupsInput", GoGetter: "TrustedKeyGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSigners", GoGetter: "TrustedSigners"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSignersInput", GoGetter: "TrustedSignersInput"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicy", GoGetter: "ViewerProtocolPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "viewerProtocolPolicyInput", GoGetter: "ViewerProtocolPolicyInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionDefaultCacheBehaviorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionLoggingConfig",
		reflect.TypeOf((*CloudfrontDistributionLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionLoggingConfigOutputReference",
		reflect.TypeOf((*CloudfrontDistributionLoggingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeCookies", GoGetter: "IncludeCookies"},
			_jsii_.MemberProperty{JsiiProperty: "includeCookiesInput", GoGetter: "IncludeCookiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "prefix", GoGetter: "Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "prefixInput", GoGetter: "PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeCookies", GoMethod: "ResetIncludeCookies"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefix", GoMethod: "ResetPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehavior",
		reflect.TypeOf((*CloudfrontDistributionOrderedCacheBehavior)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorForwardedValues",
		reflect.TypeOf((*CloudfrontDistributionOrderedCacheBehaviorForwardedValues)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies",
		reflect.TypeOf((*CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookies)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "forward", GoGetter: "Forward"},
			_jsii_.MemberProperty{JsiiProperty: "forwardInput", GoGetter: "ForwardInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhitelistedNames", GoMethod: "ResetWhitelistedNames"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNames", GoGetter: "WhitelistedNames"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistedNamesInput", GoGetter: "WhitelistedNamesInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesCookiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference",
		reflect.TypeOf((*CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesInput", GoGetter: "CookiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putCookies", GoMethod: "PutCookies"},
			_jsii_.MemberProperty{JsiiProperty: "queryString", GoGetter: "QueryString"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeys", GoGetter: "QueryStringCacheKeys"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringCacheKeysInput", GoGetter: "QueryStringCacheKeysInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringInput", GoGetter: "QueryStringInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStringCacheKeys", GoMethod: "ResetQueryStringCacheKeys"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionOrderedCacheBehaviorForwardedValuesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorFunctionAssociation",
		reflect.TypeOf((*CloudfrontDistributionOrderedCacheBehaviorFunctionAssociation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociation",
		reflect.TypeOf((*CloudfrontDistributionOrderedCacheBehaviorLambdaFunctionAssociation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOrigin",
		reflect.TypeOf((*CloudfrontDistributionOrigin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginCustomHeader",
		reflect.TypeOf((*CloudfrontDistributionOriginCustomHeader)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginCustomOriginConfig",
		reflect.TypeOf((*CloudfrontDistributionOriginCustomOriginConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginCustomOriginConfigOutputReference",
		reflect.TypeOf((*CloudfrontDistributionOriginCustomOriginConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpPort", GoGetter: "HttpPort"},
			_jsii_.MemberProperty{JsiiProperty: "httpPortInput", GoGetter: "HttpPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpsPort", GoGetter: "HttpsPort"},
			_jsii_.MemberProperty{JsiiProperty: "httpsPortInput", GoGetter: "HttpsPortInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "originKeepaliveTimeout", GoGetter: "OriginKeepaliveTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "originKeepaliveTimeoutInput", GoGetter: "OriginKeepaliveTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "originProtocolPolicy", GoGetter: "OriginProtocolPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "originProtocolPolicyInput", GoGetter: "OriginProtocolPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "originReadTimeout", GoGetter: "OriginReadTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "originReadTimeoutInput", GoGetter: "OriginReadTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "originSslProtocols", GoGetter: "OriginSslProtocols"},
			_jsii_.MemberProperty{JsiiProperty: "originSslProtocolsInput", GoGetter: "OriginSslProtocolsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginKeepaliveTimeout", GoMethod: "ResetOriginKeepaliveTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetOriginReadTimeout", GoMethod: "ResetOriginReadTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionOriginCustomOriginConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginGroup",
		reflect.TypeOf((*CloudfrontDistributionOriginGroup)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginGroupFailoverCriteria",
		reflect.TypeOf((*CloudfrontDistributionOriginGroupFailoverCriteria)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference",
		reflect.TypeOf((*CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "statusCodes", GoGetter: "StatusCodes"},
			_jsii_.MemberProperty{JsiiProperty: "statusCodesInput", GoGetter: "StatusCodesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionOriginGroupFailoverCriteriaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginGroupMember",
		reflect.TypeOf((*CloudfrontDistributionOriginGroupMember)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginOriginShield",
		reflect.TypeOf((*CloudfrontDistributionOriginOriginShield)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginOriginShieldOutputReference",
		reflect.TypeOf((*CloudfrontDistributionOriginOriginShieldOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "originShieldRegion", GoGetter: "OriginShieldRegion"},
			_jsii_.MemberProperty{JsiiProperty: "originShieldRegionInput", GoGetter: "OriginShieldRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionOriginOriginShieldOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginS3OriginConfig",
		reflect.TypeOf((*CloudfrontDistributionOriginS3OriginConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionOriginS3OriginConfigOutputReference",
		reflect.TypeOf((*CloudfrontDistributionOriginS3OriginConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentity", GoGetter: "OriginAccessIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "originAccessIdentityInput", GoGetter: "OriginAccessIdentityInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionOriginS3OriginConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionRestrictions",
		reflect.TypeOf((*CloudfrontDistributionRestrictions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionRestrictionsGeoRestriction",
		reflect.TypeOf((*CloudfrontDistributionRestrictionsGeoRestriction)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionRestrictionsGeoRestrictionOutputReference",
		reflect.TypeOf((*CloudfrontDistributionRestrictionsGeoRestrictionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "locations", GoGetter: "Locations"},
			_jsii_.MemberProperty{JsiiProperty: "locationsInput", GoGetter: "LocationsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocations", GoMethod: "ResetLocations"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionType", GoGetter: "RestrictionType"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionTypeInput", GoGetter: "RestrictionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionRestrictionsGeoRestrictionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionRestrictionsOutputReference",
		reflect.TypeOf((*CloudfrontDistributionRestrictionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "geoRestriction", GoGetter: "GeoRestriction"},
			_jsii_.MemberProperty{JsiiProperty: "geoRestrictionInput", GoGetter: "GeoRestrictionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putGeoRestriction", GoMethod: "PutGeoRestriction"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionRestrictionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedKeyGroups",
		reflect.TypeOf((*CloudfrontDistributionTrustedKeyGroups)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionTrustedKeyGroups{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedKeyGroupsItems",
		reflect.TypeOf((*CloudfrontDistributionTrustedKeyGroupsItems)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyGroupId", GoGetter: "KeyGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "keyPairIds", GoGetter: "KeyPairIds"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionTrustedKeyGroupsItems{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedSigners",
		reflect.TypeOf((*CloudfrontDistributionTrustedSigners)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionTrustedSigners{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionTrustedSignersItems",
		reflect.TypeOf((*CloudfrontDistributionTrustedSignersItems)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "awsAccountNumber", GoGetter: "AwsAccountNumber"},
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyPairIds", GoGetter: "KeyPairIds"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionTrustedSignersItems{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontDistributionViewerCertificate",
		reflect.TypeOf((*CloudfrontDistributionViewerCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontDistributionViewerCertificateOutputReference",
		reflect.TypeOf((*CloudfrontDistributionViewerCertificateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acmCertificateArn", GoGetter: "AcmCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "acmCertificateArnInput", GoGetter: "AcmCertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontDefaultCertificate", GoGetter: "CloudfrontDefaultCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontDefaultCertificateInput", GoGetter: "CloudfrontDefaultCertificateInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamCertificateId", GoGetter: "IamCertificateId"},
			_jsii_.MemberProperty{JsiiProperty: "iamCertificateIdInput", GoGetter: "IamCertificateIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "minimumProtocolVersion", GoGetter: "MinimumProtocolVersion"},
			_jsii_.MemberProperty{JsiiProperty: "minimumProtocolVersionInput", GoGetter: "MinimumProtocolVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcmCertificateArn", GoMethod: "ResetAcmCertificateArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudfrontDefaultCertificate", GoMethod: "ResetCloudfrontDefaultCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetIamCertificateId", GoMethod: "ResetIamCertificateId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumProtocolVersion", GoMethod: "ResetMinimumProtocolVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSslSupportMethod", GoMethod: "ResetSslSupportMethod"},
			_jsii_.MemberProperty{JsiiProperty: "sslSupportMethod", GoGetter: "SslSupportMethod"},
			_jsii_.MemberProperty{JsiiProperty: "sslSupportMethodInput", GoGetter: "SslSupportMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontDistributionViewerCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontFunction",
		reflect.TypeOf((*CloudfrontFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "code", GoGetter: "Code"},
			_jsii_.MemberProperty{JsiiProperty: "codeInput", GoGetter: "CodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "publish", GoGetter: "Publish"},
			_jsii_.MemberProperty{JsiiProperty: "publishInput", GoGetter: "PublishInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublish", GoMethod: "ResetPublish"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberProperty{JsiiProperty: "runtimeInput", GoGetter: "RuntimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontFunction{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontFunctionConfig",
		reflect.TypeOf((*CloudfrontFunctionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontKeyGroup",
		reflect.TypeOf((*CloudfrontKeyGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontKeyGroup{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontKeyGroupConfig",
		reflect.TypeOf((*CloudfrontKeyGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscription",
		reflect.TypeOf((*CloudfrontMonitoringSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "distributionId", GoGetter: "DistributionId"},
			_jsii_.MemberProperty{JsiiProperty: "distributionIdInput", GoGetter: "DistributionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringSubscription", GoGetter: "MonitoringSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringSubscriptionInput", GoGetter: "MonitoringSubscriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putMonitoringSubscription", GoMethod: "PutMonitoringSubscription"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontMonitoringSubscription{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscriptionConfig",
		reflect.TypeOf((*CloudfrontMonitoringSubscriptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscriptionMonitoringSubscription",
		reflect.TypeOf((*CloudfrontMonitoringSubscriptionMonitoringSubscription)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference",
		reflect.TypeOf((*CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putRealtimeMetricsSubscriptionConfig", GoMethod: "PutRealtimeMetricsSubscriptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeMetricsSubscriptionConfig", GoGetter: "RealtimeMetricsSubscriptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeMetricsSubscriptionConfigInput", GoGetter: "RealtimeMetricsSubscriptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig",
		reflect.TypeOf((*CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference",
		reflect.TypeOf((*CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeMetricsSubscriptionStatus", GoGetter: "RealtimeMetricsSubscriptionStatus"},
			_jsii_.MemberProperty{JsiiProperty: "realtimeMetricsSubscriptionStatusInput", GoGetter: "RealtimeMetricsSubscriptionStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontMonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontOriginAccessIdentity",
		reflect.TypeOf((*CloudfrontOriginAccessIdentity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "callerReference", GoGetter: "CallerReference"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cloudfrontAccessIdentityPath", GoGetter: "CloudfrontAccessIdentityPath"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "iamArn", GoGetter: "IamArn"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "s3CanonicalUserId", GoGetter: "S3CanonicalUserId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontOriginAccessIdentity{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontOriginAccessIdentityConfig",
		reflect.TypeOf((*CloudfrontOriginAccessIdentityConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicy",
		reflect.TypeOf((*CloudfrontOriginRequestPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesConfig", GoGetter: "CookiesConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesConfigInput", GoGetter: "CookiesConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "etagInput", GoGetter: "EtagInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headersConfig", GoGetter: "HeadersConfig"},
			_jsii_.MemberProperty{JsiiProperty: "headersConfigInput", GoGetter: "HeadersConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putCookiesConfig", GoMethod: "PutCookiesConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putHeadersConfig", GoMethod: "PutHeadersConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryStringsConfig", GoMethod: "PutQueryStringsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringsConfig", GoGetter: "QueryStringsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringsConfigInput", GoGetter: "QueryStringsConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetEtag", GoMethod: "ResetEtag"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontOriginRequestPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyConfig",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyCookiesConfig",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyCookiesConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyCookiesConfigCookies",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyCookiesConfigCookies)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigCookiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyCookiesConfigOutputReference",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyCookiesConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cookieBehavior", GoGetter: "CookieBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "cookieBehaviorInput", GoGetter: "CookieBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesInput", GoGetter: "CookiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putCookies", GoMethod: "PutCookies"},
			_jsii_.MemberMethod{JsiiMethod: "resetCookies", GoMethod: "ResetCookies"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontOriginRequestPolicyCookiesConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyHeadersConfig",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyHeadersConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyHeadersConfigHeaders",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyHeadersConfigHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyHeadersConfigOutputReference",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyHeadersConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headerBehavior", GoGetter: "HeaderBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "headerBehaviorInput", GoGetter: "HeaderBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "headersInput", GoGetter: "HeadersInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putHeaders", GoMethod: "PutHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderBehavior", GoMethod: "ResetHeaderBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaders", GoMethod: "ResetHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontOriginRequestPolicyHeadersConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyQueryStringsConfig",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyQueryStringsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryStrings", GoMethod: "PutQueryStrings"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringBehavior", GoGetter: "QueryStringBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringBehaviorInput", GoGetter: "QueryStringBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryStrings", GoGetter: "QueryStrings"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringsInput", GoGetter: "QueryStringsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryStrings", GoMethod: "ResetQueryStrings"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference",
		reflect.TypeOf((*CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontOriginRequestPolicyQueryStringsConfigQueryStringsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontPublicKey",
		reflect.TypeOf((*CloudfrontPublicKey)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "callerReference", GoGetter: "CallerReference"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "encodedKey", GoGetter: "EncodedKey"},
			_jsii_.MemberProperty{JsiiProperty: "encodedKeyInput", GoGetter: "EncodedKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefix", GoGetter: "NamePrefix"},
			_jsii_.MemberProperty{JsiiProperty: "namePrefixInput", GoGetter: "NamePrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamePrefix", GoMethod: "ResetNamePrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontPublicKey{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontPublicKeyConfig",
		reflect.TypeOf((*CloudfrontPublicKeyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfig",
		reflect.TypeOf((*CloudfrontRealtimeLogConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "endpointInput", GoGetter: "EndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "fields", GoGetter: "Fields"},
			_jsii_.MemberProperty{JsiiProperty: "fieldsInput", GoGetter: "FieldsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putEndpoint", GoMethod: "PutEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "samplingRate", GoGetter: "SamplingRate"},
			_jsii_.MemberProperty{JsiiProperty: "samplingRateInput", GoGetter: "SamplingRateInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontRealtimeLogConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfigConfig",
		reflect.TypeOf((*CloudfrontRealtimeLogConfigConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfigEndpoint",
		reflect.TypeOf((*CloudfrontRealtimeLogConfigEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig",
		reflect.TypeOf((*CloudfrontRealtimeLogConfigEndpointKinesisStreamConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference",
		reflect.TypeOf((*CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "roleArnInput", GoGetter: "RoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "streamArn", GoGetter: "StreamArn"},
			_jsii_.MemberProperty{JsiiProperty: "streamArnInput", GoGetter: "StreamArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontRealtimeLogConfigEndpointKinesisStreamConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontRealtimeLogConfigEndpointOutputReference",
		reflect.TypeOf((*CloudfrontRealtimeLogConfigEndpointOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStreamConfig", GoGetter: "KinesisStreamConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisStreamConfigInput", GoGetter: "KinesisStreamConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKinesisStreamConfig", GoMethod: "PutKinesisStreamConfig"},
			_jsii_.MemberProperty{JsiiProperty: "streamType", GoGetter: "StreamType"},
			_jsii_.MemberProperty{JsiiProperty: "streamTypeInput", GoGetter: "StreamTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontRealtimeLogConfigEndpointOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicy",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "commentInput", GoGetter: "CommentInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "corsConfig", GoGetter: "CorsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "corsConfigInput", GoGetter: "CorsConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customHeadersConfig", GoGetter: "CustomHeadersConfig"},
			_jsii_.MemberProperty{JsiiProperty: "customHeadersConfigInput", GoGetter: "CustomHeadersConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "etagInput", GoGetter: "EtagInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putCorsConfig", GoMethod: "PutCorsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomHeadersConfig", GoMethod: "PutCustomHeadersConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSecurityHeadersConfig", GoMethod: "PutSecurityHeadersConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetComment", GoMethod: "ResetComment"},
			_jsii_.MemberMethod{JsiiMethod: "resetCorsConfig", GoMethod: "ResetCorsConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomHeadersConfig", GoMethod: "ResetCustomHeadersConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetEtag", GoMethod: "ResetEtag"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityHeadersConfig", GoMethod: "ResetSecurityHeadersConfig"},
			_jsii_.MemberProperty{JsiiProperty: "securityHeadersConfig", GoGetter: "SecurityHeadersConfig"},
			_jsii_.MemberProperty{JsiiProperty: "securityHeadersConfigInput", GoGetter: "SecurityHeadersConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyConfig",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfig",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethodsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOriginsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCorsConfigOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCorsConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowCredentials", GoGetter: "AccessControlAllowCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowCredentialsInput", GoGetter: "AccessControlAllowCredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowHeaders", GoGetter: "AccessControlAllowHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowHeadersInput", GoGetter: "AccessControlAllowHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowMethods", GoGetter: "AccessControlAllowMethods"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowMethodsInput", GoGetter: "AccessControlAllowMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowOrigins", GoGetter: "AccessControlAllowOrigins"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowOriginsInput", GoGetter: "AccessControlAllowOriginsInput"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlExposeHeaders", GoGetter: "AccessControlExposeHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlExposeHeadersInput", GoGetter: "AccessControlExposeHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlMaxAgeSec", GoGetter: "AccessControlMaxAgeSec"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlMaxAgeSecInput", GoGetter: "AccessControlMaxAgeSecInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "originOverride", GoGetter: "OriginOverride"},
			_jsii_.MemberProperty{JsiiProperty: "originOverrideInput", GoGetter: "OriginOverrideInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccessControlAllowHeaders", GoMethod: "PutAccessControlAllowHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "putAccessControlAllowMethods", GoMethod: "PutAccessControlAllowMethods"},
			_jsii_.MemberMethod{JsiiMethod: "putAccessControlAllowOrigins", GoMethod: "PutAccessControlAllowOrigins"},
			_jsii_.MemberMethod{JsiiMethod: "putAccessControlExposeHeaders", GoMethod: "PutAccessControlExposeHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessControlExposeHeaders", GoMethod: "ResetAccessControlExposeHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessControlMaxAgeSec", GoMethod: "ResetAccessControlMaxAgeSec"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicyCorsConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCustomHeadersConfig",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCustomHeadersConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCustomHeadersConfigItems",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCustomHeadersConfigItems)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "itemsInput", GoGetter: "ItemsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetItems", GoMethod: "ResetItems"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicyCustomHeadersConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfig",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "contentSecurityPolicy", GoGetter: "ContentSecurityPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "contentSecurityPolicyInput", GoGetter: "ContentSecurityPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "overrideInput", GoGetter: "OverrideInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "overrideInput", GoGetter: "OverrideInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "frameOption", GoGetter: "FrameOption"},
			_jsii_.MemberProperty{JsiiProperty: "frameOptionInput", GoGetter: "FrameOptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "overrideInput", GoGetter: "OverrideInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "contentSecurityPolicy", GoGetter: "ContentSecurityPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "contentSecurityPolicyInput", GoGetter: "ContentSecurityPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentTypeOptions", GoGetter: "ContentTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "contentTypeOptionsInput", GoGetter: "ContentTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "frameOptions", GoGetter: "FrameOptions"},
			_jsii_.MemberProperty{JsiiProperty: "frameOptionsInput", GoGetter: "FrameOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "putContentSecurityPolicy", GoMethod: "PutContentSecurityPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putContentTypeOptions", GoMethod: "PutContentTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putFrameOptions", GoMethod: "PutFrameOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putReferrerPolicy", GoMethod: "PutReferrerPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putStrictTransportSecurity", GoMethod: "PutStrictTransportSecurity"},
			_jsii_.MemberMethod{JsiiMethod: "putXssProtection", GoMethod: "PutXssProtection"},
			_jsii_.MemberProperty{JsiiProperty: "referrerPolicy", GoGetter: "ReferrerPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "referrerPolicyInput", GoGetter: "ReferrerPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentSecurityPolicy", GoMethod: "ResetContentSecurityPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentTypeOptions", GoMethod: "ResetContentTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetFrameOptions", GoMethod: "ResetFrameOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetReferrerPolicy", GoMethod: "ResetReferrerPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetStrictTransportSecurity", GoMethod: "ResetStrictTransportSecurity"},
			_jsii_.MemberMethod{JsiiMethod: "resetXssProtection", GoMethod: "ResetXssProtection"},
			_jsii_.MemberProperty{JsiiProperty: "strictTransportSecurity", GoGetter: "StrictTransportSecurity"},
			_jsii_.MemberProperty{JsiiProperty: "strictTransportSecurityInput", GoGetter: "StrictTransportSecurityInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "xssProtection", GoGetter: "XssProtection"},
			_jsii_.MemberProperty{JsiiProperty: "xssProtectionInput", GoGetter: "XssProtectionInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "overrideInput", GoGetter: "OverrideInput"},
			_jsii_.MemberProperty{JsiiProperty: "referrerPolicy", GoGetter: "ReferrerPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "referrerPolicyInput", GoGetter: "ReferrerPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessControlMaxAgeSec", GoGetter: "AccessControlMaxAgeSec"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlMaxAgeSecInput", GoGetter: "AccessControlMaxAgeSecInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeSubdomains", GoGetter: "IncludeSubdomains"},
			_jsii_.MemberProperty{JsiiProperty: "includeSubdomainsInput", GoGetter: "IncludeSubdomainsInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "overrideInput", GoGetter: "OverrideInput"},
			_jsii_.MemberProperty{JsiiProperty: "preload", GoGetter: "Preload"},
			_jsii_.MemberProperty{JsiiProperty: "preloadInput", GoGetter: "PreloadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeSubdomains", GoMethod: "ResetIncludeSubdomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreload", GoMethod: "ResetPreload"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference",
		reflect.TypeOf((*CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "modeBlock", GoGetter: "ModeBlock"},
			_jsii_.MemberProperty{JsiiProperty: "modeBlockInput", GoGetter: "ModeBlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "overrideInput", GoGetter: "OverrideInput"},
			_jsii_.MemberProperty{JsiiProperty: "protection", GoGetter: "Protection"},
			_jsii_.MemberProperty{JsiiProperty: "protectionInput", GoGetter: "ProtectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "reportUri", GoGetter: "ReportUri"},
			_jsii_.MemberProperty{JsiiProperty: "reportUriInput", GoGetter: "ReportUriInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetModeBlock", GoMethod: "ResetModeBlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetReportUri", GoMethod: "ResetReportUri"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtectionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicy",
		reflect.TypeOf((*DataAwsCloudfrontCachePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTtl", GoGetter: "DefaultTtl"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maxTtl", GoGetter: "MaxTtl"},
			_jsii_.MemberProperty{JsiiProperty: "minTtl", GoGetter: "MinTtl"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "parametersInCacheKeyAndForwardedToOrigin", GoMethod: "ParametersInCacheKeyAndForwardedToOrigin"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontCachePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyConfig",
		reflect.TypeOf((*DataAwsCloudfrontCachePolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin",
		reflect.TypeOf((*DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "cookiesConfig", GoGetter: "CookiesConfig"},
			_jsii_.MemberProperty{JsiiProperty: "enableAcceptEncodingBrotli", GoGetter: "EnableAcceptEncodingBrotli"},
			_jsii_.MemberProperty{JsiiProperty: "enableAcceptEncodingGzip", GoGetter: "EnableAcceptEncodingGzip"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headersConfig", GoGetter: "HeadersConfig"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringsConfig", GoGetter: "QueryStringsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOrigin{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig",
		reflect.TypeOf((*DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "cookieBehavior", GoGetter: "CookieBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies",
		reflect.TypeOf((*DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookies{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig",
		reflect.TypeOf((*DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headerBehavior", GoGetter: "HeaderBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders",
		reflect.TypeOf((*DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeaders{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig",
		reflect.TypeOf((*DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringBehavior", GoGetter: "QueryStringBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "queryStrings", GoGetter: "QueryStrings"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings",
		reflect.TypeOf((*DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontCachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStrings{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontDistribution",
		reflect.TypeOf((*DataAwsCloudfrontDistribution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hostedZoneId", GoGetter: "HostedZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "inProgressValidationBatches", GoGetter: "InProgressValidationBatches"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastModifiedTime", GoGetter: "LastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
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
			j := jsiiProxy_DataAwsCloudfrontDistribution{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontDistributionConfig",
		reflect.TypeOf((*DataAwsCloudfrontDistributionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontFunction",
		reflect.TypeOf((*DataAwsCloudfrontFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "code", GoGetter: "Code"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lastModifiedTime", GoGetter: "LastModifiedTime"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "runtime", GoGetter: "Runtime"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "stageInput", GoGetter: "StageInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontFunction{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontFunctionConfig",
		reflect.TypeOf((*DataAwsCloudfrontFunctionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontLogDeliveryCanonicalUserId",
		reflect.TypeOf((*DataAwsCloudfrontLogDeliveryCanonicalUserId)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontLogDeliveryCanonicalUserId{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontLogDeliveryCanonicalUserIdConfig",
		reflect.TypeOf((*DataAwsCloudfrontLogDeliveryCanonicalUserIdConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicy",
		reflect.TypeOf((*DataAwsCloudfrontOriginRequestPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "cookiesConfig", GoMethod: "CookiesConfig"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "headersConfig", GoMethod: "HeadersConfig"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "queryStringsConfig", GoMethod: "QueryStringsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyConfig",
		reflect.TypeOf((*DataAwsCloudfrontOriginRequestPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyCookiesConfig",
		reflect.TypeOf((*DataAwsCloudfrontOriginRequestPolicyCookiesConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "cookieBehavior", GoGetter: "CookieBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "cookies", GoGetter: "Cookies"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies",
		reflect.TypeOf((*DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyCookiesConfigCookies{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyHeadersConfig",
		reflect.TypeOf((*DataAwsCloudfrontOriginRequestPolicyHeadersConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "headerBehavior", GoGetter: "HeaderBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders",
		reflect.TypeOf((*DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyHeadersConfigHeaders{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig",
		reflect.TypeOf((*DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "queryStringBehavior", GoGetter: "QueryStringBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "queryStrings", GoGetter: "QueryStrings"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings",
		reflect.TypeOf((*DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontOriginRequestPolicyQueryStringsConfigQueryStrings{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicy",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "corsConfig", GoMethod: "CorsConfig"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberMethod{JsiiMethod: "customHeadersConfig", GoMethod: "CustomHeadersConfig"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "securityHeadersConfig", GoMethod: "SecurityHeadersConfig"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyConfig",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfig",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicyCorsConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowCredentials", GoGetter: "AccessControlAllowCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowHeaders", GoGetter: "AccessControlAllowHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowMethods", GoGetter: "AccessControlAllowMethods"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlAllowOrigins", GoGetter: "AccessControlAllowOrigins"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlExposeHeaders", GoGetter: "AccessControlExposeHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "accessControlMaxAgeSec", GoGetter: "AccessControlMaxAgeSec"},
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "originOverride", GoGetter: "OriginOverride"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowHeaders{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowMethods{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlAllowOrigins{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "items", GoGetter: "Items"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCorsConfigAccessControlExposeHeaders{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "header", GoGetter: "Header"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicyCustomHeadersConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "contentSecurityPolicy", GoGetter: "ContentSecurityPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "contentTypeOptions", GoGetter: "ContentTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "frameOptions", GoGetter: "FrameOptions"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "referrerPolicy", GoGetter: "ReferrerPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "strictTransportSecurity", GoGetter: "StrictTransportSecurity"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "xssProtection", GoGetter: "XssProtection"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "contentSecurityPolicy", GoGetter: "ContentSecurityPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentSecurityPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigContentTypeOptions{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberProperty{JsiiProperty: "frameOption", GoGetter: "FrameOption"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigFrameOptions{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "referrerPolicy", GoGetter: "ReferrerPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigReferrerPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessControlMaxAgeSec", GoGetter: "AccessControlMaxAgeSec"},
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeSubdomains", GoGetter: "IncludeSubdomains"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "preload", GoGetter: "Preload"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.CloudFront.DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection",
		reflect.TypeOf((*DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexComputedListIndex", GoGetter: "ComplexComputedListIndex"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "modeBlock", GoGetter: "ModeBlock"},
			_jsii_.MemberProperty{JsiiProperty: "override", GoGetter: "Override"},
			_jsii_.MemberProperty{JsiiProperty: "protection", GoGetter: "Protection"},
			_jsii_.MemberProperty{JsiiProperty: "reportUri", GoGetter: "ReportUri"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_DataAwsCloudfrontResponseHeadersPolicySecurityHeadersConfigXssProtection{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexComputedList)
			return &j
		},
	)
}
