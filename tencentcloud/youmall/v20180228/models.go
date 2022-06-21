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

package v20180228

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ArrivedMallInfo struct {
	// 到场时间
	ArrivedTime *string `json:"ArrivedTime,omitempty" name:"ArrivedTime"`

	// 出场时间
	LeaveTime *string `json:"LeaveTime,omitempty" name:"LeaveTime"`

	// 停留时间，秒
	StaySecond *uint64 `json:"StaySecond,omitempty" name:"StaySecond"`

	// 到场抓拍图片
	InCapPic *string `json:"InCapPic,omitempty" name:"InCapPic"`

	// 出场抓拍图片
	OutCapPic *string `json:"OutCapPic,omitempty" name:"OutCapPic"`

	// 轨迹编码
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`
}

type CameraPersonInfo struct {
	// 临时id，还未生成face id时返回
	TempId *string `json:"TempId,omitempty" name:"TempId"`

	// 人脸face id
	FaceId *int64 `json:"FaceId,omitempty" name:"FaceId"`

	// 确定当次返回的哪个id有效，1-FaceId，2-TempId
	IdType *int64 `json:"IdType,omitempty" name:"IdType"`

	// 当次抓拍到的人脸图片base编码
	FacePic *string `json:"FacePic,omitempty" name:"FacePic"`

	// 当次抓拍时间戳
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 当前的person基本信息，图片以FacePic为准，结构体内未填
	PersonInfo *PersonInfo `json:"PersonInfo,omitempty" name:"PersonInfo"`
}

// Predefined struct for user
type CreateAccountRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 账号名；需要是手机号
	Name *string `json:"Name,omitempty" name:"Name"`

	// 密码；需要是(`~!@#$%^&*()_+=-）中的至少两种且八位以上
	Password *string `json:"Password,omitempty" name:"Password"`

	// 客户门店编码
	ShopCode *string `json:"ShopCode,omitempty" name:"ShopCode"`

	// 备注说明; 30个字符以内
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateAccountRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 账号名；需要是手机号
	Name *string `json:"Name,omitempty" name:"Name"`

	// 密码；需要是(`~!@#$%^&*()_+=-）中的至少两种且八位以上
	Password *string `json:"Password,omitempty" name:"Password"`

	// 客户门店编码
	ShopCode *string `json:"ShopCode,omitempty" name:"ShopCode"`

	// 备注说明; 30个字符以内
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "Name")
	delete(f, "Password")
	delete(f, "ShopCode")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAccountResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateAccountResponseParams `json:"Response"`
}

func (r *CreateAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFacePictureRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 人物类型（0表示普通顾客，1 白名单，2 表示黑名单，101表示集团白名单，102表示集团黑名单）
	PersonType *int64 `json:"PersonType,omitempty" name:"PersonType"`

	// 图片BASE编码
	Picture *string `json:"Picture,omitempty" name:"Picture"`

	// 图片名称
	PictureName *string `json:"PictureName,omitempty" name:"PictureName"`

	// 店铺ID，如果不填表示操作集团身份库
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 是否强制更新：为ture时会为用户创建一个新的指定PersonType的身份;目前这个参数已废弃，可不传
	IsForceUpload *bool `json:"IsForceUpload,omitempty" name:"IsForceUpload"`
}

type CreateFacePictureRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 人物类型（0表示普通顾客，1 白名单，2 表示黑名单，101表示集团白名单，102表示集团黑名单）
	PersonType *int64 `json:"PersonType,omitempty" name:"PersonType"`

	// 图片BASE编码
	Picture *string `json:"Picture,omitempty" name:"Picture"`

	// 图片名称
	PictureName *string `json:"PictureName,omitempty" name:"PictureName"`

	// 店铺ID，如果不填表示操作集团身份库
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 是否强制更新：为ture时会为用户创建一个新的指定PersonType的身份;目前这个参数已废弃，可不传
	IsForceUpload *bool `json:"IsForceUpload,omitempty" name:"IsForceUpload"`
}

