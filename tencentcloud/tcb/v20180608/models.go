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

package v20180608

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DatabasesInfo struct {

	// 数据库唯一标识
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 状态。包含以下取值：
	// <li>INITIALIZING：资源初始化中</li>
	// <li>RUNNING：运行中，可正常使用的状态</li>
	// <li>UNUSABLE：禁用，不可用</li>
	// <li>OVERDUE：资源过期</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeDatabaseACLRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitempty" name:"CollectionName"`
}

func (r *DescribeDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseACLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 权限标签。取值范围：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
		AclTag *string `json:"AclTag,omitempty" name:"AclTag"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabaseACLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvsRequest struct {
	*tchttp.BaseRequest

	// 环境ID，如果传了这个参数则只返回该环境的相关信息
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`
}

func (r *DescribeEnvsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境信息列表
		EnvList []*EnvInfo `json:"EnvList,omitempty" name:"EnvList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnvInfo struct {

	// 账户下该环境唯一标识
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 环境来源。包含以下取值：
	// <li>miniapp：微信小程序</li>
	// <li>qcloud ：腾讯云</li>
	Source *string `json:"Source,omitempty" name:"Source"`

	// 环境别名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 环境状态。包含以下取值：
	// <li>NORMAL：正常可用</li>
	// <li>HALTED：停服，用量超限或后台封禁</li>
	// <li>UNAVAILABLE：服务不可用，可能是尚未初始化或者初始化过程中</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 数据库列表
	Databases []*DatabasesInfo `json:"Databases,omitempty" name:"Databases" list`

	// 存储列表
	Storages []*StorageInfo `json:"Storages,omitempty" name:"Storages" list`

	// 函数列表
	Functions []*FunctionInfo `json:"Functions,omitempty" name:"Functions" list`

	// tcb产品套餐ID，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageId *string `json:"PackageId,omitempty" name:"PackageId"`

	// 套餐中文名称，参考DescribePackages接口的返回值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageName *string `json:"PackageName,omitempty" name:"PackageName"`
}

type FunctionInfo struct {

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`
}

type ModifyDatabaseACLRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 集合名称
	CollectionName *string `json:"CollectionName,omitempty" name:"CollectionName"`

	// 权限标签。取值范围：
	// <li> READONLY：所有用户可读，仅创建者和管理员可写</li>
	// <li> PRIVATE：仅创建者及管理员可读写</li>
	// <li> ADMINWRITE：所有用户可读，仅管理员可写</li>
	// <li> ADMINONLY：仅管理员可读写</li>
	AclTag *string `json:"AclTag,omitempty" name:"AclTag"`
}

func (r *ModifyDatabaseACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDatabaseACLRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDatabaseACLResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDatabaseACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDatabaseACLResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvRequest struct {
	*tchttp.BaseRequest

	// 环境ID
	EnvId *string `json:"EnvId,omitempty" name:"EnvId"`

	// 环境备注名，要以a-z开头，不能包含 a-zA-z0-9- 以外的字符
	Alias *string `json:"Alias,omitempty" name:"Alias"`
}

func (r *ModifyEnvRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEnvRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEnvResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StorageInfo struct {

	// 资源所属地域。
	// 当前支持ap-shanghai
	Region *string `json:"Region,omitempty" name:"Region"`

	// 桶名，存储资源的唯一标识
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// cdn 域名
	CdnDomain *string `json:"CdnDomain,omitempty" name:"CdnDomain"`

	// 资源所属用户的腾讯云appId
	AppId *string `json:"AppId,omitempty" name:"AppId"`
}
