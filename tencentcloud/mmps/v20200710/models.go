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

package v20200710

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AppInfoItem struct {
	// 小程序apiiid
	AppPackage *string `json:"AppPackage,omitempty" name:"AppPackage"`

	// 小程序应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 小程序应用版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 小程序隐私诊断报告下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`

	// 小程序隐私诊断报告名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTitle *string `json:"ReportTitle,omitempty" name:"ReportTitle"`

	// 小程序隐私诊断堆栈报告下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	BehaviorUrl *string `json:"BehaviorUrl,omitempty" name:"BehaviorUrl"`

	// 小程序隐私诊断堆栈报告名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BehaviorTitle *string `json:"BehaviorTitle,omitempty" name:"BehaviorTitle"`

	// 诊断风险项数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighRiskCount *int64 `json:"HighRiskCount,omitempty" name:"HighRiskCount"`

	// 隐私申明文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivacyTextName *string `json:"PrivacyTextName,omitempty" name:"PrivacyTextName"`

	// 软件MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	SoftwareMD5 *string `json:"SoftwareMD5,omitempty" name:"SoftwareMD5"`

	// 隐私文本MD5
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivacyTextMD5 *string `json:"PrivacyTextMD5,omitempty" name:"PrivacyTextMD5"`
}

type AppTaskData struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 0:默认值(待检测/待咨询), 1.检测中, 2:待评估, 3:评估中, 4:任务完成/咨询完成, 5:任务失败, 6:咨询中;
	TaskStatus *int64 `json:"TaskStatus,omitempty" name:"TaskStatus"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskErrMsg *string `json:"TaskErrMsg,omitempty" name:"TaskErrMsg"`

	// 任务来源,0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 应用信息
	AppInfo *AppInfoItem `json:"AppInfo,omitempty" name:"AppInfo"`

	// 任务启动时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务完成时间(更新时间)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 联系人信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`
}

