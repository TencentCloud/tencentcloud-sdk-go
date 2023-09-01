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

package v20200324

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Age struct {
	// 人体年龄信息，返回值为以下集合中的一个{小孩,青年,中年,老年}。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0],代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

type AttributesOptions struct {
	// 返回年龄信息
	Age *bool `json:"Age,omitnil" name:"Age"`

	// 返回随身挎包信息
	Bag *bool `json:"Bag,omitnil" name:"Bag"`

	// 返回性别信息
	Gender *bool `json:"Gender,omitnil" name:"Gender"`

	// 返回朝向信息
	Orientation *bool `json:"Orientation,omitnil" name:"Orientation"`

	// 返回上装信息
	UpperBodyCloth *bool `json:"UpperBodyCloth,omitnil" name:"UpperBodyCloth"`

	// 返回下装信息
	LowerBodyCloth *bool `json:"LowerBodyCloth,omitnil" name:"LowerBodyCloth"`
}

type Bag struct {
	// 挎包信息，返回值为以下集合中的一个{双肩包, 斜挎包, 手拎包, 无包}。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0],代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

type BodyAttributeInfo struct {
	// 人体年龄信息。 
	// AttributesType 不含 Age 或检测超过 5 个人体时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Age *Age `json:"Age,omitnil" name:"Age"`

	// 人体是否挎包。 
	// AttributesType 不含 Bag 或检测超过 5 个人体时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bag *Bag `json:"Bag,omitnil" name:"Bag"`

	// 人体性别信息。 
	// AttributesType 不含 Gender 或检测超过 5 个人体时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Gender *Gender `json:"Gender,omitnil" name:"Gender"`

	// 人体朝向信息。   
	// AttributesType 不含 UpperBodyCloth 或检测超过 5 个人体时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Orientation *Orientation `json:"Orientation,omitnil" name:"Orientation"`

	// 人体上衣属性信息。
	// AttributesType 不含 Orientation 或检测超过 5 个人体时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpperBodyCloth *UpperBodyCloth `json:"UpperBodyCloth,omitnil" name:"UpperBodyCloth"`

	// 人体下衣属性信息。  
	// AttributesType 不含 LowerBodyCloth 或检测超过 5 个人体时，此参数仍返回，但不具备参考意义。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LowerBodyCloth *LowerBodyCloth `json:"LowerBodyCloth,omitnil" name:"LowerBodyCloth"`
}

type BodyDetectResult struct {
	// 检测出的人体置信度。 
	// 误识率百分之十对应的阈值是0.14；误识率百分之五对应的阈值是0.32；误识率百分之二对应的阈值是0.62；误识率百分之一对应的阈值是0.81。 
	// 通常情况建议使用阈值0.32，可适用大多数情况。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`

	// 图中检测出来的人体框
	BodyRect *BodyRect `json:"BodyRect,omitnil" name:"BodyRect"`

	// 图中检测出的人体属性信息。
	BodyAttributeInfo *BodyAttributeInfo `json:"BodyAttributeInfo,omitnil" name:"BodyAttributeInfo"`
}

type BodyJointsResult struct {
	// 图中检测出来的人体框。
	BoundBox *BoundRect `json:"BoundBox,omitnil" name:"BoundBox"`

	// 14个人体关键点的坐标，人体关键点详见KeyPointInfo。
	BodyJoints []*KeyPointInfo `json:"BodyJoints,omitnil" name:"BodyJoints"`

	// 检测出的人体置信度，0-1之间，数值越高越准确。
	Confidence *float64 `json:"Confidence,omitnil" name:"Confidence"`
}

type BodyRect struct {
	// 人体框左上角横坐标。
	X *uint64 `json:"X,omitnil" name:"X"`

	// 人体框左上角纵坐标。
	Y *uint64 `json:"Y,omitnil" name:"Y"`

	// 人体宽度。
	Width *uint64 `json:"Width,omitnil" name:"Width"`

	// 人体高度。
	Height *uint64 `json:"Height,omitnil" name:"Height"`
}

type BoundRect struct {
	// 人体框左上角横坐标。
	X *int64 `json:"X,omitnil" name:"X"`

	// 人体框左上角纵坐标。
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// 人体宽度。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 人体高度。
	Height *int64 `json:"Height,omitnil" name:"Height"`
}

type Candidate struct {
	// 人员ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 人体动作轨迹ID。
	TraceId *string `json:"TraceId,omitnil" name:"TraceId"`

	// 候选者的匹配得分。 
	// 十万人体库下，误识率百分之五对应的分数为70分；误识率百分之二对应的分数为80分；误识率百分之一对应的分数为90分。
	//  
	// 二十万人体库下，误识率百分之五对应的分数为80分；误识率百分之二对应的分数为90分；误识率百分之一对应的分数为95分。
	//  
	// 通常情况建议使用分数80分（保召回）。若希望获得较高精度，建议使用分数90分（保准确）。
	Score *float64 `json:"Score,omitnil" name:"Score"`
}

// Predefined struct for user
type CreateGroupRequestParams struct {
	// 人体库名称，[1,60]个字符，可修改，不可重复。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 人体库 ID，不可修改，不可重复。支持英文、数字、-%@#&_，长度限制64B。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 人体库信息备注，[0，40]个字符。
	Tag *string `json:"Tag,omitnil" name:"Tag"`

	// 人体识别所用的算法模型版本。 
	// 目前入参仅支持 “1.0”1个输入。 默认为"1.0"。  
	// 不同算法模型版本对应的人体识别算法不同，新版本的整体效果会优于旧版本，后续我们将推出更新版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitnil" name:"BodyModelVersion"`
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest
	
	// 人体库名称，[1,60]个字符，可修改，不可重复。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 人体库 ID，不可修改，不可重复。支持英文、数字、-%@#&_，长度限制64B。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 人体库信息备注，[0，40]个字符。
	Tag *string `json:"Tag,omitnil" name:"Tag"`

	// 人体识别所用的算法模型版本。 
	// 目前入参仅支持 “1.0”1个输入。 默认为"1.0"。  
	// 不同算法模型版本对应的人体识别算法不同，新版本的整体效果会优于旧版本，后续我们将推出更新版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitnil" name:"BodyModelVersion"`
}

func (r *CreateGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "GroupId")
	delete(f, "Tag")
	delete(f, "BodyModelVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateGroupResponseParams `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonRequestParams struct {
	// 待加入的人员库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 人员名称。[1，60]个字符，可修改，可重复。
	PersonName *string `json:"PersonName,omitnil" name:"PersonName"`

	// 人员ID，单个腾讯云账号下不可修改，不可重复。 
	// 支持英文、数字、-%@#&_，，长度限制64B。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 人体动作轨迹信息。
	Trace *Trace `json:"Trace,omitnil" name:"Trace"`
}

type CreatePersonRequest struct {
	*tchttp.BaseRequest
	
	// 待加入的人员库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 人员名称。[1，60]个字符，可修改，可重复。
	PersonName *string `json:"PersonName,omitnil" name:"PersonName"`

	// 人员ID，单个腾讯云账号下不可修改，不可重复。 
	// 支持英文、数字、-%@#&_，，长度限制64B。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 人体动作轨迹信息。
	Trace *Trace `json:"Trace,omitnil" name:"Trace"`
}

func (r *CreatePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "PersonName")
	delete(f, "PersonId")
	delete(f, "Trace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonResponseParams struct {
	// 人员动作轨迹唯一标识。
	TraceId *string `json:"TraceId,omitnil" name:"TraceId"`

	// 人体识别所用的算法模型版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitnil" name:"BodyModelVersion"`

	// 输入的人体动作轨迹图片中的合法性校验结果。
	// 只有为0时结果才有意义。
	// -1001: 输入图片不合法。-1002: 输入图片不能构成动作轨迹。
	InputRetCode *int64 `json:"InputRetCode,omitnil" name:"InputRetCode"`

	// 输入的人体动作轨迹图片中的合法性校验结果详情。 
	// -1101:图片无效，-1102:url不合法。-1103:图片过大。-1104:图片下载失败。-1105:图片解码失败。-1109:图片分辨率过高。-2023:动作轨迹中有非同人图片。-2024: 动作轨迹提取失败。-2025: 人体检测失败。
	// RetCode 的顺序和入参中Images 或 Urls 的顺序一致。
	InputRetCodeDetails []*int64 `json:"InputRetCodeDetails,omitnil" name:"InputRetCodeDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreatePersonResponse struct {
	*tchttp.BaseResponse
	Response *CreatePersonResponseParams `json:"Response"`
}

func (r *CreatePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSegmentationTaskRequestParams struct {
	// 需要分割的视频URL，可外网访问。
	VideoUrl *string `json:"VideoUrl,omitnil" name:"VideoUrl"`

	// 背景图片URL。 
	// 可以将视频背景替换为输入的图片。 
	// 如果不输入背景图片，则输出人像区域mask。
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil" name:"BackgroundImageUrl"`

	// 预留字段，后期用于展示更多识别信息。
	Config *string `json:"Config,omitnil" name:"Config"`
}

type CreateSegmentationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 需要分割的视频URL，可外网访问。
	VideoUrl *string `json:"VideoUrl,omitnil" name:"VideoUrl"`

	// 背景图片URL。 
	// 可以将视频背景替换为输入的图片。 
	// 如果不输入背景图片，则输出人像区域mask。
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitnil" name:"BackgroundImageUrl"`

	// 预留字段，后期用于展示更多识别信息。
	Config *string `json:"Config,omitnil" name:"Config"`
}

func (r *CreateSegmentationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSegmentationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VideoUrl")
	delete(f, "BackgroundImageUrl")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSegmentationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSegmentationTaskResponseParams struct {
	// 任务标识ID,可以用与追溯任务状态，查看任务结果
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 预估处理时间，单位为秒
	EstimatedProcessingTime *float64 `json:"EstimatedProcessingTime,omitnil" name:"EstimatedProcessingTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSegmentationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSegmentationTaskResponseParams `json:"Response"`
}

func (r *CreateSegmentationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSegmentationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceRequestParams struct {
	// 人员ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 人体动作轨迹信息。
	Trace *Trace `json:"Trace,omitnil" name:"Trace"`
}

type CreateTraceRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 人体动作轨迹信息。
	Trace *Trace `json:"Trace,omitnil" name:"Trace"`
}

func (r *CreateTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "Trace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTraceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTraceResponseParams struct {
	// 人员动作轨迹唯一标识。
	TraceId *string `json:"TraceId,omitnil" name:"TraceId"`

	// 人体识别所用的算法模型版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitnil" name:"BodyModelVersion"`

	// 输入的人体动作轨迹图片中的合法性校验结果。
	// 只有为0时结果才有意义。
	// -1001: 输入图片不合法。-1002: 输入图片不能构成轨迹。
	InputRetCode *int64 `json:"InputRetCode,omitnil" name:"InputRetCode"`

	// 输入的人体动作轨迹图片中的合法性校验结果详情。 
	// -1101:图片无效，-1102:url不合法。-1103:图片过大。-1104:图片下载失败。-1105:图片解码失败。-1109:图片分辨率过高。-2023:动作轨迹中有非同人图片。-2024: 动作轨迹提取失败。-2025: 人体检测失败。
	InputRetCodeDetails []*int64 `json:"InputRetCodeDetails,omitnil" name:"InputRetCodeDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTraceResponse struct {
	*tchttp.BaseResponse
	Response *CreateTraceResponseParams `json:"Response"`
}

func (r *CreateTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupRequestParams struct {
	// 人体库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest
	
	// 人体库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGroupResponseParams `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonRequestParams struct {
	// 人员ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`
}

type DeletePersonRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`
}

func (r *DeletePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeletePersonResponse struct {
	*tchttp.BaseResponse
	Response *DeletePersonResponseParams `json:"Response"`
}

func (r *DeletePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSegmentationTaskRequestParams struct {
	// 在提交分割任务成功时返回的任务标识ID。
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`
}

type DescribeSegmentationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 在提交分割任务成功时返回的任务标识ID。
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`
}

func (r *DescribeSegmentationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSegmentationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSegmentationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSegmentationTaskResponseParams struct {
	// 当前任务状态：
	// QUEUING 排队中
	// PROCESSING 处理中
	// FINISHED 处理完成
	TaskStatus *string `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// 分割后视频URL, 存储于腾讯云COS
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultVideoUrl *string `json:"ResultVideoUrl,omitnil" name:"ResultVideoUrl"`

	// 分割后视频MD5，用于校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultVideoMD5 *string `json:"ResultVideoMD5,omitnil" name:"ResultVideoMD5"`

	// 视频基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoBasicInformation *VideoBasicInformation `json:"VideoBasicInformation,omitnil" name:"VideoBasicInformation"`

	// 分割任务错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSegmentationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSegmentationTaskResponseParams `json:"Response"`
}

func (r *DescribeSegmentationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSegmentationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectBodyJointsRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M。  
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 人体局部关键点识别，开启后对人体局部图（例如部分身体部位）进行关键点识别，输出人体关键点坐标，默认不开启
	// 
	// 注意：若开启人体局部图片关键点识别，则BoundBox、Confidence返回为空。
	LocalBodySwitch *bool `json:"LocalBodySwitch,omitnil" name:"LocalBodySwitch"`
}

type DetectBodyJointsRequest struct {
	*tchttp.BaseRequest
	
	// 图片 base64 数据，base64 编码后大小不可超过5M。  
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。  
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 人体局部关键点识别，开启后对人体局部图（例如部分身体部位）进行关键点识别，输出人体关键点坐标，默认不开启
	// 
	// 注意：若开启人体局部图片关键点识别，则BoundBox、Confidence返回为空。
	LocalBodySwitch *bool `json:"LocalBodySwitch,omitnil" name:"LocalBodySwitch"`
}

func (r *DetectBodyJointsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectBodyJointsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "LocalBodySwitch")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectBodyJointsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectBodyJointsResponseParams struct {
	// 图中检测出的人体框和人体关键点， 包含14个人体关键点的坐标，建议根据人体框置信度筛选出合格的人体；
	BodyJointsResults []*BodyJointsResult `json:"BodyJointsResults,omitnil" name:"BodyJointsResults"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DetectBodyJointsResponse struct {
	*tchttp.BaseResponse
	Response *DetectBodyJointsResponseParams `json:"Response"`
}

func (r *DetectBodyJointsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectBodyJointsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectBodyRequestParams struct {
	// 人体图片 Base64 数据。
	// 图片 base64 编码后大小不可超过5M。
	// 图片分辨率不得超过 1920 * 1080 。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 最多检测的人体数目，默认值为1（仅检测图片中面积最大的那个人体）； 最大值10 ，检测图片中面积最大的10个人体。
	MaxBodyNum *uint64 `json:"MaxBodyNum,omitnil" name:"MaxBodyNum"`

	// 人体图片 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片 base64 编码后大小不可超过5M。 
	// 图片分辨率不得超过 1920 * 1080 。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 是否返回年龄、性别、朝向等属性。 
	// 可选项有 Age、Bag、Gender、UpperBodyCloth、LowerBodyCloth、Orientation。  
	// 如果此参数为空则为不需要返回。 
	// 需要将属性组成一个用逗号分隔的字符串，属性之间的顺序没有要求。 
	// 关于各属性的详细描述，参见下文出参。 
	// 最多返回面积最大的 5 个人体属性信息，超过 5 个人体（第 6 个及以后的人体）的 BodyAttributesInfo 不具备参考意义。
	AttributesOptions *AttributesOptions `json:"AttributesOptions,omitnil" name:"AttributesOptions"`
}

type DetectBodyRequest struct {
	*tchttp.BaseRequest
	
	// 人体图片 Base64 数据。
	// 图片 base64 编码后大小不可超过5M。
	// 图片分辨率不得超过 1920 * 1080 。
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 最多检测的人体数目，默认值为1（仅检测图片中面积最大的那个人体）； 最大值10 ，检测图片中面积最大的10个人体。
	MaxBodyNum *uint64 `json:"MaxBodyNum,omitnil" name:"MaxBodyNum"`

	// 人体图片 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片 base64 编码后大小不可超过5M。 
	// 图片分辨率不得超过 1920 * 1080 。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 是否返回年龄、性别、朝向等属性。 
	// 可选项有 Age、Bag、Gender、UpperBodyCloth、LowerBodyCloth、Orientation。  
	// 如果此参数为空则为不需要返回。 
	// 需要将属性组成一个用逗号分隔的字符串，属性之间的顺序没有要求。 
	// 关于各属性的详细描述，参见下文出参。 
	// 最多返回面积最大的 5 个人体属性信息，超过 5 个人体（第 6 个及以后的人体）的 BodyAttributesInfo 不具备参考意义。
	AttributesOptions *AttributesOptions `json:"AttributesOptions,omitnil" name:"AttributesOptions"`
}

func (r *DetectBodyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectBodyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "MaxBodyNum")
	delete(f, "Url")
	delete(f, "AttributesOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetectBodyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetectBodyResponseParams struct {
	// 图中检测出来的人体框。
	BodyDetectResults []*BodyDetectResult `json:"BodyDetectResults,omitnil" name:"BodyDetectResults"`

	// 人体识别所用的算法模型版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitnil" name:"BodyModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DetectBodyResponse struct {
	*tchttp.BaseResponse
	Response *DetectBodyResponseParams `json:"Response"`
}

func (r *DetectBodyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetectBodyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Gender struct {
	// 性别信息，返回值为以下集合中的一个 {男性, 女性}
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0],代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

// Predefined struct for user
type GetGroupListRequestParams struct {
	// 起始序号，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type GetGroupListRequest struct {
	*tchttp.BaseRequest
	
	// 起始序号，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *GetGroupListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetGroupListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetGroupListResponseParams struct {
	// 返回的人体库信息。
	GroupInfos []*GroupInfo `json:"GroupInfos,omitnil" name:"GroupInfos"`

	// 人体库总数量。
	GroupNum *uint64 `json:"GroupNum,omitnil" name:"GroupNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetGroupListResponse struct {
	*tchttp.BaseResponse
	Response *GetGroupListResponseParams `json:"Response"`
}

func (r *GetGroupListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListRequestParams struct {
	// 人体库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 起始序号，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type GetPersonListRequest struct {
	*tchttp.BaseRequest
	
	// 人体库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 起始序号，默认值为0。
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回数量，默认值为10，最大值为1000。
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *GetPersonListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetPersonListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetPersonListResponseParams struct {
	// 返回的人员信息。
	PersonInfos []*PersonInfo `json:"PersonInfos,omitnil" name:"PersonInfos"`

	// 该人体库的人员数量。
	PersonNum *uint64 `json:"PersonNum,omitnil" name:"PersonNum"`

	// 人体识别所用的算法模型版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitnil" name:"BodyModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetPersonListResponse struct {
	*tchttp.BaseResponse
	Response *GetPersonListResponseParams `json:"Response"`
}

func (r *GetPersonListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetPersonListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSummaryInfoRequestParams struct {

}

type GetSummaryInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *GetSummaryInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSummaryInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetSummaryInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetSummaryInfoResponseParams struct {
	// 人体库总数量。
	GroupCount *uint64 `json:"GroupCount,omitnil" name:"GroupCount"`

	// 人员总数量
	PersonCount *uint64 `json:"PersonCount,omitnil" name:"PersonCount"`

	// 人员轨迹总数量
	TraceCount *uint64 `json:"TraceCount,omitnil" name:"TraceCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GetSummaryInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetSummaryInfoResponseParams `json:"Response"`
}

func (r *GetSummaryInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetSummaryInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {
	// 人体库名称。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 人体库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 人体库信息备注。
	Tag *string `json:"Tag,omitnil" name:"Tag"`

	// 人体识别所用的算法模型版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitnil" name:"BodyModelVersion"`

	// Group的创建时间和日期 CreationTimestamp。CreationTimestamp 的值是自 Unix 纪元时间到Group创建时间的毫秒数。  
	// Unix 纪元时间是 1970 年 1 月 1 日星期四，协调世界时 (UTC) 。
	CreationTimestamp *uint64 `json:"CreationTimestamp,omitnil" name:"CreationTimestamp"`
}

type ImageRect struct {
	// 左上角横坐标。
	X *int64 `json:"X,omitnil" name:"X"`

	// 左上角纵坐标。
	Y *int64 `json:"Y,omitnil" name:"Y"`

	// 人体宽度。
	Width *int64 `json:"Width,omitnil" name:"Width"`

	// 人体高度。
	Height *int64 `json:"Height,omitnil" name:"Height"`

	// 分割选项名称。
	Label *string `json:"Label,omitnil" name:"Label"`
}

type KeyPointInfo struct {
	// 代表不同位置的人体关键点信息，返回值为以下集合中的一个 [头部,颈部,右肩,右肘,右腕,左肩,左肘,左腕,右髋,右膝,右踝,左髋,左膝,左踝]
	KeyPointType *string `json:"KeyPointType,omitnil" name:"KeyPointType"`

	// 人体关键点横坐标
	X *float64 `json:"X,omitnil" name:"X"`

	// 人体关键点纵坐标
	Y *float64 `json:"Y,omitnil" name:"Y"`

	// 关键点坐标置信度，分数取值在0-1之间，阈值建议为0.25，小于0.25认为在图中无人体关键点。
	BodyScore *float64 `json:"BodyScore,omitnil" name:"BodyScore"`
}

type LowerBodyCloth struct {
	// 下衣颜色信息。
	Color *LowerBodyClothColor `json:"Color,omitnil" name:"Color"`

	// 下衣长度信息 。
	Length *LowerBodyClothLength `json:"Length,omitnil" name:"Length"`

	// 下衣类型信息。
	Type *LowerBodyClothType `json:"Type,omitnil" name:"Type"`
}

type LowerBodyClothColor struct {
	// 下衣颜色信息，返回值为以下集合中的一个{ 黑色系, 灰白色系, 彩色} 。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0],代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

type LowerBodyClothLength struct {
	// 下衣长度信息，返回值为以下集合中的一个，{长, 短} 。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0],代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

type LowerBodyClothType struct {
	// 下衣类型，返回值为以下集合中的一个 {裤子,裙子} 。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0],代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

// Predefined struct for user
type ModifyGroupRequestParams struct {
	// 人体库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 人体库名称。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 人体库信息备注。
	Tag *string `json:"Tag,omitnil" name:"Tag"`
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest
	
	// 人体库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 人体库名称。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 人体库信息备注。
	Tag *string `json:"Tag,omitnil" name:"Tag"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "Tag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGroupResponseParams `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonInfoRequestParams struct {
	// 人员ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 人员名称。
	PersonName *string `json:"PersonName,omitnil" name:"PersonName"`
}

type ModifyPersonInfoRequest struct {
	*tchttp.BaseRequest
	
	// 人员ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 人员名称。
	PersonName *string `json:"PersonName,omitnil" name:"PersonName"`
}

func (r *ModifyPersonInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PersonId")
	delete(f, "PersonName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPersonInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyPersonInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPersonInfoResponseParams `json:"Response"`
}

func (r *ModifyPersonInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Orientation struct {
	// 人体朝向信息，返回值为以下集合中的一个 {正向, 背向, 左, 右}。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0],代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

type PersonInfo struct {
	// 人员名称。
	PersonName *string `json:"PersonName,omitnil" name:"PersonName"`

	// 人员ID。
	PersonId *string `json:"PersonId,omitnil" name:"PersonId"`

	// 包含的人体动作轨迹图片信息列表。
	TraceInfos []*TraceInfo `json:"TraceInfos,omitnil" name:"TraceInfos"`
}

// Predefined struct for user
type SearchTraceRequestParams struct {
	// 希望搜索的人体库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 人体动作轨迹信息。
	Trace *Trace `json:"Trace,omitnil" name:"Trace"`

	// 单张被识别的人体动作轨迹返回的最相似人员数量。
	// 默认值为5，最大值为100。
	//  例，设MaxPersonNum为8，则返回Top8相似的人员信息。 值越大，需要处理的时间越长。建议不要超过10。
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil" name:"MaxPersonNum"`

	// 出参Score中，只有超过TraceMatchThreshold值的结果才会返回。
	// 默认为0。范围[0, 100.0]。
	TraceMatchThreshold *float64 `json:"TraceMatchThreshold,omitnil" name:"TraceMatchThreshold"`
}

type SearchTraceRequest struct {
	*tchttp.BaseRequest
	
	// 希望搜索的人体库ID。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 人体动作轨迹信息。
	Trace *Trace `json:"Trace,omitnil" name:"Trace"`

	// 单张被识别的人体动作轨迹返回的最相似人员数量。
	// 默认值为5，最大值为100。
	//  例，设MaxPersonNum为8，则返回Top8相似的人员信息。 值越大，需要处理的时间越长。建议不要超过10。
	MaxPersonNum *uint64 `json:"MaxPersonNum,omitnil" name:"MaxPersonNum"`

	// 出参Score中，只有超过TraceMatchThreshold值的结果才会返回。
	// 默认为0。范围[0, 100.0]。
	TraceMatchThreshold *float64 `json:"TraceMatchThreshold,omitnil" name:"TraceMatchThreshold"`
}

func (r *SearchTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchTraceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Trace")
	delete(f, "MaxPersonNum")
	delete(f, "TraceMatchThreshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchTraceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchTraceResponseParams struct {
	// 识别出的最相似候选人。
	Candidates []*Candidate `json:"Candidates,omitnil" name:"Candidates"`

	// 输入的人体动作轨迹图片中的合法性校验结果。
	// 只有为0时结果才有意义。
	// -1001: 输入图片不合法。-1002: 输入图片不能构成动作轨迹。
	InputRetCode *int64 `json:"InputRetCode,omitnil" name:"InputRetCode"`

	// 输入的人体动作轨迹图片中的合法性校验结果详情。 
	// -1101:图片无效，-1102:url不合法。-1103:图片过大。-1104:图片下载失败。-1105:图片解码失败。-1109:图片分辨率过高。-2023:动作轨迹中有非同人图片。-2024: 动作轨迹提取失败。-2025: 人体检测失败。
	InputRetCodeDetails []*int64 `json:"InputRetCodeDetails,omitnil" name:"InputRetCodeDetails"`

	// 人体识别所用的算法模型版本。
	BodyModelVersion *string `json:"BodyModelVersion,omitnil" name:"BodyModelVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SearchTraceResponse struct {
	*tchttp.BaseResponse
	Response *SearchTraceResponseParams `json:"Response"`
}

func (r *SearchTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SegmentCustomizedPortraitPicRequestParams struct {
	// 此参数为分割选项，请根据需要选择自己所想从图片中分割的部分。注意所有选项均为非必选，如未选择则值默认为false, 但是必须要保证多于一个选项的描述为true。
	SegmentationOptions *SegmentationOptions `json:"SegmentationOptions,omitnil" name:"SegmentationOptions"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 图片分辨率须小于2000*2000。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 图片的 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片分辨率须小于2000*2000 ，图片 base64 编码后大小不可超过5M。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil" name:"Url"`
}

type SegmentCustomizedPortraitPicRequest struct {
	*tchttp.BaseRequest
	
	// 此参数为分割选项，请根据需要选择自己所想从图片中分割的部分。注意所有选项均为非必选，如未选择则值默认为false, 但是必须要保证多于一个选项的描述为true。
	SegmentationOptions *SegmentationOptions `json:"SegmentationOptions,omitnil" name:"SegmentationOptions"`

	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 图片分辨率须小于2000*2000。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 图片的 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片分辨率须小于2000*2000 ，图片 base64 编码后大小不可超过5M。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil" name:"Url"`
}

func (r *SegmentCustomizedPortraitPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SegmentCustomizedPortraitPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SegmentationOptions")
	delete(f, "Image")
	delete(f, "Url")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SegmentCustomizedPortraitPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SegmentCustomizedPortraitPicResponseParams struct {
	// 根据指定标签分割输出的透明背景人像图片的 base64 数据。
	PortraitImage *string `json:"PortraitImage,omitnil" name:"PortraitImage"`

	// 指定标签处理后的Mask。一个通过 Base64 编码的文件，解码后文件由 Float 型浮点数组成。这些浮点数代表原图从左上角开始的每一行的每一个像素点，每一个浮点数的值是原图相应像素点位于人体轮廓内的置信度（0-1）转化的灰度值（0-255）
	MaskImage *string `json:"MaskImage,omitnil" name:"MaskImage"`

	// 坐标信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageRects []*ImageRect `json:"ImageRects,omitnil" name:"ImageRects"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SegmentCustomizedPortraitPicResponse struct {
	*tchttp.BaseResponse
	Response *SegmentCustomizedPortraitPicResponseParams `json:"Response"`
}

func (r *SegmentCustomizedPortraitPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SegmentCustomizedPortraitPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SegmentPortraitPicRequestParams struct {
	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 图片分辨率须小于2000*2000。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 图片的 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片分辨率须小于2000*2000 ，图片 base64 编码后大小不可超过5M。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 返回图像方式（base64 或 Url ) ，二选一。url有效期为30分钟。
	RspImgType *string `json:"RspImgType,omitnil" name:"RspImgType"`

	// 适用场景类型。
	// 
	// 取值：GEN/GS。GEN为通用场景模式；GS为绿幕场景模式，针对绿幕场景下的人像分割效果更好。
	// 两种模式选择一种传入，默认为GEN。
	Scene *string `json:"Scene,omitnil" name:"Scene"`
}

type SegmentPortraitPicRequest struct {
	*tchttp.BaseRequest
	
	// 图片 base64 数据，base64 编码后大小不可超过5M。
	// 图片分辨率须小于2000*2000。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitnil" name:"Image"`

	// 图片的 Url 。
	// Url、Image必须提供一个，如果都提供，只使用 Url。
	// 图片分辨率须小于2000*2000 ，图片 base64 编码后大小不可超过5M。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitnil" name:"Url"`

	// 返回图像方式（base64 或 Url ) ，二选一。url有效期为30分钟。
	RspImgType *string `json:"RspImgType,omitnil" name:"RspImgType"`

	// 适用场景类型。
	// 
	// 取值：GEN/GS。GEN为通用场景模式；GS为绿幕场景模式，针对绿幕场景下的人像分割效果更好。
	// 两种模式选择一种传入，默认为GEN。
	Scene *string `json:"Scene,omitnil" name:"Scene"`
}

func (r *SegmentPortraitPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SegmentPortraitPicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Image")
	delete(f, "Url")
	delete(f, "RspImgType")
	delete(f, "Scene")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SegmentPortraitPicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SegmentPortraitPicResponseParams struct {
	// 处理后的图片 base64 数据，透明背景图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultImage *string `json:"ResultImage,omitnil" name:"ResultImage"`

	// 一个通过 base64 编码的文件，解码后文件由 Float 型浮点数组成。这些浮点数代表原图从左上角开始的每一行的每一个像素点，每一个浮点数的值是原图相应像素点位于人体轮廓内的置信度（0-1）转化的灰度值（0-255）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultMask *string `json:"ResultMask,omitnil" name:"ResultMask"`

	// 图片是否存在前景。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasForeground *bool `json:"HasForeground,omitnil" name:"HasForeground"`

	// 支持将处理过的图片 base64 数据，透明背景图以Url的形式返回值，Url有效期为30分钟。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultImageUrl *string `json:"ResultImageUrl,omitnil" name:"ResultImageUrl"`

	// 一个通过 base64 编码的文件，解码后文件由 Float 型浮点数组成。支持以Url形式的返回值；Url有效期为30分钟。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultMaskUrl *string `json:"ResultMaskUrl,omitnil" name:"ResultMaskUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SegmentPortraitPicResponse struct {
	*tchttp.BaseResponse
	Response *SegmentPortraitPicResponseParams `json:"Response"`
}

func (r *SegmentPortraitPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SegmentPortraitPicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SegmentationOptions struct {
	// 分割选项-背景
	Background *bool `json:"Background,omitnil" name:"Background"`

	// 分割选项-头发
	Hair *bool `json:"Hair,omitnil" name:"Hair"`

	// 分割选项-左眉
	LeftEyebrow *bool `json:"LeftEyebrow,omitnil" name:"LeftEyebrow"`

	// 分割选项-右眉
	RightEyebrow *bool `json:"RightEyebrow,omitnil" name:"RightEyebrow"`

	// 分割选项-左眼
	LeftEye *bool `json:"LeftEye,omitnil" name:"LeftEye"`

	// 分割选项-右眼
	RightEye *bool `json:"RightEye,omitnil" name:"RightEye"`

	// 分割选项-鼻子
	Nose *bool `json:"Nose,omitnil" name:"Nose"`

	// 分割选项-上唇
	UpperLip *bool `json:"UpperLip,omitnil" name:"UpperLip"`

	// 分割选项-下唇
	LowerLip *bool `json:"LowerLip,omitnil" name:"LowerLip"`

	// 分割选项-牙齿
	Tooth *bool `json:"Tooth,omitnil" name:"Tooth"`

	// 分割选项-口腔（不包含牙齿）
	Mouth *bool `json:"Mouth,omitnil" name:"Mouth"`

	// 分割选项-左耳
	LeftEar *bool `json:"LeftEar,omitnil" name:"LeftEar"`

	// 分割选项-右耳
	RightEar *bool `json:"RightEar,omitnil" name:"RightEar"`

	// 分割选项-面部(不包含眼、耳、口、鼻等五官及头发。)
	Face *bool `json:"Face,omitnil" name:"Face"`

	// 复合分割选项-头部(包含所有的头部元素，相关装饰除外)
	Head *bool `json:"Head,omitnil" name:"Head"`

	// 分割选项-身体（包含脖子）
	Body *bool `json:"Body,omitnil" name:"Body"`

	// 分割选项-帽子
	Hat *bool `json:"Hat,omitnil" name:"Hat"`

	// 分割选项-头饰
	Headdress *bool `json:"Headdress,omitnil" name:"Headdress"`

	// 分割选项-耳环
	Earrings *bool `json:"Earrings,omitnil" name:"Earrings"`

	// 分割选项-项链
	Necklace *bool `json:"Necklace,omitnil" name:"Necklace"`

	// 分割选项-随身物品（ 例如伞、包、手机等。 ）
	Belongings *bool `json:"Belongings,omitnil" name:"Belongings"`
}

// Predefined struct for user
type TerminateSegmentationTaskRequestParams struct {
	// 在提交分割任务成功时返回的任务标识ID。
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`
}

type TerminateSegmentationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 在提交分割任务成功时返回的任务标识ID。
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`
}

func (r *TerminateSegmentationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateSegmentationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateSegmentationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateSegmentationTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TerminateSegmentationTaskResponse struct {
	*tchttp.BaseResponse
	Response *TerminateSegmentationTaskResponseParams `json:"Response"`
}

func (r *TerminateSegmentationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateSegmentationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Trace struct {
	// 人体动作轨迹图片 Base64 数组。 
	// 数组长度最小为1最大为5。 
	// 单个图片 base64 编码后大小不可超过2M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Images []*string `json:"Images,omitnil" name:"Images"`

	// 人体动作轨迹图片 Url 数组。 
	// 数组长度最小为1最大为5。 
	// 单个图片 base64 编码后大小不可超过2M。 
	// Urls、Images必须提供一个，如果都提供，只使用 Urls。 
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Urls []*string `json:"Urls,omitnil" name:"Urls"`

	// 若输入的Images 和 Urls 是已经裁剪后的人体小图，则可以忽略本参数。 
	// 若否，或图片中包含多个人体，则需要通过本参数来指定图片中的人体框。 
	// 顺序对应 Images 或 Urls 中的顺序。  
	// 当不输入本参数时，我们将认为输入图片已是经过裁剪后的人体小图，不会进行人体检测而直接进行特征提取处理。
	BodyRects []*BodyRect `json:"BodyRects,omitnil" name:"BodyRects"`
}

type TraceInfo struct {
	// 人体动作轨迹ID。
	TraceId *string `json:"TraceId,omitnil" name:"TraceId"`

	// 包含的人体动作轨迹图片Id列表。
	BodyIds []*string `json:"BodyIds,omitnil" name:"BodyIds"`
}

type UpperBodyCloth struct {
	// 上衣纹理信息。
	Texture *UpperBodyClothTexture `json:"Texture,omitnil" name:"Texture"`

	// 上衣颜色信息。
	Color *UpperBodyClothColor `json:"Color,omitnil" name:"Color"`

	// 上衣衣袖信息。
	Sleeve *UpperBodyClothSleeve `json:"Sleeve,omitnil" name:"Sleeve"`
}

type UpperBodyClothColor struct {
	// 上衣颜色信息，返回值为以下集合中的一个 {红色系, 黄色系, 绿色系, 蓝色系, 黑色系, 灰白色系。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0],代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

type UpperBodyClothSleeve struct {
	// 上衣衣袖信息, 返回值为以下集合中的一个 {长袖, 短袖}。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0],代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

type UpperBodyClothTexture struct {
	// 上衣纹理信息，返回值为以下集合中的一个, {纯色, 格子, 大色块}。
	Type *string `json:"Type,omitnil" name:"Type"`

	// Type识别概率值，[0.0,1.0], 代表判断正确的概率。如0.8则代表有Type值有80%概率正确。
	Probability *float64 `json:"Probability,omitnil" name:"Probability"`
}

type VideoBasicInformation struct {
	// 视频宽度
	FrameWidth *int64 `json:"FrameWidth,omitnil" name:"FrameWidth"`

	// 视频高度
	FrameHeight *int64 `json:"FrameHeight,omitnil" name:"FrameHeight"`

	// 视频帧速率(FPS)
	FramesPerSecond *int64 `json:"FramesPerSecond,omitnil" name:"FramesPerSecond"`

	// 视频时长
	Duration *float64 `json:"Duration,omitnil" name:"Duration"`

	// 视频帧数
	TotalFrames *int64 `json:"TotalFrames,omitnil" name:"TotalFrames"`
}