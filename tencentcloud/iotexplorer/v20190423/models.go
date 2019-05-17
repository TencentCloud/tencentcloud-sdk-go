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

package v20190423

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ControlDeviceDataRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 属性数据
	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *ControlDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ControlDeviceDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ControlDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ControlDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ControlDeviceDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Project *ProjectEntry `json:"Project,omitempty" name:"Project"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品分组模板ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 产品类型
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// 加密类型
	EncryptionType *string `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// 连接类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 数据协议
	DataProtocol *int64 `json:"DataProtocol,omitempty" name:"DataProtocol"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitempty" name:"ProductDesc"`

	// 产品的项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品描述
		Product *ProductEntry `json:"Product,omitempty" name:"Product"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DeleteStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceDataHistoryRequest struct {
	*tchttp.BaseRequest

	// 区间开始时间
	MinTime *uint64 `json:"MinTime,omitempty" name:"MinTime"`

	// 区间结束时间
	MaxTime *uint64 `json:"MaxTime,omitempty" name:"MaxTime"`

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`

	// 属性字段名称
	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

	// 返回条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 检索上下文
	Context *string `json:"Context,omitempty" name:"Context"`
}

func (r *DescribeDeviceDataHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceDataHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceDataHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 属性字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		FieldName *string `json:"FieldName,omitempty" name:"FieldName"`

		// 数据是否已全部返回，true 表示数据全部返回，false 表示还有数据待返回，可将 Context 作为入参，继续查询返回结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Listover *bool `json:"Listover,omitempty" name:"Listover"`

		// 检索上下文，当 ListOver 为false时，可以用此上下文，继续读取后续数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		Context *string `json:"Context,omitempty" name:"Context"`

		// 历史数据结果数组，返回对应时间点及取值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Results []*DeviceDataHistoryItem `json:"Results,omitempty" name:"Results" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceDataHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceDataHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceDataRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 设备名称
	DeviceName *string `json:"DeviceName,omitempty" name:"DeviceName"`
}

func (r *DescribeDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 设备数据
		Data *string `json:"Data,omitempty" name:"Data"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDeviceDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModelDefinitionRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeModelDefinitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModelDefinitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品数据模板
		Model *ProductModelDefinition `json:"Model,omitempty" name:"Model"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeModelDefinitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回信息
		Project *ProjectEntryEx `json:"Project,omitempty" name:"Project"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

func (r *DescribeStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品详情
		Product *ProductEntry `json:"Product,omitempty" name:"Product"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeviceDataHistoryItem struct {

	// 时间点，毫秒时间戳
	Time *string `json:"Time,omitempty" name:"Time"`

	// 字段取值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type GetProjectListRequest struct {
	*tchttp.BaseRequest

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 个数限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetProjectListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProjectListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProjectListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目列表
		Projects []*ProjectEntryEx `json:"Projects,omitempty" name:"Projects" list`

		// 列表项个数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProjectListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProjectListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetStudioProductListRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`

	// Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetStudioProductListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetStudioProductListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetStudioProductListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品列表
		Products []*ProductEntry `json:"Products,omitempty" name:"Products" list`

		// 产品数量
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetStudioProductListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetStudioProductListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModelDefinitionRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 数据模板定义
	ModelSchema *string `json:"ModelSchema,omitempty" name:"ModelSchema"`
}

func (r *ModifyModelDefinitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModelDefinitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyModelDefinitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModelDefinitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyModelDefinitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`
}

func (r *ModifyProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProjectRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目详情
		Project *ProjectEntry `json:"Project,omitempty" name:"Project"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyProjectResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitempty" name:"ProductDesc"`

	// 模型ID
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
}

func (r *ModifyStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品描述
		Product *ProductEntry `json:"Product,omitempty" name:"Product"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProductEntry struct {

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 产品分组模板ID
	CategoryId *int64 `json:"CategoryId,omitempty" name:"CategoryId"`

	// 加密类型
	EncryptionType *string `json:"EncryptionType,omitempty" name:"EncryptionType"`

	// 连接类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 数据协议
	DataProtocol *int64 `json:"DataProtocol,omitempty" name:"DataProtocol"`

	// 产品描述
	ProductDesc *string `json:"ProductDesc,omitempty" name:"ProductDesc"`

	// 状态
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`

	// 创建时间
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 区域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 产品类型
	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品ModuleId
	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
}

type ProductModelDefinition struct {

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 模型定义
	ModelDefine *string `json:"ModelDefine,omitempty" name:"ModelDefine"`

	// 更新时间，秒级时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 创建时间，秒级时间戳
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type ProjectEntry struct {

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`

	// 创建时间，unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间，unix时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ProjectEntryEx struct {

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目描述
	ProjectDesc *string `json:"ProjectDesc,omitempty" name:"ProjectDesc"`

	// 项目创建时间，unix时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目更新时间，unix时间戳
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 产品数量
	ProductCount *uint64 `json:"ProductCount,omitempty" name:"ProductCount"`

	// NativeApp数量
	NativeAppCount *uint64 `json:"NativeAppCount,omitempty" name:"NativeAppCount"`

	// WebApp数量
	WebAppCount *uint64 `json:"WebAppCount,omitempty" name:"WebAppCount"`
}

type ReleaseStudioProductRequest struct {
	*tchttp.BaseRequest

	// 产品ID
	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`

	// 产品DevStatus
	DevStatus *string `json:"DevStatus,omitempty" name:"DevStatus"`
}

func (r *ReleaseStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ReleaseStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ReleaseStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStudioProductRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 列表Limit
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 列表Offset
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchStudioProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStudioProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchStudioProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 产品列表
		Products []*ProductEntry `json:"Products,omitempty" name:"Products" list`

		// 产品数量
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchStudioProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchStudioProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
