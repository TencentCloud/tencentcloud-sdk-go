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

package v20181115

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeDomainInfoRequestParams struct {
	// 要查询的域名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 附加字段，是否返回上下文。当为0时不返回上下文，当为1时返回上下文。
	Option *uint64 `json:"Option,omitempty" name:"Option"`
}

type DescribeDomainInfoRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的域名
	Key *string `json:"Key,omitempty" name:"Key"`

	// 附加字段，是否返回上下文。当为0时不返回上下文，当为1时返回上下文。
	Option *uint64 `json:"Option,omitempty" name:"Option"`
}

func (r *DescribeDomainInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "Option")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDomainInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDomainInfoResponseParams struct {
	// 是否有数据，0代表有数据，1代表没有数据
	ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 判定结果，如：black、white、grey
	Result *string `json:"Result,omitempty" name:"Result"`

	// 置信度，取值0-100
	Confidence *uint64 `json:"Confidence,omitempty" name:"Confidence"`

	// 威胁类型。
	// botnet = 僵尸网络
	// trojan = 木马
	// ransomware = 勒索软件
	// worm = 蠕虫
	// dga = 域名生成算法
	// c2 = c&c
	// compromised = 失陷主机
	// dynamicIP = 动态IP
	// proxy = 代理
	// idc = idc 机房
	// whitelist = 白名单
	// tor = 暗网
	// miner = 挖矿
	// maleware site = 恶意站点
	// malware IP = 恶意IP
	// 等等
	ThreatTypes []*string `json:"ThreatTypes,omitempty" name:"ThreatTypes"`

	// 恶意标签，对应的团伙，家族等信息。
	Tags []*TagType `json:"Tags,omitempty" name:"Tags"`

	// 对应的历史上的威胁情报事件
	Intelligences []*IntelligenceType `json:"Intelligences,omitempty" name:"Intelligences"`

	// 情报相关的上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDomainInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDomainInfoResponseParams `json:"Response"`
}

