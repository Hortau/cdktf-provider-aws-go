package elastictranscoder

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipeline",
		reflect.TypeOf((*ElastictranscoderPipeline)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "awsKmsKeyArn", GoGetter: "AwsKmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "awsKmsKeyArnInput", GoGetter: "AwsKmsKeyArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "contentConfig", GoGetter: "ContentConfig"},
			_jsii_.MemberProperty{JsiiProperty: "contentConfigInput", GoGetter: "ContentConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentConfigPermissions", GoGetter: "ContentConfigPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "contentConfigPermissionsInput", GoGetter: "ContentConfigPermissionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputBucket", GoGetter: "InputBucket"},
			_jsii_.MemberProperty{JsiiProperty: "inputBucketInput", GoGetter: "InputBucketInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notifications", GoGetter: "Notifications"},
			_jsii_.MemberProperty{JsiiProperty: "notificationsInput", GoGetter: "NotificationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputBucket", GoGetter: "OutputBucket"},
			_jsii_.MemberProperty{JsiiProperty: "outputBucketInput", GoGetter: "OutputBucketInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putContentConfig", GoMethod: "PutContentConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putNotifications", GoMethod: "PutNotifications"},
			_jsii_.MemberMethod{JsiiMethod: "putThumbnailConfig", GoMethod: "PutThumbnailConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAwsKmsKeyArn", GoMethod: "ResetAwsKmsKeyArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentConfig", GoMethod: "ResetContentConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentConfigPermissions", GoMethod: "ResetContentConfigPermissions"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNotifications", GoMethod: "ResetNotifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputBucket", GoMethod: "ResetOutputBucket"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetThumbnailConfig", GoMethod: "ResetThumbnailConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetThumbnailConfigPermissions", GoMethod: "ResetThumbnailConfigPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "roleInput", GoGetter: "RoleInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "thumbnailConfig", GoGetter: "ThumbnailConfig"},
			_jsii_.MemberProperty{JsiiProperty: "thumbnailConfigInput", GoGetter: "ThumbnailConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "thumbnailConfigPermissions", GoGetter: "ThumbnailConfigPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "thumbnailConfigPermissionsInput", GoGetter: "ThumbnailConfigPermissionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_ElastictranscoderPipeline{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineConfig",
		reflect.TypeOf((*ElastictranscoderPipelineConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineContentConfig",
		reflect.TypeOf((*ElastictranscoderPipelineContentConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineContentConfigOutputReference",
		reflect.TypeOf((*ElastictranscoderPipelineContentConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucket", GoMethod: "ResetBucket"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageClass", GoMethod: "ResetStorageClass"},
			_jsii_.MemberProperty{JsiiProperty: "storageClass", GoGetter: "StorageClass"},
			_jsii_.MemberProperty{JsiiProperty: "storageClassInput", GoGetter: "StorageClassInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_ElastictranscoderPipelineContentConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineContentConfigPermissions",
		reflect.TypeOf((*ElastictranscoderPipelineContentConfigPermissions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineNotifications",
		reflect.TypeOf((*ElastictranscoderPipelineNotifications)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineNotificationsOutputReference",
		reflect.TypeOf((*ElastictranscoderPipelineNotificationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "completed", GoGetter: "Completed"},
			_jsii_.MemberProperty{JsiiProperty: "completedInput", GoGetter: "CompletedInput"},
			_jsii_.MemberProperty{JsiiProperty: "error", GoGetter: "Error"},
			_jsii_.MemberProperty{JsiiProperty: "errorInput", GoGetter: "ErrorInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "progressing", GoGetter: "Progressing"},
			_jsii_.MemberProperty{JsiiProperty: "progressingInput", GoGetter: "ProgressingInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompleted", GoMethod: "ResetCompleted"},
			_jsii_.MemberMethod{JsiiMethod: "resetError", GoMethod: "ResetError"},
			_jsii_.MemberMethod{JsiiMethod: "resetProgressing", GoMethod: "ResetProgressing"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarning", GoMethod: "ResetWarning"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "warning", GoGetter: "Warning"},
			_jsii_.MemberProperty{JsiiProperty: "warningInput", GoGetter: "WarningInput"},
		},
		func() interface{} {
			j := jsiiProxy_ElastictranscoderPipelineNotificationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineThumbnailConfig",
		reflect.TypeOf((*ElastictranscoderPipelineThumbnailConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineThumbnailConfigOutputReference",
		reflect.TypeOf((*ElastictranscoderPipelineThumbnailConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetBucket", GoMethod: "ResetBucket"},
			_jsii_.MemberMethod{JsiiMethod: "resetStorageClass", GoMethod: "ResetStorageClass"},
			_jsii_.MemberProperty{JsiiProperty: "storageClass", GoGetter: "StorageClass"},
			_jsii_.MemberProperty{JsiiProperty: "storageClassInput", GoGetter: "StorageClassInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_ElastictranscoderPipelineThumbnailConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPipelineThumbnailConfigPermissions",
		reflect.TypeOf((*ElastictranscoderPipelineThumbnailConfigPermissions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPreset",
		reflect.TypeOf((*ElastictranscoderPreset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "audio", GoGetter: "Audio"},
			_jsii_.MemberProperty{JsiiProperty: "audioCodecOptions", GoGetter: "AudioCodecOptions"},
			_jsii_.MemberProperty{JsiiProperty: "audioCodecOptionsInput", GoGetter: "AudioCodecOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "audioInput", GoGetter: "AudioInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerInput", GoGetter: "ContainerInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putAudio", GoMethod: "PutAudio"},
			_jsii_.MemberMethod{JsiiMethod: "putAudioCodecOptions", GoMethod: "PutAudioCodecOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putThumbnails", GoMethod: "PutThumbnails"},
			_jsii_.MemberMethod{JsiiMethod: "putVideo", GoMethod: "PutVideo"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudio", GoMethod: "ResetAudio"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudioCodecOptions", GoMethod: "ResetAudioCodecOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetThumbnails", GoMethod: "ResetThumbnails"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetVideo", GoMethod: "ResetVideo"},
			_jsii_.MemberMethod{JsiiMethod: "resetVideoCodecOptions", GoMethod: "ResetVideoCodecOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetVideoWatermarks", GoMethod: "ResetVideoWatermarks"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "thumbnails", GoGetter: "Thumbnails"},
			_jsii_.MemberProperty{JsiiProperty: "thumbnailsInput", GoGetter: "ThumbnailsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "video", GoGetter: "Video"},
			_jsii_.MemberProperty{JsiiProperty: "videoCodecOptions", GoGetter: "VideoCodecOptions"},
			_jsii_.MemberProperty{JsiiProperty: "videoCodecOptionsInput", GoGetter: "VideoCodecOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "videoInput", GoGetter: "VideoInput"},
			_jsii_.MemberProperty{JsiiProperty: "videoWatermarks", GoGetter: "VideoWatermarks"},
			_jsii_.MemberProperty{JsiiProperty: "videoWatermarksInput", GoGetter: "VideoWatermarksInput"},
		},
		func() interface{} {
			j := jsiiProxy_ElastictranscoderPreset{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetAudio",
		reflect.TypeOf((*ElastictranscoderPresetAudio)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetAudioCodecOptions",
		reflect.TypeOf((*ElastictranscoderPresetAudioCodecOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetAudioCodecOptionsOutputReference",
		reflect.TypeOf((*ElastictranscoderPresetAudioCodecOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bitDepth", GoGetter: "BitDepth"},
			_jsii_.MemberProperty{JsiiProperty: "bitDepthInput", GoGetter: "BitDepthInput"},
			_jsii_.MemberProperty{JsiiProperty: "bitOrder", GoGetter: "BitOrder"},
			_jsii_.MemberProperty{JsiiProperty: "bitOrderInput", GoGetter: "BitOrderInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "profile", GoGetter: "Profile"},
			_jsii_.MemberProperty{JsiiProperty: "profileInput", GoGetter: "ProfileInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBitDepth", GoMethod: "ResetBitDepth"},
			_jsii_.MemberMethod{JsiiMethod: "resetBitOrder", GoMethod: "ResetBitOrder"},
			_jsii_.MemberMethod{JsiiMethod: "resetProfile", GoMethod: "ResetProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetSigned", GoMethod: "ResetSigned"},
			_jsii_.MemberProperty{JsiiProperty: "signed", GoGetter: "Signed"},
			_jsii_.MemberProperty{JsiiProperty: "signedInput", GoGetter: "SignedInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_ElastictranscoderPresetAudioCodecOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetAudioOutputReference",
		reflect.TypeOf((*ElastictranscoderPresetAudioOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "audioPackingMode", GoGetter: "AudioPackingMode"},
			_jsii_.MemberProperty{JsiiProperty: "audioPackingModeInput", GoGetter: "AudioPackingModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "bitRate", GoGetter: "BitRate"},
			_jsii_.MemberProperty{JsiiProperty: "bitRateInput", GoGetter: "BitRateInput"},
			_jsii_.MemberProperty{JsiiProperty: "channels", GoGetter: "Channels"},
			_jsii_.MemberProperty{JsiiProperty: "channelsInput", GoGetter: "ChannelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "codec", GoGetter: "Codec"},
			_jsii_.MemberProperty{JsiiProperty: "codecInput", GoGetter: "CodecInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudioPackingMode", GoMethod: "ResetAudioPackingMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetBitRate", GoMethod: "ResetBitRate"},
			_jsii_.MemberMethod{JsiiMethod: "resetChannels", GoMethod: "ResetChannels"},
			_jsii_.MemberMethod{JsiiMethod: "resetCodec", GoMethod: "ResetCodec"},
			_jsii_.MemberMethod{JsiiMethod: "resetSampleRate", GoMethod: "ResetSampleRate"},
			_jsii_.MemberProperty{JsiiProperty: "sampleRate", GoGetter: "SampleRate"},
			_jsii_.MemberProperty{JsiiProperty: "sampleRateInput", GoGetter: "SampleRateInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_ElastictranscoderPresetAudioOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetConfig",
		reflect.TypeOf((*ElastictranscoderPresetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetThumbnails",
		reflect.TypeOf((*ElastictranscoderPresetThumbnails)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetThumbnailsOutputReference",
		reflect.TypeOf((*ElastictranscoderPresetThumbnailsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aspectRatio", GoGetter: "AspectRatio"},
			_jsii_.MemberProperty{JsiiProperty: "aspectRatioInput", GoGetter: "AspectRatioInput"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "formatInput", GoGetter: "FormatInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInput", GoGetter: "IntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "maxHeight", GoGetter: "MaxHeight"},
			_jsii_.MemberProperty{JsiiProperty: "maxHeightInput", GoGetter: "MaxHeightInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxWidth", GoGetter: "MaxWidth"},
			_jsii_.MemberProperty{JsiiProperty: "maxWidthInput", GoGetter: "MaxWidthInput"},
			_jsii_.MemberProperty{JsiiProperty: "paddingPolicy", GoGetter: "PaddingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "paddingPolicyInput", GoGetter: "PaddingPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAspectRatio", GoMethod: "ResetAspectRatio"},
			_jsii_.MemberMethod{JsiiMethod: "resetFormat", GoMethod: "ResetFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetInterval", GoMethod: "ResetInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxHeight", GoMethod: "ResetMaxHeight"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxWidth", GoMethod: "ResetMaxWidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetPaddingPolicy", GoMethod: "ResetPaddingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetResolution", GoMethod: "ResetResolution"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizingPolicy", GoMethod: "ResetSizingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "resolution", GoGetter: "Resolution"},
			_jsii_.MemberProperty{JsiiProperty: "resolutionInput", GoGetter: "ResolutionInput"},
			_jsii_.MemberProperty{JsiiProperty: "sizingPolicy", GoGetter: "SizingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "sizingPolicyInput", GoGetter: "SizingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_ElastictranscoderPresetThumbnailsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetVideo",
		reflect.TypeOf((*ElastictranscoderPresetVideo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetVideoOutputReference",
		reflect.TypeOf((*ElastictranscoderPresetVideoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aspectRatio", GoGetter: "AspectRatio"},
			_jsii_.MemberProperty{JsiiProperty: "aspectRatioInput", GoGetter: "AspectRatioInput"},
			_jsii_.MemberProperty{JsiiProperty: "bitRate", GoGetter: "BitRate"},
			_jsii_.MemberProperty{JsiiProperty: "bitRateInput", GoGetter: "BitRateInput"},
			_jsii_.MemberProperty{JsiiProperty: "codec", GoGetter: "Codec"},
			_jsii_.MemberProperty{JsiiProperty: "codecInput", GoGetter: "CodecInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayAspectRatio", GoGetter: "DisplayAspectRatio"},
			_jsii_.MemberProperty{JsiiProperty: "displayAspectRatioInput", GoGetter: "DisplayAspectRatioInput"},
			_jsii_.MemberProperty{JsiiProperty: "fixedGop", GoGetter: "FixedGop"},
			_jsii_.MemberProperty{JsiiProperty: "fixedGopInput", GoGetter: "FixedGopInput"},
			_jsii_.MemberProperty{JsiiProperty: "frameRate", GoGetter: "FrameRate"},
			_jsii_.MemberProperty{JsiiProperty: "frameRateInput", GoGetter: "FrameRateInput"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isSingleItem", GoGetter: "IsSingleItem"},
			_jsii_.MemberProperty{JsiiProperty: "keyframesMaxDist", GoGetter: "KeyframesMaxDist"},
			_jsii_.MemberProperty{JsiiProperty: "keyframesMaxDistInput", GoGetter: "KeyframesMaxDistInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxFrameRate", GoGetter: "MaxFrameRate"},
			_jsii_.MemberProperty{JsiiProperty: "maxFrameRateInput", GoGetter: "MaxFrameRateInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxHeight", GoGetter: "MaxHeight"},
			_jsii_.MemberProperty{JsiiProperty: "maxHeightInput", GoGetter: "MaxHeightInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxWidth", GoGetter: "MaxWidth"},
			_jsii_.MemberProperty{JsiiProperty: "maxWidthInput", GoGetter: "MaxWidthInput"},
			_jsii_.MemberProperty{JsiiProperty: "paddingPolicy", GoGetter: "PaddingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "paddingPolicyInput", GoGetter: "PaddingPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAspectRatio", GoMethod: "ResetAspectRatio"},
			_jsii_.MemberMethod{JsiiMethod: "resetBitRate", GoMethod: "ResetBitRate"},
			_jsii_.MemberMethod{JsiiMethod: "resetCodec", GoMethod: "ResetCodec"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayAspectRatio", GoMethod: "ResetDisplayAspectRatio"},
			_jsii_.MemberMethod{JsiiMethod: "resetFixedGop", GoMethod: "ResetFixedGop"},
			_jsii_.MemberMethod{JsiiMethod: "resetFrameRate", GoMethod: "ResetFrameRate"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyframesMaxDist", GoMethod: "ResetKeyframesMaxDist"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxFrameRate", GoMethod: "ResetMaxFrameRate"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxHeight", GoMethod: "ResetMaxHeight"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxWidth", GoMethod: "ResetMaxWidth"},
			_jsii_.MemberMethod{JsiiMethod: "resetPaddingPolicy", GoMethod: "ResetPaddingPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetResolution", GoMethod: "ResetResolution"},
			_jsii_.MemberMethod{JsiiMethod: "resetSizingPolicy", GoMethod: "ResetSizingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "resolution", GoGetter: "Resolution"},
			_jsii_.MemberProperty{JsiiProperty: "resolutionInput", GoGetter: "ResolutionInput"},
			_jsii_.MemberProperty{JsiiProperty: "sizingPolicy", GoGetter: "SizingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "sizingPolicyInput", GoGetter: "SizingPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
		},
		func() interface{} {
			j := jsiiProxy_ElastictranscoderPresetVideoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"hashicorp_aws.ElasticTranscoder.ElastictranscoderPresetVideoWatermarks",
		reflect.TypeOf((*ElastictranscoderPresetVideoWatermarks)(nil)).Elem(),
	)
}
