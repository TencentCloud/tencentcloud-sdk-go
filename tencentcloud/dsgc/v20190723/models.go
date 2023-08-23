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

package v20190723

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type CreateDSPADbMetaResourcesRequestParams struct {
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源类型，支持：cdb（云数据库 MySQL）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbpg（TDSQL-C PostgreSQL版）、cynosdbmysql（TDSQL-C MySQL版）
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 用来标记本次更新是否已经是最后一次，可选值：continue（后续还需要更新）、finished（本次是最后一次更新）。
	UpdateStatus *string `json:"UpdateStatus,omitempty" name:"UpdateStatus"`

	// 本次更新的ID号，用来标记一次完整的更新过程。
	UpdateId *string `json:"UpdateId,omitempty" name:"UpdateId"`

	// 云上资源列表。
	Items []*DspaCloudResourceMeta `json:"Items,omitempty" name:"Items"`
}

type CreateDSPADbMetaResourcesRequest struct {
	*tchttp.BaseRequest
	
	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源类型，支持：cdb（云数据库 MySQL）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbpg（TDSQL-C PostgreSQL版）、cynosdbmysql（TDSQL-C MySQL版）
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 用来标记本次更新是否已经是最后一次，可选值：continue（后续还需要更新）、finished（本次是最后一次更新）。
	UpdateStatus *string `json:"UpdateStatus,omitempty" name:"UpdateStatus"`

	// 本次更新的ID号，用来标记一次完整的更新过程。
	UpdateId *string `json:"UpdateId,omitempty" name:"UpdateId"`

	// 云上资源列表。
	Items []*DspaCloudResourceMeta `json:"Items,omitempty" name:"Items"`
}

func (r *CreateDSPADbMetaResourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPADbMetaResourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DspaId")
	delete(f, "MetaType")
	delete(f, "ResourceRegion")
	delete(f, "UpdateStatus")
	delete(f, "UpdateId")
	delete(f, "Items")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDSPADbMetaResourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDSPADbMetaResourcesResponseParams struct {
	// 本次更新的ID号，用来标记一次完整的更新过程。
	UpdateId *string `json:"UpdateId,omitempty" name:"UpdateId"`

	// 资源类型，支持：cdb（云数据库 MySQL）、dcdb（TDSQL MySQL版）、mariadb（云数据库 MariaDB）、postgres（云数据库 PostgreSQL）、cynosdbpg（TDSQL-C PostgreSQL版）、cynosdbmysql（TDSQL-C MySQL版）
	MetaType *string `json:"MetaType,omitempty" name:"MetaType"`

	// DSPA实例ID。
	DspaId *string `json:"DspaId,omitempty" name:"DspaId"`

	// 资源所处地域。
	ResourceRegion *string `json:"ResourceRegion,omitempty" name:"ResourceRegion"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDSPADbMetaResourcesResponse struct {
	*tchttp.BaseResponse
	Response *CreateDSPADbMetaResourcesResponseParams `json:"Response"`
}

func (r *CreateDSPADbMetaResourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDSPADbMetaResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DspaCloudResourceMeta struct {
	// 用户资源ID。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 资源名称。
	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`

	// 资源VIP。
	ResourceVip *string `json:"ResourceVip,omitempty" name:"ResourceVip"`

	// 资源端口。
	ResourceVPort *uint64 `json:"ResourceVPort,omitempty" name:"ResourceVPort"`

	// 资源被创建时间。
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" name:"ResourceCreateTime"`

	// 用户资源VPC ID 字符串。
	ResourceUniqueVpcId *string `json:"ResourceUniqueVpcId,omitempty" name:"ResourceUniqueVpcId"`

	// 用户资源Subnet ID 字符串。
	ResourceUniqueSubnetId *string `json:"ResourceUniqueSubnetId,omitempty" name:"ResourceUniqueSubnetId"`
}