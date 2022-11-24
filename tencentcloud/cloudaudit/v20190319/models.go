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

package v20190319

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AttributeKeyDetail struct {
	// 输入框类型
	LabelType *string `json:"LabelType,omitempty" name:"LabelType"`

	// 初始化展示
	Starter *string `json:"Starter,omitempty" name:"Starter"`

	// 展示排序
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// AttributeKey值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 中文标签
	Label *string `json:"Label,omitempty" name:"Label"`
}

type AuditSummary struct {
	// 跟踪集状态，1：开启，0：关闭
	AuditStatus *int64 `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// COS存储桶名称
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// 日志前缀
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`
}

type CmqRegionInfo struct {
	// 地域描述
	CmqRegionName *string `json:"CmqRegionName,omitempty" name:"CmqRegionName"`

	// cmq地域
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`
}

type CosRegionInfo struct {
	// cos地域
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 地域描述
	CosRegionName *string `json:"CosRegionName,omitempty" name:"CosRegionName"`
}

// Predefined struct for user
type CreateAuditRequestParams struct {
	// 是否开启cmq消息通知。1：是，0：否。目前仅支持cmq的队列服务。如果开启cmq消息通知服务，云审计会将您的日志内容实时投递到您指定地域的指定队列中。
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// 管理事件的读写属性。1：只读，2：只写，3：全部。
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// 跟踪集名称。3-128字符，只能包含 ASCII 编码字母 a-z，A-Z，数字 0-9，下划线 _。
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// cos地域。目前支持的地域可以使用ListCosEnableRegion来获取。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 是否创建新的cos存储桶。1：是，0：否。
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// cos的存储桶名称。仅支持小写英文字母和数字即[a-z，0-9]、中划线“-”及其组合。用户自定义的字符串支持1 - 40个字符。存储桶命名不能以“-”开头或结尾。如果不是新创建的存储桶，云审计不会去校验该存储桶是否真的存在，请谨慎填写，避免日志投递不成功，导致您的数据丢失。
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// CMK的全局唯一标识符，如果不是新创建的kms，该值是必填值。可以通过ListKeyAliasByRegion来获取。云审计不会校验KeyId的合法性，请您谨慎填写，避免给您的数据造成损失。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 队列名称。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。如果IsEnableCmqNotify值是1的话，此值属于必填字段。如果不是新创建的队列，云审计不会去校验该队列是否真的存在，请谨慎填写，避免日志通知不成功，导致您的数据丢失。
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// kms地域。目前支持的地域可以使用ListKmsEnableRegion来获取。必须要和cos的地域保持一致。
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// 是否开启kms加密。1：是，0：否。如果开启KMS加密，数据在投递到cos时，会将数据加密。
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// 队列所在的地域。可以通过ListCmqEnableRegion获取支持的cmq地域。如果IsEnableCmqNotify值是1的话，此值属于必填字段。
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// 日志文件前缀。3-40个字符，只能包含 ASCII 编码字母 a-z，A-Z，数字 0-9。可以不填，默认以账号ID作为日志前缀。
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// 是否创建新的队列。1：是，0：否。如果IsEnableCmqNotify值是1的话，此值属于必填字段。
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`
}

