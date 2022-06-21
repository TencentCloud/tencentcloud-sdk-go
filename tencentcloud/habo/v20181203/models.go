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

package v20181203

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeStatusRequestParams struct {
	// 购买服务后获得的授权帐号，用于保证请求有效性
	Pk *string `json:"Pk,omitempty" name:"Pk"`

	// 需要获取分析结果的样本md5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type DescribeStatusRequest struct {
	*tchttp.BaseRequest
	
	// 购买服务后获得的授权帐号，用于保证请求有效性
	Pk *string `json:"Pk,omitempty" name:"Pk"`

	// 需要获取分析结果的样本md5
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *DescribeStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Pk")
	delete(f, "Md5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStatusResponseParams struct {
	// 接口调用状态，1表示成功，非1表示失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 成功时返回success，失败时返回具体的失败原因，如样本分析未完成
	Info *string `json:"Info,omitempty" name:"Info"`

	// 成功时返回样本日志下载地址，该地址10分钟内有效
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStatusResponseParams `json:"Response"`
}

func (r *DescribeStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAnalyseRequestParams struct {
	// 购买服务后获得的授权帐号，用于保证请求有效性
	Pk *string `json:"Pk,omitempty" name:"Pk"`

	// 样本md5，用于对下载获得的样本完整性进行校验
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 待分析样本下载地址
	DlUrl *string `json:"DlUrl,omitempty" name:"DlUrl"`
}

type StartAnalyseRequest struct {
	*tchttp.BaseRequest
	
	// 购买服务后获得的授权帐号，用于保证请求有效性
	Pk *string `json:"Pk,omitempty" name:"Pk"`

	// 样本md5，用于对下载获得的样本完整性进行校验
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 待分析样本下载地址
	DlUrl *string `json:"DlUrl,omitempty" name:"DlUrl"`
}

func (r *StartAnalyseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAnalyseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Pk")
	delete(f, "Md5")
	delete(f, "DlUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartAnalyseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAnalyseResponseParams struct {
	// 接口调用状态，1表示成功，非1表示失败
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 成功时返回success，失败时返回具体的失败原因
	Info *string `json:"Info,omitempty" name:"Info"`

	// 保留字段
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartAnalyseResponse struct {
	*tchttp.BaseResponse
	Response *StartAnalyseResponseParams `json:"Response"`
}

func (r *StartAnalyseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAnalyseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}