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

package v20220105

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AppInfoItem struct {
	// App包名
	AppPackage *string `json:"AppPackage,omitnil" name:"AppPackage"`

	// App名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// App版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppVersion *string `json:"AppVersion,omitnil" name:"AppVersion"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// App隐私诊断报告下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitnil" name:"ReportUrl"`

	// App隐私诊断报告名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTitle *string `json:"ReportTitle,omitnil" name:"ReportTitle"`

	// App诊断堆栈报告下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	BehaviorUrl *string `json:"BehaviorUrl,omitnil" name:"BehaviorUrl"`

	// App诊断堆栈报告名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BehaviorTitle *string `json:"BehaviorTitle,omitnil" name:"BehaviorTitle"`

	// 诊断高风险项数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighRiskCount *int64 `json:"HighRiskCount,omitnil" name:"HighRiskCount"`

	// 隐私申明文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivacyTextName *string `json:"PrivacyTextName,omitnil" name:"PrivacyTextName"`

	// 软件MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftwareMD5 *string `json:"SoftwareMD5,omitnil" name:"SoftwareMD5"`

	// 隐私文本MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivacyTextMD5 *string `json:"PrivacyTextMD5,omitnil" name:"PrivacyTextMD5"`
}

type AppTaskData struct {
	// 任务ID
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 0:默认值(待检测/待咨询), 1.检测中, 2:待评估, 3:评估中, 4:任务完成/咨询完成, 5:任务失败, 6:咨询中;
	TaskStatus *int64 `json:"TaskStatus,omitnil" name:"TaskStatus"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskErrMsg *string `json:"TaskErrMsg,omitnil" name:"TaskErrMsg"`

	// 任务来源,0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android)
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用信息
	AppInfo *AppInfoItem `json:"AppInfo,omitnil" name:"AppInfo"`

	// 任务启动时间
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 任务完成时间(更新时间)
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 联系人信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactName *string `json:"ContactName,omitnil" name:"ContactName"`
}

// Predefined struct for user
type CreateAppScanTaskRepeatRequestParams struct {
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 原诊断任务ID
	OrgTaskID *string `json:"OrgTaskID,omitnil" name:"OrgTaskID"`

	// App包名
	AppPackage *string `json:"AppPackage,omitnil" name:"AppPackage"`

	// 上传的文件ID(任务来源为1时必填)
	FileID *string `json:"FileID,omitnil" name:"FileID"`

	// 软件下载链接地址(任务来源为2时必填)
	AppDownloadUrl *string `json:"AppDownloadUrl,omitnil" name:"AppDownloadUrl"`

	// 隐私文本下载地址(任务来源为2时必填)
	PrivacyTextUrl *string `json:"PrivacyTextUrl,omitnil" name:"PrivacyTextUrl"`

	// 应用名称
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// 隐私申明文件名称
	PrivacyTextName *string `json:"PrivacyTextName,omitnil" name:"PrivacyTextName"`

	// 软件Sha1值(PrivacyTextMD5不为空时必填)
	AppSha1 *string `json:"AppSha1,omitnil" name:"AppSha1"`

	// 隐私申明文本md5(AppSha1不为空时必填)
	PrivacyTextMD5 *string `json:"PrivacyTextMD5,omitnil" name:"PrivacyTextMD5"`
}

type CreateAppScanTaskRepeatRequest struct {
	*tchttp.BaseRequest
	
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 原诊断任务ID
	OrgTaskID *string `json:"OrgTaskID,omitnil" name:"OrgTaskID"`

	// App包名
	AppPackage *string `json:"AppPackage,omitnil" name:"AppPackage"`

	// 上传的文件ID(任务来源为1时必填)
	FileID *string `json:"FileID,omitnil" name:"FileID"`

	// 软件下载链接地址(任务来源为2时必填)
	AppDownloadUrl *string `json:"AppDownloadUrl,omitnil" name:"AppDownloadUrl"`

	// 隐私文本下载地址(任务来源为2时必填)
	PrivacyTextUrl *string `json:"PrivacyTextUrl,omitnil" name:"PrivacyTextUrl"`

	// 应用名称
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// 隐私申明文件名称
	PrivacyTextName *string `json:"PrivacyTextName,omitnil" name:"PrivacyTextName"`

	// 软件Sha1值(PrivacyTextMD5不为空时必填)
	AppSha1 *string `json:"AppSha1,omitnil" name:"AppSha1"`

	// 隐私申明文本md5(AppSha1不为空时必填)
	PrivacyTextMD5 *string `json:"PrivacyTextMD5,omitnil" name:"PrivacyTextMD5"`
}

func (r *CreateAppScanTaskRepeatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppScanTaskRepeatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Platform")
	delete(f, "TaskType")
	delete(f, "OrgTaskID")
	delete(f, "AppPackage")
	delete(f, "FileID")
	delete(f, "AppDownloadUrl")
	delete(f, "PrivacyTextUrl")
	delete(f, "AppName")
	delete(f, "PrivacyTextName")
	delete(f, "AppSha1")
	delete(f, "PrivacyTextMD5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppScanTaskRepeatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppScanTaskRepeatResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitnil" name:"Result"`

	// 任务id
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAppScanTaskRepeatResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppScanTaskRepeatResponseParams `json:"Response"`
}

func (r *CreateAppScanTaskRepeatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppScanTaskRepeatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppScanTaskRequestParams struct {
	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// App包名
	AppPackage *string `json:"AppPackage,omitnil" name:"AppPackage"`

	// App名称(任务来源为2时必填)
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// App版本
	AppVersion *string `json:"AppVersion,omitnil" name:"AppVersion"`

	// 上传的软件文件ID(任务来源为1时必填)
	FileID *string `json:"FileID,omitnil" name:"FileID"`

	// 软件下载链接地址(任务来源为2时必填)
	AppDownloadUrl *string `json:"AppDownloadUrl,omitnil" name:"AppDownloadUrl"`

	// 隐私文本下载地址(任务来源为2时必填)
	PrivacyTextUrl *string `json:"PrivacyTextUrl,omitnil" name:"PrivacyTextUrl"`

	// 联系人信息
	ContactName *string `json:"ContactName,omitnil" name:"ContactName"`

	// 联系电话
	TelNumber *string `json:"TelNumber,omitnil" name:"TelNumber"`

	// 公司邮箱
	Email *string `json:"Email,omitnil" name:"Email"`

	// 公司名称
	CorpName *string `json:"CorpName,omitnil" name:"CorpName"`

	// 商务对接人员
	SalesPerson *string `json:"SalesPerson,omitnil" name:"SalesPerson"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否同意隐私条款，0:不同意(默认), 1:同意
	IsAgreePrivacy *int64 `json:"IsAgreePrivacy,omitnil" name:"IsAgreePrivacy"`

	// 隐私申明文件名称
	PrivacyTextName *string `json:"PrivacyTextName,omitnil" name:"PrivacyTextName"`

	// 软件Sha1值(PrivacyTextMD5不为空时必填)
	AppSha1 *string `json:"AppSha1,omitnil" name:"AppSha1"`

	// 隐私申明文本md5(AppSha1不为空时必填)
	PrivacyTextMD5 *string `json:"PrivacyTextMD5,omitnil" name:"PrivacyTextMD5"`
}

type CreateAppScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// App包名
	AppPackage *string `json:"AppPackage,omitnil" name:"AppPackage"`

	// App名称(任务来源为2时必填)
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// App版本
	AppVersion *string `json:"AppVersion,omitnil" name:"AppVersion"`

	// 上传的软件文件ID(任务来源为1时必填)
	FileID *string `json:"FileID,omitnil" name:"FileID"`

	// 软件下载链接地址(任务来源为2时必填)
	AppDownloadUrl *string `json:"AppDownloadUrl,omitnil" name:"AppDownloadUrl"`

	// 隐私文本下载地址(任务来源为2时必填)
	PrivacyTextUrl *string `json:"PrivacyTextUrl,omitnil" name:"PrivacyTextUrl"`

	// 联系人信息
	ContactName *string `json:"ContactName,omitnil" name:"ContactName"`

	// 联系电话
	TelNumber *string `json:"TelNumber,omitnil" name:"TelNumber"`

	// 公司邮箱
	Email *string `json:"Email,omitnil" name:"Email"`

	// 公司名称
	CorpName *string `json:"CorpName,omitnil" name:"CorpName"`

	// 商务对接人员
	SalesPerson *string `json:"SalesPerson,omitnil" name:"SalesPerson"`

	// 备注信息
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 是否同意隐私条款，0:不同意(默认), 1:同意
	IsAgreePrivacy *int64 `json:"IsAgreePrivacy,omitnil" name:"IsAgreePrivacy"`

	// 隐私申明文件名称
	PrivacyTextName *string `json:"PrivacyTextName,omitnil" name:"PrivacyTextName"`

	// 软件Sha1值(PrivacyTextMD5不为空时必填)
	AppSha1 *string `json:"AppSha1,omitnil" name:"AppSha1"`

	// 隐私申明文本md5(AppSha1不为空时必填)
	PrivacyTextMD5 *string `json:"PrivacyTextMD5,omitnil" name:"PrivacyTextMD5"`
}

func (r *CreateAppScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "Source")
	delete(f, "Platform")
	delete(f, "AppPackage")
	delete(f, "AppName")
	delete(f, "AppVersion")
	delete(f, "FileID")
	delete(f, "AppDownloadUrl")
	delete(f, "PrivacyTextUrl")
	delete(f, "ContactName")
	delete(f, "TelNumber")
	delete(f, "Email")
	delete(f, "CorpName")
	delete(f, "SalesPerson")
	delete(f, "Remark")
	delete(f, "IsAgreePrivacy")
	delete(f, "PrivacyTextName")
	delete(f, "AppSha1")
	delete(f, "PrivacyTextMD5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppScanTaskResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitnil" name:"Result"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateAppScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppScanTaskResponseParams `json:"Response"`
}

func (r *CreateAppScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelTaskReportUrlRequestParams struct {
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1: iOS，2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务id
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 报告类型, 0:诊断报告, 1:堆栈报告, 2:视频证据(预留), 3:报告json结果
	ReportType *int64 `json:"ReportType,omitnil" name:"ReportType"`

	// 子渠道APP MD5值
	AppMD5 *string `json:"AppMD5,omitnil" name:"AppMD5"`
}

type DescribeChannelTaskReportUrlRequest struct {
	*tchttp.BaseRequest
	
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1: iOS，2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务id
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 报告类型, 0:诊断报告, 1:堆栈报告, 2:视频证据(预留), 3:报告json结果
	ReportType *int64 `json:"ReportType,omitnil" name:"ReportType"`

	// 子渠道APP MD5值
	AppMD5 *string `json:"AppMD5,omitnil" name:"AppMD5"`
}

func (r *DescribeChannelTaskReportUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelTaskReportUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Platform")
	delete(f, "TaskID")
	delete(f, "TaskType")
	delete(f, "ReportType")
	delete(f, "AppMD5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeChannelTaskReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeChannelTaskReportUrlResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitnil" name:"Result"`

	// 诊断报告/堆栈信息/报告json结果下载链接
	ReportUrl *string `json:"ReportUrl,omitnil" name:"ReportUrl"`

	// 诊断报告/堆栈/报告json结果的名称
	ReportTitle *string `json:"ReportTitle,omitnil" name:"ReportTitle"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeChannelTaskReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeChannelTaskReportUrlResponseParams `json:"Response"`
}

func (r *DescribeChannelTaskReportUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeChannelTaskReportUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileTicketRequestParams struct {
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`
}

type DescribeFileTicketRequest struct {
	*tchttp.BaseRequest
	
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`
}

func (r *DescribeFileTicketRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileTicketRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileTicketRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileTicketResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitnil" name:"Result"`

	// 上传url(任务来源为2时:Post方法（100:apk,101:txt）, 任务来源为1时:put方法)
	UploadUrl *string `json:"UploadUrl,omitnil" name:"UploadUrl"`

	// 上传url鉴权信息(任务来源为1时上传需要, Authorization参数值)
	// 注意：此字段可能返回 null，表示取不到有效值。
	UploadSign *string `json:"UploadSign,omitnil" name:"UploadSign"`

	// 上传文件ID(任务来源为1时提交诊断任务需要)
	// 注意：此字段可能返回 null，表示取不到有效值。
	FildID *string `json:"FildID,omitnil" name:"FildID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeFileTicketResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileTicketResponseParams `json:"Response"`
}

func (r *DescribeFileTicketResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileTicketResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageInfoRequestParams struct {
	// 资源计费项名称(为空时，则根据Source，TaskType和Platform进行查询)
	PriceName *string `json:"PriceName,omitnil" name:"PriceName"`

	// 任务类型, 0:基础版, 1:专家版
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 应用平台, 0:android
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`
}

type DescribeResourceUsageInfoRequest struct {
	*tchttp.BaseRequest
	
	// 资源计费项名称(为空时，则根据Source，TaskType和Platform进行查询)
	PriceName *string `json:"PriceName,omitnil" name:"PriceName"`

	// 任务类型, 0:基础版, 1:专家版
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 应用平台, 0:android
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`
}

func (r *DescribeResourceUsageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUsageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PriceName")
	delete(f, "TaskType")
	delete(f, "Platform")
	delete(f, "Source")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceUsageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageInfoResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义，暂时未定
	Result *int64 `json:"Result,omitnil" name:"Result"`

	// 资源使用信息
	Data *ResourceUsageInfoData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeResourceUsageInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceUsageInfoResponseParams `json:"Response"`
}

func (r *DescribeResourceUsageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceUsageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskListRequestParams struct {
	// 任务来源, -1:所有, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务状态,可多值查询,例如:"1,2,3" 0:默认值(待检测/待咨询), 1.检测中, 2:待评估, 3:评估中, 4:任务完成/咨询完成, 5:任务失败, 6:咨询中;
	TaskStatuses *string `json:"TaskStatuses,omitnil" name:"TaskStatuses"`

	// 任务类型,可多值查询,采用逗号分隔,例如:"0,1" 0:基础版, 1:专家版, 2:本地化
	TaskTypes *string `json:"TaskTypes,omitnil" name:"TaskTypes"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 页码大小
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 应用名称或小程序名称(可选参数)
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// 查询时间范围, 查询开始时间(2021-09-30 或 2021-09-30 10:57:34)
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询时间范围, 查询结束时间(2021-09-30 或 2021-09-30 10:57:34)
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeScanTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 任务来源, -1:所有, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务状态,可多值查询,例如:"1,2,3" 0:默认值(待检测/待咨询), 1.检测中, 2:待评估, 3:评估中, 4:任务完成/咨询完成, 5:任务失败, 6:咨询中;
	TaskStatuses *string `json:"TaskStatuses,omitnil" name:"TaskStatuses"`

	// 任务类型,可多值查询,采用逗号分隔,例如:"0,1" 0:基础版, 1:专家版, 2:本地化
	TaskTypes *string `json:"TaskTypes,omitnil" name:"TaskTypes"`

	// 页码
	PageNo *int64 `json:"PageNo,omitnil" name:"PageNo"`

	// 页码大小
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// 应用名称或小程序名称(可选参数)
	AppName *string `json:"AppName,omitnil" name:"AppName"`

	// 查询时间范围, 查询开始时间(2021-09-30 或 2021-09-30 10:57:34)
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 查询时间范围, 查询结束时间(2021-09-30 或 2021-09-30 10:57:34)
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}

func (r *DescribeScanTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Platform")
	delete(f, "TaskStatuses")
	delete(f, "TaskTypes")
	delete(f, "PageNo")
	delete(f, "PageSize")
	delete(f, "AppName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskListResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitnil" name:"Result"`

	// 任务总数量
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 诊断任务数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*AppTaskData `json:"Data,omitnil" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScanTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanTaskListResponseParams `json:"Response"`
}

func (r *DescribeScanTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskReportUrlRequestParams struct {
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务id
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 报告类型, 0:诊断报告, 1:堆栈报告, 2:视频证据(预留), 3:报告json结果
	ReportType *int64 `json:"ReportType,omitnil" name:"ReportType"`
}

type DescribeScanTaskReportUrlRequest struct {
	*tchttp.BaseRequest
	
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务id
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`

	// 报告类型, 0:诊断报告, 1:堆栈报告, 2:视频证据(预留), 3:报告json结果
	ReportType *int64 `json:"ReportType,omitnil" name:"ReportType"`
}

func (r *DescribeScanTaskReportUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskReportUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Platform")
	delete(f, "TaskID")
	delete(f, "TaskType")
	delete(f, "ReportType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskReportUrlResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitnil" name:"Result"`

	// 诊断报告/堆栈信息/报告json结果下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitnil" name:"ReportUrl"`

	// 诊断报告/堆栈/报告json结果的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTitle *string `json:"ReportTitle,omitnil" name:"ReportTitle"`

	// 诊断json结果内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportResult *string `json:"ReportResult,omitnil" name:"ReportResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScanTaskReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanTaskReportUrlResponseParams `json:"Response"`
}

func (r *DescribeScanTaskReportUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskReportUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskStatusRequestParams struct {
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务id
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`
}

type DescribeScanTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android), 3:app漏洞扫描;
	Source *int64 `json:"Source,omitnil" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitnil" name:"Platform"`

	// 任务id
	TaskID *string `json:"TaskID,omitnil" name:"TaskID"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitnil" name:"TaskType"`
}

func (r *DescribeScanTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Source")
	delete(f, "Platform")
	delete(f, "TaskID")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskStatusResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitnil" name:"Result"`

	// 0:默认值(待检测/待咨询), 1.检测中,  4:任务完成/咨询完成, 5:任务失败, 6:咨询中;
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 诊断失败的错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitnil" name:"ErrMsg"`

	// 任务流详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowSteps []*TaskFlowStepsInfo `json:"FlowSteps,omitnil" name:"FlowSteps"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeScanTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeScanTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceUsageInfoData struct {
	// 资源计费项名称
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 资源总数
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 未使用资源数
	UnusedCount *int64 `json:"UnusedCount,omitnil" name:"UnusedCount"`
}

type TaskFlowStepsInfo struct {
	// 流程编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowNo *string `json:"FlowNo,omitnil" name:"FlowNo"`

	// 流程名称
	FlowName *string `json:"FlowName,omitnil" name:"FlowName"`

	// 流程状态, 其他值:进行中, 2:成功, 3:失败
	FlowStatus *int64 `json:"FlowStatus,omitnil" name:"FlowStatus"`

	// 流程状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowStateDesc *string `json:"FlowStateDesc,omitnil" name:"FlowStateDesc"`

	// 流程启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 流程完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`
}