func (r *DescribeDomainInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDomainInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileInfoRequestParams struct {
	// 要查询文件的MD5
	Key *string `json:"Key,omitempty" name:"Key"`

	// 附加字段，是否返回上下文。当为0时不返回上下文，当为1时返回上下文。
	Option *uint64 `json:"Option,omitempty" name:"Option"`
}

type DescribeFileInfoRequest struct {
	*tchttp.BaseRequest
	
	// 要查询文件的MD5
	Key *string `json:"Key,omitempty" name:"Key"`

	// 附加字段，是否返回上下文。当为0时不返回上下文，当为1时返回上下文。
	Option *uint64 `json:"Option,omitempty" name:"Option"`
}

func (r *DescribeFileInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "Option")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFileInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFileInfoResponseParams struct {
	// 是否有数据，0代表有数据，1代表没有数据
	ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 判定结果，如：black、white、grey
	Result *string `json:"Result,omitempty" name:"Result"`

	// 置信度，取值0-100
	Confidence *uint64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文件类型，文件hash
	// （md5,sha1,sha256）,文件大小等等文件
	// 基础信息
	FileInfo []*FileInfoType `json:"FileInfo,omitempty" name:"FileInfo"`

	// 恶意标签，对应的团伙，家族等信息。
	Tags []*TagType `json:"Tags,omitempty" name:"Tags"`

	// 对应的历史上的威胁情报事件
	Intelligences []*IntelligenceType `json:"Intelligences,omitempty" name:"Intelligences"`

	// 情报相关的上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFileInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFileInfoResponseParams `json:"Response"`
}

func (r *DescribeFileInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFileInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpInfoRequestParams struct {
	// 要查询的IP
	Key *string `json:"Key,omitempty" name:"Key"`

	// 附加字段，是否返回上下文。当为0时不返回上下文，当为1时返回上下文。
	Option *uint64 `json:"Option,omitempty" name:"Option"`
}

type DescribeIpInfoRequest struct {
	*tchttp.BaseRequest
	
	// 要查询的IP
	Key *string `json:"Key,omitempty" name:"Key"`

	// 附加字段，是否返回上下文。当为0时不返回上下文，当为1时返回上下文。
	Option *uint64 `json:"Option,omitempty" name:"Option"`
}

func (r *DescribeIpInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "Option")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIpInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIpInfoResponseParams struct {
	// 是否有数据，0代表有数据，1代表没有数据
	ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 判定结果，如：black、white、grey
	Result *string `json:"Result,omitempty" name:"Result"`

	// 置信度，取值0-100
	Confidence *uint64 `json:"Confidence,omitempty" name:"Confidence"`

	// 威胁类型。
	// botnet = 僵尸网络
	// trojan = 木马
	// ransomware = 勒索软件
	// worm = 蠕虫
	// dga = 域名生成算法
	// c2 = c&c
	// compromised = 失陷主机
	// dynamicIP = 动态IP
	// proxy = 代理
	// idc = idc 机房
	// whitelist = 白名单
	// tor = 暗网
	// miner = 挖矿
	// maleware site = 恶意站点
	// malware IP = 恶意IP
	// 等等
	ThreatTypes []*string `json:"ThreatTypes,omitempty" name:"ThreatTypes"`

	// 恶意标签，对应的团伙，家族等信息。
	Tags []*TagType `json:"Tags,omitempty" name:"Tags"`

	// 对应的历史上的威胁情报事件
	Intelligences []*IntelligenceType `json:"Intelligences,omitempty" name:"Intelligences"`

	// 情报相关的上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIpInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIpInfoResponseParams `json:"Response"`
}

func (r *DescribeIpInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIpInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThreatInfoRequestParams struct {
	// 查询对象，域名或IP
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询类型，当前取值为domain或ip
	Type *string `json:"Type,omitempty" name:"Type"`

	// 附加字段，是否返回上下文。当为0时不返回上下文，当为1时返回上下文。
	Option *uint64 `json:"Option,omitempty" name:"Option"`
}

type DescribeThreatInfoRequest struct {
	*tchttp.BaseRequest
	
	// 查询对象，域名或IP
	Key *string `json:"Key,omitempty" name:"Key"`

	// 查询类型，当前取值为domain或ip
	Type *string `json:"Type,omitempty" name:"Type"`

	// 附加字段，是否返回上下文。当为0时不返回上下文，当为1时返回上下文。
	Option *uint64 `json:"Option,omitempty" name:"Option"`
}

func (r *DescribeThreatInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThreatInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "Type")
	delete(f, "Option")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeThreatInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThreatInfoResponseParams struct {
	// 是否有数据，0代表有数据，1代表没有数据
	ReturnCode *uint64 `json:"ReturnCode,omitempty" name:"ReturnCode"`

	// 判定结果，如：black、white、grey
	Result *string `json:"Result,omitempty" name:"Result"`

	// 置信度，取值0-100
	Confidence *uint64 `json:"Confidence,omitempty" name:"Confidence"`

	// 威胁类型。
	// botnet = 僵尸网络
	// trojan = 木马
	// ransomware = 勒索软件
	// worm = 蠕虫
	// dga = 域名生成算法
	// c2 = c&c
	// compromised = 失陷主机
	// dynamicIP = 动态IP
	// proxy = 代理
	// idc = idc 机房
	// whitelist = 白名单
	// tor = 暗网
	// miner = 挖矿
	// maleware site = 恶意站点
	// malware IP = 恶意IP
	// 等等
	ThreatTypes []*string `json:"ThreatTypes,omitempty" name:"ThreatTypes"`

	// 恶意标签，对应的团伙，家族等信息。
	Tags []*string `json:"Tags,omitempty" name:"Tags"`

	// 当前状态
	// active = 活跃
	// sinkholed = sinkholed
	// inactive = 不活跃
	// unknown = 未知
	// expired = 过期
	Status *string `json:"Status,omitempty" name:"Status"`

	// 情报相关的上下文，参数option=1 的时候提供
	// 每个数据默认为3 条
	Context *string `json:"Context,omitempty" name:"Context"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeThreatInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeThreatInfoResponseParams `json:"Response"`
}

func (r *DescribeThreatInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThreatInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FileInfoType struct {
	// 判定渠道
	DetectId *string `json:"DetectId,omitempty" name:"DetectId"`

	// 检测优先级
	DetectPriority *string `json:"DetectPriority,omitempty" name:"DetectPriority"`

	// 引擎优先级
	EnginePriority *string `json:"EnginePriority,omitempty" name:"EnginePriority"`

	// 样本是否存在
	FileExist *string `json:"FileExist,omitempty" name:"FileExist"`

	// 文件上传
	FileForceUpload *string `json:"FileForceUpload,omitempty" name:"FileForceUpload"`

	// 文件大小
	FileSize *string `json:"FileSize,omitempty" name:"FileSize"`

	// 文件上传时间
	FileupTime *string `json:"FileupTime,omitempty" name:"FileupTime"`

	// 病毒文件全名
	FullVirusName *string `json:"FullVirusName,omitempty" name:"FullVirusName"`

	// IDC位置
	IdcPosition *string `json:"IdcPosition,omitempty" name:"IdcPosition"`

	// 文件md5值
	Md5Type *string `json:"Md5Type,omitempty" name:"Md5Type"`

	// PE结构是否存在
	PeExist *string `json:"PeExist,omitempty" name:"PeExist"`

	// PE结构上传
	PeForceUpload *string `json:"PeForceUpload,omitempty" name:"PeForceUpload"`

	// 安全性等级
	SafeLevel *string `json:"SafeLevel,omitempty" name:"SafeLevel"`

	// 扫描时间
	ScanModiTime *string `json:"ScanModiTime,omitempty" name:"ScanModiTime"`

	// 子判定渠道
	SubdetectId *string `json:"SubdetectId,omitempty" name:"SubdetectId"`

	// 病毒名
	UserDefName *string `json:"UserDefName,omitempty" name:"UserDefName"`

	// 病毒类型
	VirusType *string `json:"VirusType,omitempty" name:"VirusType"`

	// 白名单分数
	WhiteScore *string `json:"WhiteScore,omitempty" name:"WhiteScore"`
}

type IntelligenceType struct {
	// 来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// 标记
	Stamp *string `json:"Stamp,omitempty" name:"Stamp"`

	// 时间
	Time *uint64 `json:"Time,omitempty" name:"Time"`
}

type TagType struct {
	// 标签
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 标签对应的中文解释
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}