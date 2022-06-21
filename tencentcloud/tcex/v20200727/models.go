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

package v20200727

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AlgorithmResult struct {
	// 算法ID
	AlgoId *string `json:"AlgoId,omitempty" name:"AlgoId"`

	// 算法名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgoName *string `json:"AlgoName,omitempty" name:"AlgoName"`

	// 算法返回的结果。
	// - 当算法类型为“OCR（1）”时，结果为文本字符串
	// - 当算法类型为“文本分类（2）”时，结果字符串为json对象数组：
	//   Class：分类结果
	//   Confidence：置信度
	// - 算法类型为“情感分析（3）”时，结果字符串为json对象：
	//   Positive：正面情感概率
	//   Negative：负面情感概率
	//   Neutral：中性情感概率
	// - 当算法类型为“合同要素抽取（4）”时，结果字符串为json对象数组：
	//   NodeName：一级要素名称
	//   ItemName：二级要素名称
	//   Content：要素文本内容
	// - 当算法类型为“实体识别（5）”时，结果字符串为json对象数组：
	//   - Entity：实体类型
	//   - Content：实体文本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *string `json:"Result,omitempty" name:"Result"`

	// 算法调用错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Error *string `json:"Error,omitempty" name:"Error"`

	// 算法类型：
	// 1：OCR算法
	// 2：文本分类算法
	// 3：情感分析算法
	// 4：合同要素抽取算法
	// 5、实体识别算法
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlgoType *int64 `json:"AlgoType,omitempty" name:"AlgoType"`
}

// Predefined struct for user
type DescribeInvocationResultRequestParams struct {
	// 调用id，为调用InvokeService接口返回的RequestId
	InvokeId *string `json:"InvokeId,omitempty" name:"InvokeId"`
}

type DescribeInvocationResultRequest struct {
	*tchttp.BaseRequest
	
	// 调用id，为调用InvokeService接口返回的RequestId
	InvokeId *string `json:"InvokeId,omitempty" name:"InvokeId"`
}

func (r *DescribeInvocationResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InvokeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInvocationResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInvocationResultResponseParams struct {
	// 服务的调用结果
	Results []*AlgorithmResult `json:"Results,omitempty" name:"Results"`

	// 0:获取结果失败
	// 1：结果还没有生成，继续轮询
	// 2：获取结果成功
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInvocationResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInvocationResultResponseParams `json:"Response"`
}

func (r *DescribeInvocationResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInvocationResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeServiceRequestParams struct {
	// 待调用的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 要调用服务的状态：0表示调试版本，1表示上线版本
	ServiceStatus *int64 `json:"ServiceStatus,omitempty" name:"ServiceStatus"`

	// 用于测试的文档的URL。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 用于测试的文本，当此值不为空时，调用内容以此参数的值为准。
	Input *string `json:"Input,omitempty" name:"Input"`
}

type InvokeServiceRequest struct {
	*tchttp.BaseRequest
	
	// 待调用的服务ID。
	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`

	// 要调用服务的状态：0表示调试版本，1表示上线版本
	ServiceStatus *int64 `json:"ServiceStatus,omitempty" name:"ServiceStatus"`

	// 用于测试的文档的URL。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 用于测试的文本，当此值不为空时，调用内容以此参数的值为准。
	Input *string `json:"Input,omitempty" name:"Input"`
}

func (r *InvokeServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceId")
	delete(f, "ServiceStatus")
	delete(f, "FileUrl")
	delete(f, "Input")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "InvokeServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type InvokeServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type InvokeServiceResponse struct {
	*tchttp.BaseResponse
	Response *InvokeServiceResponseParams `json:"Response"`
}

func (r *InvokeServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *InvokeServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}