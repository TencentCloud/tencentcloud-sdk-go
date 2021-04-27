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
    "errors"

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

// It is highly **NOT** recommended to use this function
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
		return errors.New("CreateAuditRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否创建成功。
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return errors.New("DeleteAuditRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否删除成功
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return errors.New("DescribeAuditRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

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
	} `json:"Response"`
}

func (r *DescribeAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

	// 检索条件（目前支持 RequestId：请求 ID、EventName：事件名称、ActionType：操作类型（Write：写；Read：读）、PrincipalId：子账号、ResourceType：资源类型、ResourceName：资源名称、AccessKeyId：密钥 ID、SensitiveAction：是否敏感操作、ApiErrorCode：API 错误码、CamErrorCode：CAM 错误码）
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes" list`
}

func (r *DescribeEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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
	if len(f) > 0 {
		return errors.New("DescribeEventsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志集合是否结束
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// 查看更多日志的凭证
		NextToken *uint64 `json:"NextToken,omitempty" name:"NextToken"`

		// 日志集合
	// 注意：此字段可能返回 null，表示取不到有效值。
		Events []*Event `json:"Events,omitempty" name:"Events" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttributeKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebsiteType")
	if len(f) > 0 {
		return errors.New("GetAttributeKeyRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetAttributeKeyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// AttributeKey的有效取值范围
		AttributeKeyDetails []*AttributeKeyDetail `json:"AttributeKeyDetails,omitempty" name:"AttributeKeyDetails" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAttributeKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAttributeKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquireAuditCreditRequest struct {
	*tchttp.BaseRequest
}

func (r *InquireAuditCreditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InquireAuditCreditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("InquireAuditCreditRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type InquireAuditCreditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可创建跟踪集的数量
		AuditAmount *int64 `json:"AuditAmount,omitempty" name:"AuditAmount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquireAuditCreditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

type ListAuditsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListAuditsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuditsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	if len(f) > 0 {
		return errors.New("ListAuditsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListAuditsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询跟踪集概要集合
	// 注意：此字段可能返回 null，表示取不到有效值。
		AuditSummarys []*AuditSummary `json:"AuditSummarys,omitempty" name:"AuditSummarys" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAuditsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAuditsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCmqEnableRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebsiteType")
	if len(f) > 0 {
		return errors.New("ListCmqEnableRegionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListCmqEnableRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云审计支持的cmq的可用区
		EnableRegions []*CmqRegionInfo `json:"EnableRegions,omitempty" name:"EnableRegions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCmqEnableRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCmqEnableRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCosEnableRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WebsiteType")
	if len(f) > 0 {
		return errors.New("ListCosEnableRegionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListCosEnableRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云审计支持的cos可用区
		EnableRegions []*CosRegionInfo `json:"EnableRegions,omitempty" name:"EnableRegions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListCosEnableRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListCosEnableRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("ListKeyAliasByRegionRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ListKeyAliasByRegionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// CMK的总数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 密钥别名
		KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitempty" name:"KeyMetadatas" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListKeyAliasByRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListKeyAliasByRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LookUpEventsRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 检索条件
	LookupAttributes []*LookupAttribute `json:"LookupAttributes,omitempty" name:"LookupAttributes" list`

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

// It is highly **NOT** recommended to use this function
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
		return errors.New("LookUpEventsRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type LookUpEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查看更多日志的凭证
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// 日志集合
	// 注意：此字段可能返回 null，表示取不到有效值。
		Events []*Event `json:"Events,omitempty" name:"Events" list`

		// 日志集合是否结束
	// 注意：此字段可能返回 null，表示取不到有效值。
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LookUpEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
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

type Resource struct {

	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`

	// 资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLoggingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return errors.New("StartLoggingRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StartLoggingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否开启成功
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartLoggingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartLoggingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLoggingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AuditName")
	if len(f) > 0 {
		return errors.New("StopLoggingRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type StopLoggingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否关闭成功
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopLoggingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopLoggingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

// It is highly **NOT** recommended to use this function
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
		return errors.New("UpdateAuditRequest has unknown keys!")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAuditResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 是否更新成功
		IsSuccess *int64 `json:"IsSuccess,omitempty" name:"IsSuccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