// Predefined struct for user
type CreateAppScanTaskRepeatRequestParams struct {
	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 小程序AppID
	AppPackage *string `json:"AppPackage,omitempty" name:"AppPackage"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 原诊断任务ID
	OrgTaskID *string `json:"OrgTaskID,omitempty" name:"OrgTaskID"`
}

type CreateAppScanTaskRepeatRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 小程序AppID
	AppPackage *string `json:"AppPackage,omitempty" name:"AppPackage"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 原诊断任务ID
	OrgTaskID *string `json:"OrgTaskID,omitempty" name:"OrgTaskID"`
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
	delete(f, "TaskType")
	delete(f, "Source")
	delete(f, "AppPackage")
	delete(f, "Platform")
	delete(f, "OrgTaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppScanTaskRepeatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppScanTaskRepeatResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 小程序AppID
	AppPackage *string `json:"AppPackage,omitempty" name:"AppPackage"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 小程序名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 小程序版本
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 联系人信息
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 联系电话
	TelNumber *string `json:"TelNumber,omitempty" name:"TelNumber"`

	// 公司名称
	CorpName *string `json:"CorpName,omitempty" name:"CorpName"`

	// 商务对接人员
	SalesPerson *string `json:"SalesPerson,omitempty" name:"SalesPerson"`

	// 公司邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

type CreateAppScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 小程序AppID
	AppPackage *string `json:"AppPackage,omitempty" name:"AppPackage"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 小程序名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 小程序版本
	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`

	// 联系人信息
	ContactName *string `json:"ContactName,omitempty" name:"ContactName"`

	// 联系电话
	TelNumber *string `json:"TelNumber,omitempty" name:"TelNumber"`

	// 公司名称
	CorpName *string `json:"CorpName,omitempty" name:"CorpName"`

	// 商务对接人员
	SalesPerson *string `json:"SalesPerson,omitempty" name:"SalesPerson"`

	// 公司邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
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
	delete(f, "AppPackage")
	delete(f, "Platform")
	delete(f, "AppName")
	delete(f, "AppVersion")
	delete(f, "ContactName")
	delete(f, "TelNumber")
	delete(f, "CorpName")
	delete(f, "SalesPerson")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppScanTaskResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateFlySecMiniAppProfessionalScanTaskRequestParams struct {
	// 小程序AppID
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 小程序名称
	MiniAppName *string `json:"MiniAppName,omitempty" name:"MiniAppName"`

	// 诊断模式 2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 公司名称
	CorpName *string `json:"CorpName,omitempty" name:"CorpName"`

	// 手机号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 电子邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateFlySecMiniAppProfessionalScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 小程序AppID
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 小程序名称
	MiniAppName *string `json:"MiniAppName,omitempty" name:"MiniAppName"`

	// 诊断模式 2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 公司名称
	CorpName *string `json:"CorpName,omitempty" name:"CorpName"`

	// 手机号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 电子邮箱
	Email *string `json:"Email,omitempty" name:"Email"`

	// 备注信息
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateFlySecMiniAppProfessionalScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlySecMiniAppProfessionalScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MiniAppID")
	delete(f, "MiniAppName")
	delete(f, "Mode")
	delete(f, "CorpName")
	delete(f, "Mobile")
	delete(f, "Email")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlySecMiniAppProfessionalScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlySecMiniAppProfessionalScanTaskResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlySecMiniAppProfessionalScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlySecMiniAppProfessionalScanTaskResponseParams `json:"Response"`
}

func (r *CreateFlySecMiniAppProfessionalScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlySecMiniAppProfessionalScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlySecMiniAppScanTaskRepeatRequestParams struct {
	// 小程序AppID
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 诊断模式 1:基础诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 原任务id
	OrgTaskID *string `json:"OrgTaskID,omitempty" name:"OrgTaskID"`

	// 小程序测试账号(自有账号体系需提供,其他情况不需要)
	MiniAppTestAccount *string `json:"MiniAppTestAccount,omitempty" name:"MiniAppTestAccount"`

	// 小程序测试密码(自有账号体系需提供,其他情况不需要)
	MiniAppTestPwd *string `json:"MiniAppTestPwd,omitempty" name:"MiniAppTestPwd"`

	// 诊断扫描版本 0:正式版 1:体验版
	ScanVersion *int64 `json:"ScanVersion,omitempty" name:"ScanVersion"`
}

type CreateFlySecMiniAppScanTaskRepeatRequest struct {
	*tchttp.BaseRequest
	
	// 小程序AppID
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 诊断模式 1:基础诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 原任务id
	OrgTaskID *string `json:"OrgTaskID,omitempty" name:"OrgTaskID"`

	// 小程序测试账号(自有账号体系需提供,其他情况不需要)
	MiniAppTestAccount *string `json:"MiniAppTestAccount,omitempty" name:"MiniAppTestAccount"`

	// 小程序测试密码(自有账号体系需提供,其他情况不需要)
	MiniAppTestPwd *string `json:"MiniAppTestPwd,omitempty" name:"MiniAppTestPwd"`

	// 诊断扫描版本 0:正式版 1:体验版
	ScanVersion *int64 `json:"ScanVersion,omitempty" name:"ScanVersion"`
}

func (r *CreateFlySecMiniAppScanTaskRepeatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlySecMiniAppScanTaskRepeatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MiniAppID")
	delete(f, "Mode")
	delete(f, "OrgTaskID")
	delete(f, "MiniAppTestAccount")
	delete(f, "MiniAppTestPwd")
	delete(f, "ScanVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlySecMiniAppScanTaskRepeatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlySecMiniAppScanTaskRepeatResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlySecMiniAppScanTaskRepeatResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlySecMiniAppScanTaskRepeatResponseParams `json:"Response"`
}

func (r *CreateFlySecMiniAppScanTaskRepeatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlySecMiniAppScanTaskRepeatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlySecMiniAppScanTaskRequestParams struct {
	// 小程序AppID
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 诊断模式 1:基础诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 小程序测试账号(自有账号体系需提供,其他情况不需要)
	MiniAppTestAccount *string `json:"MiniAppTestAccount,omitempty" name:"MiniAppTestAccount"`

	// 小程序测试密码(自有账号体系需提供,其他情况不需要)
	MiniAppTestPwd *string `json:"MiniAppTestPwd,omitempty" name:"MiniAppTestPwd"`

	// 小程序所属行业
	Industry *string `json:"Industry,omitempty" name:"Industry"`

	// 小程序调查问卷json字符串
	SurveyContent *string `json:"SurveyContent,omitempty" name:"SurveyContent"`

	// 手机号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 邮箱地址
	Email *string `json:"Email,omitempty" name:"Email"`

	// 商务合作接口人
	SalesPerson *string `json:"SalesPerson,omitempty" name:"SalesPerson"`

	// 诊断扫描版本 0:正式版 1:体验版
	ScanVersion *int64 `json:"ScanVersion,omitempty" name:"ScanVersion"`
}

type CreateFlySecMiniAppScanTaskRequest struct {
	*tchttp.BaseRequest
	
	// 小程序AppID
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 诊断模式 1:基础诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 小程序测试账号(自有账号体系需提供,其他情况不需要)
	MiniAppTestAccount *string `json:"MiniAppTestAccount,omitempty" name:"MiniAppTestAccount"`

	// 小程序测试密码(自有账号体系需提供,其他情况不需要)
	MiniAppTestPwd *string `json:"MiniAppTestPwd,omitempty" name:"MiniAppTestPwd"`

	// 小程序所属行业
	Industry *string `json:"Industry,omitempty" name:"Industry"`

	// 小程序调查问卷json字符串
	SurveyContent *string `json:"SurveyContent,omitempty" name:"SurveyContent"`

	// 手机号码
	Mobile *string `json:"Mobile,omitempty" name:"Mobile"`

	// 邮箱地址
	Email *string `json:"Email,omitempty" name:"Email"`

	// 商务合作接口人
	SalesPerson *string `json:"SalesPerson,omitempty" name:"SalesPerson"`

	// 诊断扫描版本 0:正式版 1:体验版
	ScanVersion *int64 `json:"ScanVersion,omitempty" name:"ScanVersion"`
}

func (r *CreateFlySecMiniAppScanTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlySecMiniAppScanTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MiniAppID")
	delete(f, "Mode")
	delete(f, "MiniAppTestAccount")
	delete(f, "MiniAppTestPwd")
	delete(f, "Industry")
	delete(f, "SurveyContent")
	delete(f, "Mobile")
	delete(f, "Email")
	delete(f, "SalesPerson")
	delete(f, "ScanVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFlySecMiniAppScanTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFlySecMiniAppScanTaskResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFlySecMiniAppScanTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateFlySecMiniAppScanTaskResponseParams `json:"Response"`
}

func (r *CreateFlySecMiniAppScanTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFlySecMiniAppScanTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicDiagnosisResourceUsageInfoRequestParams struct {
	// 诊断模式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
}

type DescribeBasicDiagnosisResourceUsageInfoRequest struct {
	*tchttp.BaseRequest
	
	// 诊断模式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`
}

func (r *DescribeBasicDiagnosisResourceUsageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDiagnosisResourceUsageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBasicDiagnosisResourceUsageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBasicDiagnosisResourceUsageInfoResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 资源类型
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 资源未使用次数
	UnusedCount *int64 `json:"UnusedCount,omitempty" name:"UnusedCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBasicDiagnosisResourceUsageInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBasicDiagnosisResourceUsageInfoResponseParams `json:"Response"`
}

func (r *DescribeBasicDiagnosisResourceUsageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBasicDiagnosisResourceUsageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppReportUrlRequestParams struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 小程序appid
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 诊断方式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 诊断报告类型 0:基础诊断报告，1:总裁版诊断报告，2:诊断报告json结果
	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`
}

type DescribeFlySecMiniAppReportUrlRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 小程序appid
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 诊断方式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 诊断报告类型 0:基础诊断报告，1:总裁版诊断报告，2:诊断报告json结果
	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`
}

func (r *DescribeFlySecMiniAppReportUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppReportUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	delete(f, "MiniAppID")
	delete(f, "Mode")
	delete(f, "ReportType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlySecMiniAppReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppReportUrlResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 诊断报告下载链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlySecMiniAppReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlySecMiniAppReportUrlResponseParams `json:"Response"`
}

func (r *DescribeFlySecMiniAppReportUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppReportUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppScanReportListRequestParams struct {
	// 小程序AppID
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 诊断方式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 诊断状态 -1:查询全部, 0:排队中, 1:成功, 2:失败, 3:进行中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 查询数量, 0:查询所有, 其他值:最近几次的诊断数量
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 小程序版本
	MiniAppVersion *string `json:"MiniAppVersion,omitempty" name:"MiniAppVersion"`
}

type DescribeFlySecMiniAppScanReportListRequest struct {
	*tchttp.BaseRequest
	
	// 小程序AppID
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 诊断方式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 诊断状态 -1:查询全部, 0:排队中, 1:成功, 2:失败, 3:进行中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 查询数量, 0:查询所有, 其他值:最近几次的诊断数量
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 小程序版本
	MiniAppVersion *string `json:"MiniAppVersion,omitempty" name:"MiniAppVersion"`
}

func (r *DescribeFlySecMiniAppScanReportListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppScanReportListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MiniAppID")
	delete(f, "Mode")
	delete(f, "Status")
	delete(f, "Size")
	delete(f, "MiniAppVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlySecMiniAppScanReportListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppScanReportListResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 诊断报告数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*FlySecMiniAppReportData `json:"Data,omitempty" name:"Data"`

	// 诊断任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlySecMiniAppScanReportListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlySecMiniAppScanReportListResponseParams `json:"Response"`
}

func (r *DescribeFlySecMiniAppScanReportListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppScanReportListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppScanTaskListRequestParams struct {
	// 诊断方式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 诊断状态 -1:查询全部, 0:排队中, 1:成功, 2:失败, 3:进行中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 查询数量, 0:查询所有, 其他值:最近几次的诊断数量
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 小程序appid(为空的时候,则查询当前用户诊断的所有小程序)
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`
}

type DescribeFlySecMiniAppScanTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 诊断方式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 诊断状态 -1:查询全部, 0:排队中, 1:成功, 2:失败, 3:进行中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 查询数量, 0:查询所有, 其他值:最近几次的诊断数量
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 小程序appid(为空的时候,则查询当前用户诊断的所有小程序)
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`
}

func (r *DescribeFlySecMiniAppScanTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppScanTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Mode")
	delete(f, "Status")
	delete(f, "Size")
	delete(f, "MiniAppID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlySecMiniAppScanTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppScanTaskListResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 诊断任务数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*FlySecMiniAppTaskData `json:"Data,omitempty" name:"Data"`

	// 诊断任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlySecMiniAppScanTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlySecMiniAppScanTaskListResponseParams `json:"Response"`
}

func (r *DescribeFlySecMiniAppScanTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppScanTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppScanTaskParamRequestParams struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

type DescribeFlySecMiniAppScanTaskParamRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *DescribeFlySecMiniAppScanTaskParamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppScanTaskParamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlySecMiniAppScanTaskParamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppScanTaskParamResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 小程序AppID
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 诊断模式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 小程序测试账号(自有账号体系需提供,其他情况不需要)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniAppTestAccount *string `json:"MiniAppTestAccount,omitempty" name:"MiniAppTestAccount"`

	// 小程序测试密码(自有账号体系需提供,其他情况不需要)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniAppTestPwd *string `json:"MiniAppTestPwd,omitempty" name:"MiniAppTestPwd"`

	// 诊断扫描版本 0:正式版 1:体验版
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanVersion *int64 `json:"ScanVersion,omitempty" name:"ScanVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlySecMiniAppScanTaskParamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlySecMiniAppScanTaskParamResponseParams `json:"Response"`
}

