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

package v20200506

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeTransactionsRequestParams struct {
	// 事务分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 事务开始时间查询起始时间戳，UTC，精确到毫秒
	TransactionBeginFrom *int64 `json:"TransactionBeginFrom,omitempty" name:"TransactionBeginFrom"`

	// 事务开始时间查询截止时间戳，UTC，精确到毫秒
	TransactionBeginTo *int64 `json:"TransactionBeginTo,omitempty" name:"TransactionBeginTo"`

	// 仅查询异常状态的事务，true：仅查询异常，false或不传入：查询所有
	SearchError *bool `json:"SearchError,omitempty" name:"SearchError"`

	// 主事务ID，不传入时查询全量，高优先级
	TransactionId *int64 `json:"TransactionId,omitempty" name:"TransactionId"`

	// 主事务ID列表，不传入时查询全量，低优先级
	TransactionIdList []*int64 `json:"TransactionIdList,omitempty" name:"TransactionIdList"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeTransactionsRequest struct {
	*tchttp.BaseRequest
	
	// 事务分组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 事务开始时间查询起始时间戳，UTC，精确到毫秒
	TransactionBeginFrom *int64 `json:"TransactionBeginFrom,omitempty" name:"TransactionBeginFrom"`

	// 事务开始时间查询截止时间戳，UTC，精确到毫秒
	TransactionBeginTo *int64 `json:"TransactionBeginTo,omitempty" name:"TransactionBeginTo"`

	// 仅查询异常状态的事务，true：仅查询异常，false或不传入：查询所有
	SearchError *bool `json:"SearchError,omitempty" name:"SearchError"`

	// 主事务ID，不传入时查询全量，高优先级
	TransactionId *int64 `json:"TransactionId,omitempty" name:"TransactionId"`

	// 主事务ID列表，不传入时查询全量，低优先级
	TransactionIdList []*int64 `json:"TransactionIdList,omitempty" name:"TransactionIdList"`

	// 每页数量
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTransactionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTransactionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "TransactionBeginFrom")
	delete(f, "TransactionBeginTo")
	delete(f, "SearchError")
	delete(f, "TransactionId")
	delete(f, "TransactionIdList")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTransactionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTransactionsResponseParams struct {
	// 主事务分页列表
	Result *PagedTransaction `json:"Result,omitempty" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTransactionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTransactionsResponseParams `json:"Response"`
}

func (r *DescribeTransactionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTransactionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PagedTransaction struct {
	// 总条数，特定在该接口中总是会返回null
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 主事务分组列表
	Content []*Transaction `json:"Content,omitempty" name:"Content"`
}

type Transaction struct {
	// 主事务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionId *int64 `json:"TransactionId,omitempty" name:"TransactionId"`

	// 主事务开始时间戳，UTC，精确到毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionBegin *int64 `json:"TransactionBegin,omitempty" name:"TransactionBegin"`

	// 主事务结束时间戳，UTC，精确到毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionEnd *int64 `json:"TransactionEnd,omitempty" name:"TransactionEnd"`

	// 主事务提交时间戳，UTC，精确到毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionCommit *int64 `json:"TransactionCommit,omitempty" name:"TransactionCommit"`

	// 主事务回滚时间戳，UTC，精确到毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionRollback *int64 `json:"TransactionRollback,omitempty" name:"TransactionRollback"`

	// 主事务异常停止时间戳，UTC，精确到毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransactionError *int64 `json:"TransactionError,omitempty" name:"TransactionError"`

	// 主事务超时时长，单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitempty" name:"Timeout"`

	// 主事务状态：0:Trying, 1:Confirming, 2: Confirmed, 3:Canceling, 4: Canceled
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 主事务结束标识：0:运行中, 1: 已结束
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndFlag *int64 `json:"EndFlag,omitempty" name:"EndFlag"`

	// 主事务超时标识：0:运行中, 1: 已超时
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutFlag *int64 `json:"TimeoutFlag,omitempty" name:"TimeoutFlag"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 事务分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 主事务来源服务标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Server *string `json:"Server,omitempty" name:"Server"`

	// 分支事务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BranchQuantity *int64 `json:"BranchQuantity,omitempty" name:"BranchQuantity"`

	// 重试标识：true：可以重试；false：不可重试
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryFlag *bool `json:"RetryFlag,omitempty" name:"RetryFlag"`
}