type CreateAuditRequest struct {
	*tchttp.BaseRequest
	
	// 是否开启cmq消息通知。1：是，0：否。目前仅支持cmq的队列服务。如果开启cmq消息通知服务，云审计会将您的日志内容实时投递到您指定地域的指定队列中。
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// 管理事件的读写属性。1：只读，2：只写，3：全部。
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// 跟踪集名称。3-128字符，只能包含 ASCII 编码字母 a-z，A-Z，数字 0-9，下划线 _。
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// cos地域。目前支持的地域可以使用ListCosEnableRegion来获取。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 是否创建新的cos存储桶。1：是，0：否。
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// cos的存储桶名称。仅支持小写英文字母和数字即[a-z，0-9]、中划线“-”及其组合。用户自定义的字符串支持1 - 40个字符。存储桶命名不能以“-”开头或结尾。如果不是新创建的存储桶，云审计不会去校验该存储桶是否真的存在，请谨慎填写，避免日志投递不成功，导致您的数据丢失。
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// CMK的全局唯一标识符，如果不是新创建的kms，该值是必填值。可以通过ListKeyAliasByRegion来获取。云审计不会校验KeyId的合法性，请您谨慎填写，避免给您的数据造成损失。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 队列名称。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。如果IsEnableCmqNotify值是1的话，此值属于必填字段。如果不是新创建的队列，云审计不会去校验该队列是否真的存在，请谨慎填写，避免日志通知不成功，导致您的数据丢失。
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// kms地域。目前支持的地域可以使用ListKmsEnableRegion来获取。必须要和cos的地域保持一致。
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// 是否开启kms加密。1：是，0：否。如果开启KMS加密，数据在投递到cos时，会将数据加密。
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// 队列所在的地域。可以通过ListCmqEnableRegion获取支持的cmq地域。如果IsEnableCmqNotify值是1的话，此值属于必填字段。
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// 日志文件前缀。3-40个字符，只能包含 ASCII 编码字母 a-z，A-Z，数字 0-9。可以不填，默认以账号ID作为日志前缀。
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// 是否创建新的队列。1：是，0：否。如果IsEnableCmqNotify值是1的话，此值属于必填字段。
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`
}

func (r *CreateAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsEnableCmqNotify")
	delete(f, "ReadWriteAttribute")
	delete(f, "AuditName")
	delete(f, "CosRegion")
	delete(f, "IsCreateNewBucket")
	delete(f, "CosBucketName")
	delete(f, "KeyId")
	delete(f, "CmqQueueName")
	delete(f, "KmsRegion")
	delete(f, "IsEnableKmsEncry")
	delete(f, "CmqRegion")
	delete(f, "LogFilePrefix")
	delete(f, "IsCreateNewQueue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditResponseParams struct {
	// 是否创建成功。
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAuditResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditResponseParams `json:"Response"`
}

func (r *CreateAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditTrackRequestParams struct {
	// 跟踪集名称，仅支持大小写字母、数字、-以及_的组合，3-48个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 跟踪事件类型（读：Read；写：Write；全部：*）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 跟踪事件所属产品（支持全部产品或单个产品，如：cos，全部：*）
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 跟踪集状态（未开启：0；开启：1）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 跟踪事件接口名列表（ResourceType为 * 时，EventNames必须为全部：["*"]；指定ResourceType时，支持全部接口：["*"]；支持部分接口：["cos", "cls"]，接口列表上限10个）
	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`

	// 数据投递存储（目前支持 cos、cls）
	Storage *Storage `json:"Storage,omitempty" name:"Storage"`

	// 是否开启将集团成员操作日志投递到集团管理账号或者可信服务管理账号(0：未开启，1：开启，只能集团管理账号或者可信服务管理账号开启此项功能)
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitempty" name:"TrackForAllMembers"`
}

type CreateAuditTrackRequest struct {
	*tchttp.BaseRequest
	
	// 跟踪集名称，仅支持大小写字母、数字、-以及_的组合，3-48个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 跟踪事件类型（读：Read；写：Write；全部：*）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 跟踪事件所属产品（支持全部产品或单个产品，如：cos，全部：*）
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 跟踪集状态（未开启：0；开启：1）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 跟踪事件接口名列表（ResourceType为 * 时，EventNames必须为全部：["*"]；指定ResourceType时，支持全部接口：["*"]；支持部分接口：["cos", "cls"]，接口列表上限10个）
	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`

	// 数据投递存储（目前支持 cos、cls）
	Storage *Storage `json:"Storage,omitempty" name:"Storage"`

	// 是否开启将集团成员操作日志投递到集团管理账号或者可信服务管理账号(0：未开启，1：开启，只能集团管理账号或者可信服务管理账号开启此项功能)
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitempty" name:"TrackForAllMembers"`
}