func (r *DescribeFlySecMiniAppScanTaskParamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppScanTaskParamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppScanTaskStatusRequestParams struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

type DescribeFlySecMiniAppScanTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
}

func (r *DescribeFlySecMiniAppScanTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppScanTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFlySecMiniAppScanTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFlySecMiniAppScanTaskStatusResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 诊断状态, 0:排队中, 1:成功, 2:失败, 3:进行中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 诊断失败错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errno *int64 `json:"Errno,omitempty" name:"Errno"`

	// 小程序名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniAppName *string `json:"MiniAppName,omitempty" name:"MiniAppName"`

	// 小程序版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniAppVersion *string `json:"MiniAppVersion,omitempty" name:"MiniAppVersion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFlySecMiniAppScanTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFlySecMiniAppScanTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeFlySecMiniAppScanTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFlySecMiniAppScanTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageInfoRequestParams struct {

}

type DescribeResourceUsageInfoRequest struct {
	*tchttp.BaseRequest
	
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
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceUsageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceUsageInfoResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Ret *int64 `json:"Ret,omitempty" name:"Ret"`

	// 安全资源数据列表
	Data []*ResourceUsageInfoData `json:"Data,omitempty" name:"Data"`

	// 安全资源数量
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 任务来源, -1:所有, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 任务状态,可多值查询,例如:"1,2,3" 0:默认值(待检测/待咨询), 1.检测中, 2:待评估, 3:评估中, 4:任务完成/咨询完成, 5:任务失败, 6:咨询中;
	TaskStatuses *string `json:"TaskStatuses,omitempty" name:"TaskStatuses"`

	// 任务类型,可多值查询,采用逗号分隔,例如:"0,1" 0:基础版, 1:专家版, 2:本地化
	TaskTypes *string `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// 页码
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 页码大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 应用名称或小程序名称(可选参数)
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 查询时间范围, 查询开始时间(2021-09-30 或 2021-09-30 10:57:34)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询时间范围, 查询结束时间(2021-09-30 或 2021-09-30 10:57:34)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type DescribeScanTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 任务来源, -1:所有, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 任务状态,可多值查询,例如:"1,2,3" 0:默认值(待检测/待咨询), 1.检测中, 2:待评估, 3:评估中, 4:任务完成/咨询完成, 5:任务失败, 6:咨询中;
	TaskStatuses *string `json:"TaskStatuses,omitempty" name:"TaskStatuses"`

	// 任务类型,可多值查询,采用逗号分隔,例如:"0,1" 0:基础版, 1:专家版, 2:本地化
	TaskTypes *string `json:"TaskTypes,omitempty" name:"TaskTypes"`

	// 页码
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 页码大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 应用名称或小程序名称(可选参数)
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 查询时间范围, 查询开始时间(2021-09-30 或 2021-09-30 10:57:34)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询时间范围, 查询结束时间(2021-09-30 或 2021-09-30 10:57:34)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 诊断任务数据列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*AppTaskData `json:"Data,omitempty" name:"Data"`

	// 任务总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 报告类型, 0:诊断报告, 1:堆栈报告(预留), 2:视频证据(预留), 3:报告json结果
	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
}

type DescribeScanTaskReportUrlRequest struct {
	*tchttp.BaseRequest
	
	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`

	// 报告类型, 0:诊断报告, 1:堆栈报告(预留), 2:视频证据(预留), 3:报告json结果
	ReportType *int64 `json:"ReportType,omitempty" name:"ReportType"`

	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
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
	delete(f, "TaskID")
	delete(f, "Platform")
	delete(f, "ReportType")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskReportUrlResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 诊断报告/堆栈信息下载链接
	ReportUrl *string `json:"ReportUrl,omitempty" name:"ReportUrl"`

	// 诊断报告/堆栈名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTitle *string `json:"ReportTitle,omitempty" name:"ReportTitle"`

	// 诊断json结果内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportResult *string `json:"ReportResult,omitempty" name:"ReportResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`
}

type DescribeScanTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型, 0:基础版, 1:专家版, 2:本地化
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务来源, 0:小程序诊断, 1:预留字段(暂未使用), 2:app诊断(android);
	Source *int64 `json:"Source,omitempty" name:"Source"`

	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 应用平台, 0:android, 1:ios, 2:小程序
	Platform *int64 `json:"Platform,omitempty" name:"Platform"`
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
	delete(f, "TaskType")
	delete(f, "Source")
	delete(f, "TaskID")
	delete(f, "Platform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanTaskStatusResponseParams struct {
	// 返回值, 0:成功, 其他值请查看“返回值”定义
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 0:默认值(待检测/待咨询), 1.检测中,  4:任务完成/咨询完成, 5:任务失败, 6:咨询中;
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 诊断失败的错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 任务流详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowSteps []*TaskFlowStepsInfo `json:"FlowSteps,omitempty" name:"FlowSteps"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type FlySecMiniAppReportData struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 小程序appid
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 小程序名称
	MiniAppName *string `json:"MiniAppName,omitempty" name:"MiniAppName"`

	// 小程序版本
	MiniAppVersion *string `json:"MiniAppVersion,omitempty" name:"MiniAppVersion"`

	// 诊断模式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 诊断状态, 0:排队中, 1:成功, 2:失败, 3:进行中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 诊断时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 诊断得分
	RiskScore *string `json:"RiskScore,omitempty" name:"RiskScore"`

	// 诊断风险等级 1:高风险 2:中风险 3:低风险
	RiskLevel *int64 `json:"RiskLevel,omitempty" name:"RiskLevel"`

	// 诊断8大维度得分情况(每项总分100分)
	RiskItems *FlySecMiniAppRiskItems `json:"RiskItems,omitempty" name:"RiskItems"`
}

type FlySecMiniAppRiskItems struct {
	// 代码防护(基础诊断)
	RiskItem1Score *int64 `json:"RiskItem1Score,omitempty" name:"RiskItem1Score"`

	// 开发测试信息泄露(基础诊断)
	RiskItem2Score *int64 `json:"RiskItem2Score,omitempty" name:"RiskItem2Score"`

	// 编码规范(基础诊断)
	RiskItem3Score *int64 `json:"RiskItem3Score,omitempty" name:"RiskItem3Score"`

	// 配置风险(基础诊断)
	RiskItem4Score *int64 `json:"RiskItem4Score,omitempty" name:"RiskItem4Score"`

	// 账号安全(基础诊断)
	RiskItem5Score *int64 `json:"RiskItem5Score,omitempty" name:"RiskItem5Score"`

	// 用户信息安全(基础诊断)
	RiskItem6Score *int64 `json:"RiskItem6Score,omitempty" name:"RiskItem6Score"`

	// 内部信息泄露(基础诊断)
	RiskItem7Score *int64 `json:"RiskItem7Score,omitempty" name:"RiskItem7Score"`

	// 其他安全(基础诊断)
	RiskItem8Score *int64 `json:"RiskItem8Score,omitempty" name:"RiskItem8Score"`
}

type FlySecMiniAppTaskData struct {
	// 任务id
	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`

	// 小程序appid
	MiniAppID *string `json:"MiniAppID,omitempty" name:"MiniAppID"`

	// 小程序名称
	MiniAppName *string `json:"MiniAppName,omitempty" name:"MiniAppName"`

	// 小程序版本
	MiniAppVersion *string `json:"MiniAppVersion,omitempty" name:"MiniAppVersion"`

	// 诊断模式 1:基础诊断，2:深度诊断
	Mode *int64 `json:"Mode,omitempty" name:"Mode"`

	// 诊断时间
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 诊断状态, 0:排队中, 1:成功, 2:失败, 3:进行中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 诊断失败错误码
	Error *int64 `json:"Error,omitempty" name:"Error"`
}

type ResourceUsageInfoData struct {
	// 资源名称, 具体名称请查看产品配置
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源总数
	Total *int64 `json:"Total,omitempty" name:"Total"`

	// 资源未使用次数
	UnusedCount *int64 `json:"UnusedCount,omitempty" name:"UnusedCount"`
}

type TaskFlowStepsInfo struct {
	// 流程编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowNo *string `json:"FlowNo,omitempty" name:"FlowNo"`

	// 流程名称
	FlowName *string `json:"FlowName,omitempty" name:"FlowName"`

	// 流程状态, 其他值:进行中, 2:成功, 3:失败
	FlowStatus *int64 `json:"FlowStatus,omitempty" name:"FlowStatus"`

	// 流程状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowStateDesc *string `json:"FlowStateDesc,omitempty" name:"FlowStateDesc"`

	// 流程启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 流程完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}