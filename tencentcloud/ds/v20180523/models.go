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

package v20180523

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CheckVcodeRequestParams struct {
	// 模块名VerifyCode
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CheckVcode
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 帐号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`

	// 验证码
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

type CheckVcodeRequest struct {
	*tchttp.BaseRequest
	
	// 模块名VerifyCode
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CheckVcode
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 帐号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`

	// 验证码
	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

func (r *CheckVcodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckVcodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "AccountResId")
	delete(f, "ContractResId")
	delete(f, "VerifyCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckVcodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckVcodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckVcodeResponse struct {
	*tchttp.BaseResponse
	Response *CheckVcodeResponseParams `json:"Response"`
}

func (r *CheckVcodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckVcodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContractByUploadRequestParams struct {
	// 模块名ContractMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CreateContractByUpload
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 签署人信息
	SignInfos []*SignInfo `json:"SignInfos,omitempty" name:"SignInfos"`

	// 合同上传链接地址
	ContractFile *string `json:"ContractFile,omitempty" name:"ContractFile"`

	// 合同名称
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 备注
	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`

	// 合同发起方腾讯云帐号ID（由平台自动填写）
	Initiator *string `json:"Initiator,omitempty" name:"Initiator"`

	// 合同长时间未签署的过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type CreateContractByUploadRequest struct {
	*tchttp.BaseRequest
	
	// 模块名ContractMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CreateContractByUpload
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 签署人信息
	SignInfos []*SignInfo `json:"SignInfos,omitempty" name:"SignInfos"`

	// 合同上传链接地址
	ContractFile *string `json:"ContractFile,omitempty" name:"ContractFile"`

	// 合同名称
	ContractName *string `json:"ContractName,omitempty" name:"ContractName"`

	// 备注
	Remarks *string `json:"Remarks,omitempty" name:"Remarks"`

	// 合同发起方腾讯云帐号ID（由平台自动填写）
	Initiator *string `json:"Initiator,omitempty" name:"Initiator"`

	// 合同长时间未签署的过期时间
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

func (r *CreateContractByUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContractByUploadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "SignInfos")
	delete(f, "ContractFile")
	delete(f, "ContractName")
	delete(f, "Remarks")
	delete(f, "Initiator")
	delete(f, "ExpireTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateContractByUploadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateContractByUploadResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateContractByUploadResponse struct {
	*tchttp.BaseResponse
	Response *CreateContractByUploadResponseParams `json:"Response"`
}

func (r *CreateContractByUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateContractByUploadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnterpriseAccountRequestParams struct {
	// 模块名AccountMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CreateEnterpriseAccount
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 企业用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 企业用户证件类型，8代表营业执照，详情请见常见问题
	IdentType *int64 `json:"IdentType,omitempty" name:"IdentType"`

	// 企业用户营业执照号码
	IdentNo *string `json:"IdentNo,omitempty" name:"IdentNo"`

	// 企业联系人手机号
	MobilePhone *string `json:"MobilePhone,omitempty" name:"MobilePhone"`

	// 经办人姓名
	TransactorName *string `json:"TransactorName,omitempty" name:"TransactorName"`

	// 经办人证件类型，0代表身份证
	TransactorIdentType *int64 `json:"TransactorIdentType,omitempty" name:"TransactorIdentType"`

	// 经办人证件号码
	TransactorIdentNo *string `json:"TransactorIdentNo,omitempty" name:"TransactorIdentNo"`

	// 经办人手机号
	TransactorPhone *string `json:"TransactorPhone,omitempty" name:"TransactorPhone"`

	// 企业联系人邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

type CreateEnterpriseAccountRequest struct {
	*tchttp.BaseRequest
	
	// 模块名AccountMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CreateEnterpriseAccount
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 企业用户名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 企业用户证件类型，8代表营业执照，详情请见常见问题
	IdentType *int64 `json:"IdentType,omitempty" name:"IdentType"`

	// 企业用户营业执照号码
	IdentNo *string `json:"IdentNo,omitempty" name:"IdentNo"`

	// 企业联系人手机号
	MobilePhone *string `json:"MobilePhone,omitempty" name:"MobilePhone"`

	// 经办人姓名
	TransactorName *string `json:"TransactorName,omitempty" name:"TransactorName"`

	// 经办人证件类型，0代表身份证
	TransactorIdentType *int64 `json:"TransactorIdentType,omitempty" name:"TransactorIdentType"`

	// 经办人证件号码
	TransactorIdentNo *string `json:"TransactorIdentNo,omitempty" name:"TransactorIdentNo"`

	// 经办人手机号
	TransactorPhone *string `json:"TransactorPhone,omitempty" name:"TransactorPhone"`

	// 企业联系人邮箱
	Email *string `json:"Email,omitempty" name:"Email"`
}

func (r *CreateEnterpriseAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnterpriseAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Name")
	delete(f, "IdentType")
	delete(f, "IdentNo")
	delete(f, "MobilePhone")
	delete(f, "TransactorName")
	delete(f, "TransactorIdentType")
	delete(f, "TransactorIdentNo")
	delete(f, "TransactorPhone")
	delete(f, "Email")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnterpriseAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnterpriseAccountResponseParams struct {
	// 帐号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateEnterpriseAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnterpriseAccountResponseParams `json:"Response"`
}

func (r *CreateEnterpriseAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnterpriseAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonalAccountRequestParams struct {
	// 模块名AccountMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CreatePersonalAccount
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 个人用户姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 个人用户证件类型，0代表身份证，详情请见常见问题
	IdentType *int64 `json:"IdentType,omitempty" name:"IdentType"`

	// 个人用户证件号码
	IdentNo *string `json:"IdentNo,omitempty" name:"IdentNo"`

	// 个人用户手机号
	MobilePhone *string `json:"MobilePhone,omitempty" name:"MobilePhone"`
}

type CreatePersonalAccountRequest struct {
	*tchttp.BaseRequest
	
	// 模块名AccountMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CreatePersonalAccount
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 个人用户姓名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 个人用户证件类型，0代表身份证，详情请见常见问题
	IdentType *int64 `json:"IdentType,omitempty" name:"IdentType"`

	// 个人用户证件号码
	IdentNo *string `json:"IdentNo,omitempty" name:"IdentNo"`

	// 个人用户手机号
	MobilePhone *string `json:"MobilePhone,omitempty" name:"MobilePhone"`
}

func (r *CreatePersonalAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonalAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "Name")
	delete(f, "IdentType")
	delete(f, "IdentNo")
	delete(f, "MobilePhone")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePersonalAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePersonalAccountResponseParams struct {
	// 账号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreatePersonalAccountResponse struct {
	*tchttp.BaseResponse
	Response *CreatePersonalAccountResponseParams `json:"Response"`
}

func (r *CreatePersonalAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePersonalAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealRequestParams struct {
	// 模块名SealMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CreateSeal
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 帐号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 签章链接，图片必须为png格式
	ImgUrl *string `json:"ImgUrl,omitempty" name:"ImgUrl"`

	// 图片数据，base64编码
	ImgData *string `json:"ImgData,omitempty" name:"ImgData"`
}

type CreateSealRequest struct {
	*tchttp.BaseRequest
	
	// 模块名SealMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名CreateSeal
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 帐号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 签章链接，图片必须为png格式
	ImgUrl *string `json:"ImgUrl,omitempty" name:"ImgUrl"`

	// 图片数据，base64编码
	ImgData *string `json:"ImgData,omitempty" name:"ImgData"`
}

func (r *CreateSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "AccountResId")
	delete(f, "ImgUrl")
	delete(f, "ImgData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSealResponseParams struct {
	// 签章ID
	SealResId *string `json:"SealResId,omitempty" name:"SealResId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSealResponse struct {
	*tchttp.BaseResponse
	Response *CreateSealResponseParams `json:"Response"`
}

func (r *CreateSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountRequestParams struct {
	// 模块名AccountMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名DeleteAccount
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 帐号ID列表
	AccountList []*string `json:"AccountList,omitempty" name:"AccountList"`
}

type DeleteAccountRequest struct {
	*tchttp.BaseRequest
	
	// 模块名AccountMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名DeleteAccount
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 帐号ID列表
	AccountList []*string `json:"AccountList,omitempty" name:"AccountList"`
}

func (r *DeleteAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "AccountList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAccountResponseParams struct {
	// 删除成功帐号ID列表
	DelSuccessList []*string `json:"DelSuccessList,omitempty" name:"DelSuccessList"`

	// 删除失败帐号ID列表
	DelFailedList []*string `json:"DelFailedList,omitempty" name:"DelFailedList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAccountResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAccountResponseParams `json:"Response"`
}

func (r *DeleteAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSealRequestParams struct {
	// 模块名SealMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名DeleteSeal
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 帐号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 签章ID
	SealResId *string `json:"SealResId,omitempty" name:"SealResId"`
}

type DeleteSealRequest struct {
	*tchttp.BaseRequest
	
	// 模块名SealMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名DeleteSeal
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 帐号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 签章ID
	SealResId *string `json:"SealResId,omitempty" name:"SealResId"`
}

func (r *DeleteSealRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSealRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "AccountResId")
	delete(f, "SealResId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSealRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSealResponseParams struct {
	// 签章ID
	SealResId *string `json:"SealResId,omitempty" name:"SealResId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteSealResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSealResponseParams `json:"Response"`
}

func (r *DeleteSealResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSealResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusRequestParams struct {
	// 模块名CommonMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名DescribeTaskStatus
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 模块名CommonMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名DescribeTaskStatus
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 任务ID
	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskStatusResponseParams struct {
	// 任务结果
	TaskResult *string `json:"TaskResult,omitempty" name:"TaskResult"`

	// 任务类型，010代表合同上传结果，020代表合同下载结果
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadContractRequestParams struct {
	// 模块名ContractMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名DownloadContract
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`
}

type DownloadContractRequest struct {
	*tchttp.BaseRequest
	
	// 模块名ContractMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名DownloadContract
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`
}

func (r *DownloadContractRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadContractRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ContractResId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadContractRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DownloadContractResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DownloadContractResponse struct {
	*tchttp.BaseResponse
	Response *DownloadContractResponseParams `json:"Response"`
}

func (r *DownloadContractResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadContractResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendVcodeRequestParams struct {
	// 模块名VerifyCode
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名SendVcode
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`

	// 帐号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`
}

type SendVcodeRequest struct {
	*tchttp.BaseRequest
	
	// 模块名VerifyCode
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名SendVcode
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`

	// 帐号ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`
}

func (r *SendVcodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendVcodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ContractResId")
	delete(f, "AccountResId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendVcodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendVcodeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SendVcodeResponse struct {
	*tchttp.BaseResponse
	Response *SendVcodeResponseParams `json:"Response"`
}

func (r *SendVcodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendVcodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignContractByCoordinateRequestParams struct {
	// 模块名ContractMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名SignContractByCoordinate
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`

	// 帐户ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 签署坐标，坐标原点在文件左下角，坐标单位为磅，坐标不得超过合同文件边界
	SignLocations []*SignLocation `json:"SignLocations,omitempty" name:"SignLocations"`

	// 授权时间（由平台自动填充）
	AuthorizationTime *string `json:"AuthorizationTime,omitempty" name:"AuthorizationTime"`

	// 授权IP地址（由平台自动填充）
	Position *string `json:"Position,omitempty" name:"Position"`

	// 签章ID
	SealResId *string `json:"SealResId,omitempty" name:"SealResId"`

	// 选用证书类型：1  表示RSA证书， 2 表示国密证书， 参数不传时默认为1
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 签名图片，base64编码
	ImageData *string `json:"ImageData,omitempty" name:"ImageData"`
}

type SignContractByCoordinateRequest struct {
	*tchttp.BaseRequest
	
	// 模块名ContractMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名SignContractByCoordinate
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`

	// 帐户ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 签署坐标，坐标原点在文件左下角，坐标单位为磅，坐标不得超过合同文件边界
	SignLocations []*SignLocation `json:"SignLocations,omitempty" name:"SignLocations"`

	// 授权时间（由平台自动填充）
	AuthorizationTime *string `json:"AuthorizationTime,omitempty" name:"AuthorizationTime"`

	// 授权IP地址（由平台自动填充）
	Position *string `json:"Position,omitempty" name:"Position"`

	// 签章ID
	SealResId *string `json:"SealResId,omitempty" name:"SealResId"`

	// 选用证书类型：1  表示RSA证书， 2 表示国密证书， 参数不传时默认为1
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 签名图片，base64编码
	ImageData *string `json:"ImageData,omitempty" name:"ImageData"`
}

func (r *SignContractByCoordinateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignContractByCoordinateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ContractResId")
	delete(f, "AccountResId")
	delete(f, "SignLocations")
	delete(f, "AuthorizationTime")
	delete(f, "Position")
	delete(f, "SealResId")
	delete(f, "CertType")
	delete(f, "ImageData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SignContractByCoordinateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignContractByCoordinateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SignContractByCoordinateResponse struct {
	*tchttp.BaseResponse
	Response *SignContractByCoordinateResponseParams `json:"Response"`
}

func (r *SignContractByCoordinateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignContractByCoordinateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignContractByKeywordRequestParams struct {
	// 模块名ContractMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名SignContractByKeyword
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`

	// 账户ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 签署关键字，偏移坐标原点为关键字中心
	SignKeyword *SignKeyword `json:"SignKeyword,omitempty" name:"SignKeyword"`

	// 授权时间（由平台自动填充）
	AuthorizationTime *string `json:"AuthorizationTime,omitempty" name:"AuthorizationTime"`

	// 授权IP地址（由平台自动填充）
	Position *string `json:"Position,omitempty" name:"Position"`

	// 签章ID
	SealResId *string `json:"SealResId,omitempty" name:"SealResId"`

	// 选用证书类型：1  表示RSA证书， 2 表示国密证书， 参数不传时默认为1
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 签名图片，base64编码
	ImageData *string `json:"ImageData,omitempty" name:"ImageData"`
}

type SignContractByKeywordRequest struct {
	*tchttp.BaseRequest
	
	// 模块名ContractMng
	Module *string `json:"Module,omitempty" name:"Module"`

	// 操作名SignContractByKeyword
	Operation *string `json:"Operation,omitempty" name:"Operation"`

	// 合同ID
	ContractResId *string `json:"ContractResId,omitempty" name:"ContractResId"`

	// 账户ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 签署关键字，偏移坐标原点为关键字中心
	SignKeyword *SignKeyword `json:"SignKeyword,omitempty" name:"SignKeyword"`

	// 授权时间（由平台自动填充）
	AuthorizationTime *string `json:"AuthorizationTime,omitempty" name:"AuthorizationTime"`

	// 授权IP地址（由平台自动填充）
	Position *string `json:"Position,omitempty" name:"Position"`

	// 签章ID
	SealResId *string `json:"SealResId,omitempty" name:"SealResId"`

	// 选用证书类型：1  表示RSA证书， 2 表示国密证书， 参数不传时默认为1
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 签名图片，base64编码
	ImageData *string `json:"ImageData,omitempty" name:"ImageData"`
}

func (r *SignContractByKeywordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignContractByKeywordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Module")
	delete(f, "Operation")
	delete(f, "ContractResId")
	delete(f, "AccountResId")
	delete(f, "SignKeyword")
	delete(f, "AuthorizationTime")
	delete(f, "Position")
	delete(f, "SealResId")
	delete(f, "CertType")
	delete(f, "ImageData")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SignContractByKeywordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SignContractByKeywordResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SignContractByKeywordResponse struct {
	*tchttp.BaseResponse
	Response *SignContractByKeywordResponseParams `json:"Response"`
}

func (r *SignContractByKeywordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SignContractByKeywordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SignInfo struct {
	// 账户ID
	AccountResId *string `json:"AccountResId,omitempty" name:"AccountResId"`

	// 授权时间（上传合同可不传该参数）
	AuthorizationTime *string `json:"AuthorizationTime,omitempty" name:"AuthorizationTime"`

	// 授权IP地址（上传合同可不传该参数）
	Location *string `json:"Location,omitempty" name:"Location"`

	// 签章ID
	SealId *string `json:"SealId,omitempty" name:"SealId"`

	// 签名图片，优先级比SealId高
	ImageData *string `json:"ImageData,omitempty" name:"ImageData"`

	// 默认值：1  表示RSA证书， 2 表示国密证书， 参数不传时默认为1
	CertType *int64 `json:"CertType,omitempty" name:"CertType"`

	// 签名域的标签值
	SignLocation *string `json:"SignLocation,omitempty" name:"SignLocation"`
}

type SignKeyword struct {
	// 关键字
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// X轴偏移坐标
	OffsetCoordX *string `json:"OffsetCoordX,omitempty" name:"OffsetCoordX"`

	// Y轴偏移坐标
	OffsetCoordY *string `json:"OffsetCoordY,omitempty" name:"OffsetCoordY"`

	// 签章图片宽度
	ImageWidth *string `json:"ImageWidth,omitempty" name:"ImageWidth"`

	// 签章图片高度
	ImageHeight *string `json:"ImageHeight,omitempty" name:"ImageHeight"`
}

type SignLocation struct {
	// 签名域页数
	SignOnPage *string `json:"SignOnPage,omitempty" name:"SignOnPage"`

	// 签名域左下角X轴坐标轴
	SignLocationLBX *string `json:"SignLocationLBX,omitempty" name:"SignLocationLBX"`

	// 签名域左下角Y轴坐标轴
	SignLocationLBY *string `json:"SignLocationLBY,omitempty" name:"SignLocationLBY"`

	// 签名域右上角X轴坐标轴
	SignLocationRUX *string `json:"SignLocationRUX,omitempty" name:"SignLocationRUX"`

	// 签名域右上角Y轴坐标轴
	SignLocationRUY *string `json:"SignLocationRUY,omitempty" name:"SignLocationRUY"`
}