func (r *CreateFacePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFacePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "PersonType")
	delete(f, "Picture")
	delete(f, "PictureName")
	delete(f, "ShopId")
	delete(f, "IsForceUpload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFacePictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFacePictureResponseParams struct {
	// 人物ID
	PersonId *int64 `json:"PersonId,omitempty" name:"PersonId"`

	// 0.正常建档 1.重复身份 2.未检测到人脸 3.检测到多个人脸 4.人脸大小过小 5.人脸质量不达标 6.其他错误
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 图片url
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFacePictureResponse struct {
	*tchttp.BaseResponse
	Response *CreateFacePictureResponseParams `json:"Response"`
}

func (r *CreateFacePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFacePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DailyTracePoint struct {
	// 轨迹日期
	TraceDate *string `json:"TraceDate,omitempty" name:"TraceDate"`

	// 轨迹点序列
	TracePointSet []*PersonTracePoint `json:"TracePointSet,omitempty" name:"TracePointSet"`
}

// Predefined struct for user
type DeletePersonFeatureRequestParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 顾客ID
	PersonId *int64 `json:"PersonId,omitempty" name:"PersonId"`
}

type DeletePersonFeatureRequest struct {
	*tchttp.BaseRequest
	
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 顾客ID
	PersonId *int64 `json:"PersonId,omitempty" name:"PersonId"`
}

func (r *DeletePersonFeatureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonFeatureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "PersonId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePersonFeatureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePersonFeatureResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeletePersonFeatureResponse struct {
	*tchttp.BaseResponse
	Response *DeletePersonFeatureResponseParams `json:"Response"`
}

func (r *DeletePersonFeatureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePersonFeatureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCameraPersonRequestParams struct {
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 摄像头id
	CameraId *int64 `json:"CameraId,omitempty" name:"CameraId"`

	// 拉取开始时间戳，单位秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 拉取结束时间戳，单位秒，不超过StartTime+10秒，超过默认为StartTime+10
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// pos机id
	PosId *string `json:"PosId,omitempty" name:"PosId"`

	// 拉取图片数，默认为1，最大为3
	Num *int64 `json:"Num,omitempty" name:"Num"`

	// 是否需要base64的图片，0-不需要，1-需要，默认0
	IsNeedPic *int64 `json:"IsNeedPic,omitempty" name:"IsNeedPic"`
}

type DescribeCameraPersonRequest struct {
	*tchttp.BaseRequest
	
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 摄像头id
	CameraId *int64 `json:"CameraId,omitempty" name:"CameraId"`

	// 拉取开始时间戳，单位秒
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 拉取结束时间戳，单位秒，不超过StartTime+10秒，超过默认为StartTime+10
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// pos机id
	PosId *string `json:"PosId,omitempty" name:"PosId"`

	// 拉取图片数，默认为1，最大为3
	Num *int64 `json:"Num,omitempty" name:"Num"`

	// 是否需要base64的图片，0-不需要，1-需要，默认0
	IsNeedPic *int64 `json:"IsNeedPic,omitempty" name:"IsNeedPic"`
}

func (r *DescribeCameraPersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCameraPersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "CameraId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PosId")
	delete(f, "Num")
	delete(f, "IsNeedPic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCameraPersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCameraPersonResponseParams struct {
	// 集团id
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺id
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 摄像机id
	CameraId *int64 `json:"CameraId,omitempty" name:"CameraId"`

	// pos机id
	PosId *string `json:"PosId,omitempty" name:"PosId"`

	// 抓取的顾客信息
	Infos []*CameraPersonInfo `json:"Infos,omitempty" name:"Infos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeCameraPersonResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCameraPersonResponseParams `json:"Response"`
}

func (r *DescribeCameraPersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCameraPersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPersonArrivedMallRequestParams struct {
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeClusterPersonArrivedMallRequest struct {
	*tchttp.BaseRequest
	
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeClusterPersonArrivedMallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPersonArrivedMallRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MallId")
	delete(f, "PersonId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPersonArrivedMallRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPersonArrivedMallResponseParams struct {
	// 卖场系统编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 卖场客户编码
	MallCode *string `json:"MallCode,omitempty" name:"MallCode"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 到场信息
	ArrivedMallSet []*ArrivedMallInfo `json:"ArrivedMallSet,omitempty" name:"ArrivedMallSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterPersonArrivedMallResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterPersonArrivedMallResponseParams `json:"Response"`
}

func (r *DescribeClusterPersonArrivedMallResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPersonArrivedMallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPersonTraceRequestParams struct {
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeClusterPersonTraceRequest struct {
	*tchttp.BaseRequest
	
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeClusterPersonTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPersonTraceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MallId")
	delete(f, "PersonId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterPersonTraceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterPersonTraceResponseParams struct {
	// 卖场系统编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 卖场用户编码
	MallCode *string `json:"MallCode,omitempty" name:"MallCode"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 轨迹序列
	TracePointSet []*DailyTracePoint `json:"TracePointSet,omitempty" name:"TracePointSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterPersonTraceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterPersonTraceResponseParams `json:"Response"`
}

func (r *DescribeClusterPersonTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterPersonTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFaceIdByTempIdRequestParams struct {
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 临时id
	TempId *string `json:"TempId,omitempty" name:"TempId"`

	// 摄像头id
	CameraId *int64 `json:"CameraId,omitempty" name:"CameraId"`

	// pos机id
	PosId *string `json:"PosId,omitempty" name:"PosId"`

	// 图片url过期时间：在当前时间+PictureExpires秒后，图片url无法继续正常访问；单位s；默认值1*24*60*60（1天）
	PictureExpires *int64 `json:"PictureExpires,omitempty" name:"PictureExpires"`
}

type DescribeFaceIdByTempIdRequest struct {
	*tchttp.BaseRequest
	
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 临时id
	TempId *string `json:"TempId,omitempty" name:"TempId"`

	// 摄像头id
	CameraId *int64 `json:"CameraId,omitempty" name:"CameraId"`

	// pos机id
	PosId *string `json:"PosId,omitempty" name:"PosId"`

	// 图片url过期时间：在当前时间+PictureExpires秒后，图片url无法继续正常访问；单位s；默认值1*24*60*60（1天）
	PictureExpires *int64 `json:"PictureExpires,omitempty" name:"PictureExpires"`
}

func (r *DescribeFaceIdByTempIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFaceIdByTempIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "TempId")
	delete(f, "CameraId")
	delete(f, "PosId")
	delete(f, "PictureExpires")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFaceIdByTempIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFaceIdByTempIdResponseParams struct {
	// 集团id
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺id
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 摄像机id
	CameraId *int64 `json:"CameraId,omitempty" name:"CameraId"`

	// pos机id
	PosId *string `json:"PosId,omitempty" name:"PosId"`

	// 请求的临时id
	TempId *string `json:"TempId,omitempty" name:"TempId"`

	// 临时id对应的face id
	FaceId *int64 `json:"FaceId,omitempty" name:"FaceId"`

	// 顾客属性信息
	PersonInfo *PersonInfo `json:"PersonInfo,omitempty" name:"PersonInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFaceIdByTempIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFaceIdByTempIdResponseParams `json:"Response"`
}

func (r *DescribeFaceIdByTempIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFaceIdByTempIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryNetworkInfoRequestParams struct {
	// 请求时间戳
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取，为0则拉取集团全部店铺当前
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 拉取开始日期，格式：2018-09-05
	StartDay *string `json:"StartDay,omitempty" name:"StartDay"`

	// 拉取结束日期，格式L:2018-09-05，超过StartDay 90天，按StartDay+90天算
	EndDay *string `json:"EndDay,omitempty" name:"EndDay"`

	// 拉取条数，默认10
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 拉取偏移，返回offset之后的数据
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeHistoryNetworkInfoRequest struct {
	*tchttp.BaseRequest
	
	// 请求时间戳
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取，为0则拉取集团全部店铺当前
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 拉取开始日期，格式：2018-09-05
	StartDay *string `json:"StartDay,omitempty" name:"StartDay"`

	// 拉取结束日期，格式L:2018-09-05，超过StartDay 90天，按StartDay+90天算
	EndDay *string `json:"EndDay,omitempty" name:"EndDay"`

	// 拉取条数，默认10
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 拉取偏移，返回offset之后的数据
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeHistoryNetworkInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryNetworkInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Time")
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "StartDay")
	delete(f, "EndDay")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHistoryNetworkInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHistoryNetworkInfoResponseParams struct {
	// 网络状态数据
	InstanceSet *NetworkHistoryInfo `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeHistoryNetworkInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHistoryNetworkInfoResponseParams `json:"Response"`
}

func (r *DescribeHistoryNetworkInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHistoryNetworkInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInfoRequestParams struct {
	// 请求时间戳
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取，不填则拉取集团全部店铺当前
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`
}

type DescribeNetworkInfoRequest struct {
	*tchttp.BaseRequest
	
	// 请求时间戳
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取，不填则拉取集团全部店铺当前
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`
}

func (r *DescribeNetworkInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Time")
	delete(f, "CompanyId")
	delete(f, "ShopId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkInfoResponseParams struct {
	// 网络状态详情
	InstanceSet *NetworkLastInfo `json:"InstanceSet,omitempty" name:"InstanceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNetworkInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkInfoResponseParams `json:"Response"`
}

func (r *DescribeNetworkInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonArrivedMallRequestParams struct {
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribePersonArrivedMallRequest struct {
	*tchttp.BaseRequest
	
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribePersonArrivedMallRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonArrivedMallRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MallId")
	delete(f, "PersonId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonArrivedMallRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonArrivedMallResponseParams struct {
	// 卖场系统编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 卖场用户编码
	MallCode *string `json:"MallCode,omitempty" name:"MallCode"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 到场轨迹
	ArrivedMallSet []*ArrivedMallInfo `json:"ArrivedMallSet,omitempty" name:"ArrivedMallSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePersonArrivedMallResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonArrivedMallResponseParams `json:"Response"`
}

func (r *DescribePersonArrivedMallResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonArrivedMallResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonInfoByFacePictureRequestParams struct {
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 人脸图片BASE编码
	Picture *string `json:"Picture,omitempty" name:"Picture"`
}

type DescribePersonInfoByFacePictureRequest struct {
	*tchttp.BaseRequest
	
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 人脸图片BASE编码
	Picture *string `json:"Picture,omitempty" name:"Picture"`
}

func (r *DescribePersonInfoByFacePictureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonInfoByFacePictureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "Picture")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonInfoByFacePictureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonInfoByFacePictureResponseParams struct {
	// 集团id
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺id
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 顾客face id
	PersonId *int64 `json:"PersonId,omitempty" name:"PersonId"`

	// 顾客底图url
	PictureUrl *string `json:"PictureUrl,omitempty" name:"PictureUrl"`

	// 顾客类型（0表示普通顾客，1 白名单，2 表示黑名单，101表示集团白名单，102表示集团黑名单）
	PersonType *int64 `json:"PersonType,omitempty" name:"PersonType"`

	// 顾客首次进店时间
	FirstVisitTime *string `json:"FirstVisitTime,omitempty" name:"FirstVisitTime"`

	// 顾客历史到访次数
	VisitTimes *int64 `json:"VisitTimes,omitempty" name:"VisitTimes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePersonInfoByFacePictureResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonInfoByFacePictureResponseParams `json:"Response"`
}

func (r *DescribePersonInfoByFacePictureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonInfoByFacePictureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonInfoRequestParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 起始ID，第一次拉取时StartPersonId传0，后续送入的值为上一页最后一条数据项的PersonId
	StartPersonId *uint64 `json:"StartPersonId,omitempty" name:"StartPersonId"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 图片url过期时间：在当前时间+PictureExpires秒后，图片url无法继续正常访问；单位s；默认值1*24*60*60（1天）
	PictureExpires *uint64 `json:"PictureExpires,omitempty" name:"PictureExpires"`

	// 身份类型(0表示普通顾客，1 白名单，2 表示黑名单）
	PersonType *uint64 `json:"PersonType,omitempty" name:"PersonType"`
}

type DescribePersonInfoRequest struct {
	*tchttp.BaseRequest
	
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 起始ID，第一次拉取时StartPersonId传0，后续送入的值为上一页最后一条数据项的PersonId
	StartPersonId *uint64 `json:"StartPersonId,omitempty" name:"StartPersonId"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 图片url过期时间：在当前时间+PictureExpires秒后，图片url无法继续正常访问；单位s；默认值1*24*60*60（1天）
	PictureExpires *uint64 `json:"PictureExpires,omitempty" name:"PictureExpires"`

	// 身份类型(0表示普通顾客，1 白名单，2 表示黑名单）
	PersonType *uint64 `json:"PersonType,omitempty" name:"PersonType"`
}

func (r *DescribePersonInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "StartPersonId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "PictureExpires")
	delete(f, "PersonType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonInfoResponseParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 用户信息
	PersonInfoSet []*PersonInfo `json:"PersonInfoSet,omitempty" name:"PersonInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePersonInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonInfoResponseParams `json:"Response"`
}

func (r *DescribePersonInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonRequestParams struct {
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 查询偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询数量，默认20，最大查询数量100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribePersonRequest struct {
	*tchttp.BaseRequest
	
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 查询偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 查询数量，默认20，最大查询数量100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePersonRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MallId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonResponseParams struct {
	// 总计客户数量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 客户信息
	PersonSet []*PersonProfile `json:"PersonSet,omitempty" name:"PersonSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePersonResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonResponseParams `json:"Response"`
}

func (r *DescribePersonResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonTraceDetailRequestParams struct {
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 轨迹编码
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`
}

type DescribePersonTraceDetailRequest struct {
	*tchttp.BaseRequest
	
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 轨迹编码
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`
}

func (r *DescribePersonTraceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonTraceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MallId")
	delete(f, "PersonId")
	delete(f, "TraceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonTraceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonTraceDetailResponseParams struct {
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 轨迹编码
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 轨迹点坐标序列
	CoordinateSet []*PersonCoordinate `json:"CoordinateSet,omitempty" name:"CoordinateSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePersonTraceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonTraceDetailResponseParams `json:"Response"`
}

func (r *DescribePersonTraceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonTraceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonTraceRequestParams struct {
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribePersonTraceRequest struct {
	*tchttp.BaseRequest
	
	// 卖场编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 查询开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribePersonTraceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonTraceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MallId")
	delete(f, "PersonId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonTraceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonTraceResponseParams struct {
	// 卖场系统编码
	MallId *string `json:"MallId,omitempty" name:"MallId"`

	// 卖场用户编码
	MallCode *string `json:"MallCode,omitempty" name:"MallCode"`

	// 客户编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 轨迹列表
	TraceRouteSet []*PersonTraceRoute `json:"TraceRouteSet,omitempty" name:"TraceRouteSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePersonTraceResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonTraceResponseParams `json:"Response"`
}

func (r *DescribePersonTraceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonVisitInfoRequestParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 开始日期，格式yyyy-MM-dd，已废弃，请使用StartDateTime
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd，已废弃，请使用EndDateTime
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 图片url过期时间：在当前时间+PictureExpires秒后，图片url无法继续正常访问；单位s；默认值1*24*60*60（1天）
	PictureExpires *uint64 `json:"PictureExpires,omitempty" name:"PictureExpires"`

	// 开始时间，格式yyyy-MM-dd HH:mm:ss
	StartDateTime *string `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// 结束时间，格式yyyy-MM-dd HH:mm:ss
	EndDateTime *string `json:"EndDateTime,omitempty" name:"EndDateTime"`
}

type DescribePersonVisitInfoRequest struct {
	*tchttp.BaseRequest
	
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 开始日期，格式yyyy-MM-dd，已废弃，请使用StartDateTime
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd，已废弃，请使用EndDateTime
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 图片url过期时间：在当前时间+PictureExpires秒后，图片url无法继续正常访问；单位s；默认值1*24*60*60（1天）
	PictureExpires *uint64 `json:"PictureExpires,omitempty" name:"PictureExpires"`

	// 开始时间，格式yyyy-MM-dd HH:mm:ss
	StartDateTime *string `json:"StartDateTime,omitempty" name:"StartDateTime"`

	// 结束时间，格式yyyy-MM-dd HH:mm:ss
	EndDateTime *string `json:"EndDateTime,omitempty" name:"EndDateTime"`
}

func (r *DescribePersonVisitInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonVisitInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "PictureExpires")
	delete(f, "StartDateTime")
	delete(f, "EndDateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePersonVisitInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePersonVisitInfoResponseParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 用户到访明细
	PersonVisitInfoSet []*PersonVisitInfo `json:"PersonVisitInfoSet,omitempty" name:"PersonVisitInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribePersonVisitInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribePersonVisitInfoResponseParams `json:"Response"`
}

func (r *DescribePersonVisitInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePersonVisitInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShopHourTrafficInfoRequestParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式：yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式：yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeShopHourTrafficInfoRequest struct {
	*tchttp.BaseRequest
	
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式：yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式：yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeShopHourTrafficInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShopHourTrafficInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShopHourTrafficInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShopHourTrafficInfoResponseParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 查询结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 分时客流信息
	ShopHourTrafficInfoSet []*ShopHourTrafficInfo `json:"ShopHourTrafficInfoSet,omitempty" name:"ShopHourTrafficInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeShopHourTrafficInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShopHourTrafficInfoResponseParams `json:"Response"`
}

func (r *DescribeShopHourTrafficInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShopHourTrafficInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShopInfoRequestParams struct {
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeShopInfoRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeShopInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShopInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShopInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShopInfoResponseParams struct {
	// 门店总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 门店列表信息
	ShopInfoSet []*ShopInfo `json:"ShopInfoSet,omitempty" name:"ShopInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeShopInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShopInfoResponseParams `json:"Response"`
}

func (r *DescribeShopInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShopInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShopTrafficInfoRequestParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 介绍日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeShopTrafficInfoRequest struct {
	*tchttp.BaseRequest
	
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 介绍日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeShopTrafficInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShopTrafficInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeShopTrafficInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeShopTrafficInfoResponseParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 查询结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 客流信息列表
	ShopDayTrafficInfoSet []*ShopDayTrafficInfo `json:"ShopDayTrafficInfoSet,omitempty" name:"ShopDayTrafficInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeShopTrafficInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeShopTrafficInfoResponseParams `json:"Response"`
}

func (r *DescribeShopTrafficInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeShopTrafficInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrajectoryDataRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 限制返回数据的最大条数，最大 400（负数代为 400）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 顾客性别顾虑，0是男，1是女，其它代表不分性别
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`
}

type DescribeTrajectoryDataRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 限制返回数据的最大条数，最大 400（负数代为 400）
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 顾客性别顾虑，0是男，1是女，其它代表不分性别
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`
}

func (r *DescribeTrajectoryDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrajectoryDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Limit")
	delete(f, "Gender")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrajectoryDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrajectoryDataResponseParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 总人数
	TotalPerson *int64 `json:"TotalPerson,omitempty" name:"TotalPerson"`

	// 总动迹数目
	TotalTrajectory *int64 `json:"TotalTrajectory,omitempty" name:"TotalTrajectory"`

	// 返回动迹中的总人数
	Person *int64 `json:"Person,omitempty" name:"Person"`

	// 返回动迹的数目
	Trajectory *int64 `json:"Trajectory,omitempty" name:"Trajectory"`

	// 返回动迹的具体信息
	Data []*TrajectorySunData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTrajectoryDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrajectoryDataResponseParams `json:"Response"`
}

func (r *DescribeTrajectoryDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrajectoryDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowAgeInfoByZoneIdRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeZoneFlowAgeInfoByZoneIdRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeZoneFlowAgeInfoByZoneIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowAgeInfoByZoneIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "ZoneId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneFlowAgeInfoByZoneIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowAgeInfoByZoneIdResponseParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 区域名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 当前年龄段占比
	Data []*float64 `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneFlowAgeInfoByZoneIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneFlowAgeInfoByZoneIdResponseParams `json:"Response"`
}

func (r *DescribeZoneFlowAgeInfoByZoneIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowAgeInfoByZoneIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowAndStayTimeRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeZoneFlowAndStayTimeRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeZoneFlowAndStayTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowAndStayTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneFlowAndStayTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowAndStayTimeResponseParams struct {
	// 集团id
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺id
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 各区域人流数目和停留时长
	Data []*ZoneFlowAndAvrStayTime `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneFlowAndStayTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneFlowAndStayTimeResponseParams `json:"Response"`
}

func (r *DescribeZoneFlowAndStayTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowAndStayTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowDailyByZoneIdRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeZoneFlowDailyByZoneIdRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeZoneFlowDailyByZoneIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowDailyByZoneIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "ZoneId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneFlowDailyByZoneIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowDailyByZoneIdResponseParams struct {
	// 集团id
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺id
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 区域名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 每日人流量
	Data []*ZoneDayFlow `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneFlowDailyByZoneIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneFlowDailyByZoneIdResponseParams `json:"Response"`
}

func (r *DescribeZoneFlowDailyByZoneIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowDailyByZoneIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowGenderAvrStayTimeByZoneIdRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeZoneFlowGenderAvrStayTimeByZoneIdRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeZoneFlowGenderAvrStayTimeByZoneIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowGenderAvrStayTimeByZoneIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "ZoneId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneFlowGenderAvrStayTimeByZoneIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowGenderAvrStayTimeByZoneIdResponseParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 区域名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 不同年龄段男女停留时间（返回格式为数组，从第 1 个到最后一个数据，年龄段分别为 0-17，18 - 23,  24 - 30, 31 - 40, 41 - 50, 51 - 60, 61 - 100）
	Data []*ZoneAgeGroupAvrStayTime `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneFlowGenderAvrStayTimeByZoneIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneFlowGenderAvrStayTimeByZoneIdResponseParams `json:"Response"`
}

func (r *DescribeZoneFlowGenderAvrStayTimeByZoneIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowGenderAvrStayTimeByZoneIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowGenderInfoByZoneIdRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeZoneFlowGenderInfoByZoneIdRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeZoneFlowGenderInfoByZoneIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowGenderInfoByZoneIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "ZoneId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneFlowGenderInfoByZoneIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowGenderInfoByZoneIdResponseParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 区域名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 男性占比
	MalePercent *float64 `json:"MalePercent,omitempty" name:"MalePercent"`

	// 女性占比
	FemalePercent *float64 `json:"FemalePercent,omitempty" name:"FemalePercent"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneFlowGenderInfoByZoneIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneFlowGenderInfoByZoneIdResponseParams `json:"Response"`
}

func (r *DescribeZoneFlowGenderInfoByZoneIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowGenderInfoByZoneIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowHourlyByZoneIdRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeZoneFlowHourlyByZoneIdRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeZoneFlowHourlyByZoneIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowHourlyByZoneIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "ZoneId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneFlowHourlyByZoneIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneFlowHourlyByZoneIdResponseParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 区域名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 各个分时人流量
	Data []*ZoneHourFlow `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneFlowHourlyByZoneIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneFlowHourlyByZoneIdResponseParams `json:"Response"`
}

func (r *DescribeZoneFlowHourlyByZoneIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneFlowHourlyByZoneIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneTrafficInfoRequestParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeZoneTrafficInfoRequest struct {
	*tchttp.BaseRequest
	
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 开始日期，格式yyyy-MM-dd
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 结束日期，格式yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeZoneTrafficInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneTrafficInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZoneTrafficInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZoneTrafficInfoResponseParams struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 查询结果总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 区域客流信息列表
	ZoneTrafficInfoSet []*ZoneTrafficInfo `json:"ZoneTrafficInfoSet,omitempty" name:"ZoneTrafficInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeZoneTrafficInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZoneTrafficInfoResponseParams `json:"Response"`
}

func (r *DescribeZoneTrafficInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZoneTrafficInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenderAgeTrafficDetail struct {
	// 性别: 0男1女
	Gender *uint64 `json:"Gender,omitempty" name:"Gender"`

	// 年龄区间，枚举值：0-17、18-23、24-30、31-40、41-50、51-60、>60
	AgeGap *string `json:"AgeGap,omitempty" name:"AgeGap"`

	// 客流量
	TrafficCount *uint64 `json:"TrafficCount,omitempty" name:"TrafficCount"`
}

type HourTrafficInfoDetail struct {
	// 小时 取值为：0，1，2，3，4，5，6，7，8，9，10，11，12，13，14，15，16，17，18，19，20，21，22，23
	Hour *uint64 `json:"Hour,omitempty" name:"Hour"`

	// 分时客流量
	HourTrafficTotalCount *uint64 `json:"HourTrafficTotalCount,omitempty" name:"HourTrafficTotalCount"`
}

// Predefined struct for user
type ModifyPersonFeatureInfoRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 需要修改的顾客id
	PersonId *int64 `json:"PersonId,omitempty" name:"PersonId"`

	// 图片BASE编码
	Picture *string `json:"Picture,omitempty" name:"Picture"`

	// 图片名称（尽量不要重复）
	PictureName *string `json:"PictureName,omitempty" name:"PictureName"`

	// 人物类型，仅能操作黑白名单顾客（1 白名单，2 表示黑名单，101表示集团白名单，102表示集团黑名单）
	PersonType *int64 `json:"PersonType,omitempty" name:"PersonType"`

	// 店铺ID，如果不填表示操作集团身份库
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`
}

type ModifyPersonFeatureInfoRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 需要修改的顾客id
	PersonId *int64 `json:"PersonId,omitempty" name:"PersonId"`

	// 图片BASE编码
	Picture *string `json:"Picture,omitempty" name:"Picture"`

	// 图片名称（尽量不要重复）
	PictureName *string `json:"PictureName,omitempty" name:"PictureName"`

	// 人物类型，仅能操作黑白名单顾客（1 白名单，2 表示黑名单，101表示集团白名单，102表示集团黑名单）
	PersonType *int64 `json:"PersonType,omitempty" name:"PersonType"`

	// 店铺ID，如果不填表示操作集团身份库
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`
}

func (r *ModifyPersonFeatureInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonFeatureInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "PersonId")
	delete(f, "Picture")
	delete(f, "PictureName")
	delete(f, "PersonType")
	delete(f, "ShopId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPersonFeatureInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonFeatureInfoResponseParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺ID，如果不填表示操作集团身份库
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 请求的顾客id
	PersonId *int64 `json:"PersonId,omitempty" name:"PersonId"`

	// 图片实际绑定person_id，可能与请求的person_id不同，以此id为准
	PersonIdBind *int64 `json:"PersonIdBind,omitempty" name:"PersonIdBind"`

	// 请求的顾客类型
	PersonType *int64 `json:"PersonType,omitempty" name:"PersonType"`

	// 与请求的person_id类型相同、与请求图片特征相似的一个或多个person_id，需要额外确认这些id是否是同一个人
	SimilarPersonIds []*int64 `json:"SimilarPersonIds,omitempty" name:"SimilarPersonIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPersonFeatureInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPersonFeatureInfoResponseParams `json:"Response"`
}

func (r *ModifyPersonFeatureInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonFeatureInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonTagInfoRequestParams struct {
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取，为0则拉取集团全部店铺当前
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 需要设置的顾客信息，批量设置最大为10个
	Tags []*PersonTagInfo `json:"Tags,omitempty" name:"Tags"`
}

type ModifyPersonTagInfoRequest struct {
	*tchttp.BaseRequest
	
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取，为0则拉取集团全部店铺当前
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 需要设置的顾客信息，批量设置最大为10个
	Tags []*PersonTagInfo `json:"Tags,omitempty" name:"Tags"`
}

func (r *ModifyPersonTagInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonTagInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPersonTagInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonTagInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPersonTagInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPersonTagInfoResponseParams `json:"Response"`
}

func (r *ModifyPersonTagInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonTagInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonTypeRequestParams struct {
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 顾客ID
	PersonId *uint64 `json:"PersonId,omitempty" name:"PersonId"`

	// 身份类型(0表示普通顾客，1 白名单，2 表示黑名单）
	PersonType *uint64 `json:"PersonType,omitempty" name:"PersonType"`

	// 身份子类型:
	// PersonType=0时(普通顾客)，0普通顾客
	// PersonType=1时(白名单)，0店员，1商场人员，2其他类型人员，3区域经理，4注册会员，5VIP用户
	// PersonType=2时(黑名单)，0普通黑名单，1小偷)
	PersonSubType *uint64 `json:"PersonSubType,omitempty" name:"PersonSubType"`
}

type ModifyPersonTypeRequest struct {
	*tchttp.BaseRequest
	
	// 集团ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 顾客ID
	PersonId *uint64 `json:"PersonId,omitempty" name:"PersonId"`

	// 身份类型(0表示普通顾客，1 白名单，2 表示黑名单）
	PersonType *uint64 `json:"PersonType,omitempty" name:"PersonType"`

	// 身份子类型:
	// PersonType=0时(普通顾客)，0普通顾客
	// PersonType=1时(白名单)，0店员，1商场人员，2其他类型人员，3区域经理，4注册会员，5VIP用户
	// PersonType=2时(黑名单)，0普通黑名单，1小偷)
	PersonSubType *uint64 `json:"PersonSubType,omitempty" name:"PersonSubType"`
}

func (r *ModifyPersonTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "ShopId")
	delete(f, "PersonId")
	delete(f, "PersonType")
	delete(f, "PersonSubType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPersonTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPersonTypeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyPersonTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPersonTypeResponseParams `json:"Response"`
}

func (r *ModifyPersonTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPersonTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkAndShopInfo struct {
	// 集团id
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺id
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 店铺省份
	Province *string `json:"Province,omitempty" name:"Province"`

	// 店铺城市
	City *string `json:"City,omitempty" name:"City"`

	// 店铺名
	ShopName *string `json:"ShopName,omitempty" name:"ShopName"`

	// 上传带宽，单位Mb/s，-1：未知
	Upload *float64 `json:"Upload,omitempty" name:"Upload"`

	// 下载带宽，单位Mb/s，-1：未知
	Download *float64 `json:"Download,omitempty" name:"Download"`

	// 最小延迟，单位ms，-1：未知
	MinRtt *float64 `json:"MinRtt,omitempty" name:"MinRtt"`

	// 平均延迟，单位ms，-1：未知
	AvgRtt *float64 `json:"AvgRtt,omitempty" name:"AvgRtt"`

	// 最大延迟，单位ms，-1：未知
	MaxRtt *float64 `json:"MaxRtt,omitempty" name:"MaxRtt"`

	// 平均偏差延迟，单位ms，-1：未知
	MdevRtt *float64 `json:"MdevRtt,omitempty" name:"MdevRtt"`

	// 丢包率百分比，-1：未知
	Loss *float64 `json:"Loss,omitempty" name:"Loss"`

	// 更新时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 上报网络状态设备
	Mac *string `json:"Mac,omitempty" name:"Mac"`
}

type NetworkHistoryInfo struct {
	// 总数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 集团id
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 店铺id
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 店铺省份
	Province *string `json:"Province,omitempty" name:"Province"`

	// 店铺城市
	City *string `json:"City,omitempty" name:"City"`

	// 店铺名称
	ShopName *string `json:"ShopName,omitempty" name:"ShopName"`

	// 网络信息
	Infos []*NetworkInfo `json:"Infos,omitempty" name:"Infos"`
}

type NetworkInfo struct {
	// 上传带宽，单位Mb/s，-1：未知
	Upload *float64 `json:"Upload,omitempty" name:"Upload"`

	// 下载带宽，单位Mb/s，-1：未知
	Download *float64 `json:"Download,omitempty" name:"Download"`

	// 最小延迟，单位ms，-1：未知
	MinRtt *float64 `json:"MinRtt,omitempty" name:"MinRtt"`

	// 平均延迟，单位ms，-1：未知
	AvgRtt *float64 `json:"AvgRtt,omitempty" name:"AvgRtt"`

	// 最大延迟，单位ms，-1：未知
	MaxRtt *float64 `json:"MaxRtt,omitempty" name:"MaxRtt"`

	// 平均偏差延迟，单位ms，-1：未知
	MdevRtt *float64 `json:"MdevRtt,omitempty" name:"MdevRtt"`

	// 丢包率百分比，-1：未知
	Loss *float64 `json:"Loss,omitempty" name:"Loss"`

	// 更新时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 上报网络状态设备
	Mac *string `json:"Mac,omitempty" name:"Mac"`
}

type NetworkLastInfo struct {
	// 总数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 网络状态
	Infos []*NetworkAndShopInfo `json:"Infos,omitempty" name:"Infos"`
}

type PersonCoordinate struct {
	// CAD图X坐标
	CADX *float64 `json:"CADX,omitempty" name:"CADX"`

	// CAD图Y坐标
	CADY *float64 `json:"CADY,omitempty" name:"CADY"`

	// 抓拍时间点
	CapTime *string `json:"CapTime,omitempty" name:"CapTime"`

	// 抓拍图片
	CapPic *string `json:"CapPic,omitempty" name:"CapPic"`

	// 卖场区域类型
	MallAreaType *int64 `json:"MallAreaType,omitempty" name:"MallAreaType"`

	// 坐标编号
	PosId *int64 `json:"PosId,omitempty" name:"PosId"`

	// 门店编号
	ShopId *int64 `json:"ShopId,omitempty" name:"ShopId"`

	// 事件
	Event *string `json:"Event,omitempty" name:"Event"`
}

type PersonInfo struct {
	// 用户ID
	PersonId *uint64 `json:"PersonId,omitempty" name:"PersonId"`

	// 人脸图片Base64内容，已弃用，返回默认空值
	PersonPicture *string `json:"PersonPicture,omitempty" name:"PersonPicture"`

	// 性别：0男1女
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 年龄
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 身份类型（0表示普通顾客，1 白名单，2 表示黑名单）
	PersonType *int64 `json:"PersonType,omitempty" name:"PersonType"`

	// 人脸图片Url，在有效期内可以访问下载
	PersonPictureUrl *string `json:"PersonPictureUrl,omitempty" name:"PersonPictureUrl"`

	// 身份子类型:
	// PersonType=0时(普通顾客)，0普通顾客
	// PersonType=1时(白名单)，0店员，1商场人员，2其他类型人员，3区域经理，4注册用户，5VIP用户
	// PersonType=2时(黑名单)，0普通黑名单，1小偷)
	PersonSubType *int64 `json:"PersonSubType,omitempty" name:"PersonSubType"`

	// 到访次数，-1表示未知
	VisitTimes *int64 `json:"VisitTimes,omitempty" name:"VisitTimes"`

	// 到访天数，-1表示未知
	VisitDays *int64 `json:"VisitDays,omitempty" name:"VisitDays"`
}

type PersonProfile struct {
	// 客人编码
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 性别
	Gender *uint64 `json:"Gender,omitempty" name:"Gender"`

	// 年龄
	Age *uint64 `json:"Age,omitempty" name:"Age"`

	// 首次到场时间
	FirstArrivedTime *string `json:"FirstArrivedTime,omitempty" name:"FirstArrivedTime"`

	// 来访次数
	ArrivedCount *uint64 `json:"ArrivedCount,omitempty" name:"ArrivedCount"`

	// 客户图片
	PicUrl *string `json:"PicUrl,omitempty" name:"PicUrl"`

	// 置信度
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`
}

type PersonTagInfo struct {
	// 顾客原类型
	OldType *int64 `json:"OldType,omitempty" name:"OldType"`

	// 顾客新类型
	NewType *int64 `json:"NewType,omitempty" name:"NewType"`

	// 顾客face id
	PersonId *int64 `json:"PersonId,omitempty" name:"PersonId"`
}

type PersonTracePoint struct {
	// 卖场区域编码
	MallAreaId *uint64 `json:"MallAreaId,omitempty" name:"MallAreaId"`

	// 门店编码
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 卖场区域类型
	MallAreaType *uint64 `json:"MallAreaType,omitempty" name:"MallAreaType"`

	// 轨迹事件
	TraceEventType *uint64 `json:"TraceEventType,omitempty" name:"TraceEventType"`

	// 轨迹事件发生时间点
	TraceEventTime *string `json:"TraceEventTime,omitempty" name:"TraceEventTime"`

	// 抓拍图片
	CapPic *string `json:"CapPic,omitempty" name:"CapPic"`

	// 购物袋类型
	ShoppingBagType *uint64 `json:"ShoppingBagType,omitempty" name:"ShoppingBagType"`

	// 购物袋数量
	ShoppingBagCount *uint64 `json:"ShoppingBagCount,omitempty" name:"ShoppingBagCount"`
}

type PersonTraceRoute struct {
	// 轨迹编码
	TraceId *string `json:"TraceId,omitempty" name:"TraceId"`

	// 轨迹点序列
	TracePointSet []*PersonTracePoint `json:"TracePointSet,omitempty" name:"TracePointSet"`
}

type PersonVisitInfo struct {
	// 用户ID
	PersonId *uint64 `json:"PersonId,omitempty" name:"PersonId"`

	// 用户到访ID
	VisitId *uint64 `json:"VisitId,omitempty" name:"VisitId"`

	// 到访时间：Unix时间戳
	InTime *uint64 `json:"InTime,omitempty" name:"InTime"`

	// 抓拍到的头像Base64内容，已弃用，返回默认空值
	CapturedPicture *string `json:"CapturedPicture,omitempty" name:"CapturedPicture"`

	// 口罩类型：0不戴口罩，1戴口罩
	MaskType *uint64 `json:"MaskType,omitempty" name:"MaskType"`

	// 眼镜类型：0不戴眼镜，1普通眼镜 , 2墨镜
	GlassType *uint64 `json:"GlassType,omitempty" name:"GlassType"`

	// 发型：0 短发,  1长发
	HairType *uint64 `json:"HairType,omitempty" name:"HairType"`

	// 抓拍到的头像Url，在有效期内可以访问下载
	CapturedPictureUrl *string `json:"CapturedPictureUrl,omitempty" name:"CapturedPictureUrl"`

	// 抓拍头像的场景图信息
	SceneInfo *SceneInfo `json:"SceneInfo,omitempty" name:"SceneInfo"`
}

// Predefined struct for user
type RegisterCallbackRequestParams struct {
	// 集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 通知回调地址，完整url，示例（http://youmall.tencentcloudapi.com/）
	BackUrl *string `json:"BackUrl,omitempty" name:"BackUrl"`

	// 请求时间戳
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 是否需要顾客图片，1-需要图片，其它-不需要图片
	NeedFacePic *uint64 `json:"NeedFacePic,omitempty" name:"NeedFacePic"`
}

type RegisterCallbackRequest struct {
	*tchttp.BaseRequest
	
	// 集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 通知回调地址，完整url，示例（http://youmall.tencentcloudapi.com/）
	BackUrl *string `json:"BackUrl,omitempty" name:"BackUrl"`

	// 请求时间戳
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 是否需要顾客图片，1-需要图片，其它-不需要图片
	NeedFacePic *uint64 `json:"NeedFacePic,omitempty" name:"NeedFacePic"`
}

func (r *RegisterCallbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCallbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CompanyId")
	delete(f, "BackUrl")
	delete(f, "Time")
	delete(f, "NeedFacePic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterCallbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterCallbackResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterCallbackResponse struct {
	*tchttp.BaseResponse
	Response *RegisterCallbackResponseParams `json:"Response"`
}

func (r *RegisterCallbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SceneInfo struct {
	// 场景图
	ScenePictureURL *string `json:"ScenePictureURL,omitempty" name:"ScenePictureURL"`

	// 抓拍头像左上角X坐标在场景图中的像素点位置
	HeadX *int64 `json:"HeadX,omitempty" name:"HeadX"`

	// 抓拍头像左上角Y坐标在场景图中的像素点位置
	HeadY *int64 `json:"HeadY,omitempty" name:"HeadY"`

	// 抓拍头像在场景图中占有的像素宽度
	HeadWidth *int64 `json:"HeadWidth,omitempty" name:"HeadWidth"`

	// 抓拍头像在场景图中占有的像素高度
	HeadHeight *int64 `json:"HeadHeight,omitempty" name:"HeadHeight"`
}

type ShopDayTrafficInfo struct {
	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 客流量
	DayTrafficTotalCount *uint64 `json:"DayTrafficTotalCount,omitempty" name:"DayTrafficTotalCount"`

	// 性别年龄分组下的客流信息
	GenderAgeTrafficDetailSet []*GenderAgeTrafficDetail `json:"GenderAgeTrafficDetailSet,omitempty" name:"GenderAgeTrafficDetailSet"`
}

type ShopHourTrafficInfo struct {
	// 日期，格式yyyy-MM-dd
	Date *string `json:"Date,omitempty" name:"Date"`

	// 分时客流详细信息
	HourTrafficInfoDetailSet []*HourTrafficInfoDetail `json:"HourTrafficInfoDetailSet,omitempty" name:"HourTrafficInfoDetailSet"`
}

type ShopInfo struct {
	// 公司ID
	CompanyId *string `json:"CompanyId,omitempty" name:"CompanyId"`

	// 门店ID
	ShopId *uint64 `json:"ShopId,omitempty" name:"ShopId"`

	// 门店名称
	ShopName *string `json:"ShopName,omitempty" name:"ShopName"`

	// 客户门店编码
	ShopCode *string `json:"ShopCode,omitempty" name:"ShopCode"`

	// 省
	Province *string `json:"Province,omitempty" name:"Province"`

	// 市
	City *string `json:"City,omitempty" name:"City"`

	// 公司名称
	CompanyName *string `json:"CompanyName,omitempty" name:"CompanyName"`
}

type TrajectorySunData struct {
	// 区域动线，形如 x-x-x-x-x，其中 x 为区域 ID
	Zones *string `json:"Zones,omitempty" name:"Zones"`

	// 该动线出现次数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 该动线平均停留时间（秒）
	AvgStayTime *int64 `json:"AvgStayTime,omitempty" name:"AvgStayTime"`
}

type ZoneAgeGroupAvrStayTime struct {
	// 男性平均停留时间
	MaleAvrStayTime *float64 `json:"MaleAvrStayTime,omitempty" name:"MaleAvrStayTime"`

	// 女性平均停留时间
	FemaleAvrStayTime *float64 `json:"FemaleAvrStayTime,omitempty" name:"FemaleAvrStayTime"`
}

type ZoneDayFlow struct {
	// 日期，如 2018-08-6
	Day *string `json:"Day,omitempty" name:"Day"`

	// 客流量
	FlowCount *int64 `json:"FlowCount,omitempty" name:"FlowCount"`
}

type ZoneFlowAndAvrStayTime struct {
	// 区域id
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 区域名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 人流量
	FlowCount *uint64 `json:"FlowCount,omitempty" name:"FlowCount"`

	// 平均停留时长
	AvrStayTime *uint64 `json:"AvrStayTime,omitempty" name:"AvrStayTime"`
}

type ZoneHourFlow struct {
	// 分时 0~23
	Hour *int64 `json:"Hour,omitempty" name:"Hour"`

	// 客流量
	FlowCount *int64 `json:"FlowCount,omitempty" name:"FlowCount"`
}

type ZoneTrafficInfo struct {
	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 门店区域客流详细信息
	ZoneTrafficInfoDetailSet []*ZoneTrafficInfoDetail `json:"ZoneTrafficInfoDetailSet,omitempty" name:"ZoneTrafficInfoDetailSet"`
}

type ZoneTrafficInfoDetail struct {
	// 区域ID
	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 区域名称
	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`

	// 客流量
	TrafficTotalCount *uint64 `json:"TrafficTotalCount,omitempty" name:"TrafficTotalCount"`

	// 平均停留时间
	AvgStayTime *uint64 `json:"AvgStayTime,omitempty" name:"AvgStayTime"`
}