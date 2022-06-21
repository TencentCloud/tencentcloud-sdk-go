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

package v20190118

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type GetLocalEngineRequestParams struct {
	// 购买服务后获得的授权信息，用于保证请求有效性
	Key *string `json:"Key,omitempty" name:"Key"`
}

type GetLocalEngineRequest struct {
	*tchttp.BaseRequest
	
	// 购买服务后获得的授权信息，用于保证请求有效性
	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *GetLocalEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLocalEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLocalEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLocalEngineResponseParams struct {
	// 接口调用状态，成功返回200，失败返回400
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 接口调用描述信息，成功返回"scan success"，失败返回"scan error"
	Info *string `json:"Info,omitempty" name:"Info"`

	// 本地引擎下载地址
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetLocalEngineResponse struct {
	*tchttp.BaseResponse
	Response *GetLocalEngineResponseParams `json:"Response"`
}

func (r *GetLocalEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLocalEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetScanResultRequestParams struct {
	// 购买服务后获得的授权信息，用于保证请求有效性
	Key *string `json:"Key,omitempty" name:"Key"`

	// 需要获取扫描接口的md5（只允许单个md5）
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type GetScanResultRequest struct {
	*tchttp.BaseRequest
	
	// 购买服务后获得的授权信息，用于保证请求有效性
	Key *string `json:"Key,omitempty" name:"Key"`

	// 需要获取扫描接口的md5（只允许单个md5）
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *GetScanResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetScanResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "Md5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetScanResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetScanResultResponseParams struct {
	// 接口调用状态，成功返回200，失败返回400
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 接口调用描述信息，成功返回"scan success"，失败返回"scan error"
	Info *string `json:"Info,omitempty" name:"Info"`

	// 实际结果信息，包括md5、scan_status、virus_name三个字段；virus_name报毒名："torjan.**":黑样本的报毒名、".":样本不报毒、"" :样本无检出信息，需上传扫描；
	// scan_status样本状态：-1无检出信息需上传扫描、0样本扫描中、1样本扫描结束且不报毒、2样本扫描结束且报黑、3样本下载失败；
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetScanResultResponse struct {
	*tchttp.BaseResponse
	Response *GetScanResultResponseParams `json:"Response"`
}

func (r *GetScanResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetScanResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanFileHashRequestParams struct {
	// 购买服务后获得的授权信息，用于保证请求有效性
	Key *string `json:"Key,omitempty" name:"Key"`

	// 需要查询的md5值（支持单个和多个，多个md5间用逗号分格）
	Md5s *string `json:"Md5s,omitempty" name:"Md5s"`

	// 保留字段默认填0
	WithCategory *string `json:"WithCategory,omitempty" name:"WithCategory"`

	// 松严规则控制字段默认填10（5-松、10-标准、15-严）
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" name:"SensitiveLevel"`
}

type ScanFileHashRequest struct {
	*tchttp.BaseRequest
	
	// 购买服务后获得的授权信息，用于保证请求有效性
	Key *string `json:"Key,omitempty" name:"Key"`

	// 需要查询的md5值（支持单个和多个，多个md5间用逗号分格）
	Md5s *string `json:"Md5s,omitempty" name:"Md5s"`

	// 保留字段默认填0
	WithCategory *string `json:"WithCategory,omitempty" name:"WithCategory"`

	// 松严规则控制字段默认填10（5-松、10-标准、15-严）
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" name:"SensitiveLevel"`
}

func (r *ScanFileHashRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanFileHashRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "Md5s")
	delete(f, "WithCategory")
	delete(f, "SensitiveLevel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanFileHashRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanFileHashResponseParams struct {
	// 接口调用状态，成功返回200，失败返回400
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 接口调用描述信息，成功返回"scan success"，失败返回"scan error"
	Info *string `json:"Info,omitempty" name:"Info"`

	// 云查实际结果信息，包括md5、return_state、virus_state、virus_name字符逗号间隔；        
	// return_state查询状态：-1/0代表失败、1/2代表成功；
	// virus_state文状件态：0文件不存在、1白、2黑、3未知、4感染性、5低可信白；
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanFileHashResponse struct {
	*tchttp.BaseResponse
	Response *ScanFileHashResponseParams `json:"Response"`
}

func (r *ScanFileHashResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanFileHashResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanFileRequestParams struct {
	// 购买服务后获得的授权信息，用于保证请求有效性
	Key *string `json:"Key,omitempty" name:"Key"`

	// 文件下载url地址
	Sample *string `json:"Sample,omitempty" name:"Sample"`

	// 文件的md5值
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

type ScanFileRequest struct {
	*tchttp.BaseRequest
	
	// 购买服务后获得的授权信息，用于保证请求有效性
	Key *string `json:"Key,omitempty" name:"Key"`

	// 文件下载url地址
	Sample *string `json:"Sample,omitempty" name:"Sample"`

	// 文件的md5值
	Md5 *string `json:"Md5,omitempty" name:"Md5"`
}

func (r *ScanFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "Sample")
	delete(f, "Md5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanFileResponseParams struct {
	// 接口调用状态，成功返回200，失败返回400
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 接口调用描述信息，成功返回"success"，失败返回"invalid request"
	Info *string `json:"Info,omitempty" name:"Info"`

	// 异步扫描任务提交成功返回success
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanFileResponse struct {
	*tchttp.BaseResponse
	Response *ScanFileResponseParams `json:"Response"`
}

func (r *ScanFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}