func (r *CreateAuditTrackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditTrackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ActionType")
	delete(f, "ResourceType")
	delete(f, "Status")
	delete(f, "EventNames")
	delete(f, "Storage")
	delete(f, "TrackForAllMembers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditTrackResponseParams struct {
	// 跟踪集 ID
	TrackId *uint64 `json:"TrackId,omitempty" name:"TrackId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAuditTrackResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditTrackResponseParams `json:"Response"`
}

func (r *CreateAuditTrackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditRequestParams struct {
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

type DeleteAuditRequest struct {
	*tchttp.BaseRequest
	
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

func (r *DeleteAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditResponseParams struct {
	// 是否删除成功
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAuditResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditResponseParams `json:"Response"`
}

func (r *DeleteAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditTrackRequestParams struct {
	// 跟踪集 ID
	TrackId *uint64 `json:"TrackId,omitempty" name:"TrackId"`
}

type DeleteAuditTrackRequest struct {
	*tchttp.BaseRequest
	
	// 跟踪集 ID
	TrackId *uint64 `json:"TrackId,omitempty" name:"TrackId"`
}

func (r *DeleteAuditTrackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditTrackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditTrackResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteAuditTrackResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditTrackResponseParams `json:"Response"`
}

func (r *DeleteAuditTrackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditRequestParams struct {
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

type DescribeAuditRequest struct {
	*tchttp.BaseRequest
	
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

func (r *DescribeAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditResponseParams struct {
	// 是否开启cmq消息通知。1：是，0：否。
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// 管理事件读写属性，1：只读，2：只写，3：全部
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// CMK的全局唯一标识符。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// 跟踪集状态，1：开启，0：停止。
	AuditStatus *int64 `json:"AuditStatus,omitempty" name:"AuditStatus"`

	// 跟踪集名称。
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// cos存储桶所在地域。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 队列名称。
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// CMK别名。
	KmsAlias *string `json:"KmsAlias,omitempty" name:"KmsAlias"`

	// kms地域。
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// 是否开启kms加密。1：是，0：否。如果开启KMS加密，数据在投递到cos时，会将数据加密。
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// cos存储桶名称。
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 队列所在地域。
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// 日志前缀。
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditResponseParams `json:"Response"`
}

func (r *DescribeAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditTrackRequestParams struct {
	// 跟踪集 ID
	TrackId *uint64 `json:"TrackId,omitempty" name:"TrackId"`
}

type DescribeAuditTrackRequest struct {
	*tchttp.BaseRequest
	
	// 跟踪集 ID
	TrackId *uint64 `json:"TrackId,omitempty" name:"TrackId"`
}

func (r *DescribeAuditTrackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditTrackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrackId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditTrackResponseParams struct {
	// 跟踪集名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 跟踪事件类型（读：Read；写：Write；全部：*）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 跟踪事件所属产品（如：cos，全部：*）
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 跟踪集状态（未开启：0；开启：1）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 跟踪事件接口名列表（全部：[*]）
	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`

	// 数据投递存储（目前支持 cos、cls）
	Storage *Storage `json:"Storage,omitempty" name:"Storage"`

	// 跟踪集创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 是否开启将集团成员操作日志投递到集团管理账号或者可信服务管理账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitempty" name:"TrackForAllMembers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditTrackResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditTrackResponseParams `json:"Response"`
}

func (r *DescribeAuditTrackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditTracksRequestParams struct {
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页数目
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeAuditTracksRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页数目
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeAuditTracksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditTracksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditTracksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditTracksResponseParams struct {
	// 跟踪集列表
	Tracks []*Tracks `json:"Tracks,omitempty" name:"Tracks"`

	// 总数目
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAuditTracksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditTracksResponseParams `json:"Response"`
}

func (r *DescribeAuditTracksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditTracksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsRequestParams struct {
	// 起始时间戳（单位秒，不超过当前时间 90 天）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳（单位秒，查询时间跨度小于 30 天）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 查看更多日志的凭证
	NextToken *uint64 `json:"NextToken,omitempty" name:"NextToken"`

	// 返回日志的最大条数（最大 50 条）
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// 检索条件（目前支持 RequestId：请求 ID、EventName：事件名称、ActionType：操作类型（Write：写；Read：读）、PrincipalId：子账号、ResourceType：资源类型、ResourceName：资源名称、AccessKeyId：密钥 ID、SensitiveAction：是否敏感操作、ApiErrorCode：API 错误码、CamErrorCode：CAM 错误码、Tags：标签（AttributeValue格式：[{"key":"*","value":"*"}]））
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes"`

	// 是否返回 IP 归属地（1 返回，0 不返回）
	IsReturnLocation *uint64 `json:"IsReturnLocation,omitempty" name:"IsReturnLocation"`
}

type DescribeEventsRequest struct {
	*tchttp.BaseRequest
	
	// 起始时间戳（单位秒，不超过当前时间 90 天）
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳（单位秒，查询时间跨度小于 30 天）
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 查看更多日志的凭证
	NextToken *uint64 `json:"NextToken,omitempty" name:"NextToken"`

	// 返回日志的最大条数（最大 50 条）
	MaxResults *uint64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// 检索条件（目前支持 RequestId：请求 ID、EventName：事件名称、ActionType：操作类型（Write：写；Read：读）、PrincipalId：子账号、ResourceType：资源类型、ResourceName：资源名称、AccessKeyId：密钥 ID、SensitiveAction：是否敏感操作、ApiErrorCode：API 错误码、CamErrorCode：CAM 错误码、Tags：标签（AttributeValue格式：[{"key":"*","value":"*"}]））
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes"`

	// 是否返回 IP 归属地（1 返回，0 不返回）
	IsReturnLocation *uint64 `json:"IsReturnLocation,omitempty" name:"IsReturnLocation"`
}

func (r *DescribeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "LookupAttributes")
	delete(f, "IsReturnLocation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventsResponseParams struct {
	// 日志集合是否结束。true表示结束，无需进行翻页。
	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

	// 查看更多日志的凭证
	NextToken *uint64 `json:"NextToken,omitempty" name:"NextToken"`

	// 日志集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*Event `json:"Events,omitempty" name:"Events"`

	// 此字段已经废弃。翻页请使用ListOver配合NextToken，在ListOver为false进行下一页数据读取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventsResponseParams `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Event struct {
	// 日志ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 用户名
	Username *string `json:"Username,omitempty" name:"Username"`

	// 事件时间
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// 日志详情
	CloudAuditEvent *string `json:"CloudAuditEvent,omitempty" name:"CloudAuditEvent"`

	// 资源类型中文描述（此字段请按需使用，如果您是其他语言使用者，可以忽略该字段描述）
	ResourceTypeCn *string `json:"ResourceTypeCn,omitempty" name:"ResourceTypeCn"`

	// 鉴权错误码
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 事件名称
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 证书ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// 请求来源
	EventSource *string `json:"EventSource,omitempty" name:"EventSource"`

	// 请求ID
	RequestID *string `json:"RequestID,omitempty" name:"RequestID"`

	// 资源地域
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 主账号ID
	AccountID *int64 `json:"AccountID,omitempty" name:"AccountID"`

	// 源IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceIPAddress *string `json:"SourceIPAddress,omitempty" name:"SourceIPAddress"`

	// 事件名称中文描述（此字段请按需使用，如果您是其他语言使用者，可以忽略该字段描述）
	EventNameCn *string `json:"EventNameCn,omitempty" name:"EventNameCn"`

	// 资源对
	Resources *Resource `json:"Resources,omitempty" name:"Resources"`

	// 事件地域
	EventRegion *string `json:"EventRegion,omitempty" name:"EventRegion"`

	// IP 归属地
	Location *string `json:"Location,omitempty" name:"Location"`
}

// Predefined struct for user
type GetAttributeKeyRequestParams struct {
	// 网站类型，取值范围是zh和en。如果不传值默认zh
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

type GetAttributeKeyRequest struct {
	*tchttp.BaseRequest
	
	// 网站类型，取值范围是zh和en。如果不传值默认zh
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

func (r *GetAttributeKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttributeKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebsiteType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAttributeKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAttributeKeyResponseParams struct {
	// AttributeKey的有效取值范围
	AttributeKeyDetails []*AttributeKeyDetail `json:"AttributeKeyDetails,omitempty" name:"AttributeKeyDetails"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetAttributeKeyResponse struct {
	*tchttp.BaseResponse
	Response *GetAttributeKeyResponseParams `json:"Response"`
}

func (r *GetAttributeKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttributeKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquireAuditCreditRequestParams struct {

}

type InquireAuditCreditRequest struct {
	*tchttp.BaseRequest
	
}

func (r *InquireAuditCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireAuditCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InquireAuditCreditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InquireAuditCreditResponseParams struct {
	// 可创建跟踪集的数量
	AuditAmount *int64 `json:"AuditAmount,omitempty" name:"AuditAmount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InquireAuditCreditResponse struct {
	*tchttp.BaseResponse
	Response *InquireAuditCreditResponseParams `json:"Response"`
}

func (r *InquireAuditCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireAuditCreditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyMetadata struct {
	// 作为密钥更容易辨识，更容易被人看懂的别名
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// CMK的全局唯一标识
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

// Predefined struct for user
type ListAuditsRequestParams struct {

}

type ListAuditsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListAuditsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuditsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAuditsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAuditsResponseParams struct {
	// 查询跟踪集概要集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditSummarys []*AuditSummary `json:"AuditSummarys,omitempty" name:"AuditSummarys"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListAuditsResponse struct {
	*tchttp.BaseResponse
	Response *ListAuditsResponseParams `json:"Response"`
}

func (r *ListAuditsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuditsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCmqEnableRegionRequestParams struct {
	// 站点类型。zh表示中国区，en表示国际区。默认中国区。
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

type ListCmqEnableRegionRequest struct {
	*tchttp.BaseRequest
	
	// 站点类型。zh表示中国区，en表示国际区。默认中国区。
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

func (r *ListCmqEnableRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCmqEnableRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebsiteType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCmqEnableRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCmqEnableRegionResponseParams struct {
	// 云审计支持的cmq的可用区
	EnableRegions []*CmqRegionInfo `json:"EnableRegions,omitempty" name:"EnableRegions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListCmqEnableRegionResponse struct {
	*tchttp.BaseResponse
	Response *ListCmqEnableRegionResponseParams `json:"Response"`
}

func (r *ListCmqEnableRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCmqEnableRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCosEnableRegionRequestParams struct {
	// 站点类型。zh表示中国区，en表示国际区。默认中国区。
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

type ListCosEnableRegionRequest struct {
	*tchttp.BaseRequest
	
	// 站点类型。zh表示中国区，en表示国际区。默认中国区。
	WebsiteType *string `json:"WebsiteType,omitempty" name:"WebsiteType"`
}

func (r *ListCosEnableRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCosEnableRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebsiteType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListCosEnableRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListCosEnableRegionResponseParams struct {
	// 云审计支持的cos可用区
	EnableRegions []*CosRegionInfo `json:"EnableRegions,omitempty" name:"EnableRegions"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListCosEnableRegionResponse struct {
	*tchttp.BaseResponse
	Response *ListCosEnableRegionResponseParams `json:"Response"`
}

func (r *ListCosEnableRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCosEnableRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeyAliasByRegionRequestParams struct {
	// Kms地域
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次获最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type ListKeyAliasByRegionRequest struct {
	*tchttp.BaseRequest
	
	// Kms地域
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// 含义跟 SQL 查询的 Limit 一致，表示本次获最多获取 Limit 个元素。缺省值为10，最大值为200
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListKeyAliasByRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeyAliasByRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KmsRegion")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListKeyAliasByRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListKeyAliasByRegionResponseParams struct {
	// CMK的总数量
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 密钥别名
	KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitempty" name:"KeyMetadatas"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ListKeyAliasByRegionResponse struct {
	*tchttp.BaseResponse
	Response *ListKeyAliasByRegionResponseParams `json:"Response"`
}

func (r *ListKeyAliasByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeyAliasByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LookUpEventsRequestParams struct {
	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 检索条件
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes"`

	// 查看更多日志的凭证
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 返回日志的最大条数
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// 云审计模式，有效值：standard | quick，其中standard是标准模式，quick是极速模式。默认为标准模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type LookUpEventsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 检索条件
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes"`

	// 查看更多日志的凭证
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 返回日志的最大条数
	MaxResults *int64 `json:"MaxResults,omitempty" name:"MaxResults"`

	// 云审计模式，有效值：standard | quick，其中standard是标准模式，quick是极速模式。默认为标准模式
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *LookUpEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LookUpEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "LookupAttributes")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LookUpEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LookUpEventsResponseParams struct {
	// 查看更多日志的凭证
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 日志集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*Event `json:"Events,omitempty" name:"Events"`

	// 日志集合是否结束
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LookUpEventsResponse struct {
	*tchttp.BaseResponse
	Response *LookUpEventsResponseParams `json:"Response"`
}

func (r *LookUpEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LookUpEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LookupAttribute struct {
	// AttributeKey的有效取值范围是:RequestId、EventName、ReadOnly、Username、ResourceType、ResourceName和AccessKeyId，EventId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeKey *string `json:"AttributeKey,omitempty" name:"AttributeKey"`

	// AttributeValue的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttributeValue *string `json:"AttributeValue,omitempty" name:"AttributeValue"`
}

// Predefined struct for user
type ModifyAuditTrackRequestParams struct {
	// 跟踪集 ID
	TrackId *uint64 `json:"TrackId,omitempty" name:"TrackId"`

	// 跟踪集名称，仅支持大小写字母、数字、-以及_的组合，3-48个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 跟踪事件类型（读：Read；写：Write；全部：*）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 跟踪事件所属产品（支持全部产品或单个产品，如：cos，全部：*）
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 跟踪集状态（未开启：0；开启：1）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 跟踪事件接口名列表（ResourceType为 * 时，EventNames必须为全部：["*"]；指定ResourceType时，支持全部接口：["*"]；支持部分接口：["cos", "cls"]，接口列表上限10个）
	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`

	// 数据投递存储（目前支持 cos、cls）
	Storage *Storage `json:"Storage,omitempty" name:"Storage"`

	// 是否开启将集团成员操作日志投递到集团管理账号或者可信服务管理账号(0：未开启，1：开启，只能集团管理账号或者可信服务管理账号开启此项功能)
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitempty" name:"TrackForAllMembers"`
}

type ModifyAuditTrackRequest struct {
	*tchttp.BaseRequest
	
	// 跟踪集 ID
	TrackId *uint64 `json:"TrackId,omitempty" name:"TrackId"`

	// 跟踪集名称，仅支持大小写字母、数字、-以及_的组合，3-48个字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 跟踪事件类型（读：Read；写：Write；全部：*）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 跟踪事件所属产品（支持全部产品或单个产品，如：cos，全部：*）
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 跟踪集状态（未开启：0；开启：1）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 跟踪事件接口名列表（ResourceType为 * 时，EventNames必须为全部：["*"]；指定ResourceType时，支持全部接口：["*"]；支持部分接口：["cos", "cls"]，接口列表上限10个）
	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`

	// 数据投递存储（目前支持 cos、cls）
	Storage *Storage `json:"Storage,omitempty" name:"Storage"`

	// 是否开启将集团成员操作日志投递到集团管理账号或者可信服务管理账号(0：未开启，1：开启，只能集团管理账号或者可信服务管理账号开启此项功能)
	TrackForAllMembers *uint64 `json:"TrackForAllMembers,omitempty" name:"TrackForAllMembers"`
}

func (r *ModifyAuditTrackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditTrackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TrackId")
	delete(f, "Name")
	delete(f, "ActionType")
	delete(f, "ResourceType")
	delete(f, "Status")
	delete(f, "EventNames")
	delete(f, "Storage")
	delete(f, "TrackForAllMembers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditTrackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditTrackResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAuditTrackResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditTrackResponseParams `json:"Response"`
}

func (r *ModifyAuditTrackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditTrackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

// Predefined struct for user
type StartLoggingRequestParams struct {
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

type StartLoggingRequest struct {
	*tchttp.BaseRequest
	
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

func (r *StartLoggingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLoggingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartLoggingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartLoggingResponseParams struct {
	// 是否开启成功
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartLoggingResponse struct {
	*tchttp.BaseResponse
	Response *StartLoggingResponseParams `json:"Response"`
}

func (r *StartLoggingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLoggingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLoggingRequestParams struct {
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

type StopLoggingRequest struct {
	*tchttp.BaseRequest
	
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`
}

func (r *StopLoggingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLoggingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopLoggingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopLoggingResponseParams struct {
	// 是否关闭成功
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopLoggingResponse struct {
	*tchttp.BaseResponse
	Response *StopLoggingResponseParams `json:"Response"`
}

func (r *StopLoggingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLoggingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Storage struct {
	// 存储类型（目前支持 cos、cls）
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`

	// 存储所在地域
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 存储名称(cos：存储名称为用户自定义的存储桶名称，不包含"-APPID"，仅支持小写字母、数字以及中划线"-"的组合，不能超过50字符，且不支持中划线"-"开头或结尾； cls：存储名称为日志主题id，字符长度为1-50个字符)
	StorageName *string `json:"StorageName,omitempty" name:"StorageName"`

	// 存储目录前缀，cos日志文件前缀仅支持字母和数字的组合，3-40个字符
	StoragePrefix *string `json:"StoragePrefix,omitempty" name:"StoragePrefix"`
}

type Tracks struct {
	// 跟踪集名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 跟踪事件类型（读：Read；写：Write；全部：*）
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 跟踪事件所属产品（如：cos，全部：*）
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 跟踪集状态（未开启：0；开启：1）
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 跟踪事件接口名列表（全部：[*]）
	EventNames []*string `json:"EventNames,omitempty" name:"EventNames"`

	// 数据投递存储（目前支持 cos、cls）
	Storage *Storage `json:"Storage,omitempty" name:"Storage"`

	// 跟踪集创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 跟踪集 ID
	TrackId *uint64 `json:"TrackId,omitempty" name:"TrackId"`
}

// Predefined struct for user
type UpdateAuditRequestParams struct {
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// 是否开启cmq消息通知。1：是，0：否。目前仅支持cmq的队列服务。如果开启cmq消息通知服务，云审计会将您的日志内容实时投递到您指定地域的指定队列中。
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// 管理事件的读写属性。1：只读，2：只写，3：全部。
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// CMK的全局唯一标识符，如果不是新创建的kms，该值是必填值。可以通过ListKeyAliasByRegion来获取。云审计不会校验KeyId的合法性，请您谨慎填写，避免给您的数据造成损失。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// cos地域。目前支持的地域可以使用ListCosEnableRegion来获取。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 队列名称。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。如果IsEnableCmqNotify值是1的话，此值属于必填字段。如果不是新创建的队列，云审计不会去校验该队列是否真的存在，请谨慎填写，避免日志通知不成功，导致您的数据丢失。
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// 是否创建新的cos存储桶。1：是，0：否。
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// kms地域。目前支持的地域可以使用ListKmsEnableRegion来获取。必须要和cos的地域保持一致。
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// 是否开启kms加密。1：是，0：否。如果开启KMS加密，数据在投递到cos时，会将数据加密。
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// cos的存储桶名称。仅支持小写英文字母和数字即[a-z，0-9]、中划线“-”及其组合。用户自定义的字符串支持1 - 40个字符。存储桶命名不能以“-”开头或结尾。如果不是新创建的存储桶，云审计不会去校验该存储桶是否真的存在，请谨慎填写，避免日志投递不成功，导致您的数据丢失。
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 队列所在的地域。可以通过ListCmqEnableRegion获取支持的cmq地域。如果IsEnableCmqNotify值是1的话，此值属于必填字段。
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// 日志文件前缀。3-40个字符，只能包含 ASCII 编码字母 a-z，A-Z，数字 0-9。
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// 是否创建新的队列。1：是，0：否。如果IsEnableCmqNotify值是1的话，此值属于必填字段。
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`
}

type UpdateAuditRequest struct {
	*tchttp.BaseRequest
	
	// 跟踪集名称
	AuditName *string `json:"AuditName,omitempty" name:"AuditName"`

	// 是否开启cmq消息通知。1：是，0：否。目前仅支持cmq的队列服务。如果开启cmq消息通知服务，云审计会将您的日志内容实时投递到您指定地域的指定队列中。
	IsEnableCmqNotify *int64 `json:"IsEnableCmqNotify,omitempty" name:"IsEnableCmqNotify"`

	// 管理事件的读写属性。1：只读，2：只写，3：全部。
	ReadWriteAttribute *int64 `json:"ReadWriteAttribute,omitempty" name:"ReadWriteAttribute"`

	// CMK的全局唯一标识符，如果不是新创建的kms，该值是必填值。可以通过ListKeyAliasByRegion来获取。云审计不会校验KeyId的合法性，请您谨慎填写，避免给您的数据造成损失。
	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`

	// cos地域。目前支持的地域可以使用ListCosEnableRegion来获取。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 队列名称。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。如果IsEnableCmqNotify值是1的话，此值属于必填字段。如果不是新创建的队列，云审计不会去校验该队列是否真的存在，请谨慎填写，避免日志通知不成功，导致您的数据丢失。
	CmqQueueName *string `json:"CmqQueueName,omitempty" name:"CmqQueueName"`

	// 是否创建新的cos存储桶。1：是，0：否。
	IsCreateNewBucket *int64 `json:"IsCreateNewBucket,omitempty" name:"IsCreateNewBucket"`

	// kms地域。目前支持的地域可以使用ListKmsEnableRegion来获取。必须要和cos的地域保持一致。
	KmsRegion *string `json:"KmsRegion,omitempty" name:"KmsRegion"`

	// 是否开启kms加密。1：是，0：否。如果开启KMS加密，数据在投递到cos时，会将数据加密。
	IsEnableKmsEncry *int64 `json:"IsEnableKmsEncry,omitempty" name:"IsEnableKmsEncry"`

	// cos的存储桶名称。仅支持小写英文字母和数字即[a-z，0-9]、中划线“-”及其组合。用户自定义的字符串支持1 - 40个字符。存储桶命名不能以“-”开头或结尾。如果不是新创建的存储桶，云审计不会去校验该存储桶是否真的存在，请谨慎填写，避免日志投递不成功，导致您的数据丢失。
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// 队列所在的地域。可以通过ListCmqEnableRegion获取支持的cmq地域。如果IsEnableCmqNotify值是1的话，此值属于必填字段。
	CmqRegion *string `json:"CmqRegion,omitempty" name:"CmqRegion"`

	// 日志文件前缀。3-40个字符，只能包含 ASCII 编码字母 a-z，A-Z，数字 0-9。
	LogFilePrefix *string `json:"LogFilePrefix,omitempty" name:"LogFilePrefix"`

	// 是否创建新的队列。1：是，0：否。如果IsEnableCmqNotify值是1的话，此值属于必填字段。
	IsCreateNewQueue *int64 `json:"IsCreateNewQueue,omitempty" name:"IsCreateNewQueue"`
}

func (r *UpdateAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	delete(f, "IsEnableCmqNotify")
	delete(f, "ReadWriteAttribute")
	delete(f, "KeyId")
	delete(f, "CosRegion")
	delete(f, "CmqQueueName")
	delete(f, "IsCreateNewBucket")
	delete(f, "KmsRegion")
	delete(f, "IsEnableKmsEncry")
	delete(f, "CosBucketName")
	delete(f, "CmqRegion")
	delete(f, "LogFilePrefix")
	delete(f, "IsCreateNewQueue")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAuditResponseParams struct {
	// 是否更新成功
	IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateAuditResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAuditResponseParams `json:"Response"`
}

func (r *UpdateAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}