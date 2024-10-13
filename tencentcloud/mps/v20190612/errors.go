// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190612

const (
	// 此产品的特有错误码

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 操作失败：bucket 已经设置通知。
	FAILEDOPERATION_BUCKETNOTIFYALREADYEXIST = "FailedOperation.BucketNotifyAlreadyExist"

	// 操作失败：COS 已经停服。
	FAILEDOPERATION_COSSTATUSINAVLID = "FailedOperation.CosStatusInavlid"

	// 生成资源失败。
	FAILEDOPERATION_GENERATERESOURCE = "FailedOperation.GenerateResource"

	// 操作失败：获取源通知错误。
	FAILEDOPERATION_GETSOURCENOTIFY = "FailedOperation.GetSourceNotify"

	// 操作失败：非法 mps 用户。
	FAILEDOPERATION_INVALIDMPSUSER = "FailedOperation.InvalidMpsUser"

	// 操作失败：无效用户。
	FAILEDOPERATION_INVALIDUSER = "FailedOperation.InvalidUser"

	// 操作失败：设置源通知错误。
	FAILEDOPERATION_SETSOURCENOTIFY = "FailedOperation.SetSourceNotify"

	// 查询的任务不存在
	FAILEDOPERATION_TASKNOTFOUND = "FailedOperation.TaskNotFound"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 数据错误。
	INTERNALERROR_ACCESSDBERROR = "InternalError.AccessDBError"

	// 内部错误：生成模板 ID 失败。
	INTERNALERROR_GENDEFINITION = "InternalError.GenDefinition"

	// 内部错误：上传水印图片失败。
	INTERNALERROR_UPLOADWATERMARKERROR = "InternalError.UploadWatermarkError"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// InvalidParameter.EndTime
	INVALIDPARAMETER_ENDTIME = "InvalidParameter.EndTime"

	// InvalidParameter.ExceededQuantityLimit
	INVALIDPARAMETER_EXCEEDEDQUANTITYLIMIT = "InvalidParameter.ExceededQuantityLimit"

	// InvalidParameter.Id
	INVALIDPARAMETER_ID = "InvalidParameter.Id"

	// InvalidParameter.Input
	INVALIDPARAMETER_INPUT = "InvalidParameter.Input"

	// InvalidParameter.InputOutputId
	INVALIDPARAMETER_INPUTOUTPUTID = "InvalidParameter.InputOutputId"

	// InvalidParameter.MaxBandwidth
	INVALIDPARAMETER_MAXBANDWIDTH = "InvalidParameter.MaxBandwidth"

	// InvalidParameter.Name
	INVALIDPARAMETER_NAME = "InvalidParameter.Name"

	// InvalidParameter.NotFound
	INVALIDPARAMETER_NOTFOUND = "InvalidParameter.NotFound"

	// InvalidParameter.Output
	INVALIDPARAMETER_OUTPUT = "InvalidParameter.Output"

	// InvalidParameter.OutputGroups
	INVALIDPARAMETER_OUTPUTGROUPS = "InvalidParameter.OutputGroups"

	// InvalidParameter.OutputId
	INVALIDPARAMETER_OUTPUTID = "InvalidParameter.OutputId"

	// InvalidParameter.PageNum
	INVALIDPARAMETER_PAGENUM = "InvalidParameter.PageNum"

	// InvalidParameter.PageSize
	INVALIDPARAMETER_PAGESIZE = "InvalidParameter.PageSize"

	// InvalidParameter.Period
	INVALIDPARAMETER_PERIOD = "InvalidParameter.Period"

	// InvalidParameter.Pipeline
	INVALIDPARAMETER_PIPELINE = "InvalidParameter.Pipeline"

	// InvalidParameter.Protocol
	INVALIDPARAMETER_PROTOCOL = "InvalidParameter.Protocol"

	// InvalidParameter.SortType
	INVALIDPARAMETER_SORTTYPE = "InvalidParameter.SortType"

	// InvalidParameter.StartTime
	INVALIDPARAMETER_STARTTIME = "InvalidParameter.StartTime"

	// InvalidParameter.State
	INVALIDPARAMETER_STATE = "InvalidParameter.State"

	// InvalidParameter.Type
	INVALIDPARAMETER_TYPE = "InvalidParameter.Type"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// AAC帧时间戳间隔不合理配置错误
	INVALIDPARAMETERVALUE_AACDURATIONDEVIATION = "InvalidParameterValue.AACDurationDeviation"

	// 音视频交织不合理配置错误
	INVALIDPARAMETERVALUE_AVTIMESTAMPINTERLEAVE = "InvalidParameterValue.AVTimestampInterleave"

	// 参数错误：音频流码率。
	INVALIDPARAMETERVALUE_AUDIOBITRATE = "InvalidParameterValue.AudioBitrate"

	// 音频流码率超出范围配置错误
	INVALIDPARAMETERVALUE_AUDIOBITRATEOUTOFRANGE = "InvalidParameterValue.AudioBitrateOutofRange"

	// 参数值错误：AudioChannel。
	INVALIDPARAMETERVALUE_AUDIOCHANNEL = "InvalidParameterValue.AudioChannel"

	// 音频声道变化配置错误
	INVALIDPARAMETERVALUE_AUDIOCHANNELSCHANGED = "InvalidParameterValue.AudioChannelsChanged"

	// 参数错误：音频流编码格式。
	INVALIDPARAMETERVALUE_AUDIOCODEC = "InvalidParameterValue.AudioCodec"

	// 音频解码错误配置错误
	INVALIDPARAMETERVALUE_AUDIODECODEFAILED = "InvalidParameterValue.AudioDecodeFailed"

	// 音频丢帧配置错误
	INVALIDPARAMETERVALUE_AUDIODROPPINGFRAMES = "InvalidParameterValue.AudioDroppingFrames"

	// 音频流中存在重复帧配置错误
	INVALIDPARAMETERVALUE_AUDIODUPLICATEDFRAME = "InvalidParameterValue.AudioDuplicatedFrame"

	// 双通道音频相位相反配置错误
	INVALIDPARAMETERVALUE_AUDIOOUTOFPHASE = "InvalidParameterValue.AudioOutOfPhase"

	// 参数错误：音频流采样率。
	INVALIDPARAMETERVALUE_AUDIOSAMPLERATE = "InvalidParameterValue.AudioSampleRate"

	// 音频采样率变化配置错误
	INVALIDPARAMETERVALUE_AUDIOSAMPLERATECHANGED = "InvalidParameterValue.AudioSampleRateChanged"

	// 无音频流配置错误
	INVALIDPARAMETERVALUE_AUDIOSTREAMLACK = "InvalidParameterValue.AudioStreamLack"

	// 无效的音频/视频码率。
	INVALIDPARAMETERVALUE_BITRATE = "InvalidParameterValue.Bitrate"

	// 黑白边检测配置错误
	INVALIDPARAMETERVALUE_BLACKWHITEEDGE = "InvalidParameterValue.BlackWhiteEdge"

	// 参数值错误：BlockConfidence 参数取值非法。
	INVALIDPARAMETERVALUE_BLOCKCONFIDENCE = "InvalidParameterValue.BlockConfidence"

	// 模糊检测配置错误
	INVALIDPARAMETERVALUE_BLUR = "InvalidParameterValue.Blur"

	// 参数值错误：智能分类控制字段参数错误。
	INVALIDPARAMETERVALUE_CLASSIFCATIONCONFIGURE = "InvalidParameterValue.ClassifcationConfigure"

	// 无效的音频/视频编编码格式。
	INVALIDPARAMETERVALUE_CODEC = "InvalidParameterValue.Codec"

	// 参数值错误：ColumnCount。
	INVALIDPARAMETERVALUE_COLUMNCOUNT = "InvalidParameterValue.ColumnCount"

	// 参数错误：对该模板的描述。
	INVALIDPARAMETERVALUE_COMMENT = "InvalidParameterValue.Comment"

	// 参数错误：封装格式。
	INVALIDPARAMETERVALUE_CONTAINER = "InvalidParameterValue.Container"

	// 参数值错误：ContainerType。
	INVALIDPARAMETERVALUE_CONTAINERTYPE = "InvalidParameterValue.ContainerType"

	// 参数值错误：CoordinateOrigin。
	INVALIDPARAMETERVALUE_COORDINATEORIGIN = "InvalidParameterValue.CoordinateOrigin"

	// 参数值错误：智能封面控制字段参数错误。
	INVALIDPARAMETERVALUE_COVERCONFIGURE = "InvalidParameterValue.CoverConfigure"

	// 花屏检测配置错误
	INVALIDPARAMETERVALUE_CRASHSCREEN = "InvalidParameterValue.CrashScreen"

	// 视频的宽高比异常配置错误
	INVALIDPARAMETERVALUE_DARORSARINVALID = "InvalidParameterValue.DarOrSarInvalid"

	// 参数值错误：人脸默认库过滤标签非法。
	INVALIDPARAMETERVALUE_DEFAULTLIBRARYLABELSET = "InvalidParameterValue.DefaultLibraryLabelSet"

	// 参数错误：Definition。
	INVALIDPARAMETERVALUE_DEFINITION = "InvalidParameterValue.Definition"

	// 参数错误：Definitions。
	INVALIDPARAMETERVALUE_DEFINITIONS = "InvalidParameterValue.Definitions"

	// 参数值错误：不允许删除默认模板。
	INVALIDPARAMETERVALUE_DELETEDEFAULTTEMPLATE = "InvalidParameterValue.DeleteDefaultTemplate"

	// DestinationLanguage参数错误
	INVALIDPARAMETERVALUE_DESTINATIONLANGUAGE = "InvalidParameterValue.DestinationLanguage"

	// 无效的禁止码率低转高开关值。
	INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEOBITRATE = "InvalidParameterValue.DisableHigherVideoBitrate"

	// 无效的禁止分辨率低转高开关值。
	INVALIDPARAMETERVALUE_DISABLEHIGHERVIDEORESOLUTION = "InvalidParameterValue.DisableHigherVideoResolution"

	// DTS抖动过大配置错误
	INVALIDPARAMETERVALUE_DTSJITTER = "InvalidParameterValue.DtsJitter"

	// 模板配置的使能检测项为空
	INVALIDPARAMETERVALUE_EMPTYDETECTITEM = "InvalidParameterValue.EmptyDetectItem"

	// 参数值错误：人脸重复。
	INVALIDPARAMETERVALUE_FACEDUPLICATE = "InvalidParameterValue.FaceDuplicate"

	// 参数值错误：人脸库参数非法。
	INVALIDPARAMETERVALUE_FACELIBRARY = "InvalidParameterValue.FaceLibrary"

	// 参数值错误：人脸分数参数取值非法。
	INVALIDPARAMETERVALUE_FACESCORE = "InvalidParameterValue.FaceScore"

	// 参数错误：填充方式错误。
	INVALIDPARAMETERVALUE_FILLTYPE = "InvalidParameterValue.FillType"

	// 参数值错误：Format。
	INVALIDPARAMETERVALUE_FORMAT = "InvalidParameterValue.Format"

	// 参数值错误：Format 为 webp 时，Width、Height 均为空。
	INVALIDPARAMETERVALUE_FORMATWEBPLACKWIDTHANDHEIGHT = "InvalidParameterValue.FormatWebpLackWidthAndHeight"

	// 参数值错误：Format 为 webp 时，不允许 Width、Height 都为 0。
	INVALIDPARAMETERVALUE_FORMATWEBPWIDTHANDHEIGHTBOTHZERO = "InvalidParameterValue.FormatWebpWidthAndHeightBothZero"

	// 参数错误：视频帧率。
	INVALIDPARAMETERVALUE_FPS = "InvalidParameterValue.Fps"

	// 通过PTS计算得到的流帧率抖动过大配置错误
	INVALIDPARAMETERVALUE_FPSJITTER = "InvalidParameterValue.FpsJitter"

	// 参数值错误：智能按帧标签控制字段参数错误。
	INVALIDPARAMETERVALUE_FRAMETAGCONFIGURE = "InvalidParameterValue.FrameTagConfigure"

	// 参数值错误：FunctionArg。
	INVALIDPARAMETERVALUE_FUNCTIONARG = "InvalidParameterValue.FunctionArg"

	// 参数值错误：FunctionName。
	INVALIDPARAMETERVALUE_FUNCTIONNAME = "InvalidParameterValue.FunctionName"

	// 无效的Gop值。
	INVALIDPARAMETERVALUE_GOP = "InvalidParameterValue.Gop"

	// 无效的m3u8文件配置错误
	INVALIDPARAMETERVALUE_HLSBADM3U8FORMAT = "InvalidParameterValue.HLSBadM3u8Format"

	// 无效的master m3u8文件配置错误
	INVALIDPARAMETERVALUE_HLSINVALIDMASTERM3U8 = "InvalidParameterValue.HLSInvalidMasterM3u8"

	// 无效的media m3u8文件配置错误
	INVALIDPARAMETERVALUE_HLSINVALIDMEDIAM3U8 = "InvalidParameterValue.HLSInvalidMediaM3u8"

	// master m3u8缺少标准推荐的参数配置错误
	INVALIDPARAMETERVALUE_HLSMASTERM3U8RECOMMENDED = "InvalidParameterValue.HLSMasterM3u8Recommended"

	// media m3u8存在EXT-X-DISCONTINUITY配置错误
	INVALIDPARAMETERVALUE_HLSMEDIAM3U8DISCONTINUITYEXIST = "InvalidParameterValue.HLSMediaM3u8DiscontinuityExist"

	// media m3u8缺少标准推荐的参数配置错误
	INVALIDPARAMETERVALUE_HLSMEDIAM3U8RECOMMENDED = "InvalidParameterValue.HLSMediaM3u8Recommended"

	// 切片间DTS跳变且没有EXT-X-DISCONTINUITY配置错误
	INVALIDPARAMETERVALUE_HLSMEDIASEGMENTSDTSJITTERDEVIATION = "InvalidParameterValue.HLSMediaSegmentsDTSJitterDeviation"

	// 切片间PTS跳变且没有EXT-X-DISCONTINUITY配置错误
	INVALIDPARAMETERVALUE_HLSMEDIASEGMENTSPTSJITTERDEVIATION = "InvalidParameterValue.HLSMediaSegmentsPTSJitterDeviation"

	// 切片的流数目发生变化配置错误
	INVALIDPARAMETERVALUE_HLSMEDIASEGMENTSSTREAMNUMCHANGE = "InvalidParameterValue.HLSMediaSegmentsStreamNumChange"

	// 参数错误：高度。
	INVALIDPARAMETERVALUE_HEIGHT = "InvalidParameterValue.Height"

	// 亮度异常：过曝配置错误
	INVALIDPARAMETERVALUE_HIGHLIGHTING = "InvalidParameterValue.HighLighting"

	// 爆音检测配置错误
	INVALIDPARAMETERVALUE_HIGHVOICE = "InvalidParameterValue.HighVoice"

	// ImageContent参数值无效。
	INVALIDPARAMETERVALUE_IMAGECONTENT = "InvalidParameterValue.ImageContent"

	// 参数错误：图片水印模板。
	INVALIDPARAMETERVALUE_IMAGETEMPLATE = "InvalidParameterValue.ImageTemplate"

	// 解析内容 Content 的值不合法。
	INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"

	// 无效的操作类型。
	INVALIDPARAMETERVALUE_INVALIDOPERATIONTYPE = "InvalidParameterValue.InvalidOperationType"

	// 参数值错误：LabelSet 参数取值非法。
	INVALIDPARAMETERVALUE_LABELSET = "InvalidParameterValue.LabelSet"

	// 缺失音频流恢复配置错误
	INVALIDPARAMETERVALUE_LACKAUDIORECOVER = "InvalidParameterValue.LackAudioRecover"

	// 缺失视频流恢复配置错误
	INVALIDPARAMETERVALUE_LACKVIDEORECOVER = "InvalidParameterValue.LackVideoRecover"

	// 参数错误：Limit。
	INVALIDPARAMETERVALUE_LIMIT = "InvalidParameterValue.Limit"

	// 无参考打分配置错误
	INVALIDPARAMETERVALUE_LOWEVALUATION = "InvalidParameterValue.LowEvaluation"

	// 亮度异常：低光照配置错误
	INVALIDPARAMETERVALUE_LOWLIGHTING = "InvalidParameterValue.LowLighting"

	// 低音检测配置错误
	INVALIDPARAMETERVALUE_LOWVOICE = "InvalidParameterValue.LowVoice"

	// 参数值错误：不允许修改默认模板。
	INVALIDPARAMETERVALUE_MODIFYDEFAULTTEMPLATE = "InvalidParameterValue.ModifyDefaultTemplate"

	// 马赛克检测配置错误
	INVALIDPARAMETERVALUE_MOSAIC = "InvalidParameterValue.Mosaic"

	// MP4中codec fourcc不符合Apple HLS要求配置错误
	INVALIDPARAMETERVALUE_MP4INVALIDCODECFOURCC = "InvalidParameterValue.Mp4InvalidCodecFourcc"

	// 参数值错误：Name 超过长度限制。
	INVALIDPARAMETERVALUE_NAME = "InvalidParameterValue.Name"

	// 静音检测配置错误
	INVALIDPARAMETERVALUE_NOVOICE = "InvalidParameterValue.NoVoice"

	// 不支持状态不为处理中的任务。
	INVALIDPARAMETERVALUE_NOTPROCESSINGTASK = "InvalidParameterValue.NotProcessingTask"

	// 参数值错误：物体库参数非法。
	INVALIDPARAMETERVALUE_OBJECTLIBRARY = "InvalidParameterValue.ObjectLibrary"

	// 流参数集信息发生变化配置错误
	INVALIDPARAMETERVALUE_PARAMETERSETSCHANGED = "InvalidParameterValue.ParameterSetsChanged"

	// 参数值错误：人脸图片格式错误。
	INVALIDPARAMETERVALUE_PICFORMATERROR = "InvalidParameterValue.PicFormatError"

	// PTS抖动过大配置错误
	INVALIDPARAMETERVALUE_PTSJITTER = "InvalidParameterValue.PtsJitter"

	// 媒体流的 pts 小于 dts配置错误
	INVALIDPARAMETERVALUE_PTSLESSTHANDTS = "InvalidParameterValue.PtsLessThanDts"

	// 参数值错误：Quality。
	INVALIDPARAMETERVALUE_QUALITY = "InvalidParameterValue.Quality"

	// 网络接收帧率抖动过大配置错误
	INVALIDPARAMETERVALUE_RECEIVEFPSJITTER = "InvalidParameterValue.ReceiveFpsJitter"

	// 网络接收视频帧率过小配置错误
	INVALIDPARAMETERVALUE_RECEIVEFPSTOOSMALL = "InvalidParameterValue.ReceiveFpsTooSmall"

	// 参数值错误：RemoveAudio。
	INVALIDPARAMETERVALUE_REMOVEAUDIO = "InvalidParameterValue.RemoveAudio"

	// 参数值错误：RemoveVideo。
	INVALIDPARAMETERVALUE_REMOVEVIDEO = "InvalidParameterValue.RemoveVideo"

	// 参数错误：RepeatType 无效。
	INVALIDPARAMETERVALUE_REPEATTYPE = "InvalidParameterValue.RepeatType"

	// 参数错误：分辨率错误。
	INVALIDPARAMETERVALUE_RESOLUTION = "InvalidParameterValue.Resolution"

	// 无效的ResolutionAdaptive。
	INVALIDPARAMETERVALUE_RESOLUTIONADAPTIVE = "InvalidParameterValue.ResolutionAdaptive"

	// 参数值错误：ReviewConfidence 参数取值非法。
	INVALIDPARAMETERVALUE_REVIEWCONFIDENCE = "InvalidParameterValue.ReviewConfidence"

	// 参数值错误：RowCount。
	INVALIDPARAMETERVALUE_ROWCOUNT = "InvalidParameterValue.RowCount"

	// 参数值错误：SampleInterval。
	INVALIDPARAMETERVALUE_SAMPLEINTERVAL = "InvalidParameterValue.SampleInterval"

	// 无效的音频采样率。
	INVALIDPARAMETERVALUE_SAMPLERATE = "InvalidParameterValue.SampleRate"

	// 参数值错误：SampleType。
	INVALIDPARAMETERVALUE_SAMPLETYPE = "InvalidParameterValue.SampleType"

	// SessionContext 过长。
	INVALIDPARAMETERVALUE_SESSIONCONTEXTTOOLONG = "InvalidParameterValue.SessionContextTooLong"

	// 去重识别码重复，请求被去重。
	INVALIDPARAMETERVALUE_SESSIONID = "InvalidParameterValue.SessionId"

	// SessionId 过长。
	INVALIDPARAMETERVALUE_SESSIONIDTOOLONG = "InvalidParameterValue.SessionIdTooLong"

	// 参数错误：音频通道方式。
	INVALIDPARAMETERVALUE_SOUNDSYSTEM = "InvalidParameterValue.SoundSystem"

	// SourceLanguage参数错误
	INVALIDPARAMETERVALUE_SOURCELANGUAGE = "InvalidParameterValue.SourceLanguage"

	// 源文件错误。
	INVALIDPARAMETERVALUE_SRCFILE = "InvalidParameterValue.SrcFile"

	// 流结束配置错误
	INVALIDPARAMETERVALUE_STREAMEND = "InvalidParameterValue.StreamEnd"

	// NALU起始码错误配置错误
	INVALIDPARAMETERVALUE_STREAMNALUERROR = "InvalidParameterValue.StreamNALUError"

	// 流打开失败配置错误
	INVALIDPARAMETERVALUE_STREAMOPENFAILED = "InvalidParameterValue.StreamOpenFailed"

	// 流解析失败配置错误
	INVALIDPARAMETERVALUE_STREAMPARSEFAILED = "InvalidParameterValue.StreamParseFailed"

	// 参数值错误：SubtitleFormat 参数非法。
	INVALIDPARAMETERVALUE_SUBTITLEFORMAT = "InvalidParameterValue.SubtitleFormat"

	// 参数值错误：SVG 为空。
	INVALIDPARAMETERVALUE_SVGTEMPLATE = "InvalidParameterValue.SvgTemplate"

	// 参数值错误：SVG 高度。
	INVALIDPARAMETERVALUE_SVGTEMPLATEHEIGHT = "InvalidParameterValue.SvgTemplateHeight"

	// 参数值错误：SVG 宽度。
	INVALIDPARAMETERVALUE_SVGTEMPLATEWIDTH = "InvalidParameterValue.SvgTemplateWidth"

	// 参数值错误：Switch 参数取值非法。
	INVALIDPARAMETERVALUE_SWITCH = "InvalidParameterValue.Switch"

	// 参数值错误：TEHD Type 无效。
	INVALIDPARAMETERVALUE_TEHDTYPE = "InvalidParameterValue.TEHDType"

	// 参数值错误：智能标签控制字段参数错误。
	INVALIDPARAMETERVALUE_TAGCONFIGURE = "InvalidParameterValue.TagConfigure"

	// 任务 ID 不存在。
	INVALIDPARAMETERVALUE_TASKID = "InvalidParameterValue.TaskId"

	// 参数错误：文字透明度。
	INVALIDPARAMETERVALUE_TEXTALPHA = "InvalidParameterValue.TextAlpha"

	// 参数错误：文字模板。
	INVALIDPARAMETERVALUE_TEXTTEMPLATE = "InvalidParameterValue.TextTemplate"

	// MP4存在tmcd轨道配置错误
	INVALIDPARAMETERVALUE_TIMECODETRACKEXIST = "InvalidParameterValue.TimecodeTrackExist"

	// DTS时间戳回退配置错误
	INVALIDPARAMETERVALUE_TIMESTAMPFALLBACK = "InvalidParameterValue.TimestampFallback"

	// MPEG2-TS流有多个program配置错误
	INVALIDPARAMETERVALUE_TSMULTIPROGRAMS = "InvalidParameterValue.TsMultiPrograms"

	// mpegts的H26x流缺失 AUD NALU配置错误
	INVALIDPARAMETERVALUE_TSSTREAMNOAUD = "InvalidParameterValue.TsStreamNoAud"

	// 参数错误：Type 参数值错误。
	INVALIDPARAMETERVALUE_TYPE = "InvalidParameterValue.Type"

	// 未知诊断类别
	INVALIDPARAMETERVALUE_UNKNOWNCATEGORY = "InvalidParameterValue.UnknownCategory"

	// 参数值错误：人脸用户自定义库过滤标签非法。
	INVALIDPARAMETERVALUE_USERDEFINELIBRARYLABELSET = "InvalidParameterValue.UserDefineLibraryLabelSet"

	// 参数错误：视频流码率。
	INVALIDPARAMETERVALUE_VIDEOBITRATE = "InvalidParameterValue.VideoBitrate"

	// 视频流码率超出范围配置错误
	INVALIDPARAMETERVALUE_VIDEOBITRATEOUTOFRANGE = "InvalidParameterValue.VideoBitrateOutofRange"

	// 参数错误：视频流的编码格式。
	INVALIDPARAMETERVALUE_VIDEOCODEC = "InvalidParameterValue.VideoCodec"

	// 视频解码错误配置错误
	INVALIDPARAMETERVALUE_VIDEODECODEFAILED = "InvalidParameterValue.VideoDecodeFailed"

	// 视频丢帧配置错误
	INVALIDPARAMETERVALUE_VIDEODROPPINGFRAMES = "InvalidParameterValue.VideoDroppingFrames"

	// 视频流中存在重复帧配置错误
	INVALIDPARAMETERVALUE_VIDEODUPLICATEDFRAME = "InvalidParameterValue.VideoDuplicatedFrame"

	// 首帧不是IDR帧配置错误
	INVALIDPARAMETERVALUE_VIDEOFIRSTFRAMENOTIDR = "InvalidParameterValue.VideoFirstFrameNotIdr"

	// 视频冻结配置错误
	INVALIDPARAMETERVALUE_VIDEOFREEZEDFRAME = "InvalidParameterValue.VideoFreezedFrame"

	// 视频分辨率变化配置错误
	INVALIDPARAMETERVALUE_VIDEORESOLUTIONCHANGED = "InvalidParameterValue.VideoResolutionChanged"

	// 视频画面旋转配置错误
	INVALIDPARAMETERVALUE_VIDEOROTATION = "InvalidParameterValue.VideoRotation"

	// 无视频流配置错误
	INVALIDPARAMETERVALUE_VIDEOSTREAMLACK = "InvalidParameterValue.VideoStreamLack"

	// 参数错误：宽度。
	INVALIDPARAMETERVALUE_WIDTH = "InvalidParameterValue.Width"

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式。
	INVALIDPARAMETERVALUE_XPOS = "InvalidParameterValue.XPos"

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式。
	INVALIDPARAMETERVALUE_YPOS = "InvalidParameterValue.YPos"

	// 超过限制值：模板数超限。
	LIMITEXCEEDED_TOOMUCHTEMPLATE = "LimitExceeded.TooMuchTemplate"

	// 资源不存在。
	RESOURCENOTFOUND = "ResourceNotFound"

	// 资源不存在：Cos bucket 名称无效。
	RESOURCENOTFOUND_COSBUCKETNAMEINVALID = "ResourceNotFound.CosBucketNameInvalid"

	// 资源不存在：Cos bucket 不存在。
	RESOURCENOTFOUND_COSBUCKETNOTEXIST = "ResourceNotFound.CosBucketNotExist"

	// 资源不存在：人物。
	RESOURCENOTFOUND_PERSON = "ResourceNotFound.Person"

	// 资源不存在：模板不存在。
	RESOURCENOTFOUND_TEMPLATENOTEXIST = "ResourceNotFound.TemplateNotExist"

	// 资源不存在：关键词。
	RESOURCENOTFOUND_WORD = "ResourceNotFound.Word"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
)
